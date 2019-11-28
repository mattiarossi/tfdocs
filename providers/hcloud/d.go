package hcloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "hcloud_datacenter",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud Datacenter.`,
			Description: `
Provides details about a specific Hetzner Cloud Datacenter.
Use this resource to get detailed information about specific datacenter.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_datacenters",
			Category:         "Data Sources",
			ShortDescription: `List all available Hetzner Cloud Datacenters.`,
			Description: `
Provides a list of available Hetzner Cloud Datacenters.
This resource may be useful to create highly available infrastructure, distributed across several datacenters.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_floating_ip",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud Floating IP.`,
			Description: `

Provides details about a Hetzner Cloud Floating IP.

This resource can be useful when you need to determine a Floating IP ID based on the IP address.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_image",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud Image.`,
			Description: `
Provides details about a Hetzner Cloud Image.
This resource is useful if you want to use a non-terraform managed image.
`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_location",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud Location.`,
			Description: `
Provides details about a specific Hetzner Cloud Location.
Use this resource to get detailed information about specific location.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_locations",
			Category:         "Data Sources",
			ShortDescription: `List all available Hetzner Cloud Locations.`,
			Description: `
Provides a list of available Hetzner Cloud Locations.
This resource may be useful to create highly available infrastructure, distributed across several locations.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_network",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud network.`,
			Description: `
Provides details about a Hetzner Cloud network.
This resource is useful if you want to use a non-terraform managed network.
`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_server",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud Server.`,
			Description: `
Provides details about a Hetzner Cloud Server.
This resource is useful if you want to use a non-terraform managed server.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_ssh_key",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud SSH Key.`,
			Description: `
Provides details about a Hetzner Cloud SSH Key.
This resource is useful if you want to use a non-terraform managed SSH Key.
`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_ssh_keys",
			Category:         "Data Sources",
			ShortDescription: `Provides details about multiple Hetzner Cloud SSH Keys.`,
			Description: `

Provides details about Hetzner Cloud SSH Keys.
This resource is useful if you want to use a non-terraform managed SSH Key.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "hcloud_volume",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Hetzner Cloud volume.`,
			Description: `
Provides details about a Hetzner Cloud volume.
This resource is useful if you want to use a non-terraform managed volume.
`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"hcloud_datacenter":  0,
		"hcloud_datacenters": 1,
		"hcloud_floating_ip": 2,
		"hcloud_image":       3,
		"hcloud_location":    4,
		"hcloud_locations":   5,
		"hcloud_network":     6,
		"hcloud_server":      7,
		"hcloud_ssh_key":     8,
		"hcloud_ssh_keys":    9,
		"hcloud_volume":      10,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
