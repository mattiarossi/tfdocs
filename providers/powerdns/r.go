package powerdns

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "powerdns_record",
			Category:         "Resources",
			ShortDescription: `Provides a PowerDNS record resource.`,
			Description:      ``,
			Keywords: []string{
				"record",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The name of zone to contain this record.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the record.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The record type.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Required) The TTL of the record.`,
				},
				resource.Attribute{
					Name:        "records",
					Description: `(Required) A string list of records.`,
				},
				resource.Attribute{
					Name:        "set_ptr",
					Description: `(Optional) [`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "powerdns_zone",
			Category:         "Resources",
			ShortDescription: `Provides a PowerDNS zone.`,
			Description:      ``,
			Keywords: []string{
				"zone",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of zone.`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Required) The kind of the zone.`,
				},
				resource.Attribute{
					Name:        "nameservers",
					Description: `(Optional) The zone nameservers.`,
				},
				resource.Attribute{
					Name:        "masters",
					Description: `(Optional) List of IP addresses configured as a master for this zone (“Slave” kind zones only)`,
				},
				resource.Attribute{
					Name:        "soa_edit_api",
					Description: `(Optional) This should map to one of the [supported API values](https://doc.powerdns.com/authoritative/dnsupdate.html#soa-edit-dnsupdate-settings)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"powerdns_record": 0,
		"powerdns_zone":   1,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
