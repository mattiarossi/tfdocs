package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/cycloidio/tfdocs/resource"
)

const (
	iconFile  = "icons.json"
	aliasFile = "aliases.json"
)

var (
	erbEmbed = regexp.MustCompile("<%=?[^>]*>")
)

func main() {
	providersPath := filepath.Join("terraform-website", "ext", "providers")
	assetsPath := filepath.Join("assets")
	aliasesPath := filepath.Join("aliases")
	fileInfos, err := ioutil.ReadDir(providersPath)
	if err != nil {
		panic(err)
	}
	for _, fi := range fileInfos {
		if !fi.IsDir() {
			continue
		}
		provider := fi.Name()
		// avi has all the Resource names as 'ActionGroupConfig' instead of 'avi_actiongroupconfig'
		// as any other provider would have
		if provider == "avi" {
			continue
		}
		var icons map[string]string
		b, err := ioutil.ReadFile(filepath.Join(assetsPath, provider, iconFile))
		if !os.IsNotExist(err) {
			err = json.Unmarshal(b, &icons)
			if err != nil {
				panic(err)
			}
		}

		var aliases = make(map[string]map[string]string)
		b, err = ioutil.ReadFile(filepath.Join(aliasesPath, provider, aliasFile))
		if !os.IsNotExist(err) {
			err = json.Unmarshal(b, &aliases)
			if err != nil {
				panic(err)
			}
		}
		fmt.Println(aliases)

		docsPath := filepath.Join(providersPath, provider, "website", "docs")
		sidebarPath := filepath.Join(providersPath, provider, "website", fmt.Sprintf("%s.erb", provider))
		secondSidebarPath := filepath.Join(providersPath, provider, "website", "layout.erb")

		// categories has as key the resource name and as value the category
		// they belong to
		// The DataSources are not here
		categories := make(map[string]string)

		// The type it's hard to guess.
		// * On the .markdown, the `Resource: ` sometimes has with or without aws_, and some
		// times it does not match
		// * The `page_title` the smae
		// * File name does not match always
		// The only valid solution is to read the sidebar and get the URL which is the file
		// and use the Text to set as type
		fileType := make(map[string]string)

		body, err := ioutil.ReadFile(sidebarPath)
		if err != nil {
			if strings.Contains(err.Error(), "no such file or directory") {
				body, err = ioutil.ReadFile(secondSidebarPath)
				if err != nil {
					panic(err)
				}
			} else {
				panic(err)
			}
		}

		body = erbEmbed.ReplaceAll(body, nil)

		doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(body))
		if err != nil {
			panic(err)
		}

		doc.Find(".nav.docs-sidenav > li").Each(func(_ int, s *goquery.Selection) {
			as := s.Find("a")
			cat := as.First().Text()

			// If they have more than one 'ul' it means that it has
			// the Category as first lvl and then the actual differenciation
			// between Resource or DataSource in 2nd lvl
			ulsLen := s.Find("ul").Length()
			if ulsLen > 1 {
				s.ChildrenFiltered("ul").ChildrenFiltered("li").Each(func(_ int, s *goquery.Selection) {
					as = s.Find("a")
					resourceType := as.First().Text()

					isResource := true
					if strings.Contains(resourceType, "Data Sources") || strings.Contains(resourceType, "Data Resources") {
						isResource = false
					}
					as.Slice(1, goquery.ToEnd).Each(func(i int, s *goquery.Selection) {
						parseATag(s, provider, cat, isResource, fileType, categories)
					})
				})
			} else if ulsLen == 1 {
				as.Slice(1, goquery.ToEnd).Each(func(i int, s *goquery.Selection) {
					isResource := true
					if strings.Contains(cat, "Data Sources") || strings.Contains(cat, "Data Resources") {
						isResource = false
					}
					parseATag(s, provider, cat, isResource, fileType, categories)
				})
			}
		})

		for _, t := range []string{"r", "d"} {
			files, err := ioutil.ReadDir(filepath.Join(docsPath, t))
			if err != nil {
				// This means that the provides does not have 'r' or 'd'
				if strings.Contains(err.Error(), "no such file or directory") {
					continue
				}
				panic(err)
			}

			err = os.MkdirAll(filepath.Join("providers", provider), os.ModePerm)
			if err != nil && !os.IsExist(err) {
				panic(err)
			}
			out, err := os.OpenFile(filepath.Join("providers", provider, fmt.Sprintf("%s.go", t)), os.O_APPEND|os.O_TRUNC|os.O_RDWR|os.O_CREATE, 0644)
			if err != nil {
				panic(err)
			}
			defer out.Close()

			resources := make([]resource.Resource, 0)
			for _, f := range files {
				if f.IsDir() || f.Name() == ".keep" {
					continue
				}
				body, err := ioutil.ReadFile(filepath.Join(docsPath, t, f.Name()))
				if err != nil {
					panic(err)
				}

				fileName := fmt.Sprintf("%s_%s", t, strings.Split(f.Name(), ".")[0])
				rt, ok := fileType[fileName]
				if !ok {
					log.Printf("the provider %q file %q of %q has no type, name used was %q", provider, f.Name(), t, fileName)
					continue
				}

				r := resource.Resource{
					Type:             rt,
					ShortDescription: getShortDescription(body),
					Description:      getDescription(body),
					Arguments:        getArguments(body),
					Attributes:       getAttributes(body),
				}
				if t == "d" {
					r.Category = "Data Sources"
				} else {
					c, ok := categories[rt]
					if !ok {
						log.Printf("the provider %q the resource %q has no category", provider, rt)
						continue
					}
					r.Category = c
					r.Keywords = categoryAndTypeToKeywords(provider, c, rt)
				}
				if icons != nil {
					if ic, ok := icons[rt]; ok {
						r.Icon = ic
					} else {
						icons[rt] = ""
					}
				}
				resources = append(resources, r)

				if nt, ok := aliases[t][rt]; ok {
					r.Type = nt
					resources = append(resources, r)
				}
			}

			buff := &bytes.Buffer{}

			td := TemplateData{
				Resources: resources,
				Type:      t,
				Provider:  provider,
			}

			err = resourceTmpl.Execute(buff, td)
			if err != nil {
				panic(err)
			}

			b, err := format.Source(buff.Bytes())
			if err != nil {
				buff = &bytes.Buffer{}
				log.Printf("ERROR: could not import for provider %q and type %q setting empty", provider, t)
				td.Resources = nil
				err = resourceTmpl.Execute(buff, td)
				if err != nil {
					panic(err)
				}
				b, err = format.Source(buff.Bytes())
				if err != nil {
					panic(err)
				}
			}

			io.Copy(out, bytes.NewBuffer(b))

		}
		if icons != nil {
			f, err := os.OpenFile(filepath.Join(assetsPath, provider, iconFile), os.O_RDWR|os.O_TRUNC, 0755)
			if err != nil {
				log.Fatal(err)
			}
			e := json.NewEncoder(f)
			e.SetEscapeHTML(false)
			e.SetIndent("", " ")
			err = e.Encode(icons)
			if err != nil {
				log.Fatal(err)
			}
			f.Close()
		}

	}
}

// parseATag parses the 'a' tag into a resource and sets it where it should go
func parseATag(s *goquery.Selection, provider, cat string, isResource bool, fileType, categories map[string]string) {
	text := s.Text()
	// Set the href of the link with the actual type of it
	href, _ := s.Attr("href")
	paths := strings.Split(href, "/")
	fileName := strings.Split(paths[len(paths)-1], ".")[0]

	// We want to ignore the 'Data Sources'
	// as they all are the same category
	if !isResource {
		fileType[fmt.Sprintf("d_%s", fileName)] = text
		return
	}

	fileType[fmt.Sprintf("r_%s", fileName)] = text
	if v, ok := categories[text]; ok {
		log.Printf("WARNING: the provider %q with resource %q found with category %q and also on %q", provider, text, v, cat)
		cat += " " + v
	}
	categories[text] = cat
}

//var categoryReplacer = strings.NewReplacer("(", " ", ")", " ", "/", " ")
var categoryReplacer = regexp.MustCompile(`\W`)

func categoryAndTypeToKeywords(provider, c, rt string) []string {
	cws := strings.Split(
		strings.ToLower(categoryReplacer.ReplaceAllString(c, " ")), " ")
	tws := strings.Split(rt, "_")

	res := make([]string, 0, len(cws))
	ws := make(map[string]struct{})
	for _, w := range append(cws, tws...) {
		// The word it's already there
		if _, ok := ws[w]; ok {
			continue
		}
		if w == "resources" || w == provider || w == "" {
			continue
		}
		res = append(res, w)
		ws[w] = struct{}{}
	}

	return res
}

//var resourceTypeRe = regexp.MustCompile(`(?ms)#\s(?:Resource|Data\sSource):\s([\w\d-_]+)\n`)

var resourceTypeRe = regexp.MustCompile(`(?ms)page_title:\s"[\w\d\-_]+:\s([\w\d\-_]+)"`)

func getResourceType(b []byte) string {
	res := resourceTypeRe.FindSubmatch(b)
	if len(res) != 2 {
		return ""
	}
	return standardizeSpaces(string(res[1]))
}

var shortDescriptionRe = regexp.MustCompile(`(?ms)description:\s(?:\|-)?([^#]+)+---`)

func getShortDescription(b []byte) string {
	res := shortDescriptionRe.FindSubmatch(b)
	if len(res) != 2 {
		return ""
	}
	return escapeRawStringQuotes(standardizeSpaces(string(res[1])))
}

var descriptionRe = regexp.MustCompile(`(?ms)#\s(?:Resource|Data\sSource):\s[\w\d-_]+(.+?)##[^#]+`)

func getDescription(b []byte) string {
	res := descriptionRe.FindSubmatch(b)
	if len(res) != 2 {
		return ""
	}
	return escapeRawStringQuotes(string(res[1]))
}

var argumentsBodyRe = regexp.MustCompile(`(?ms)##\sArgument\sReference(.+)(?:##[^#]+|$)+?`)

var argumentRe = regexp.MustCompile("(?ms)\\*\\s`([^`]+)`\\s-\\s([^*]+)")

func getArguments(b []byte) []resource.Attribute {
	res := argumentsBodyRe.FindSubmatch(b)
	if len(res) != 2 {
		return nil
	}
	atts := argumentRe.FindAllSubmatch(res[1], -1)
	attributes := make([]resource.Attribute, 0, len(atts))
	for _, at := range atts {
		attributes = append(attributes, resource.Attribute{
			Name:        standardizeSpaces(string(at[1])),
			Description: escapeRawStringQuotes(standardizeSpaces(string(at[2]))),
		})
	}
	return attributes
}

var attributesRe = regexp.MustCompile(`(?ms)##\sAttributes?\sReference(.+)(?:##[^#]+|$)+?`)

func getAttributes(b []byte) []resource.Attribute {
	res := attributesRe.FindSubmatch(b)
	if len(res) != 2 {
		return nil
	}
	atts := argumentRe.FindAllSubmatch(res[1], -1)
	attributes := make([]resource.Attribute, 0, len(atts))
	for _, at := range atts {
		attributes = append(attributes, resource.Attribute{
			Name:        standardizeSpaces(string(at[1])),
			Description: escapeRawStringQuotes(standardizeSpaces(string(at[2]))),
		})
	}
	return attributes
}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func escapeRawStringQuotes(s string) string {
	return strings.Replace(s, "`", "`+\"`\"+`", -1)
}
