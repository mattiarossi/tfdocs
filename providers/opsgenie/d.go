package opsgenie

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "opsgenie_user",
			Category:         "Data Sources",
			ShortDescription: `Gets information about a specific user within OpsGenie`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The username (email) to use to find a user in OpsGenie. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "full_name",
					Description: `The full name of the found user.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `The role of the found user.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "full_name",
					Description: `The full name of the found user.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `The role of the found user.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"opsgenie_user": 0,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
