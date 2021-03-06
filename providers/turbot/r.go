package turbot

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "turbot_file",
			Category:         "File",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"file",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "turbot_folder",
			Category:         "Folder",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"folder",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "turbot_google_directory",
			Category:         "Google Directory",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"google",
				"directory",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "turbot_grant",
			Category:         "Grant",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"grant",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "turbot_grant_activation",
			Category:         "Grant Activation",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"grant",
				"activation",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "turbot_local_directory",
			Category:         "Local Directory",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"directory",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "turbot_local_directory_user",
			Category:         "Local Directory User",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"local",
				"directory",
				"user",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "turbot_mod",
			Category:         "Mod",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"mod",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "turbot_policy_setting",
			Category:         "Policy Setting",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"policy",
				"setting",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "turbot_profile",
			Category:         "Profile",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"profile",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "turbot_resource",
			Category:         "Resource",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"resource",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "turbot_saml_directory",
			Category:         "SAML Directory",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"saml",
				"directory",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "turbot_shadow_resource",
			Category:         "Shadow Resource",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"shadow",
				"resource",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "turbot_smart_folder",
			Category:         "Smart Folder",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"smart",
				"folder",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "turbot_smart_folder_attachment",
			Category:         "Smart Folder Attachment",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"smart",
				"folder",
				"attachment",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "turbot_turbot_directory",
			Category:         "Turbot Directory",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"directory",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"turbot_file":                    0,
		"turbot_folder":                  1,
		"turbot_google_directory":        2,
		"turbot_grant":                   3,
		"turbot_grant_activation":        4,
		"turbot_local_directory":         5,
		"turbot_local_directory_user":    6,
		"turbot_mod":                     7,
		"turbot_policy_setting":          8,
		"turbot_profile":                 9,
		"turbot_resource":                10,
		"turbot_saml_directory":          11,
		"turbot_shadow_resource":         12,
		"turbot_smart_folder":            13,
		"turbot_smart_folder_attachment": 14,
		"turbot_turbot_directory":        15,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
