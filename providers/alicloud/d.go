package alicloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "alicloud_account",
			Category:         "Data Sources",
			ShortDescription: `Provides information about the current Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Account ID (e.g. "1239306421830812"). It can be used to construct an ARN.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_actiontrails",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of action trail to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results action trail name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of trail names.`,
				},
				resource.Attribute{
					Name:        "actiontrails",
					Description: `A list of actiontrails. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the trail.`,
				},
				resource.Attribute{
					Name:        "event_rw",
					Description: `Indicates whether the event is a read or a write event.`,
				},
				resource.Attribute{
					Name:        "oss_bucket_name",
					Description: `The name of the specified OSS bucket.`,
				},
				resource.Attribute{
					Name:        "oss_key_prefix",
					Description: `The prefix of the specified OSS bucket name.`,
				},
				resource.Attribute{
					Name:        "role_name",
					Description: `The role in ActionTrail.`,
				},
				resource.Attribute{
					Name:        "sls_project_arn",
					Description: `The unique ARN of the Log Service project.`,
				},
				resource.Attribute{
					Name:        "sls_write_role_arn",
					Description: `The unique ARN of the Log Service role.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `A list of trail names.`,
				},
				resource.Attribute{
					Name:        "actiontrails",
					Description: `A list of actiontrails. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the trail.`,
				},
				resource.Attribute{
					Name:        "event_rw",
					Description: `Indicates whether the event is a read or a write event.`,
				},
				resource.Attribute{
					Name:        "oss_bucket_name",
					Description: `The name of the specified OSS bucket.`,
				},
				resource.Attribute{
					Name:        "oss_key_prefix",
					Description: `The prefix of the specified OSS bucket name.`,
				},
				resource.Attribute{
					Name:        "role_name",
					Description: `The role in ActionTrail.`,
				},
				resource.Attribute{
					Name:        "sls_project_arn",
					Description: `The unique ARN of the Log Service project.`,
				},
				resource.Attribute{
					Name:        "sls_write_role_arn",
					Description: `The unique ARN of the Log Service role.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_adb_clusters",
			Category:         "Data Sources",
			ShortDescription: `Provides a collection of ADB clusters according to the specified filters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description_regex",
					Description: `(Optional) A regex string to filter results by cluster description.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of ADB cluster IDs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.68.0+) A mapping of tags to assign to the resource. - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string. - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of ADB cluster IDs.`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `A list of ADB cluster descriptions.`,
				},
				resource.Attribute{
					Name:        "clusters",
					Description: `A list of ADB clusters. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the ADB cluster.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the ADB cluster.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Billing method. Value options: ` + "`" + `PostPaid` + "`" + ` for Pay-As-You-Go and ` + "`" + `PrePaid` + "`" + ` for subscription.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `The DBClusterNetworkType of the ADB cluster.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the cluster belongs to.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `The ZoneId of the ADB cluster.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `Expiration time. Pay-As-You-Go clusters never expire.`,
				},
				resource.Attribute{
					Name:        "expired",
					Description: `The expired of the ADB cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cluster.`,
				},
				resource.Attribute{
					Name:        "lock_mode",
					Description: `The LockMode of the ADB cluster.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The CreateTime of the ADB cluster.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC the cluster belongs to.`,
				},
				resource.Attribute{
					Name:        "db_node_count",
					Description: `The DBNodeCount of the ADB cluster.`,
				},
				resource.Attribute{
					Name:        "db_node_class",
					Description: `The DBNodeClass of the ADB cluster.`,
				},
				resource.Attribute{
					Name:        "db_node_storage",
					Description: `The DBNodeStorage of the ADB cluster.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of ADB cluster IDs.`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `A list of ADB cluster descriptions.`,
				},
				resource.Attribute{
					Name:        "clusters",
					Description: `A list of ADB clusters. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the ADB cluster.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the ADB cluster.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Billing method. Value options: ` + "`" + `PostPaid` + "`" + ` for Pay-As-You-Go and ` + "`" + `PrePaid` + "`" + ` for subscription.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `The DBClusterNetworkType of the ADB cluster.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the cluster belongs to.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `The ZoneId of the ADB cluster.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `Expiration time. Pay-As-You-Go clusters never expire.`,
				},
				resource.Attribute{
					Name:        "expired",
					Description: `The expired of the ADB cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cluster.`,
				},
				resource.Attribute{
					Name:        "lock_mode",
					Description: `The LockMode of the ADB cluster.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The CreateTime of the ADB cluster.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC the cluster belongs to.`,
				},
				resource.Attribute{
					Name:        "db_node_count",
					Description: `The DBNodeCount of the ADB cluster.`,
				},
				resource.Attribute{
					Name:        "db_node_class",
					Description: `The DBNodeClass of the ADB cluster.`,
				},
				resource.Attribute{
					Name:        "db_node_storage",
					Description: `The DBNodeStorage of the ADB cluster.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_adb_zones",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of availability zones for ADB that can be used by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "multi",
					Description: `(Optional) Indicate whether the zones can be used in a multi AZ configuration. Default to ` + "`" + `false` + "`" + `. Multi AZ is usually used to launch ADB instances.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "multi_zone_ids",
					Description: `A list of zone ids in which the multi zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "multi_zone_ids",
					Description: `A list of zone ids in which the multi zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_alidns_domain_groups",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Domain Groups available to the Alidns.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of instance IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by the domain group name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of instance IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of domain group names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the instance.`,
				},
				resource.Attribute{
					Name:        "domain_count",
					Description: `Number of domain names in the group.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `Id of the domain group.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `The name of the domain group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of instance IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of domain group names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the instance.`,
				},
				resource.Attribute{
					Name:        "domain_count",
					Description: `Number of domain names in the group.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `Id of the domain group.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `The name of the domain group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_alidns_records",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of domain records available to the Alidns.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_name",
					Description: `(Required) The domain name associated to the records.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Record type. Valid values: ` + "`" + `A` + "`" + `, ` + "`" + `NS` + "`" + `, ` + "`" + `MX` + "`" + `, ` + "`" + `TXT` + "`" + `, ` + "`" + `CNAME` + "`" + `, ` + "`" + `SRV` + "`" + `, ` + "`" + `AAAA` + "`" + `, ` + "`" + `REDIRECT_URL` + "`" + `, ` + "`" + `FORWORD_URL` + "`" + ` .`,
				},
				resource.Attribute{
					Name:        "line",
					Description: `(Optional) ISP line. Valid values: ` + "`" + `default` + "`" + `, ` + "`" + `telecom` + "`" + `, ` + "`" + `unicom` + "`" + `, ` + "`" + `mobile` + "`" + `, ` + "`" + `oversea` + "`" + `, ` + "`" + `edu` + "`" + `, ` + "`" + `drpeng` + "`" + `, ` + "`" + `btvn` + "`" + `, .etc. For checking all resolution lines enumeration please visit [Alibaba Cloud DNS doc](https://www.alibabacloud.com/help/doc-detail/34339.htm)`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Record status. Valid values: ` + "`" + `ENABLE` + "`" + ` and ` + "`" + `DISABLE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "is_locked",
					Description: `(Optional, type: bool) Whether the record is locked or not.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of record IDs.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Optional) Sorting direction. Valid values: ` + "`" + `DESC` + "`" + `,` + "`" + `ASC` + "`" + `. Default to ` + "`" + `AESC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Optional) Domain name group ID.`,
				},
				resource.Attribute{
					Name:        "key_word",
					Description: `(Optional) Keywords.`,
				},
				resource.Attribute{
					Name:        "lang",
					Description: `(Optional) User language.`,
				},
				resource.Attribute{
					Name:        "order_by",
					Description: `(Optional) Sort by. Sort from newest to oldest according to the time added by resolution.`,
				},
				resource.Attribute{
					Name:        "rr_key_word",
					Description: `(Optional) The keywords recorded by the host are searched according to the ` + "`" + `%RRKeyWord%` + "`" + ` mode, and are not case sensitive.`,
				},
				resource.Attribute{
					Name:        "search_mode",
					Description: `(Optional) Search mode, Valid values: ` + "`" + `LIKE` + "`" + `, ` + "`" + `EXACT` + "`" + `, ` + "`" + `ADVANCED` + "`" + `, ` + "`" + `LIKE` + "`" + ` (fuzzy), ` + "`" + `EXACT` + "`" + ` (accurate) search supports KeyWord field, ` + "`" + `ADVANCED` + "`" + ` (advanced) mode supports other fields.`,
				},
				resource.Attribute{
					Name:        "type_key_word",
					Description: `(Optional) Analyze type keywords, search by full match, not case sensitive.`,
				},
				resource.Attribute{
					Name:        "value_key_word",
					Description: `(Optional) The keywords of the recorded value are searched according to the ` + "`" + `%ValueKeyWord%` + "`" + ` mode, and are not case sensitive.`,
				},
				resource.Attribute{
					Name:        "rr_regex",
					Description: `(Optional) Host record regex.`,
				},
				resource.Attribute{
					Name:        "value_regex",
					Description: `(Optional) Host record value regex.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of record IDs.`,
				},
				resource.Attribute{
					Name:        "records",
					Description: `A list of records. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "record_id",
					Description: `ID of the record.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `Name of the domain record belongs to.`,
				},
				resource.Attribute{
					Name:        "rr",
					Description: `Host record of the domain.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Host record value of the domain.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the record.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `TTL of the record.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Priority of the ` + "`" + `MX` + "`" + ` record.`,
				},
				resource.Attribute{
					Name:        "line",
					Description: `ISP line of the record.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the record.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Indicates whether the record is locked.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of record IDs.`,
				},
				resource.Attribute{
					Name:        "records",
					Description: `A list of records. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the resource.`,
				},
				resource.Attribute{
					Name:        "record_id",
					Description: `ID of the record.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `Name of the domain record belongs to.`,
				},
				resource.Attribute{
					Name:        "rr",
					Description: `Host record of the domain.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Host record value of the domain.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the record.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `TTL of the record.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Priority of the ` + "`" + `MX` + "`" + ` record.`,
				},
				resource.Attribute{
					Name:        "line",
					Description: `ISP line of the record.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the record.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Indicates whether the record is locked.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_alikafka_consumer_groups",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of alikafka consumer groups available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) ID of the ALIKAFKA Instance that owns the consumer groups.`,
				},
				resource.Attribute{
					Name:        "consumer_id_regex",
					Description: `(Optional) A regex string to filter results by the consumer group id.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "consumer_ids",
					Description: `A list of consumer group ids.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "consumer_ids",
					Description: `A list of consumer group ids.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_alikafka_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of alikafka instances available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of instance IDs to filter results.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by the instance name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of instance IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of instance names.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the instance.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of the instance.`,
				},
				resource.Attribute{
					Name:        "service_status",
					Description: `The current status of the instance. -1: unknown status, 0: wait deploy, 1: initializing, 2: preparing, 3 starting, 5: in service, 7: wait upgrade, 8: upgrading, 10: released, 15: freeze, 101: deploy error, 102: upgrade error.`,
				},
				resource.Attribute{
					Name:        "deploy_type",
					Description: `The deploy type of the instance. 0: sharing instance, 1: vpc instance, 2: vpc instance(support ip mapping), 3: eip instance, 4: eip/vpc instance, 5: vpc instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of attaching VPC to instance.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `The ID of attaching vswitch to instance.`,
				},
				resource.Attribute{
					Name:        "io_max",
					Description: `The peak value of io of the instance.`,
				},
				resource.Attribute{
					Name:        "eip_max",
					Description: `The peak bandwidth of the instance.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `The disk type of the instance. 0: efficient cloud disk , 1: SSD.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `The disk size of the instance.`,
				},
				resource.Attribute{
					Name:        "topic_quota",
					Description: `The max num of topic can be create of the instance.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `The ID of attaching zone to instance.`,
				},
				resource.Attribute{
					Name:        "paid_type",
					Description: `The paid type of the instance.`,
				},
				resource.Attribute{
					Name:        "spec_type",
					Description: `The spec type of the instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of instance IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of instance names.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the instance.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of the instance.`,
				},
				resource.Attribute{
					Name:        "service_status",
					Description: `The current status of the instance. -1: unknown status, 0: wait deploy, 1: initializing, 2: preparing, 3 starting, 5: in service, 7: wait upgrade, 8: upgrading, 10: released, 15: freeze, 101: deploy error, 102: upgrade error.`,
				},
				resource.Attribute{
					Name:        "deploy_type",
					Description: `The deploy type of the instance. 0: sharing instance, 1: vpc instance, 2: vpc instance(support ip mapping), 3: eip instance, 4: eip/vpc instance, 5: vpc instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of attaching VPC to instance.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `The ID of attaching vswitch to instance.`,
				},
				resource.Attribute{
					Name:        "io_max",
					Description: `The peak value of io of the instance.`,
				},
				resource.Attribute{
					Name:        "eip_max",
					Description: `The peak bandwidth of the instance.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `The disk type of the instance. 0: efficient cloud disk , 1: SSD.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `The disk size of the instance.`,
				},
				resource.Attribute{
					Name:        "topic_quota",
					Description: `The max num of topic can be create of the instance.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `The ID of attaching zone to instance.`,
				},
				resource.Attribute{
					Name:        "paid_type",
					Description: `The paid type of the instance.`,
				},
				resource.Attribute{
					Name:        "spec_type",
					Description: `The spec type of the instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_alikafka_sasl_acls",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of alikafka sasl acls available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) ID of the ALIKAFKA Instance that owns the sasl acls.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Get results for the specified username.`,
				},
				resource.Attribute{
					Name:        "acl_resource_type",
					Description: `(Required) Get results for the specified resource type.`,
				},
				resource.Attribute{
					Name:        "acl_resource_name",
					Description: `(Required) Get results for the specified resource name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "acls",
					Description: `A list of sasl acls. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username of the sasl acl.`,
				},
				resource.Attribute{
					Name:        "acl_resource_type",
					Description: `The resource type of the sasl acl.`,
				},
				resource.Attribute{
					Name:        "acl_resource_name",
					Description: `The resource name of the sasl acl.`,
				},
				resource.Attribute{
					Name:        "acl_resource_pattern_type",
					Description: `The resource pattern type of the sasl acl.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The host of the sasl acl.`,
				},
				resource.Attribute{
					Name:        "acl_operation_type",
					Description: `The operation type of the sasl acl.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "acls",
					Description: `A list of sasl acls. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username of the sasl acl.`,
				},
				resource.Attribute{
					Name:        "acl_resource_type",
					Description: `The resource type of the sasl acl.`,
				},
				resource.Attribute{
					Name:        "acl_resource_name",
					Description: `The resource name of the sasl acl.`,
				},
				resource.Attribute{
					Name:        "acl_resource_pattern_type",
					Description: `The resource pattern type of the sasl acl.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The host of the sasl acl.`,
				},
				resource.Attribute{
					Name:        "acl_operation_type",
					Description: `The operation type of the sasl acl.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_alikafka_sasl_users",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of alikafka sasl users available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) ID of the ALIKAFKA Instance that owns the sasl users.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by the username.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of sasl usernames.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `A list of sasl users. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username of the user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The password of the user.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `A list of sasl usernames.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `A list of sasl users. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username of the user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The password of the user.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_alikafka_topics",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of alikafka topics available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by the topic name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of topic names.`,
				},
				resource.Attribute{
					Name:        "topics",
					Description: `A list of topics. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "topic",
					Description: `The name of the topic.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Time of creation.`,
				},
				resource.Attribute{
					Name:        "local_topic",
					Description: `whether the current topic is kafka local topic or not.`,
				},
				resource.Attribute{
					Name:        "compact_topic",
					Description: `whether the current topic is kafka compact topic or not.`,
				},
				resource.Attribute{
					Name:        "partition_num",
					Description: `Partition number of the topic.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Remark of the topic.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status code of the topic. There are three values to describe the topic status: 0 stands for the topic is in service, 1 stands for freezing and 2 stands for pause.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `A list of topic names.`,
				},
				resource.Attribute{
					Name:        "topics",
					Description: `A list of topics. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "topic",
					Description: `The name of the topic.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Time of creation.`,
				},
				resource.Attribute{
					Name:        "local_topic",
					Description: `whether the current topic is kafka local topic or not.`,
				},
				resource.Attribute{
					Name:        "compact_topic",
					Description: `whether the current topic is kafka compact topic or not.`,
				},
				resource.Attribute{
					Name:        "partition_num",
					Description: `Partition number of the topic.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Remark of the topic.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status code of the topic. There are three values to describe the topic status: 0 stands for the topic is in service, 1 stands for freezing and 2 stands for pause.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_api_gateway_apis",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of apis to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api_id",
					Description: `(Deprecated, Optional) (It has been deprecated from version 1.52.2, and use field 'ids' to replace.) ID of the specified API.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Optional) ID of the specified group.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter api gateway apis by name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, Available 1.52.2+) A list of api IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of api IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of api names.`,
				},
				resource.Attribute{
					Name:        "apis",
					Description: `A list of apis. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `API ID, which is generated by the system and globally unique.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `API name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `API description.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The ID of the region where the API is located.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `The group id that the apis belong to.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `The group name that the apis belong to.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of api IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of api names.`,
				},
				resource.Attribute{
					Name:        "apis",
					Description: `A list of apis. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `API ID, which is generated by the system and globally unique.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `API name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `API description.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The ID of the region where the API is located.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `The group id that the apis belong to.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `The group name that the apis belong to.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_api_gateway_apps",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of apps to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter apps by name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, Available in 1.52.2+) A list of app IDs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.55.3+) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of app IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of app names.`,
				},
				resource.Attribute{
					Name:        "apps",
					Description: `A list of apps. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `App ID, which is generated by the system and globally unique.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `App name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `App description.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Creation time (Greenwich mean time).`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `Last modification time (Greenwich mean time).`,
				},
				resource.Attribute{
					Name:        "app_code",
					Description: `App code.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of app IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of app names.`,
				},
				resource.Attribute{
					Name:        "apps",
					Description: `A list of apps. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `App ID, which is generated by the system and globally unique.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `App name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `App description.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Creation time (Greenwich mean time).`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `Last modification time (Greenwich mean time).`,
				},
				resource.Attribute{
					Name:        "app_code",
					Description: `App code.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_api_gateway_groups",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of api groups to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter api gateway groups by name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, Available 1.52.1+) A list of api group IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of api group IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of api group names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of api groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `API group ID, which is generated by the system and globally unique.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `API group name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `API group description.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The ID of the region where the API group is located.`,
				},
				resource.Attribute{
					Name:        "sub_domain",
					Description: `Second-level domain name automatically assigned to the API group.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Creation time (Greenwich mean time).`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `Last modification time (Greenwich mean time).`,
				},
				resource.Attribute{
					Name:        "traffic_limit",
					Description: `Upper QPS limit of the API group; default value: 500, which can be increased by submitting an application.`,
				},
				resource.Attribute{
					Name:        "billing_status",
					Description: `Billing status. - NORMAL: The API group is normal. - LOCKED: Locked due to outstanding payment.`,
				},
				resource.Attribute{
					Name:        "illegal_status",
					Description: `Locking in invalid state. - NORMAL: The API group is normal. - LOCKED: Locked due to illegality.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of api group IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of api group names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of api groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `API group ID, which is generated by the system and globally unique.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `API group name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `API group description.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The ID of the region where the API group is located.`,
				},
				resource.Attribute{
					Name:        "sub_domain",
					Description: `Second-level domain name automatically assigned to the API group.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Creation time (Greenwich mean time).`,
				},
				resource.Attribute{
					Name:        "modified_time",
					Description: `Last modification time (Greenwich mean time).`,
				},
				resource.Attribute{
					Name:        "traffic_limit",
					Description: `Upper QPS limit of the API group; default value: 500, which can be increased by submitting an application.`,
				},
				resource.Attribute{
					Name:        "billing_status",
					Description: `Billing status. - NORMAL: The API group is normal. - LOCKED: Locked due to outstanding payment.`,
				},
				resource.Attribute{
					Name:        "illegal_status",
					Description: `Locking in invalid state. - NORMAL: The API group is normal. - LOCKED: Locked due to illegality.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_caller_identity",
			Category:         "Data Sources",
			ShortDescription: `Provides the identity of the current Alibaba Cloud user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Principal ID.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The Alibaba Cloud Resource Name (ARN) of the user making the call.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Account ID.`,
				},
				resource.Attribute{
					Name:        "identity_type",
					Description: `The type of the princiapal. RAMUser for users.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cas_certificates",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of certs available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by the certificate name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, Available in 1.52.0+) A list of cert IDs. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of cert IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of cert names.`,
				},
				resource.Attribute{
					Name:        "certificates",
					Description: `A list of apis. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The cert's id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The cert's name.`,
				},
				resource.Attribute{
					Name:        "common",
					Description: `The cert's common name.`,
				},
				resource.Attribute{
					Name:        "finger_print",
					Description: `The cert's finger.`,
				},
				resource.Attribute{
					Name:        "issuer",
					Description: `The cert's .`,
				},
				resource.Attribute{
					Name:        "org_name",
					Description: `The cert's organization.`,
				},
				resource.Attribute{
					Name:        "province",
					Description: `The cert's province.`,
				},
				resource.Attribute{
					Name:        "city",
					Description: `The cert's city.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `The cert's country.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `The cert's not valid before time.`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `The cert's not valid after time.`,
				},
				resource.Attribute{
					Name:        "sans",
					Description: `The cert's subject alternative name.`,
				},
				resource.Attribute{
					Name:        "expired",
					Description: `The cert is expired or not.`,
				},
				resource.Attribute{
					Name:        "buy_in_aliyun",
					Description: `The cert is buy from aliyun or not.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of cert IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of cert names.`,
				},
				resource.Attribute{
					Name:        "certificates",
					Description: `A list of apis. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The cert's id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The cert's name.`,
				},
				resource.Attribute{
					Name:        "common",
					Description: `The cert's common name.`,
				},
				resource.Attribute{
					Name:        "finger_print",
					Description: `The cert's finger.`,
				},
				resource.Attribute{
					Name:        "issuer",
					Description: `The cert's .`,
				},
				resource.Attribute{
					Name:        "org_name",
					Description: `The cert's organization.`,
				},
				resource.Attribute{
					Name:        "province",
					Description: `The cert's province.`,
				},
				resource.Attribute{
					Name:        "city",
					Description: `The cert's city.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `The cert's country.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `The cert's not valid before time.`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `The cert's not valid after time.`,
				},
				resource.Attribute{
					Name:        "sans",
					Description: `The cert's subject alternative name.`,
				},
				resource.Attribute{
					Name:        "expired",
					Description: `The cert is expired or not.`,
				},
				resource.Attribute{
					Name:        "buy_in_aliyun",
					Description: `The cert is buy from aliyun or not.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cassandra_clusters",
			Category:         "Data Sources",
			ShortDescription: `Provides a collection of Cassandra clusters according to the specified filters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to apply to the cluster name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) The list of Cassandra cluster ids.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in 1.73.0) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) The name of file that can save the collection of clusters after running ` + "`" + `terraform plan` + "`" + `. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `The list of Cassandra cluster ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `The name list of Cassandra clusters.`,
				},
				resource.Attribute{
					Name:        "clusters",
					Description: `A list of Cassandra clusters. Its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Cassandra cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `The ID of the Cassandra cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `The name of the Cassandra cluster.`,
				},
				resource.Attribute{
					Name:        "major_version",
					Description: `The major version of the cluster.`,
				},
				resource.Attribute{
					Name:        "minor_version",
					Description: `The minor version of the cluster.`,
				},
				resource.Attribute{
					Name:        "lock_mode",
					Description: `The lock mode of the cluster.`,
				},
				resource.Attribute{
					Name:        "data_center_count",
					Description: `The count of data centers`,
				},
				resource.Attribute{
					Name:        "pay_type",
					Description: `Billing method. Value options are ` + "`" + `Subscription` + "`" + ` for Pay-As-You-Go and ` + "`" + `PayAsYouGo` + "`" + ` for yearly or monthly subscription.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cluster.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of the cluster.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expire time of the cluster.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `The list of Cassandra cluster ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `The name list of Cassandra clusters.`,
				},
				resource.Attribute{
					Name:        "clusters",
					Description: `A list of Cassandra clusters. Its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Cassandra cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `The ID of the Cassandra cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `The name of the Cassandra cluster.`,
				},
				resource.Attribute{
					Name:        "major_version",
					Description: `The major version of the cluster.`,
				},
				resource.Attribute{
					Name:        "minor_version",
					Description: `The minor version of the cluster.`,
				},
				resource.Attribute{
					Name:        "lock_mode",
					Description: `The lock mode of the cluster.`,
				},
				resource.Attribute{
					Name:        "data_center_count",
					Description: `The count of data centers`,
				},
				resource.Attribute{
					Name:        "pay_type",
					Description: `Billing method. Value options are ` + "`" + `Subscription` + "`" + ` for Pay-As-You-Go and ` + "`" + `PayAsYouGo` + "`" + ` for yearly or monthly subscription.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cluster.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of the cluster.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expire time of the cluster.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cassandra_data_centers",
			Category:         "Data Sources",
			ShortDescription: `Provides a collection of Cassandra Data Centers according to the specified filters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to apply to the cluster name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) The list of Cassandra data center ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `(Optional) The name list of Cassandra data centers.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The cluster id of dataCenters belongs to.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) The name of file that can save the collection of data centers after running ` + "`" + `terraform plan` + "`" + `. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `The list of Cassandra data center ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `The name list of Cassandra data centers.`,
				},
				resource.Attribute{
					Name:        "centers",
					Description: `A list of Cassandra data centers. Its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `The ID of the Cassandra cluster.`,
				},
				resource.Attribute{
					Name:        "commodity_instance",
					Description: `The commodity ID of the Cassandra dataCenter.`,
				},
				resource.Attribute{
					Name:        "data_center_id",
					Description: `The id of the Cassandra dataCenter.`,
				},
				resource.Attribute{
					Name:        "data_center_name",
					Description: `The name of the Cassandra dataCenter.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `One node disk size, unit:GB.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Cloud_ssd or cloud_efficiency.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The instance type of the Cassandra dataCenter, eg: cassandra.c.large.`,
				},
				resource.Attribute{
					Name:        "lock_mode",
					Description: `The lock mode of the dataCenter.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `The node count of dataCenter.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the dataCenter.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of the dataCenter.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expire time of the dataCenter.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `Zone ID the dataCenter belongs to.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC ID the dataCenter belongs to.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `VSwitch ID the dataCenter belongs to.`,
				},
				resource.Attribute{
					Name:        "pay_type",
					Description: `Billing method. Value options are ` + "`" + `Subscription` + "`" + ` for Pay-As-You-Go and ` + "`" + `PayAsYouGo` + "`" + ` for yearly or monthly subscription.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `The list of Cassandra data center ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `The name list of Cassandra data centers.`,
				},
				resource.Attribute{
					Name:        "centers",
					Description: `A list of Cassandra data centers. Its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `The ID of the Cassandra cluster.`,
				},
				resource.Attribute{
					Name:        "commodity_instance",
					Description: `The commodity ID of the Cassandra dataCenter.`,
				},
				resource.Attribute{
					Name:        "data_center_id",
					Description: `The id of the Cassandra dataCenter.`,
				},
				resource.Attribute{
					Name:        "data_center_name",
					Description: `The name of the Cassandra dataCenter.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `One node disk size, unit:GB.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Cloud_ssd or cloud_efficiency.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The instance type of the Cassandra dataCenter, eg: cassandra.c.large.`,
				},
				resource.Attribute{
					Name:        "lock_mode",
					Description: `The lock mode of the dataCenter.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `The node count of dataCenter.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the dataCenter.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of the dataCenter.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The expire time of the dataCenter.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `Zone ID the dataCenter belongs to.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC ID the dataCenter belongs to.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `VSwitch ID the dataCenter belongs to.`,
				},
				resource.Attribute{
					Name:        "pay_type",
					Description: `Billing method. Value options are ` + "`" + `Subscription` + "`" + ` for Pay-As-You-Go and ` + "`" + `PayAsYouGo` + "`" + ` for yearly or monthly subscription.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cassandra_zones",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of availability zones for Cassandra that can be used by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "multi",
					Description: `(Optional) Indicate whether the zones can be used in a multi AZ configuration. Default to ` + "`" + `false` + "`" + `. Multi AZ is usually used to launch Cassandra clusters.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "multi_zone_ids",
					Description: `A list of zone ids in which the multi zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "multi_zone_ids",
					Description: `A list of zone ids in which the multi zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cen_bandwidth_limits",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of CEN Bandwidth Limits owned by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_ids",
					Description: `(Optional) A list of CEN instances IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "limits",
					Description: `A list of CEN Bandwidth Limits. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the CEN instance.`,
				},
				resource.Attribute{
					Name:        "local_region_id",
					Description: `ID of local region.`,
				},
				resource.Attribute{
					Name:        "opposite_region_id",
					Description: `ID of opposite region.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the CEN Bandwidth Limit, including "Active" and "Modifying".`,
				},
				resource.Attribute{
					Name:        "bandwidth_limit",
					Description: `The bandwidth limit configured for the interconnected regions communication.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "limits",
					Description: `A list of CEN Bandwidth Limits. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the CEN instance.`,
				},
				resource.Attribute{
					Name:        "local_region_id",
					Description: `ID of local region.`,
				},
				resource.Attribute{
					Name:        "opposite_region_id",
					Description: `ID of opposite region.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the CEN Bandwidth Limit, including "Active" and "Modifying".`,
				},
				resource.Attribute{
					Name:        "bandwidth_limit",
					Description: `The bandwidth limit configured for the interconnected regions communication.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cen_bandwidth_packages",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of CEN Bandwidth Packages owned by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) ID of a CEN instance.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) Limit search to a list of specific CEN Bandwidth Package IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter CEN Bandwidth Package by name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "packages",
					Description: `A list of CEN bandwidth package. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the CEN Bandwidth Package.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of CEN instance that owns the CEN Bandwidth Package.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the CEN Bandwidth Package.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the CEN Bandwidth Package.`,
				},
				resource.Attribute{
					Name:        "business_status",
					Description: `Status of the CEN Bandwidth Package, including "Normal", "FinancialLocked" and "SecurityLocked".`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the CEN Bandwidth Package in CEN instance, including "Idle" and "InUse".`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The bandwidth in Mbps of the CEN bandwidth package.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Creation time of the CEN bandwidth package.`,
				},
				resource.Attribute{
					Name:        "bandwidth_package_charge_type",
					Description: `The billing method, including "POSTPAY" and "PREPAY".`,
				},
				resource.Attribute{
					Name:        "geographic_region_a_id",
					Description: `Region ID of the interconnected regions.`,
				},
				resource.Attribute{
					Name:        "geographic_region_b_id",
					Description: `Region ID of the interconnected regions.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "packages",
					Description: `A list of CEN bandwidth package. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the CEN Bandwidth Package.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of CEN instance that owns the CEN Bandwidth Package.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the CEN Bandwidth Package.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the CEN Bandwidth Package.`,
				},
				resource.Attribute{
					Name:        "business_status",
					Description: `Status of the CEN Bandwidth Package, including "Normal", "FinancialLocked" and "SecurityLocked".`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the CEN Bandwidth Package in CEN instance, including "Idle" and "InUse".`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The bandwidth in Mbps of the CEN bandwidth package.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Creation time of the CEN bandwidth package.`,
				},
				resource.Attribute{
					Name:        "bandwidth_package_charge_type",
					Description: `The billing method, including "POSTPAY" and "PREPAY".`,
				},
				resource.Attribute{
					Name:        "geographic_region_a_id",
					Description: `Region ID of the interconnected regions.`,
				},
				resource.Attribute{
					Name:        "geographic_region_b_id",
					Description: `Region ID of the interconnected regions.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cen_flowlogs",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of CEN flow logs owned by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of CEN flow log IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter CEN flow logs by name.`,
				},
				resource.Attribute{
					Name:        "cen_id",
					Description: `(Optional) The ID of the CEN Instance.`,
				},
				resource.Attribute{
					Name:        "project_name",
					Description: `(Optional) The name of the SLS project.`,
				},
				resource.Attribute{
					Name:        "log_store_name",
					Description: `(Optional) The name of the log store which is in the ` + "`" + `project_name` + "`" + ` SLS project.`,
				},
				resource.Attribute{
					Name:        "flow_log_name",
					Description: `(Optional) The name of flowlog.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of flowlog.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of flowlog. Valid values: ["Active", "Inactive"]. Default to "Active".`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of CEN flow log IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of CEN flow log names.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of CEN flow logs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the CEN flow log.`,
				},
				resource.Attribute{
					Name:        "flow_log_id",
					Description: `ID of the CEN flow log.`,
				},
				resource.Attribute{
					Name:        "cen_id",
					Description: `The ID of the CEN Instance.`,
				},
				resource.Attribute{
					Name:        "project_name",
					Description: `The name of the SLS project.`,
				},
				resource.Attribute{
					Name:        "log_store_name",
					Description: `The name of the log store which is in the ` + "`" + `project_name` + "`" + ` SLS project.`,
				},
				resource.Attribute{
					Name:        "flow_log_name",
					Description: `The name of flowlog.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of flowlog.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of flowlog.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of CEN flow log IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of CEN flow log names.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of CEN flow logs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the CEN flow log.`,
				},
				resource.Attribute{
					Name:        "flow_log_id",
					Description: `ID of the CEN flow log.`,
				},
				resource.Attribute{
					Name:        "cen_id",
					Description: `The ID of the CEN Instance.`,
				},
				resource.Attribute{
					Name:        "project_name",
					Description: `The name of the SLS project.`,
				},
				resource.Attribute{
					Name:        "log_store_name",
					Description: `The name of the log store which is in the ` + "`" + `project_name` + "`" + ` SLS project.`,
				},
				resource.Attribute{
					Name:        "flow_log_name",
					Description: `The name of flowlog.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of flowlog.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of flowlog.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cen_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of CEN(Cloud Enterprise Network) instances owned by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of CEN instances IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter CEN instances by name.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.81.0+) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of CEN instances IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of CEN instances names.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of CEN instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "cen_bandwidth_package_ids",
					Description: `List of CEN Bandwidth Package IDs in the specified CEN instance.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the CEN instance.`,
				},
				resource.Attribute{
					Name:        "cen_id",
					Description: `ID of the CEN instance.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the CEN instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the CEN instance.`,
				},
				resource.Attribute{
					Name:        "protection_level",
					Description: `Indicates the allowed level of CIDR block overlapping.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the CEN instance, including "Creating", "Active" and "Deleting".`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the Cen Instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of CEN instances IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of CEN instances names.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of CEN instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "cen_bandwidth_package_ids",
					Description: `List of CEN Bandwidth Package IDs in the specified CEN instance.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the CEN instance.`,
				},
				resource.Attribute{
					Name:        "cen_id",
					Description: `ID of the CEN instance.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the CEN instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the CEN instance.`,
				},
				resource.Attribute{
					Name:        "protection_level",
					Description: `Indicates the allowed level of CIDR block overlapping.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the CEN instance, including "Creating", "Active" and "Deleting".`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the Cen Instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cen_private_zones",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of CEN(Cloud Enterprise Network) Private Zones owned by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cen_id",
					Description: `(Required) The ID of the CEN instance.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of CEN private zone IDs.`,
				},
				resource.Attribute{
					Name:        "host_region_id",
					Description: `(Optional) The service region is the target region of the PrivateZone service accessed through CEN.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the PrivateZone service, including ` + "`" + `Creating` + "`" + `, ` + "`" + `Active` + "`" + ` and ` + "`" + `Deleting` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of CEN private zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of CEN private zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the private zone.`,
				},
				resource.Attribute{
					Name:        "cen_id",
					Description: `The ID of the CEN instance.`,
				},
				resource.Attribute{
					Name:        "private_zone_dns_servers",
					Description: `The DNS IP addresses of the PrivateZone service.`,
				},
				resource.Attribute{
					Name:        "access_region_id",
					Description: `The access region. The access region is the region of the cloud resource that accesses the PrivateZone service through CEN.`,
				},
				resource.Attribute{
					Name:        "host_region_id",
					Description: `The service region. The service region is the target region of the PrivateZone service accessed through CEN.`,
				},
				resource.Attribute{
					Name:        "host_vpc_id",
					Description: `The VPC that belongs to the service region.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the PrivateZone service.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of CEN private zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of CEN private zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the private zone.`,
				},
				resource.Attribute{
					Name:        "cen_id",
					Description: `The ID of the CEN instance.`,
				},
				resource.Attribute{
					Name:        "private_zone_dns_servers",
					Description: `The DNS IP addresses of the PrivateZone service.`,
				},
				resource.Attribute{
					Name:        "access_region_id",
					Description: `The access region. The access region is the region of the cloud resource that accesses the PrivateZone service through CEN.`,
				},
				resource.Attribute{
					Name:        "host_region_id",
					Description: `The service region. The service region is the target region of the PrivateZone service accessed through CEN.`,
				},
				resource.Attribute{
					Name:        "host_vpc_id",
					Description: `The VPC that belongs to the service region.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the PrivateZone service.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cen_region_route_entries",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of CEN Route Entries from specific region owned by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) ID of the CEN instance.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `(Required) ID of the region. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "entries",
					Description: `A list of CEN Route Entries. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The destination CIDR block of the route entry.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the route entry.`,
				},
				resource.Attribute{
					Name:        "next_hop_id",
					Description: `ID of the next hop.`,
				},
				resource.Attribute{
					Name:        "next_hop_type",
					Description: `Type of the next hop.`,
				},
				resource.Attribute{
					Name:        "next_hop_region_id",
					Description: `ID of the region where the next hop is located.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "entries",
					Description: `A list of CEN Route Entries. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The destination CIDR block of the route entry.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the route entry.`,
				},
				resource.Attribute{
					Name:        "next_hop_id",
					Description: `ID of the next hop.`,
				},
				resource.Attribute{
					Name:        "next_hop_type",
					Description: `Type of the next hop.`,
				},
				resource.Attribute{
					Name:        "next_hop_region_id",
					Description: `ID of the region where the next hop is located.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cen_route_entries",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of CEN Route Entries owned by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) ID of the CEN instance.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Required) ID of the route table of the VPC or VBR.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Optional) The destination CIDR block of the route entry to query.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "entries",
					Description: `A list of CEN Route Entries. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `ID of the route table.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The destination CIDR block of the route entry.`,
				},
				resource.Attribute{
					Name:        "next_hop_id",
					Description: `ID of the next hop.`,
				},
				resource.Attribute{
					Name:        "next_hop_type",
					Description: `Type of the next hop, including "Instance", "HaVip" and "RouterInterface".`,
				},
				resource.Attribute{
					Name:        "route_type",
					Description: `Type of the route entry, including "System", "Custom" and "BGP".`,
				},
				resource.Attribute{
					Name:        "operational_mode",
					Description: `Whether to allow the route entry to be published or removed to or from CEN.`,
				},
				resource.Attribute{
					Name:        "publish_status",
					Description: `The publish status of the route entry in CEN, including "Published" and "NonPublished".`,
				},
				resource.Attribute{
					Name:        "conflicts",
					Description: `A list of conflicted Route Entries. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The destination CIDR block of the conflicted route entry.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `ID of the region where the conflicted route entry is located.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the CEN child instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The type of the CEN child instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Reasons of exceptions.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "entries",
					Description: `A list of CEN Route Entries. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `ID of the route table.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The destination CIDR block of the route entry.`,
				},
				resource.Attribute{
					Name:        "next_hop_id",
					Description: `ID of the next hop.`,
				},
				resource.Attribute{
					Name:        "next_hop_type",
					Description: `Type of the next hop, including "Instance", "HaVip" and "RouterInterface".`,
				},
				resource.Attribute{
					Name:        "route_type",
					Description: `Type of the route entry, including "System", "Custom" and "BGP".`,
				},
				resource.Attribute{
					Name:        "operational_mode",
					Description: `Whether to allow the route entry to be published or removed to or from CEN.`,
				},
				resource.Attribute{
					Name:        "publish_status",
					Description: `The publish status of the route entry in CEN, including "Published" and "NonPublished".`,
				},
				resource.Attribute{
					Name:        "conflicts",
					Description: `A list of conflicted Route Entries. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The destination CIDR block of the conflicted route entry.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `ID of the region where the conflicted route entry is located.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the CEN child instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The type of the CEN child instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Reasons of exceptions.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cen_route_maps",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of CEN(Cloud Enterprise Network) Route Maps owned by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cen_id",
					Description: `(Required) The ID of the CEN instance.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of CEN route map IDs.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the route map, including ` + "`" + `Creating` + "`" + `, ` + "`" + `Active` + "`" + ` and ` + "`" + `Deleting` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description_regex",
					Description: `(Optional) A regex string to filter CEN route map by description.`,
				},
				resource.Attribute{
					Name:        "cen_region_id",
					Description: `(Optional) The ID of the region to which the CEN instance belongs.`,
				},
				resource.Attribute{
					Name:        "transmit_direction",
					Description: `(Optional) The direction in which the route map is applied, including ` + "`" + `RegionIn` + "`" + ` and ` + "`" + `RegionOut` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of CEN route map IDs.`,
				},
				resource.Attribute{
					Name:        "maps",
					Description: `A list of CEN instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the route map.`,
				},
				resource.Attribute{
					Name:        "cen_id",
					Description: `The ID of the CEN instance.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the route map.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the route map.`,
				},
				resource.Attribute{
					Name:        "cen_region_id",
					Description: `The ID of the region to which the CEN instance belongs.`,
				},
				resource.Attribute{
					Name:        "as_path_match_mode",
					Description: `A match statement. It indicates the mode in which the as-path attribute is matched.`,
				},
				resource.Attribute{
					Name:        "cidr_match_mode",
					Description: `A match statement. It indicates the mode in which the prefix attribute is matched.`,
				},
				resource.Attribute{
					Name:        "community_match_mode",
					Description: `A match statement. It indicates the mode in which the community attribute is matched.`,
				},
				resource.Attribute{
					Name:        "community_operate_mode",
					Description: `An action statement. It indicates the mode in which the community attribute is operated.`,
				},
				resource.Attribute{
					Name:        "destination_child_instance_types",
					Description: `A match statement that indicates the list of IDs of the destination instances.`,
				},
				resource.Attribute{
					Name:        "destination_cidr_blocks",
					Description: `A match statement that indicates the prefix list.`,
				},
				resource.Attribute{
					Name:        "destination_instance_ids",
					Description: `A match statement that indicates the list of IDs of the destination instances.`,
				},
				resource.Attribute{
					Name:        "destination_instance_ids_reverse_match",
					Description: `Indicates whether to enable the reverse match method of the DestinationInstanceIds match condition.`,
				},
				resource.Attribute{
					Name:        "destination_route_table_ids",
					Description: `A match statement that indicates the list of IDs of the destination route tables.`,
				},
				resource.Attribute{
					Name:        "map_result",
					Description: `The action that is performed to a route if the route meets all the match conditions.`,
				},
				resource.Attribute{
					Name:        "match_asns",
					Description: `A match statement that indicates the As path list.`,
				},
				resource.Attribute{
					Name:        "match_community_set",
					Description: `A match statement that indicates the community set.`,
				},
				resource.Attribute{
					Name:        "next_priority",
					Description: `The priority of the next route map that is associated with the current route map.`,
				},
				resource.Attribute{
					Name:        "operate_community_set",
					Description: `An action statement that operates the community attribute.`,
				},
				resource.Attribute{
					Name:        "preference",
					Description: `An action statement that modifies the preference of the route.`,
				},
				resource.Attribute{
					Name:        "prepend_as_path",
					Description: `Indicates AS Path prepending when a regional gateway receives or publishes a route.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority of the route map.`,
				},
				resource.Attribute{
					Name:        "route_map_id",
					Description: `The ID of the route map.`,
				},
				resource.Attribute{
					Name:        "route_types",
					Description: `A match statement that indicates the list of route types.`,
				},
				resource.Attribute{
					Name:        "source_child_instance_types",
					Description: `A match statement that indicates the list of IDs of the source instances.`,
				},
				resource.Attribute{
					Name:        "source_instance_ids",
					Description: `A match statement that indicates the list of IDs of the source instances.`,
				},
				resource.Attribute{
					Name:        "source_instance_ids_reverse_match",
					Description: `Indicates whether to enable the reverse match method of the SourceInstanceIds match condition.`,
				},
				resource.Attribute{
					Name:        "source_region_ids",
					Description: `A match statement that indicates the list of IDs of the source regions.`,
				},
				resource.Attribute{
					Name:        "source_route_table_ids",
					Description: `A match statement that indicates the list of IDs of the source route tables.`,
				},
				resource.Attribute{
					Name:        "transmit_direction",
					Description: `The direction in which the route map is applied.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of CEN route map IDs.`,
				},
				resource.Attribute{
					Name:        "maps",
					Description: `A list of CEN instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the route map.`,
				},
				resource.Attribute{
					Name:        "cen_id",
					Description: `The ID of the CEN instance.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the route map.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the route map.`,
				},
				resource.Attribute{
					Name:        "cen_region_id",
					Description: `The ID of the region to which the CEN instance belongs.`,
				},
				resource.Attribute{
					Name:        "as_path_match_mode",
					Description: `A match statement. It indicates the mode in which the as-path attribute is matched.`,
				},
				resource.Attribute{
					Name:        "cidr_match_mode",
					Description: `A match statement. It indicates the mode in which the prefix attribute is matched.`,
				},
				resource.Attribute{
					Name:        "community_match_mode",
					Description: `A match statement. It indicates the mode in which the community attribute is matched.`,
				},
				resource.Attribute{
					Name:        "community_operate_mode",
					Description: `An action statement. It indicates the mode in which the community attribute is operated.`,
				},
				resource.Attribute{
					Name:        "destination_child_instance_types",
					Description: `A match statement that indicates the list of IDs of the destination instances.`,
				},
				resource.Attribute{
					Name:        "destination_cidr_blocks",
					Description: `A match statement that indicates the prefix list.`,
				},
				resource.Attribute{
					Name:        "destination_instance_ids",
					Description: `A match statement that indicates the list of IDs of the destination instances.`,
				},
				resource.Attribute{
					Name:        "destination_instance_ids_reverse_match",
					Description: `Indicates whether to enable the reverse match method of the DestinationInstanceIds match condition.`,
				},
				resource.Attribute{
					Name:        "destination_route_table_ids",
					Description: `A match statement that indicates the list of IDs of the destination route tables.`,
				},
				resource.Attribute{
					Name:        "map_result",
					Description: `The action that is performed to a route if the route meets all the match conditions.`,
				},
				resource.Attribute{
					Name:        "match_asns",
					Description: `A match statement that indicates the As path list.`,
				},
				resource.Attribute{
					Name:        "match_community_set",
					Description: `A match statement that indicates the community set.`,
				},
				resource.Attribute{
					Name:        "next_priority",
					Description: `The priority of the next route map that is associated with the current route map.`,
				},
				resource.Attribute{
					Name:        "operate_community_set",
					Description: `An action statement that operates the community attribute.`,
				},
				resource.Attribute{
					Name:        "preference",
					Description: `An action statement that modifies the preference of the route.`,
				},
				resource.Attribute{
					Name:        "prepend_as_path",
					Description: `Indicates AS Path prepending when a regional gateway receives or publishes a route.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority of the route map.`,
				},
				resource.Attribute{
					Name:        "route_map_id",
					Description: `The ID of the route map.`,
				},
				resource.Attribute{
					Name:        "route_types",
					Description: `A match statement that indicates the list of route types.`,
				},
				resource.Attribute{
					Name:        "source_child_instance_types",
					Description: `A match statement that indicates the list of IDs of the source instances.`,
				},
				resource.Attribute{
					Name:        "source_instance_ids",
					Description: `A match statement that indicates the list of IDs of the source instances.`,
				},
				resource.Attribute{
					Name:        "source_instance_ids_reverse_match",
					Description: `Indicates whether to enable the reverse match method of the SourceInstanceIds match condition.`,
				},
				resource.Attribute{
					Name:        "source_region_ids",
					Description: `A match statement that indicates the list of IDs of the source regions.`,
				},
				resource.Attribute{
					Name:        "source_route_table_ids",
					Description: `A match statement that indicates the list of IDs of the source route tables.`,
				},
				resource.Attribute{
					Name:        "transmit_direction",
					Description: `The direction in which the route map is applied.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cloud_connect_networks",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of CCN(Cloud Connect Network) instances owned by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of CCN instances IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter CCN instances by name. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of CCN instances IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of CCN instances names.`,
				},
				resource.Attribute{
					Name:        "networks",
					Description: `A list of CCN instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the CCN instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the CCN instance.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `CidrBlock of the CCN instance.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `IsDefault of the CCN instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of CCN instances IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of CCN instances names.`,
				},
				resource.Attribute{
					Name:        "networks",
					Description: `A list of CCN instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the CCN instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the CCN instance.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `CidrBlock of the CCN instance.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `IsDefault of the CCN instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_common_bandwidth_packages",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Common Bandwidth Packages owned by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Common Bandwidth Packages IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew, Available in 1.58.0+) The Id of resource group which the common bandwidth package belongs. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Common Bandwidth Packages IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Common Bandwidth Packages names.`,
				},
				resource.Attribute{
					Name:        "packages",
					Description: `A list of Common Bandwidth Packages. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Common Bandwidth Package.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The peak bandwidth of the Internet Shared Bandwidth instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Common Bandwidth Package.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the Common Bandwidth Package.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Common Bandwidth Package instance.`,
				},
				resource.Attribute{
					Name:        "business_status",
					Description: `The business status of the Common Bandwidth Package instance.`,
				},
				resource.Attribute{
					Name:        "isp",
					Description: `ISP of the Common Bandwidth Package.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Time of creation.`,
				},
				resource.Attribute{
					Name:        "public_ip_addresses",
					Description: `Public ip addresses that in the Common Bandwidth Pakcage.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The Id of resource group which the common bandwidth package belongs. ## Public ip addresses Block The public ip addresses mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "allocation_id",
					Description: `The ID of the EIP instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Common Bandwidth Packages IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Common Bandwidth Packages names.`,
				},
				resource.Attribute{
					Name:        "packages",
					Description: `A list of Common Bandwidth Packages. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Common Bandwidth Package.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The peak bandwidth of the Internet Shared Bandwidth instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Common Bandwidth Package.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the Common Bandwidth Package.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Common Bandwidth Package instance.`,
				},
				resource.Attribute{
					Name:        "business_status",
					Description: `The business status of the Common Bandwidth Package instance.`,
				},
				resource.Attribute{
					Name:        "isp",
					Description: `ISP of the Common Bandwidth Package.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Time of creation.`,
				},
				resource.Attribute{
					Name:        "public_ip_addresses",
					Description: `Public ip addresses that in the Common Bandwidth Pakcage.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The Id of resource group which the common bandwidth package belongs. ## Public ip addresses Block The public ip addresses mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "allocation_id",
					Description: `The ID of the EIP instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cr_ee_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Container Registry Enterprise Edition instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of ids to filter results by instance id.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by instance name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Container Registry Enterprise Edition instances. Its element is an instance uuid.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of instance names.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of matched Container Registry Enterprise Editioninstances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of Container Registry Enterprise Edition instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of Container Registry Enterprise Edition instance.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region of Container Registry Enterprise Edition instance.`,
				},
				resource.Attribute{
					Name:        "specification",
					Description: `Specification of Container Registry Enterprise Edition instance.`,
				},
				resource.Attribute{
					Name:        "namespace_quota",
					Description: `The max number of namespaces that an instance can create.`,
				},
				resource.Attribute{
					Name:        "namespace_usage",
					Description: `The number of namespaces already created.`,
				},
				resource.Attribute{
					Name:        "repo_quota",
					Description: `The max number of repos that an instance can create.`,
				},
				resource.Attribute{
					Name:        "repo_usage",
					Description: `The number of repos already created.`,
				},
				resource.Attribute{
					Name:        "vpc_endpoints",
					Description: `A list of domains for access on vpc network.`,
				},
				resource.Attribute{
					Name:        "public_endpoints",
					Description: `A list of domains for access on internet network.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Container Registry Enterprise Edition instances. Its element is an instance uuid.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of instance names.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of matched Container Registry Enterprise Editioninstances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of Container Registry Enterprise Edition instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of Container Registry Enterprise Edition instance.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region of Container Registry Enterprise Edition instance.`,
				},
				resource.Attribute{
					Name:        "specification",
					Description: `Specification of Container Registry Enterprise Edition instance.`,
				},
				resource.Attribute{
					Name:        "namespace_quota",
					Description: `The max number of namespaces that an instance can create.`,
				},
				resource.Attribute{
					Name:        "namespace_usage",
					Description: `The number of namespaces already created.`,
				},
				resource.Attribute{
					Name:        "repo_quota",
					Description: `The max number of repos that an instance can create.`,
				},
				resource.Attribute{
					Name:        "repo_usage",
					Description: `The number of repos already created.`,
				},
				resource.Attribute{
					Name:        "vpc_endpoints",
					Description: `A list of domains for access on vpc network.`,
				},
				resource.Attribute{
					Name:        "public_endpoints",
					Description: `A list of domains for access on internet network.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cr_ee_namespaces",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Container Registry Enterprise Edition namespaces.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) ID of Container Registry Enterprise Edition instance.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of ids to filter results by namespace id.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by namespace name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Container Registry Enterprise Edition namespaces. Its element is a namespace uuid.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of namespace names.`,
				},
				resource.Attribute{
					Name:        "namespaces",
					Description: `A list of matched Container Registry Enterprise Edition namespaces. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of Container Registry Enterprise Edition instance.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of Container Registry Enterprise Edition namespace.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of Container Registry Enterprise Edition namespace.`,
				},
				resource.Attribute{
					Name:        "auto_create",
					Description: `Boolean, when it set to true, repositories are automatically created when pushing new images. If it set to false, you create repository for images before pushing.`,
				},
				resource.Attribute{
					Name:        "default_visibility",
					Description: `` + "`" + `PUBLIC` + "`" + ` or ` + "`" + `PRIVATE` + "`" + `, default repository visibility in this namespace.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Container Registry Enterprise Edition namespaces. Its element is a namespace uuid.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of namespace names.`,
				},
				resource.Attribute{
					Name:        "namespaces",
					Description: `A list of matched Container Registry Enterprise Edition namespaces. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of Container Registry Enterprise Edition instance.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of Container Registry Enterprise Edition namespace.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of Container Registry Enterprise Edition namespace.`,
				},
				resource.Attribute{
					Name:        "auto_create",
					Description: `Boolean, when it set to true, repositories are automatically created when pushing new images. If it set to false, you create repository for images before pushing.`,
				},
				resource.Attribute{
					Name:        "default_visibility",
					Description: `` + "`" + `PUBLIC` + "`" + ` or ` + "`" + `PRIVATE` + "`" + `, default repository visibility in this namespace.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cr_ee_repos",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Container Registry Enterprise Edition repositories.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) ID of Container Registry Enterprise Edition instance.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Name of Container Registry Enterprise Edition namespace where the repositories are located in.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of ids to filter results by repository id.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by repository name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "enable_details",
					Description: `(Optional) Boolean, false by default, only repository attributes are exported. Set to true if tags belong to this repository are needed. See ` + "`" + `tags` + "`" + ` in attributes. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Container Registry Enterprise Edition repositories. Its element is a repository id.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of repository names.`,
				},
				resource.Attribute{
					Name:        "repos",
					Description: `A list of matched Container Registry Enterprise Edition namespaces. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of Container Registry Enterprise Edition instance.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `Name of Container Registry Enterprise Edition namespace where repo is located.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of Container Registry Enterprise Edition repository.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of Container Registry Enterprise Edition repository.`,
				},
				resource.Attribute{
					Name:        "summary",
					Description: `The repository general information.`,
				},
				resource.Attribute{
					Name:        "repo_type",
					Description: `` + "`" + `PUBLIC` + "`" + ` or ` + "`" + `PRIVATE` + "`" + `, repository's visibility.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of image tags belong to this repository. Each contains several attributes, see ` + "`" + `Block Tag` + "`" + `. ### Block Tag`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `Tag of this image.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `Id of this image.`,
				},
				resource.Attribute{
					Name:        "digest",
					Description: `Digest of this image.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of this image.`,
				},
				resource.Attribute{
					Name:        "image_size",
					Description: `Status of this image, in bytes.`,
				},
				resource.Attribute{
					Name:        "image_update",
					Description: `Last update time of this image, unix time in nanoseconds.`,
				},
				resource.Attribute{
					Name:        "image_create",
					Description: `Create time of this image, unix time in nanoseconds.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Container Registry Enterprise Edition repositories. Its element is a repository id.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of repository names.`,
				},
				resource.Attribute{
					Name:        "repos",
					Description: `A list of matched Container Registry Enterprise Edition namespaces. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of Container Registry Enterprise Edition instance.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `Name of Container Registry Enterprise Edition namespace where repo is located.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of Container Registry Enterprise Edition repository.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of Container Registry Enterprise Edition repository.`,
				},
				resource.Attribute{
					Name:        "summary",
					Description: `The repository general information.`,
				},
				resource.Attribute{
					Name:        "repo_type",
					Description: `` + "`" + `PUBLIC` + "`" + ` or ` + "`" + `PRIVATE` + "`" + `, repository's visibility.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of image tags belong to this repository. Each contains several attributes, see ` + "`" + `Block Tag` + "`" + `. ### Block Tag`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `Tag of this image.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `Id of this image.`,
				},
				resource.Attribute{
					Name:        "digest",
					Description: `Digest of this image.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of this image.`,
				},
				resource.Attribute{
					Name:        "image_size",
					Description: `Status of this image, in bytes.`,
				},
				resource.Attribute{
					Name:        "image_update",
					Description: `Last update time of this image, unix time in nanoseconds.`,
				},
				resource.Attribute{
					Name:        "image_create",
					Description: `Create time of this image, unix time in nanoseconds.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cr_ee_sync_rules",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Container Registry Enterprise Edition sync rules.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) ID of Container Registry Enterprise Edition local instance.`,
				},
				resource.Attribute{
					Name:        "namespace_name",
					Description: `(Optional) Name of Container Registry Enterprise Edition local namespace.`,
				},
				resource.Attribute{
					Name:        "repo_name",
					Description: `(Optional) Name of Container Registry Enterprise Edition local repo.`,
				},
				resource.Attribute{
					Name:        "target_instance_id",
					Description: `(Optional) ID of Container Registry Enterprise Edition target instance.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by sync rule name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of ids to filter results by sync rule id.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Container Registry Enterprise Edition sync rules. Its element is a sync rule uuid.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of sync rule names.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `A list of matched Container Registry Enterprise Edition sync rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of Container Registry Enterprise Edition sync rule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of Container Registry Enterprise Edition sync rule.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region of Container Registry Enterprise Edition local instance.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of Container Registry Enterprise Edition local instance.`,
				},
				resource.Attribute{
					Name:        "namespace_name",
					Description: `Name of Container Registry Enterprise Edition local namespace.`,
				},
				resource.Attribute{
					Name:        "repo_name",
					Description: `Name of Container Registry Enterprise Edition local repo.`,
				},
				resource.Attribute{
					Name:        "target_region_id",
					Description: `Region of Container Registry Enterprise Edition target instance.`,
				},
				resource.Attribute{
					Name:        "target_instance_id",
					Description: `ID of Container Registry Enterprise Edition target instance.`,
				},
				resource.Attribute{
					Name:        "target_namespace_name",
					Description: `Name of Container Registry Enterprise Edition target namespace.`,
				},
				resource.Attribute{
					Name:        "target_repo_name",
					Description: `Name of Container Registry Enterprise Edition target repo.`,
				},
				resource.Attribute{
					Name:        "tag_filter",
					Description: `The regular expression used to filter image tags for synchronization in the source repository.`,
				},
				resource.Attribute{
					Name:        "sync_direction",
					Description: `` + "`" + `FROM` + "`" + ` or ` + "`" + `TO` + "`" + `, the direction of synchronization. ` + "`" + `FROM` + "`" + ` indicates that the local instance is the source instance. ` + "`" + `TO` + "`" + ` indicates that the local instance is the target instance to be synchronized.`,
				},
				resource.Attribute{
					Name:        "sync_scope",
					Description: `` + "`" + `REPO` + "`" + ` or ` + "`" + `NAMESPACE` + "`" + `,the scope that the synchronization rule applies.`,
				},
				resource.Attribute{
					Name:        "sync_trigger",
					Description: `` + "`" + `PASSIVE` + "`" + ` or ` + "`" + `INITIATIVE` + "`" + `, the policy configured to trigger the synchronization rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Container Registry Enterprise Edition sync rules. Its element is a sync rule uuid.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of sync rule names.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `A list of matched Container Registry Enterprise Edition sync rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of Container Registry Enterprise Edition sync rule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of Container Registry Enterprise Edition sync rule.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region of Container Registry Enterprise Edition local instance.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of Container Registry Enterprise Edition local instance.`,
				},
				resource.Attribute{
					Name:        "namespace_name",
					Description: `Name of Container Registry Enterprise Edition local namespace.`,
				},
				resource.Attribute{
					Name:        "repo_name",
					Description: `Name of Container Registry Enterprise Edition local repo.`,
				},
				resource.Attribute{
					Name:        "target_region_id",
					Description: `Region of Container Registry Enterprise Edition target instance.`,
				},
				resource.Attribute{
					Name:        "target_instance_id",
					Description: `ID of Container Registry Enterprise Edition target instance.`,
				},
				resource.Attribute{
					Name:        "target_namespace_name",
					Description: `Name of Container Registry Enterprise Edition target namespace.`,
				},
				resource.Attribute{
					Name:        "target_repo_name",
					Description: `Name of Container Registry Enterprise Edition target repo.`,
				},
				resource.Attribute{
					Name:        "tag_filter",
					Description: `The regular expression used to filter image tags for synchronization in the source repository.`,
				},
				resource.Attribute{
					Name:        "sync_direction",
					Description: `` + "`" + `FROM` + "`" + ` or ` + "`" + `TO` + "`" + `, the direction of synchronization. ` + "`" + `FROM` + "`" + ` indicates that the local instance is the source instance. ` + "`" + `TO` + "`" + ` indicates that the local instance is the target instance to be synchronized.`,
				},
				resource.Attribute{
					Name:        "sync_scope",
					Description: `` + "`" + `REPO` + "`" + ` or ` + "`" + `NAMESPACE` + "`" + `,the scope that the synchronization rule applies.`,
				},
				resource.Attribute{
					Name:        "sync_trigger",
					Description: `` + "`" + `PASSIVE` + "`" + ` or ` + "`" + `INITIATIVE` + "`" + `, the policy configured to trigger the synchronization rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cr_namespaces",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Container Registry namespaces.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by namespace name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Container Registry namespaces. Its element is a namespace name.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of namespace names.`,
				},
				resource.Attribute{
					Name:        "namespaces",
					Description: `A list of matched Container Registry namespaces. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of Container Registry namespace.`,
				},
				resource.Attribute{
					Name:        "auto_create",
					Description: `Boolean, when it set to true, repositories are automatically created when pushing new images. If it set to false, you create repository for images before pushing.`,
				},
				resource.Attribute{
					Name:        "default_visibility",
					Description: `` + "`" + `PUBLIC` + "`" + ` or ` + "`" + `PRIVATE` + "`" + `, default repository visibility in this namespace.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Container Registry namespaces. Its element is a namespace name.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of namespace names.`,
				},
				resource.Attribute{
					Name:        "namespaces",
					Description: `A list of matched Container Registry namespaces. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of Container Registry namespace.`,
				},
				resource.Attribute{
					Name:        "auto_create",
					Description: `Boolean, when it set to true, repositories are automatically created when pushing new images. If it set to false, you create repository for images before pushing.`,
				},
				resource.Attribute{
					Name:        "default_visibility",
					Description: `` + "`" + `PUBLIC` + "`" + ` or ` + "`" + `PRIVATE` + "`" + `, default repository visibility in this namespace.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cr_repos",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Container Registry repositories.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional) Name of container registry namespace where the repositories are located in.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by repository name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "enable_details",
					Description: `(Optional) Boolean, false by default, only repository attributes are exported. Set to true if domain list and tags belong to this repository are needed. See ` + "`" + `tags` + "`" + ` in attributes. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Container Registry Repositories. Its element is set to ` + "`" + `names` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of repository names.`,
				},
				resource.Attribute{
					Name:        "repos",
					Description: `A list of matched Container Registry Namespaces. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `Name of container registry namespace where repo is located.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of container registry namespace.`,
				},
				resource.Attribute{
					Name:        "summary",
					Description: `The repository general information.`,
				},
				resource.Attribute{
					Name:        "repo_type",
					Description: `` + "`" + `PUBLIC` + "`" + ` or ` + "`" + `PRIVATE` + "`" + `, repository's visibility.`,
				},
				resource.Attribute{
					Name:        "domain_list",
					Description: `The repository domain list.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Domain of public endpoint.`,
				},
				resource.Attribute{
					Name:        "internal",
					Description: `Domain of internal endpoint, only in some regions.`,
				},
				resource.Attribute{
					Name:        "vpc",
					Description: `Domain of vpc endpoint.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of image tags belong to this repository. Each contains several attributes, see ` + "`" + `Block Tag` + "`" + `. ### Block Tag`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `Tag of this image.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `Id of this image.`,
				},
				resource.Attribute{
					Name:        "digest",
					Description: `Digest of this image.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of this image.`,
				},
				resource.Attribute{
					Name:        "image_size",
					Description: `Status of this image, in bytes.`,
				},
				resource.Attribute{
					Name:        "image_update",
					Description: `Last update time of this image, unix time in nanoseconds.`,
				},
				resource.Attribute{
					Name:        "image_create",
					Description: `Create time of this image, unix time in nanoseconds.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Container Registry Repositories. Its element is set to ` + "`" + `names` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of repository names.`,
				},
				resource.Attribute{
					Name:        "repos",
					Description: `A list of matched Container Registry Namespaces. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `Name of container registry namespace where repo is located.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of container registry namespace.`,
				},
				resource.Attribute{
					Name:        "summary",
					Description: `The repository general information.`,
				},
				resource.Attribute{
					Name:        "repo_type",
					Description: `` + "`" + `PUBLIC` + "`" + ` or ` + "`" + `PRIVATE` + "`" + `, repository's visibility.`,
				},
				resource.Attribute{
					Name:        "domain_list",
					Description: `The repository domain list.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Domain of public endpoint.`,
				},
				resource.Attribute{
					Name:        "internal",
					Description: `Domain of internal endpoint, only in some regions.`,
				},
				resource.Attribute{
					Name:        "vpc",
					Description: `Domain of vpc endpoint.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of image tags belong to this repository. Each contains several attributes, see ` + "`" + `Block Tag` + "`" + `. ### Block Tag`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `Tag of this image.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `Id of this image.`,
				},
				resource.Attribute{
					Name:        "digest",
					Description: `Digest of this image.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of this image.`,
				},
				resource.Attribute{
					Name:        "image_size",
					Description: `Status of this image, in bytes.`,
				},
				resource.Attribute{
					Name:        "image_update",
					Description: `Last update time of this image, unix time in nanoseconds.`,
				},
				resource.Attribute{
					Name:        "image_create",
					Description: `Create time of this image, unix time in nanoseconds.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cs_kubernetes_clusters",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Container Service Kubernetes Clusters to be used by the alicloud_cs_kubernetes_clusters resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) Cluster IDs to filter.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by cluster name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "enabled_details",
					Description: `(Optional) Boolean, false by default, only ` + "`" + `id` + "`" + ` and ` + "`" + `name` + "`" + ` are exported. Set to true if more details are needed, e.g., ` + "`" + `master_disk_category` + "`" + `, ` + "`" + `slb_internet_enabled` + "`" + `, ` + "`" + `connections` + "`" + `. See full list in attributes. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Kubernetes clusters' ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of matched Kubernetes clusters' names.`,
				},
				resource.Attribute{
					Name:        "clusters",
					Description: `A list of matched Kubernetes clusters. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the container cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the container cluster.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The ID of availability zone.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `The keypair of ssh login cluster node, you have to create it first.`,
				},
				resource.Attribute{
					Name:        "worker_numbers",
					Description: `The ECS instance node number in the current container cluster.`,
				},
				resource.Attribute{
					Name:        "vswitch_ids",
					Description: `The ID of VSwitches where the current cluster is located.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of VPC where the current cluster is located.`,
				},
				resource.Attribute{
					Name:        "slb_internet_enabled",
					Description: `Whether internet load balancer for API Server is created`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of security group where the current cluster worker node is located.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of node image.`,
				},
				resource.Attribute{
					Name:        "nat_gateway_id",
					Description: `The ID of nat gateway used to launch kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "master_instance_types",
					Description: `The instance type of master node.`,
				},
				resource.Attribute{
					Name:        "worker_instance_types",
					Description: `The instance type of worker node.`,
				},
				resource.Attribute{
					Name:        "master_disk_category",
					Description: `The system disk category of master node.`,
				},
				resource.Attribute{
					Name:        "master_disk_size",
					Description: `The system disk size of master node.`,
				},
				resource.Attribute{
					Name:        "worker_disk_category",
					Description: `The system disk category of worker node.`,
				},
				resource.Attribute{
					Name:        "worker_disk_size",
					Description: `The system disk size of worker node.`,
				},
				resource.Attribute{
					Name:        "worker_data_disk_category",
					Description: `The data disk size of worker node.`,
				},
				resource.Attribute{
					Name:        "worker_data_disk_size",
					Description: `The data disk category of worker node.`,
				},
				resource.Attribute{
					Name:        "master_nodes",
					Description: `List of cluster master nodes. It contains several attributes to ` + "`" + `Block Nodes` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "worker_nodes",
					Description: `List of cluster worker nodes. It contains several attributes to ` + "`" + `Block Nodes` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "connections",
					Description: `Map of kubernetes cluster connection information. It contains several attributes to ` + "`" + `Block Connections` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "node_cidr_mask",
					Description: `The network mask used on pods for each node.`,
				},
				resource.Attribute{
					Name:        "log_config",
					Description: `A list of one element containing information about the associated log store. It contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of collecting logs.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `Log Service project name. ### Block Nodes`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the node.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Node name.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address of node.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Deprecated from version 1.9.4) ### Block Connections`,
				},
				resource.Attribute{
					Name:        "api_server_internet",
					Description: `API Server Internet endpoint.`,
				},
				resource.Attribute{
					Name:        "api_server_intranet",
					Description: `API Server Intranet endpoint.`,
				},
				resource.Attribute{
					Name:        "master_public_ip",
					Description: `Master node SSH IP address.`,
				},
				resource.Attribute{
					Name:        "service_domain",
					Description: `Service Access Domain.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Kubernetes clusters' ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of matched Kubernetes clusters' names.`,
				},
				resource.Attribute{
					Name:        "clusters",
					Description: `A list of matched Kubernetes clusters. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the container cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the container cluster.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The ID of availability zone.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `The keypair of ssh login cluster node, you have to create it first.`,
				},
				resource.Attribute{
					Name:        "worker_numbers",
					Description: `The ECS instance node number in the current container cluster.`,
				},
				resource.Attribute{
					Name:        "vswitch_ids",
					Description: `The ID of VSwitches where the current cluster is located.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of VPC where the current cluster is located.`,
				},
				resource.Attribute{
					Name:        "slb_internet_enabled",
					Description: `Whether internet load balancer for API Server is created`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of security group where the current cluster worker node is located.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The ID of node image.`,
				},
				resource.Attribute{
					Name:        "nat_gateway_id",
					Description: `The ID of nat gateway used to launch kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "master_instance_types",
					Description: `The instance type of master node.`,
				},
				resource.Attribute{
					Name:        "worker_instance_types",
					Description: `The instance type of worker node.`,
				},
				resource.Attribute{
					Name:        "master_disk_category",
					Description: `The system disk category of master node.`,
				},
				resource.Attribute{
					Name:        "master_disk_size",
					Description: `The system disk size of master node.`,
				},
				resource.Attribute{
					Name:        "worker_disk_category",
					Description: `The system disk category of worker node.`,
				},
				resource.Attribute{
					Name:        "worker_disk_size",
					Description: `The system disk size of worker node.`,
				},
				resource.Attribute{
					Name:        "worker_data_disk_category",
					Description: `The data disk size of worker node.`,
				},
				resource.Attribute{
					Name:        "worker_data_disk_size",
					Description: `The data disk category of worker node.`,
				},
				resource.Attribute{
					Name:        "master_nodes",
					Description: `List of cluster master nodes. It contains several attributes to ` + "`" + `Block Nodes` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "worker_nodes",
					Description: `List of cluster worker nodes. It contains several attributes to ` + "`" + `Block Nodes` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "connections",
					Description: `Map of kubernetes cluster connection information. It contains several attributes to ` + "`" + `Block Connections` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "node_cidr_mask",
					Description: `The network mask used on pods for each node.`,
				},
				resource.Attribute{
					Name:        "log_config",
					Description: `A list of one element containing information about the associated log store. It contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of collecting logs.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `Log Service project name. ### Block Nodes`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the node.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Node name.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address of node.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Deprecated from version 1.9.4) ### Block Connections`,
				},
				resource.Attribute{
					Name:        "api_server_internet",
					Description: `API Server Internet endpoint.`,
				},
				resource.Attribute{
					Name:        "api_server_intranet",
					Description: `API Server Intranet endpoint.`,
				},
				resource.Attribute{
					Name:        "master_public_ip",
					Description: `Master node SSH IP address.`,
				},
				resource.Attribute{
					Name:        "service_domain",
					Description: `Service Access Domain.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cs_managed_kubernetes_clusters",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Container Service Managed Kubernetes Clusters to be used by the alicloud_cs_managed_kubernetes_clusters resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) Cluster IDs to filter.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by cluster name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "enabled_details",
					Description: `(Optional) Boolean, false by default, only ` + "`" + `id` + "`" + ` and ` + "`" + `name` + "`" + ` are exported. Set to true if more details are needed, e.g., ` + "`" + `master_disk_category` + "`" + `, ` + "`" + `slb_internet_enabled` + "`" + `, ` + "`" + `connections` + "`" + `. See full list in attributes. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Kubernetes clusters' ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of matched Kubernetes clusters' names.`,
				},
				resource.Attribute{
					Name:        "clusters",
					Description: `A list of matched Kubernetes clusters. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the container cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the container cluster.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The ID of availability zone.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `The keypair of ssh login cluster node, you have to create it first.`,
				},
				resource.Attribute{
					Name:        "worker_numbers",
					Description: `The ECS instance node number in the current container cluster.`,
				},
				resource.Attribute{
					Name:        "vswitch_ids",
					Description: `The ID of VSwitches where the current cluster is located.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of VPC where the current cluster is located.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of security group where the current cluster worker node is located.`,
				},
				resource.Attribute{
					Name:        "nat_gateway_id",
					Description: `The ID of nat gateway used to launch kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "worker_nodes",
					Description: `List of cluster worker nodes. It contains several attributes to ` + "`" + `Block Nodes` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "connections",
					Description: `Map of kubernetes cluster connection information. It contains several attributes to ` + "`" + `Block Connections` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_config",
					Description: `A list of one element containing information about the associated log store. It contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of collecting logs.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `Log Service project name. ### Block Nodes`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the node.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Node name.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address of node.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Deprecated from version 1.9.4) ### Block Connections`,
				},
				resource.Attribute{
					Name:        "api_server_internet",
					Description: `API Server Internet endpoint.`,
				},
				resource.Attribute{
					Name:        "api_server_intranet",
					Description: `API Server Intranet endpoint.`,
				},
				resource.Attribute{
					Name:        "master_public_ip",
					Description: `Master node SSH IP address.`,
				},
				resource.Attribute{
					Name:        "service_domain",
					Description: `Service Access Domain.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Kubernetes clusters' ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of matched Kubernetes clusters' names.`,
				},
				resource.Attribute{
					Name:        "clusters",
					Description: `A list of matched Kubernetes clusters. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the container cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the container cluster.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The ID of availability zone.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `The keypair of ssh login cluster node, you have to create it first.`,
				},
				resource.Attribute{
					Name:        "worker_numbers",
					Description: `The ECS instance node number in the current container cluster.`,
				},
				resource.Attribute{
					Name:        "vswitch_ids",
					Description: `The ID of VSwitches where the current cluster is located.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of VPC where the current cluster is located.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of security group where the current cluster worker node is located.`,
				},
				resource.Attribute{
					Name:        "nat_gateway_id",
					Description: `The ID of nat gateway used to launch kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "worker_nodes",
					Description: `List of cluster worker nodes. It contains several attributes to ` + "`" + `Block Nodes` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "connections",
					Description: `Map of kubernetes cluster connection information. It contains several attributes to ` + "`" + `Block Connections` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_config",
					Description: `A list of one element containing information about the associated log store. It contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of collecting logs.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `Log Service project name. ### Block Nodes`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the node.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Node name.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address of node.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Deprecated from version 1.9.4) ### Block Connections`,
				},
				resource.Attribute{
					Name:        "api_server_internet",
					Description: `API Server Internet endpoint.`,
				},
				resource.Attribute{
					Name:        "api_server_intranet",
					Description: `API Server Intranet endpoint.`,
				},
				resource.Attribute{
					Name:        "master_public_ip",
					Description: `Master node SSH IP address.`,
				},
				resource.Attribute{
					Name:        "service_domain",
					Description: `Service Access Domain.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_cs_serverless_kubernetes_clusters",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Container Service Serverless Kubernetes Clusters to be used by the alicloud_cs_serverless_kubernetes_clusters resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) Cluster IDs to filter.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by cluster name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "enabled_details",
					Description: `(Optional) Boolean, false by default, only ` + "`" + `id` + "`" + ` and ` + "`" + `name` + "`" + ` are exported. Set to true if more details are needed, e.g., ` + "`" + `deletion_protection` + "`" + `, ` + "`" + `connections` + "`" + `. See full list in attributes. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Kubernetes clusters' ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of matched Kubernetes clusters' names.`,
				},
				resource.Attribute{
					Name:        "clusters",
					Description: `A list of matched Kubernetes clusters. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the container cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the container cluster.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `The ID of VSwitch where the current cluster is located.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of VPC where the current cluster is located.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of security group where the current cluster is located.`,
				},
				resource.Attribute{
					Name:        "nat_gateway_id",
					Description: `The ID of nat gateway used to launch kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `Whether the cluster support delete protection.`,
				},
				resource.Attribute{
					Name:        "connections",
					Description: `Map of serverless cluster connection information. It contains several attributes to ` + "`" + `Block Connections` + "`" + `. ### Block Connections`,
				},
				resource.Attribute{
					Name:        "api_server_internet",
					Description: `API Server Internet endpoint.`,
				},
				resource.Attribute{
					Name:        "api_server_intranet",
					Description: `API Server Intranet endpoint.`,
				},
				resource.Attribute{
					Name:        "master_public_ip",
					Description: `Master node SSH IP address.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of matched Kubernetes clusters' ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of matched Kubernetes clusters' names.`,
				},
				resource.Attribute{
					Name:        "clusters",
					Description: `A list of matched Kubernetes clusters. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the container cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the container cluster.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `The ID of VSwitch where the current cluster is located.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of VPC where the current cluster is located.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of security group where the current cluster is located.`,
				},
				resource.Attribute{
					Name:        "nat_gateway_id",
					Description: `The ID of nat gateway used to launch kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `Whether the cluster support delete protection.`,
				},
				resource.Attribute{
					Name:        "connections",
					Description: `Map of serverless cluster connection information. It contains several attributes to ` + "`" + `Block Connections` + "`" + `. ### Block Connections`,
				},
				resource.Attribute{
					Name:        "api_server_internet",
					Description: `API Server Internet endpoint.`,
				},
				resource.Attribute{
					Name:        "api_server_intranet",
					Description: `API Server Intranet endpoint.`,
				},
				resource.Attribute{
					Name:        "master_public_ip",
					Description: `Master node SSH IP address.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_db_instance_classes",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of RDS instacne classes info.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional) The Zone to launch the DB instance.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional) Filter the results by charge type. Valid values: ` + "`" + `PrePaid` + "`" + ` and ` + "`" + `PostPaid` + "`" + `. Default to ` + "`" + `PostPaid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Optional) Database type. Options are ` + "`" + `MySQL` + "`" + `, ` + "`" + `SQLServer` + "`" + `, ` + "`" + `PostgreSQL` + "`" + ` and ` + "`" + `PPAS` + "`" + `. If no value is specified, all types are returned.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Optional) DB Instance category. the value like [` + "`" + `Basic` + "`" + `, ` + "`" + `HighAvailability` + "`" + `, ` + "`" + `Finance` + "`" + `], [detail info](https://www.alibabacloud.com/help/doc-detail/69795.htm).`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Optional) Database version required by the user. Value options can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/26228.htm) ` + "`" + `EngineVersion` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "db_instance_class",
					Description: `(Optional, Available in 1.51.0+) The DB instance class type by the user.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Optional) The DB instance storage space required by the user. Valid values: ` + "`" + `cloud_ssd` + "`" + ` and ` + "`" + `local_ssd` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "multi_zone",
					Description: `(Optional, Available in v1.48.0+) Whether to show multi available zone. Default false to not show multi availability zone.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform apply` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Available in 1.60.0+) A list of Rds instance class codes.`,
				},
				resource.Attribute{
					Name:        "instance_classes",
					Description: `A list of Rds available resource. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "zone_ids",
					Description: `A list of Zone to launch the DB instance.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Zone to launch the DB instance`,
				},
				resource.Attribute{
					Name:        "sub_zone_ids",
					Description: `A list of sub zone ids which in the id - e.g If ` + "`" + `id` + "`" + ` is ` + "`" + `cn-beijing-MAZ5(a,b)` + "`" + `, ` + "`" + `sub_zone_ids` + "`" + ` will be ` + "`" + `["cn-beijing-a", "cn-beijing-b"]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_class",
					Description: `DB Instance available class.`,
				},
				resource.Attribute{
					Name:        "storage_range",
					Description: `DB Instance available storage range.`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `DB Instance available storage min value.`,
				},
				resource.Attribute{
					Name:        "max",
					Description: `DB Instance available storage max value.`,
				},
				resource.Attribute{
					Name:        "step",
					Description: `DB Instance available storage increase step.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Available in 1.60.0+) A list of Rds instance class codes.`,
				},
				resource.Attribute{
					Name:        "instance_classes",
					Description: `A list of Rds available resource. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "zone_ids",
					Description: `A list of Zone to launch the DB instance.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Zone to launch the DB instance`,
				},
				resource.Attribute{
					Name:        "sub_zone_ids",
					Description: `A list of sub zone ids which in the id - e.g If ` + "`" + `id` + "`" + ` is ` + "`" + `cn-beijing-MAZ5(a,b)` + "`" + `, ` + "`" + `sub_zone_ids` + "`" + ` will be ` + "`" + `["cn-beijing-a", "cn-beijing-b"]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_class",
					Description: `DB Instance available class.`,
				},
				resource.Attribute{
					Name:        "storage_range",
					Description: `DB Instance available storage range.`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `DB Instance available storage min value.`,
				},
				resource.Attribute{
					Name:        "max",
					Description: `DB Instance available storage max value.`,
				},
				resource.Attribute{
					Name:        "step",
					Description: `DB Instance available storage increase step.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_db_instance_engines",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of RDS instacne engines resource info.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional) The Zone to launch the DB instance.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional) Filter the results by charge type. Valid values: ` + "`" + `PrePaid` + "`" + ` and ` + "`" + `PostPaid` + "`" + `. Default to ` + "`" + `PostPaid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Optional) Database type. Options are ` + "`" + `MySQL` + "`" + `, ` + "`" + `SQLServer` + "`" + `, ` + "`" + `PostgreSQL` + "`" + ` and ` + "`" + `PPAS` + "`" + `. If no value is specified, all types are returned.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Optional) Database version required by the user. Value options can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/26228.htm) ` + "`" + `EngineVersion` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "multi_zone",
					Description: `(Optional, Available in v1.48.0+) Whether to show multi available zone. Default false to not show multi availability zone.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform apply` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "instance_engines",
					Description: `A list of Rds available resource. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "zone_ids",
					Description: `A list of Zone to launch the DB instance.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Zone to launch the DB instance`,
				},
				resource.Attribute{
					Name:        "sub_zone_ids",
					Description: `A list of sub zone ids which in the id - e.g If ` + "`" + `id` + "`" + ` is ` + "`" + `cn-beijing-MAZ5(a,b)` + "`" + `, ` + "`" + `sub_zone_ids` + "`" + ` will be ` + "`" + `["cn-beijing-a", "cn-beijing-b"]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Database type.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `DB Instance version.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `DB Instance category.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_engines",
					Description: `A list of Rds available resource. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "zone_ids",
					Description: `A list of Zone to launch the DB instance.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Zone to launch the DB instance`,
				},
				resource.Attribute{
					Name:        "sub_zone_ids",
					Description: `A list of sub zone ids which in the id - e.g If ` + "`" + `id` + "`" + ` is ` + "`" + `cn-beijing-MAZ5(a,b)` + "`" + `, ` + "`" + `sub_zone_ids` + "`" + ` will be ` + "`" + `["cn-beijing-a", "cn-beijing-b"]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Database type.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `DB Instance version.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `DB Instance category.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_db_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a collection of RDS instances according to the specified filters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by instance name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, Available 1.52.0+) A list of RDS instance IDs.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Optional) Database type. Options are ` + "`" + `MySQL` + "`" + `, ` + "`" + `SQLServer` + "`" + `, ` + "`" + `PostgreSQL` + "`" + ` and ` + "`" + `PPAS` + "`" + `. If no value is specified, all types are returned.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Status of the instance.`,
				},
				resource.Attribute{
					Name:        "db_type",
					Description: `(Optional) ` + "`" + `Primary` + "`" + ` for primary instance, ` + "`" + `Readonly` + "`" + ` for read-only instance, ` + "`" + `Guard` + "`" + ` for disaster recovery instance, and ` + "`" + `Temp` + "`" + ` for temporary instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) Used to retrieve instances belong to specified VPC.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional) Used to retrieve instances belong to specified ` + "`" + `vswitch` + "`" + ` resources.`,
				},
				resource.Attribute{
					Name:        "connection_mode",
					Description: `(Optional) ` + "`" + `Standard` + "`" + ` for standard access mode and ` + "`" + `Safe` + "`" + ` for high security access mode.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A map of tags assigned to the DB instances. Note: Before 1.60.0, the value's format is a ` + "`" + `json` + "`" + ` string which including ` + "`" + `TagKey` + "`" + ` and ` + "`" + `TagValue` + "`" + `. ` + "`" + `TagKey` + "`" + ` cannot be null, and ` + "`" + `TagValue` + "`" + ` can be empty. Format example ` + "`" + `"{\"key1\":\"value1\"}"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of RDS instance IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of RDS instance names.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of RDS instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the RDS instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the RDS instance.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Billing method. Value options: ` + "`" + `Postpaid` + "`" + ` for Pay-As-You-Go and ` + "`" + `Prepaid` + "`" + ` for subscription.`,
				},
				resource.Attribute{
					Name:        "db_type",
					Description: `` + "`" + `Primary` + "`" + ` for primary instance, ` + "`" + `Readonly` + "`" + ` for read-only instance, ` + "`" + `Guard` + "`" + ` for disaster recovery instance, and ` + "`" + `Temp` + "`" + ` for temporary instance.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the instance.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `Expiration time. Pay-As-You-Go instances never expire.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Database type. Options are ` + "`" + `MySQL` + "`" + `, ` + "`" + `SQLServer` + "`" + `, ` + "`" + `PostgreSQL` + "`" + ` and ` + "`" + `PPAS` + "`" + `. If no value is specified, all types are returned.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Database version.`,
				},
				resource.Attribute{
					Name:        "net_type",
					Description: `` + "`" + `Internet` + "`" + ` for public network or ` + "`" + `Intranet` + "`" + ` for private network.`,
				},
				resource.Attribute{
					Name:        "connection_mode",
					Description: `` + "`" + `Standard` + "`" + ` for standard access mode and ` + "`" + `Safe` + "`" + ` for high security access mode.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Sizing of the RDS instance.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone.`,
				},
				resource.Attribute{
					Name:        "master_instance_id",
					Description: `ID of the primary instance. If this parameter is not returned, the current instance is a primary instance.`,
				},
				resource.Attribute{
					Name:        "guard_instance_id",
					Description: `If a disaster recovery instance is attached to the current instance, the ID of the disaster recovery instance applies.`,
				},
				resource.Attribute{
					Name:        "temp_instance_id",
					Description: `If a temporary instance is attached to the current instance, the ID of the temporary instance applies.`,
				},
				resource.Attribute{
					Name:        "readonly_instance_ids",
					Description: `A list of IDs of read-only instances attached to the primary instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `ID of the VSwitch the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Available in 1.70.3+) RDS database connection port.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `(Available in 1.70.3+) RDS database connection string.`,
				},
				resource.Attribute{
					Name:        "instance_storage",
					Description: `(Available in 1.70.3+) User-defined DB instance storage space.`,
				},
				resource.Attribute{
					Name:        "db_instance_storage_type",
					Description: `(Available in 1.70.3+) The storage type of the instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of RDS instance IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of RDS instance names.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of RDS instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the RDS instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the RDS instance.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Billing method. Value options: ` + "`" + `Postpaid` + "`" + ` for Pay-As-You-Go and ` + "`" + `Prepaid` + "`" + ` for subscription.`,
				},
				resource.Attribute{
					Name:        "db_type",
					Description: `` + "`" + `Primary` + "`" + ` for primary instance, ` + "`" + `Readonly` + "`" + ` for read-only instance, ` + "`" + `Guard` + "`" + ` for disaster recovery instance, and ` + "`" + `Temp` + "`" + ` for temporary instance.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the instance.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `Expiration time. Pay-As-You-Go instances never expire.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Database type. Options are ` + "`" + `MySQL` + "`" + `, ` + "`" + `SQLServer` + "`" + `, ` + "`" + `PostgreSQL` + "`" + ` and ` + "`" + `PPAS` + "`" + `. If no value is specified, all types are returned.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Database version.`,
				},
				resource.Attribute{
					Name:        "net_type",
					Description: `` + "`" + `Internet` + "`" + ` for public network or ` + "`" + `Intranet` + "`" + ` for private network.`,
				},
				resource.Attribute{
					Name:        "connection_mode",
					Description: `` + "`" + `Standard` + "`" + ` for standard access mode and ` + "`" + `Safe` + "`" + ` for high security access mode.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Sizing of the RDS instance.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone.`,
				},
				resource.Attribute{
					Name:        "master_instance_id",
					Description: `ID of the primary instance. If this parameter is not returned, the current instance is a primary instance.`,
				},
				resource.Attribute{
					Name:        "guard_instance_id",
					Description: `If a disaster recovery instance is attached to the current instance, the ID of the disaster recovery instance applies.`,
				},
				resource.Attribute{
					Name:        "temp_instance_id",
					Description: `If a temporary instance is attached to the current instance, the ID of the temporary instance applies.`,
				},
				resource.Attribute{
					Name:        "readonly_instance_ids",
					Description: `A list of IDs of read-only instances attached to the primary instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `ID of the VSwitch the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Available in 1.70.3+) RDS database connection port.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `(Available in 1.70.3+) RDS database connection string.`,
				},
				resource.Attribute{
					Name:        "instance_storage",
					Description: `(Available in 1.70.3+) User-defined DB instance storage space.`,
				},
				resource.Attribute{
					Name:        "db_instance_storage_type",
					Description: `(Available in 1.70.3+) The storage type of the instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_db_zones",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of availability zones for RDS that can be used by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "multi",
					Description: `(Optional) Indicate whether the zones can be used in a multi AZ configuration. Default to ` + "`" + `false` + "`" + `. Multi AZ is usually used to launch RDS instances.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional) Filter the results by a specific instance charge type. Valid values: ` + "`" + `PrePaid` + "`" + ` and ` + "`" + `PostPaid` + "`" + `. Default to ` + "`" + `PostPaid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "multi_zone_ids",
					Description: `A list of zone ids in which the multi zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "multi_zone_ids",
					Description: `A list of zone ids in which the multi zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ddosbgp_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Anti-DDoS Advanced(Ddosbgp) instances available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by the instance name.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) A region of instance.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of instance IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of instance IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of instance names.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of apis. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The instance's id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The instance's remark.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The instance's type.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The instance's region.`,
				},
				resource.Attribute{
					Name:        "base_bandwidth",
					Description: `The instance's base defend bandwidth.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The instance's elastic defend bandwidth.`,
				},
				resource.Attribute{
					Name:        "ip_type",
					Description: `The instance's IP version.`,
				},
				resource.Attribute{
					Name:        "ip_count",
					Description: `The instance's count of ip config.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of instance IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of instance names.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of apis. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The instance's id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The instance's remark.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The instance's type.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The instance's region.`,
				},
				resource.Attribute{
					Name:        "base_bandwidth",
					Description: `The instance's base defend bandwidth.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The instance's elastic defend bandwidth.`,
				},
				resource.Attribute{
					Name:        "ip_type",
					Description: `The instance's IP version.`,
				},
				resource.Attribute{
					Name:        "ip_count",
					Description: `The instance's count of ip config.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ddoscoo_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of BGP-Line Anti-DDoS Pro instances available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by the instance name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of instance IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of instance IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of instance names.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of apis. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The instance's id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The instance's remark.`,
				},
				resource.Attribute{
					Name:        "base_bandwidth",
					Description: `The instance's base defend bandwidth.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The instance's elastic defend bandwidth.`,
				},
				resource.Attribute{
					Name:        "service_bandwidth",
					Description: `The instance's business bandwidth.`,
				},
				resource.Attribute{
					Name:        "port_count",
					Description: `The instance's count of port retransmission config.`,
				},
				resource.Attribute{
					Name:        "domain_count",
					Description: `The instance's count of domain retransmission config.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of instance IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of instance names.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of apis. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The instance's id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The instance's remark.`,
				},
				resource.Attribute{
					Name:        "base_bandwidth",
					Description: `The instance's base defend bandwidth.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `The instance's elastic defend bandwidth.`,
				},
				resource.Attribute{
					Name:        "service_bandwidth",
					Description: `The instance's business bandwidth.`,
				},
				resource.Attribute{
					Name:        "port_count",
					Description: `The instance's count of port retransmission config.`,
				},
				resource.Attribute{
					Name:        "domain_count",
					Description: `The instance's count of domain retransmission config.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_disks",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of disks to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of disks IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by disk name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Disk type. Possible values: ` + "`" + `system` + "`" + ` and ` + "`" + `data` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Optional) Disk category. Possible values: ` + "`" + `cloud` + "`" + ` (basic cloud disk), ` + "`" + `cloud_efficiency` + "`" + ` (ultra cloud disk), ` + "`" + `ephemeral_ssd` + "`" + ` (local SSD cloud disk), ` + "`" + `cloud_ssd` + "`" + ` (SSD cloud disk), and ` + "`" + `cloud_essd` + "`" + ` (ESSD cloud disk).`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `(Optional) Indicate whether the disk is encrypted or not. Possible values: ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) Filter the results by the specified ECS instance ID.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew, Available in 1.57.0+) The Id of resource group which the disk belongs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A map of tags assigned to the disks. It must be in the format: ` + "`" + `` + "`" + `` + "`" + ` data "alicloud_disks" "disks_ds" { tags = { tagKey1 = "tagValue1", tagKey2 = "tagValue2" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "disks",
					Description: `A list of disks. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the disk.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Disk name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Disk description.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the disk belongs to.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone of the disk.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status. Possible values: ` + "`" + `In_use` + "`" + `, ` + "`" + `Available` + "`" + `, ` + "`" + `Attaching` + "`" + `, ` + "`" + `Detaching` + "`" + `, ` + "`" + `Creating` + "`" + ` and ` + "`" + `ReIniting` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Disk type. Possible values: ` + "`" + `system` + "`" + ` and ` + "`" + `data` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `Disk category. Possible values: ` + "`" + `cloud` + "`" + ` (basic cloud disk), ` + "`" + `cloud_efficiency` + "`" + ` (ultra cloud disk), ` + "`" + `ephemeral_ssd` + "`" + ` (local SSD cloud disk), ` + "`" + `cloud_ssd` + "`" + ` (SSD cloud disk), and ` + "`" + `cloud_essd` + "`" + ` (ESSD cloud disk).`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Indicate whether the disk is encrypted or not. Possible values: ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Disk size in GiB.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `ID of the image from which the disk is created. It is null unless the disk is created using an image.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `Snapshot used to create the disk. It is null if no snapshot is used to create the disk.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the related instance. It is ` + "`" + `null` + "`" + ` unless the ` + "`" + `status` + "`" + ` is ` + "`" + `In_use` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Disk creation time.`,
				},
				resource.Attribute{
					Name:        "attached_time",
					Description: `Disk attachment time.`,
				},
				resource.Attribute{
					Name:        "detached_time",
					Description: `Disk detachment time.`,
				},
				resource.Attribute{
					Name:        "expiration_time",
					Description: `Disk expiration time.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the disk.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The Id of resource group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "disks",
					Description: `A list of disks. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the disk.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Disk name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Disk description.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the disk belongs to.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone of the disk.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status. Possible values: ` + "`" + `In_use` + "`" + `, ` + "`" + `Available` + "`" + `, ` + "`" + `Attaching` + "`" + `, ` + "`" + `Detaching` + "`" + `, ` + "`" + `Creating` + "`" + ` and ` + "`" + `ReIniting` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Disk type. Possible values: ` + "`" + `system` + "`" + ` and ` + "`" + `data` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `Disk category. Possible values: ` + "`" + `cloud` + "`" + ` (basic cloud disk), ` + "`" + `cloud_efficiency` + "`" + ` (ultra cloud disk), ` + "`" + `ephemeral_ssd` + "`" + ` (local SSD cloud disk), ` + "`" + `cloud_ssd` + "`" + ` (SSD cloud disk), and ` + "`" + `cloud_essd` + "`" + ` (ESSD cloud disk).`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Indicate whether the disk is encrypted or not. Possible values: ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Disk size in GiB.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `ID of the image from which the disk is created. It is null unless the disk is created using an image.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `Snapshot used to create the disk. It is null if no snapshot is used to create the disk.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the related instance. It is ` + "`" + `null` + "`" + ` unless the ` + "`" + `status` + "`" + ` is ` + "`" + `In_use` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Disk creation time.`,
				},
				resource.Attribute{
					Name:        "attached_time",
					Description: `Disk attachment time.`,
				},
				resource.Attribute{
					Name:        "detached_time",
					Description: `Disk detachment time.`,
				},
				resource.Attribute{
					Name:        "expiration_time",
					Description: `Disk expiration time.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the disk.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The Id of resource group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_dms_enterprise_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of available DMS Enterprise Instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Filter the results by status of the DMS Enterprise Instances. Valid values: ` + "`" + `NORMAL` + "`" + `, ` + "`" + `UNAVAILABLE` + "`" + `, ` + "`" + `UNKNOWN` + "`" + `, ` + "`" + `DELETED` + "`" + `, ` + "`" + `DISABLE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "env_type",
					Description: `(Optional) The type of the environment to which the database instance belongs.`,
				},
				resource.Attribute{
					Name:        "instance_source",
					Description: `(Optional) The source of the database instance.`,
				},
				resource.Attribute{
					Name:        "instance_state",
					Description: `(Optional) The status of the database instance.`,
				},
				resource.Attribute{
					Name:        "net_type",
					Description: `(Optional) The network type of the database instance. Valid values: CLASSIC and VPC. For more information about the valid values, see the description of the RegisterInstance operation.`,
				},
				resource.Attribute{
					Name:        "search_key",
					Description: `(Optional) The keyword used to query database instances.`,
				},
				resource.Attribute{
					Name:        "tid",
					Description: `(Optional) The ID of the tenant in Data Management (DMS) Enterprise.`,
				},
				resource.Attribute{
					Name:        "instance_alias_regex",
					Description: `(Optional) A regex string to filter the results by the DMS Enterprise Instance instance_alias.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of KMS keys. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "data_link_name",
					Description: `The name of the data link for the database instance.`,
				},
				resource.Attribute{
					Name:        "database_password",
					Description: `The logon password of the database instance.`,
				},
				resource.Attribute{
					Name:        "database_user",
					Description: `The logon username of the database instance.`,
				},
				resource.Attribute{
					Name:        "dba_id",
					Description: `The ID of the database administrator (DBA) of the database instance.`,
				},
				resource.Attribute{
					Name:        "dba_nick_name",
					Description: `The nickname of the DBA.`,
				},
				resource.Attribute{
					Name:        "ddl_online",
					Description: `Indicates whether the online data description language (DDL) service was enabled for the database instance.`,
				},
				resource.Attribute{
					Name:        "ecs_instance_id",
					Description: `The ID of the Elastic Compute Service (ECS) instance to which the database instance belongs.`,
				},
				resource.Attribute{
					Name:        "ecs_region",
					Description: `The region where the database instance resides.`,
				},
				resource.Attribute{
					Name:        "env_type",
					Description: `The type of the environment to which the database instance belongs..`,
				},
				resource.Attribute{
					Name:        "export_timeout",
					Description: `The timeout period for exporting the database instance.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The endpoint of the database instance.`,
				},
				resource.Attribute{
					Name:        "instance_alias",
					Description: `The alias of the database instance.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the database instance.`,
				},
				resource.Attribute{
					Name:        "instance_source",
					Description: `The ID of the database instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The ID of the database instance.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The connection port of the database instance.`,
				},
				resource.Attribute{
					Name:        "query_timeout",
					Description: `The timeout period for querying the database instance.`,
				},
				resource.Attribute{
					Name:        "safe_rule_id",
					Description: `The ID of the security rule for the database instance.`,
				},
				resource.Attribute{
					Name:        "sid",
					Description: `The system ID (SID) of the database instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the database instance.`,
				},
				resource.Attribute{
					Name:        "use_dsql",
					Description: `Indicates whether cross-database query was enabled for the database instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of the Virtual Private Cloud (VPC) to which the database instance belongs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instances",
					Description: `A list of KMS keys. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "data_link_name",
					Description: `The name of the data link for the database instance.`,
				},
				resource.Attribute{
					Name:        "database_password",
					Description: `The logon password of the database instance.`,
				},
				resource.Attribute{
					Name:        "database_user",
					Description: `The logon username of the database instance.`,
				},
				resource.Attribute{
					Name:        "dba_id",
					Description: `The ID of the database administrator (DBA) of the database instance.`,
				},
				resource.Attribute{
					Name:        "dba_nick_name",
					Description: `The nickname of the DBA.`,
				},
				resource.Attribute{
					Name:        "ddl_online",
					Description: `Indicates whether the online data description language (DDL) service was enabled for the database instance.`,
				},
				resource.Attribute{
					Name:        "ecs_instance_id",
					Description: `The ID of the Elastic Compute Service (ECS) instance to which the database instance belongs.`,
				},
				resource.Attribute{
					Name:        "ecs_region",
					Description: `The region where the database instance resides.`,
				},
				resource.Attribute{
					Name:        "env_type",
					Description: `The type of the environment to which the database instance belongs..`,
				},
				resource.Attribute{
					Name:        "export_timeout",
					Description: `The timeout period for exporting the database instance.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The endpoint of the database instance.`,
				},
				resource.Attribute{
					Name:        "instance_alias",
					Description: `The alias of the database instance.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the database instance.`,
				},
				resource.Attribute{
					Name:        "instance_source",
					Description: `The ID of the database instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The ID of the database instance.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The connection port of the database instance.`,
				},
				resource.Attribute{
					Name:        "query_timeout",
					Description: `The timeout period for querying the database instance.`,
				},
				resource.Attribute{
					Name:        "safe_rule_id",
					Description: `The ID of the security rule for the database instance.`,
				},
				resource.Attribute{
					Name:        "sid",
					Description: `The system ID (SID) of the database instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the database instance.`,
				},
				resource.Attribute{
					Name:        "use_dsql",
					Description: `Indicates whether cross-database query was enabled for the database instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of the Virtual Private Cloud (VPC) to which the database instance belongs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_dms_enterprise_users",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of available DMS Enterprise Users.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) The role of the user to query.`,
				},
				resource.Attribute{
					Name:        "search_key",
					Description: `(Optional) The keyword used to query users.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the user.`,
				},
				resource.Attribute{
					Name:        "tid",
					Description: `(Optional) The ID of the tenant in DMS Enterprise.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of DMS Enterprise User IDs (UID).`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of DMS Enterprise User IDs (UID).`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `A list of DMS Enterprise Users. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "mobile",
					Description: `The DingTalk number or mobile number of the user.`,
				},
				resource.Attribute{
					Name:        "nick_name",
					Description: `The nickname of the user.`,
				},
				resource.Attribute{
					Name:        "parent_uid",
					Description: `The Alibaba Cloud unique ID (UID) of the parent account if the user corresponds to a Resource Access Management (RAM) user.`,
				},
				resource.Attribute{
					Name:        "role_ids",
					Description: `The list ids of the role that the user plays.`,
				},
				resource.Attribute{
					Name:        "role_names",
					Description: `The list names of the role that he user plays.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the user.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Alibaba Cloud unique ID (UID) of the user.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The ID of the user.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of DMS Enterprise User IDs (UID).`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `A list of DMS Enterprise Users. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "mobile",
					Description: `The DingTalk number or mobile number of the user.`,
				},
				resource.Attribute{
					Name:        "nick_name",
					Description: `The nickname of the user.`,
				},
				resource.Attribute{
					Name:        "parent_uid",
					Description: `The Alibaba Cloud unique ID (UID) of the parent account if the user corresponds to a Resource Access Management (RAM) user.`,
				},
				resource.Attribute{
					Name:        "role_ids",
					Description: `The list ids of the role that the user plays.`,
				},
				resource.Attribute{
					Name:        "role_names",
					Description: `The list names of the role that he user plays.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the user.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Alibaba Cloud unique ID (UID) of the user.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The ID of the user.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_dns_domain_groups",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of groups available to the domain.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_dns_domain_records",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of records available to the domain.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_dns_domain_txt_guid",
			Category:         "Data Sources",
			ShortDescription: `Provides the generation of txt records to realize the retrieval and verification of domain names.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_name",
					Description: `(Required) Verified domain name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Txt verification function. Value:` + "`" + `ADD_SUB_DOMAIN` + "`" + `, ` + "`" + `RETRIEVAL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lang",
					Description: `(Optional) User language.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "rr",
					Description: `Host record.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Record the value.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "rr",
					Description: `Host record.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Record the value.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_dns_domains",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of domains available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_name_regex",
					Description: `(Optional) A regex string to filter results by the domain name.`,
				},
				resource.Attribute{
					Name:        "group_name_regex",
					Description: `(Optional) A regex string to filter results by the group name.`,
				},
				resource.Attribute{
					Name:        "ali_domain",
					Description: `(Optional, type: bool) Specifies whether the domain is from Alibaba Cloud or not.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) Cloud analysis product ID.`,
				},
				resource.Attribute{
					Name:        "version_code",
					Description: `(Optional) Cloud analysis version code.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew, Available in 1.59.0+) The Id of resource group which the dns belongs.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Optional, ForceNew, Available in 1.87.0+) Domain group ID, if not filled, the default is all groups.`,
				},
				resource.Attribute{
					Name:        "key_word",
					Description: `(Optional, ForceNew, Available in 1.87.0+) The keywords are searched according to the ` + "`" + `%KeyWord%` + "`" + ` mode, which is not case sensitive.`,
				},
				resource.Attribute{
					Name:        "lang",
					Description: `(Optional, ForceNew, Available in 1.87.0+) User language.`,
				},
				resource.Attribute{
					Name:        "search_mode",
					Description: `(Optional, ForceNew, Available in 1.87.0+) Search mode, ` + "`" + `LIKE` + "`" + ` fuzzy search, ` + "`" + `EXACT` + "`" + ` exact search.`,
				},
				resource.Attribute{
					Name:        "starmark",
					Description: `(Optional, ForceNew, Available in 1.87.0+) Whether to query the domain name star.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, ForceNew, Available in 1.87.0+) A mapping of tags to assign to the resource. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of domain IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of domain names.`,
				},
				resource.Attribute{
					Name:        "domains",
					Description: `A list of domains. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `ID of the domain.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `Name of the domain.`,
				},
				resource.Attribute{
					Name:        "ali_domain",
					Description: `Indicates whether the domain is an Alibaba Cloud domain.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `Id of group that contains the domain.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `Name of group that contains the domain.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Cloud analysis product ID of the domain.`,
				},
				resource.Attribute{
					Name:        "version_code",
					Description: `Cloud analysis version code of the domain.`,
				},
				resource.Attribute{
					Name:        "puny_code",
					Description: `Punycode of the Chinese domain.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `DNS list of the domain in the analysis system.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The Id of resource group which the dns belongs.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `DNS list of domain names in the resolution system.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Id of resource.`,
				},
				resource.Attribute{
					Name:        "in_black_hole",
					Description: `Whether it is in black hole.`,
				},
				resource.Attribute{
					Name:        "in_clean",
					Description: `Whether it is cleaning.`,
				},
				resource.Attribute{
					Name:        "min_ttl",
					Description: `Minimum TTL.`,
				},
				resource.Attribute{
					Name:        "version_code",
					Description: `Cloud resolution version ID.`,
				},
				resource.Attribute{
					Name:        "record_line_tree_json",
					Description: `Tree-like analytical line list.`,
				},
				resource.Attribute{
					Name:        "region_lines",
					Description: `Whether it is a regional route.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `The Id of resource group which the dns belongs.`,
				},
				resource.Attribute{
					Name:        "slave_dns",
					Description: `Whether to allow auxiliary dns.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of domain IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of domain names.`,
				},
				resource.Attribute{
					Name:        "domains",
					Description: `A list of domains. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `ID of the domain.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `Name of the domain.`,
				},
				resource.Attribute{
					Name:        "ali_domain",
					Description: `Indicates whether the domain is an Alibaba Cloud domain.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `Id of group that contains the domain.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `Name of group that contains the domain.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Cloud analysis product ID of the domain.`,
				},
				resource.Attribute{
					Name:        "version_code",
					Description: `Cloud analysis version code of the domain.`,
				},
				resource.Attribute{
					Name:        "puny_code",
					Description: `Punycode of the Chinese domain.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `DNS list of the domain in the analysis system.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The Id of resource group which the dns belongs.`,
				},
				resource.Attribute{
					Name:        "dns_servers",
					Description: `DNS list of domain names in the resolution system.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Id of resource.`,
				},
				resource.Attribute{
					Name:        "in_black_hole",
					Description: `Whether it is in black hole.`,
				},
				resource.Attribute{
					Name:        "in_clean",
					Description: `Whether it is cleaning.`,
				},
				resource.Attribute{
					Name:        "min_ttl",
					Description: `Minimum TTL.`,
				},
				resource.Attribute{
					Name:        "version_code",
					Description: `Cloud resolution version ID.`,
				},
				resource.Attribute{
					Name:        "record_line_tree_json",
					Description: `Tree-like analytical line list.`,
				},
				resource.Attribute{
					Name:        "region_lines",
					Description: `Whether it is a regional route.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `The Id of resource group which the dns belongs.`,
				},
				resource.Attribute{
					Name:        "slave_dns",
					Description: `Whether to allow auxiliary dns.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_dns_groups",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of groups available to the dns.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by group name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, Available 1.52.2+) A list of group IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of group IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of group names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `Id of the group.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `Name of the group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of group IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of group names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `Id of the group.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `Name of the group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_dns_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of instances available to the DNS.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of instance IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of instance IDs.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the instance.`,
				},
				resource.Attribute{
					Name:        "domain_numbers",
					Description: `Number of domain names bound.`,
				},
				resource.Attribute{
					Name:        "dns_security",
					Description: `DNS security level.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Id of the instance resource.`,
				},
				resource.Attribute{
					Name:        "version_code",
					Description: `Paid package version.`,
				},
				resource.Attribute{
					Name:        "version_name",
					Description: `Paid package version name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of instance IDs.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the instance.`,
				},
				resource.Attribute{
					Name:        "domain_numbers",
					Description: `Number of domain names bound.`,
				},
				resource.Attribute{
					Name:        "dns_security",
					Description: `DNS security level.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Id of the instance resource.`,
				},
				resource.Attribute{
					Name:        "version_code",
					Description: `Paid package version.`,
				},
				resource.Attribute{
					Name:        "version_name",
					Description: `Paid package version name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_dns_records",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of records available to the dns.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_name",
					Description: `(Required) The domain name associated to the records.`,
				},
				resource.Attribute{
					Name:        "host_record_regex",
					Description: `(Optional) Host record regex.`,
				},
				resource.Attribute{
					Name:        "value_regex",
					Description: `(Optional) Host record value regex.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Record type. Valid items are ` + "`" + `A` + "`" + `, ` + "`" + `NS` + "`" + `, ` + "`" + `MX` + "`" + `, ` + "`" + `TXT` + "`" + `, ` + "`" + `CNAME` + "`" + `, ` + "`" + `SRV` + "`" + `, ` + "`" + `AAAA` + "`" + `, ` + "`" + `REDIRECT_URL` + "`" + `, ` + "`" + `FORWORD_URL` + "`" + ` .`,
				},
				resource.Attribute{
					Name:        "line",
					Description: `(Optional) ISP line. Valid items are ` + "`" + `default` + "`" + `, ` + "`" + `telecom` + "`" + `, ` + "`" + `unicom` + "`" + `, ` + "`" + `mobile` + "`" + `, ` + "`" + `oversea` + "`" + `, ` + "`" + `edu` + "`" + `, ` + "`" + `drpeng` + "`" + `, ` + "`" + `btvn` + "`" + `, .etc. For checking all resolution lines enumeration please visit [Alibaba Cloud DNS doc](https://www.alibabacloud.com/help/doc-detail/34339.htm)`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Record status. Valid items are ` + "`" + `ENABLE` + "`" + ` and ` + "`" + `DISABLE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "is_locked",
					Description: `(Optional, type: bool) Whether the record is locked or not.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, Available 1.52.2+) A list of record IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of record IDs.`,
				},
				resource.Attribute{
					Name:        "urls",
					Description: `A list of entire URLs. Each item format as ` + "`" + `<host_record>.<domain_name>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "records",
					Description: `A list of records. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "record_id",
					Description: `ID of the record.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `Name of the domain the record belongs to.`,
				},
				resource.Attribute{
					Name:        "host_record",
					Description: `Host record of the domain.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Host record value of the domain.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the record.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `TTL of the record.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Priority of the ` + "`" + `MX` + "`" + ` record.`,
				},
				resource.Attribute{
					Name:        "line",
					Description: `ISP line of the record.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the record.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Indicates whether the record is locked.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of record IDs.`,
				},
				resource.Attribute{
					Name:        "urls",
					Description: `A list of entire URLs. Each item format as ` + "`" + `<host_record>.<domain_name>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "records",
					Description: `A list of records. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "record_id",
					Description: `ID of the record.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `Name of the domain the record belongs to.`,
				},
				resource.Attribute{
					Name:        "host_record",
					Description: `Host record of the domain.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Host record value of the domain.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the record.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `TTL of the record.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Priority of the ` + "`" + `MX` + "`" + ` record.`,
				},
				resource.Attribute{
					Name:        "line",
					Description: `ISP line of the record.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the record.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `Indicates whether the record is locked.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_resolution_lines",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of domains available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_name",
					Description: `(Optional) Domain Name.`,
				},
				resource.Attribute{
					Name:        "line_codes",
					Description: `(Optional) A list of lines codes.`,
				},
				resource.Attribute{
					Name:        "line_display_names",
					Description: `(Optional) A list of line display names.`,
				},
				resource.Attribute{
					Name:        "user_client_ip",
					Description: `(Optional) The ip of user client.`,
				},
				resource.Attribute{
					Name:        "lang",
					Description: `(Optional) language.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "line_codes",
					Description: `A list of lines codes.`,
				},
				resource.Attribute{
					Name:        "line_display_names",
					Description: `A list of line display names.`,
				},
				resource.Attribute{
					Name:        "lines",
					Description: `A list of cloud resolution line. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "line_codes",
					Description: `Line code.`,
				},
				resource.Attribute{
					Name:        "line_display_name",
					Description: `Line display name.`,
				},
				resource.Attribute{
					Name:        "line_name",
					Description: `Line name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "line_codes",
					Description: `A list of lines codes.`,
				},
				resource.Attribute{
					Name:        "line_display_names",
					Description: `A list of line display names.`,
				},
				resource.Attribute{
					Name:        "lines",
					Description: `A list of cloud resolution line. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "line_codes",
					Description: `Line code.`,
				},
				resource.Attribute{
					Name:        "line_display_name",
					Description: `Line display name.`,
				},
				resource.Attribute{
					Name:        "line_name",
					Description: `Line name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_drds_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a collection of DRDS instances according to the specified filters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional, Deprecated) A regex string to filter results by instance description. It is deprecated since v1.91.0 and will be removed in a future release, please use 'description_regex' instead.`,
				},
				resource.Attribute{
					Name:        "description_regex",
					Description: `(Optional, Available in 1.91.0+) A regex string to filter results by instance description.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of DRDS instance IDs. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of DRDS instance IDs.`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `A list of DRDS descriptions.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of DRDS instances.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the DRDS instance.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The DRDS instance description.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the RDS instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The DRDS Instance type.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the instance.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `` + "`" + `Classic` + "`" + ` for public classic network or ` + "`" + `VPC` + "`" + ` for private network.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `Zone ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The DRDS Instance version.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of DRDS instance IDs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of DRDS instance IDs.`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `A list of DRDS descriptions.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of DRDS instances.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the DRDS instance.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The DRDS instance description.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the RDS instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The DRDS Instance type.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the instance.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `` + "`" + `Classic` + "`" + ` for public classic network or ` + "`" + `VPC` + "`" + ` for private network.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `Zone ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The DRDS Instance version.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of DRDS instance IDs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_eci_image_caches",
			Category:         "Data Sources",
			ShortDescription: `Provides a collection of ECI Image Cache to the specified filters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list ids of ECI Image Cache.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by the image cache name.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional) Find the mirror cache containing it according to the image name.`,
				},
				resource.Attribute{
					Name:        "image_cache_name",
					Description: `(Optional) The name of ECI Image Cache.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The id of snapshot.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of ECI Image Cache.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list ids of ECI Image Cache.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of ECI Image Cache names.`,
				},
				resource.Attribute{
					Name:        "caches",
					Description: `A list of caches. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the ECI Image Cache.`,
				},
				resource.Attribute{
					Name:        "container_group_id",
					Description: `The id of container group.`,
				},
				resource.Attribute{
					Name:        "expire_date_time",
					Description: `The time of expired.`,
				},
				resource.Attribute{
					Name:        "image_cache_id",
					Description: `The id of the ECI Image Cache.`,
				},
				resource.Attribute{
					Name:        "image_cache_name",
					Description: `The name of the ECI Image Cache.`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `The list of cached images.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `The progress of ECI Image Cache.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The id of snapshot.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of ECI Image Cache.`,
				},
				resource.Attribute{
					Name:        "events",
					Description: `Image cache pulls image event information.`,
				},
				resource.Attribute{
					Name:        "first_timestamp",
					Description: `Start time.`,
				},
				resource.Attribute{
					Name:        "last_timestamp",
					Description: `End time.`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `Number of events.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of event.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of event.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list ids of ECI Image Cache.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of ECI Image Cache names.`,
				},
				resource.Attribute{
					Name:        "caches",
					Description: `A list of caches. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the ECI Image Cache.`,
				},
				resource.Attribute{
					Name:        "container_group_id",
					Description: `The id of container group.`,
				},
				resource.Attribute{
					Name:        "expire_date_time",
					Description: `The time of expired.`,
				},
				resource.Attribute{
					Name:        "image_cache_id",
					Description: `The id of the ECI Image Cache.`,
				},
				resource.Attribute{
					Name:        "image_cache_name",
					Description: `The name of the ECI Image Cache.`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `The list of cached images.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `The progress of ECI Image Cache.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The id of snapshot.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of ECI Image Cache.`,
				},
				resource.Attribute{
					Name:        "events",
					Description: `Image cache pulls image event information.`,
				},
				resource.Attribute{
					Name:        "first_timestamp",
					Description: `Start time.`,
				},
				resource.Attribute{
					Name:        "last_timestamp",
					Description: `End time.`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `Number of events.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of event.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of event.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ecs_dedicated_hosts",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of available ECS Dedicated Hosts.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of ECS Dedicated Host ids.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by the ECS Dedicated Host name.`,
				},
				resource.Attribute{
					Name:        "dedicated_host_id",
					Description: `(Optional) The ID of ECS Dedicated Host.`,
				},
				resource.Attribute{
					Name:        "dedicated_host_name",
					Description: `(Optional) The name of ECS Dedicated Host.`,
				},
				resource.Attribute{
					Name:        "dedicated_host_type",
					Description: `(Optional) The type of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional) The ID of the resource group to which the ECS Dedicated Host belongs.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the ECS Dedicated Host. validate value: ` + "`" + `Available` + "`" + `, ` + "`" + `Creating` + "`" + `, ` + "`" + `PermanentFailure` + "`" + `, ` + "`" + `Released` + "`" + `, ` + "`" + `UnderAssessment` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional) The zone ID of the ECS Dedicated Host.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) Save the result to the file. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of ECS Dedicated Host ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of ECS Dedicated Host names.`,
				},
				resource.Attribute{
					Name:        "hosts",
					Description: `A list of ECS Dedicated Hosts. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the ECS Dedicated Host.`,
				},
				resource.Attribute{
					Name:        "action_on_maintenance",
					Description: `The policy used to migrate the instances from the dedicated host when the dedicated host fails or needs to be repaired online.`,
				},
				resource.Attribute{
					Name:        "auto_placement",
					Description: `Specifies whether to add the dedicated host to the resource pool for automatic deployment.`,
				},
				resource.Attribute{
					Name:        "auto_release_time",
					Description: `The automatic release time of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "cores",
					Description: `A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "dedicated_host_id",
					Description: `ID of the ECS Dedicated Host.`,
				},
				resource.Attribute{
					Name:        "dedicated_host_name",
					Description: `The name of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "dedicated_host_type",
					Description: `The type of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `The expiration time of the subscription dedicated host.`,
				},
				resource.Attribute{
					Name:        "gpu_spec",
					Description: `The GPU model.`,
				},
				resource.Attribute{
					Name:        "machine_id",
					Description: `The machine code of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "payment_type",
					Description: `The billing method of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "physical_gpus",
					Description: `The number of physical GPUs.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The ID of the resource group to which the dedicated host belongs.`,
				},
				resource.Attribute{
					Name:        "sale_cycle",
					Description: `The unit of the subscription billing method.`,
				},
				resource.Attribute{
					Name:        "sockets",
					Description: `The number of physical CPUs.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The service status of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "supported_instance_types_list",
					Description: `The list of ECS instanc`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of ECS Dedicated Host ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of ECS Dedicated Host names.`,
				},
				resource.Attribute{
					Name:        "hosts",
					Description: `A list of ECS Dedicated Hosts. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the ECS Dedicated Host.`,
				},
				resource.Attribute{
					Name:        "action_on_maintenance",
					Description: `The policy used to migrate the instances from the dedicated host when the dedicated host fails or needs to be repaired online.`,
				},
				resource.Attribute{
					Name:        "auto_placement",
					Description: `Specifies whether to add the dedicated host to the resource pool for automatic deployment.`,
				},
				resource.Attribute{
					Name:        "auto_release_time",
					Description: `The automatic release time of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "cores",
					Description: `A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "dedicated_host_id",
					Description: `ID of the ECS Dedicated Host.`,
				},
				resource.Attribute{
					Name:        "dedicated_host_name",
					Description: `The name of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "dedicated_host_type",
					Description: `The type of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `The expiration time of the subscription dedicated host.`,
				},
				resource.Attribute{
					Name:        "gpu_spec",
					Description: `The GPU model.`,
				},
				resource.Attribute{
					Name:        "machine_id",
					Description: `The machine code of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "payment_type",
					Description: `The billing method of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "physical_gpus",
					Description: `The number of physical GPUs.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The ID of the resource group to which the dedicated host belongs.`,
				},
				resource.Attribute{
					Name:        "sale_cycle",
					Description: `The unit of the subscription billing method.`,
				},
				resource.Attribute{
					Name:        "sockets",
					Description: `The number of physical CPUs.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The service status of the dedicated host.`,
				},
				resource.Attribute{
					Name:        "supported_instance_types_list",
					Description: `The list of ECS instanc`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_edas_applications",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of EDAS applications available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) An ids string to filter results by the application id.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by the application name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of application IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of applications names.`,
				},
				resource.Attribute{
					Name:        "applications",
					Description: `A list of applications.`,
				},
				resource.Attribute{
					Name:        "app_name",
					Description: `The name of your EDAS application. Only letters '-' '_' and numbers are allowed. The length cannot exceed 36 characters.`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `The ID of the application that you want to deploy.`,
				},
				resource.Attribute{
					Name:        "application_type",
					Description: `The type of the package for the deployment of the application that you want to create. The valid values are: WAR and JAR. We strongly recommend you to set this parameter when creating the application.`,
				},
				resource.Attribute{
					Name:        "build_package_id",
					Description: `The package ID of Enterprise Distributed Application Service (EDAS) Container.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `The ID of the cluster that you want to create the application.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `The type of the cluster that you want to create. Valid values: 1: Swarm cluster. 2: ECS cluster. 3: Kubernates cluster.`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `Number of instances.`,
				},
				resource.Attribute{
					Name:        "running_instance_count",
					Description: `Number of running instances.`,
				},
				resource.Attribute{
					Name:        "health_check_url",
					Description: `The URL for health checking of the application.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Application's creation time.`,
				},
				resource.Attribute{
					Name:        "slb_id",
					Description: `The ID of the SLB instance that is going to be bound.`,
				},
				resource.Attribute{
					Name:        "slb_ip",
					Description: `The IP address that is allocated to the bound SLB instance.`,
				},
				resource.Attribute{
					Name:        "slb_port",
					Description: `The port of intranet SLB.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The ID of the namespace the application belongs to.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of application IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of applications names.`,
				},
				resource.Attribute{
					Name:        "applications",
					Description: `A list of applications.`,
				},
				resource.Attribute{
					Name:        "app_name",
					Description: `The name of your EDAS application. Only letters '-' '_' and numbers are allowed. The length cannot exceed 36 characters.`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `The ID of the application that you want to deploy.`,
				},
				resource.Attribute{
					Name:        "application_type",
					Description: `The type of the package for the deployment of the application that you want to create. The valid values are: WAR and JAR. We strongly recommend you to set this parameter when creating the application.`,
				},
				resource.Attribute{
					Name:        "build_package_id",
					Description: `The package ID of Enterprise Distributed Application Service (EDAS) Container.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `The ID of the cluster that you want to create the application.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `The type of the cluster that you want to create. Valid values: 1: Swarm cluster. 2: ECS cluster. 3: Kubernates cluster.`,
				},
				resource.Attribute{
					Name:        "instance_count",
					Description: `Number of instances.`,
				},
				resource.Attribute{
					Name:        "running_instance_count",
					Description: `Number of running instances.`,
				},
				resource.Attribute{
					Name:        "health_check_url",
					Description: `The URL for health checking of the application.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Application's creation time.`,
				},
				resource.Attribute{
					Name:        "slb_id",
					Description: `The ID of the SLB instance that is going to be bound.`,
				},
				resource.Attribute{
					Name:        "slb_ip",
					Description: `The IP address that is allocated to the bound SLB instance.`,
				},
				resource.Attribute{
					Name:        "slb_port",
					Description: `The port of intranet SLB.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The ID of the namespace the application belongs to.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_edas_clusters",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of EDAS clusters available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "logical_region_id",
					Description: `(Required) ID of the namespace in EDAS.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) An ids string to filter results by the cluster id.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by the cluster name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of cluster IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of cluster names.`,
				},
				resource.Attribute{
					Name:        "clusters",
					Description: `A list of clusters.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `The ID of the cluster that you want to create the application.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `The name of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `The type of the cluster, Valid values: 1: Swarm cluster. 2: ECS cluster. 3: Kubernates cluster.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Cluster's creation time.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The time when the cluster was last updated.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `The total number of CPUs in the cluster.`,
				},
				resource.Attribute{
					Name:        "cpu_used",
					Description: `The number of used CPUs in the cluster.`,
				},
				resource.Attribute{
					Name:        "mem",
					Description: `The total amount of memory in the cluser. Unit: MB.`,
				},
				resource.Attribute{
					Name:        "mem_used",
					Description: `The amount of used memory in the cluser. Unit: MB.`,
				},
				resource.Attribute{
					Name:        "network_mode",
					Description: `The network type of the cluster. Valid values: 1: classic network. 2: VPC.`,
				},
				resource.Attribute{
					Name:        "node_num",
					Description: `The number of the Elastic Compute Service (ECS) instances that are deployed to the cluster.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of the Virtual Private Cloud (VPC) for the cluster.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The ID of the namespace the application belongs to.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of cluster IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of cluster names.`,
				},
				resource.Attribute{
					Name:        "clusters",
					Description: `A list of clusters.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `The ID of the cluster that you want to create the application.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `The name of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `The type of the cluster, Valid values: 1: Swarm cluster. 2: ECS cluster. 3: Kubernates cluster.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Cluster's creation time.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The time when the cluster was last updated.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `The total number of CPUs in the cluster.`,
				},
				resource.Attribute{
					Name:        "cpu_used",
					Description: `The number of used CPUs in the cluster.`,
				},
				resource.Attribute{
					Name:        "mem",
					Description: `The total amount of memory in the cluser. Unit: MB.`,
				},
				resource.Attribute{
					Name:        "mem_used",
					Description: `The amount of used memory in the cluser. Unit: MB.`,
				},
				resource.Attribute{
					Name:        "network_mode",
					Description: `The network type of the cluster. Valid values: 1: classic network. 2: VPC.`,
				},
				resource.Attribute{
					Name:        "node_num",
					Description: `The number of the Elastic Compute Service (ECS) instances that are deployed to the cluster.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of the Virtual Private Cloud (VPC) for the cluster.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The ID of the namespace the application belongs to.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_edas_deploy_groups",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of EDAS deploy groups available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "app_id",
					Description: `(Required) ID of the EDAS application.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) An ids string to filter results by the deploy group id.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by the deploy group name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of deploy group IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of deploy group names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of consumer group ids.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `The ID of the instance group.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `The name of the instance group. The length cannot exceed 64 characters.`,
				},
				resource.Attribute{
					Name:        "group_type",
					Description: `The type of the instance group. Valid values: 0: Default group. 1: Phased release is disabled for traffic management. 2: Phased release is enabled for traffic management.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the instance group was created.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The time when the instance group was updated.`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `The ID of the application that you want to deploy.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `The ID of the cluster that you want to create the application.`,
				},
				resource.Attribute{
					Name:        "package_version_id",
					Description: `The version of the deployment package for the instance group that was created.`,
				},
				resource.Attribute{
					Name:        "app_version_id",
					Description: `The version of the deployment package for the application.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of deploy group IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of deploy group names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of consumer group ids.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `The ID of the instance group.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `The name of the instance group. The length cannot exceed 64 characters.`,
				},
				resource.Attribute{
					Name:        "group_type",
					Description: `The type of the instance group. Valid values: 0: Default group. 1: Phased release is disabled for traffic management. 2: Phased release is enabled for traffic management.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the instance group was created.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `The time when the instance group was updated.`,
				},
				resource.Attribute{
					Name:        "app_id",
					Description: `The ID of the application that you want to deploy.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `The ID of the cluster that you want to create the application.`,
				},
				resource.Attribute{
					Name:        "package_version_id",
					Description: `The version of the deployment package for the instance group that was created.`,
				},
				resource.Attribute{
					Name:        "app_version_id",
					Description: `The version of the deployment package for the application.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_eips",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of EIP owned by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of EIP IDs.`,
				},
				resource.Attribute{
					Name:        "ip_addresses",
					Description: `(Optional) A list of EIP public IP addresses.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.55.3+) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "in_use",
					Description: `(Deprecated) Deprecated since the version 1.8.0 of this provider.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew, Available in 1.58.0+) The Id of resource group which the eips belongs. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of EIP IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `(Optional) A list of EIP names.`,
				},
				resource.Attribute{
					Name:        "eips",
					Description: `A list of EIPs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the EIP.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `EIP status. Possible values are: ` + "`" + `Associating` + "`" + `, ` + "`" + `Unassociating` + "`" + `, ` + "`" + `InUse` + "`" + ` and ` + "`" + `Available` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `Public IP Address of the the EIP.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `EIP internet max bandwidth in Mbps.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `EIP internet charge type.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the instance that is being bound.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The instance type of that the EIP is bound.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Time of creation.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The Id of resource group which the eips belongs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of EIP IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `(Optional) A list of EIP names.`,
				},
				resource.Attribute{
					Name:        "eips",
					Description: `A list of EIPs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the EIP.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `EIP status. Possible values are: ` + "`" + `Associating` + "`" + `, ` + "`" + `Unassociating` + "`" + `, ` + "`" + `InUse` + "`" + ` and ` + "`" + `Available` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `Public IP Address of the the EIP.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `EIP internet max bandwidth in Mbps.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `EIP internet charge type.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the instance that is being bound.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The instance type of that the EIP is bound.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Time of creation.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The Id of resource group which the eips belongs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_elasticsearch_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a collection of Elasticsearch instances according to the specified filters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description_regex",
					Description: `(Optional) A regex string to apply to the instance description.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, Available 1.52.1+) A list of Elasticsearch instance IDs.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Elasticsearch version. Options are ` + "`" + `5.5.3_with_X-Pack` + "`" + `, ` + "`" + `6.3.2_with_X-Pack` + "`" + ` and ` + "`" + `6.7.0_with_X-Pack` + "`" + `. If no value is specified, all versions are returned.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available 1.74.0+) A map of tags assigned to instances.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of Elasticsearch instance IDs.`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `A list of Elasticsearch instance descriptions.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of Elasticsearch instances. Its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Elasticsearch instance.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Elasticsearch instance.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `Billing method. Value options: ` + "`" + `PostPaid` + "`" + ` for Pay-As-You-Go and ` + "`" + `PrePaid` + "`" + ` for subscription.`,
				},
				resource.Attribute{
					Name:        "data_node_amount",
					Description: `The Elasticsearch cluster's data node quantity, between 2 and 50.`,
				},
				resource.Attribute{
					Name:        "data_node_spec",
					Description: `The data node specifications of the elasticsearch instance.`,
				},
				resource.Attribute{
					Name:        "data_node_disk_size",
					Description: `The single data node storage space. Unit: GB.`,
				},
				resource.Attribute{
					Name:        "data_node_disk_type",
					Description: `The data node disk type. Included values: ` + "`" + `cloud_ssd` + "`" + ` and ` + "`" + `cloud_efficiency` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `VSwitch ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Elasticsearch version includes ` + "`" + `5.5.3_with_X-Pack` + "`" + `, ` + "`" + `6.3.2_with_X-Pack` + "`" + ` and ` + "`" + `6.7.0_with_X-Pack` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cerated_at",
					Description: `The creation time of the instance. It's a GTM format, such as: "2019-01-08T15:50:50.623Z".`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The last modified time of the instance. It's a GMT format, such as: "2019-01-08T15:50:50.623Z".`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance. It includes ` + "`" + `active` + "`" + `, ` + "`" + `activating` + "`" + `, ` + "`" + `inactive` + "`" + ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of Elasticsearch instance IDs.`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `A list of Elasticsearch instance descriptions.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of Elasticsearch instances. Its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Elasticsearch instance.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Elasticsearch instance.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `Billing method. Value options: ` + "`" + `PostPaid` + "`" + ` for Pay-As-You-Go and ` + "`" + `PrePaid` + "`" + ` for subscription.`,
				},
				resource.Attribute{
					Name:        "data_node_amount",
					Description: `The Elasticsearch cluster's data node quantity, between 2 and 50.`,
				},
				resource.Attribute{
					Name:        "data_node_spec",
					Description: `The data node specifications of the elasticsearch instance.`,
				},
				resource.Attribute{
					Name:        "data_node_disk_size",
					Description: `The single data node storage space. Unit: GB.`,
				},
				resource.Attribute{
					Name:        "data_node_disk_type",
					Description: `The data node disk type. Included values: ` + "`" + `cloud_ssd` + "`" + ` and ` + "`" + `cloud_efficiency` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `VSwitch ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Elasticsearch version includes ` + "`" + `5.5.3_with_X-Pack` + "`" + `, ` + "`" + `6.3.2_with_X-Pack` + "`" + ` and ` + "`" + `6.7.0_with_X-Pack` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cerated_at",
					Description: `The creation time of the instance. It's a GTM format, such as: "2019-01-08T15:50:50.623Z".`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The last modified time of the instance. It's a GMT format, such as: "2019-01-08T15:50:50.623Z".`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance. It includes ` + "`" + `active` + "`" + `, ` + "`" + `activating` + "`" + `, ` + "`" + `inactive` + "`" + ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_elasticsearch_zones",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of availability zones for Elasticsearch that can be used by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "multi",
					Description: `(Optional) Indicate whether the zones can be used in a multi AZ configuration. Default to ` + "`" + `false` + "`" + `. Multi AZ is usually used to launch Elasticsearch instances.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "multi_zone_ids",
					Description: `A list of zone ids in which the multi zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "multi_zone_ids",
					Description: `A list of zone ids in which the multi zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_emr_disk_types",
			Category:         "Data Sources",
			ShortDescription: `Provides a collection of data disk and system disk types when create emr cluster according to the specified filters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "destination_resource",
					Description: `(Required) The destination resource of emr cluster instance`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Required) Filter the results by charge type. Valid values: ` + "`" + `PrePaid` + "`" + ` and ` + "`" + `PostPaid` + "`" + `. Default to ` + "`" + `PostPaid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Required) The cluster type of the emr cluster instance. Possible values: ` + "`" + `HADOOP` + "`" + `, ` + "`" + `KAFKA` + "`" + `, ` + "`" + `ZOOKEEPER` + "`" + `, ` + "`" + `DRUID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required) The ecs instance type of create emr cluster instance.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional) The Zone to create emr cluster instance.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of data disk and system disk type IDs.`,
				},
				resource.Attribute{
					Name:        "types",
					Description: `A list of emr instance types. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the data disk or system disk`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `The mininum value of the data disk to supported the specific instance type`,
				},
				resource.Attribute{
					Name:        "max",
					Description: `The maximum value of the data disk to supported the specific instance type`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of data disk and system disk type IDs.`,
				},
				resource.Attribute{
					Name:        "types",
					Description: `A list of emr instance types. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value of the data disk or system disk`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `The mininum value of the data disk to supported the specific instance type`,
				},
				resource.Attribute{
					Name:        "max",
					Description: `The maximum value of the data disk to supported the specific instance type`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_emr_instance_types",
			Category:         "Data Sources",
			ShortDescription: `Provides a collection of ecs instance types when create emr cluster according to the specified filters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "destination_resource",
					Description: `(Required) The destination resource of emr cluster instance`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Required) Filter the results by charge type. Valid values: ` + "`" + `PrePaid` + "`" + ` and ` + "`" + `PostPaid` + "`" + `. Default to ` + "`" + `PostPaid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Required) The cluster type of the emr cluster instance. Possible values: ` + "`" + `HADOOP` + "`" + `, ` + "`" + `KAFKA` + "`" + `, ` + "`" + `ZOOKEEPER` + "`" + `, ` + "`" + `DRUID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional, Available in 1.71.2+) Filter the specific ecs instance type to create emr cluster.`,
				},
				resource.Attribute{
					Name:        "support_local_storage",
					Description: `(Optional, Available in 1.61.0+) Whether the current storage disk is local or not.`,
				},
				resource.Attribute{
					Name:        "support_node_type",
					Description: `(Optional, Available in 1.63.0+) The specific supported node type list. Possible values may be any one or combination of these: ["MASTER", "CORE", "TASK", "GATEWAY"]`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional,Available in 1.69.0+) The supported resources of specific zoneId.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of emr instance types IDs.`,
				},
				resource.Attribute{
					Name:        "types",
					Description: `A list of emr instance types. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the instance type.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `The available zone id in Alibaba Cloud account`,
				},
				resource.Attribute{
					Name:        "local_storage_capacity",
					Description: `Local capacity of the applied ecs instance for emr cluster. Unit: GB.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of emr instance types IDs.`,
				},
				resource.Attribute{
					Name:        "types",
					Description: `A list of emr instance types. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the instance type.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `The available zone id in Alibaba Cloud account`,
				},
				resource.Attribute{
					Name:        "local_storage_capacity",
					Description: `Local capacity of the applied ecs instance for emr cluster. Unit: GB.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_emr_main_versions",
			Category:         "Data Sources",
			ShortDescription: `Provides a collection of emr main versions when create emr cluster according to the specified filters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "emr_version",
					Description: `(Optional) The version of the emr cluster instance. Possible values: ` + "`" + `EMR-4.0.0` + "`" + `, ` + "`" + `EMR-3.23.0` + "`" + `, ` + "`" + `EMR-3.22.0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Optional, Available in 1.70.1+) The supported clusterType of this emr version. Possible values may be any one or combination of these: ["HADOOP", "DRUID", "KAFKA", "ZOOKEEPER", "FLINK", "CLICKHOUSE"]`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of emr instance types IDs.`,
				},
				resource.Attribute{
					Name:        "main_versions",
					Description: `A list of versions of the emr cluster instance. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "emr_version",
					Description: `The version of the emr cluster instance.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The image id of the emr cluster instance.`,
				},
				resource.Attribute{
					Name:        "cluster_types",
					Description: `A list of cluster types the emr cluster supported. Possible values: ` + "`" + `HADOOP` + "`" + `, ` + "`" + `ZOOKEEPER` + "`" + `, ` + "`" + `KAFKA` + "`" + `, ` + "`" + `DRUID` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of emr instance types IDs.`,
				},
				resource.Attribute{
					Name:        "main_versions",
					Description: `A list of versions of the emr cluster instance. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "emr_version",
					Description: `The version of the emr cluster instance.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The image id of the emr cluster instance.`,
				},
				resource.Attribute{
					Name:        "cluster_types",
					Description: `A list of cluster types the emr cluster supported. Possible values: ` + "`" + `HADOOP` + "`" + `, ` + "`" + `ZOOKEEPER` + "`" + `, ` + "`" + `KAFKA` + "`" + `, ` + "`" + `DRUID` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ess_alarms",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of alarms available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Optional) Scaling group id the alarms belong to.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting alarms by name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of alarm IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "metric_type",
					Description: `(Optional) The type for the alarm's associated metric. Supported value: system, custom. "system" means the metric data is collected by Aliyun Cloud Monitor Service(CMS), "custom" means the metric data is upload to CMS by users. Defaults to system. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of alarm ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of alarm names.`,
				},
				resource.Attribute{
					Name:        "alarms",
					Description: `A list of alarms. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of alarm.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name for ess alarm.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description for the alarm.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `Whether to enable specific ess alarm.`,
				},
				resource.Attribute{
					Name:        "alarm_actions",
					Description: `The list of actions to execute when this alarm transition into an ALARM state. Each action is specified as ess scaling rule ari.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `The scaling group associated with this alarm.`,
				},
				resource.Attribute{
					Name:        "metric_type",
					Description: `The type for the alarm's associated metric.`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `The name for the alarm's associated metric. See [Block_metricNames_and_dimensions](#block-metricnames_and_dimensions) below for details.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `The period in seconds over which the specified statistic is applied.`,
				},
				resource.Attribute{
					Name:        "statistics",
					Description: `The statistic to apply to the alarm's associated metric.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `The value against which the specified statistics is compared.`,
				},
				resource.Attribute{
					Name:        "comparison_operator",
					Description: `The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand.`,
				},
				resource.Attribute{
					Name:        "evaluation_count",
					Description: `The number of times that needs to satisfies comparison condition before transition into ALARM state.`,
				},
				resource.Attribute{
					Name:        "cloud_monitor_group_id",
					Description: `Defines the application group id defined by CMS which is assigned when you upload custom metric to CMS, only available for custom metirc.`,
				},
				resource.Attribute{
					Name:        "dimensions",
					Description: `The dimension map for the alarm's associated metric.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of alarm task.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of alarm ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of alarm names.`,
				},
				resource.Attribute{
					Name:        "alarms",
					Description: `A list of alarms. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of alarm.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name for ess alarm.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description for the alarm.`,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `Whether to enable specific ess alarm.`,
				},
				resource.Attribute{
					Name:        "alarm_actions",
					Description: `The list of actions to execute when this alarm transition into an ALARM state. Each action is specified as ess scaling rule ari.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `The scaling group associated with this alarm.`,
				},
				resource.Attribute{
					Name:        "metric_type",
					Description: `The type for the alarm's associated metric.`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `The name for the alarm's associated metric. See [Block_metricNames_and_dimensions](#block-metricnames_and_dimensions) below for details.`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `The period in seconds over which the specified statistic is applied.`,
				},
				resource.Attribute{
					Name:        "statistics",
					Description: `The statistic to apply to the alarm's associated metric.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `The value against which the specified statistics is compared.`,
				},
				resource.Attribute{
					Name:        "comparison_operator",
					Description: `The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand.`,
				},
				resource.Attribute{
					Name:        "evaluation_count",
					Description: `The number of times that needs to satisfies comparison condition before transition into ALARM state.`,
				},
				resource.Attribute{
					Name:        "cloud_monitor_group_id",
					Description: `Defines the application group id defined by CMS which is assigned when you upload custom metric to CMS, only available for custom metirc.`,
				},
				resource.Attribute{
					Name:        "dimensions",
					Description: `The dimension map for the alarm's associated metric.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of alarm task.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ess_lifecycle_hooks",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of lifecycle hooks available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Optional) Scaling group id the lifecycle hooks belong to.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting lifecycle hook by name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of lifecycle hook IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of lifecycle hook ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of lifecycle hook names.`,
				},
				resource.Attribute{
					Name:        "hooks",
					Description: `A list of lifecycle hooks. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lifecycle hook.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `ID of the scaling group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the lifecycle hook.`,
				},
				resource.Attribute{
					Name:        "default_result",
					Description: `Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses.`,
				},
				resource.Attribute{
					Name:        "heartbeat_timeout",
					Description: `Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the default_result parameter.`,
				},
				resource.Attribute{
					Name:        "lifecycle_transition",
					Description: `Type of Scaling activity attached to lifecycle hook.`,
				},
				resource.Attribute{
					Name:        "notification_arn",
					Description: `The Arn of notification target.`,
				},
				resource.Attribute{
					Name:        "notification_metadata",
					Description: `Additional information that you want to include when Auto Scaling sends a message to the notification target.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of lifecycle hook ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of lifecycle hook names.`,
				},
				resource.Attribute{
					Name:        "hooks",
					Description: `A list of lifecycle hooks. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the lifecycle hook.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `ID of the scaling group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the lifecycle hook.`,
				},
				resource.Attribute{
					Name:        "default_result",
					Description: `Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses.`,
				},
				resource.Attribute{
					Name:        "heartbeat_timeout",
					Description: `Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the default_result parameter.`,
				},
				resource.Attribute{
					Name:        "lifecycle_transition",
					Description: `Type of Scaling activity attached to lifecycle hook.`,
				},
				resource.Attribute{
					Name:        "notification_arn",
					Description: `The Arn of notification target.`,
				},
				resource.Attribute{
					Name:        "notification_metadata",
					Description: `Additional information that you want to include when Auto Scaling sends a message to the notification target.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ess_notifications",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of notifications available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Required) Scaling group id the notifications belong to.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional)A list of notification ids.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of notification ids.`,
				},
				resource.Attribute{
					Name:        "notifications",
					Description: `A list of notifications. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the notification.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `ID of the scaling group.`,
				},
				resource.Attribute{
					Name:        "notification_arn",
					Description: `The Alibaba Cloud Resource Name (ARN) for the notification object.`,
				},
				resource.Attribute{
					Name:        "notification_types",
					Description: `The notification types of Auto Scaling events and resource changes.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of notification ids.`,
				},
				resource.Attribute{
					Name:        "notifications",
					Description: `A list of notifications. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the notification.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `ID of the scaling group.`,
				},
				resource.Attribute{
					Name:        "notification_arn",
					Description: `The Alibaba Cloud Resource Name (ARN) for the notification object.`,
				},
				resource.Attribute{
					Name:        "notification_types",
					Description: `The notification types of Auto Scaling events and resource changes.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ess_scaling_configurations",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of scaling configurations available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Optional) Scaling group id the scaling configurations belong to.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting scaling configurations by name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of scaling configuration IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of scaling configuration ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of scaling configuration names.`,
				},
				resource.Attribute{
					Name:        "configurations",
					Description: `A list of scaling rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `ID of the scaling group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `Image ID of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Instance type of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Security group ID of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `Internet charge type of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_in",
					Description: `Internet max bandwidth in of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `Internet max bandwidth of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "system_disk_category",
					Description: `System disk category of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `System disk size of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "data_disks",
					Description: `Data disks of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Size of data disk.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `Category of data disk.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `Size of data disk.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Device attribute of data disk.`,
				},
				resource.Attribute{
					Name:        "delete_with_instance",
					Description: `Delete_with_instance attribute of data disk.`,
				},
				resource.Attribute{
					Name:        "lifecycle_state",
					Description: `Lifecycle state of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Creation time of the scaling configuration.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of scaling configuration ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of scaling configuration names.`,
				},
				resource.Attribute{
					Name:        "configurations",
					Description: `A list of scaling rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `ID of the scaling group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `Image ID of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Instance type of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Security group ID of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `Internet charge type of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_in",
					Description: `Internet max bandwidth in of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `Internet max bandwidth of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "system_disk_category",
					Description: `System disk category of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "system_disk_size",
					Description: `System disk size of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "data_disks",
					Description: `Data disks of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Size of data disk.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `Category of data disk.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `Size of data disk.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Device attribute of data disk.`,
				},
				resource.Attribute{
					Name:        "delete_with_instance",
					Description: `Delete_with_instance attribute of data disk.`,
				},
				resource.Attribute{
					Name:        "lifecycle_state",
					Description: `Lifecycle state of the scaling configuration.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Creation time of the scaling configuration.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ess_scaling_groups",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of scaling groups available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting scaling groups by name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of scaling group IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of scaling group ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of scaling group names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of scaling groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the scaling group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the scaling group.`,
				},
				resource.Attribute{
					Name:        "launch_template_id",
					Description: `Active launch template ID for scaling group.`,
				},
				resource.Attribute{
					Name:        "launch_template_version",
					Description: `Version of active launch template.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the scaling group belongs to.`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `The minimum number of ECS instances.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `The maximum number of ECS instances.`,
				},
				resource.Attribute{
					Name:        "cooldown_time",
					Description: `Default cooldown time of scaling group.`,
				},
				resource.Attribute{
					Name:        "removal_policies",
					Description: `Removal policy used to select the ECS instance to remove from the scaling group.`,
				},
				resource.Attribute{
					Name:        "load_balancer_ids",
					Description: `Slb instances id which the ECS instance attached to.`,
				},
				resource.Attribute{
					Name:        "db_instance_ids",
					Description: `Db instances id which the ECS instance attached to.`,
				},
				resource.Attribute{
					Name:        "vswitch_ids",
					Description: `Vswitches id in which the ECS instance launched.`,
				},
				resource.Attribute{
					Name:        "lifecycle_state",
					Description: `Lifecycle state of scaling group.`,
				},
				resource.Attribute{
					Name:        "total_capacity",
					Description: `Number of instances in scaling group.`,
				},
				resource.Attribute{
					Name:        "active_capacity",
					Description: `Number of active instances in scaling group.`,
				},
				resource.Attribute{
					Name:        "pending_capacity",
					Description: `Number of pending instances in scaling group.`,
				},
				resource.Attribute{
					Name:        "removing_capacity",
					Description: `Number of removing instances in scaling group.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Creation time of scaling group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of scaling group ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of scaling group names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of scaling groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the scaling group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the scaling group.`,
				},
				resource.Attribute{
					Name:        "launch_template_id",
					Description: `Active launch template ID for scaling group.`,
				},
				resource.Attribute{
					Name:        "launch_template_version",
					Description: `Version of active launch template.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the scaling group belongs to.`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `The minimum number of ECS instances.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `The maximum number of ECS instances.`,
				},
				resource.Attribute{
					Name:        "cooldown_time",
					Description: `Default cooldown time of scaling group.`,
				},
				resource.Attribute{
					Name:        "removal_policies",
					Description: `Removal policy used to select the ECS instance to remove from the scaling group.`,
				},
				resource.Attribute{
					Name:        "load_balancer_ids",
					Description: `Slb instances id which the ECS instance attached to.`,
				},
				resource.Attribute{
					Name:        "db_instance_ids",
					Description: `Db instances id which the ECS instance attached to.`,
				},
				resource.Attribute{
					Name:        "vswitch_ids",
					Description: `Vswitches id in which the ECS instance launched.`,
				},
				resource.Attribute{
					Name:        "lifecycle_state",
					Description: `Lifecycle state of scaling group.`,
				},
				resource.Attribute{
					Name:        "total_capacity",
					Description: `Number of instances in scaling group.`,
				},
				resource.Attribute{
					Name:        "active_capacity",
					Description: `Number of active instances in scaling group.`,
				},
				resource.Attribute{
					Name:        "pending_capacity",
					Description: `Number of pending instances in scaling group.`,
				},
				resource.Attribute{
					Name:        "removing_capacity",
					Description: `Number of removing instances in scaling group.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Creation time of scaling group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ess_scaling_rules",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of scaling rules available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `(Optional) Scaling group id the scaling rules belong to.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of scaling rule.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting scaling rules by name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of scaling rule IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of scaling rule ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of scaling rule names.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `A list of scaling rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `ID of the scaling group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "cooldown",
					Description: `Cooldown time of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "adjustment_type",
					Description: `Adjustment type of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "adjustment_value",
					Description: `Adjustment value of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "min_adjustment_magnitude",
					Description: `Min adjustment magnitude of scaling rule.`,
				},
				resource.Attribute{
					Name:        "scaling_rule_ari",
					Description: `Ari of scaling rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of scaling rule ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of scaling rule names.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `A list of scaling rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "scaling_group_id",
					Description: `ID of the scaling group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "cooldown",
					Description: `Cooldown time of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "adjustment_type",
					Description: `Adjustment type of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "adjustment_value",
					Description: `Adjustment value of the scaling rule.`,
				},
				resource.Attribute{
					Name:        "min_adjustment_magnitude",
					Description: `Min adjustment magnitude of scaling rule.`,
				},
				resource.Attribute{
					Name:        "scaling_rule_ari",
					Description: `Ari of scaling rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ess_scheduled_tasks",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of scheduled tasks available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "scheduled_task_id",
					Description: `(Optional) The id of the scheduled task.`,
				},
				resource.Attribute{
					Name:        "scheduled_action",
					Description: `(Optional) The operation to be performed when a scheduled task is triggered.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting scheduled tasks by name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of scheduled task IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of scheduled task ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of scheduled task names.`,
				},
				resource.Attribute{
					Name:        "tasks",
					Description: `A list of scheduled tasks. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the scheduled task id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the scheduled task name.`,
				},
				resource.Attribute{
					Name:        "scheduled_action",
					Description: `The operation to be performed when a scheduled task is triggered.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the scheduled task.`,
				},
				resource.Attribute{
					Name:        "launch_expiration_time",
					Description: `The time period during which a failed scheduled task is retried.`,
				},
				resource.Attribute{
					Name:        "launch_time",
					Description: `The time at which the scheduled task is triggered.`,
				},
				resource.Attribute{
					Name:        "recurrence_type",
					Description: `Specifies the recurrence type of the scheduled task.`,
				},
				resource.Attribute{
					Name:        "recurrence_value",
					Description: `Specifies how often a scheduled task recurs.`,
				},
				resource.Attribute{
					Name:        "recurrence_end_time",
					Description: `Specifies the end time after which the scheduled task is no longer repeated.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of scheduled task ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of scheduled task names.`,
				},
				resource.Attribute{
					Name:        "tasks",
					Description: `A list of scheduled tasks. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the scheduled task id.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the scheduled task name.`,
				},
				resource.Attribute{
					Name:        "scheduled_action",
					Description: `The operation to be performed when a scheduled task is triggered.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the scheduled task.`,
				},
				resource.Attribute{
					Name:        "launch_expiration_time",
					Description: `The time period during which a failed scheduled task is retried.`,
				},
				resource.Attribute{
					Name:        "launch_time",
					Description: `The time at which the scheduled task is triggered.`,
				},
				resource.Attribute{
					Name:        "recurrence_type",
					Description: `Specifies the recurrence type of the scheduled task.`,
				},
				resource.Attribute{
					Name:        "recurrence_value",
					Description: `Specifies how often a scheduled task recurs.`,
				},
				resource.Attribute{
					Name:        "recurrence_end_time",
					Description: `Specifies the end time after which the scheduled task is no longer repeated.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_fc_functions",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of FC functions to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `Name of the service that contains the functions to find.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by function name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of functions ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of functions names.`,
				},
				resource.Attribute{
					Name:        "functions",
					Description: `A list of functions. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Function ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Function name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Function description.`,
				},
				resource.Attribute{
					Name:        "runtime",
					Description: `Function runtime. The list of possible values is [available here](https://www.alibabacloud.com/help/doc-detail/52077.htm).`,
				},
				resource.Attribute{
					Name:        "handler",
					Description: `Function [entry point](https://www.alibabacloud.com/help/doc-detail/62213.htm) in the code.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Maximum amount of time the function can run in seconds.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `Amount of memory in MB the function can use at runtime.`,
				},
				resource.Attribute{
					Name:        "code_size",
					Description: `Function code size in bytes.`,
				},
				resource.Attribute{
					Name:        "code_checksum",
					Description: `Checksum (crc64) of the function code.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Function creation time.`,
				},
				resource.Attribute{
					Name:        "last_modification_time",
					Description: `Function last modification time.`,
				},
				resource.Attribute{
					Name:        "environment_variables",
					Description: `A map that defines environment variables for the function.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of functions ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of functions names.`,
				},
				resource.Attribute{
					Name:        "functions",
					Description: `A list of functions. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Function ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Function name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Function description.`,
				},
				resource.Attribute{
					Name:        "runtime",
					Description: `Function runtime. The list of possible values is [available here](https://www.alibabacloud.com/help/doc-detail/52077.htm).`,
				},
				resource.Attribute{
					Name:        "handler",
					Description: `Function [entry point](https://www.alibabacloud.com/help/doc-detail/62213.htm) in the code.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `Maximum amount of time the function can run in seconds.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `Amount of memory in MB the function can use at runtime.`,
				},
				resource.Attribute{
					Name:        "code_size",
					Description: `Function code size in bytes.`,
				},
				resource.Attribute{
					Name:        "code_checksum",
					Description: `Checksum (crc64) of the function code.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Function creation time.`,
				},
				resource.Attribute{
					Name:        "last_modification_time",
					Description: `Function last modification time.`,
				},
				resource.Attribute{
					Name:        "environment_variables",
					Description: `A map that defines environment variables for the function.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_fc_services",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of FC services to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by FC service name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of FC services ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of FC services names.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `A list of FC services. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `FC service ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `FC service name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `FC service description.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `FC service role ARN.`,
				},
				resource.Attribute{
					Name:        "internet_access",
					Description: `Indicate whether the service can access to internet or not.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `FC service creation time.`,
				},
				resource.Attribute{
					Name:        "last_modification_time",
					Description: `FC service last modification time.`,
				},
				resource.Attribute{
					Name:        "log_config",
					Description: `A list of one element containing information about the associated log store. It contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `Log Service project name.`,
				},
				resource.Attribute{
					Name:        "logstore",
					Description: `Log Service store name.`,
				},
				resource.Attribute{
					Name:        "vpc_config",
					Description: `A list of one element containing information about accessible VPC resources. It contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Associated VPC ID.`,
				},
				resource.Attribute{
					Name:        "vswitch_ids",
					Description: `Associated VSwitch IDs.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Associated security group ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of FC services ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of FC services names.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `A list of FC services. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `FC service ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `FC service name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `FC service description.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `FC service role ARN.`,
				},
				resource.Attribute{
					Name:        "internet_access",
					Description: `Indicate whether the service can access to internet or not.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `FC service creation time.`,
				},
				resource.Attribute{
					Name:        "last_modification_time",
					Description: `FC service last modification time.`,
				},
				resource.Attribute{
					Name:        "log_config",
					Description: `A list of one element containing information about the associated log store. It contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `Log Service project name.`,
				},
				resource.Attribute{
					Name:        "logstore",
					Description: `Log Service store name.`,
				},
				resource.Attribute{
					Name:        "vpc_config",
					Description: `A list of one element containing information about accessible VPC resources. It contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Associated VPC ID.`,
				},
				resource.Attribute{
					Name:        "vswitch_ids",
					Description: `Associated VSwitch IDs.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Associated security group ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_fc_triggers",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of FC triggers to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `FC service name.`,
				},
				resource.Attribute{
					Name:        "function_name",
					Description: `FC function name.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by FC trigger name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of FC triggers ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of FC triggers names.`,
				},
				resource.Attribute{
					Name:        "triggers",
					Description: `A list of FC triggers. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `FC trigger ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `FC trigger name.`,
				},
				resource.Attribute{
					Name:        "source_arn",
					Description: `Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the trigger. Valid values: ` + "`" + `oss` + "`" + `, ` + "`" + `log` + "`" + `, ` + "`" + `timer` + "`" + `, ` + "`" + `http` + "`" + ` and ` + "`" + `mns_topic` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "invocation_role",
					Description: `RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `JSON-encoded trigger configuration. See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `FC trigger creation time.`,
				},
				resource.Attribute{
					Name:        "last_modification_time",
					Description: `FC trigger last modification time.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of FC triggers ids.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of FC triggers names.`,
				},
				resource.Attribute{
					Name:        "triggers",
					Description: `A list of FC triggers. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `FC trigger ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `FC trigger name.`,
				},
				resource.Attribute{
					Name:        "source_arn",
					Description: `Event source resource address. See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the trigger. Valid values: ` + "`" + `oss` + "`" + `, ` + "`" + `log` + "`" + `, ` + "`" + `timer` + "`" + `, ` + "`" + `http` + "`" + ` and ` + "`" + `mns_topic` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "invocation_role",
					Description: `RAM role arn attached to the Function Compute trigger. Role used by the event source to call the function. The value format is "acs:ram::$account-id:role/$role-name". See [Create a trigger](https://www.alibabacloud.com/help/doc-detail/53102.htm) for more details.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `JSON-encoded trigger configuration. See [Configure triggers and events](https://www.alibabacloud.com/help/doc-detail/70140.htm) for more details.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `FC trigger creation time.`,
				},
				resource.Attribute{
					Name:        "last_modification_time",
					Description: `FC trigger last modification time.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_fc_zones",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of availability zones for FunctionCompute that can be used by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_file_crc64_checksum",
			Category:         "Data Sources",
			ShortDescription: `Provides compute file crc64 checksum.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filename",
					Description: `(Required) The name of the file to be computed crc64 checksum. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `file crc64 ID`,
				},
				resource.Attribute{
					Name:        "checksum",
					Description: `the file checksum of crc64.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `file crc64 ID`,
				},
				resource.Attribute{
					Name:        "checksum",
					Description: `the file checksum of crc64.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_forward_entries",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Forward Entries owned by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Forward Entries IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional, Available in 1.44.0+) A regex string to filter results by forward entry name.`,
				},
				resource.Attribute{
					Name:        "external_ip",
					Description: `(Optional) The public IP address.`,
				},
				resource.Attribute{
					Name:        "internal_ip",
					Description: `(Optional) The private IP address.`,
				},
				resource.Attribute{
					Name:        "forward_table_id",
					Description: `(Required) The ID of the Forward table.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of Forward Entries IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Forward Entries names.`,
				},
				resource.Attribute{
					Name:        "entries",
					Description: `A list of Forward Entries. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Forward Entry.`,
				},
				resource.Attribute{
					Name:        "external_ip",
					Description: `The public IP address.`,
				},
				resource.Attribute{
					Name:        "external_port",
					Description: `The public port.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The protocol type.`,
				},
				resource.Attribute{
					Name:        "internal_ip",
					Description: `The private IP address.`,
				},
				resource.Attribute{
					Name:        "internal_port",
					Description: `The private port.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The forward entry name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the Forward Entry.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of Forward Entries IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Forward Entries names.`,
				},
				resource.Attribute{
					Name:        "entries",
					Description: `A list of Forward Entries. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Forward Entry.`,
				},
				resource.Attribute{
					Name:        "external_ip",
					Description: `The public IP address.`,
				},
				resource.Attribute{
					Name:        "external_port",
					Description: `The public port.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The protocol type.`,
				},
				resource.Attribute{
					Name:        "internal_ip",
					Description: `The private IP address.`,
				},
				resource.Attribute{
					Name:        "internal_port",
					Description: `The private port.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The forward entry name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the Forward Entry.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_gpdb_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a collection of AnalyticDB for PostgreSQL instances according to the specified filters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of instance IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to apply to the instance name.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) Instance availability zone.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional) Used to retrieve instances belong to specified ` + "`" + `vswitch` + "`" + ` resources.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.55.3+) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) The name of file that can save the collection of instances after running ` + "`" + `terraform plan` + "`" + `. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `The ids list of AnalyticDB for PostgreSQL instances.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `The names list of AnalyticDB for PostgreSQL instance.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of AnalyticDB for PostgreSQL instances. Its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The instance id.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of an instance.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Billing method. Value options are ` + "`" + `PostPaid` + "`" + ` for Pay-As-You-Go and ` + "`" + `PrePaid` + "`" + ` for yearly or monthly subscription.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Instance availability zone.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `The time when you create an instance. The format is YYYY-MM-DDThh:mm:ssZ, such as 2011-05-30T12:11:4Z.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Database engine type. Supported option is ` + "`" + `gpdb` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Database engine version.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `Classic network or VPC.`,
				},
				resource.Attribute{
					Name:        "instance_class",
					Description: `The group type.`,
				},
				resource.Attribute{
					Name:        "instance_group_count",
					Description: `The number of groups.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `The ids list of AnalyticDB for PostgreSQL instances.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `The names list of AnalyticDB for PostgreSQL instance.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of AnalyticDB for PostgreSQL instances. Its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The instance id.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of an instance.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Billing method. Value options are ` + "`" + `PostPaid` + "`" + ` for Pay-As-You-Go and ` + "`" + `PrePaid` + "`" + ` for yearly or monthly subscription.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Instance availability zone.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `The time when you create an instance. The format is YYYY-MM-DDThh:mm:ssZ, such as 2011-05-30T12:11:4Z.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Database engine type. Supported option is ` + "`" + `gpdb` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Database engine version.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `Classic network or VPC.`,
				},
				resource.Attribute{
					Name:        "instance_class",
					Description: `The group type.`,
				},
				resource.Attribute{
					Name:        "instance_group_count",
					Description: `The number of groups.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_hbase_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a collection of HBase instances according to the specified filters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to apply to the instance name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) The ids list of HBase instances`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) Instance availability zone.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in 1.73.0) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) The name of file that can save the collection of instances after running ` + "`" + `terraform plan` + "`" + `. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `The ids list of HBase instances`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `The names list of HBase instances`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of HBase instances. Its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the HBase instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the HBase instance.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `Zone ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `the engine of the instance.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `the engine_version of the instance.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `Classic network or VPC.`,
				},
				resource.Attribute{
					Name:        "master_instance_type",
					Description: `hbase.sn1.large, hbase.sn1.large, hbase.n1.2xlarge and so on.`,
				},
				resource.Attribute{
					Name:        "master_node_count",
					Description: `the node count of master`,
				},
				resource.Attribute{
					Name:        "core_instance_type",
					Description: `hbase.sn1.large, hbase.sn1.large, hbase.n1.2xlarge and so on.`,
				},
				resource.Attribute{
					Name:        "core_node_count",
					Description: `same with "core_instance_quantity"`,
				},
				resource.Attribute{
					Name:        "core_disk_type",
					Description: `cloud_ssd or cloud_efficiency`,
				},
				resource.Attribute{
					Name:        "core_disk_size",
					Description: `core node disk size, unit:GB.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `VSwitch ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "pay_type",
					Description: `Billing method. Value options are ` + "`" + `PostPaid` + "`" + ` for Pay-As-You-Go and ` + "`" + `PrePaid` + "`" + ` for yearly or monthly subscription.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "backup_status",
					Description: `the Backup Status of the instance.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `the create time of the instance.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `the expire time of the instance.`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `the switch of delete protection.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `The ids list of HBase instances`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `The names list of HBase instances`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of HBase instances. Its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the HBase instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the HBase instance.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `Zone ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `the engine of the instance.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `the engine_version of the instance.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `Classic network or VPC.`,
				},
				resource.Attribute{
					Name:        "master_instance_type",
					Description: `hbase.sn1.large, hbase.sn1.large, hbase.n1.2xlarge and so on.`,
				},
				resource.Attribute{
					Name:        "master_node_count",
					Description: `the node count of master`,
				},
				resource.Attribute{
					Name:        "core_instance_type",
					Description: `hbase.sn1.large, hbase.sn1.large, hbase.n1.2xlarge and so on.`,
				},
				resource.Attribute{
					Name:        "core_node_count",
					Description: `same with "core_instance_quantity"`,
				},
				resource.Attribute{
					Name:        "core_disk_type",
					Description: `cloud_ssd or cloud_efficiency`,
				},
				resource.Attribute{
					Name:        "core_disk_size",
					Description: `core node disk size, unit:GB.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `VSwitch ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "pay_type",
					Description: `Billing method. Value options are ` + "`" + `PostPaid` + "`" + ` for Pay-As-You-Go and ` + "`" + `PrePaid` + "`" + ` for yearly or monthly subscription.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "backup_status",
					Description: `the Backup Status of the instance.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `the create time of the instance.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `the expire time of the instance.`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `the switch of delete protection.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_hbase_zones",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of availability zones for HBase that can be used by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "multi",
					Description: `(Optional) Indicate whether the zones can be used in a multi AZ configuration. Default to ` + "`" + `false` + "`" + `. Multi AZ is usually used to launch HBase instances.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "multi_zone_ids",
					Description: `A list of zone ids in which the multi zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "multi_zone_ids",
					Description: `A list of zone ids in which the multi zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_images",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of images available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting images by name.`,
				},
				resource.Attribute{
					Name:        "most_recent",
					Description: `(Optional, type: bool) If more than one result are returned, select the most recent one.`,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `(Optional) Filter results by a specific image owner. Valid items are ` + "`" + `system` + "`" + `, ` + "`" + `self` + "`" + `, ` + "`" + `others` + "`" + `, ` + "`" + `marketplace` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ->`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of image IDs.`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `A list of images. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the image.`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `Platform type of the image system: i386 or x86_64.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Time of creation.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the image.`,
				},
				resource.Attribute{
					Name:        "image_owner_alias",
					Description: `Alias of the image owner.`,
				},
				resource.Attribute{
					Name:        "os_name",
					Description: `Display Chinese name of the OS.`,
				},
				resource.Attribute{
					Name:        "os_name_en",
					Description: `Display English name of the OS.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the image. Possible values: ` + "`" + `UnAvailable` + "`" + `, ` + "`" + `Available` + "`" + `, ` + "`" + `Creating` + "`" + ` and ` + "`" + `CreateFailed` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Size of the image.`,
				},
				resource.Attribute{
					Name:        "disk_device_mappings",
					Description: `Description of the system with disks and snapshots under the image.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Device information of the created disk: such as /dev/xvdb.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Size of the created disk.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `Snapshot ID.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `Product code of the image on the image market.`,
				},
				resource.Attribute{
					Name:        "is_subscribed",
					Description: `Whether the user has subscribed to the terms of service for the image product corresponding to the ProductCode.`,
				},
				resource.Attribute{
					Name:        "image_version",
					Description: `Version of the image.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `Progress of image creation, presented in percentages.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of image IDs.`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `A list of images. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the image.`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `Platform type of the image system: i386 or x86_64.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Time of creation.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the image.`,
				},
				resource.Attribute{
					Name:        "image_owner_alias",
					Description: `Alias of the image owner.`,
				},
				resource.Attribute{
					Name:        "os_name",
					Description: `Display Chinese name of the OS.`,
				},
				resource.Attribute{
					Name:        "os_name_en",
					Description: `Display English name of the OS.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the image. Possible values: ` + "`" + `UnAvailable` + "`" + `, ` + "`" + `Available` + "`" + `, ` + "`" + `Creating` + "`" + ` and ` + "`" + `CreateFailed` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Size of the image.`,
				},
				resource.Attribute{
					Name:        "disk_device_mappings",
					Description: `Description of the system with disks and snapshots under the image.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Device information of the created disk: such as /dev/xvdb.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Size of the created disk.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `Snapshot ID.`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `Product code of the image on the image market.`,
				},
				resource.Attribute{
					Name:        "is_subscribed",
					Description: `Whether the user has subscribed to the terms of service for the image product corresponding to the ProductCode.`,
				},
				resource.Attribute{
					Name:        "image_version",
					Description: `Version of the image.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `Progress of image creation, presented in percentages.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_instance_type_families",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of ECS Instance Type Families to be used by the alicloud_instance resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional, ForceNew) The Zone to launch the instance.`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `(Optional) The generation of the instance type family, Valid values: ` + "`" + `ecs-1` + "`" + `, ` + "`" + `ecs-2` + "`" + `, ` + "`" + `ecs-3` + "`" + ` and ` + "`" + `ecs-4` + "`" + `. For more information, see [Instance type families](https://www.alibabacloud.com/help/doc-detail/25378.htm).`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional, ForceNew) Valid values are ` + "`" + `PrePaid` + "`" + `, ` + "`" + `PostPaid` + "`" + `, Default to ` + "`" + `PostPaid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "spot_strategy",
					Description: `(Optional, ForceNew) Filter the results by ECS spot type. Valid values: ` + "`" + `NoSpot` + "`" + `, ` + "`" + `SpotWithPriceLimit` + "`" + ` and ` + "`" + `SpotAsPriceGo` + "`" + `. Default to ` + "`" + `NoSpot` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of instance type family IDs.`,
				},
				resource.Attribute{
					Name:        "instance_types",
					Description: `A list of image type families. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the instance type family.`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `The generation of the instance type family.`,
				},
				resource.Attribute{
					Name:        "zone_ids",
					Description: `A list of Zone to launch the instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of instance type family IDs.`,
				},
				resource.Attribute{
					Name:        "instance_types",
					Description: `A list of image type families. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the instance type family.`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `The generation of the instance type family.`,
				},
				resource.Attribute{
					Name:        "zone_ids",
					Description: `A list of Zone to launch the instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_instance_types",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of ECS Instance Types to be used by the alicloud_instance resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) The zone where instance types are supported.`,
				},
				resource.Attribute{
					Name:        "cpu_core_count",
					Description: `(Optional) Filter the results to a specific number of cpu cores.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `(Optional) Filter the results to a specific memory size in GB.`,
				},
				resource.Attribute{
					Name:        "gpu_amount",
					Description: `(Optional, Available in 1.69.0+) The GPU amount of an instance type.`,
				},
				resource.Attribute{
					Name:        "gpu_spec",
					Description: `(Optional, Available in 1.69.0+) The GPU spec of an instance type.`,
				},
				resource.Attribute{
					Name:        "instance_type_family",
					Description: `(Optional) Filter the results based on their family name. For example: 'ecs.n4'.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional) Filter the results by charge type. Valid values: ` + "`" + `PrePaid` + "`" + ` and ` + "`" + `PostPaid` + "`" + `. Default to ` + "`" + `PostPaid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `(Optional) Filter the results by network type. Valid values: ` + "`" + `Classic` + "`" + ` and ` + "`" + `Vpc` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "spot_strategy",
					Description: `(Optional) Filter the results by ECS spot type. Valid values: ` + "`" + `NoSpot` + "`" + `, ` + "`" + `SpotWithPriceLimit` + "`" + ` and ` + "`" + `SpotAsPriceGo` + "`" + `. Default to ` + "`" + `NoSpot` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "eni_amount",
					Description: `(Optional) Filter the result whose network interface number is no more than ` + "`" + `eni_amount` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "kubernetes_node_role",
					Description: `(Optional) Filter the result which is used to create a [kubernetes cluster](https://www.terraform.io/docs/providers/alicloud/r/cs_kubernetes.html) and [managed kubernetes cluster](https://www.terraform.io/docs/providers/alicloud/r/cs_managed_kubernetes.html). Optional Values: ` + "`" + `Master` + "`" + ` and ` + "`" + `Worker` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "is_outdated",
					Description: `(Optional, type: bool) If true, outdated instance types are included in the results. Default to false.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of instance type IDs.`,
				},
				resource.Attribute{
					Name:        "instance_types",
					Description: `A list of image types. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the instance type.`,
				},
				resource.Attribute{
					Name:        "cpu_core_count",
					Description: `Number of CPU cores.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `Size of memory, measured in GB.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `The instance type family.`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `List of availability zones that support the instance type.`,
				},
				resource.Attribute{
					Name:        "gpu",
					Description: `The GPU attribution of an instance type:`,
				},
				resource.Attribute{
					Name:        "amount",
					Description: `The amount of GPU of an instance type.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `The category of GPU of an instance type.`,
				},
				resource.Attribute{
					Name:        "burstable_instance",
					Description: `The burstable instance attribution:`,
				},
				resource.Attribute{
					Name:        "initial_credit",
					Description: `The initial CPU credit of a burstable instance.`,
				},
				resource.Attribute{
					Name:        "baseline_credit",
					Description: `The compute performance benchmark CPU credit of a burstable instance.`,
				},
				resource.Attribute{
					Name:        "eni_amount",
					Description: `The maximum number of network interfaces that an instance type can be attached to.`,
				},
				resource.Attribute{
					Name:        "local_storage",
					Description: `Local storage of an instance type:`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `The capacity of a local storage in GB.`,
				},
				resource.Attribute{
					Name:        "amount",
					Description: `The number of local storage devices that an instance has been attached to.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `The category of local storage that an instance has been attached to.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of instance type IDs.`,
				},
				resource.Attribute{
					Name:        "instance_types",
					Description: `A list of image types. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the instance type.`,
				},
				resource.Attribute{
					Name:        "cpu_core_count",
					Description: `Number of CPU cores.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `Size of memory, measured in GB.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `The instance type family.`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `List of availability zones that support the instance type.`,
				},
				resource.Attribute{
					Name:        "gpu",
					Description: `The GPU attribution of an instance type:`,
				},
				resource.Attribute{
					Name:        "amount",
					Description: `The amount of GPU of an instance type.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `The category of GPU of an instance type.`,
				},
				resource.Attribute{
					Name:        "burstable_instance",
					Description: `The burstable instance attribution:`,
				},
				resource.Attribute{
					Name:        "initial_credit",
					Description: `The initial CPU credit of a burstable instance.`,
				},
				resource.Attribute{
					Name:        "baseline_credit",
					Description: `The compute performance benchmark CPU credit of a burstable instance.`,
				},
				resource.Attribute{
					Name:        "eni_amount",
					Description: `The maximum number of network interfaces that an instance type can be attached to.`,
				},
				resource.Attribute{
					Name:        "local_storage",
					Description: `Local storage of an instance type:`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `The capacity of a local storage in GB.`,
				},
				resource.Attribute{
					Name:        "amount",
					Description: `The number of local storage devices that an instance has been attached to.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `The category of local storage that an instance has been attached to.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of ECS instances to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of ECS instance IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by instance name.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) The image ID of some ECS instance used.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Instance status. Valid values: "Creating", "Starting", "Running", "Stopping" and "Stopped". If undefined, all statuses are considered.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) ID of the VPC linked to the instances.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional) ID of the VSwitch linked to the instances.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) Availability zone where instances are located.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew, Available in 1.57.0+) The Id of resource group which the instance belongs.`,
				},
				resource.Attribute{
					Name:        "ram_role_name",
					Description: `(Optional, ForceNew, Available in 1.69.0+) The RAM role name which the instance attaches.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A map of tags assigned to the ECS instances. It must be in the format: ` + "`" + `` + "`" + `` + "`" + ` data "alicloud_instances" "taggedInstances" { tags = { tagKey1 = "tagValue1", tagKey2 = "tagValue2" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of ECS instance IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of instances names.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the instance.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Instance current status.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Instance name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Instance description.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Instance type.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `ID of the VSwitch the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `Image ID the instance is using.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Instance private IP address.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Instance public IP address.`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `EIP address the VPC instance is using.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `List of security group IDs the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `Key pair the instance is using.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Instance creation time.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `Instance charge type.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `Instance network charge type.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `Max output bandwidth for internet.`,
				},
				resource.Attribute{
					Name:        "spot_strategy",
					Description: `Spot strategy the instance is using.`,
				},
				resource.Attribute{
					Name:        "disk_device_mappings",
					Description: `Description of the attached disks.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Device information of the created disk: such as /dev/xvdb.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Size of the created disk.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `Cloud disk category.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Cloud disk type: system disk or data disk.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the ECS instance.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The Id of resource group.`,
				},
				resource.Attribute{
					Name:        "ram_role_name",
					Description: `The Ram role name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of ECS instance IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of instances names.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the instance.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Instance current status.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Instance name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Instance description.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Instance type.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `ID of the VSwitch the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `Image ID the instance is using.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Instance private IP address.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Instance public IP address.`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `EIP address the VPC instance is using.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `List of security group IDs the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `Key pair the instance is using.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Instance creation time.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `Instance charge type.`,
				},
				resource.Attribute{
					Name:        "internet_charge_type",
					Description: `Instance network charge type.`,
				},
				resource.Attribute{
					Name:        "internet_max_bandwidth_out",
					Description: `Max output bandwidth for internet.`,
				},
				resource.Attribute{
					Name:        "spot_strategy",
					Description: `Spot strategy the instance is using.`,
				},
				resource.Attribute{
					Name:        "disk_device_mappings",
					Description: `Description of the attached disks.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Device information of the created disk: such as /dev/xvdb.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Size of the created disk.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `Cloud disk category.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Cloud disk type: system disk or data disk.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the ECS instance.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The Id of resource group.`,
				},
				resource.Attribute{
					Name:        "ram_role_name",
					Description: `The Ram role name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_key_pairs",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of available key pairs that can be used by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to apply to the resulting key pairs.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, Available 1.52.1+) A list of key pair IDs.`,
				},
				resource.Attribute{
					Name:        "finger_print",
					Description: `(Optional) A finger print used to retrieve specified key pair.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew, Available in 1.57.0+) The Id of resource group which the key pair belongs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.66.0+) A mapping of tags to assign to the resource. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of key pair names.`,
				},
				resource.Attribute{
					Name:        "key_pairs",
					Description: `A list of key pairs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the key pair.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `Name of the key pair.`,
				},
				resource.Attribute{
					Name:        "finger_print",
					Description: `Finger print of the key pair.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of ECS instances that has been bound this key pair.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The ID of the availability zone where the ECS instance is located.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the ECS instance.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `The name of the ECS instance.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `The ID of the VSwitch attached to the ECS instance.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP address or EIP of the ECS instance.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address of the ECS instance.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The Id of resource group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.66.0+) A mapping of tags to assign to the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `A list of key pair names.`,
				},
				resource.Attribute{
					Name:        "key_pairs",
					Description: `A list of key pairs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the key pair.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `Name of the key pair.`,
				},
				resource.Attribute{
					Name:        "finger_print",
					Description: `Finger print of the key pair.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of ECS instances that has been bound this key pair.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The ID of the availability zone where the ECS instance is located.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the ECS instance.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `The name of the ECS instance.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `The ID of the VSwitch attached to the ECS instance.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `The public IP address or EIP of the ECS instance.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IP address of the ECS instance.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The Id of resource group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.66.0+) A mapping of tags to assign to the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_kms_aliases",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of available KMS Aliases.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of KMS aliases IDs. The value is same as KMS alias_name.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter the results by the KMS alias name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of kms aliases IDs. The value is same as KMS alias_name.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of KMS alias name.`,
				},
				resource.Attribute{
					Name:        "aliases",
					Description: `A list of KMS User alias. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the alias. The value is same as KMS alias_name.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `ID of the key.`,
				},
				resource.Attribute{
					Name:        "alias_name",
					Description: `The unique identifier of the alias.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of kms aliases IDs. The value is same as KMS alias_name.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of KMS alias name.`,
				},
				resource.Attribute{
					Name:        "aliases",
					Description: `A list of KMS User alias. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the alias. The value is same as KMS alias_name.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `ID of the key.`,
				},
				resource.Attribute{
					Name:        "alias_name",
					Description: `The unique identifier of the alias.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_kms_ciphertext",
			Category:         "Data Sources",
			ShortDescription: `Encrypt data with KMS.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "plaintext",
					Description: `The plaintext to be encrypted which must be encoded in Base64.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `The globally unique ID of the CMK.`,
				},
				resource.Attribute{
					Name:        "encryption_context",
					Description: `(Optional) The Encryption context. If you specify this parameter here, it is also required when you call the Decrypt API operation. For more information, see [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ciphertext_blob",
					Description: `The ciphertext of the data key encrypted with the primary CMK version.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ciphertext_blob",
					Description: `The ciphertext of the data key encrypted with the primary CMK version.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_kms_key_versions",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of available KMS KeyVersions.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key_id",
					Description: `(Required) The id of kms key.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of KMS KeyVersion IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of KMS KeyVersion IDs.`,
				},
				resource.Attribute{
					Name:        "versions",
					Description: `A list of KMS KeyVersions. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Date and time when the key version was created (UTC time).`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `ID of the key.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the KMS KeyVersion resource.`,
				},
				resource.Attribute{
					Name:        "key_version_id",
					Description: `ID of the key version.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of KMS KeyVersion IDs.`,
				},
				resource.Attribute{
					Name:        "versions",
					Description: `A list of KMS KeyVersions. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Date and time when the key version was created (UTC time).`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `ID of the key.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the KMS KeyVersion resource.`,
				},
				resource.Attribute{
					Name:        "key_version_id",
					Description: `ID of the key version.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_kms_keys",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of available KMS Keys.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of KMS key IDs.`,
				},
				resource.Attribute{
					Name:        "description_regex",
					Description: `(Optional) A regex string to filter the results by the KMS key description.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Filter the results by status of the KMS keys. Valid values: ` + "`" + `Enabled` + "`" + `, ` + "`" + `Disabled` + "`" + `, ` + "`" + `PendingDeletion` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of KMS key IDs.`,
				},
				resource.Attribute{
					Name:        "keys",
					Description: `A list of KMS keys. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the key.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The Alibaba Cloud Resource Name (ARN) of the key.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the key.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the key. Possible values: ` + "`" + `Enabled` + "`" + `, ` + "`" + `Disabled` + "`" + ` and ` + "`" + `PendingDeletion` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Creation date of key.`,
				},
				resource.Attribute{
					Name:        "delete_date",
					Description: `Deletion date of key.`,
				},
				resource.Attribute{
					Name:        "creator",
					Description: `The owner of the key.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of KMS key IDs.`,
				},
				resource.Attribute{
					Name:        "keys",
					Description: `A list of KMS keys. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the key.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `The Alibaba Cloud Resource Name (ARN) of the key.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the key.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the key. Possible values: ` + "`" + `Enabled` + "`" + `, ` + "`" + `Disabled` + "`" + ` and ` + "`" + `PendingDeletion` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Creation date of key.`,
				},
				resource.Attribute{
					Name:        "delete_date",
					Description: `Deletion date of key.`,
				},
				resource.Attribute{
					Name:        "creator",
					Description: `The owner of the key.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_kms_plaintext",
			Category:         "Data Sources",
			ShortDescription: `Decrypt a ciphertext encrypted with KMS.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "encryption_context",
					Description: `(Optional) The Encryption context. If you specify this parameter in the Encrypt or GenerateDataKey API operation, it is also required when you call the Decrypt API operation. For more information, see [Encryption Context](https://www.alibabacloud.com/help/doc-detail/42975.htm).`,
				},
				resource.Attribute{
					Name:        "ciphertext_blob",
					Description: `The ciphertext to be decrypted. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "plaintext",
					Description: `The decrypted plaintext.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `The globally unique ID of the CMK. It is the ID of the CMK used to decrypt ciphertext.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "plaintext",
					Description: `The decrypted plaintext.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `The globally unique ID of the CMK. It is the ID of the CMK used to decrypt ciphertext.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_kms_secret_versions",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of available KMS Secret Versions.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "include_deprecated",
					Description: `(Optional, ForceNew) Specifies whether to return deprecated secret versions. Default to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, ForceNew) A list of KMS Secret Version ids.`,
				},
				resource.Attribute{
					Name:        "secret_name",
					Description: `(Required, ForceNew) The name of the secret.`,
				},
				resource.Attribute{
					Name:        "version_stage",
					Description: `(Optional, ForceNew, Available in 1.89.0+) The stage of the secret version.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "enable_details",
					Description: `(Optional) Default to false and only output ` + "`" + `secret_name` + "`" + `, ` + "`" + `version_id` + "`" + `, ` + "`" + `version_stages` + "`" + `. Set it to true can output more details. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of Kms Secret Version ids.`,
				},
				resource.Attribute{
					Name:        "versions",
					Description: `A list of KMS Secret Versions. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "secret_data",
					Description: `The secret value. Secrets Manager decrypts the stored secret value in ciphertext and returns it. (Returned when ` + "`" + `enable_details` + "`" + ` is true).`,
				},
				resource.Attribute{
					Name:        "secret_data_type",
					Description: `The type of the secret value. (Returned when ` + "`" + `enable_details` + "`" + ` is true).`,
				},
				resource.Attribute{
					Name:        "secret_name",
					Description: `The name of the secret.`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `The version number of the secret value.`,
				},
				resource.Attribute{
					Name:        "version_stages",
					Description: `Stage labels that mark the secret version.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of Kms Secret Version ids.`,
				},
				resource.Attribute{
					Name:        "versions",
					Description: `A list of KMS Secret Versions. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "secret_data",
					Description: `The secret value. Secrets Manager decrypts the stored secret value in ciphertext and returns it. (Returned when ` + "`" + `enable_details` + "`" + ` is true).`,
				},
				resource.Attribute{
					Name:        "secret_data_type",
					Description: `The type of the secret value. (Returned when ` + "`" + `enable_details` + "`" + ` is true).`,
				},
				resource.Attribute{
					Name:        "secret_name",
					Description: `The name of the secret.`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `The version number of the secret value.`,
				},
				resource.Attribute{
					Name:        "version_stages",
					Description: `Stage labels that mark the secret version.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_kms_secrets",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of available KMS Secrets.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "fetch_tags",
					Description: `(Optional) Whether to include the predetermined resource tag in the return value. Default to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of KMS Secret ids. The value is same as KMS secret_name.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter the results by the KMS secret_name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of Kms Secret ids. The value is same as KMS secret_name.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of KMS Secret names.`,
				},
				resource.Attribute{
					Name:        "secrets",
					Description: `A list of KMS Secrets. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Kms Secret. The value is same as KMS secret_name.`,
				},
				resource.Attribute{
					Name:        "secret_name",
					Description: `Name of the KMS Secret.`,
				},
				resource.Attribute{
					Name:        "planned_delete_time",
					Description: `Schedule deletion time.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of Kms Secret ids. The value is same as KMS secret_name.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of KMS Secret names.`,
				},
				resource.Attribute{
					Name:        "secrets",
					Description: `A list of KMS Secrets. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Kms Secret. The value is same as KMS secret_name.`,
				},
				resource.Attribute{
					Name:        "secret_name",
					Description: `Name of the KMS Secret.`,
				},
				resource.Attribute{
					Name:        "planned_delete_time",
					Description: `Schedule deletion time.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_kvstore_instance_classes",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of KVStore instacne classes info.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The Zone to launch the KVStore instance.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional) Filter the results by charge type. Valid values: ` + "`" + `PrePaid` + "`" + ` and ` + "`" + `PostPaid` + "`" + `. Default to ` + "`" + `PrePaid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Optional) Database type. Options are ` + "`" + `Redis` + "`" + `, ` + "`" + `Memcache` + "`" + `. Default to ` + "`" + `Redis` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Optional) Database version required by the user. Value options of Redis can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/60873.htm) ` + "`" + `EngineVersion` + "`" + `. Value of Memcache should be empty.`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `(Optional) The KVStore instance system architecture required by the user. Valid values: ` + "`" + `standard` + "`" + `, ` + "`" + `cluster` + "`" + ` and ` + "`" + `rwsplit` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "performance_type",
					Description: `(Optional, Deprecated) It has been deprecated from 1.68.0.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Optional, Deprecated) It has been deprecated from 1.68.0.`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `(Optional) The KVStore instance node type required by the user. Valid values: ` + "`" + `double` + "`" + `, ` + "`" + `single` + "`" + `, ` + "`" + `readone` + "`" + `, ` + "`" + `readthree` + "`" + ` and ` + "`" + `readfive` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "package_type",
					Description: `(Optional, Deprecated) It has been deprecated from 1.68.0.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform apply` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "edition_type",
					Description: `(Optional, Available in 1.68.0+) The KVStore instance edition type required by the user. Valid values: ` + "`" + `Community` + "`" + ` and ` + "`" + `Enterprise` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "series_type",
					Description: `(Optional, Available in 1.68.0+) The KVStore instance series type required by the user. Valid values: ` + "`" + `enhanced_performance_type` + "`" + ` and ` + "`" + `hybrid_storage` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "shard_number",
					Description: `(Optional, Available in 1.68.0+) The number of shard.Valid values: ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `8` + "`" + `, ` + "`" + `16` + "`" + `, ` + "`" + `32` + "`" + `, ` + "`" + `64` + "`" + `, ` + "`" + `128` + "`" + `, ` + "`" + `256` + "`" + `. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "instance_classes",
					Description: `A list of KVStore available instance classes.`,
				},
				resource.Attribute{
					Name:        "classes",
					Description: `A list of KVStore available instance classes when the ` + "`" + `sorted_by` + "`" + ` is "Price". include:`,
				},
				resource.Attribute{
					Name:        "instance_class",
					Description: `KVStore available instance class.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_classes",
					Description: `A list of KVStore available instance classes.`,
				},
				resource.Attribute{
					Name:        "classes",
					Description: `A list of KVStore available instance classes when the ` + "`" + `sorted_by` + "`" + ` is "Price". include:`,
				},
				resource.Attribute{
					Name:        "instance_class",
					Description: `KVStore available instance class.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_kvstore_instance_engines",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of KVStore instacne engines info.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The Zone to launch the KVStore instance.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional) Filter the results by charge type. Valid values: ` + "`" + `PrePaid` + "`" + ` and ` + "`" + `PostPaid` + "`" + `. Default to ` + "`" + `PrePaid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Optional) Database type. Options are ` + "`" + `Redis` + "`" + `, ` + "`" + `Memcache` + "`" + `. Default to ` + "`" + `Redis` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Optional) Database version required by the user. Value options of Redis can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/60873.htm) ` + "`" + `EngineVersion` + "`" + `. Value of Memcache should be empty.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform apply` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "instance_engines",
					Description: `A list of KVStore available instance engines. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `The Zone to launch the KVStore instance.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Database type.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `KVStore Instance version.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_engines",
					Description: `A list of KVStore available instance engines. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `The Zone to launch the KVStore instance.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Database type.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `KVStore Instance version.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_kvstore_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a collection of kvstore instances according to the specified filters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to apply to the instance name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, Available 1.52.2+) A list of RKV instance IDs.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional) Database type. Options are ` + "`" + `Memcache` + "`" + `, and ` + "`" + `Redis` + "`" + `. If no value is specified, all types are returned.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Status of the instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) Used to retrieve instances belong to specified VPC.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional) Used to retrieve instances belong to specified ` + "`" + `vswitch` + "`" + ` resources.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Query the instance bound to the tag. The format of the incoming value is ` + "`" + `json` + "`" + ` string, including ` + "`" + `TagKey` + "`" + ` and ` + "`" + `TagValue` + "`" + `. ` + "`" + `TagKey` + "`" + ` cannot be null, and ` + "`" + `TagValue` + "`" + ` can be empty. Format example ` + "`" + `{"key1":"value1"}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) The name of file that can save the collection of instances after running ` + "`" + `terraform plan` + "`" + `. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of RKV instance IDs.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of RKV instances. Its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the RKV instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the RKV instance.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Billing method. Value options: ` + "`" + `PostPaid` + "`" + ` for Pay-As-You-Go and ` + "`" + `PrePaid` + "`" + ` for subscription.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the instance.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `Expiration time. Pay-As-You-Go instances are never expire.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional) Database type. Options are ` + "`" + `Memcache` + "`" + `, and ` + "`" + `Redis` + "`" + `. If no value is specified, all types are returned.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `VSwitch ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP address of the instance.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username of the instance.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `Capacity of the applied ApsaraDB for Redis instance. Unit: MB.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `Instance bandwidth limit. Unit: Mbit/s.`,
				},
				resource.Attribute{
					Name:        "connections",
					Description: `Instance connection quantity limit. Unit: count.`,
				},
				resource.Attribute{
					Name:        "connections_domain",
					Description: `Instance connection domain (only Intranet access supported).`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Connection port of the instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of RKV instance IDs.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of RKV instances. Its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the RKV instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the RKV instance.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Billing method. Value options: ` + "`" + `PostPaid` + "`" + ` for Pay-As-You-Go and ` + "`" + `PrePaid` + "`" + ` for subscription.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation time of the instance.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `Expiration time. Pay-As-You-Go instances are never expire.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional) Database type. Options are ` + "`" + `Memcache` + "`" + `, and ` + "`" + `Redis` + "`" + `. If no value is specified, all types are returned.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Availability zone.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `VSwitch ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP address of the instance.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username of the instance.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `Capacity of the applied ApsaraDB for Redis instance. Unit: MB.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `Instance bandwidth limit. Unit: Mbit/s.`,
				},
				resource.Attribute{
					Name:        "connections",
					Description: `Instance connection quantity limit. Unit: count.`,
				},
				resource.Attribute{
					Name:        "connections_domain",
					Description: `Instance connection domain (only Intranet access supported).`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Connection port of the instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_kvstore_zones",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of availability zones for KVStore that can be used by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "multi",
					Description: `(Optional) Indicate whether the zones can be used in a multi AZ configuration. Default to ` + "`" + `false` + "`" + `. Multi AZ is usually used to launch KVStore instances.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional) Filter the results by a specific instance charge type. Valid values: ` + "`" + `PrePaid` + "`" + ` and ` + "`" + `PostPaid` + "`" + `. Default to ` + "`" + `PostPaid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "multi_zone_ids",
					Description: `A list of zone ids in which the multi zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "multi_zone_ids",
					Description: `A list of zone ids in which the multi zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_market_product",
			Category:         "Data Sources",
			ShortDescription: `Provides details of a Market product item.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "product_code",
					Description: `(Required) The product code of the market product.`,
				},
				resource.Attribute{
					Name:        "available_region",
					Description: `(Available in 1.71.1+) A available region id used to filter market place Ecs images. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "product",
					Description: `A product. It contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `The code of the product.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the product.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the product.`,
				},
				resource.Attribute{
					Name:        "skus",
					Description: `A list of one element containing sku attributes of an object. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "sku_code",
					Description: `The sku code of this product sku.`,
				},
				resource.Attribute{
					Name:        "sku_name",
					Description: `The sku name of this product sku.`,
				},
				resource.Attribute{
					Name:        "package_versions",
					Description: `The list of package version details of this product sku, Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "package_name",
					Description: `The package name of this product sku package.`,
				},
				resource.Attribute{
					Name:        "package_version",
					Description: `The package version of this product sku package. Currently, the API products can return package_version, but others can not for ensure.`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `The list of custom ECS images, Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The Ecs image id.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `The Ecs image display name.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The Ecs image region.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "product",
					Description: `A product. It contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `The code of the product.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the product.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the product.`,
				},
				resource.Attribute{
					Name:        "skus",
					Description: `A list of one element containing sku attributes of an object. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "sku_code",
					Description: `The sku code of this product sku.`,
				},
				resource.Attribute{
					Name:        "sku_name",
					Description: `The sku name of this product sku.`,
				},
				resource.Attribute{
					Name:        "package_versions",
					Description: `The list of package version details of this product sku, Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "package_name",
					Description: `The package name of this product sku package.`,
				},
				resource.Attribute{
					Name:        "package_version",
					Description: `The package version of this product sku package. Currently, the API products can return package_version, but others can not for ensure.`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `The list of custom ECS images, Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The Ecs image id.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `The Ecs image display name.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The Ecs image region.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_market_products",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Market product items.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional, Available 1.66.0+) A regex string to apply to the product name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, Available 1.66.0+) A list of product code.`,
				},
				resource.Attribute{
					Name:        "sort",
					Description: `(Optional, ForceNew) This field determines how to sort the filtered results, Valid values: ` + "`" + `user_count-desc` + "`" + `, ` + "`" + `created_on-desc` + "`" + `, ` + "`" + `price-desc` + "`" + ` and ` + "`" + `score-desc` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "category_id",
					Description: `(Optional, ForceNew) The Category ID of products. For more information, see [DescribeProducts](https://help.aliyun.com/document_detail/89834.htm).`,
				},
				resource.Attribute{
					Name:        "product_type",
					Description: `(Optional, ForceNew) The type of products, Valid values: ` + "`" + `APP` + "`" + `, ` + "`" + `SERVICE` + "`" + `, ` + "`" + `MIRROR` + "`" + `, ` + "`" + `DOWNLOAD` + "`" + ` and ` + "`" + `API_SERVICE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "search_term",
					Description: `(Optional, ForceNew, Available 1.69.0+) Search term in this query.`,
				},
				resource.Attribute{
					Name:        "supplier_id",
					Description: `(Optional, ForceNew, Available 1.71.1+) The supplier id of the product.`,
				},
				resource.Attribute{
					Name:        "supplier_name_keyword",
					Description: `(Optional, ForceNew, Available 1.71.1+) The supplier name keyword of the product.`,
				},
				resource.Attribute{
					Name:        "suggested_price",
					Description: `(Optional, ForceNew, Available 1.71.1+) The suggested price of the product.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of product codes.`,
				},
				resource.Attribute{
					Name:        "products",
					Description: `A list of products. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `The code of the product.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the product.`,
				},
				resource.Attribute{
					Name:        "category_id",
					Description: `The category id of the product.`,
				},
				resource.Attribute{
					Name:        "supplier_id",
					Description: `The supplier id of the product.`,
				},
				resource.Attribute{
					Name:        "supplier_name",
					Description: `The supplier name of the product.`,
				},
				resource.Attribute{
					Name:        "short_description",
					Description: `The short description of the product.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags of the product.`,
				},
				resource.Attribute{
					Name:        "suggested_price",
					Description: `The suggested price of the product.`,
				},
				resource.Attribute{
					Name:        "target_url",
					Description: `The detail page URL of the product.`,
				},
				resource.Attribute{
					Name:        "image_url",
					Description: `The image URL of the product.`,
				},
				resource.Attribute{
					Name:        "score",
					Description: `The rating information of the product.`,
				},
				resource.Attribute{
					Name:        "operation_system",
					Description: `The operation system of the product.`,
				},
				resource.Attribute{
					Name:        "warranty_date",
					Description: `The warranty date of the product.`,
				},
				resource.Attribute{
					Name:        "delivery_date",
					Description: `The delivery date of the product.`,
				},
				resource.Attribute{
					Name:        "delivery_way",
					Description: `The delivery way of the product.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of product codes.`,
				},
				resource.Attribute{
					Name:        "products",
					Description: `A list of products. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `The code of the product.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the product.`,
				},
				resource.Attribute{
					Name:        "category_id",
					Description: `The category id of the product.`,
				},
				resource.Attribute{
					Name:        "supplier_id",
					Description: `The supplier id of the product.`,
				},
				resource.Attribute{
					Name:        "supplier_name",
					Description: `The supplier name of the product.`,
				},
				resource.Attribute{
					Name:        "short_description",
					Description: `The short description of the product.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags of the product.`,
				},
				resource.Attribute{
					Name:        "suggested_price",
					Description: `The suggested price of the product.`,
				},
				resource.Attribute{
					Name:        "target_url",
					Description: `The detail page URL of the product.`,
				},
				resource.Attribute{
					Name:        "image_url",
					Description: `The image URL of the product.`,
				},
				resource.Attribute{
					Name:        "score",
					Description: `The rating information of the product.`,
				},
				resource.Attribute{
					Name:        "operation_system",
					Description: `The operation system of the product.`,
				},
				resource.Attribute{
					Name:        "warranty_date",
					Description: `The warranty date of the product.`,
				},
				resource.Attribute{
					Name:        "delivery_date",
					Description: `The delivery date of the product.`,
				},
				resource.Attribute{
					Name:        "delivery_way",
					Description: `The delivery way of the product.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_mns_queues",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of mns queues available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_prefix",
					Description: `(Optional) A string to filter resulting queues by their name prefixs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of queue names.`,
				},
				resource.Attribute{
					Name:        "queues",
					Description: `A list of queues. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the queue, The value is set to ` + "`" + `name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the queue`,
				},
				resource.Attribute{
					Name:        "delay_seconds",
					Description: `This attribute defines the length of time, in seconds, after which every message sent to the queue is dequeued.`,
				},
				resource.Attribute{
					Name:        "maximum_message_size",
					Description: `This indicates the maximum length, in bytes, of any message body sent to the queue.`,
				},
				resource.Attribute{
					Name:        "message_retention_period",
					Description: `Messages are deleted from the queue after a specified length of time, whether they have been activated or not. This attribute defines the viability period, in seconds, for every message in the queue.`,
				},
				resource.Attribute{
					Name:        "visibility_timeouts",
					Description: `Dequeued messages change from active (visible) status to inactive (invisible) status. This attribute defines the length of time, in seconds, that messages remain invisible. Messages return to active status after the set period.`,
				},
				resource.Attribute{
					Name:        "polling_wait_seconds",
					Description: `Long polling is measured in seconds. When this attribute is set to 0, long polling is disabled. When it is not set to 0, long polling is enabled and message dequeue requests will be processed only when valid messages are received or when long polling times out.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `A list of queue names.`,
				},
				resource.Attribute{
					Name:        "queues",
					Description: `A list of queues. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the queue, The value is set to ` + "`" + `name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the queue`,
				},
				resource.Attribute{
					Name:        "delay_seconds",
					Description: `This attribute defines the length of time, in seconds, after which every message sent to the queue is dequeued.`,
				},
				resource.Attribute{
					Name:        "maximum_message_size",
					Description: `This indicates the maximum length, in bytes, of any message body sent to the queue.`,
				},
				resource.Attribute{
					Name:        "message_retention_period",
					Description: `Messages are deleted from the queue after a specified length of time, whether they have been activated or not. This attribute defines the viability period, in seconds, for every message in the queue.`,
				},
				resource.Attribute{
					Name:        "visibility_timeouts",
					Description: `Dequeued messages change from active (visible) status to inactive (invisible) status. This attribute defines the length of time, in seconds, that messages remain invisible. Messages return to active status after the set period.`,
				},
				resource.Attribute{
					Name:        "polling_wait_seconds",
					Description: `Long polling is measured in seconds. When this attribute is set to 0, long polling is disabled. When it is not set to 0, long polling is enabled and message dequeue requests will be processed only when valid messages are received or when long polling times out.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_mns_topic_subscriptions",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of mns topic subscriptions available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "topic_name",
					Description: `(Required) Two topics on a single account in the same region cannot have the same name. A topic name must start with an English letter or a digit, and can contain English letters, digits, and hyphens, with the length not exceeding 256 characters.`,
				},
				resource.Attribute{
					Name:        "name_prefix",
					Description: `(Optional) A string to filter resulting subscriptions of the topic by their name prefixs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of subscription names.`,
				},
				resource.Attribute{
					Name:        "subscriptions",
					Description: `A list of subscriptions. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the topic subscription. The value is set to ` + "`" + `name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the subscription.`,
				},
				resource.Attribute{
					Name:        "notify_strategy",
					Description: `The NotifyStrategy attribute of Subscription. This attribute specifies the retry strategy when message sending fails.`,
				},
				resource.Attribute{
					Name:        "notify_content_format",
					Description: `The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `Describe the terminal address of the message received in this subscription.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `A list of subscription names.`,
				},
				resource.Attribute{
					Name:        "subscriptions",
					Description: `A list of subscriptions. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the topic subscription. The value is set to ` + "`" + `name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the subscription.`,
				},
				resource.Attribute{
					Name:        "notify_strategy",
					Description: `The NotifyStrategy attribute of Subscription. This attribute specifies the retry strategy when message sending fails.`,
				},
				resource.Attribute{
					Name:        "notify_content_format",
					Description: `The NotifyContentFormat attribute of Subscription. This attribute specifies the content format of the messages pushed to users.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `Describe the terminal address of the message received in this subscription.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_mns_topics",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of mns topics available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_prefix",
					Description: `(Optional) A string to filter resulting topics by their name prefixs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of topic names.`,
				},
				resource.Attribute{
					Name:        "topics",
					Description: `A list of topics. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the topic. The value is set to ` + "`" + `name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the topic.`,
				},
				resource.Attribute{
					Name:        "maximum_message_size",
					Description: `This indicates the maximum length, in bytes, of any message body sent to the topic.`,
				},
				resource.Attribute{
					Name:        "logging_enabled",
					Description: `Whether to enable logging.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `A list of topic names.`,
				},
				resource.Attribute{
					Name:        "topics",
					Description: `A list of topics. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the topic. The value is set to ` + "`" + `name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the topic.`,
				},
				resource.Attribute{
					Name:        "maximum_message_size",
					Description: `This indicates the maximum length, in bytes, of any message body sent to the topic.`,
				},
				resource.Attribute{
					Name:        "logging_enabled",
					Description: `Whether to enable logging.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_mongodb_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a collection of MongoDB instances according to the specified filters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to apply to the instance name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, Available 1.53.0+) The ids list of MongoDB instances`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional) Type of the instance to be queried. If it is set to ` + "`" + `sharding` + "`" + `, the sharded cluster instances are listed. If it is set to ` + "`" + `replicate` + "`" + `, replica set instances are listed. Default value ` + "`" + `replicate` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_class",
					Description: `(Optional) Sizing of the instance to be queried.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) Instance availability zone.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.66.0+) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) The name of file that can save the collection of instances after running ` + "`" + `terraform plan` + "`" + `. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `The ids list of MongoDB instances`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `The names list of MongoDB instances`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of MongoDB instances. Its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the MongoDB instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the MongoDB instance.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Billing method. Value options are ` + "`" + `PostPaid` + "`" + ` for Pay-As-You-Go and ` + "`" + `PrePaid` + "`" + ` for yearly or monthly subscription.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Instance type. Optional values ` + "`" + `sharding` + "`" + ` or ` + "`" + `replicate` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Creation time of the instance in RFC3339 format.`,
				},
				resource.Attribute{
					Name:        "expiration_time",
					Description: `Expiration time in RFC3339 format. Pay-As-You-Go instances are never expire.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "replication",
					Description: `Replication factor corresponds to number of nodes. Optional values are ` + "`" + `1` + "`" + ` for single node and ` + "`" + `3` + "`" + ` for three nodes replica set.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Database engine type. Supported option is ` + "`" + `MongoDB` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Database engine version.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `Classic network or VPC.`,
				},
				resource.Attribute{
					Name:        "instance_class",
					Description: `Sizing of the MongoDB instance.`,
				},
				resource.Attribute{
					Name:        "lock_mode",
					Description: `Lock status of the instance.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `Storage size.`,
				},
				resource.Attribute{
					Name:        "mongos",
					Description: `Array composed of Mongos.`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `Mongos instance ID.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Mongos instance description.`,
				},
				resource.Attribute{
					Name:        "class",
					Description: `Mongos instance specification.`,
				},
				resource.Attribute{
					Name:        "shards",
					Description: `Array composed of shards.`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `Shard instance ID.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Shard instance description.`,
				},
				resource.Attribute{
					Name:        "class",
					Description: `Shard instance specification.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `Shard disk.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Instance availability zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `The ids list of MongoDB instances`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `The names list of MongoDB instances`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of MongoDB instances. Its every element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the MongoDB instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the MongoDB instance.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Billing method. Value options are ` + "`" + `PostPaid` + "`" + ` for Pay-As-You-Go and ` + "`" + `PrePaid` + "`" + ` for yearly or monthly subscription.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `Instance type. Optional values ` + "`" + `sharding` + "`" + ` or ` + "`" + `replicate` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Creation time of the instance in RFC3339 format.`,
				},
				resource.Attribute{
					Name:        "expiration_time",
					Description: `Expiration time in RFC3339 format. Pay-As-You-Go instances are never expire.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the instance.`,
				},
				resource.Attribute{
					Name:        "replication",
					Description: `Replication factor corresponds to number of nodes. Optional values are ` + "`" + `1` + "`" + ` for single node and ` + "`" + `3` + "`" + ` for three nodes replica set.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Database engine type. Supported option is ` + "`" + `MongoDB` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Database engine version.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `Classic network or VPC.`,
				},
				resource.Attribute{
					Name:        "instance_class",
					Description: `Sizing of the MongoDB instance.`,
				},
				resource.Attribute{
					Name:        "lock_mode",
					Description: `Lock status of the instance.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `Storage size.`,
				},
				resource.Attribute{
					Name:        "mongos",
					Description: `Array composed of Mongos.`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `Mongos instance ID.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Mongos instance description.`,
				},
				resource.Attribute{
					Name:        "class",
					Description: `Mongos instance specification.`,
				},
				resource.Attribute{
					Name:        "shards",
					Description: `Array composed of shards.`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `Shard instance ID.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Shard instance description.`,
				},
				resource.Attribute{
					Name:        "class",
					Description: `Shard instance specification.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `Shard disk.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Instance availability zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_mongodb_zones",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of availability zones for mongoDB that can be used by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "multi",
					Description: `(Optional) Indicate whether the zones can be used in a multi AZ configuration. Default to ` + "`" + `false` + "`" + `. Multi AZ is usually used to launch MongoDB instances.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "multi_zone_ids",
					Description: `A list of zone ids in which the multi zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "multi_zone_ids",
					Description: `A list of zone ids in which the multi zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_nas_access_groups",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Access Groups owned by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Required) A regex string to filter AccessGroups by name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Filter results by a specific AccessGroupType.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Filter results by a specific Description.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of AccessGroup IDs, the value is set to ` + "`" + `names` + "`" + ` .`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of AccessGroup names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of AccessGroups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `AccessGroupName of the AccessGroup.`,
				},
				resource.Attribute{
					Name:        "rule_count",
					Description: `RuleCount of the AccessGroup.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `AccessGroupType of the AccessGroup.`,
				},
				resource.Attribute{
					Name:        "mount_target_count",
					Description: `MountTargetCount block of the AccessGroup`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Destription of the AccessGroup.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of AccessGroup IDs, the value is set to ` + "`" + `names` + "`" + ` .`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of AccessGroup names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of AccessGroups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `AccessGroupName of the AccessGroup.`,
				},
				resource.Attribute{
					Name:        "rule_count",
					Description: `RuleCount of the AccessGroup.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `AccessGroupType of the AccessGroup.`,
				},
				resource.Attribute{
					Name:        "mount_target_count",
					Description: `MountTargetCount block of the AccessGroup`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Destription of the AccessGroup.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_nas_access_rules",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of AccessRules owned by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_group_name",
					Description: `(Required ForceNew) Filter results by a specific AccessGroupName.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, Available in 1.53.0+) A list of rule IDs.`,
				},
				resource.Attribute{
					Name:        "source_cidr_ip",
					Description: `(Optional) Filter results by a specific SourceCidrIp.`,
				},
				resource.Attribute{
					Name:        "user_access",
					Description: `(Optional) Filter results by a specific UserAccess.`,
				},
				resource.Attribute{
					Name:        "rw_access",
					Description: `(Optional) Filter results by a specific RWAccess.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of rule IDs, Each element set to ` + "`" + `access_rule_id` + "`" + ` (Each element formats as ` + "`" + `<access_group_name>:<access rule id>` + "`" + ` before 1.53.0).`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `A list of AccessRules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "source_cidr_ip",
					Description: `SourceCidrIp of the AccessRule.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Priority of the AccessRule.`,
				},
				resource.Attribute{
					Name:        "access_rule_id",
					Description: `AccessRuleId of the AccessRule.`,
				},
				resource.Attribute{
					Name:        "user_access",
					Description: `UserAccess of the AccessRule`,
				},
				resource.Attribute{
					Name:        "rw_access",
					Description: `RWAccess of the AccessRule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of rule IDs, Each element set to ` + "`" + `access_rule_id` + "`" + ` (Each element formats as ` + "`" + `<access_group_name>:<access rule id>` + "`" + ` before 1.53.0).`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `A list of AccessRules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "source_cidr_ip",
					Description: `SourceCidrIp of the AccessRule.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Priority of the AccessRule.`,
				},
				resource.Attribute{
					Name:        "access_rule_id",
					Description: `AccessRuleId of the AccessRule.`,
				},
				resource.Attribute{
					Name:        "user_access",
					Description: `UserAccess of the AccessRule`,
				},
				resource.Attribute{
					Name:        "rw_access",
					Description: `RWAccess of the AccessRule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_nas_file_systems",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of FileSystems owned by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of FileSystemId.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Optional) Filter results by a specific StorageType.`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `(Optional) Filter results by a specific ProtocolType.`,
				},
				resource.Attribute{
					Name:        "description_regex",
					Description: `(Optional) A regex string to filter the results by the ：FileSystem description.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of FileSystem Id.`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `A list of FileSystem descriptions.`,
				},
				resource.Attribute{
					Name:        "systems",
					Description: `A list of VPCs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the FileSystem.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `ID of the region where the FileSystem is located.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Destription of the FileSystem.`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `ProtocolType block of the FileSystem`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `StorageType block of the FileSystem.`,
				},
				resource.Attribute{
					Name:        "metered_size",
					Description: `MeteredSize of the FileSystem.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Time of creation.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of FileSystem Id.`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `A list of FileSystem descriptions.`,
				},
				resource.Attribute{
					Name:        "systems",
					Description: `A list of VPCs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the FileSystem.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `ID of the region where the FileSystem is located.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Destription of the FileSystem.`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `ProtocolType block of the FileSystem`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `StorageType block of the FileSystem.`,
				},
				resource.Attribute{
					Name:        "metered_size",
					Description: `MeteredSize of the FileSystem.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Time of creation.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_nas_mount_targets",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of mns topic subscriptions available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "file_system_id",
					Description: `(Required ForceNew) The ID of the FileSystem that owns the MountTarget.`,
				},
				resource.Attribute{
					Name:        "access_group_name",
					Description: `(Optional) Filter results by a specific AccessGroupName.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Filter results by a specific NetworkType.`,
				},
				resource.Attribute{
					Name:        "mount_target_domain",
					Description: `(Deprecated, Optional) Filter results by a specific MountTargetDomain.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) Filter results by a specific VpcId.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional) Filter results by a specific VSwitchId.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, Available 1.53.0+) A list of MountTargetDomain.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of MountTargetDomain.`,
				},
				resource.Attribute{
					Name:        "targets",
					Description: `A list of MountTargetDomains. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the MountTargetDomain.`,
				},
				resource.Attribute{
					Name:        "mount_target_domain",
					Description: `MountTargetDomain of the MountTarget.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VpcId of The MountTarget.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `VSwitchId of The MountTarget.`,
				},
				resource.Attribute{
					Name:        "access_group_name",
					Description: `AccessGroup of The MountTarget.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of MountTargetDomain.`,
				},
				resource.Attribute{
					Name:        "targets",
					Description: `A list of MountTargetDomains. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the MountTargetDomain.`,
				},
				resource.Attribute{
					Name:        "mount_target_domain",
					Description: `MountTargetDomain of the MountTarget.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VpcId of The MountTarget.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `VSwitchId of The MountTarget.`,
				},
				resource.Attribute{
					Name:        "access_group_name",
					Description: `AccessGroup of The MountTarget.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_nas_protocols",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of FileType owned by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The file system type. Valid Values: Performance and Capacity.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional) String to filter results by zone id.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "protocols",
					Description: `A list of supported protocol type..`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "protocols",
					Description: `A list of supported protocol type..`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_nat_gateways",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Nat Gateways owned by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of NAT gateways IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter nat gateways by name.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Nat gateways IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Nat gateways names.`,
				},
				resource.Attribute{
					Name:        "gateways",
					Description: `A list of Nat gateways. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Time of creation.`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `The specification of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "snat_table_id",
					Description: `The snat table id.`,
				},
				resource.Attribute{
					Name:        "forward_table_id",
					Description: `The forward table id.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "ip_lists",
					Description: `The ip address of the bind eip.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Nat gateways IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Nat gateways names.`,
				},
				resource.Attribute{
					Name:        "gateways",
					Description: `A list of Nat gateways. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Time of creation.`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `The specification of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "snat_table_id",
					Description: `The snat table id.`,
				},
				resource.Attribute{
					Name:        "forward_table_id",
					Description: `The forward table id.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "ip_lists",
					Description: `The ip address of the bind eip.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_network_interfaces",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to get a list of elastic network interfaces according to the specified filters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "interfaces",
					Description: `A list of ENIs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the ENI.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the ENI.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC that the ENI belongs to.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `ID of the VSwitch that the ENI is linked to.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `ID of the availability zone that the ENI belongs to.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP of the ENI.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Primary private IP of the ENI.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `A list of secondary private IP address that is assigned to the ENI.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `MAC address of the ENI.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `A list of security group that the ENI belongs to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the ENI.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the ENI.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the instance that the ENI is attached to.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Creation time of the ENI.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the ENI.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The Id of resource group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ons_groups",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of ons groups available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) ID of the ONS Instance that owns the groups.`,
				},
				resource.Attribute{
					Name:        "group_id_regex",
					Description: `(Optional) A regex string to filter results by the group name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of group names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The name of the group.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `The ID of the group owner, which is the Alibaba Cloud UID.`,
				},
				resource.Attribute{
					Name:        "independent_naming",
					Description: `Indicates whether namespaces are available. Read [Fields in SubscribeInfoDo](https://www.alibabacloud.com/help/doc-detail/29619.html) for further details.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Remark of the group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of group names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The name of the group.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `The ID of the group owner, which is the Alibaba Cloud UID.`,
				},
				resource.Attribute{
					Name:        "independent_naming",
					Description: `Indicates whether namespaces are available. Read [Fields in SubscribeInfoDo](https://www.alibabacloud.com/help/doc-detail/29619.html) for further details.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Remark of the group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ons_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of ons instances available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of instance IDs to filter results.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by the instance name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of instance IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of instance names.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `Name of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The type of the instance. Read [Fields in InstanceVO](https://www.alibabacloud.com/help/doc-detail/106351.html) for further details.`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `The status of the instance. Read [Fields in InstanceVO](https://www.alibabacloud.com/help/doc-detail/106351.html) for further details.`,
				},
				resource.Attribute{
					Name:        "release_time",
					Description: `The automatic release time of an Enterprise Platinum Edition instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of instance IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of instance names.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `Name of the instance.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `The type of the instance. Read [Fields in InstanceVO](https://www.alibabacloud.com/help/doc-detail/106351.html) for further details.`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `The status of the instance. Read [Fields in InstanceVO](https://www.alibabacloud.com/help/doc-detail/106351.html) for further details.`,
				},
				resource.Attribute{
					Name:        "release_time",
					Description: `The automatic release time of an Enterprise Platinum Edition instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ons_topics",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of ons topics available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) ID of the ONS Instance that owns the topics.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by the topic name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of topic names.`,
				},
				resource.Attribute{
					Name:        "topics",
					Description: `A list of topics. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "topic",
					Description: `The name of the topic.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `The ID of the topic owner, which is the Alibaba Cloud UID.`,
				},
				resource.Attribute{
					Name:        "relation",
					Description: `The relation ID. Read [Fields in PublishInfoDo](https://www.alibabacloud.com/help/doc-detail/29590.html) for further details.`,
				},
				resource.Attribute{
					Name:        "relation_name",
					Description: `The name of the relation, for example, owner, publishable, subscribable, and publishable and subscribable.`,
				},
				resource.Attribute{
					Name:        "message_type",
					Description: `The type of the message. Read [Fields in PublishInfoDo](https://www.alibabacloud.com/help/doc-detail/29590.html) for further details.`,
				},
				resource.Attribute{
					Name:        "independent_naming",
					Description: `Indicates whether namespaces are available. Read [Fields in PublishInfoDo](https://www.alibabacloud.com/help/doc-detail/29590.html) for further details.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Time of creation.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Remark of the topic.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `A list of topic names.`,
				},
				resource.Attribute{
					Name:        "topics",
					Description: `A list of topics. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "topic",
					Description: `The name of the topic.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `The ID of the topic owner, which is the Alibaba Cloud UID.`,
				},
				resource.Attribute{
					Name:        "relation",
					Description: `The relation ID. Read [Fields in PublishInfoDo](https://www.alibabacloud.com/help/doc-detail/29590.html) for further details.`,
				},
				resource.Attribute{
					Name:        "relation_name",
					Description: `The name of the relation, for example, owner, publishable, subscribable, and publishable and subscribable.`,
				},
				resource.Attribute{
					Name:        "message_type",
					Description: `The type of the message. Read [Fields in PublishInfoDo](https://www.alibabacloud.com/help/doc-detail/29590.html) for further details.`,
				},
				resource.Attribute{
					Name:        "independent_naming",
					Description: `Indicates whether namespaces are available. Read [Fields in PublishInfoDo](https://www.alibabacloud.com/help/doc-detail/29590.html) for further details.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Time of creation.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Remark of the topic.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_oos_templates",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of OOS Templates.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of OOS Template ids. Each element in the list is same as template_name.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter the results by the template_name.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Optional) The category of template.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `(Optional) The creator of the template.`,
				},
				resource.Attribute{
					Name:        "created_date",
					Description: `(Optional) The template whose creation time is less than or equal to the specified time. The format is: YYYY-MM-DDThh:mm::ssZ.`,
				},
				resource.Attribute{
					Name:        "has_trigger",
					Description: `(Optional) Is it triggered successfully.`,
				},
				resource.Attribute{
					Name:        "share_type",
					Description: `(Optional) The sharing type of the template. Valid values: ` + "`" + `Private` + "`" + `, ` + "`" + `Public` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "template_format",
					Description: `(Optional) The format of the template. Valid values: ` + "`" + `JSON` + "`" + `, ` + "`" + `YAML` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "template_type",
					Description: `(Optional) The type of OOS Template.`,
				},
				resource.Attribute{
					Name:        "sort_field",
					Description: `(Optional) Sort field. Valid values: ` + "`" + `TotalExecutionCount` + "`" + `, ` + "`" + `Popularity` + "`" + `, ` + "`" + `TemplateName` + "`" + ` and ` + "`" + `CreatedDate` + "`" + `. Default to ` + "`" + `TotalExecutionCount` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sort_order",
					Description: `(Optional) Sort order. Valid values: ` + "`" + `Ascending` + "`" + `, ` + "`" + `Descending` + "`" + `. Default to ` + "`" + `Descending` + "`" + ``,
				},
				resource.Attribute{
					Name:        "created_date_after",
					Description: `(Optional) Create a template whose time is greater than or equal to the specified time. The format is: YYYY-MM-DDThh:mm:ssZ.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of OOS Template ids. Each element in the list is same as template_name.`,
				},
				resource.Attribute{
					Name:        "templates",
					Description: `A list of OOS Templates. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the OOS Template. The value is same as template_name.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `Name of the OOS Template.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the OOS Template.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `ID of the OOS Template resource.`,
				},
				resource.Attribute{
					Name:        "template_version",
					Description: `Version of the OOS Template.`,
				},
				resource.Attribute{
					Name:        "updated_by",
					Description: `The user who updated the template.`,
				},
				resource.Attribute{
					Name:        "updated_date",
					Description: `The time when the template was updated.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of OOS Template ids. Each element in the list is same as template_name.`,
				},
				resource.Attribute{
					Name:        "templates",
					Description: `A list of OOS Templates. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the OOS Template. The value is same as template_name.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `Name of the OOS Template.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the OOS Template.`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `ID of the OOS Template resource.`,
				},
				resource.Attribute{
					Name:        "template_version",
					Description: `Version of the OOS Template.`,
				},
				resource.Attribute{
					Name:        "updated_by",
					Description: `The user who updated the template.`,
				},
				resource.Attribute{
					Name:        "updated_date",
					Description: `The time when the template was updated.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_oss_bucket_objects",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of bucket objects to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket_name",
					Description: `Name of the bucket that contains the objects to find.`,
				},
				resource.Attribute{
					Name:        "key_regex",
					Description: `(Optional) A regex string to filter results by key.`,
				},
				resource.Attribute{
					Name:        "key_prefix",
					Description: `(Optional) Filter results by the given key prefix (such as "path/to/folder/logs-").`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "objects",
					Description: `A list of bucket objects. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Object key.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `Object access control list. Possible values: ` + "`" + `default` + "`" + `, ` + "`" + `private` + "`" + `, ` + "`" + `public-read` + "`" + ` and ` + "`" + `public-read-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `Standard MIME type describing the format of the object data, e.g. "application/octet-stream".`,
				},
				resource.Attribute{
					Name:        "content_length",
					Description: `Size of the object in bytes.`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `Caching behavior along the request/reply chain. Read [RFC2616 Cache-Control](https://www.ietf.org/rfc/rfc2616.txt) for further details.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `Presentational information for the object. Read [RFC2616 Content-Disposition](https://www.ietf.org/rfc/rfc2616.txt) for further details.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `Content encodings that have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [RFC2616 Content-Encoding](https://www.ietf.org/rfc/rfc2616.txt) for further details.`,
				},
				resource.Attribute{
					Name:        "content_md5",
					Description: `MD5 value of the content. Read [MD5](https://www.alibabacloud.com/help/doc-detail/31978.htm) for computing method.`,
				},
				resource.Attribute{
					Name:        "expires",
					Description: `Expiration date for the the request/response. Read [RFC2616 Expires](https://www.ietf.org/rfc/rfc2616.txt) for further details.`,
				},
				resource.Attribute{
					Name:        "server_side_encryption",
					Description: `Server-side encryption of the object in OSS. It can be empty or ` + "`" + `AES256` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sse_kms_key_id",
					Description: `If present, specifies the ID of the Key Management Service(KMS) master encryption key that was used for the object.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `ETag generated for the object (MD5 sum of the object content).`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `Object storage type. Possible values: ` + "`" + `Standard` + "`" + `, ` + "`" + `IA` + "`" + ` and ` + "`" + `Archive` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "last_modification_time",
					Description: `Last modification time of the object.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "objects",
					Description: `A list of bucket objects. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Object key.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `Object access control list. Possible values: ` + "`" + `default` + "`" + `, ` + "`" + `private` + "`" + `, ` + "`" + `public-read` + "`" + ` and ` + "`" + `public-read-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `Standard MIME type describing the format of the object data, e.g. "application/octet-stream".`,
				},
				resource.Attribute{
					Name:        "content_length",
					Description: `Size of the object in bytes.`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `Caching behavior along the request/reply chain. Read [RFC2616 Cache-Control](https://www.ietf.org/rfc/rfc2616.txt) for further details.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `Presentational information for the object. Read [RFC2616 Content-Disposition](https://www.ietf.org/rfc/rfc2616.txt) for further details.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `Content encodings that have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [RFC2616 Content-Encoding](https://www.ietf.org/rfc/rfc2616.txt) for further details.`,
				},
				resource.Attribute{
					Name:        "content_md5",
					Description: `MD5 value of the content. Read [MD5](https://www.alibabacloud.com/help/doc-detail/31978.htm) for computing method.`,
				},
				resource.Attribute{
					Name:        "expires",
					Description: `Expiration date for the the request/response. Read [RFC2616 Expires](https://www.ietf.org/rfc/rfc2616.txt) for further details.`,
				},
				resource.Attribute{
					Name:        "server_side_encryption",
					Description: `Server-side encryption of the object in OSS. It can be empty or ` + "`" + `AES256` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sse_kms_key_id",
					Description: `If present, specifies the ID of the Key Management Service(KMS) master encryption key that was used for the object.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `ETag generated for the object (MD5 sum of the object content).`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `Object storage type. Possible values: ` + "`" + `Standard` + "`" + `, ` + "`" + `IA` + "`" + ` and ` + "`" + `Archive` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "last_modification_time",
					Description: `Last modification time of the object.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_oss_buckets",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of OSS buckets to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by bucket name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of bucket names.`,
				},
				resource.Attribute{
					Name:        "buckets",
					Description: `A list of buckets. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Bucket name.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `Bucket access control list. Possible values: ` + "`" + `private` + "`" + `, ` + "`" + `public-read` + "`" + ` and ` + "`" + `public-read-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extranet_endpoint",
					Description: `Internet domain name for accessing the bucket from outside.`,
				},
				resource.Attribute{
					Name:        "intranet_endpoint",
					Description: `Intranet domain name for accessing the bucket from an ECS instance in the same region.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Region of the data center where the bucket is located.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Bucket owner.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `Object storage type. Possible values: ` + "`" + `Standard` + "`" + `, ` + "`" + `IA` + "`" + ` and ` + "`" + `Archive` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "redundancy_type",
					Description: `Redundancy type. Possible values: ` + "`" + `LRS` + "`" + `, and ` + "`" + `ZRS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Bucket creation date.`,
				},
				resource.Attribute{
					Name:        "cors_rules",
					Description: `A list of CORS rule configurations. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "allowed_origins",
					Description: `The origins allowed for cross-domain requests. Multiple elements can be used to specify multiple allowed origins. Each rule allows up to one wildcard "\`,
				},
				resource.Attribute{
					Name:        "allowed_methods",
					Description: `Specify the allowed methods for cross-domain requests. Possible values: ` + "`" + `GET` + "`" + `, ` + "`" + `PUT` + "`" + `, ` + "`" + `DELETE` + "`" + `, ` + "`" + `POST` + "`" + ` and ` + "`" + `HEAD` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "allowed_headers",
					Description: `Control whether the headers specified by Access-Control-Request-Headers in the OPTIONS prefetch command are allowed. Each header specified by Access-Control-Request-Headers must match a value in AllowedHeader. Each rule allows up to one wildcard “`,
				},
				resource.Attribute{
					Name:        "expose_headers",
					Description: `Specify the response headers allowing users to access from an application (for example, a Javascript XMLHttpRequest object). The wildcard "\`,
				},
				resource.Attribute{
					Name:        "max_age_seconds",
					Description: `Specify the cache time for the returned result of a browser prefetch (OPTIONS) request to a specific resource.`,
				},
				resource.Attribute{
					Name:        "website",
					Description: `A list of one element containing configuration parameters used when the bucket is used as a website. It contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "index_document",
					Description: `Key of the HTML document containing the home page.`,
				},
				resource.Attribute{
					Name:        "error_document",
					Description: `Key of the HTML document containing the error page.`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `A list of one element containing configuration parameters used for storing access log information. It contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "target_bucket",
					Description: `Bucket for storing access logs.`,
				},
				resource.Attribute{
					Name:        "target_prefix",
					Description: `Prefix of the saved access log file paths.`,
				},
				resource.Attribute{
					Name:        "referer_config",
					Description: `A list of one element containing referer configuration. It contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "allow_empty",
					Description: `Indicate whether the access request referer field can be empty.`,
				},
				resource.Attribute{
					Name:        "referers",
					Description: `Referer access whitelist.`,
				},
				resource.Attribute{
					Name:        "lifecycle_rule",
					Description: `A list CORS of lifecycle configurations. When Lifecycle is enabled, OSS automatically deletes the objects or transitions the objects (to another storage class) corresponding the lifecycle rules on a regular basis. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique ID of the rule.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `Prefix applicable to a rule. Only those objects with a matching prefix can be affected by the rule.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Indicate whether the rule is enabled or not.`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `A list of one element containing expiration attributes of an object. It contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `Date after which the rule to take effect. The format is like 2017-03-09.`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `Indicate the number of days after the last object update until the rules take effect.`,
				},
				resource.Attribute{
					Name:        "server_side_encryption_rule",
					Description: `A configuration of default encryption for a bucket. It contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "sse_algorithm",
					Description: `The server-side encryption algorithm to use.`,
				},
				resource.Attribute{
					Name:        "kms_master_key_id",
					Description: `The alibaba cloud KMS master key ID used for the SSE-KMS encryption.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags.`,
				},
				resource.Attribute{
					Name:        "versioning",
					Description: `If present , the versioning state has been set on the bucket. It contains the following attribute.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `A bucket versioning state. Possible values:` + "`" + `Enabled` + "`" + ` and ` + "`" + `Suspended` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `A list of bucket names.`,
				},
				resource.Attribute{
					Name:        "buckets",
					Description: `A list of buckets. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Bucket name.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `Bucket access control list. Possible values: ` + "`" + `private` + "`" + `, ` + "`" + `public-read` + "`" + ` and ` + "`" + `public-read-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extranet_endpoint",
					Description: `Internet domain name for accessing the bucket from outside.`,
				},
				resource.Attribute{
					Name:        "intranet_endpoint",
					Description: `Intranet domain name for accessing the bucket from an ECS instance in the same region.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Region of the data center where the bucket is located.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Bucket owner.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `Object storage type. Possible values: ` + "`" + `Standard` + "`" + `, ` + "`" + `IA` + "`" + ` and ` + "`" + `Archive` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "redundancy_type",
					Description: `Redundancy type. Possible values: ` + "`" + `LRS` + "`" + `, and ` + "`" + `ZRS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Bucket creation date.`,
				},
				resource.Attribute{
					Name:        "cors_rules",
					Description: `A list of CORS rule configurations. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "allowed_origins",
					Description: `The origins allowed for cross-domain requests. Multiple elements can be used to specify multiple allowed origins. Each rule allows up to one wildcard "\`,
				},
				resource.Attribute{
					Name:        "allowed_methods",
					Description: `Specify the allowed methods for cross-domain requests. Possible values: ` + "`" + `GET` + "`" + `, ` + "`" + `PUT` + "`" + `, ` + "`" + `DELETE` + "`" + `, ` + "`" + `POST` + "`" + ` and ` + "`" + `HEAD` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "allowed_headers",
					Description: `Control whether the headers specified by Access-Control-Request-Headers in the OPTIONS prefetch command are allowed. Each header specified by Access-Control-Request-Headers must match a value in AllowedHeader. Each rule allows up to one wildcard “`,
				},
				resource.Attribute{
					Name:        "expose_headers",
					Description: `Specify the response headers allowing users to access from an application (for example, a Javascript XMLHttpRequest object). The wildcard "\`,
				},
				resource.Attribute{
					Name:        "max_age_seconds",
					Description: `Specify the cache time for the returned result of a browser prefetch (OPTIONS) request to a specific resource.`,
				},
				resource.Attribute{
					Name:        "website",
					Description: `A list of one element containing configuration parameters used when the bucket is used as a website. It contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "index_document",
					Description: `Key of the HTML document containing the home page.`,
				},
				resource.Attribute{
					Name:        "error_document",
					Description: `Key of the HTML document containing the error page.`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `A list of one element containing configuration parameters used for storing access log information. It contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "target_bucket",
					Description: `Bucket for storing access logs.`,
				},
				resource.Attribute{
					Name:        "target_prefix",
					Description: `Prefix of the saved access log file paths.`,
				},
				resource.Attribute{
					Name:        "referer_config",
					Description: `A list of one element containing referer configuration. It contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "allow_empty",
					Description: `Indicate whether the access request referer field can be empty.`,
				},
				resource.Attribute{
					Name:        "referers",
					Description: `Referer access whitelist.`,
				},
				resource.Attribute{
					Name:        "lifecycle_rule",
					Description: `A list CORS of lifecycle configurations. When Lifecycle is enabled, OSS automatically deletes the objects or transitions the objects (to another storage class) corresponding the lifecycle rules on a regular basis. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique ID of the rule.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `Prefix applicable to a rule. Only those objects with a matching prefix can be affected by the rule.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Indicate whether the rule is enabled or not.`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `A list of one element containing expiration attributes of an object. It contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `Date after which the rule to take effect. The format is like 2017-03-09.`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `Indicate the number of days after the last object update until the rules take effect.`,
				},
				resource.Attribute{
					Name:        "server_side_encryption_rule",
					Description: `A configuration of default encryption for a bucket. It contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "sse_algorithm",
					Description: `The server-side encryption algorithm to use.`,
				},
				resource.Attribute{
					Name:        "kms_master_key_id",
					Description: `The alibaba cloud KMS master key ID used for the SSE-KMS encryption.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags.`,
				},
				resource.Attribute{
					Name:        "versioning",
					Description: `If present , the versioning state has been set on the bucket. It contains the following attribute.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `A bucket versioning state. Possible values:` + "`" + `Enabled` + "`" + ` and ` + "`" + `Suspended` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ots_instance_attachments",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of ots instance attachments to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Required) The name of OTS instance.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by vpc name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of vpc names.`,
				},
				resource.Attribute{
					Name:        "vpc_ids",
					Description: `A list of vpc ids.`,
				},
				resource.Attribute{
					Name:        "attachments",
					Description: `A list of instance attachments. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID, the value is same as "instance_name".`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `The domain of the instance attachment.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The access endpoint of the instance attachment.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region of the instance attachment.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `The instance name.`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `The name of attaching VPC to instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of attaching VPC to instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `A list of vpc names.`,
				},
				resource.Attribute{
					Name:        "vpc_ids",
					Description: `A list of vpc ids.`,
				},
				resource.Attribute{
					Name:        "attachments",
					Description: `A list of instance attachments. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID, the value is same as "instance_name".`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `The domain of the instance attachment.`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The access endpoint of the instance attachment.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region of the instance attachment.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `The instance name.`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `The name of attaching VPC to instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of attaching VPC to instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ots_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of ots instances to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of instance IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by instance name.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A map of tags assigned to the instance. It must be in the format: ` + "`" + `` + "`" + `` + "`" + ` data "alicloud_ots_instances" "instances_ds" { tags = { tagKey1 = "tagValue1", tagKey2 = "tagValue2" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of instance IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of instance names.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Instance name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Instance status. Possible values: ` + "`" + `Running` + "`" + `, ` + "`" + `Disabled` + "`" + `, ` + "`" + `Deleting` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "write_capacity",
					Description: `The maximum adjustable write capacity unit of the instance.`,
				},
				resource.Attribute{
					Name:        "read_capacity",
					Description: `The maximum adjustable read capacity unit of the instance.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `The cluster type of the instance. Possible values: ` + "`" + `SSD` + "`" + `, ` + "`" + `HYBRID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of the instance.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The user id of the instance.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The network type of the instance. Possible values: ` + "`" + `NORMAL` + "`" + `, ` + "`" + `VPC` + "`" + `, ` + "`" + `VPC_CONSOLE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the instance.`,
				},
				resource.Attribute{
					Name:        "entity_quota",
					Description: `The instance quota which indicating the maximum number of tables.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags of the instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of instance IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of instance names.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Instance name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Instance status. Possible values: ` + "`" + `Running` + "`" + `, ` + "`" + `Disabled` + "`" + `, ` + "`" + `Deleting` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "write_capacity",
					Description: `The maximum adjustable write capacity unit of the instance.`,
				},
				resource.Attribute{
					Name:        "read_capacity",
					Description: `The maximum adjustable read capacity unit of the instance.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `The cluster type of the instance. Possible values: ` + "`" + `SSD` + "`" + `, ` + "`" + `HYBRID` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create time of the instance.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The user id of the instance.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `The network type of the instance. Possible values: ` + "`" + `NORMAL` + "`" + `, ` + "`" + `VPC` + "`" + `, ` + "`" + `VPC_CONSOLE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the instance.`,
				},
				resource.Attribute{
					Name:        "entity_quota",
					Description: `The instance quota which indicating the maximum number of tables.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags of the instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ots_tables",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of ots tables to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_name",
					Description: `The name of OTS instance.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of table IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by table name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of table IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of table names.`,
				},
				resource.Attribute{
					Name:        "tables",
					Description: `A list of tables. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the table. The value is ` + "`" + `<instance_name>:<table_name>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `The OTS instance name.`,
				},
				resource.Attribute{
					Name:        "table_name",
					Description: `The table name of the OTS which could not be changed.`,
				},
				resource.Attribute{
					Name:        "primary_key",
					Description: `The property of ` + "`" + `TableMeta` + "`" + ` which indicates the structure information of a table.`,
				},
				resource.Attribute{
					Name:        "time_to_live",
					Description: `The retention time of data stored in this table.`,
				},
				resource.Attribute{
					Name:        "max_version",
					Description: `The maximum number of versions stored in this table.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of table IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of table names.`,
				},
				resource.Attribute{
					Name:        "tables",
					Description: `A list of tables. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the table. The value is ` + "`" + `<instance_name>:<table_name>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `The OTS instance name.`,
				},
				resource.Attribute{
					Name:        "table_name",
					Description: `The table name of the OTS which could not be changed.`,
				},
				resource.Attribute{
					Name:        "primary_key",
					Description: `The property of ` + "`" + `TableMeta` + "`" + ` which indicates the structure information of a table.`,
				},
				resource.Attribute{
					Name:        "time_to_live",
					Description: `The retention time of data stored in this table.`,
				},
				resource.Attribute{
					Name:        "max_version",
					Description: `The maximum number of versions stored in this table.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_polardb_accounts",
			Category:         "Data Sources",
			ShortDescription: `Provides a collection of PolarDB endpoints according to the specified filters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "db_cluster_id",
					Description: `(Required) The polarDB cluster ID.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by account name. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `Account name of the cluster.`,
				},
				resource.Attribute{
					Name:        "accounts",
					Description: `A list of PolarDB cluster accounts. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "account_description",
					Description: `Account description.`,
				},
				resource.Attribute{
					Name:        "account_lock_state",
					Description: `Account lock state, Valid values are ` + "`" + `Lock` + "`" + `, ` + "`" + `UnLock` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `Account name.`,
				},
				resource.Attribute{
					Name:        "account_status",
					Description: `Cluster address type.` + "`" + `Cluster` + "`" + `: the default address of the Cluster.` + "`" + `Primary` + "`" + `: Primary address.` + "`" + `Custom` + "`" + `: Custom cluster addresses.`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `Account type, Valid values are ` + "`" + `Normal` + "`" + `, ` + "`" + `Super` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "database_privileges",
					Description: `A list of database privilege. Each element contains the following attributes.`,
				},
				resource.Attribute{
					Name:        "account_privilege",
					Description: `Account privilege of database`,
				},
				resource.Attribute{
					Name:        "db_name",
					Description: `The account owned database name`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `Account name of the cluster.`,
				},
				resource.Attribute{
					Name:        "accounts",
					Description: `A list of PolarDB cluster accounts. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "account_description",
					Description: `Account description.`,
				},
				resource.Attribute{
					Name:        "account_lock_state",
					Description: `Account lock state, Valid values are ` + "`" + `Lock` + "`" + `, ` + "`" + `UnLock` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `Account name.`,
				},
				resource.Attribute{
					Name:        "account_status",
					Description: `Cluster address type.` + "`" + `Cluster` + "`" + `: the default address of the Cluster.` + "`" + `Primary` + "`" + `: Primary address.` + "`" + `Custom` + "`" + `: Custom cluster addresses.`,
				},
				resource.Attribute{
					Name:        "account_type",
					Description: `Account type, Valid values are ` + "`" + `Normal` + "`" + `, ` + "`" + `Super` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "database_privileges",
					Description: `A list of database privilege. Each element contains the following attributes.`,
				},
				resource.Attribute{
					Name:        "account_privilege",
					Description: `Account privilege of database`,
				},
				resource.Attribute{
					Name:        "db_name",
					Description: `The account owned database name`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_polardb_clusters",
			Category:         "Data Sources",
			ShortDescription: `Provides a collection of PolarDB clusters according to the specified filters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description_regex",
					Description: `(Optional) A regex string to filter results by cluster description.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of PolarDB cluster IDs.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) status of the cluster.`,
				},
				resource.Attribute{
					Name:        "db_type",
					Description: `(Optional) Database type. Options are ` + "`" + `MySQL` + "`" + `, ` + "`" + `Oracle` + "`" + ` and ` + "`" + `PostgreSQL` + "`" + `. If no value is specified, all types are returned.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.68.0+) A mapping of tags to assign to the resource. - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string. - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of RDS cluster IDs.`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `A list of RDS cluster descriptions.`,
				},
				resource.Attribute{
					Name:        "clusters",
					Description: `A list of PolarDB clusters. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the PolarDB cluster.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the PolarDB cluster.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Billing method. Value options: ` + "`" + `PostPaid` + "`" + ` for Pay-As-You-Go and ` + "`" + `PrePaid` + "`" + ` for subscription.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `The DBClusterNetworkType of the PolarDB cluster.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the cluster belongs to.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `The ZoneId of the PolarDB cluster.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `Expiration time. Pay-As-You-Go clusters never expire.`,
				},
				resource.Attribute{
					Name:        "expired",
					Description: `The expired of the PolarDB cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cluster.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Database type. Options are ` + "`" + `MySQL` + "`" + `, ` + "`" + `Oracle` + "`" + ` and ` + "`" + `PostgreSQL` + "`" + `. If no value is specified, all types are returned.`,
				},
				resource.Attribute{
					Name:        "db_type",
					Description: `` + "`" + `Primary` + "`" + ` for primary cluster, ` + "`" + `ReadOnly` + "`" + ` for read-only cluster, ` + "`" + `Guard` + "`" + ` for disaster recovery cluster, and ` + "`" + `Temp` + "`" + ` for temporary cluster.`,
				},
				resource.Attribute{
					Name:        "db_version",
					Description: `The DBVersion of the PolarDB cluster.`,
				},
				resource.Attribute{
					Name:        "lock_mode",
					Description: `The LockMode of the PolarDB cluster.`,
				},
				resource.Attribute{
					Name:        "delete_lock",
					Description: `The DeleteLock of the PolarDB cluster.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The CreateTime of the PolarDB cluster.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC the cluster belongs to.`,
				},
				resource.Attribute{
					Name:        "db_node_number",
					Description: `The DBNodeNumber of the PolarDB cluster.`,
				},
				resource.Attribute{
					Name:        "db_node_class",
					Description: `The DBNodeClass of the PolarDB cluster.`,
				},
				resource.Attribute{
					Name:        "storage_used",
					Description: `The StorageUsed of the PolarDB cluster.`,
				},
				resource.Attribute{
					Name:        "db_nodes",
					Description: `The DBNodes of the PolarDB cluster.`,
				},
				resource.Attribute{
					Name:        "db_node_class",
					Description: `The db_node_class of the db_nodes.`,
				},
				resource.Attribute{
					Name:        "max_iops",
					Description: `The max_iops of the db_nodes.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The region_id of the db_nodes.`,
				},
				resource.Attribute{
					Name:        "db_node_role",
					Description: `The db_node_role of the db_nodes.`,
				},
				resource.Attribute{
					Name:        "max_connections",
					Description: `The max_connections of the db_nodes.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `The zone_id of the db_nodes.`,
				},
				resource.Attribute{
					Name:        "db_node_status",
					Description: `The db_node_status of the db_nodes.`,
				},
				resource.Attribute{
					Name:        "db_node_id",
					Description: `The db_node_id of the db_nodes.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create_time of the db_nodes.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of RDS cluster IDs.`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `A list of RDS cluster descriptions.`,
				},
				resource.Attribute{
					Name:        "clusters",
					Description: `A list of PolarDB clusters. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the PolarDB cluster.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the PolarDB cluster.`,
				},
				resource.Attribute{
					Name:        "charge_type",
					Description: `Billing method. Value options: ` + "`" + `PostPaid` + "`" + ` for Pay-As-You-Go and ` + "`" + `PrePaid` + "`" + ` for subscription.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `The DBClusterNetworkType of the PolarDB cluster.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the cluster belongs to.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `The ZoneId of the PolarDB cluster.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `Expiration time. Pay-As-You-Go clusters never expire.`,
				},
				resource.Attribute{
					Name:        "expired",
					Description: `The expired of the PolarDB cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cluster.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Database type. Options are ` + "`" + `MySQL` + "`" + `, ` + "`" + `Oracle` + "`" + ` and ` + "`" + `PostgreSQL` + "`" + `. If no value is specified, all types are returned.`,
				},
				resource.Attribute{
					Name:        "db_type",
					Description: `` + "`" + `Primary` + "`" + ` for primary cluster, ` + "`" + `ReadOnly` + "`" + ` for read-only cluster, ` + "`" + `Guard` + "`" + ` for disaster recovery cluster, and ` + "`" + `Temp` + "`" + ` for temporary cluster.`,
				},
				resource.Attribute{
					Name:        "db_version",
					Description: `The DBVersion of the PolarDB cluster.`,
				},
				resource.Attribute{
					Name:        "lock_mode",
					Description: `The LockMode of the PolarDB cluster.`,
				},
				resource.Attribute{
					Name:        "delete_lock",
					Description: `The DeleteLock of the PolarDB cluster.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The CreateTime of the PolarDB cluster.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC the cluster belongs to.`,
				},
				resource.Attribute{
					Name:        "db_node_number",
					Description: `The DBNodeNumber of the PolarDB cluster.`,
				},
				resource.Attribute{
					Name:        "db_node_class",
					Description: `The DBNodeClass of the PolarDB cluster.`,
				},
				resource.Attribute{
					Name:        "storage_used",
					Description: `The StorageUsed of the PolarDB cluster.`,
				},
				resource.Attribute{
					Name:        "db_nodes",
					Description: `The DBNodes of the PolarDB cluster.`,
				},
				resource.Attribute{
					Name:        "db_node_class",
					Description: `The db_node_class of the db_nodes.`,
				},
				resource.Attribute{
					Name:        "max_iops",
					Description: `The max_iops of the db_nodes.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The region_id of the db_nodes.`,
				},
				resource.Attribute{
					Name:        "db_node_role",
					Description: `The db_node_role of the db_nodes.`,
				},
				resource.Attribute{
					Name:        "max_connections",
					Description: `The max_connections of the db_nodes.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `The zone_id of the db_nodes.`,
				},
				resource.Attribute{
					Name:        "db_node_status",
					Description: `The db_node_status of the db_nodes.`,
				},
				resource.Attribute{
					Name:        "db_node_id",
					Description: `The db_node_id of the db_nodes.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The create_time of the db_nodes.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_polardb_databases",
			Category:         "Data Sources",
			ShortDescription: `Provides a collection of PolarDB endpoints according to the specified filters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "db_cluster_id",
					Description: `(Required) The polarDB cluster ID.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by database name. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `database name of the cluster.`,
				},
				resource.Attribute{
					Name:        "databases",
					Description: `A list of PolarDB cluster databases. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "character_set_name",
					Description: `The character set name of database.`,
				},
				resource.Attribute{
					Name:        "db_description",
					Description: `Database description.`,
				},
				resource.Attribute{
					Name:        "db_name",
					Description: `Database name.`,
				},
				resource.Attribute{
					Name:        "db_status",
					Description: `The status of database.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `The engine of database.`,
				},
				resource.Attribute{
					Name:        "accounts",
					Description: `A list of accounts of database. Each element contains the following attributes.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `Account name.`,
				},
				resource.Attribute{
					Name:        "account_status",
					Description: `Account status.`,
				},
				resource.Attribute{
					Name:        "privilege_status",
					Description: `The privilege status of account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `database name of the cluster.`,
				},
				resource.Attribute{
					Name:        "databases",
					Description: `A list of PolarDB cluster databases. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "character_set_name",
					Description: `The character set name of database.`,
				},
				resource.Attribute{
					Name:        "db_description",
					Description: `Database description.`,
				},
				resource.Attribute{
					Name:        "db_name",
					Description: `Database name.`,
				},
				resource.Attribute{
					Name:        "db_status",
					Description: `The status of database.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `The engine of database.`,
				},
				resource.Attribute{
					Name:        "accounts",
					Description: `A list of accounts of database. Each element contains the following attributes.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `Account name.`,
				},
				resource.Attribute{
					Name:        "account_status",
					Description: `Account status.`,
				},
				resource.Attribute{
					Name:        "privilege_status",
					Description: `The privilege status of account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_polardb_endpoints",
			Category:         "Data Sources",
			ShortDescription: `Provides a collection of PolarDB endpoints according to the specified filters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "db_cluster_id",
					Description: `(Required, ForceNew) PolarDB cluster ID.`,
				},
				resource.Attribute{
					Name:        "db_endpoint_id",
					Description: `(Optional) endpoint of the cluster. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "endpoints",
					Description: `A list of PolarDB cluster endpoints. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "db_endpoint_id",
					Description: `The endpoint ID.`,
				},
				resource.Attribute{
					Name:        "auto_add_new_nodes",
					Description: `Whether the new node is automatically added to the default cluster address.Options are ` + "`" + `Enable` + "`" + ` and ` + "`" + `Disable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "endpoint_config",
					Description: `The Endpoint configuration. ` + "`" + `ConsistLevel` + "`" + `: session consistency level, value:` + "`" + `0` + "`" + `: final consistency,` + "`" + `1` + "`" + `: session consistency;` + "`" + `LoadBalanceStrategy` + "`" + `: load balancing strategy. Based on the automatic scheduling of load, the value is: ` + "`" + `load` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "endpoint_type",
					Description: `Cluster address type.` + "`" + `Cluster` + "`" + `: the default address of the Cluster.` + "`" + `Primary` + "`" + `: Primary address.` + "`" + `Custom` + "`" + `: Custom cluster addresses.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `A list of nodes that connect to the address configuration.`,
				},
				resource.Attribute{
					Name:        "read_write_mode",
					Description: `Read-write mode:` + "`" + `ReadWrite` + "`" + `: readable and writable (automatic read-write separation).` + "`" + `ReadOnly` + "`" + `: ReadOnly.`,
				},
				resource.Attribute{
					Name:        "address_items",
					Description: `A list of endpoint addresses. Each element contains the following attributes.`,
				},
				resource.Attribute{
					Name:        "net_type",
					Description: `IP network type:` + "`" + `Public` + "`" + ` or ` + "`" + `Private` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `Connection instance string.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Intranet connection port.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `ID of the VSwitch the cluster belongs to.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The ip address of connection string.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "endpoints",
					Description: `A list of PolarDB cluster endpoints. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "db_endpoint_id",
					Description: `The endpoint ID.`,
				},
				resource.Attribute{
					Name:        "auto_add_new_nodes",
					Description: `Whether the new node is automatically added to the default cluster address.Options are ` + "`" + `Enable` + "`" + ` and ` + "`" + `Disable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "endpoint_config",
					Description: `The Endpoint configuration. ` + "`" + `ConsistLevel` + "`" + `: session consistency level, value:` + "`" + `0` + "`" + `: final consistency,` + "`" + `1` + "`" + `: session consistency;` + "`" + `LoadBalanceStrategy` + "`" + `: load balancing strategy. Based on the automatic scheduling of load, the value is: ` + "`" + `load` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "endpoint_type",
					Description: `Cluster address type.` + "`" + `Cluster` + "`" + `: the default address of the Cluster.` + "`" + `Primary` + "`" + `: Primary address.` + "`" + `Custom` + "`" + `: Custom cluster addresses.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `A list of nodes that connect to the address configuration.`,
				},
				resource.Attribute{
					Name:        "read_write_mode",
					Description: `Read-write mode:` + "`" + `ReadWrite` + "`" + `: readable and writable (automatic read-write separation).` + "`" + `ReadOnly` + "`" + `: ReadOnly.`,
				},
				resource.Attribute{
					Name:        "address_items",
					Description: `A list of endpoint addresses. Each element contains the following attributes.`,
				},
				resource.Attribute{
					Name:        "net_type",
					Description: `IP network type:` + "`" + `Public` + "`" + ` or ` + "`" + `Private` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "connection_string",
					Description: `Connection instance string.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Intranet connection port.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC the instance belongs to.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `ID of the VSwitch the cluster belongs to.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The ip address of connection string.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_polardb_node_classes",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of PolarDB node classes info.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "pay_type",
					Description: `(Required) Filter the results by charge type. Valid values: ` + "`" + `PrePaid` + "`" + ` and ` + "`" + `PostPaid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "db_type",
					Description: `(Optional) Database type. Options are ` + "`" + `MySQL` + "`" + `, ` + "`" + `PostgreSQL` + "`" + `, ` + "`" + `Oracle` + "`" + `. If db_type is set, db_version also needs to be set.`,
				},
				resource.Attribute{
					Name:        "db_version",
					Description: `(Optional) Database version required by the user. Value options can refer to the latest docs [detail info](https://www.alibabacloud.com/help/doc-detail/98169.htm) ` + "`" + `DBVersion` + "`" + `. If db_version is set, db_type also needs to be set.`,
				},
				resource.Attribute{
					Name:        "db_node_class",
					Description: `(Optional) The PolarDB node class type by the user.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `(Optional) The Region to launch the PolarDB cluster.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional) The Zone to launch the PolarDB cluster.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform apply` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "classes",
					Description: `A list of PolarDB node classes. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `The Zone to launch the PolarDB cluster.`,
				},
				resource.Attribute{
					Name:        "supported_engines",
					Description: `A list of PolarDB node classes in the zone.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `In the zone, the database type supports classes in the following available_resources.`,
				},
				resource.Attribute{
					Name:        "available_resources",
					Description: `A list of PolarDB node available classes.`,
				},
				resource.Attribute{
					Name:        "db_node_class",
					Description: `PolarDB node available class.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "classes",
					Description: `A list of PolarDB node classes. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `The Zone to launch the PolarDB cluster.`,
				},
				resource.Attribute{
					Name:        "supported_engines",
					Description: `A list of PolarDB node classes in the zone.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `In the zone, the database type supports classes in the following available_resources.`,
				},
				resource.Attribute{
					Name:        "available_resources",
					Description: `A list of PolarDB node available classes.`,
				},
				resource.Attribute{
					Name:        "db_node_class",
					Description: `PolarDB node available class.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_polardb_zones",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of availability zones for PolarDB that can be used by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "multi",
					Description: `(Optional) Indicate whether the zones can be used in a multi AZ configuration. Default to ` + "`" + `false` + "`" + `. Multi AZ is usually used to launch PolarDB instances.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "multi_zone_ids",
					Description: `A list of zone ids in which the multi zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "multi_zone_ids",
					Description: `A list of zone ids in which the multi zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_pvtz_zone_records",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Private Zone Records which owned by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "keyword",
					Description: `(Optional) Keyword for record rr and value.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) ID of the Private Zone.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, Available in 1.53.0+) A list of Private Zone Record IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of Private Zone Record IDs.`,
				},
				resource.Attribute{
					Name:        "records",
					Description: `A list of zone records. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Private Zone Record.`,
				},
				resource.Attribute{
					Name:        "resource_record",
					Description: `Resource record of the Private Zone Record.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the Private Zone Record.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value of the Private Zone Record.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `Ttl of the Private Zone Record.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Priority of the Private Zone Record.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of Private Zone Record IDs.`,
				},
				resource.Attribute{
					Name:        "records",
					Description: `A list of zone records. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Private Zone Record.`,
				},
				resource.Attribute{
					Name:        "resource_record",
					Description: `Resource record of the Private Zone Record.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the Private Zone Record.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value of the Private Zone Record.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `Ttl of the Private Zone Record.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Priority of the Private Zone Record.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_pvtz_zones",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Private Zones which owned by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "keyword",
					Description: `(Optional) keyword for zone name.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, Available 1.53.0+) A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of zone names.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Private Zone.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Remark of the Private Zone.`,
				},
				resource.Attribute{
					Name:        "record_count",
					Description: `Count of the Private Zone Record.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the Private Zone.`,
				},
				resource.Attribute{
					Name:        "is_ptr",
					Description: `Whether the Private Zone is ptr`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Time of creation of the Private Zone.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Time of update of the Private Zone.`,
				},
				resource.Attribute{
					Name:        "bind_vpcs",
					Description: `List of the VPCs is bound to the Private Zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of zone names.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Private Zone.`,
				},
				resource.Attribute{
					Name:        "remark",
					Description: `Remark of the Private Zone.`,
				},
				resource.Attribute{
					Name:        "record_count",
					Description: `Count of the Private Zone Record.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the Private Zone.`,
				},
				resource.Attribute{
					Name:        "is_ptr",
					Description: `Whether the Private Zone is ptr`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Time of creation of the Private Zone.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Time of update of the Private Zone.`,
				},
				resource.Attribute{
					Name:        "bind_vpcs",
					Description: `List of the VPCs is bound to the Private Zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ram_account_aliases",
			Category:         "Data Sources",
			ShortDescription: `Provides an alias of the Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "account_alias",
					Description: `Alias of the account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "account_alias",
					Description: `Alias of the account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ram_account_alias",
			Category:         "Data Sources",
			ShortDescription: `Provides an alias of the Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ram_groups",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of ram groups available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter the returned groups by their names.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `(Optional) Filter the results by a specific the user name.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `(Optional) Filter the results by a specific policy type. Valid items are ` + "`" + `Custom` + "`" + ` and ` + "`" + `System` + "`" + `. If you set this parameter, you must set ` + "`" + `policy_name` + "`" + ` as well.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `(Optional) Filter the results by a specific policy name. If you set this parameter without setting ` + "`" + `policy_type` + "`" + `, it will be automatically set to ` + "`" + `System` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of ram group names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the group.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments of the group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `A list of ram group names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the group.`,
				},
				resource.Attribute{
					Name:        "comments",
					Description: `Comments of the group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ram_policies",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of ram policies available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting policies by name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Filter results by a specific policy type. Valid values are ` + "`" + `Custom` + "`" + ` and ` + "`" + `System` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `(Optional) Filter results by a specific user name. Returned policies are attached to the specified user.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `(Optional) Filter results by a specific group name. Returned policies are attached to the specified group.`,
				},
				resource.Attribute{
					Name:        "role_name",
					Description: `(Optional) Filter results by a specific role name. Returned policies are attached to the specified role.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of ram group names.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `A list of policies. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the policy.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the policy.`,
				},
				resource.Attribute{
					Name:        "default_version",
					Description: `Default version of the policy.`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `Creation date of the policy.`,
				},
				resource.Attribute{
					Name:        "update_date",
					Description: `Update date of the policy.`,
				},
				resource.Attribute{
					Name:        "attachment_count",
					Description: `Attachment count of the policy.`,
				},
				resource.Attribute{
					Name:        "document",
					Description: `Policy document of the policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `A list of ram group names.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `A list of policies. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the policy.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the policy.`,
				},
				resource.Attribute{
					Name:        "default_version",
					Description: `Default version of the policy.`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `Creation date of the policy.`,
				},
				resource.Attribute{
					Name:        "update_date",
					Description: `Update date of the policy.`,
				},
				resource.Attribute{
					Name:        "attachment_count",
					Description: `Attachment count of the policy.`,
				},
				resource.Attribute{
					Name:        "document",
					Description: `Policy document of the policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ram_roles",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of ram roles available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by the role name.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `(Optional) Filter results by a specific policy type. Valid values are ` + "`" + `Custom` + "`" + ` and ` + "`" + `System` + "`" + `. If you set this parameter, you must set ` + "`" + `policy_name` + "`" + ` as well.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `(Optional) Filter results by a specific policy name. If you set this parameter without setting ` + "`" + `policy_type` + "`" + `, the later will be automatically set to ` + "`" + `System` + "`" + `. The resulting roles will be attached to the specified policy.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of ram role IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of ram role names.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `A list of roles. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the role.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the role.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `Resource descriptor of the role.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the role.`,
				},
				resource.Attribute{
					Name:        "assume_role_policy_document",
					Description: `Authorization strategy of the role. This parameter is deprecated and replaced by ` + "`" + `document` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "document",
					Description: `Authorization strategy of the role.`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `Creation date of the role.`,
				},
				resource.Attribute{
					Name:        "update_date",
					Description: `Update date of the role.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of ram role IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of ram role names.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `A list of roles. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the role.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the role.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `Resource descriptor of the role.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the role.`,
				},
				resource.Attribute{
					Name:        "assume_role_policy_document",
					Description: `Authorization strategy of the role. This parameter is deprecated and replaced by ` + "`" + `document` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "document",
					Description: `Authorization strategy of the role.`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `Creation date of the role.`,
				},
				resource.Attribute{
					Name:        "update_date",
					Description: `Update date of the role.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ram_users",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of ram users available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter resulting users by their names.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `(Optional) Filter results by a specific group name. Returned users are in the specified group.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `(Optional) Filter results by a specific policy type. Valid values are ` + "`" + `Custom` + "`" + ` and ` + "`" + `System` + "`" + `. If you set this parameter, you must set ` + "`" + `policy_name` + "`" + ` as well.`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `(Optional) Filter results by a specific policy name. If you set this parameter without setting ` + "`" + `policy_type` + "`" + `, the later will be automatically set to ` + "`" + `System` + "`" + `. Returned users are attached to the specified policy.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of ram user IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of ram user names.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `A list of users. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The original id is user name, but it is user id in 1.37.0+.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the user.`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `Creation date of the user.`,
				},
				resource.Attribute{
					Name:        "last_login_date",
					Description: `Last login date of the user. Removed from version 1.79.0.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of ram user IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of ram user names.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `A list of users. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The original id is user name, but it is user id in 1.37.0+.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the user.`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `Creation date of the user.`,
				},
				resource.Attribute{
					Name:        "last_login_date",
					Description: `Last login date of the user. Removed from version 1.79.0.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_regions",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of regions that can be used by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the region to select, such as ` + "`" + `eu-central-1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "current",
					Description: `(Optional) Set to true to match only the region configured in the provider.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ~>`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of region IDs.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `A list of regions. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the region.`,
				},
				resource.Attribute{
					Name:        "local_name",
					Description: `Name of the region in the local language.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of region IDs.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `A list of regions. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the region.`,
				},
				resource.Attribute{
					Name:        "local_name",
					Description: `Name of the region in the local language.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_resource_manager_accounts",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Resource Manager Accounts to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of account IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of account IDs.`,
				},
				resource.Attribute{
					Name:        "accounts",
					Description: `A list of accounts. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `The ID of the folder.`,
				},
				resource.Attribute{
					Name:        "join_method",
					Description: `The way in which the member account joined the resource directory.`,
				},
				resource.Attribute{
					Name:        "join_time",
					Description: `The time when the member account joined the resource directory.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `The time when the member account was modified.`,
				},
				resource.Attribute{
					Name:        "resource_directory_id",
					Description: `The ID of the resource directory.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the member account.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the member account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of account IDs.`,
				},
				resource.Attribute{
					Name:        "accounts",
					Description: `A list of accounts. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `The ID of the folder.`,
				},
				resource.Attribute{
					Name:        "join_method",
					Description: `The way in which the member account joined the resource directory.`,
				},
				resource.Attribute{
					Name:        "join_time",
					Description: `The time when the member account joined the resource directory.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `The time when the member account was modified.`,
				},
				resource.Attribute{
					Name:        "resource_directory_id",
					Description: `The ID of the resource directory.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the member account.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the member account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_resource_manager_folders",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of resource manager folders to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of resource manager folders IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by folder name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of folder IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of folder names.`,
				},
				resource.Attribute{
					Name:        "folders",
					Description: `A list of folders. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the folder.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of folder IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of folder names.`,
				},
				resource.Attribute{
					Name:        "folders",
					Description: `A list of folders. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the folder.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_resource_manager_handshakes",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Resource Manager Handshakes to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Resource Manager Handshake IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of Resource Manager Handshake IDs.`,
				},
				resource.Attribute{
					Name:        "handshakes",
					Description: `A list of Resource Manager Handshakes. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The time when the invitation expires.`,
				},
				resource.Attribute{
					Name:        "master_account_id",
					Description: `The ID of the master account of the resource directory.`,
				},
				resource.Attribute{
					Name:        "master_account_name",
					Description: `The name of the master account of the resource directory.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `The time when the invitation was modified.`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `The invitation note.`,
				},
				resource.Attribute{
					Name:        "resource_directory_id",
					Description: `The ID of the resource directory.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the invitation.`,
				},
				resource.Attribute{
					Name:        "target_entity",
					Description: `The ID or logon email address of the invited account.`,
				},
				resource.Attribute{
					Name:        "target_type",
					Description: `The type of the invited account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of Resource Manager Handshake IDs.`,
				},
				resource.Attribute{
					Name:        "handshakes",
					Description: `A list of Resource Manager Handshakes. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `The time when the invitation expires.`,
				},
				resource.Attribute{
					Name:        "master_account_id",
					Description: `The ID of the master account of the resource directory.`,
				},
				resource.Attribute{
					Name:        "master_account_name",
					Description: `The name of the master account of the resource directory.`,
				},
				resource.Attribute{
					Name:        "modify_time",
					Description: `The time when the invitation was modified.`,
				},
				resource.Attribute{
					Name:        "note",
					Description: `The invitation note.`,
				},
				resource.Attribute{
					Name:        "resource_directory_id",
					Description: `The ID of the resource directory.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the invitation.`,
				},
				resource.Attribute{
					Name:        "target_entity",
					Description: `The ID or logon email address of the invited account.`,
				},
				resource.Attribute{
					Name:        "target_type",
					Description: `The type of the invited account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_resource_manager_policies",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Resource Manager Policies to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Resource Manager Policy IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by policy name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `(Optional) The type of the policy. If you do not specify this parameter, the system lists all types of policies. Valid values: ` + "`" + `Custom` + "`" + ` and ` + "`" + `System` + "`" + `. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of policy IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of policy names.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `A list of policies. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the policy.`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `The time when the policy was created.`,
				},
				resource.Attribute{
					Name:        "update_date",
					Description: `The time when the policy was updated.`,
				},
				resource.Attribute{
					Name:        "default_version",
					Description: `The default version of the policy.`,
				},
				resource.Attribute{
					Name:        "attachment_count",
					Description: `The number of times the policy is referenced.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of policy IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of policy names.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `A list of policies. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the policy.`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `The time when the policy was created.`,
				},
				resource.Attribute{
					Name:        "update_date",
					Description: `The time when the policy was updated.`,
				},
				resource.Attribute{
					Name:        "default_version",
					Description: `The default version of the policy.`,
				},
				resource.Attribute{
					Name:        "attachment_count",
					Description: `The number of times the policy is referenced.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_resource_manager_policy_versions",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Resource Manager Policy Versions to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of policy version IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "policy_name",
					Description: `(Required) The name of the policy.`,
				},
				resource.Attribute{
					Name:        "policy_type",
					Description: `(Required) The type of the policy. Valid values:` + "`" + `Custom` + "`" + ` and ` + "`" + `System` + "`" + `. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of policy version IDs.`,
				},
				resource.Attribute{
					Name:        "versions",
					Description: `A list of policy versions. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource, the value is ` + "`" + `<policy_name>` + "`" + `:` + "`" + `<version_id>` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of policy version IDs.`,
				},
				resource.Attribute{
					Name:        "versions",
					Description: `A list of policy versions. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource, the value is ` + "`" + `<policy_name>` + "`" + `:` + "`" + `<version_id>` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_resource_manager_resource_directories",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Resource Manager Resource Directories to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "directories",
					Description: `A list of resource directories. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of resource directory.`,
				},
				resource.Attribute{
					Name:        "resource_directory_id",
					Description: `The ID of the resource directory.`,
				},
				resource.Attribute{
					Name:        "root_folder_id",
					Description: `The ID of the root folder.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "directories",
					Description: `A list of resource directories. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of resource directory.`,
				},
				resource.Attribute{
					Name:        "resource_directory_id",
					Description: `The ID of the resource directory.`,
				},
				resource.Attribute{
					Name:        "root_folder_id",
					Description: `The ID of the root folder.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_resource_manager_resource_groups",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of resource groups to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of resource group IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by resource group name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the resource group. Possible values:` + "`" + `Creating` + "`" + `,` + "`" + `Deleted` + "`" + `,` + "`" + `OK` + "`" + ` and ` + "`" + `PendingDelete` + "`" + `. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of resource group IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of resource group names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of resource groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The unique identifier of the resource group.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the resource group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The ID of the Alibaba Cloud account to which the resource group belongs.`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `The time when the resource group was created.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the resource group. Possible values:` + "`" + `Creating` + "`" + `,` + "`" + `Deleted` + "`" + `,` + "`" + `OK` + "`" + ` and ` + "`" + `PendingDelete` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of resource group IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of resource group names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of resource groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The unique identifier of the resource group.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the resource group.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `The ID of the Alibaba Cloud account to which the resource group belongs.`,
				},
				resource.Attribute{
					Name:        "create_date",
					Description: `The time when the resource group was created.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the resource group. Possible values:` + "`" + `Creating` + "`" + `,` + "`" + `Deleted` + "`" + `,` + "`" + `OK` + "`" + ` and ` + "`" + `PendingDelete` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_resource_manager_roles",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Resource Manager Roles to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Resource Manager Role IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by role name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of role IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of role names.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `A list of roles. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the role.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of role IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of role names.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `A list of roles. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the role.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_route_entries",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Route Entries owned by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Required, ForceNew) The ID of the router table to which the route entry belongs.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) The instance ID of the next hop.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of the route entry.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Optional) The destination CIDR block of the route entry.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "entries",
					Description: `A list of Route Entries. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the route entry.`,
				},
				resource.Attribute{
					Name:        "next_hop_type",
					Description: `The type of the next hop.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the route entry.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The instance ID of the next hop.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the router table to which the route entry belongs.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The destination CIDR block of the route entry.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "entries",
					Description: `A list of Route Entries. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the route entry.`,
				},
				resource.Attribute{
					Name:        "next_hop_type",
					Description: `The type of the next hop.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the route entry.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The instance ID of the next hop.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `The ID of the router table to which the route entry belongs.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `The destination CIDR block of the route entry.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_route_tables",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Route Tables owned by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Route Tables IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter route tables by name.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) Vpc id of the route table.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.55.3+) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew, Available in 1.60.0+) The Id of resource group which route tables belongs. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Route Tables IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Route Tables names.`,
				},
				resource.Attribute{
					Name:        "tables",
					Description: `A list of Route Tables. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Route Table.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `Router Id of the route table.`,
				},
				resource.Attribute{
					Name:        "route_table_type",
					Description: `The type of route table.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the route table.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the route table instance.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Time of creation.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The Id of resource group which route tables belongs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Route Tables IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Route Tables names.`,
				},
				resource.Attribute{
					Name:        "tables",
					Description: `A list of Route Tables. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Route Table.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `Router Id of the route table.`,
				},
				resource.Attribute{
					Name:        "route_table_type",
					Description: `The type of route table.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the route table.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the route table instance.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Time of creation.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The Id of resource group which route tables belongs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_router_interfaces",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of router interfaces to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string used to filter by router interface name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Expected status. Valid values are ` + "`" + `Active` + "`" + `, ` + "`" + `Inactive` + "`" + ` and ` + "`" + `Idle` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "specification",
					Description: `(Optional) Specification of the link, such as ` + "`" + `Small.1` + "`" + ` (10Mb), ` + "`" + `Middle.1` + "`" + ` (100Mb), ` + "`" + `Large.2` + "`" + ` (2Gb), ...etc.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `(Optional) ID of the VRouter located in the local region.`,
				},
				resource.Attribute{
					Name:        "router_type",
					Description: `(Optional) Router type in the local region. Valid values are ` + "`" + `VRouter` + "`" + ` and ` + "`" + `VBR` + "`" + ` (physical connection).`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) Role of the router interface. Valid values are ` + "`" + `InitiatingSide` + "`" + ` (connection initiator) and ` + "`" + `AcceptingSide` + "`" + ` (connection receiver). The value of this parameter must be ` + "`" + `InitiatingSide` + "`" + ` if the ` + "`" + `router_type` + "`" + ` is set to ` + "`" + `VBR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "opposite_interface_id",
					Description: `(Optional) ID of the peer router interface.`,
				},
				resource.Attribute{
					Name:        "opposite_interface_owner_id",
					Description: `(Optional) Account ID of the owner of the peer router interface.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, Available in 1.44.0+) A list of router interface IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of router interface IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of router interface names.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `A list of router interfaces. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Router interface ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Router interface status. Possible values: ` + "`" + `Active` + "`" + `, ` + "`" + `Inactive` + "`" + ` and ` + "`" + `Idle` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Router interface name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Router interface description.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Router interface role. Possible values: ` + "`" + `InitiatingSide` + "`" + ` and ` + "`" + `AcceptingSide` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "specification",
					Description: `Router interface specification. Possible values: ` + "`" + `Small.1` + "`" + `, ` + "`" + `Middle.1` + "`" + `, ` + "`" + `Large.2` + "`" + `, ...etc.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `ID of the VRouter located in the local region.`,
				},
				resource.Attribute{
					Name:        "router_type",
					Description: `Router type in the local region. Possible values: ` + "`" + `VRouter` + "`" + ` and ` + "`" + `VBR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC that owns the router in the local region.`,
				},
				resource.Attribute{
					Name:        "access_point_id",
					Description: `ID of the access point used by the VBR.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Router interface creation time.`,
				},
				resource.Attribute{
					Name:        "opposite_region_id",
					Description: `Peer router region ID.`,
				},
				resource.Attribute{
					Name:        "opposite_interface_id",
					Description: `Peer router interface ID.`,
				},
				resource.Attribute{
					Name:        "opposite_router_id",
					Description: `Peer router ID.`,
				},
				resource.Attribute{
					Name:        "opposite_router_type",
					Description: `Router type in the peer region. Possible values: ` + "`" + `VRouter` + "`" + ` and ` + "`" + `VBR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "opposite_interface_owner_id",
					Description: `Account ID of the owner of the peer router interface.`,
				},
				resource.Attribute{
					Name:        "health_check_source_ip",
					Description: `Source IP address used to perform health check on the physical connection.`,
				},
				resource.Attribute{
					Name:        "health_check_target_ip",
					Description: `Destination IP address used to perform health check on the physical connection.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of router interface IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of router interface names.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `A list of router interfaces. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Router interface ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Router interface status. Possible values: ` + "`" + `Active` + "`" + `, ` + "`" + `Inactive` + "`" + ` and ` + "`" + `Idle` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Router interface name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Router interface description.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Router interface role. Possible values: ` + "`" + `InitiatingSide` + "`" + ` and ` + "`" + `AcceptingSide` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "specification",
					Description: `Router interface specification. Possible values: ` + "`" + `Small.1` + "`" + `, ` + "`" + `Middle.1` + "`" + `, ` + "`" + `Large.2` + "`" + `, ...etc.`,
				},
				resource.Attribute{
					Name:        "router_id",
					Description: `ID of the VRouter located in the local region.`,
				},
				resource.Attribute{
					Name:        "router_type",
					Description: `Router type in the local region. Possible values: ` + "`" + `VRouter` + "`" + ` and ` + "`" + `VBR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC that owns the router in the local region.`,
				},
				resource.Attribute{
					Name:        "access_point_id",
					Description: `ID of the access point used by the VBR.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Router interface creation time.`,
				},
				resource.Attribute{
					Name:        "opposite_region_id",
					Description: `Peer router region ID.`,
				},
				resource.Attribute{
					Name:        "opposite_interface_id",
					Description: `Peer router interface ID.`,
				},
				resource.Attribute{
					Name:        "opposite_router_id",
					Description: `Peer router ID.`,
				},
				resource.Attribute{
					Name:        "opposite_router_type",
					Description: `Router type in the peer region. Possible values: ` + "`" + `VRouter` + "`" + ` and ` + "`" + `VBR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "opposite_interface_owner_id",
					Description: `Account ID of the owner of the peer router interface.`,
				},
				resource.Attribute{
					Name:        "health_check_source_ip",
					Description: `Source IP address used to perform health check on the physical connection.`,
				},
				resource.Attribute{
					Name:        "health_check_target_ip",
					Description: `Destination IP address used to perform health check on the physical connection.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_sag_acls",
			Category:         "Data Sources",
			ShortDescription: `Provides the access control list (ACL) function in the form of whitelists and blacklists for different SAG instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Sag Acl IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter Sag Acl instances by name. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of Sag Acl IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Sag Acls names.`,
				},
				resource.Attribute{
					Name:        "acls",
					Description: `A list of Sag Acls. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the ACL. For example "acl-xxx".`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Acl.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of Sag Acl IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Sag Acls names.`,
				},
				resource.Attribute{
					Name:        "acls",
					Description: `A list of Sag Acls. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the ACL. For example "acl-xxx".`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Acl.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_security_group_rules",
			Category:         "Data Sources",
			ShortDescription: `Provides a collection of Security Group Rules available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required) The ID of the security group that owns the rules.`,
				},
				resource.Attribute{
					Name:        "nic_type",
					Description: `(Optional) Refers to the network type. Can be either ` + "`" + `internet` + "`" + ` or ` + "`" + `intranet` + "`" + `. The default value is ` + "`" + `internet` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Optional) Authorization direction. Valid values are: ` + "`" + `ingress` + "`" + ` or ` + "`" + `egress` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `(Optional) The IP protocol. Valid values are: ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, ` + "`" + `gre` + "`" + ` and ` + "`" + `all` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Optional) Authorization policy. Can be either ` + "`" + `accept` + "`" + ` or ` + "`" + `drop` + "`" + `. The default value is ` + "`" + `accept` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `A list of rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `The name of the security group that owns the rules.`,
				},
				resource.Attribute{
					Name:        "group_desc",
					Description: `The description of the security group that owns the rules.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `A list of security group rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The protocol. Can be ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, ` + "`" + `gre` + "`" + ` or ` + "`" + `all` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port_range",
					Description: `The range of port numbers.`,
				},
				resource.Attribute{
					Name:        "source_cidr_ip",
					Description: `Source IP address segment for ingress authorization.`,
				},
				resource.Attribute{
					Name:        "source_security_group_id",
					Description: `Source security group ID for ingress authorization.`,
				},
				resource.Attribute{
					Name:        "source_group_owner_account",
					Description: `Alibaba Cloud account of the source security group.`,
				},
				resource.Attribute{
					Name:        "dest_cidr_ip",
					Description: `Target IP address segment for egress authorization.`,
				},
				resource.Attribute{
					Name:        "dest_security_group_id",
					Description: `Target security group id for ingress authorization.`,
				},
				resource.Attribute{
					Name:        "dest_group_owner_account",
					Description: `Alibaba Cloud account of the target security group.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `Authorization policy. Can be either ` + "`" + `accept` + "`" + ` or ` + "`" + `drop` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "nic_type",
					Description: `Network type, ` + "`" + `internet` + "`" + ` or ` + "`" + `intranet` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Rule priority.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `Authorization direction, ` + "`" + `ingress` + "`" + ` or ` + "`" + `egress` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "rules",
					Description: `A list of rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `The name of the security group that owns the rules.`,
				},
				resource.Attribute{
					Name:        "group_desc",
					Description: `The description of the security group that owns the rules.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `A list of security group rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `The protocol. Can be ` + "`" + `tcp` + "`" + `, ` + "`" + `udp` + "`" + `, ` + "`" + `icmp` + "`" + `, ` + "`" + `gre` + "`" + ` or ` + "`" + `all` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port_range",
					Description: `The range of port numbers.`,
				},
				resource.Attribute{
					Name:        "source_cidr_ip",
					Description: `Source IP address segment for ingress authorization.`,
				},
				resource.Attribute{
					Name:        "source_security_group_id",
					Description: `Source security group ID for ingress authorization.`,
				},
				resource.Attribute{
					Name:        "source_group_owner_account",
					Description: `Alibaba Cloud account of the source security group.`,
				},
				resource.Attribute{
					Name:        "dest_cidr_ip",
					Description: `Target IP address segment for egress authorization.`,
				},
				resource.Attribute{
					Name:        "dest_security_group_id",
					Description: `Target security group id for ingress authorization.`,
				},
				resource.Attribute{
					Name:        "dest_group_owner_account",
					Description: `Alibaba Cloud account of the target security group.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `Authorization policy. Can be either ` + "`" + `accept` + "`" + ` or ` + "`" + `drop` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "nic_type",
					Description: `Network type, ` + "`" + `internet` + "`" + ` or ` + "`" + `intranet` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Rule priority.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `Authorization direction, ` + "`" + `ingress` + "`" + ` or ` + "`" + `egress` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_security_groups",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Security Groups available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, Available 1.52.0+) A list of Security Group IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter the resulting security groups by their names.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) Used to retrieve security groups that belong to the specified VPC ID.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew, Available in 1.58.0+) The Id of resource group which the security_group belongs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A map of tags assigned to the ECS instances. It must be in the format: ` + "`" + `` + "`" + `` + "`" + ` data "alicloud_security_groups" "taggedSecurityGroups" { tags = { tagKey1 = "tagValue1", tagKey2 = "tagValue2" } } ` + "`" + `` + "`" + `` + "`" + ` ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of Security Group IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Security Group names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of Security Groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the security group.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of the VPC that owns the security group.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The Id of resource group which the security_group belongs.`,
				},
				resource.Attribute{
					Name:        "security_group_type",
					Description: `The type of the security group.`,
				},
				resource.Attribute{
					Name:        "inner_access",
					Description: `Whether to allow inner network access.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Creation time of the security group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the ECS instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of Security Group IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of Security Group names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of Security Groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the security group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the security group.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of the VPC that owns the security group.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The Id of resource group which the security_group belongs.`,
				},
				resource.Attribute{
					Name:        "security_group_type",
					Description: `The type of the security group.`,
				},
				resource.Attribute{
					Name:        "inner_access",
					Description: `Whether to allow inner network access.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Creation time of the security group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the ECS instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_slb_acls",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of server load balancer acls (access control lists) to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of acls IDs to filter results.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by acl name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew, Available in 1.60.0+) The Id of resource group which acl belongs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.66.0+) A mapping of tags to assign to the resource. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of SLB acls IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of SLB acls names.`,
				},
				resource.Attribute{
					Name:        "acls",
					Description: `A list of SLB acls. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Acl ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Acl name.`,
				},
				resource.Attribute{
					Name:        "entry_list",
					Description: `A list of entry (IP addresses or CIDR blocks). Each entry contains two sub-fields as ` + "`" + `Entry Block` + "`" + ` follows.`,
				},
				resource.Attribute{
					Name:        "related_listeners",
					Description: `A list of listener are attached by the acl. Each listener contains four sub-fields as ` + "`" + `Listener Block` + "`" + ` follows.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `Resource group ID. ## Entry Block The entry mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `the comment of the entry. ## Listener Block The Listener mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `the id of load balancer instance, the listener belongs to.`,
				},
				resource.Attribute{
					Name:        "frontend_port",
					Description: `the listener port.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of SLB acls IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of SLB acls names.`,
				},
				resource.Attribute{
					Name:        "acls",
					Description: `A list of SLB acls. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Acl ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Acl name.`,
				},
				resource.Attribute{
					Name:        "entry_list",
					Description: `A list of entry (IP addresses or CIDR blocks). Each entry contains two sub-fields as ` + "`" + `Entry Block` + "`" + ` follows.`,
				},
				resource.Attribute{
					Name:        "related_listeners",
					Description: `A list of listener are attached by the acl. Each listener contains four sub-fields as ` + "`" + `Listener Block` + "`" + ` follows.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `Resource group ID. ## Entry Block The entry mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `the comment of the entry. ## Listener Block The Listener mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `the id of load balancer instance, the listener belongs to.`,
				},
				resource.Attribute{
					Name:        "frontend_port",
					Description: `the listener port.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_slb_attachments",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of server load balancer attachments to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `ID of the SLB with attachments.`,
				},
				resource.Attribute{
					Name:        "instance_ids",
					Description: `(Optional) List of attached ECS instance IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "slb_attachments",
					Description: `A list of SLB attachments. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the attached ECS instance.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Weight associated to the ECS instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "slb_attachments",
					Description: `A list of SLB attachments. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the attached ECS instance.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Weight associated to the ECS instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_slb_backend_servers",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of server load balancer backend servers to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `ID of the SLB with attachments.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) List of attached ECS instance IDs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "backend_servers",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `backend server ID.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Weight associated to the ECS instance.`,
				},
				resource.Attribute{
					Name:        "server_type",
					Description: `Type of the backend server.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "backend_servers",
					Description: ``,
				},
				resource.Attribute{
					Name:        "id",
					Description: `backend server ID.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Weight associated to the ECS instance.`,
				},
				resource.Attribute{
					Name:        "server_type",
					Description: `Type of the backend server.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_slb_ca_certificates",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of slb CA certificates.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of ca certificates IDs to filter results.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by ca certificate name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew, Available in 1.60.0+) The Id of resource group which ca certificates belongs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.66.0+) A mapping of tags to assign to the resource. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of SLB ca certificates IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of SLB ca certificates names.`,
				},
				resource.Attribute{
					Name:        "certificates",
					Description: `A list of SLB ca certificates. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `CA certificate ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `CA certificate name.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `CA certificate fingerprint.`,
				},
				resource.Attribute{
					Name:        "common_name",
					Description: `CA certificate common name.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `CA certificate expired time.`,
				},
				resource.Attribute{
					Name:        "expired_timestamp",
					Description: `CA certificate expired timestamp.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `CA certificate created time.`,
				},
				resource.Attribute{
					Name:        "created_timestamp",
					Description: `CA certificate created timestamp.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The resource group Id of CA certificate.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The region Id of CA certificate.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Available in v1.66.0+) A mapping of tags to assign to the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of SLB ca certificates IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of SLB ca certificates names.`,
				},
				resource.Attribute{
					Name:        "certificates",
					Description: `A list of SLB ca certificates. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `CA certificate ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `CA certificate name.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `CA certificate fingerprint.`,
				},
				resource.Attribute{
					Name:        "common_name",
					Description: `CA certificate common name.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `CA certificate expired time.`,
				},
				resource.Attribute{
					Name:        "expired_timestamp",
					Description: `CA certificate expired timestamp.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `CA certificate created time.`,
				},
				resource.Attribute{
					Name:        "created_timestamp",
					Description: `CA certificate created timestamp.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The resource group Id of CA certificate.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The region Id of CA certificate.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Available in v1.66.0+) A mapping of tags to assign to the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_slb_domain_extensions",
			Category:         "Data Sources",
			ShortDescription: `Provides a Load Banlancer domain extension Resource and add it to one Listener.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) IDs of the SLB domain extensions.`,
				},
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `(Required) The ID of the SLB instance.`,
				},
				resource.Attribute{
					Name:        "frontend_port",
					Description: `(Required) The frontend port used by the HTTPS listener of the SLB instance. Valid values: 1–65535. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "extensions",
					Description: `A list of SLB domain extension. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the domain extension.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `The domain name.`,
				},
				resource.Attribute{
					Name:        "server_certificate_id",
					Description: `The ID of the certificate used by the domain name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "extensions",
					Description: `A list of SLB domain extension. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the domain extension.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `The domain name.`,
				},
				resource.Attribute{
					Name:        "server_certificate_id",
					Description: `The ID of the certificate used by the domain name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_slb_listeners",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of server load balancer listeners to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `(Required) ID of the SLB with listeners.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Filter listeners by the specified protocol. Valid values: ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `, ` + "`" + `tcp` + "`" + ` and ` + "`" + `udp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "frontend_port",
					Description: `(Optional) Filter listeners by the specified frontend port.`,
				},
				resource.Attribute{
					Name:        "description_regex",
					Description: `(Optional, Available in 1.69.0+) A regex string to filter results by SLB listener description.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "slb_listeners",
					Description: `A list of SLB listeners. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "frontend_port",
					Description: `Frontend port used to receive incoming traffic and distribute it to the backend servers.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `Port opened on the backend server to receive requests.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Listener protocol. Possible values: ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `, ` + "`" + `tcp` + "`" + ` and ` + "`" + `udp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Listener status.`,
				},
				resource.Attribute{
					Name:        "security_status",
					Description: `Security status. Only available when the protocol is ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `Peak bandwidth. If the value is set to -1, the listener is not limited by bandwidth.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `Algorithm used to distribute traffic. Possible values: ` + "`" + `wrr` + "`" + ` (weighted round robin), ` + "`" + `wlc` + "`" + ` (weighted least connection) and ` + "`" + `rr` + "`" + ` (round robin).`,
				},
				resource.Attribute{
					Name:        "server_group_id",
					Description: `ID of the linked VServer group.`,
				},
				resource.Attribute{
					Name:        "master_slave_server_group_id",
					Description: `ID of the active/standby server group.`,
				},
				resource.Attribute{
					Name:        "persistence_timeout",
					Description: `Timeout value of the TCP connection in seconds. If the value is 0, the session persistence function is disabled. Only available when the protocol is ` + "`" + `tcp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "established_timeout",
					Description: `Connection timeout in seconds for the Layer 4 TCP listener. Only available when the protocol is ` + "`" + `tcp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sticky_session",
					Description: `Indicate whether session persistence is enabled or not. If enabled, all session requests from the same client are sent to the same backend server. Possible values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sticky_session_type",
					Description: `Method used to handle the cookie. Possible values are ` + "`" + `insert` + "`" + ` (cookie added to the response) and ` + "`" + `server` + "`" + ` (cookie set by the backend server). Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + ` and sticky_session is ` + "`" + `on` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cookie_timeout",
					Description: `Cookie timeout in seconds. Only available when the sticky_session_type is ` + "`" + `insert` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cookie",
					Description: `Cookie configured by the backend server. Only available when the sticky_session_type is ` + "`" + `server` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Indicate whether health check is enabled of not. Possible values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_type",
					Description: `Health check method. Possible values are ` + "`" + `tcp` + "`" + ` and ` + "`" + `http` + "`" + `. Only available when the protocol is ` + "`" + `tcp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_domain",
					Description: `Domain name used for health check. The SLB sends HTTP head requests to the backend server, the domain is useful when the backend server verifies the host field in the requests. Only available when the protocol is ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + ` or ` + "`" + `tcp` + "`" + ` (in this case health_check_type must be ` + "`" + `http` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "health_check_uri",
					Description: `URI used for health check. Only available when the protocol is ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + ` or ` + "`" + `tcp` + "`" + ` (in this case health_check_type must be ` + "`" + `http` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "health_check_connect_port",
					Description: `Port used for health check.`,
				},
				resource.Attribute{
					Name:        "health_check_connect_timeout",
					Description: `Amount of time in seconds to wait for the response for a health check.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `Number of consecutive successes of health check performed on the same ECS instance (from failure to success).`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `Number of consecutive failures of health check performed on the same ECS instance (from success to failure).`,
				},
				resource.Attribute{
					Name:        "health_check_timeout",
					Description: `Amount of time in seconds to wait for the response from a health check. If an ECS instance sends no response within the specified timeout period, the health check fails. Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_interval",
					Description: `Time interval between two consecutive health checks.`,
				},
				resource.Attribute{
					Name:        "health_check_http_code",
					Description: `HTTP status codes indicating that the health check is normal. It can contain several comma-separated values such as "http_2xx,http_3xx". Only available when the protocol is ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + ` or ` + "`" + `tcp` + "`" + ` (in this case health_check_type must be ` + "`" + `http` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "gzip",
					Description: `Indicate whether Gzip compression is enabled or not. Possible values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ssl_certificate_id",
					Description: `ID of the server certificate. Only available when the protocol is ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ca_certificate_id",
					Description: `ID of the CA certificate (only required when two-way authentication is used). Only available when the protocol is ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "x_forwarded_for",
					Description: `Indicate whether the HTTP header field "X-Forwarded-For" is added or not; it allows the backend server to know about the user's IP address. Possible values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "x_forwarded_for_slb_ip",
					Description: `Indicate whether the HTTP header field "X-Forwarded-For_SLBIP" is added or not; it allows the backend server to know about the SLB IP address. Possible values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "x_forwarded_for_slb_id",
					Description: `Indicate whether the HTTP header field "X-Forwarded-For_SLBID" is added or not; it allows the backend server to know about the SLB ID. Possible values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "x_forwarded_for_slb_proto",
					Description: `Indicate whether the HTTP header field "X-Forwarded-For_proto" is added or not; it allows the backend server to know about the user's protocol. Possible values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `Timeout of http or https listener established connection idle timeout. Valid value range: [1-60] in seconds. Default to 15.`,
				},
				resource.Attribute{
					Name:        "request_timeout",
					Description: `Timeout of http or https listener request (which does not get response from backend) timeout. Valid value range: [1-180] in seconds. Default to 60.`,
				},
				resource.Attribute{
					Name:        "enable_http2",
					Description: `Whether to enable https listener support http2 or not. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Default to ` + "`" + `on` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tls_cipher_policy",
					Description: `Https listener TLS cipher policy. Valid values are ` + "`" + `tls_cipher_policy_1_0` + "`" + `, ` + "`" + `tls_cipher_policy_1_1` + "`" + `, ` + "`" + `tls_cipher_policy_1_2` + "`" + `, ` + "`" + `tls_cipher_policy_1_2_strict` + "`" + `. Default to ` + "`" + `tls_cipher_policy_1_0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of slb listener.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "slb_listeners",
					Description: `A list of SLB listeners. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "frontend_port",
					Description: `Frontend port used to receive incoming traffic and distribute it to the backend servers.`,
				},
				resource.Attribute{
					Name:        "backend_port",
					Description: `Port opened on the backend server to receive requests.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Listener protocol. Possible values: ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `, ` + "`" + `tcp` + "`" + ` and ` + "`" + `udp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Listener status.`,
				},
				resource.Attribute{
					Name:        "security_status",
					Description: `Security status. Only available when the protocol is ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `Peak bandwidth. If the value is set to -1, the listener is not limited by bandwidth.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `Algorithm used to distribute traffic. Possible values: ` + "`" + `wrr` + "`" + ` (weighted round robin), ` + "`" + `wlc` + "`" + ` (weighted least connection) and ` + "`" + `rr` + "`" + ` (round robin).`,
				},
				resource.Attribute{
					Name:        "server_group_id",
					Description: `ID of the linked VServer group.`,
				},
				resource.Attribute{
					Name:        "master_slave_server_group_id",
					Description: `ID of the active/standby server group.`,
				},
				resource.Attribute{
					Name:        "persistence_timeout",
					Description: `Timeout value of the TCP connection in seconds. If the value is 0, the session persistence function is disabled. Only available when the protocol is ` + "`" + `tcp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "established_timeout",
					Description: `Connection timeout in seconds for the Layer 4 TCP listener. Only available when the protocol is ` + "`" + `tcp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sticky_session",
					Description: `Indicate whether session persistence is enabled or not. If enabled, all session requests from the same client are sent to the same backend server. Possible values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sticky_session_type",
					Description: `Method used to handle the cookie. Possible values are ` + "`" + `insert` + "`" + ` (cookie added to the response) and ` + "`" + `server` + "`" + ` (cookie set by the backend server). Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + ` and sticky_session is ` + "`" + `on` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cookie_timeout",
					Description: `Cookie timeout in seconds. Only available when the sticky_session_type is ` + "`" + `insert` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cookie",
					Description: `Cookie configured by the backend server. Only available when the sticky_session_type is ` + "`" + `server` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Indicate whether health check is enabled of not. Possible values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_type",
					Description: `Health check method. Possible values are ` + "`" + `tcp` + "`" + ` and ` + "`" + `http` + "`" + `. Only available when the protocol is ` + "`" + `tcp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_domain",
					Description: `Domain name used for health check. The SLB sends HTTP head requests to the backend server, the domain is useful when the backend server verifies the host field in the requests. Only available when the protocol is ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + ` or ` + "`" + `tcp` + "`" + ` (in this case health_check_type must be ` + "`" + `http` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "health_check_uri",
					Description: `URI used for health check. Only available when the protocol is ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + ` or ` + "`" + `tcp` + "`" + ` (in this case health_check_type must be ` + "`" + `http` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "health_check_connect_port",
					Description: `Port used for health check.`,
				},
				resource.Attribute{
					Name:        "health_check_connect_timeout",
					Description: `Amount of time in seconds to wait for the response for a health check.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `Number of consecutive successes of health check performed on the same ECS instance (from failure to success).`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `Number of consecutive failures of health check performed on the same ECS instance (from success to failure).`,
				},
				resource.Attribute{
					Name:        "health_check_timeout",
					Description: `Amount of time in seconds to wait for the response from a health check. If an ECS instance sends no response within the specified timeout period, the health check fails. Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_interval",
					Description: `Time interval between two consecutive health checks.`,
				},
				resource.Attribute{
					Name:        "health_check_http_code",
					Description: `HTTP status codes indicating that the health check is normal. It can contain several comma-separated values such as "http_2xx,http_3xx". Only available when the protocol is ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + ` or ` + "`" + `tcp` + "`" + ` (in this case health_check_type must be ` + "`" + `http` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "gzip",
					Description: `Indicate whether Gzip compression is enabled or not. Possible values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ssl_certificate_id",
					Description: `ID of the server certificate. Only available when the protocol is ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ca_certificate_id",
					Description: `ID of the CA certificate (only required when two-way authentication is used). Only available when the protocol is ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "x_forwarded_for",
					Description: `Indicate whether the HTTP header field "X-Forwarded-For" is added or not; it allows the backend server to know about the user's IP address. Possible values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "x_forwarded_for_slb_ip",
					Description: `Indicate whether the HTTP header field "X-Forwarded-For_SLBIP" is added or not; it allows the backend server to know about the SLB IP address. Possible values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "x_forwarded_for_slb_id",
					Description: `Indicate whether the HTTP header field "X-Forwarded-For_SLBID" is added or not; it allows the backend server to know about the SLB ID. Possible values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "x_forwarded_for_slb_proto",
					Description: `Indicate whether the HTTP header field "X-Forwarded-For_proto" is added or not; it allows the backend server to know about the user's protocol. Possible values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Only available when the protocol is ` + "`" + `http` + "`" + ` or ` + "`" + `https` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `Timeout of http or https listener established connection idle timeout. Valid value range: [1-60] in seconds. Default to 15.`,
				},
				resource.Attribute{
					Name:        "request_timeout",
					Description: `Timeout of http or https listener request (which does not get response from backend) timeout. Valid value range: [1-180] in seconds. Default to 60.`,
				},
				resource.Attribute{
					Name:        "enable_http2",
					Description: `Whether to enable https listener support http2 or not. Valid values are ` + "`" + `on` + "`" + ` and ` + "`" + `off` + "`" + `. Default to ` + "`" + `on` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tls_cipher_policy",
					Description: `Https listener TLS cipher policy. Valid values are ` + "`" + `tls_cipher_policy_1_0` + "`" + `, ` + "`" + `tls_cipher_policy_1_1` + "`" + `, ` + "`" + `tls_cipher_policy_1_2` + "`" + `, ` + "`" + `tls_cipher_policy_1_2_strict` + "`" + `. Default to ` + "`" + `tls_cipher_policy_1_0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of slb listener.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_slb_master_slave_server_groups",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of master slave server groups related to a server load balancer to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `ID of the SLB.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of master slave server group IDs to filter results.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by master slave server group name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of SLB master slave server groups IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of SLB master slave server groups names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of SLB master slave server groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `master slave server group ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `master slave server group name.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `ECS instances associated to the group. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the attached ECS instance.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Weight associated to the ECS instance.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port used by the master slave server group.`,
				},
				resource.Attribute{
					Name:        "server_type",
					Description: `The server type of the attached ECS instance.`,
				},
				resource.Attribute{
					Name:        "is_backup",
					Description: `(Removed from v1.63.0) Determine if the server is executing.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of SLB master slave server groups IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of SLB master slave server groups names.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `A list of SLB master slave server groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `master slave server group ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `master slave server group name.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `ECS instances associated to the group. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the attached ECS instance.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Weight associated to the ECS instance.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port used by the master slave server group.`,
				},
				resource.Attribute{
					Name:        "server_type",
					Description: `The server type of the attached ECS instance.`,
				},
				resource.Attribute{
					Name:        "is_backup",
					Description: `(Removed from v1.63.0) Determine if the server is executing.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_slb_rules",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of server load balancer rules to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `ID of the SLB with listener rules.`,
				},
				resource.Attribute{
					Name:        "frontend_port",
					Description: `SLB listener port.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of rules IDs to filter results.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by rule name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of SLB listener rules IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of SLB listener rules names.`,
				},
				resource.Attribute{
					Name:        "slb_rules",
					Description: `A list of SLB listener rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Rule ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Rule name.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain name in the HTTP request where the rule applies (e.g. "`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Path in the HTTP request where the rule applies (e.g. "/image").`,
				},
				resource.Attribute{
					Name:        "server_group_id",
					Description: `ID of the linked VServer group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of SLB listener rules IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of SLB listener rules names.`,
				},
				resource.Attribute{
					Name:        "slb_rules",
					Description: `A list of SLB listener rules. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Rule ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Rule name.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain name in the HTTP request where the rule applies (e.g. "`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Path in the HTTP request where the rule applies (e.g. "/image").`,
				},
				resource.Attribute{
					Name:        "server_group_id",
					Description: `ID of the linked VServer group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_slb_server_certificates",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of slb server certificates.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of server certificates IDs to filter results.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by server certificate name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew, Available in 1.58.0+) The Id of resource group which the slb server certificates belongs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.66.0+) A mapping of tags to assign to the resource. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of SLB server certificates IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of SLB server certificates names.`,
				},
				resource.Attribute{
					Name:        "certificates",
					Description: `A list of SLB server certificates. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Server certificate ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Server certificate name.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `Server certificate fingerprint.`,
				},
				resource.Attribute{
					Name:        "common_name",
					Description: `Server certificate common name.`,
				},
				resource.Attribute{
					Name:        "subject_alternative_names",
					Description: `Server certificate subject alternative name list.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Server certificate expired time.`,
				},
				resource.Attribute{
					Name:        "expired_timestamp",
					Description: `Server certificate expired timestamp.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Server certificate created time.`,
				},
				resource.Attribute{
					Name:        "created_timestamp",
					Description: `Server certificate created timestamp.`,
				},
				resource.Attribute{
					Name:        "alicloud_certificate_id",
					Description: `Id of server certificate issued by alibaba cloud.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The Id of resource group which the slb server certificates belongs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Available in v1.66.0+) A mapping of tags to assign to the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of SLB server certificates IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of SLB server certificates names.`,
				},
				resource.Attribute{
					Name:        "certificates",
					Description: `A list of SLB server certificates. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Server certificate ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Server certificate name.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `Server certificate fingerprint.`,
				},
				resource.Attribute{
					Name:        "common_name",
					Description: `Server certificate common name.`,
				},
				resource.Attribute{
					Name:        "subject_alternative_names",
					Description: `Server certificate subject alternative name list.`,
				},
				resource.Attribute{
					Name:        "expired_time",
					Description: `Server certificate expired time.`,
				},
				resource.Attribute{
					Name:        "expired_timestamp",
					Description: `Server certificate expired timestamp.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Server certificate created time.`,
				},
				resource.Attribute{
					Name:        "created_timestamp",
					Description: `Server certificate created timestamp.`,
				},
				resource.Attribute{
					Name:        "alicloud_certificate_id",
					Description: `Id of server certificate issued by alibaba cloud.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `The Id of resource group which the slb server certificates belongs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Available in v1.66.0+) A mapping of tags to assign to the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_slb_server_groups",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of VServer groups related to a server load balancer to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "load_balancer_id",
					Description: `ID of the SLB.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of VServer group IDs to filter results.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by VServer group name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of SLB VServer groups IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of SLB VServer groups names.`,
				},
				resource.Attribute{
					Name:        "slb_server_groups",
					Description: `A list of SLB VServer groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `VServer group ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `VServer group name.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `ECS instances associated to the group. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the attached ECS instance.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Weight associated to the ECS instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of SLB VServer groups IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of SLB VServer groups names.`,
				},
				resource.Attribute{
					Name:        "slb_server_groups",
					Description: `A list of SLB VServer groups. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `VServer group ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `VServer group name.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `ECS instances associated to the group. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `ID of the attached ECS instance.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Weight associated to the ECS instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_slb_zones",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of availability zones for SLB that can be used by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "enable_details",
					Description: `(Optional) Default to false and only output ` + "`" + `id` + "`" + ` in the ` + "`" + `zones` + "`" + ` block. Set it to true can output more details.`,
				},
				resource.Attribute{
					Name:        "available_slb_address_type",
					Description: `(Optional) Filter the results by a slb instance address type. Can be either ` + "`" + `Vpc` + "`" + `, ` + "`" + `classic_internet` + "`" + ` or ` + "`" + `classic_intranet` + "`" + ``,
				},
				resource.Attribute{
					Name:        "available_slb_address_ip_version",
					Description: `(Optional) Filter the results by a slb instance address version. Can be either ` + "`" + `ipv4` + "`" + `, or ` + "`" + `ipv6` + "`" + `. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "slb_slave_zone_ids",
					Description: `A list of slb slave zone ids in which the slb master zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "slb_slave_zone_ids",
					Description: `A list of slb slave zone ids in which the slb master zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_slbs",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of server load balancers to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of SLBs IDs.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by SLB name.`,
				},
				resource.Attribute{
					Name:        "master_availability_zone",
					Description: `(Optional) Master availability zone of the SLBs.`,
				},
				resource.Attribute{
					Name:        "slave_availability_zone",
					Description: `(Optional) Slave availability zone of the SLBs.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `(Optional) Network type of the SLBs. Valid values: ` + "`" + `vpc` + "`" + ` and ` + "`" + `classic` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) ID of the VPC linked to the SLBs.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional) ID of the VSwitch linked to the SLBs.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Optional) Service address of the SLBs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A map of tags assigned to the SLB instances. The ` + "`" + `tags` + "`" + ` can have a maximum of 5 tag. It must be in the format: ` + "`" + `` + "`" + `` + "`" + ` data "alicloud_slbs" "taggedInstances" { tags = { tagKey1 = "tagValue1", tagKey2 = "tagValue2" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew, Available in 1.60.0+) The Id of resource group which SLB belongs. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of slb IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of slb names.`,
				},
				resource.Attribute{
					Name:        "slbs",
					Description: `A list of SLBs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the SLB.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the SLB belongs to.`,
				},
				resource.Attribute{
					Name:        "master_availability_zone",
					Description: `Master availability zone of the SLBs.`,
				},
				resource.Attribute{
					Name:        "slave_availability_zone",
					Description: `Slave availability zone of the SLBs.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `SLB current status. Possible values: ` + "`" + `inactive` + "`" + `, ` + "`" + `active` + "`" + ` and ` + "`" + `locked` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `SLB name.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `Network type of the SLB. Possible values: ` + "`" + `vpc` + "`" + ` and ` + "`" + `classic` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC the SLB belongs to.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `ID of the VSwitch the SLB belongs to.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `Service address of the SLB.`,
				},
				resource.Attribute{
					Name:        "internet",
					Description: `SLB addressType: internet if ` + "`" + `true` + "`" + `, intranet if ` + "`" + `false` + "`" + `. Must be ` + "`" + `false` + "`" + ` when ` + "`" + `network_type` + "`" + ` is ` + "`" + `vpc` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `SLB creation time.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the SLB instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of slb IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of slb names.`,
				},
				resource.Attribute{
					Name:        "slbs",
					Description: `A list of SLBs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the SLB.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `Region ID the SLB belongs to.`,
				},
				resource.Attribute{
					Name:        "master_availability_zone",
					Description: `Master availability zone of the SLBs.`,
				},
				resource.Attribute{
					Name:        "slave_availability_zone",
					Description: `Slave availability zone of the SLBs.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `SLB current status. Possible values: ` + "`" + `inactive` + "`" + `, ` + "`" + `active` + "`" + ` and ` + "`" + `locked` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `SLB name.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `Network type of the SLB. Possible values: ` + "`" + `vpc` + "`" + ` and ` + "`" + `classic` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC the SLB belongs to.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `ID of the VSwitch the SLB belongs to.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `Service address of the SLB.`,
				},
				resource.Attribute{
					Name:        "internet",
					Description: `SLB addressType: internet if ` + "`" + `true` + "`" + `, intranet if ` + "`" + `false` + "`" + `. Must be ` + "`" + `false` + "`" + ` when ` + "`" + `network_type` + "`" + ` is ` + "`" + `vpc` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `SLB creation time.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the SLB instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_snapshots",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to get a list of snapshot according to the specified filters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of snapshot IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of snapshots names.`,
				},
				resource.Attribute{
					Name:        "snapshots",
					Description: `A list of snapshots. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the snapshot.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the snapshot.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `Whether the snapshot is encrypted or not.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `Progress of snapshot creation, presented in percentage.`,
				},
				resource.Attribute{
					Name:        "source_disk_id",
					Description: `Source disk ID, which is retained after the source disk of the snapshot is deleted.`,
				},
				resource.Attribute{
					Name:        "source_disk_size",
					Description: `Size of the source disk, measured in GB.`,
				},
				resource.Attribute{
					Name:        "source_disk_type",
					Description: `Source disk attribute. Value range:`,
				},
				resource.Attribute{
					Name:        "product_code",
					Description: `Product code on the image market place.`,
				},
				resource.Attribute{
					Name:        "retention_days",
					Description: `The number of days that an automatic snapshot retains in the console for your instance.`,
				},
				resource.Attribute{
					Name:        "remain_time",
					Description: `The remaining time of a snapshot creation task, in seconds.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Creation time. Time of creation. It is represented according to ISO8601, and UTC time is used. Format: YYYY-MM-DDThh:mmZ.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The snapshot status. Value range:`,
				},
				resource.Attribute{
					Name:        "usage",
					Description: `Whether the snapshots are used to create resources or not. Value range:`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the snapshot.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_snat_entries",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of Snat Entries owned by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Snat Entries IDs.`,
				},
				resource.Attribute{
					Name:        "snat_ip",
					Description: `(Optional) The public IP of the Snat Entry.`,
				},
				resource.Attribute{
					Name:        "source_cidr",
					Description: `(Optional) The source CIDR block of the Snat Entry.`,
				},
				resource.Attribute{
					Name:        "snat_table_id",
					Description: `(Required) The ID of the Snat table.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Snat Entries IDs.`,
				},
				resource.Attribute{
					Name:        "entries",
					Description: `A list of Snat Entries. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Snat Entry.`,
				},
				resource.Attribute{
					Name:        "snat_ip",
					Description: `The public IP of the Snat Entry.`,
				},
				resource.Attribute{
					Name:        "source_cidr",
					Description: `The source CIDR block of the Snat Entry.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the Snat Entry.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of Snat Entries IDs.`,
				},
				resource.Attribute{
					Name:        "entries",
					Description: `A list of Snat Entries. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Snat Entry.`,
				},
				resource.Attribute{
					Name:        "snat_ip",
					Description: `The public IP of the Snat Entry.`,
				},
				resource.Attribute{
					Name:        "source_cidr",
					Description: `The source CIDR block of the Snat Entry.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the Snat Entry.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ssl_vpn_client_certs",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of SSL-VPN client certificates which owned by an Alicloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) IDs of the SSL-VPN client certificates.`,
				},
				resource.Attribute{
					Name:        "ssl_vpn_server_id",
					Description: `(Optional) Use the SSL-VPN server ID as the search key.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string of SSL-VPN client certificate name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) Save the result to the file. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of SSL-VPN client cert IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of SSL-VPN client cert names.`,
				},
				resource.Attribute{
					Name:        "ssl_vpn_client_certs",
					Description: `A list of SSL-VPN client certificates. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the SSL-VPN client certificate.`,
				},
				resource.Attribute{
					Name:        "ssl_vpn_server_id",
					Description: `ID of the SSL-VPN Server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SSL-VPN client certificate.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `The expiration time of the client certificate.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the client certificate. valid value:expiring-soon, normal, expired.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of SSL-VPN client cert IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of SSL-VPN client cert names.`,
				},
				resource.Attribute{
					Name:        "ssl_vpn_client_certs",
					Description: `A list of SSL-VPN client certificates. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the SSL-VPN client certificate.`,
				},
				resource.Attribute{
					Name:        "ssl_vpn_server_id",
					Description: `ID of the SSL-VPN Server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SSL-VPN client certificate.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `The expiration time of the client certificate.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the client certificate. valid value:expiring-soon, normal, expired.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_ssl_vpn_servers",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of SSL-VPN servers which owned by an Alicloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) IDs of the SSL-VPN servers.`,
				},
				resource.Attribute{
					Name:        "vpn_gateway_id",
					Description: `(Optional) Use the VPN gateway ID as the search key.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string of SSL-VPN server name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) Save the result to the file. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of SSL-VPN server IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of SSL-VPN server names.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `A list of SSL-VPN servers. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "vpn_gateway_id",
					Description: `The ID of the VPN gateway instance.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the SSL-VPN server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SSL-VPN server.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation.`,
				},
				resource.Attribute{
					Name:        "compress",
					Description: `Whether to compress.`,
				},
				resource.Attribute{
					Name:        "cipher",
					Description: `The encryption algorithm used.`,
				},
				resource.Attribute{
					Name:        "proto",
					Description: `The protocol used by the SSL-VPN server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port used by the SSL-VPN server.`,
				},
				resource.Attribute{
					Name:        "client_ip_pool",
					Description: `The IP address pool of the client.`,
				},
				resource.Attribute{
					Name:        "local_subnet",
					Description: `The local subnet of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "internet_ip",
					Description: `The public IP.`,
				},
				resource.Attribute{
					Name:        "connections",
					Description: `The number of current connections.`,
				},
				resource.Attribute{
					Name:        "max_connections",
					Description: `The maximum number of connections.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of SSL-VPN server IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of SSL-VPN server names.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `A list of SSL-VPN servers. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "vpn_gateway_id",
					Description: `The ID of the VPN gateway instance.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the SSL-VPN server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SSL-VPN server.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time of creation.`,
				},
				resource.Attribute{
					Name:        "compress",
					Description: `Whether to compress.`,
				},
				resource.Attribute{
					Name:        "cipher",
					Description: `The encryption algorithm used.`,
				},
				resource.Attribute{
					Name:        "proto",
					Description: `The protocol used by the SSL-VPN server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port used by the SSL-VPN server.`,
				},
				resource.Attribute{
					Name:        "client_ip_pool",
					Description: `The IP address pool of the client.`,
				},
				resource.Attribute{
					Name:        "local_subnet",
					Description: `The local subnet of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "internet_ip",
					Description: `The public IP.`,
				},
				resource.Attribute{
					Name:        "connections",
					Description: `The number of current connections.`,
				},
				resource.Attribute{
					Name:        "max_connections",
					Description: `The maximum number of connections.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_vpcs",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of VPCs owned by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Optional) Filter results by a specific CIDR block. For example: "172.16.0.0/12".`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Filter results by a specific status. Valid value are ` + "`" + `Pending` + "`" + ` and ` + "`" + `Available` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter VPCs by name.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `(Optional, type: bool) Indicate whether the VPC is the default one in the specified region.`,
				},
				resource.Attribute{
					Name:        "vswitch_id",
					Description: `(Optional) Filter results by the specified VSwitch.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.55.3+) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, Available in 1.52.0+) A list of VPC IDs.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew, Available in 1.60.0+) The Id of resource group which VPC belongs. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of VPC IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of VPC names.`,
				},
				resource.Attribute{
					Name:        "vpcs",
					Description: `A list of VPCs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `ID of the region where the VPC is located.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `Name of the VPC.`,
				},
				resource.Attribute{
					Name:        "vswitch_ids",
					Description: `List of VSwitch IDs in the specified VPC`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `CIDR block of the VPC.`,
				},
				resource.Attribute{
					Name:        "vrouter_id",
					Description: `ID of the VRouter.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `Route table ID of the VRouter.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the VPC`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Whether the VPC is the default VPC in the region.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Time of creation.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the VPC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of VPC IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of VPC names.`,
				},
				resource.Attribute{
					Name:        "vpcs",
					Description: `A list of VPCs. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `ID of the region where the VPC is located.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `Name of the VPC.`,
				},
				resource.Attribute{
					Name:        "vswitch_ids",
					Description: `List of VSwitch IDs in the specified VPC`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `CIDR block of the VPC.`,
				},
				resource.Attribute{
					Name:        "vrouter_id",
					Description: `ID of the VRouter.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `Route table ID of the VRouter.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the VPC`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Whether the VPC is the default VPC in the region.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Time of creation.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the VPC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_vpn_connections",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of VPN connections which owned by an Alicloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) IDs of the VPN connections.`,
				},
				resource.Attribute{
					Name:        "vpn_gateway_id",
					Description: `(Optional) Use the VPN gateway ID as the search key.`,
				},
				resource.Attribute{
					Name:        "customer_gateway_id",
					Description: `(Optional)Use the VPN customer gateway ID as the search key.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string of VPN connection name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) Save the result to the file. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) IDs of the VPN connections.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `(Optional) names of the VPN connections.`,
				},
				resource.Attribute{
					Name:        "connections",
					Description: `A list of VPN connections. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "customer_gateway_id",
					Description: `ID of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "vpn_gateway_id",
					Description: `ID of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "local_subnet",
					Description: `The local subnet of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "remote_subnet",
					Description: `The remote subnet of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the VPN connection, valid value:ike_sa_not_established, ike_sa_established, ipsec_sa_not_established, ipsec_sa_established.`,
				},
				resource.Attribute{
					Name:        "ike_config",
					Description: `The configurations of phase-one negotiation.`,
				},
				resource.Attribute{
					Name:        "ipsec_config",
					Description: `The configurations of phase-two negotiation. ### Block ike_config The ike_config mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "psk",
					Description: `Used for authentication between the IPsec VPN gateway and the customer gateway.`,
				},
				resource.Attribute{
					Name:        "ike_version",
					Description: `The version of the IKE protocol.`,
				},
				resource.Attribute{
					Name:        "ike_mode",
					Description: `The negotiation mode of IKE phase-one.`,
				},
				resource.Attribute{
					Name:        "ike_enc_alg",
					Description: `The encryption algorithm of phase-one negotiation.`,
				},
				resource.Attribute{
					Name:        "ike_auth_alg",
					Description: `The authentication algorithm of phase-one negotiation.`,
				},
				resource.Attribute{
					Name:        "ike_pfs",
					Description: `The Diffie-Hellman key exchange algorithm used by phase-one negotiation.`,
				},
				resource.Attribute{
					Name:        "ike_lifetime",
					Description: `The SA lifecycle as the result of phase-one negotiation.`,
				},
				resource.Attribute{
					Name:        "ike_local_id",
					Description: `The identification of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "ike_remote_id",
					Description: `The identification of the customer gateway. ### Block ipsec_config The ipsec_config mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "ipsec_enc_alg",
					Description: `The encryption algorithm of phase-two negotiation.`,
				},
				resource.Attribute{
					Name:        "ipsec_auth_alg",
					Description: `The authentication algorithm of phase-two negotiation.`,
				},
				resource.Attribute{
					Name:        "ipsec_pfs",
					Description: `The Diffie-Hellman key exchange algorithm used by phase-two negotiation.`,
				},
				resource.Attribute{
					Name:        "ipsec_lifetime",
					Description: `The SA lifecycle as the result of phase-two negotiation.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) IDs of the VPN connections.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `(Optional) names of the VPN connections.`,
				},
				resource.Attribute{
					Name:        "connections",
					Description: `A list of VPN connections. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "customer_gateway_id",
					Description: `ID of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "vpn_gateway_id",
					Description: `ID of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "local_subnet",
					Description: `The local subnet of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "remote_subnet",
					Description: `The remote subnet of the VPN connection.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the VPN connection, valid value:ike_sa_not_established, ike_sa_established, ipsec_sa_not_established, ipsec_sa_established.`,
				},
				resource.Attribute{
					Name:        "ike_config",
					Description: `The configurations of phase-one negotiation.`,
				},
				resource.Attribute{
					Name:        "ipsec_config",
					Description: `The configurations of phase-two negotiation. ### Block ike_config The ike_config mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "psk",
					Description: `Used for authentication between the IPsec VPN gateway and the customer gateway.`,
				},
				resource.Attribute{
					Name:        "ike_version",
					Description: `The version of the IKE protocol.`,
				},
				resource.Attribute{
					Name:        "ike_mode",
					Description: `The negotiation mode of IKE phase-one.`,
				},
				resource.Attribute{
					Name:        "ike_enc_alg",
					Description: `The encryption algorithm of phase-one negotiation.`,
				},
				resource.Attribute{
					Name:        "ike_auth_alg",
					Description: `The authentication algorithm of phase-one negotiation.`,
				},
				resource.Attribute{
					Name:        "ike_pfs",
					Description: `The Diffie-Hellman key exchange algorithm used by phase-one negotiation.`,
				},
				resource.Attribute{
					Name:        "ike_lifetime",
					Description: `The SA lifecycle as the result of phase-one negotiation.`,
				},
				resource.Attribute{
					Name:        "ike_local_id",
					Description: `The identification of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "ike_remote_id",
					Description: `The identification of the customer gateway. ### Block ipsec_config The ipsec_config mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "ipsec_enc_alg",
					Description: `The encryption algorithm of phase-two negotiation.`,
				},
				resource.Attribute{
					Name:        "ipsec_auth_alg",
					Description: `The authentication algorithm of phase-two negotiation.`,
				},
				resource.Attribute{
					Name:        "ipsec_pfs",
					Description: `The Diffie-Hellman key exchange algorithm used by phase-two negotiation.`,
				},
				resource.Attribute{
					Name:        "ipsec_lifetime",
					Description: `The SA lifecycle as the result of phase-two negotiation.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_vpn_customer_gateways",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of VPN customer gateways which owned by an Alicloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) ID of the VPN customer gateways.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string of VPN customer gateways name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) Save the result to the file. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "gateways",
					Description: `A list of VPN customer gateways. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPN customer gateway .`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The ip address of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The creation time of the VPN customer gateway.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "gateways",
					Description: `A list of VPN customer gateways. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPN customer gateway .`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The ip address of the VPN customer gateway.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The creation time of the VPN customer gateway.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_vpn_gateways",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of VPNs which owned by an Alicloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) Use the VPC ID as the search key.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) IDs of the VPN.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Limit search to specific status - valid value is "Init", "Provisioning", "Active", "Updating", "Deleting".`,
				},
				resource.Attribute{
					Name:        "business_status",
					Description: `(Optional) Limit search to specific business status - valid value is "Normal", "FinancialLocked".`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string of VPN name.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) Save the result to the file. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `IDs of the VPN.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `names of the VPN.`,
				},
				resource.Attribute{
					Name:        "gateways",
					Description: `A list of VPN gateways. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPN.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC that the VPN belongs.`,
				},
				resource.Attribute{
					Name:        "internet_ip",
					Description: `The internet ip of the VPN.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The creation time of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `The expiration time of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "specification",
					Description: `The Specification of the VPN`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the VPN.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the VPN`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the VPN`,
				},
				resource.Attribute{
					Name:        "business_status",
					Description: `The business status of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `The charge type of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "enable_ipsec",
					Description: `Whether the ipsec function is enabled.`,
				},
				resource.Attribute{
					Name:        "enable_ssl",
					Description: `Whether the ssl function is enabled.`,
				},
				resource.Attribute{
					Name:        "ssl_connections",
					Description: `Total count of ssl vpn connections.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `IDs of the VPN.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `names of the VPN.`,
				},
				resource.Attribute{
					Name:        "gateways",
					Description: `A list of VPN gateways. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPN.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC that the VPN belongs.`,
				},
				resource.Attribute{
					Name:        "internet_ip",
					Description: `The internet ip of the VPN.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The creation time of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `The expiration time of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "specification",
					Description: `The Specification of the VPN`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the VPN.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the VPN`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the VPN`,
				},
				resource.Attribute{
					Name:        "business_status",
					Description: `The business status of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `The charge type of the VPN gateway.`,
				},
				resource.Attribute{
					Name:        "enable_ipsec",
					Description: `Whether the ipsec function is enabled.`,
				},
				resource.Attribute{
					Name:        "enable_ssl",
					Description: `Whether the ssl function is enabled.`,
				},
				resource.Attribute{
					Name:        "ssl_connections",
					Description: `Total count of ssl vpn connections.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_vswitches",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of VSwitch owned by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Optional) Filter results by a specific CIDR block. For example: "172.16.0.0/12".`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional) The availability zone of the VSwitch.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) A regex string to filter results by name.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `(Optional, type: bool) Indicate whether the VSwitch is created by the system.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) ID of the VPC that owns the VSwitch.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.55.3+) A mapping of tags to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional, Available in 1.52.0+) A list of VSwitch IDs.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional, ForceNew, Available in 1.60.0+) The Id of resource group which VSWitch belongs. ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of VSwitch IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of VSwitch names.`,
				},
				resource.Attribute{
					Name:        "vswitches",
					Description: `A list of VSwitches. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VSwitch.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `ID of the availability zone where the VSwitch is located.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC that owns the VSwitch.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the VSwitch.`,
				},
				resource.Attribute{
					Name:        "instance_ids",
					Description: `List of ECS instance IDs in the specified VSwitch.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `CIDR block of the VSwitch.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the VSwitch.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Whether the VSwitch is the default one in the region.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Time of creation.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of VSwitch IDs.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of VSwitch names.`,
				},
				resource.Attribute{
					Name:        "vswitches",
					Description: `A list of VSwitches. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VSwitch.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `ID of the availability zone where the VSwitch is located.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC that owns the VSwitch.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the VSwitch.`,
				},
				resource.Attribute{
					Name:        "instance_ids",
					Description: `List of ECS instance IDs in the specified VSwitch.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `CIDR block of the VSwitch.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the VSwitch.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Whether the VSwitch is the default one in the region.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `Time of creation.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_waf_domains",
			Category:         "Data Sources",
			ShortDescription: `Provides a datasource to retrieve domain names.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of WAF domain names. Each item is domain name.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) The Id of waf instance to which waf domain belongs.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of WAF domain names. Each item is domain name.`,
				},
				resource.Attribute{
					Name:        "domains",
					Description: `A list of Domains. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Name of the domain. ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of WAF domain names. Each item is domain name.`,
				},
				resource.Attribute{
					Name:        "domains",
					Description: `A list of Domains. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Name of the domain. ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_waf_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a datasource to retrieve WAF instances.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of WAF instance IDs.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of WAF instance to filter results. Optional value: ` + "`" + `0` + "`" + `: The instance has expired, ` + "`" + `1` + "`" + ` : The instance has not expired and is working properly.`,
				},
				resource.Attribute{
					Name:        "resource_group_id",
					Description: `(Optional) The ID of resource group to which WAF instance belongs.`,
				},
				resource.Attribute{
					Name:        "instance_source",
					Description: `(Optional) The source of the WAF instance.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `). ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of WAF instance IDs.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of WAF instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the WAF instance.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of WAF the instance.`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `The timestamp (in seconds) indicating when the WAF instance expires.`,
				},
				resource.Attribute{
					Name:        "in_debt",
					Description: `Indicates whether the WAF instance has overdue payments.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region where the WAF instance is located.`,
				},
				resource.Attribute{
					Name:        "remain_day",
					Description: `The number of days before the trial period of the WAF instance expires.`,
				},
				resource.Attribute{
					Name:        "trial",
					Description: `Indicates whether this is a trial instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates whether the WAF instance has expired.`,
				},
				resource.Attribute{
					Name:        "subscription_typed",
					Description: `The billing method of the WAF instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) A list of WAF instance IDs.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of WAF instances. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the WAF instance.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of WAF the instance.`,
				},
				resource.Attribute{
					Name:        "end_date",
					Description: `The timestamp (in seconds) indicating when the WAF instance expires.`,
				},
				resource.Attribute{
					Name:        "in_debt",
					Description: `Indicates whether the WAF instance has overdue payments.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region where the WAF instance is located.`,
				},
				resource.Attribute{
					Name:        "remain_day",
					Description: `The number of days before the trial period of the WAF instance expires.`,
				},
				resource.Attribute{
					Name:        "trial",
					Description: `Indicates whether this is a trial instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates whether the WAF instance has expired.`,
				},
				resource.Attribute{
					Name:        "subscription_typed",
					Description: `The billing method of the WAF instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_yundun_bastionhost_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of cloud Bastionhost(yundun_bastionhost) instances available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description_regex",
					Description: `(Optional) A regex string to filter results by the instance description.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) Matched instance IDs to filter data source result.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name to persist data source output.`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `(Optional) Descriptions to filter data source result.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.67.0+) A map of tags assigned to the bastionhost instance. It must be in the format: ` + "`" + `` + "`" + `` + "`" + ` data "alicloud_yundun_bastionhost_instances" "instance" { tags = { tagKey1 = "tagValue1" } } ` + "`" + `` + "`" + `` + "`" + ` ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of apis. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The instance's id.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The instance's remark.`,
				},
				resource.Attribute{
					Name:        "user_vswitch_id",
					Description: `The instance's vSwitch ID.`,
				},
				resource.Attribute{
					Name:        "private_domain",
					Description: `The instance's private domain name.`,
				},
				resource.Attribute{
					Name:        "public_domain",
					Description: `The instance's public domain name.`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `The instance's status.`,
				},
				resource.Attribute{
					Name:        "public_network_access",
					Description: `The instance's public network access configuration.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `The instance's security group configuration.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the bastionhost instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instances",
					Description: `A list of apis. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The instance's id.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The instance's remark.`,
				},
				resource.Attribute{
					Name:        "user_vswitch_id",
					Description: `The instance's vSwitch ID.`,
				},
				resource.Attribute{
					Name:        "private_domain",
					Description: `The instance's private domain name.`,
				},
				resource.Attribute{
					Name:        "public_domain",
					Description: `The instance's public domain name.`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `The instance's status.`,
				},
				resource.Attribute{
					Name:        "public_network_access",
					Description: `The instance's public network access configuration.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `The instance's security group configuration.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the bastionhost instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_yundun_dbaudit_instances",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of cloud DBaudit(yundun_dbaudit) instances available to the user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description_regex",
					Description: `(Optional) A regex string to filter results by the instance description.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `(Optional) Matched instance IDs to filter data source result.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name to persist data source output.`,
				},
				resource.Attribute{
					Name:        "descriptions",
					Description: `(Optional) Descriptions to filter data source result.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, Available in v1.67.0+) A map of tags assigned to the dbaudit instance. It must be in the format: ` + "`" + `` + "`" + `` + "`" + ` data "alicloud_yundun_dbaudit_instances" "instance" { tags = { tagKey1 = "tagValue1" } } ` + "`" + `` + "`" + `` + "`" + ` ## Attributes Reference The following attributes are exported in addition to the arguments listed above:`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `A list of apis. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The instance's id.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The instance's remark.`,
				},
				resource.Attribute{
					Name:        "user_vswitch_id",
					Description: `The instance's vSwitch ID.`,
				},
				resource.Attribute{
					Name:        "private_domain",
					Description: `The instance's private domain name.`,
				},
				resource.Attribute{
					Name:        "public_domain",
					Description: `The instance's public domain name.`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `The instance's status.`,
				},
				resource.Attribute{
					Name:        "public_network_access",
					Description: `The instance's public network access configuration.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the dbaudit instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instances",
					Description: `A list of apis. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The instance's id.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The instance's remark.`,
				},
				resource.Attribute{
					Name:        "user_vswitch_id",
					Description: `The instance's vSwitch ID.`,
				},
				resource.Attribute{
					Name:        "private_domain",
					Description: `The instance's private domain name.`,
				},
				resource.Attribute{
					Name:        "public_domain",
					Description: `The instance's public domain name.`,
				},
				resource.Attribute{
					Name:        "instance_status",
					Description: `The instance's status.`,
				},
				resource.Attribute{
					Name:        "public_network_access",
					Description: `The instance's public network access configuration.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of tags assigned to the dbaudit instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alicloud_zones",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of availability zones that can be used by an Alibaba Cloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "available_instance_type",
					Description: `(Optional) Filter the results by a specific instance type.`,
				},
				resource.Attribute{
					Name:        "available_resource_creation",
					Description: `(Optional) Filter the results by a specific resource type. Valid values: ` + "`" + `Instance` + "`" + `, ` + "`" + `Disk` + "`" + `, ` + "`" + `VSwitch` + "`" + `, ` + "`" + `Rds` + "`" + `, ` + "`" + `KVStore` + "`" + `, ` + "`" + `FunctionCompute` + "`" + `, ` + "`" + `Elasticsearch` + "`" + `, ` + "`" + `Slb` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "available_disk_category",
					Description: `(Optional) Filter the results by a specific disk category. Can be either ` + "`" + `cloud` + "`" + `, ` + "`" + `cloud_efficiency` + "`" + `, ` + "`" + `cloud_ssd` + "`" + `, ` + "`" + `ephemeral_ssd` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "multi",
					Description: `(Optional, type: bool) Indicate whether the zones can be used in a multi AZ configuration. Default to ` + "`" + `false` + "`" + `. Multi AZ is usually used to launch RDS instances.`,
				},
				resource.Attribute{
					Name:        "instance_charge_type",
					Description: `(Optional) Filter the results by a specific ECS instance charge type. Valid values: ` + "`" + `PrePaid` + "`" + ` and ` + "`" + `PostPaid` + "`" + `. Default to ` + "`" + `PostPaid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `(Optional) Filter the results by a specific network type. Valid values: ` + "`" + `Classic` + "`" + ` and ` + "`" + `Vpc` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "spot_strategy",
					Description: `- (Optional) Filter the results by a specific ECS spot type. Valid values: ` + "`" + `NoSpot` + "`" + `, ` + "`" + `SpotWithPriceLimit` + "`" + ` and ` + "`" + `SpotAsPriceGo` + "`" + `. Default to ` + "`" + `NoSpot` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "output_file",
					Description: `(Optional) File name where to save data source results (after running ` + "`" + `terraform plan` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "enable_details",
					Description: `(Optional, Available in 1.36.0+) Default to false and only output ` + "`" + `id` + "`" + ` in the ` + "`" + `zones` + "`" + ` block. Set it to true can output more details.`,
				},
				resource.Attribute{
					Name:        "available_slb_address_type",
					Description: `(Available in 1.45.0+) Filter the results by a slb instance address type. Can be either ` + "`" + `Vpc` + "`" + `, ` + "`" + `classic_internet` + "`" + ` or ` + "`" + `classic_intranet` + "`" + ``,
				},
				resource.Attribute{
					Name:        "available_slb_address_ip_version",
					Description: `(Available in 1.45.0+) Filter the results by a slb instance address version. Can be either ` + "`" + `ipv4` + "`" + `, or ` + "`" + `ipv6` + "`" + `. ->`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "local_name",
					Description: `Name of the zone in the local language.`,
				},
				resource.Attribute{
					Name:        "available_instance_types",
					Description: `Allowed instance types.`,
				},
				resource.Attribute{
					Name:        "available_resource_creation",
					Description: `Type of resources that can be created.`,
				},
				resource.Attribute{
					Name:        "available_disk_categories",
					Description: `Set of supported disk categories.`,
				},
				resource.Attribute{
					Name:        "multi_zone_ids",
					Description: `A list of zone ids in which the multi zone.`,
				},
				resource.Attribute{
					Name:        "slb_slave_zone_ids",
					Description: `A list of slb slave zone ids in which the slb master zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of zone IDs.`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `A list of availability zones. Each element contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the zone.`,
				},
				resource.Attribute{
					Name:        "local_name",
					Description: `Name of the zone in the local language.`,
				},
				resource.Attribute{
					Name:        "available_instance_types",
					Description: `Allowed instance types.`,
				},
				resource.Attribute{
					Name:        "available_resource_creation",
					Description: `Type of resources that can be created.`,
				},
				resource.Attribute{
					Name:        "available_disk_categories",
					Description: `Set of supported disk categories.`,
				},
				resource.Attribute{
					Name:        "multi_zone_ids",
					Description: `A list of zone ids in which the multi zone.`,
				},
				resource.Attribute{
					Name:        "slb_slave_zone_ids",
					Description: `A list of slb slave zone ids in which the slb master zone.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"alicloud_account":                               0,
		"alicloud_actiontrails":                          1,
		"alicloud_adb_clusters":                          2,
		"alicloud_adb_zones":                             3,
		"alicloud_alidns_domain_groups":                  4,
		"alicloud_alidns_records":                        5,
		"alicloud_alikafka_consumer_groups":              6,
		"alicloud_alikafka_instances":                    7,
		"alicloud_alikafka_sasl_acls":                    8,
		"alicloud_alikafka_sasl_users":                   9,
		"alicloud_alikafka_topics":                       10,
		"alicloud_api_gateway_apis":                      11,
		"alicloud_api_gateway_apps":                      12,
		"alicloud_api_gateway_groups":                    13,
		"alicloud_caller_identity":                       14,
		"alicloud_cas_certificates":                      15,
		"alicloud_cassandra_clusters":                    16,
		"alicloud_cassandra_data_centers":                17,
		"alicloud_cassandra_zones":                       18,
		"alicloud_cen_bandwidth_limits":                  19,
		"alicloud_cen_bandwidth_packages":                20,
		"alicloud_cen_flowlogs":                          21,
		"alicloud_cen_instances":                         22,
		"alicloud_cen_private_zones":                     23,
		"alicloud_cen_region_route_entries":              24,
		"alicloud_cen_route_entries":                     25,
		"alicloud_cen_route_maps":                        26,
		"alicloud_cloud_connect_networks":                27,
		"alicloud_common_bandwidth_packages":             28,
		"alicloud_cr_ee_instances":                       29,
		"alicloud_cr_ee_namespaces":                      30,
		"alicloud_cr_ee_repos":                           31,
		"alicloud_cr_ee_sync_rules":                      32,
		"alicloud_cr_namespaces":                         33,
		"alicloud_cr_repos":                              34,
		"alicloud_cs_kubernetes_clusters":                35,
		"alicloud_cs_managed_kubernetes_clusters":        36,
		"alicloud_cs_serverless_kubernetes_clusters":     37,
		"alicloud_db_instance_classes":                   38,
		"alicloud_db_instance_engines":                   39,
		"alicloud_db_instances":                          40,
		"alicloud_db_zones":                              41,
		"alicloud_ddosbgp_instances":                     42,
		"alicloud_ddoscoo_instances":                     43,
		"alicloud_disks":                                 44,
		"alicloud_dms_enterprise_instances":              45,
		"alicloud_dms_enterprise_users":                  46,
		"alicloud_dns_domain_groups":                     47,
		"alicloud_dns_domain_records":                    48,
		"alicloud_dns_domain_txt_guid":                   49,
		"alicloud_dns_domains":                           50,
		"alicloud_dns_groups":                            51,
		"alicloud_dns_instances":                         52,
		"alicloud_dns_records":                           53,
		"alicloud_resolution_lines":                      54,
		"alicloud_drds_instances":                        55,
		"alicloud_eci_image_caches":                      56,
		"alicloud_ecs_dedicated_hosts":                   57,
		"alicloud_edas_applications":                     58,
		"alicloud_edas_clusters":                         59,
		"alicloud_edas_deploy_groups":                    60,
		"alicloud_eips":                                  61,
		"alicloud_elasticsearch_instances":               62,
		"alicloud_elasticsearch_zones":                   63,
		"alicloud_emr_disk_types":                        64,
		"alicloud_emr_instance_types":                    65,
		"alicloud_emr_main_versions":                     66,
		"alicloud_ess_alarms":                            67,
		"alicloud_ess_lifecycle_hooks":                   68,
		"alicloud_ess_notifications":                     69,
		"alicloud_ess_scaling_configurations":            70,
		"alicloud_ess_scaling_groups":                    71,
		"alicloud_ess_scaling_rules":                     72,
		"alicloud_ess_scheduled_tasks":                   73,
		"alicloud_fc_functions":                          74,
		"alicloud_fc_services":                           75,
		"alicloud_fc_triggers":                           76,
		"alicloud_fc_zones":                              77,
		"alicloud_file_crc64_checksum":                   78,
		"alicloud_forward_entries":                       79,
		"alicloud_gpdb_instances":                        80,
		"alicloud_hbase_instances":                       81,
		"alicloud_hbase_zones":                           82,
		"alicloud_images":                                83,
		"alicloud_instance_type_families":                84,
		"alicloud_instance_types":                        85,
		"alicloud_instances":                             86,
		"alicloud_key_pairs":                             87,
		"alicloud_kms_aliases":                           88,
		"alicloud_kms_ciphertext":                        89,
		"alicloud_kms_key_versions":                      90,
		"alicloud_kms_keys":                              91,
		"alicloud_kms_plaintext":                         92,
		"alicloud_kms_secret_versions":                   93,
		"alicloud_kms_secrets":                           94,
		"alicloud_kvstore_instance_classes":              95,
		"alicloud_kvstore_instance_engines":              96,
		"alicloud_kvstore_instances":                     97,
		"alicloud_kvstore_zones":                         98,
		"alicloud_market_product":                        99,
		"alicloud_market_products":                       100,
		"alicloud_mns_queues":                            101,
		"alicloud_mns_topic_subscriptions":               102,
		"alicloud_mns_topics":                            103,
		"alicloud_mongodb_instances":                     104,
		"alicloud_mongodb_zones":                         105,
		"alicloud_nas_access_groups":                     106,
		"alicloud_nas_access_rules":                      107,
		"alicloud_nas_file_systems":                      108,
		"alicloud_nas_mount_targets":                     109,
		"alicloud_nas_protocols":                         110,
		"alicloud_nat_gateways":                          111,
		"alicloud_network_interfaces":                    112,
		"alicloud_ons_groups":                            113,
		"alicloud_ons_instances":                         114,
		"alicloud_ons_topics":                            115,
		"alicloud_oos_templates":                         116,
		"alicloud_oss_bucket_objects":                    117,
		"alicloud_oss_buckets":                           118,
		"alicloud_ots_instance_attachments":              119,
		"alicloud_ots_instances":                         120,
		"alicloud_ots_tables":                            121,
		"alicloud_polardb_accounts":                      122,
		"alicloud_polardb_clusters":                      123,
		"alicloud_polardb_databases":                     124,
		"alicloud_polardb_endpoints":                     125,
		"alicloud_polardb_node_classes":                  126,
		"alicloud_polardb_zones":                         127,
		"alicloud_pvtz_zone_records":                     128,
		"alicloud_pvtz_zones":                            129,
		"alicloud_ram_account_aliases":                   130,
		"alicloud_ram_account_alias":                     131,
		"alicloud_ram_groups":                            132,
		"alicloud_ram_policies":                          133,
		"alicloud_ram_roles":                             134,
		"alicloud_ram_users":                             135,
		"alicloud_regions":                               136,
		"alicloud_resource_manager_accounts":             137,
		"alicloud_resource_manager_folders":              138,
		"alicloud_resource_manager_handshakes":           139,
		"alicloud_resource_manager_policies":             140,
		"alicloud_resource_manager_policy_versions":      141,
		"alicloud_resource_manager_resource_directories": 142,
		"alicloud_resource_manager_resource_groups":      143,
		"alicloud_resource_manager_roles":                144,
		"alicloud_route_entries":                         145,
		"alicloud_route_tables":                          146,
		"alicloud_router_interfaces":                     147,
		"alicloud_sag_acls":                              148,
		"alicloud_security_group_rules":                  149,
		"alicloud_security_groups":                       150,
		"alicloud_slb_acls":                              151,
		"alicloud_slb_attachments":                       152,
		"alicloud_slb_backend_servers":                   153,
		"alicloud_slb_ca_certificates":                   154,
		"alicloud_slb_domain_extensions":                 155,
		"alicloud_slb_listeners":                         156,
		"alicloud_slb_master_slave_server_groups":        157,
		"alicloud_slb_rules":                             158,
		"alicloud_slb_server_certificates":               159,
		"alicloud_slb_server_groups":                     160,
		"alicloud_slb_zones":                             161,
		"alicloud_slbs":                                  162,
		"alicloud_snapshots":                             163,
		"alicloud_snat_entries":                          164,
		"alicloud_ssl_vpn_client_certs":                  165,
		"alicloud_ssl_vpn_servers":                       166,
		"alicloud_vpcs":                                  167,
		"alicloud_vpn_connections":                       168,
		"alicloud_vpn_customer_gateways":                 169,
		"alicloud_vpn_gateways":                          170,
		"alicloud_vswitches":                             171,
		"alicloud_waf_domains":                           172,
		"alicloud_waf_instances":                         173,
		"alicloud_yundun_bastionhost_instances":          174,
		"alicloud_yundun_dbaudit_instances":              175,
		"alicloud_zones":                                 176,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
