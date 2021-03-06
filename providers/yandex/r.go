package yandex

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "yandex_compute_disk",
			Category:         "Yandex Compute Service Resources",
			ShortDescription: `Persistent disks are durable storage devices that function similarly to the physical disks in a desktop or a server.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"service",
				"disk",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the disk. Provide this property when you create a resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the disk. Provide this property when you create a resource.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the folder that the disk belongs to. If it is not provided, the default provider folder is used.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Labels to assign to this disk. A list of key/value pairs.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Availability zone where the disk will reside.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Size of the persistent disk, specified in GB. You can specify this field when creating a persistent disk using the ` + "`" + `image_id` + "`" + ` or ` + "`" + `snapshot_id` + "`" + ` parameter, or specify it alone to create an empty persistent disk. If you specify this field along with ` + "`" + `image_id` + "`" + ` or ` + "`" + `snapshot_id` + "`" + `, the size value must not be less than the size of the source image or the size of the snapshot.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of disk to create. Provide this when creating a disk. One of ` + "`" + `network-hdd` + "`" + ` (default) or ` + "`" + `network-ssd` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) The source image to use for disk creation.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The source snapshot to use for disk creation. ~>`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the disk.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the disk. ## Timeouts This resource provides the following configuration options for [timeouts](/docs/configuration/resources.html#timeouts): - ` + "`" + `create` + "`" + ` - Default is 5 minutes. - ` + "`" + `update` + "`" + ` - Default is 5 minutes. - ` + "`" + `delete` + "`" + ` - Default is 5 minutes. ## Import A disk can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_compute_disk.default disk_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `The status of the disk.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the disk. ## Timeouts This resource provides the following configuration options for [timeouts](/docs/configuration/resources.html#timeouts): - ` + "`" + `create` + "`" + ` - Default is 5 minutes. - ` + "`" + `update` + "`" + ` - Default is 5 minutes. - ` + "`" + `delete` + "`" + ` - Default is 5 minutes. ## Import A disk can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_compute_disk.default disk_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_compute_image",
			Category:         "Yandex Compute Service Resources",
			ShortDescription: `Creates a VM image for the Yandex Compute service from an existing tarball.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"service",
				"image",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the disk.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of the image. Provide this property when you create a resource.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs to assign to the image.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `(Optional) The name of the image family to which this image belongs.`,
				},
				resource.Attribute{
					Name:        "min_disk_size",
					Description: `(Optional) Minimum size in GB of the disk that will be created from this image.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `(Optional) Operating system type that is contained in the image. Possible values: "LINUX", "WINDOWS".`,
				},
				resource.Attribute{
					Name:        "source_family",
					Description: `(Optional) The name of the family to use as the source of the new image. The ID of the latest image is taken from the "standard-images" folder. Changing the family forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "source_image",
					Description: `(Optional) The ID of an existing image to use as the source of the image. Changing this ID forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "source_snapshot",
					Description: `(Optional) The ID of a snapshot to use as the source of the image. Changing this ID forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "source_disk",
					Description: `(Optional) The ID of a disk to use as the source of the image. Changing this ID forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "source_url",
					Description: `(Optional) The URL to use as the source of the image. Changing this URL forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "product_ids",
					Description: `(Optional) License IDs that indicate which licenses are attached to this image. ~>`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the image, specified in GB.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the image.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the image. ## Timeouts ` + "`" + `yandex_compute_image` + "`" + ` provides the following configuration options for [timeouts](/docs/configuration/resources.html#timeouts): - ` + "`" + `create` + "`" + ` - Default 5 minutes - ` + "`" + `update` + "`" + ` - Default 5 minutes - ` + "`" + `delete` + "`" + ` - Default 5 minutes ## Import A VM image can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_compute_image.web-image image_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "size",
					Description: `The size of the image, specified in GB.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the image.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the image. ## Timeouts ` + "`" + `yandex_compute_image` + "`" + ` provides the following configuration options for [timeouts](/docs/configuration/resources.html#timeouts): - ` + "`" + `create` + "`" + ` - Default 5 minutes - ` + "`" + `update` + "`" + ` - Default 5 minutes - ` + "`" + `delete` + "`" + ` - Default 5 minutes ## Import A VM image can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_compute_image.web-image image_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_compute_instance",
			Category:         "Yandex Compute Service Resources",
			ShortDescription: `Manages a VM instance resource.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"service",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) Compute resources that are allocated for the instance. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "boot_disk",
					Description: `(Required) The boot disk for the instance. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "network_interface",
					Description: `(Required) Networks to attach to the instance. This can be specified multiple times. The structure is documented below. - - -`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Resource name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the instance.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs to assign to the instance.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) The availability zone where the virtual machine will be created. If it is not provided, the default provider folder is used.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) Host name for the instance. This field is used to generate the instance ` + "`" + `fqdn` + "`" + ` value. The host name must be unique within the network and region. If not specified, the host name will be equal to ` + "`" + `id` + "`" + ` of the instance and ` + "`" + `fqdn` + "`" + ` will be ` + "`" + `<id>.auto.internal` + "`" + `. Otherwise FQDN will be ` + "`" + `<hostname>.<region_id>.internal` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Metadata key/value pairs to make available from within the instance.`,
				},
				resource.Attribute{
					Name:        "platform_id",
					Description: `(Optional) The type of virtual machine to create. The default is 'standard-v1'.`,
				},
				resource.Attribute{
					Name:        "secondary_disk",
					Description: `(Optional) A list of disks to attach to the instance. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "scheduling_policy",
					Description: `(Optional) Scheduling policy configuration. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "service_account_id",
					Description: `(Optional) ID of the service account authorized for this instance.`,
				},
				resource.Attribute{
					Name:        "allow_stopping_for_update",
					Description: `(Optional) If true, allows Terraform to stop the instance in order to update its properties. If you try to update a property that requires stopping the instance without setting this field, the update will fail.`,
				},
				resource.Attribute{
					Name:        "network_acceleration_type",
					Description: `(Optional) Type of network acceleration. The default is ` + "`" + `standard` + "`" + `. Values: ` + "`" + `standard` + "`" + `, ` + "`" + `software_accelerated` + "`" + ` --- The ` + "`" + `resources` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "cores",
					Description: `(Required) CPU cores for the instance.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Required) Memory size in GB.`,
				},
				resource.Attribute{
					Name:        "core_fraction",
					Description: `(Optional) If provided, specifies baseline performance for a core as a percent. The ` + "`" + `boot_disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "auto_delete",
					Description: `(Optional) Defines whether the disk will be auto-deleted when the instance is deleted. The default value is ` + "`" + `True` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(Optional) Name that can be used to access an attached disk.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Type of access to the disk resource. By default, a disk is attached in ` + "`" + `READ_WRITE` + "`" + ` mode.`,
				},
				resource.Attribute{
					Name:        "disk_id",
					Description: `(Optional) The ID of the existing disk (such as those managed by ` + "`" + `yandex_compute_disk` + "`" + `) to attach as a boot disk.`,
				},
				resource.Attribute{
					Name:        "initialize_params",
					Description: `(Optional) Parameters for a new disk that will be created alongside the new instance. Either ` + "`" + `initialize_params` + "`" + ` or ` + "`" + `disk_id` + "`" + ` must be set. The structure is documented below. ~>`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the boot disk.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the boot disk.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Size of the disk in GB.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Disk type.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) A disk image to initialize this disk from.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) A snapshot to initialize this disk from. ~>`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) ID of the subnet to attach this interface to. The subnet must exist in the same zone where this instance will be created.`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `(Optional) Allocate an IPv4 address for the interface. The default value is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) The private IP address to assign to the instance. If empty, the address will be automatically assigned from the specified subnet.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `(Optional) If true, allocate an IPv6 address for the interface. The address will be automatically assigned from the specified subnet.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional) The private IPv6 address to assign to the instance.`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `(Optional) Provide a public address, for instance, to access the internet over NAT.`,
				},
				resource.Attribute{
					Name:        "nat_ip_address",
					Description: `(Optional) Provide a public address, for instance, to access the internet over NAT. Address should be already reserved in web UI.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) Security group ids for network interface. The ` + "`" + `secondary_disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "disk_id",
					Description: `(Required) ID of the disk that is attached to the instance.`,
				},
				resource.Attribute{
					Name:        "auto_delete",
					Description: `(Optional) Whether the disk is auto-deleted when the instance is deleted. The default value is false.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(Optional) Name that can be used to access an attached disk under ` + "`" + `/dev/disk/by-id/` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Type of access to the disk resource. By default, a disk is attached in ` + "`" + `READ_WRITE` + "`" + ` mode. The ` + "`" + `scheduling_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "preemptible",
					Description: `(Optional) Specifies if the instance is preemptible. Defaults to false. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The fully qualified DNS name of this instance.`,
				},
				resource.Attribute{
					Name:        "network_interface.0.ip_address",
					Description: `The internal IP address of the instance.`,
				},
				resource.Attribute{
					Name:        "network_interface.0.nat_ip_address",
					Description: `The external IP address of the instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of this instance.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the instance. ## Import Instances can be imported using the ` + "`" + `ID` + "`" + ` of an instance, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_compute_instance.default instance_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "fqdn",
					Description: `The fully qualified DNS name of this instance.`,
				},
				resource.Attribute{
					Name:        "network_interface.0.ip_address",
					Description: `The internal IP address of the instance.`,
				},
				resource.Attribute{
					Name:        "network_interface.0.nat_ip_address",
					Description: `The external IP address of the instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of this instance.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the instance. ## Import Instances can be imported using the ` + "`" + `ID` + "`" + ` of an instance, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_compute_instance.default instance_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_compute_instance_group",
			Category:         "Yandex Compute Service Resources",
			ShortDescription: `Manages an Instance group resource.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"service",
				"instance",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Required) The ID of the folder that the resources belong to.`,
				},
				resource.Attribute{
					Name:        "scale_policy",
					Description: `(Required) The scaling policy of the instance group. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "deploy_policy",
					Description: `(Required) The deployment policy of the instance group. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "service_account_id",
					Description: `(Required) The ID of the service account authorized for this instance group.`,
				},
				resource.Attribute{
					Name:        "instance_template",
					Description: `(Required) The template for creating new instances. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "allocation_policy",
					Description: `(Required) The allocation policy of the instance group by zone and region. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the instance group.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `(Optional) Health check specifications. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "load_balancer",
					Description: `(Optional) Load balancing specifications. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the instance group.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs to assign to the instance group.`,
				},
				resource.Attribute{
					Name:        "variables",
					Description: `(Optional) A set of key/value variables pairs to assign to the instance group.`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `(Optional) Flag that protects the instance group from accidental deletion. --- The ` + "`" + `load_balancer` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "target_group_name",
					Description: `(Optional) The name of the target group.`,
				},
				resource.Attribute{
					Name:        "target_group_description",
					Description: `(Optional) A description of the target group.`,
				},
				resource.Attribute{
					Name:        "target_group_labels",
					Description: `(Optional) A set of key/value label pairs. --- The ` + "`" + `health_check` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) The interval to wait between health checks in seconds.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The length of time to wait for a response before the health check times out in seconds.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `(Optional) The number of successful health checks before the managed instance is declared healthy.`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `(Optional) The number of failed health checks before the managed instance is declared unhealthy.`,
				},
				resource.Attribute{
					Name:        "tcp_options",
					Description: `(Optional) TCP check options. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "http_options",
					Description: `(Optional) HTTP check options. The structure is documented below. --- The ` + "`" + `http_options` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The port used for HTTP health checks.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The URL path used for health check requests. --- The ` + "`" + `tcp_options` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The port used for TCP health checks. --- The ` + "`" + `allocation_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "zones",
					Description: `(Required) A list of availability zones. --- The ` + "`" + `instance_template` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "boot_disk",
					Description: `(Required) Boot disk specifications for the instance. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) Compute resource specifications for the instance. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "network_interface",
					Description: `(Required) Network specifications for the instance. This can be used multiple times for adding multiple interfaces. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "scheduling_policy",
					Description: `(Optional) The scheduling policy configuration. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the instance.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) A set of metadata key/value pairs to make available from within the instance.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs to assign to the instance.`,
				},
				resource.Attribute{
					Name:        "platform_id",
					Description: `(Optional) The ID of the hardware platform configuration for the instance. The default is 'standard-v1'.`,
				},
				resource.Attribute{
					Name:        "secondary_disk",
					Description: `(Optional) A list of disks to attach to the instance. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "service_account_id",
					Description: `(Optional) The ID of the service account authorized for this instance.`,
				},
				resource.Attribute{
					Name:        "network_settings",
					Description: `(Optional) Network acceleration type for instance. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name template of the instance. In order to be unique it must contain at least one of instance unique placeholders: {instance.short_id} {instance.index} combination of {instance.zone_id} and {instance.index_in_zone} Example: my-instance-{instance.index} If not set, default is used: {instance_group.id}-{instance.short_id} It may also contain another placeholders, see metadata doc for full list.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) Hostname template for the instance. This field is used to generate the FQDN value of instance. The hostname must be unique within the network and region. If not specified, the hostname will be equal to id of the instance and FQDN will be ` + "`" + `<id>.auto.internal` + "`" + `. Otherwise FQDN will be ` + "`" + `<hostname>.<region_id>.internal` + "`" + `. In order to be unique it must contain at least on of instance unique placeholders: {instance.short_id} {instance.index} combination of {instance.zone_id} and {instance.index_in_zone} Example: my-instance-{instance.index} If not set, ` + "`" + `name` + "`" + ` value will be used It may also contain another placeholders, see metadata doc for full list. --- The ` + "`" + `secondary_disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) The access mode to the disk resource. By default a disk is attached in ` + "`" + `READ_WRITE` + "`" + ` mode.`,
				},
				resource.Attribute{
					Name:        "initialize_params",
					Description: `(Required) Parameters used for creating a disk alongside the instance. The structure is documented below. - - -`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(Optional) This value can be used to reference the device under ` + "`" + `/dev/disk/by-id/` + "`" + `. --- The ` + "`" + `initialize_params` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the boot disk.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) The size of the disk in GB.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The disk type.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) The disk image to initialize this disk from.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The snapshot to initialize this disk from. ~>`,
				},
				resource.Attribute{
					Name:        "preemptible",
					Description: `(Optional) Specifies if the instance is preemptible. Defaults to false. --- The ` + "`" + `network_interface` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) The ID of the network.`,
				},
				resource.Attribute{
					Name:        "subnet_ids",
					Description: `(Optional) The ID of the subnets to attach this interface to.`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `(Optional) A public address that can be used to access the internet over NAT.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) Security group ids for network interface. --- The ` + "`" + `resources` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Required) The memory size in GB.`,
				},
				resource.Attribute{
					Name:        "cores",
					Description: `(Required) The number of CPU cores for the instance. - - -`,
				},
				resource.Attribute{
					Name:        "core_fraction",
					Description: `(Optional) If provided, specifies baseline core performance as a percent. --- The ` + "`" + `boot_disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) The access mode to the disk resource. By default a disk is attached in ` + "`" + `READ_WRITE` + "`" + ` mode.`,
				},
				resource.Attribute{
					Name:        "initialize_params",
					Description: `(Required) Parameters for creating a disk alongside the instance. The structure is documented below. - - -`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(Optional) This value can be used to reference the device under ` + "`" + `/dev/disk/by-id/` + "`" + `. --- The ` + "`" + `initialize_params` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the boot disk.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) The size of the disk in GB.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The disk type.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) The disk image to initialize this disk from.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The snapshot to initialize this disk from. ~>`,
				},
				resource.Attribute{
					Name:        "max_unavailable",
					Description: `(Required) The maximum number of running instances that can be taken offline (stopped or deleted) at the same time during the update process.`,
				},
				resource.Attribute{
					Name:        "max_expansion",
					Description: `(Required) The maximum number of instances that can be temporarily allocated above the group's target size during the update process. - - -`,
				},
				resource.Attribute{
					Name:        "max_deleting",
					Description: `(Optional) The maximum number of instances that can be deleted at the same time.`,
				},
				resource.Attribute{
					Name:        "max_creating",
					Description: `(Optional) The maximum number of instances that can be created at the same time.`,
				},
				resource.Attribute{
					Name:        "startup_duration",
					Description: `(Optional) The amount of time in seconds to allow for an instance to start. Instance will be considered up and running (and start receiving traffic) only after the startup_duration has elapsed and all health checks are passed. --- The ` + "`" + `scale_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "fixed_scale",
					Description: `(Optional) The fixed scaling policy of the instance group. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "auto_scale",
					Description: `(Optional) The auto scaling policy of the instance group. The structure is documented below. ~>`,
				},
				resource.Attribute{
					Name:        "test_auto_scale",
					Description: `(Optional) The test auto scaling policy of the instance group. Use it to test how the auto scale works. The structure is documented below. --- The ` + "`" + `fixed_scale` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The number of instances in the instance group. --- The ` + "`" + `auto_scale` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "initial_size",
					Description: `(Required) The initial number of instances in the instance group.`,
				},
				resource.Attribute{
					Name:        "measurement_duration",
					Description: `(Required) The amount of time, in seconds, that metrics are averaged for. If the average value at the end of the interval is higher than the ` + "`" + `cpu_utilization_target` + "`" + `, the instance group will increase the number of virtual machines in the group.`,
				},
				resource.Attribute{
					Name:        "cpu_utilization_target",
					Description: `(Required) Target CPU load level.`,
				},
				resource.Attribute{
					Name:        "min_zone_size",
					Description: `(Optional) The minimum number of virtual machines in a single availability zone.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `(Optional) The maximum number of virtual machines in the group.`,
				},
				resource.Attribute{
					Name:        "warmup_duration",
					Description: `(Optional) The warm-up time of the virtual machine, in seconds. During this time, traffic is fed to the virtual machine, but load metrics are not taken into account.`,
				},
				resource.Attribute{
					Name:        "stabilization_duration",
					Description: `(Optional) The minimum time interval, in seconds, to monitor the load before an instance group can reduce the number of virtual machines in the group. During this time, the group will not decrease even if the average load falls below the value of ` + "`" + `cpu_utilization_target` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "custom_rule",
					Description: `(Optional) A list of custom rules. The structure is documented below. --- The ` + "`" + `test_auto_scale` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "initial_size",
					Description: `(Required) The initial number of instances in the instance group.`,
				},
				resource.Attribute{
					Name:        "measurement_duration",
					Description: `(Required) The amount of time, in seconds, that metrics are averaged for. If the average value at the end of the interval is higher than the ` + "`" + `cpu_utilization_target` + "`" + `, the instance group will increase the number of virtual machines in the group.`,
				},
				resource.Attribute{
					Name:        "cpu_utilization_target",
					Description: `(Required) Target CPU load level.`,
				},
				resource.Attribute{
					Name:        "min_zone_size",
					Description: `(Optional) The minimum number of virtual machines in a single availability zone.`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `(Optional) The maximum number of virtual machines in the group.`,
				},
				resource.Attribute{
					Name:        "warmup_duration",
					Description: `(Optional) The warm-up time of the virtual machine, in seconds. During this time, traffic is fed to the virtual machine, but load metrics are not taken into account.`,
				},
				resource.Attribute{
					Name:        "stabilization_duration",
					Description: `(Optional) The minimum time interval, in seconds, to monitor the load before an instance group can reduce the number of virtual machines in the group. During this time, the group will not decrease even if the average load falls below the value of ` + "`" + `cpu_utilization_target` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "custom_rule",
					Description: `(Optional) A list of custom rules. The structure is documented below. --- The ` + "`" + `custom_rule` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "rule_type",
					Description: `(Required) Rule type: ` + "`" + `UTILIZATION` + "`" + ` - This type means that the metric applies to one instance. First, Instance Groups calculates the average metric value for each instance, then averages the values for instances in one availability zone. This type of metric must have the ` + "`" + `instance_id` + "`" + ` label. ` + "`" + `WORKLOAD` + "`" + ` - This type means that the metric applies to instances in one availability zone. This type of metric must have the ` + "`" + `zone_id` + "`" + ` label.`,
				},
				resource.Attribute{
					Name:        "metric_type",
					Description: `(Required) Metric type, ` + "`" + `GAUGE` + "`" + ` or ` + "`" + `COUNTER` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "metric_name",
					Description: `(Required) The name of metric.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) Target metric value level.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A map of labels of metric. --- The ` + "`" + `network_settings` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Network acceleration type. By default a network is in ` + "`" + `STANDARD` + "`" + ` mode. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the instance group.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The instance group creation timestamp.`,
				},
				resource.Attribute{
					Name:        "load_balancer.0.target_group_id",
					Description: `The ID of the target group.`,
				},
				resource.Attribute{
					Name:        "load_balancer.0.status_message",
					Description: `The status message of the target group. The ` + "`" + `instances` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the managed instance.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The Fully Qualified Domain Name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the instance.`,
				},
				resource.Attribute{
					Name:        "status_message",
					Description: `The status message of the instance.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `The ID of the availability zone where the instance resides.`,
				},
				resource.Attribute{
					Name:        "network_interface",
					Description: `An array with the network interfaces attached to the managed instance. --- The ` + "`" + `network_interface` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `The index of the network interface as generated by the server.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The MAC address assigned to the network interface.`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `True if IPv4 address allocated for the network interface.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The private IP address to assign to the instance. If empty, the address is automatically assigned from the specified subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the subnet to attach this interface to. The subnet must reside in the same zone where this instance was created.`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `The instance's public address for accessing the internet over NAT.`,
				},
				resource.Attribute{
					Name:        "nat_ip_address",
					Description: `The public IP address of the instance.`,
				},
				resource.Attribute{
					Name:        "nat_ip_version",
					Description: `The IP version for the public address.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the instance group.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The instance group creation timestamp.`,
				},
				resource.Attribute{
					Name:        "load_balancer.0.target_group_id",
					Description: `The ID of the target group.`,
				},
				resource.Attribute{
					Name:        "load_balancer.0.status_message",
					Description: `The status message of the target group. The ` + "`" + `instances` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The ID of the instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the managed instance.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The Fully Qualified Domain Name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the instance.`,
				},
				resource.Attribute{
					Name:        "status_message",
					Description: `The status message of the instance.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `The ID of the availability zone where the instance resides.`,
				},
				resource.Attribute{
					Name:        "network_interface",
					Description: `An array with the network interfaces attached to the managed instance. --- The ` + "`" + `network_interface` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "index",
					Description: `The index of the network interface as generated by the server.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `The MAC address assigned to the network interface.`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `True if IPv4 address allocated for the network interface.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The private IP address to assign to the instance. If empty, the address is automatically assigned from the specified subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the subnet to attach this interface to. The subnet must reside in the same zone where this instance was created.`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `The instance's public address for accessing the internet over NAT.`,
				},
				resource.Attribute{
					Name:        "nat_ip_address",
					Description: `The public IP address of the instance.`,
				},
				resource.Attribute{
					Name:        "nat_ip_version",
					Description: `The IP version for the public address.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_compute_snapshot",
			Category:         "Yandex Compute Service Resources",
			ShortDescription: `Creates a new snapshot of a disk.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"service",
				"snapshot",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_disk_id",
					Description: `(Required) ID of the disk to create a snapshot from. - - -`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A name for the resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the resource.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs to assign to the snapshot. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `Size of the disk when the snapshot was created, specified in GB.`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `Size of the snapshot, specified in GB.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the snapshot. ## Import A snapshot can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_compute_snapshot.disk-snapshot shapshot_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "disk_size",
					Description: `Size of the disk when the snapshot was created, specified in GB.`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `Size of the snapshot, specified in GB.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the snapshot. ## Import A snapshot can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_compute_snapshot.disk-snapshot shapshot_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_cr_container_registry",
			Category:         "Yandex Container Registry Resources",
			ShortDescription: `Creates a new container registry.`,
			Description:      ``,
			Keywords: []string{
				"container",
				"registry",
				"cr",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) Folder that the resource belongs to. If value is omitted, the default provider folder is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A name of the registry.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs to assign to the registry. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the registry.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the registry. ## Import A registry can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_container_registry.default registry_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `Status of the registry.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the registry. ## Import A registry can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_container_registry.default registry_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_dataproc_cluster",
			Category:         "Yandex Data Proc Resources",
			ShortDescription: `Manages a Data Proc cluster within Yandex.Cloud.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"proc",
				"dataproc",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of a specific Data Proc cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_config",
					Description: `(Required) Configuration and resources for hosts that should be created with the cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "service_account_id",
					Description: `(Required) Service account to be used by the Data Proc agent to access resources of Yandex.Cloud. Selected service account should have ` + "`" + `mdb.dataproc.agent` + "`" + ` role on the folder where the Data Proc cluster will be located. ---`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) ID of the folder to create a cluster in. If it is not provided, the default provider folder is used.`,
				},
				resource.Attribute{
					Name:        "bucket",
					Description: `(Optional) Name of the Object Storage bucket to use for Data Proc jobs. Data Proc Agent saves output of job driver's process to specified bucket. In order for this to work service account (specified by the ` + "`" + `service_account_id` + "`" + ` argument) should be given permission to create objects within this bucket.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Data Proc cluster.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs to assign to the Data Proc cluster.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Optional) ID of the availability zone to create cluster in. If it is not provided, the default provider zone is used. --- The ` + "`" + `cluster_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `(Required) Version of Data Proc image.`,
				},
				resource.Attribute{
					Name:        "hadoop",
					Description: `(Optional) Data Proc specific options. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "subcluster_spec",
					Description: `(Required) Configuration of the Data Proc subcluster. The structure is documented below. --- The ` + "`" + `hadoop` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Optional) List of services to run on Data Proc cluster.`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `(Optional) A set of key/value pairs that are used to configure cluster services.`,
				},
				resource.Attribute{
					Name:        "ssh_public_keys",
					Description: `(Optional) List of SSH public keys to put to the hosts of the cluster. For information on how to connect to the cluster, see [the official documentation](https://cloud.yandex.com/docs/data-proc/operations/connect). --- The ` + "`" + `subcluster_spec` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Data Proc subcluster.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) Role of the subcluster in the Data Proc cluster.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) Resources allocated to each host of the Data Proc subcluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) The ID of the subnet, to which hosts of the subcluster belong. Subnets of all the subclusters must belong to the same VPC network.`,
				},
				resource.Attribute{
					Name:        "hosts_count",
					Description: `(Required) Number of hosts within Data Proc subcluster. --- The ` + "`" + `resources` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resource_preset_id",
					Description: `(Required) The ID of the preset for computational resources available to a host. All available presets are listed in the [documentation](https://cloud.yandex.com/docs/data-proc/concepts/instance-types).`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Required) Volume of the storage available to a host, in gigabytes.`,
				},
				resource.Attribute{
					Name:        "disk_type_id",
					Description: `(Optional) Type of the storage of a host. One of ` + "`" + `network-hdd` + "`" + ` (default) or ` + "`" + `network-ssd` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) ID of a new Data Proc cluster.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `(Computed) The Data Proc cluster creation timestamp.`,
				},
				resource.Attribute{
					Name:        "cluster_config.0.subcluster_spec.X.id",
					Description: `(Computed) ID of the subcluster. ## Import A cluster can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_dataproc_cluster.foo cluster_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) ID of a new Data Proc cluster.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `(Computed) The Data Proc cluster creation timestamp.`,
				},
				resource.Attribute{
					Name:        "cluster_config.0.subcluster_spec.X.id",
					Description: `(Computed) ID of the subcluster. ## Import A cluster can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_dataproc_cluster.foo cluster_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_function",
			Category:         "Yandex Cloud Functions Resources",
			ShortDescription: `Allows management of a Yandex Cloud Function.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"functions",
				"function",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "folder_id",
					Description: `Folder ID for the Yandex Cloud Function`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the Yandex Cloud Function`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs to assign to the Yandex Cloud Function`,
				},
				resource.Attribute{
					Name:        "user_hash",
					Description: `User-defined string for current function version. User must change this string any times when function changed. Function will be updated when hash is changed.`,
				},
				resource.Attribute{
					Name:        "runtime",
					Description: `Runtime for Yandex Cloud Function`,
				},
				resource.Attribute{
					Name:        "entrypoint",
					Description: `Entrypoint for Yandex Cloud Function`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Memory in megabytes (`,
				},
				resource.Attribute{
					Name:        "execution_timeout",
					Description: `Execution timeout in seconds for Yandex Cloud Function`,
				},
				resource.Attribute{
					Name:        "service_account_id",
					Description: `Service account ID for Yandex Cloud Function`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `A set of key/value environment variables for Yandex Cloud Function`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags for Yandex Cloud Function. Tag "$latest" isn't returned.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version for Yandex Cloud Function.`,
				},
				resource.Attribute{
					Name:        "image_size",
					Description: `Image size for Yandex Cloud Function.`,
				},
				resource.Attribute{
					Name:        "loggroup_id",
					Description: `Loggroup ID size for Yandex Cloud Function.`,
				},
				resource.Attribute{
					Name:        "package",
					Description: `Version deployment package for Yandex Cloud Function code. Can be only one ` + "`" + `package` + "`" + ` or ` + "`" + `content` + "`" + ` section.`,
				},
				resource.Attribute{
					Name:        "package.0.sha_256",
					Description: `SHA256 hash of the version deployment package.`,
				},
				resource.Attribute{
					Name:        "package.0.bucket_name",
					Description: `Name of the bucket that stores the code for the version.`,
				},
				resource.Attribute{
					Name:        "package.0.object_name",
					Description: `Name of the object in the bucket that stores the code for the version.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `Version deployment content for Yandex Cloud Function code. Can be only one ` + "`" + `package` + "`" + ` or ` + "`" + `content` + "`" + ` section.`,
				},
				resource.Attribute{
					Name:        "content.0.zip_filename",
					Description: `Filename to zip archive for the version. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the Yandex Cloud Function.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version for Yandex Cloud Function.`,
				},
				resource.Attribute{
					Name:        "image_size",
					Description: `Image size for Yandex Cloud Function.`,
				},
				resource.Attribute{
					Name:        "loggroup_id",
					Description: `Log group ID size for Yandex Cloud Function.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the Yandex Cloud Function.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version for Yandex Cloud Function.`,
				},
				resource.Attribute{
					Name:        "image_size",
					Description: `Image size for Yandex Cloud Function.`,
				},
				resource.Attribute{
					Name:        "loggroup_id",
					Description: `Log group ID size for Yandex Cloud Function.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_function_iam_binding",
			Category:         "Yandex Cloud Functions Resources",
			ShortDescription: `Allows management of a single IAM binding for a [Yandex Cloud Function](https://cloud.yandex.com/docs/functions/).`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"functions",
				"function",
				"iam",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "function_id",
					Description: `(Required) The [Yandex Cloud Function](https://cloud.yandex.com/docs/functions/) ID to apply a binding to.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. See [roles](https://cloud.yandex.com/docs/functions/security/)`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Required) Identities that will be granted the privilege in ` + "`" + `role` + "`" + `. Each entry can have one of the following values:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_function_trigger",
			Category:         "Yandex Cloud Functions Resources",
			ShortDescription: `Allows management of a Yandex Cloud Functions Trigger.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"functions",
				"function",
				"trigger",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) Folder ID for the Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs to assign to the Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "function",
					Description: `[Yandex.Cloud Function](https://cloud.yandex.com/docs/functions/concepts/function) settings definition for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "function.0.id",
					Description: `Yandex.Cloud Function ID for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "function.0.service_account_id",
					Description: `Service account ID for Yandex.Cloud Function for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "function.0.tag",
					Description: `Tag for Yandex.Cloud Function for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "function.0.retry_attempts",
					Description: `Retry attempts for Yandex.Cloud Function for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "function.0.retry_interval",
					Description: `Retry interval in seconds for Yandex.Cloud Function for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "dlq",
					Description: `Dead Letter Queue settings definition for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "dlq.0.queue_id",
					Description: `Queue ID for Dead Letter Queue for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "dlq.0.service_account_id",
					Description: `Service Account ID for Dead Letter Queue for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "iot",
					Description: `[IoT](https://cloud.yandex.com/docs/functions/concepts/trigger/iot-core-trigger) settings definition for Yandex Cloud Functions Trigger, if present. Only one section ` + "`" + `iot` + "`" + ` or ` + "`" + `message_queue` + "`" + ` or ` + "`" + `object_storage` + "`" + ` or ` + "`" + `timer` + "`" + ` can be defined.`,
				},
				resource.Attribute{
					Name:        "iot.0.registry_id",
					Description: `IoT Registry ID for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "iot.0.device_id",
					Description: `IoT Device ID for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "iot.0.topic",
					Description: `IoT Topic for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "message_queue",
					Description: `[Message Queue](https://cloud.yandex.com/docs/functions/concepts/trigger/ymq-trigger) settings definition for Yandex Cloud Functions Trigger, if present`,
				},
				resource.Attribute{
					Name:        "message_queue.0.queue_id",
					Description: `Message Queue ID for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "message_queue.0.service_account_id",
					Description: `Message Queue Service Account ID for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "message_queue.0.batch_cutoff",
					Description: `Batch Duration in seconds for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "message_queue.0.batch_size",
					Description: `Batch Size for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "message_queue.0.visibility_timeout",
					Description: `Visibility timeout for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "object_storage",
					Description: `[Object Storage](https://cloud.yandex.com/docs/functions/concepts/trigger/os-trigger) settings definition for Yandex Cloud Functions Trigger, if present`,
				},
				resource.Attribute{
					Name:        "object_storage.0.bucket_id",
					Description: `Object Storage Bucket ID for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "object_storage.0.prefix",
					Description: `Prefix for Object Storage for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "object_storage.0.suffix",
					Description: `Suffix for Object Storage for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "object_storage.0.create",
					Description: `Boolean flag for setting create event for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "object_storage.0.update",
					Description: `Boolean flag for setting update event for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "object_storage.0.delete",
					Description: `Boolean flag for setting delete event for Yandex Cloud Functions Trigger`,
				},
				resource.Attribute{
					Name:        "timer",
					Description: `[Timer](https://cloud.yandex.com/docs/functions/concepts/trigger/timer) settings definition for Yandex Cloud Functions Trigger, if present`,
				},
				resource.Attribute{
					Name:        "timer.0.cron_expression",
					Description: `Cron expression for timer for Yandex Cloud Functions Trigger ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the Yandex Cloud Functions Trigger`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the Yandex Cloud Functions Trigger`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_iam_service_account",
			Category:         "Yandex Identity and Access Management Resources",
			ShortDescription: `Allows management of a Yandex.Cloud IAM service account.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"and",
				"access",
				"management",
				"iam",
				"service",
				"account",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the service account. Can be updated without creating a new resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the service account.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) ID of the folder that the service account will be created in. Defaults to the provider folder configuration. ## Import A service account can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_iam_service_account.sa account_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_iam_service_account_api_key",
			Category:         "Yandex Identity and Access Management Resources",
			ShortDescription: `Allows management of a Yandex.Cloud IAM service account API key.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"and",
				"access",
				"management",
				"iam",
				"service",
				"account",
				"api",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_account_id",
					Description: `(Required) ID of the service account to an API key for. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the key.`,
				},
				resource.Attribute{
					Name:        "pgp_key",
					Description: `(Optional) An optional PGP key to encrypt the resulting secret key material. May either be a base64-encoded public key or a keybase username in the form ` + "`" + `keybase:keybaseusername` + "`" + `. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `The secret key. This is only populated when no ` + "`" + `pgp_key` + "`" + ` is provided.`,
				},
				resource.Attribute{
					Name:        "encrypted_secret_key",
					Description: `The encrypted secret key, base64 encoded. This is only populated when ` + "`" + `pgp_key` + "`" + ` is supplied.`,
				},
				resource.Attribute{
					Name:        "key_fingerprint",
					Description: `The fingerprint of the PGP key used to encrypt the secret key. This is only populated when ` + "`" + `pgp_key` + "`" + ` is supplied.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the static access key.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "secret_key",
					Description: `The secret key. This is only populated when no ` + "`" + `pgp_key` + "`" + ` is provided.`,
				},
				resource.Attribute{
					Name:        "encrypted_secret_key",
					Description: `The encrypted secret key, base64 encoded. This is only populated when ` + "`" + `pgp_key` + "`" + ` is supplied.`,
				},
				resource.Attribute{
					Name:        "key_fingerprint",
					Description: `The fingerprint of the PGP key used to encrypt the secret key. This is only populated when ` + "`" + `pgp_key` + "`" + ` is supplied.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the static access key.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_iam_service_account_iam_binding",
			Category:         "Yandex Identity and Access Management Resources",
			ShortDescription: `Allows management of a single IAM binding for a Yandex IAM service account.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"and",
				"access",
				"management",
				"iam",
				"service",
				"account",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_account_id",
					Description: `(Required) The service account ID to apply a binding to.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. Only one ` + "`" + `yandex_iam_service_account_iam_binding` + "`" + ` can be used per role.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Required) Identities that will be granted the privilege in ` + "`" + `role` + "`" + `. Each entry can have one of the following values:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_iam_service_account_iam_member",
			Category:         "Yandex Identity and Access Management Resources",
			ShortDescription: `Allows management of a single member for a single IAM binding for a Yandex IAM service account.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"and",
				"access",
				"management",
				"iam",
				"service",
				"account",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_account_id",
					Description: `(Required) The service account ID to apply a policy to.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. Only one ` + "`" + `yandex_iam_service_account_iam_binding` + "`" + ` can be used per role.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) Identity that will be granted the privilege in ` + "`" + `role` + "`" + `. Entry can have one of the following values:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_iam_service_account_iam_policy",
			Category:         "Yandex Identity and Access Management Resources",
			ShortDescription: `Allows management of the IAM policy for a Yandex IAM service account.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"and",
				"access",
				"management",
				"iam",
				"service",
				"account",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_account_id",
					Description: `(Required) The service account ID to apply a policy to.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Required) Identities that will be granted the privilege in ` + "`" + `role` + "`" + `. Each entry can have one of the following values:`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. Only one ` + "`" + `yandex_iam_service_account_iam_binding` + "`" + ` can be used per role.`,
				},
				resource.Attribute{
					Name:        "policy_data",
					Description: `(Required only by ` + "`" + `yandex_iam_service_account_iam_policy` + "`" + `) The policy data generated by a ` + "`" + `yandex_iam_policy` + "`" + ` data source. ## Import Service account IAM policy resources can be imported using the service account ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_iam_service_account_iam_policy.admin-account-iam service_account_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_iam_service_account_key",
			Category:         "Yandex Identity and Access Management Resources",
			ShortDescription: `Allows management of a Yandex.Cloud IAM service account key.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"and",
				"access",
				"management",
				"iam",
				"service",
				"account",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_account_id",
					Description: `(Required) ID of the service account to create a pair for. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the key pair.`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Optional) The output format of the keys. ` + "`" + `PEM_FILE` + "`" + ` is the default format.`,
				},
				resource.Attribute{
					Name:        "key_algorithm",
					Description: `(Optional) The algorithm used to generate the key. ` + "`" + `RSA_2048` + "`" + ` is the default algorithm. Valid values are listed in the [API reference](https://cloud.yandex.com/docs/iam/api-ref/Key).`,
				},
				resource.Attribute{
					Name:        "pgp_key",
					Description: `(Optional) An optional PGP key to encrypt the resulting private key material. May either be a base64-encoded public key or a keybase username in the form ` + "`" + `keybase:keybaseusername` + "`" + `. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `The public key.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The private key. This is only populated when no ` + "`" + `pgp_key` + "`" + ` is provided.`,
				},
				resource.Attribute{
					Name:        "encrypted_private_key",
					Description: `The encrypted private key, base64 encoded. This is only populated when ` + "`" + `pgp_key` + "`" + ` is supplied.`,
				},
				resource.Attribute{
					Name:        "key_fingerprint",
					Description: `The fingerprint of the PGP key used to encrypt the private key. This is only populated when ` + "`" + `pgp_key` + "`" + ` is supplied.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the static access key.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "public_key",
					Description: `The public key.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The private key. This is only populated when no ` + "`" + `pgp_key` + "`" + ` is provided.`,
				},
				resource.Attribute{
					Name:        "encrypted_private_key",
					Description: `The encrypted private key, base64 encoded. This is only populated when ` + "`" + `pgp_key` + "`" + ` is supplied.`,
				},
				resource.Attribute{
					Name:        "key_fingerprint",
					Description: `The fingerprint of the PGP key used to encrypt the private key. This is only populated when ` + "`" + `pgp_key` + "`" + ` is supplied.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the static access key.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_iam_service_account_static_access_key",
			Category:         "Yandex Identity and Access Management Resources",
			ShortDescription: `Allows management of a Yandex.Cloud IAM service account static access key.`,
			Description:      ``,
			Keywords: []string{
				"identity",
				"and",
				"access",
				"management",
				"iam",
				"service",
				"account",
				"static",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_account_id",
					Description: `(Required) ID of the service account which is used to get a static key. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the service account static key.`,
				},
				resource.Attribute{
					Name:        "pgp_key",
					Description: `(Optional) An optional PGP key to encrypt the resulting secret key material. May either be a base64-encoded public key or a keybase username in the form ` + "`" + `keybase:keybaseusername` + "`" + `. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `ID of the static access key.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `Private part of generated static access key. This is only populated when no ` + "`" + `pgp_key` + "`" + ` is provided.`,
				},
				resource.Attribute{
					Name:        "encrypted_secret_key",
					Description: `The encrypted secret, base64 encoded. This is only populated when ` + "`" + `pgp_key` + "`" + ` is supplied.`,
				},
				resource.Attribute{
					Name:        "key_fingerprint",
					Description: `The fingerprint of the PGP key used to encrypt the secret key. This is only populated when ` + "`" + `pgp_key` + "`" + ` is supplied.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the static access key.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "access_key",
					Description: `ID of the static access key.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `Private part of generated static access key. This is only populated when no ` + "`" + `pgp_key` + "`" + ` is provided.`,
				},
				resource.Attribute{
					Name:        "encrypted_secret_key",
					Description: `The encrypted secret, base64 encoded. This is only populated when ` + "`" + `pgp_key` + "`" + ` is supplied.`,
				},
				resource.Attribute{
					Name:        "key_fingerprint",
					Description: `The fingerprint of the PGP key used to encrypt the secret key. This is only populated when ` + "`" + `pgp_key` + "`" + ` is supplied.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the static access key.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_iot_core_device",
			Category:         "Yandex IoT Core Resources",
			ShortDescription: `Allows management of a Yandex.Cloud IoT Core Device.`,
			Description:      ``,
			Keywords: []string{
				"iot",
				"core",
				"device",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "registry_id",
					Description: `IoT Core Registry ID for the IoT Core Device`,
				},
				resource.Attribute{
					Name:        "aliases",
					Description: `A set of key/value aliases pairs to assign to the IoT Core Device`,
				},
				resource.Attribute{
					Name:        "certificates",
					Description: `A set of certificate's fingerprints for the IoT Core Device`,
				},
				resource.Attribute{
					Name:        "passwords",
					Description: `A set of passwords's id for the IoT Core Device ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the IoT Core Device`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the IoT Core Device`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_iot_core_registry",
			Category:         "Yandex IoT Core Resources",
			ShortDescription: `Allows management of a Yandex.Cloud IoT Core Registry.`,
			Description:      ``,
			Keywords: []string{
				"iot",
				"core",
				"registry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "labels",
					Description: `A set of key/value label pairs to assign to the IoT Core Registry.`,
				},
				resource.Attribute{
					Name:        "certificates",
					Description: `A set of certificate's fingerprints for the IoT Core Registry`,
				},
				resource.Attribute{
					Name:        "passwords",
					Description: `A set of passwords's id for the IoT Core Registry ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `Folder ID for the IoT Core Registry`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the IoT Core Registry`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "folder_id",
					Description: `Folder ID for the IoT Core Registry`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the IoT Core Registry`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_kms_secret_ciphertext",
			Category:         "Yandex Key Management Service Resources",
			ShortDescription: `Encrypts given plaintext with the specified Yandex KMS key and provides access to the ciphertext.`,
			Description:      ``,
			Keywords: []string{
				"key",
				"management",
				"service",
				"kms",
				"secret",
				"ciphertext",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key_id",
					Description: `(Required) ID of the symmetric KMS key to use for encryption.`,
				},
				resource.Attribute{
					Name:        "aad_context",
					Description: `(Optional) Additional authenticated data (AAD context), optional. If specified, this data will be required for decryption with the ` + "`" + `SymmetricDecryptRequest` + "`" + ``,
				},
				resource.Attribute{
					Name:        "plaintext",
					Description: `(Required) Plaintext to be encrypted. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format ` + "`" + `{{key_id}}/{{ciphertext}}` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ciphertext",
					Description: `Resulting ciphertext, encoded with "standard" base64 alphabet as defined in RFC 4648 section 4 ## Timeouts ` + "`" + `yandex_kms_secret_ciphertext` + "`" + ` provides the following configuration options for [timeouts](/docs/configuration/resources.html#timeouts): - ` + "`" + `create` + "`" + ` - Default 1 minute - ` + "`" + `delete` + "`" + ` - Default 1 minute`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format ` + "`" + `{{key_id}}/{{ciphertext}}` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ciphertext",
					Description: `Resulting ciphertext, encoded with "standard" base64 alphabet as defined in RFC 4648 section 4 ## Timeouts ` + "`" + `yandex_kms_secret_ciphertext` + "`" + ` provides the following configuration options for [timeouts](/docs/configuration/resources.html#timeouts): - ` + "`" + `create` + "`" + ` - Default 1 minute - ` + "`" + `delete` + "`" + ` - Default 1 minute`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_kms_symmetric_key",
			Category:         "Yandex Key Management Service Resources",
			ShortDescription: `Creates a Yandex KMS symmetric key that can be used for cryptographic operation.`,
			Description:      ``,
			Keywords: []string{
				"key",
				"management",
				"service",
				"kms",
				"symmetric",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the key.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of the key.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs to assign to the key.`,
				},
				resource.Attribute{
					Name:        "default_algorithm",
					Description: `(Optional) Encryption algorithm to be used with a new key version, generated with the next rotation. The default value is ` + "`" + `AES_128` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rotation_period",
					Description: `(Optional) Interval between automatic rotations. To disable automatic rotation, omit this parameter. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the key.`,
				},
				resource.Attribute{
					Name:        "rotated_at",
					Description: `Last rotation timestamp of the key.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the key. ## Timeouts ` + "`" + `yandex_kms_symmetric_key` + "`" + ` provides the following configuration options for [timeouts](/docs/configuration/resources.html#timeouts): - ` + "`" + `create` + "`" + ` - Default 1 minute - ` + "`" + `update` + "`" + ` - Default 1 minute - ` + "`" + `delete` + "`" + ` - Default 1 minute ## Import A KMS symmetric key can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_kms_symmetric_key.top-secret kms_key_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `The status of the key.`,
				},
				resource.Attribute{
					Name:        "rotated_at",
					Description: `Last rotation timestamp of the key.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the key. ## Timeouts ` + "`" + `yandex_kms_symmetric_key` + "`" + ` provides the following configuration options for [timeouts](/docs/configuration/resources.html#timeouts): - ` + "`" + `create` + "`" + ` - Default 1 minute - ` + "`" + `update` + "`" + ` - Default 1 minute - ` + "`" + `delete` + "`" + ` - Default 1 minute ## Import A KMS symmetric key can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_kms_symmetric_key.top-secret kms_key_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_kubernetes_cluster",
			Category:         "Yandex Managed Service for Kubernetes Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"managed",
				"service",
				"for",
				"kubernetes",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of a specific Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the folder that the Kubernetes cluster belongs to. If it is not provided, the default provider folder is used.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs to assign to the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) The ID of the cluster network.`,
				},
				resource.Attribute{
					Name:        "cluster_ipv4_range",
					Description: `(Optional) CIDR block. IP range for allocating pod addresses. It should not overlap with any subnet in the network the Kubernetes cluster located in. Static routes will be set up for this CIDR blocks in node subnets.`,
				},
				resource.Attribute{
					Name:        "node_ipv4_cidr_mask_size",
					Description: `(Optional) Size of the masks that are assigned to each node in the cluster. Effectively limits maximum number of pods for each node.`,
				},
				resource.Attribute{
					Name:        "service_ipv4_range",
					Description: `(Optional) CIDR block. IP range Kubernetes service Kubernetes cluster IP addresses will be allocated from. It should not overlap with any subnet in the network the Kubernetes cluster located in.`,
				},
				resource.Attribute{
					Name:        "service_account_id",
					Description: `Service account to be used for provisioning Compute Cloud and VPC resources for Kubernetes cluster. Selected service account should have ` + "`" + `edit` + "`" + ` role on the folder where the Kubernetes cluster will be located and on the folder where selected network resides.`,
				},
				resource.Attribute{
					Name:        "node_service_account_id",
					Description: `Service account to be used by the worker nodes of the Kubernetes cluster to access Container Registry or to push node logs and metrics.`,
				},
				resource.Attribute{
					Name:        "release_channel",
					Description: `Cluster release channel.`,
				},
				resource.Attribute{
					Name:        "network_policy_provider",
					Description: `(Optional) Network policy provider for the cluster. Possible values: ` + "`" + `CALICO` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "kms_provider",
					Description: `(Optional) cluster KMS provider parameters.`,
				},
				resource.Attribute{
					Name:        "master",
					Description: `Kubernetes master configuration options. The structure is documented below. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Computed) ID of a new Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Computed)Status of the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `(Computed) Health of the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `(Computed) The Kubernetes cluster creation timestamp. --- The ` + "`" + `master` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) (Computed) Version of Kubernetes that will be used for master.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(Optional) (Computed) Boolean flag. When ` + "`" + `true` + "`" + `, Kubernetes master will have visible ipv4 address.`,
				},
				resource.Attribute{
					Name:        "maintenance_policy",
					Description: `(Optional) (Computed) Maintenance policy for Kubernetes master. If policy is omitted, automatic revision upgrades of the kubernetes master are enabled and could happen at any time. Revision upgrades are performed only within the same minor version, e.g. 1.13. Minor version upgrades (e.g. 1.13->1.14) should be performed manually. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "zonal",
					Description: `(Optional) Initialize parameters for Zonal Master (single node master). The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "regional",
					Description: `(Optional) Initialize parameters for Regional Master (highly available master). The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "version_info",
					Description: `(Computed) Information about cluster version. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "internal_v4_address",
					Description: `(Computed) An IPv4 internal network address that is assigned to the master.`,
				},
				resource.Attribute{
					Name:        "external_v4_address",
					Description: `(Computed) An IPv4 external network address that is assigned to the master.`,
				},
				resource.Attribute{
					Name:        "internal_v4_endpoint",
					Description: `(Computed) Internal endpoint that can be used to connect to the master from cloud networks.`,
				},
				resource.Attribute{
					Name:        "external_v4_endpoint",
					Description: `(Computed) External endpoint that can be used to access Kubernetes cluster API from the internet (outside of the cloud).`,
				},
				resource.Attribute{
					Name:        "cluster_ca_certificate",
					Description: `(Computed) PEM-encoded public certificate that is the root of trust for the Kubernetes cluster. --- The ` + "`" + `maintenance_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "auto_upgrade",
					Description: `(Required) Boolean flag that specifies if master can be upgraded automatically. When omitted, default value is TRUE.`,
				},
				resource.Attribute{
					Name:        "maintenance_window",
					Description: `(Optional) (Computed) This structure specifies maintenance window, when update for master is allowed. When omitted, it defaults to any time. To specify time of day interval, for all days, one element should be provided, with two fields set, ` + "`" + `start_time` + "`" + ` and ` + "`" + `duration` + "`" + `. Please see ` + "`" + `zonal_cluster_resource_name` + "`" + ` config example. To allow maintenance only on specific days of week, please provide list of elements, with all fields set. Only one time interval (` + "`" + `duration` + "`" + `) is allowed for each day of week. Please see ` + "`" + `regional_cluster_resource_name` + "`" + ` config example. --- The ` + "`" + `zonal` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) ID of the availability zone.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) ID of the subnet. If no ID is specified, and there only one subnet in specified zone, an address in this subnet will be allocated. --- The ` + "`" + `regional` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Name of availability region (e.g. "ru-central1"), where master instances will be allocated.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Array of locations, where master instances will be allocated. The structure is documented below. --- The ` + "`" + `location` + "`" + ` block supports repeated values:`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) ID of the availability zone.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) ID of the subnet. --- The ` + "`" + `version_info` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "current_version",
					Description: `Current Kubernetes version, major.minor (e.g. 1.15).`,
				},
				resource.Attribute{
					Name:        "new_revision_available",
					Description: `Boolean flag. Newer revisions may include Kubernetes patches (e.g 1.15.1 -> 1.15.2) as well as some internal component updates - new features or bug fixes in yandex-specific components either on the master or nodes.`,
				},
				resource.Attribute{
					Name:        "new_revision_summary",
					Description: `Human readable description of the changes to be applied when updating to the latest revision. Empty if new_revision_available is false.`,
				},
				resource.Attribute{
					Name:        "version_deprecated",
					Description: `Boolean flag. The current version is on the deprecation schedule, component (master or node group) should be upgraded. --- The ` + "`" + `kms_provider` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `KMS key ID. ## Timeouts This resource provides the following configuration options for [timeouts](/docs/configuration/resources.html#timeouts): - ` + "`" + `create` + "`" + ` - Default is 30 minute. - ` + "`" + `update` + "`" + ` - Default is 20 minute. - ` + "`" + `delete` + "`" + ` - Default is 20 minute. ## Import A Managed Kubernetes cluster can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_kubernetes_cluster.default cluster_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Computed) ID of a new Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Computed)Status of the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `(Computed) Health of the Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `(Computed) The Kubernetes cluster creation timestamp. --- The ` + "`" + `master` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) (Computed) Version of Kubernetes that will be used for master.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `(Optional) (Computed) Boolean flag. When ` + "`" + `true` + "`" + `, Kubernetes master will have visible ipv4 address.`,
				},
				resource.Attribute{
					Name:        "maintenance_policy",
					Description: `(Optional) (Computed) Maintenance policy for Kubernetes master. If policy is omitted, automatic revision upgrades of the kubernetes master are enabled and could happen at any time. Revision upgrades are performed only within the same minor version, e.g. 1.13. Minor version upgrades (e.g. 1.13->1.14) should be performed manually. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "zonal",
					Description: `(Optional) Initialize parameters for Zonal Master (single node master). The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "regional",
					Description: `(Optional) Initialize parameters for Regional Master (highly available master). The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "version_info",
					Description: `(Computed) Information about cluster version. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "internal_v4_address",
					Description: `(Computed) An IPv4 internal network address that is assigned to the master.`,
				},
				resource.Attribute{
					Name:        "external_v4_address",
					Description: `(Computed) An IPv4 external network address that is assigned to the master.`,
				},
				resource.Attribute{
					Name:        "internal_v4_endpoint",
					Description: `(Computed) Internal endpoint that can be used to connect to the master from cloud networks.`,
				},
				resource.Attribute{
					Name:        "external_v4_endpoint",
					Description: `(Computed) External endpoint that can be used to access Kubernetes cluster API from the internet (outside of the cloud).`,
				},
				resource.Attribute{
					Name:        "cluster_ca_certificate",
					Description: `(Computed) PEM-encoded public certificate that is the root of trust for the Kubernetes cluster. --- The ` + "`" + `maintenance_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "auto_upgrade",
					Description: `(Required) Boolean flag that specifies if master can be upgraded automatically. When omitted, default value is TRUE.`,
				},
				resource.Attribute{
					Name:        "maintenance_window",
					Description: `(Optional) (Computed) This structure specifies maintenance window, when update for master is allowed. When omitted, it defaults to any time. To specify time of day interval, for all days, one element should be provided, with two fields set, ` + "`" + `start_time` + "`" + ` and ` + "`" + `duration` + "`" + `. Please see ` + "`" + `zonal_cluster_resource_name` + "`" + ` config example. To allow maintenance only on specific days of week, please provide list of elements, with all fields set. Only one time interval (` + "`" + `duration` + "`" + `) is allowed for each day of week. Please see ` + "`" + `regional_cluster_resource_name` + "`" + ` config example. --- The ` + "`" + `zonal` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) ID of the availability zone.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) ID of the subnet. If no ID is specified, and there only one subnet in specified zone, an address in this subnet will be allocated. --- The ` + "`" + `regional` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Name of availability region (e.g. "ru-central1"), where master instances will be allocated.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Array of locations, where master instances will be allocated. The structure is documented below. --- The ` + "`" + `location` + "`" + ` block supports repeated values:`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) ID of the availability zone.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) ID of the subnet. --- The ` + "`" + `version_info` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "current_version",
					Description: `Current Kubernetes version, major.minor (e.g. 1.15).`,
				},
				resource.Attribute{
					Name:        "new_revision_available",
					Description: `Boolean flag. Newer revisions may include Kubernetes patches (e.g 1.15.1 -> 1.15.2) as well as some internal component updates - new features or bug fixes in yandex-specific components either on the master or nodes.`,
				},
				resource.Attribute{
					Name:        "new_revision_summary",
					Description: `Human readable description of the changes to be applied when updating to the latest revision. Empty if new_revision_available is false.`,
				},
				resource.Attribute{
					Name:        "version_deprecated",
					Description: `Boolean flag. The current version is on the deprecation schedule, component (master or node group) should be upgraded. --- The ` + "`" + `kms_provider` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `KMS key ID. ## Timeouts This resource provides the following configuration options for [timeouts](/docs/configuration/resources.html#timeouts): - ` + "`" + `create` + "`" + ` - Default is 30 minute. - ` + "`" + `update` + "`" + ` - Default is 20 minute. - ` + "`" + `delete` + "`" + ` - Default is 20 minute. ## Import A Managed Kubernetes cluster can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_kubernetes_cluster.default cluster_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_kubernetes_node_group",
			Category:         "Yandex Managed Service for Kubernetes Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"managed",
				"service",
				"for",
				"kubernetes",
				"node",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The ID of the Kubernetes cluster that this node group belongs to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of a specific Kubernetes node group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the Kubernetes node group.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs assigned to the Kubernetes node group.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Version of Kubernetes that will be used for Kubernetes node group.`,
				},
				resource.Attribute{
					Name:        "instance_template",
					Description: `(Required) Template used to create compute instances in this Kubernetes node group. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "scale_policy",
					Description: `(Required) Scale policy of the node group. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "allocation_policy",
					Description: `This argument specify subnets (zones), that will be used by node group compute instances. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "instance_group_id",
					Description: `(Computed) ID of instance group that is used to manage this Kubernetes node group.`,
				},
				resource.Attribute{
					Name:        "maintenance_policy",
					Description: `(Optional) (Computed) Maintenance policy for this Kubernetes node group. If policy is omitted, automatic revision upgrades are enabled and could happen at any time. Revision upgrades are performed only within the same minor version, e.g. 1.13. Minor version upgrades (e.g. 1.13->1.14) should be performed manually. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "node_labels",
					Description: `(Optional, Forces new resource) A set of key/value label pairs, that are assigned to all the nodes of this Kubernetes node group.`,
				},
				resource.Attribute{
					Name:        "node_taints",
					Description: `(Optional, Forces new resource) A list of Kubernetes taints, that are applied to all the nodes of this Kubernetes node group.`,
				},
				resource.Attribute{
					Name:        "allowed_unsafe_sysctls",
					Description: `(Optional, Forces new resource) A list of allowed unsafe sysctl parameters for this node group. For more details see [documentation](https://kubernetes.io/docs/tasks/administer-cluster/sysctl-cluster/).`,
				},
				resource.Attribute{
					Name:        "version_info",
					Description: `(Computed) Information about Kubernetes node group version. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "deploy_policy",
					Description: `Deploy policy of the node group. The structure is documented below. --- The ` + "`" + `instance_template` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "platform_id",
					Description: `The ID of the hardware platform configuration for the node group compute instances.`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `Boolean flag, enables NAT for node group compute instances.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The set of metadata ` + "`" + `key:value` + "`" + ` pairs assigned to this instance template. This includes custom metadata and predefined keys.`,
				},
				resource.Attribute{
					Name:        "resources.0.memory",
					Description: `The memory size allocated to the instance.`,
				},
				resource.Attribute{
					Name:        "resources.0.cores",
					Description: `Number of CPU cores allocated to the instance.`,
				},
				resource.Attribute{
					Name:        "resources.0.core_fraction",
					Description: `Baseline core performance as a percent.`,
				},
				resource.Attribute{
					Name:        "resources.0.gpus",
					Description: `Number of GPU cores allocated to the instance.`,
				},
				resource.Attribute{
					Name:        "boot_disk",
					Description: `The specifications for boot disks that will be attached to the instance. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "scheduling_policy",
					Description: `The scheduling policy for the instances in node group. The structure is documented below. --- The ` + "`" + `boot_disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the disk in GB. Allowed minimal size: 64 GB.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The disk type. --- The ` + "`" + `scheduling_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "preemptible",
					Description: `Specifies if the instance is preemptible. Defaults to false. --- The ` + "`" + `scale_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "fixed_scale",
					Description: `Scale policy for a fixed scale node group. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "auto_scale",
					Description: `Scale policy for an autoscaled node group. The structure is documented below. --- The ` + "`" + `fixed_scale` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The number of instances in the node group. --- The ` + "`" + `auto_scale` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `Minimum number of instances in the node group.`,
				},
				resource.Attribute{
					Name:        "max",
					Description: `Maximum number of instances in the node group.`,
				},
				resource.Attribute{
					Name:        "initial",
					Description: `Initial number of instances in the node group. --- The ` + "`" + `allocation_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Repeated field, that specify subnets (zones), that will be used by node group compute instances. The structure is documented below. --- The ` + "`" + `location` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `ID of the availability zone where for one compute instance in node group.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `ID of the subnet, that will be used by one compute instance in node group. Subnet specified by ` + "`" + `subnet_id` + "`" + ` should be allocated in zone specified by 'zone' argument --- The ` + "`" + `maintenance_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "auto_upgrade",
					Description: `(Required) Boolean flag that specifies if node group can be upgraded automatically. When omitted, default value is TRUE.`,
				},
				resource.Attribute{
					Name:        "maintenance_window",
					Description: `(Optional) (Computed) Set of day intervals, when maintenance is allowed for this node group. When omitted, it defaults to any time. To specify time of day interval, for all days, one element should be provided, with two fields set, ` + "`" + `start_time` + "`" + ` and ` + "`" + `duration` + "`" + `. To allow maintenance only on specific days of week, please provide list of elements, with all fields set. Only one time interval is allowed for each day of week. Please see ` + "`" + `my_node_group` + "`" + ` config example. --- The ` + "`" + `version_info` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "current_version",
					Description: `Current Kubernetes version, major.minor (e.g. 1.15).`,
				},
				resource.Attribute{
					Name:        "new_revision_available",
					Description: `True/false flag. Newer revisions may include Kubernetes patches (e.g 1.15.1 -> 1.15.2) as well as some internal component updates - new features or bug fixes in yandex-specific components either on the master or nodes.`,
				},
				resource.Attribute{
					Name:        "new_revision_summary",
					Description: `Human readable description of the changes to be applied when updating to the latest revision. Empty if new_revision_available is false.`,
				},
				resource.Attribute{
					Name:        "version_deprecated",
					Description: `True/false flag. The current version is on the deprecation schedule, component (master or node group) should be upgraded. --- The ` + "`" + `deploy_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "max_expansion",
					Description: `The maximum number of instances that can be temporarily allocated above the group's target size during the update.`,
				},
				resource.Attribute{
					Name:        "max_unavailable",
					Description: `The maximum number of running instances that can be taken offline during update. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Computed) Status of the Kubernetes node group.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `(Computed) The Kubernetes node group creation timestamp. ## Timeouts This resource provides the following configuration options for [timeouts](/docs/configuration/resources.html#timeouts): - ` + "`" + `create` + "`" + ` - Default is 60 minute. - ` + "`" + `update` + "`" + ` - Default is 60 minute. - ` + "`" + `delete` + "`" + ` - Default is 20 minute. ## Import A Yandex Kubernetes Node Group can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_kubernetes_node_group.default node_group_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `(Computed) Status of the Kubernetes node group.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `(Computed) The Kubernetes node group creation timestamp. ## Timeouts This resource provides the following configuration options for [timeouts](/docs/configuration/resources.html#timeouts): - ` + "`" + `create` + "`" + ` - Default is 60 minute. - ` + "`" + `update` + "`" + ` - Default is 60 minute. - ` + "`" + `delete` + "`" + ` - Default is 20 minute. ## Import A Yandex Kubernetes Node Group can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_kubernetes_node_group.default node_group_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_lb_network_load_balancer",
			Category:         "Yandex Load Balancer Resources",
			ShortDescription: `A network load balancer is used to evenly distribute the load across cloud resources.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"lb",
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the network load balancer. Provided by the client when the network load balancer is created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of the network load balancer. Provide this property when you create the resource.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the folder to which the resource belongs. If omitted, the provider folder is used.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Labels to assign to this network load balancer. A list of key/value pairs.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `(Optional) ID of the availability zone where the network load balancer resides. The default is 'ru-central1'.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of the network load balancer. Must be one of 'external' or 'internal'. The default is 'external'.`,
				},
				resource.Attribute{
					Name:        "attached_target_group",
					Description: `(Optional) An AttachedTargetGroup resource. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "listener",
					Description: `(Optional) Listener specification that will be used by a network load balancer. The structure is documented below. --- The ` + "`" + `attached_target_group` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "target_group_id",
					Description: `(Required) ID of the target group.`,
				},
				resource.Attribute{
					Name:        "healthcheck",
					Description: `(Required) A HealthCheck resource. The structure is documented below. --- The ` + "`" + `healthcheck` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the health check. The name must be unique for each target group that attached to a single load balancer.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Optional) The interval between health checks. The default is 2 seconds.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Timeout for a target to return a response for the health check. The default is 1 second.`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `(Optional) Number of failed health checks before changing the status to ` + "`" + `UNHEALTHY` + "`" + `. The default is 2.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `(Optional) Number of successful health checks required in order to set the ` + "`" + `HEALTHY` + "`" + ` status for the target.`,
				},
				resource.Attribute{
					Name:        "http_options",
					Description: `(Optional) Options for HTTP health check. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "tcp_options",
					Description: `(Optional) Options for TCP health check. The structure is documented below. ~>`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Port to use for HTTP health checks.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) URL path to set for health checking requests for every target in the target group. For example ` + "`" + `/ping` + "`" + `. The default path is ` + "`" + `/` + "`" + `. --- The ` + "`" + `tcp_options` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Port to use for TCP health checks. --- The ` + "`" + `listener` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the listener. The name must be unique for each listener on a single load balancer.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Port for incoming traffic.`,
				},
				resource.Attribute{
					Name:        "target_port",
					Description: `(Optional) Port of a target. The default is the same as listener's port.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Protocol for incoming traffic. TCP or UDP and the default is TCP.`,
				},
				resource.Attribute{
					Name:        "external_address_spec",
					Description: `(Optional) External IP address specification. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "internal_address_spec",
					Description: `(Optional) Internal IP address specification. The structure is documented below. ~>`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Optional) External IP address for a listener. IP address will be allocated if it wasn't been set.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) IP version of the external addresses that the load balancer works with. Must be one of ipv4 or ipv6. The default is ipv4. --- The ` + "`" + `internal_address_spec` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) ID of the subnet to which the internal IP address belongs.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Optional) Internal IP address for a listener. IP address will be allocated if it wasn't been set.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) IP version of the internal addresses that the load balancer works with. Must be one of ipv4 or ipv6. The default is ipv4. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the network load balancer.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The network load balancer creation timestamp. ## Timeouts This resource provides the following configuration options for [timeouts](/docs/configuration/resources.html#timeouts): - ` + "`" + `create` + "`" + ` - Default is 5 minute. - ` + "`" + `update` + "`" + ` - Default is 5 minute. - ` + "`" + `delete` + "`" + ` - Default is 5 minute. ## Import A network load balancer can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_lb_network_load_balancer.default network_load_balancer_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the network load balancer.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The network load balancer creation timestamp. ## Timeouts This resource provides the following configuration options for [timeouts](/docs/configuration/resources.html#timeouts): - ` + "`" + `create` + "`" + ` - Default is 5 minute. - ` + "`" + `update` + "`" + ` - Default is 5 minute. - ` + "`" + `delete` + "`" + ` - Default is 5 minute. ## Import A network load balancer can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_lb_network_load_balancer.default network_load_balancer_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_lb_target_group",
			Category:         "Yandex Load Balancer Resources",
			ShortDescription: `A load balancer distributes the load across cloud resources that are combined into a target group.`,
			Description:      ``,
			Keywords: []string{
				"load",
				"balancer",
				"lb",
				"target",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the target group. Provided by the client when the target group is created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of the target group. Provide this property when you create the resource.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the folder to which the resource belongs. If omitted, the provider folder is used.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Labels to assign to this target group. A list of key/value pairs.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `(Optional) ID of the availability zone where the target group resides. The default is 'ru-central1'.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Optional) A Target resource. The structure is documented below. --- The ` + "`" + `target` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) IP address of the target.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) ID of the subnet that targets are connected to. All targets in the target group must be connected to the same subnet within a single availability zone. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the target group.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The target group creation timestamp. ## Timeouts This resource provides the following configuration options for [timeouts](/docs/configuration/resources.html#timeouts): - ` + "`" + `create` + "`" + ` - Default is 5 minute. - ` + "`" + `update` + "`" + ` - Default is 5 minute. - ` + "`" + `delete` + "`" + ` - Default is 5 minute. ## Import A target group can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_lb_target_group.default target_group_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the target group.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The target group creation timestamp. ## Timeouts This resource provides the following configuration options for [timeouts](/docs/configuration/resources.html#timeouts): - ` + "`" + `create` + "`" + ` - Default is 5 minute. - ` + "`" + `update` + "`" + ` - Default is 5 minute. - ` + "`" + `delete` + "`" + ` - Default is 5 minute. ## Import A target group can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_lb_target_group.default target_group_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_mdb_clickhouse_cluster",
			Category:         "Yandex Managed Service for Database Resources",
			ShortDescription: `Manages a ClickHouse cluster within Yandex.Cloud.`,
			Description:      ``,
			Keywords: []string{
				"managed",
				"service",
				"for",
				"database",
				"mdb",
				"clickhouse",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the ClickHouse cluster. Provided by the client when the cluster is created.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) ID of the network, to which the ClickHouse cluster belongs.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `(Required) Deployment environment of the ClickHouse cluster. Can be either ` + "`" + `PRESTABLE` + "`" + ` or ` + "`" + `PRODUCTION` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "clickhouse",
					Description: `(Required) Configuration of the ClickHouse subcluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Required) A user of the ClickHouse cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required) A database of the ClickHouse cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required) A host of the ClickHouse cluster. The structure is documented below. - - -`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Version of the ClickHouse server software.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the ClickHouse cluster.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs to assign to the ClickHouse cluster.`,
				},
				resource.Attribute{
					Name:        "backup_window_start",
					Description: `(Optional) Time to start the daily backup, in the UTC timezone. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `(Optional) Access policy to the ClickHouse cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "zookeeper",
					Description: `(Optional) Configuration of the ZooKeeper subcluster. The structure is documented below. - - - The ` + "`" + `clickhouse` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) Resources allocated to hosts of the ClickHouse subcluster. The structure is documented below. The ` + "`" + `resources` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resources_preset_id",
					Description: `(Required) The ID of the preset for computational resources available to a ClickHouse host (CPU, memory etc.). For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-clickhouse/concepts).`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Required) Volume of the storage available to a ClickHouse host, in gigabytes.`,
				},
				resource.Attribute{
					Name:        "disk_type_id",
					Description: `(Required) Type of the storage of ClickHouse hosts. For more information see [the official documentation](https://cloud.yandex.com/docs/managed-clickhouse/concepts/storage). The ` + "`" + `zookeeper` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Optional) Resources allocated to hosts of the ZooKeeper subcluster. The structure is documented below. The ` + "`" + `resources` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resources_preset_id",
					Description: `(Optional) The ID of the preset for computational resources available to a ZooKeeper host (CPU, memory etc.). For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-clickhouse/concepts).`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional) Volume of the storage available to a ZooKeeper host, in gigabytes.`,
				},
				resource.Attribute{
					Name:        "disk_type_id",
					Description: `(Optional) Type of the storage of ZooKeeper hosts. For more information see [the official documentation](https://cloud.yandex.com/docs/managed-clickhouse/concepts/storage). The ` + "`" + `user` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The password of the user.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `(Optional) Set of permissions granted to the user. The structure is documented below. The ` + "`" + `permission` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `(Required) The name of the database that the permission grants access to. The ` + "`" + `database` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the database. The ` + "`" + `host` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `(Computed) The fully qualified domain name of the host.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the host to be deployed. Can be either ` + "`" + `CLICKHOUSE` + "`" + ` or ` + "`" + `ZOOKEEPER` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The availability zone where the ClickHouse host will be created. For more information see [the official documentation](https://cloud.yandex.com/docs/overview/concepts/geo-scope).`,
				},
				resource.Attribute{
					Name:        "hours",
					Description: `(Optional) The hour at which backup will be started.`,
				},
				resource.Attribute{
					Name:        "minutes",
					Description: `(Optional) The minute at which backup will be started. The ` + "`" + `access` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "web_sql",
					Description: `(Optional) Allow access for DataLens. Can be either ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "data_lens",
					Description: `(Optional) Allow access for Web SQL. Can be either ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "metrika",
					Description: `(Optional) Allow access for Yandex.Metrika. Can be either ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "serverless",
					Description: `(Optional) Allow access for Serverless. Can be either ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Timestamp of cluster creation.`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `Aggregated health of the cluster. Can be either ` + "`" + `ALIVE` + "`" + `, ` + "`" + `DEGRADED` + "`" + `, ` + "`" + `DEAD` + "`" + ` or ` + "`" + `HEALTH_UNKNOWN` + "`" + `. For more information see ` + "`" + `health` + "`" + ` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-clickhouse/api-ref/Cluster/).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cluster. Can be either ` + "`" + `CREATING` + "`" + `, ` + "`" + `STARTING` + "`" + `, ` + "`" + `RUNNING` + "`" + `, ` + "`" + `UPDATING` + "`" + `, ` + "`" + `STOPPING` + "`" + `, ` + "`" + `STOPPED` + "`" + `, ` + "`" + `ERROR` + "`" + ` or ` + "`" + `STATUS_UNKNOWN` + "`" + `. For more information see ` + "`" + `status` + "`" + ` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-clickhouse/api-ref/Cluster/). ## Import A cluster can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_mdb_clickhouse_cluster.foo cluster_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Timestamp of cluster creation.`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `Aggregated health of the cluster. Can be either ` + "`" + `ALIVE` + "`" + `, ` + "`" + `DEGRADED` + "`" + `, ` + "`" + `DEAD` + "`" + ` or ` + "`" + `HEALTH_UNKNOWN` + "`" + `. For more information see ` + "`" + `health` + "`" + ` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-clickhouse/api-ref/Cluster/).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cluster. Can be either ` + "`" + `CREATING` + "`" + `, ` + "`" + `STARTING` + "`" + `, ` + "`" + `RUNNING` + "`" + `, ` + "`" + `UPDATING` + "`" + `, ` + "`" + `STOPPING` + "`" + `, ` + "`" + `STOPPED` + "`" + `, ` + "`" + `ERROR` + "`" + ` or ` + "`" + `STATUS_UNKNOWN` + "`" + `. For more information see ` + "`" + `status` + "`" + ` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-clickhouse/api-ref/Cluster/). ## Import A cluster can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_mdb_clickhouse_cluster.foo cluster_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_mdb_mongodb_cluster",
			Category:         "Yandex Managed Service for Database Resources",
			ShortDescription: `Manages a MongoDB cluster within Yandex.Cloud.`,
			Description:      ``,
			Keywords: []string{
				"managed",
				"service",
				"for",
				"database",
				"mdb",
				"mongodb",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the MongoDB cluster. Provided by the client when the cluster is created.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) ID of the network, to which the MongoDB cluster belongs.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `(Required) Deployment environment of the MongoDB cluster. Can be either ` + "`" + `PRESTABLE` + "`" + ` or ` + "`" + `PRODUCTION` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cluster_config",
					Description: `(Required) Configuration of the MongoDB subcluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Required) A user of the MongoDB cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required) A database of the MongoDB cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required) A host of the MongoDB cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) Resources allocated to hosts of the MongoDB cluster. The structure is documented below. - - -`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Version of the MongoDB server software. Can be either ` + "`" + `3.6` + "`" + `, ` + "`" + `4.0` + "`" + ` and ` + "`" + `4.2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the MongoDB cluster.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs to assign to the MongoDB cluster.`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `(Optional) Access policy to the MongoDB cluster. The structure is documented below. - - - The ` + "`" + `cluster_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Version of MongoDB (either 4.2, 4.0 or 3.6).`,
				},
				resource.Attribute{
					Name:        "feature_compatibility_version",
					Description: `(Optional) Feature compatibility version of MongoDB. If not provided version is taken. Can be either ` + "`" + `4.2` + "`" + `, ` + "`" + `4.0` + "`" + ` and ` + "`" + `3.6` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "backup_window_start",
					Description: `(Optional) Time to start the daily backup, in the UTC timezone. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `(Optional) Shows whether cluster has access to data lens. The structure is documented below. The ` + "`" + `backup_window_start` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "hours",
					Description: `(Optional) The hour at which backup will be started.`,
				},
				resource.Attribute{
					Name:        "minutes",
					Description: `(Optional) The minute at which backup will be started. The ` + "`" + `resources` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resources_preset_id",
					Description: `(Required) The ID of the preset for computational resources available to a MongoDB host (CPU, memory etc.). For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-mongodb/concepts).`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Required) Volume of the storage available to a MongoDB host, in gigabytes.`,
				},
				resource.Attribute{
					Name:        "disk_type_id",
					Description: `(Required) Type of the storage of MongoDB hosts. For more information see [the official documentation](https://cloud.yandex.com/docs/managed-clickhouse/concepts/storage). The ` + "`" + `user` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The password of the user.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `(Optional) Set of permissions granted to the user. The structure is documented below. The ` + "`" + `permission` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `(Required) The name of the database that the permission grants access to.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) The role of the user in this database. For more information see [the official documentation](https://cloud.yandex.com/docs/managed-mongodb/concepts/users-and-roles). The ` + "`" + `database` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the database. The ` + "`" + `host` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Computed) The fully qualified domain name of the host. Computed on server side.`,
				},
				resource.Attribute{
					Name:        "zone_id",
					Description: `(Required) The availability zone where the MongoDB host will be created. For more information see [the official documentation](https://cloud.yandex.com/docs/overview/concepts/geo-scope).`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) The role of the cluster (either PRIMARY or SECONDARY).`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `(Computed) The health of the host.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) The ID of the subnet, to which the host belongs. The subnet must be a part of the network to which the cluster belongs.`,
				},
				resource.Attribute{
					Name:        "shard_name",
					Description: `(Optional) The name of the shard to which the host belongs.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) type of mongo daemon which runs on this host (mongod, mongos or monogcfg). Defaults to mongod. The ` + "`" + `access` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "data_lens",
					Description: `(Optional) Allow access for DataLens. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the key.`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `Aggregated health of the cluster. Can be either ` + "`" + `ALIVE` + "`" + `, ` + "`" + `DEGRADED` + "`" + `, ` + "`" + `DEAD` + "`" + ` or ` + "`" + `HEALTH_UNKNOWN` + "`" + `. For more information see ` + "`" + `health` + "`" + ` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-mongodb/api-ref/Cluster/).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cluster. Can be either ` + "`" + `CREATING` + "`" + `, ` + "`" + `STARTING` + "`" + `, ` + "`" + `RUNNING` + "`" + `, ` + "`" + `UPDATING` + "`" + `, ` + "`" + `STOPPING` + "`" + `, ` + "`" + `STOPPED` + "`" + `, ` + "`" + `ERROR` + "`" + ` or ` + "`" + `STATUS_UNKNOWN` + "`" + `. For more information see ` + "`" + `status` + "`" + ` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-mongodb/api-ref/Cluster/).`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `The ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "sharded",
					Description: `MongoDB Cluster mode enabled/disabled. ## Import A cluster can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_mdb_mongodb_cluster.foo cluster_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the key.`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `Aggregated health of the cluster. Can be either ` + "`" + `ALIVE` + "`" + `, ` + "`" + `DEGRADED` + "`" + `, ` + "`" + `DEAD` + "`" + ` or ` + "`" + `HEALTH_UNKNOWN` + "`" + `. For more information see ` + "`" + `health` + "`" + ` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-mongodb/api-ref/Cluster/).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cluster. Can be either ` + "`" + `CREATING` + "`" + `, ` + "`" + `STARTING` + "`" + `, ` + "`" + `RUNNING` + "`" + `, ` + "`" + `UPDATING` + "`" + `, ` + "`" + `STOPPING` + "`" + `, ` + "`" + `STOPPED` + "`" + `, ` + "`" + `ERROR` + "`" + ` or ` + "`" + `STATUS_UNKNOWN` + "`" + `. For more information see ` + "`" + `status` + "`" + ` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-mongodb/api-ref/Cluster/).`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `The ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "sharded",
					Description: `MongoDB Cluster mode enabled/disabled. ## Import A cluster can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_mdb_mongodb_cluster.foo cluster_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_mdb_mysql_cluster",
			Category:         "Yandex Managed Service for Database Resources",
			ShortDescription: `Manages a MySQL cluster within Yandex.Cloud.`,
			Description:      ``,
			Keywords: []string{
				"managed",
				"service",
				"for",
				"database",
				"mdb",
				"mysql",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the MySQL cluster. Provided by the client when the cluster is created.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) ID of the network, to which the MySQL cluster uses.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `(Required) Deployment environment of the MySQL cluster.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Version of the MySQL cluster.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) Resources allocated to hosts of the MySQL cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Required) A user of the MySQL cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required) A database of the MySQL cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required) A host of the MySQL cluster. The structure is documented below. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the MySQL cluster.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs to assign to the MySQL cluster.`,
				},
				resource.Attribute{
					Name:        "backup_window_start",
					Description: `(Optional) Time to start the daily backup, in the UTC. The structure is documented below. - - - The ` + "`" + `resources` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resources_preset_id",
					Description: `(Required) The ID of the preset for computational resources available to a MySQL host (CPU, memory etc.). For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-mysql/concepts/instance-types).`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Required) Volume of the storage available to a MySQL host, in gigabytes.`,
				},
				resource.Attribute{
					Name:        "disk_type_id",
					Description: `(Required) Type of the storage of MySQL hosts. The ` + "`" + `backup_window_start` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "hours",
					Description: `(Optional) The hour at which backup will be started.`,
				},
				resource.Attribute{
					Name:        "minutes",
					Description: `(Optional) The minute at which backup will be started. The ` + "`" + `user` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The password of the user.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `(Optional) Set of permissions granted to the user. The structure is documented below. The ` + "`" + `permission` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `(Required) The name of the database that the permission grants access to.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Optional) List user's roles in the database. Allowed roles: ` + "`" + `ALL` + "`" + `,` + "`" + `ALTER` + "`" + `,` + "`" + `ALTER_ROUTINE` + "`" + `,` + "`" + `CREATE` + "`" + `,` + "`" + `CREATE_ROUTINE` + "`" + `,` + "`" + `CREATE_TEMPORARY_TABLES` + "`" + `, ` + "`" + `CREATE_VIEW` + "`" + `,` + "`" + `DELETE` + "`" + `,` + "`" + `DROP` + "`" + `,` + "`" + `EVENT` + "`" + `,` + "`" + `EXECUTE` + "`" + `,` + "`" + `INDEX` + "`" + `,` + "`" + `INSERT` + "`" + `,` + "`" + `LOCK_TABLES` + "`" + `,` + "`" + `SELECT` + "`" + `,` + "`" + `SHOW_VIEW` + "`" + `,` + "`" + `TRIGGER` + "`" + `,` + "`" + `UPDATE` + "`" + `. The ` + "`" + `database` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the database. The ` + "`" + `host` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `(Computed) The fully qualified domain name of the host.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The availability zone where the MySQL host will be created.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) The ID of the subnet, to which the host belongs. The subnet must be a part of the network to which the cluster belongs.`,
				},
				resource.Attribute{
					Name:        "assign_public_ip",
					Description: `(Optional) Sets whether the host should get a public IP address on creation. Changing this parameter for an existing host is not supported at the moment ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the cluster.`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `Aggregated health of the cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cluster. ## Import A cluster can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_mdb_mysql_cluster.foo cluster_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the cluster.`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `Aggregated health of the cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cluster. ## Import A cluster can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_mdb_mysql_cluster.foo cluster_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_mdb_postgresql_cluster",
			Category:         "Yandex Managed Service for Database Resources",
			ShortDescription: `Manages a PostgreSQL cluster within Yandex.Cloud.`,
			Description:      ``,
			Keywords: []string{
				"managed",
				"service",
				"for",
				"database",
				"mdb",
				"postgresql",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config",
					Description: `(Required) Configuration of the PostgreSQL cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required) A database of the PostgreSQL cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `(Required) Deployment environment of the PostgreSQL cluster.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required) A host of the PostgreSQL cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the PostgreSQL cluster. Provided by the client when the cluster is created.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) ID of the network, to which the PostgreSQL cluster belongs.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `(Required) A user of the PostgreSQL cluster. The structure is documented below. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the PostgreSQL cluster.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs to assign to the PostgreSQL cluster. - - - The ` + "`" + `config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) Resources allocated to hosts of the PostgreSQL cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Version of the PostgreSQL cluster.`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `(Optional) Access policy to the PostgreSQL cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "autofailover",
					Description: `(Optional) Configuration setting which enables/disables autofailover in cluster.`,
				},
				resource.Attribute{
					Name:        "backup_window_start",
					Description: `(Optional) Time to start the daily backup, in the UTC timezone. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "pooler_config",
					Description: `(Optional) Configuration of the connection pooler. The structure is documented below. The ` + "`" + `resources` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Required) Volume of the storage available to a PostgreSQL host, in gigabytes.`,
				},
				resource.Attribute{
					Name:        "disk_type_id",
					Description: `(Required) Type of the storage of PostgreSQL hosts.`,
				},
				resource.Attribute{
					Name:        "resources_preset_id",
					Description: `(Required) The ID of the preset for computational resources available to a PostgreSQL host (CPU, memory etc.). For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-postgresql/concepts/instance-types). The ` + "`" + `pooler_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "pool_discard",
					Description: `(Optional) Setting ` + "`" + `server_reset_query_always` + "`" + ` [parameter in PgBouncer](https://www.pgbouncer.org/config.html).`,
				},
				resource.Attribute{
					Name:        "pooling_mode",
					Description: `(Optional) Mode that the connection pooler is working in. See descriptions of all modes in the [documentation for PgBouncer](https://pgbouncer.github.io/usage). The ` + "`" + `backup_window_start` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "hours",
					Description: `(Optional) The hour at which backup will be started (UTC).`,
				},
				resource.Attribute{
					Name:        "minutes",
					Description: `(Optional) The minute at which backup will be started (UTC). The ` + "`" + `access` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "data_lens",
					Description: `(Optional) Allow access for [Yandex DataLens](https://cloud.yandex.com/services/datalens). The ` + "`" + `user` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The password of the user.`,
				},
				resource.Attribute{
					Name:        "grants",
					Description: `(Optional) List of the user's grants.`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `(Optional) User's ability to login.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `(Optional) Set of permissions granted to the user. The structure is documented below. The ` + "`" + `permission` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `(Required) The name of the database that the permission grants access to. The ` + "`" + `database` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the database.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `(Required) Name of the user assigned as the owner of the database.`,
				},
				resource.Attribute{
					Name:        "extension",
					Description: `(Optional) Set of database extensions. The structure is documented below`,
				},
				resource.Attribute{
					Name:        "lc_collate",
					Description: `(Optional) POSIX locale for string sorting order. Forbidden to change in an existing database.`,
				},
				resource.Attribute{
					Name:        "lc_type",
					Description: `(Optional) POSIX locale for character classification. Forbidden to change in an existing database. The ` + "`" + `extension` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the database extension. For more information on available extensions see [the official documentation](https://cloud.yandex.com/docs/managed-postgresql/operations/cluster-extensions).`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Version of the extension. The ` + "`" + `host` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The availability zone where the PostgreSQL host will be created.`,
				},
				resource.Attribute{
					Name:        "assign_public_ip",
					Description: `(Optional) Sets whether the host should get a public IP address on creation. Changing this parameter for an existing host is not supported at the moment`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) The ID of the subnet, to which the host belongs. The subnet must be a part of the network to which the cluster belongs.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `(Computed) The fully qualified domain name of the host. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Timestamp of cluster creation.`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `Aggregated health of the cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cluster. ## Import A cluster can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_mdb_postgresql_cluster.foo cluster_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Timestamp of cluster creation.`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `Aggregated health of the cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cluster. ## Import A cluster can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_mdb_postgresql_cluster.foo cluster_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_mdb_redis_cluster",
			Category:         "Yandex Managed Service for Database Resources",
			ShortDescription: `Manages a Redis cluster within Yandex.Cloud.`,
			Description:      ``,
			Keywords: []string{
				"managed",
				"service",
				"for",
				"database",
				"mdb",
				"redis",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Redis cluster. Provided by the client when the cluster is created.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) ID of the network, to which the Redis cluster belongs.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `(Required) Deployment environment of the Redis cluster. Can be either ` + "`" + `PRESTABLE` + "`" + ` or ` + "`" + `PRODUCTION` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `(Required) Configuration of the Redis cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) Resources allocated to hosts of the Redis cluster. The structure is documented below.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required) A host of the Redis cluster. The structure is documented below. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Redis cluster.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs to assign to the Redis cluster.`,
				},
				resource.Attribute{
					Name:        "sharded",
					Description: `(Optional) Redis Cluster mode enabled/disabled. - - - The ` + "`" + `config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password for the Redis cluster.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Close the connection after a client is idle for N seconds.`,
				},
				resource.Attribute{
					Name:        "maxmemory_policy",
					Description: `(Optional) Redis key eviction policy for a dataset that reaches maximum memory. Can be any of the listed in [the official RedisDB documentation](https://docs.redislabs.com/latest/rs/administering/database-operations/eviction-policy/).`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Version of Redis (either 5.0 or 6.0). The ` + "`" + `resources` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resources_preset_id",
					Description: `(Required) The ID of the preset for computational resources available to a host (CPU, memory etc.). For more information, see [the official documentation](https://cloud.yandex.com/docs/managed-redis/concepts).`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Required) Volume of the storage available to a host, in gigabytes. The ` + "`" + `host` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The availability zone where the Redis host will be created. For more information see [the official documentation](https://cloud.yandex.com/docs/overview/concepts/geo-scope).`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the key.`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `Aggregated health of the cluster. Can be either ` + "`" + `ALIVE` + "`" + `, ` + "`" + `DEGRADED` + "`" + `, ` + "`" + `DEAD` + "`" + ` or ` + "`" + `HEALTH_UNKNOWN` + "`" + `. For more information see ` + "`" + `health` + "`" + ` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-redis/api-ref/Cluster/).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cluster. Can be either ` + "`" + `CREATING` + "`" + `, ` + "`" + `STARTING` + "`" + `, ` + "`" + `RUNNING` + "`" + `, ` + "`" + `UPDATING` + "`" + `, ` + "`" + `STOPPING` + "`" + `, ` + "`" + `STOPPED` + "`" + `, ` + "`" + `ERROR` + "`" + ` or ` + "`" + `STATUS_UNKNOWN` + "`" + `. For more information see ` + "`" + `status` + "`" + ` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-redis/api-ref/Cluster/). ## Import A cluster can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_mdb_redis_cluster.foo cluster_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the key.`,
				},
				resource.Attribute{
					Name:        "health",
					Description: `Aggregated health of the cluster. Can be either ` + "`" + `ALIVE` + "`" + `, ` + "`" + `DEGRADED` + "`" + `, ` + "`" + `DEAD` + "`" + ` or ` + "`" + `HEALTH_UNKNOWN` + "`" + `. For more information see ` + "`" + `health` + "`" + ` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-redis/api-ref/Cluster/).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cluster. Can be either ` + "`" + `CREATING` + "`" + `, ` + "`" + `STARTING` + "`" + `, ` + "`" + `RUNNING` + "`" + `, ` + "`" + `UPDATING` + "`" + `, ` + "`" + `STOPPING` + "`" + `, ` + "`" + `STOPPED` + "`" + `, ` + "`" + `ERROR` + "`" + ` or ` + "`" + `STATUS_UNKNOWN` + "`" + `. For more information see ` + "`" + `status` + "`" + ` field of JSON representation in [the official documentation](https://cloud.yandex.com/docs/managed-redis/api-ref/Cluster/). ## Import A cluster can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_mdb_redis_cluster.foo cluster_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_message_queue",
			Category:         "Yandex Message Queue Resources",
			ShortDescription: `Allows management of a Yandex.Cloud Message Queue.`,
			Description:      ``,
			Keywords: []string{
				"message",
				"queue",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, forces new resource) Queue name. The maximum length is 80 characters. You can use numbers, letters, underscores, and hyphens in the name. The name of a FIFO queue must end with the ` + "`" + `.fifo` + "`" + ` suffix. If not specified, random name will be generated. Conflicts with ` + "`" + `name_prefix` + "`" + `. For more information see [documentation](https://cloud.yandex.com/docs/message-queue/api-ref/queue/CreateQueue).`,
				},
				resource.Attribute{
					Name:        "name_prefix",
					Description: `(Optional, forces new resource) Generates random name with the specified prefix. Conflicts with ` + "`" + `name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "visibility_timeout_seconds",
					Description: `(Optional) [Visibility timeout](https://cloud.yandex.com/docs/message-queue/concepts/visibility-timeout) for messages in a queue, specified in seconds. Valid values: from 0 to 43200 seconds (12 hours). Default: 30.`,
				},
				resource.Attribute{
					Name:        "message_retention_seconds",
					Description: `(Optional) The length of time in seconds to retain a message. Valid values: from 60 seconds (1 minute) to 1209600 seconds (14 days). Default: 345600 (4 days). For more information see [documentation](https://cloud.yandex.com/docs/message-queue/api-ref/queue/CreateQueue).`,
				},
				resource.Attribute{
					Name:        "max_message_size",
					Description: `(Optional) Maximum message size in bytes. Valid values: from 1024 bytes (1 KB) to 262144 bytes (256 KB). Default: 262144 (256 KB). For more information see [documentation](https://cloud.yandex.com/docs/message-queue/api-ref/queue/CreateQueue).`,
				},
				resource.Attribute{
					Name:        "delay_seconds",
					Description: `(Optional) Number of seconds to [delay the message from being available for processing](https://cloud.yandex.com/docs/message-queue/concepts/delay-queues#delay-queues). Valid values: from 0 to 900 seconds (15 minutes). Default: 0.`,
				},
				resource.Attribute{
					Name:        "receive_wait_time_seconds",
					Description: `(Optional) Wait time for the [ReceiveMessage](https://cloud.yandex.com/docs/message-queue/api-ref/message/ReceiveMessage) method (for long polling), in seconds. Valid values: from 0 to 20 seconds. Default: 0. For more information about long polling see [documentation](https://cloud.yandex.com/docs/message-queue/concepts/long-polling).`,
				},
				resource.Attribute{
					Name:        "redrive_policy",
					Description: `(Optional) Message redrive policy in [Dead Letter Queue](https://cloud.yandex.com/docs/message-queue/concepts/dlq). The source queue and DLQ must be the same type: for FIFO queues, the DLQ must also be a FIFO queue. For more information about redrive policy see [documentation](https://cloud.yandex.com/docs/message-queue/api-ref/queue/CreateQueue). Also you can use example in this page.`,
				},
				resource.Attribute{
					Name:        "fifo_queue",
					Description: `(Optional, forces new resource) Is this queue [FIFO](https://cloud.yandex.com/docs/message-queue/concepts/queue#fifo-queues). If this parameter is not used, a standard queue is created. You cannot change the parameter value for a created queue.`,
				},
				resource.Attribute{
					Name:        "content_based_deduplication",
					Description: `(Optional) Enables [content-based deduplication](https://cloud.yandex.com/docs/message-queue/concepts/deduplication#content-based-deduplication). Can be used only if queue is [FIFO](https://cloud.yandex.com/docs/message-queue/concepts/queue#fifo-queues).`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Optional) The [access key](https://cloud.yandex.com/docs/iam/operations/sa/create-access-key) to use when applying changes. If omitted, ` + "`" + `ymq_access_key` + "`" + ` specified in provider config is used. For more information see [documentation](https://cloud.yandex.com/docs/message-queue/quickstart).`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Optional) The [secret key](https://cloud.yandex.com/docs/iam/operations/sa/create-access-key) to use when applying changes. If omitted, ` + "`" + `ymq_secret_key` + "`" + ` specified in provider config is used. For more information see [documentation](https://cloud.yandex.com/docs/message-queue/quickstart). ## Attributes Reference Message Queue also has the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `URL of the Yandex Message Queue.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `ARN of the Yandex Message Queue. It is used for setting up a [redrive policy](https://cloud.yandex.com/docs/message-queue/concepts/dlq). See [documentation](https://cloud.yandex.com/docs/message-queue/api-ref/queue/SetQueueAttributes). ## Import Yandex Message Queues can be imported using its ` + "`" + `queue url` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_message_queue.example_import_queue https://message-queue.api.cloud.yandex.net/abcdefghijklmn123456/opqrstuvwxyz87654321/ymq_terraform_import_example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `URL of the Yandex Message Queue.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `ARN of the Yandex Message Queue. It is used for setting up a [redrive policy](https://cloud.yandex.com/docs/message-queue/concepts/dlq). See [documentation](https://cloud.yandex.com/docs/message-queue/api-ref/queue/SetQueueAttributes). ## Import Yandex Message Queues can be imported using its ` + "`" + `queue url` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_message_queue.example_import_queue https://message-queue.api.cloud.yandex.net/abcdefghijklmn123456/opqrstuvwxyz87654321/ymq_terraform_import_example ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_resourcemanager_cloud_iam_binding",
			Category:         "Yandex Resource Manager Resources",
			ShortDescription: `Allows management of a single IAM binding for a Yandex Resource Manager cloud.`,
			Description:      ``,
			Keywords: []string{
				"resource",
				"manager",
				"resourcemanager",
				"cloud",
				"iam",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_id",
					Description: `(Required) ID of the cloud to attach the policy to.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be assigned. Only one ` + "`" + `yandex_resourcemanager_cloud_iam_binding` + "`" + ` can be used per role.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Required) An array of identities that will be granted the privilege in the ` + "`" + `role` + "`" + `. Each entry can have one of the following values:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_resourcemanager_cloud_iam_member",
			Category:         "Yandex Resource Manager Resources",
			ShortDescription: `Allows management of a single member for a single IAM binding on a Yandex Resource Manager cloud.`,
			Description:      ``,
			Keywords: []string{
				"resource",
				"manager",
				"resourcemanager",
				"cloud",
				"iam",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_id",
					Description: `(Required) ID of the cloud to attach a policy to.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be assigned.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The identity that will be granted the privilege that is specified in the ` + "`" + `role` + "`" + ` field. This field can have one of the following values:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_resourcemanager_folder_iam_binding",
			Category:         "Yandex Resource Manager Resources",
			ShortDescription: `Allows management of a single IAM binding for a Yandex Resource Manager folder.`,
			Description:      ``,
			Keywords: []string{
				"resource",
				"manager",
				"resourcemanager",
				"folder",
				"iam",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Required) ID of the folder to attach a policy to.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be assigned. Only one ` + "`" + `yandex_resourcemanager_folder_iam_binding` + "`" + ` can be used per role.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Required) An array of identities that will be granted the privilege that is specified in the ` + "`" + `role` + "`" + ` field. Each entry can have one of the following values:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_resourcemanager_folder_iam_member",
			Category:         "Yandex Resource Manager Resources",
			ShortDescription: `Allows management of a single member for a single IAM binding for a Yandex Resource Manager folder.`,
			Description:      ``,
			Keywords: []string{
				"resource",
				"manager",
				"resourcemanager",
				"folder",
				"iam",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Required) ID of the folder to attach a policy to.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be assigned.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The identity that will be granted the privilege that is specified in the ` + "`" + `role` + "`" + ` field. This field can have one of the following values:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_resourcemanager_folder_iam_policy",
			Category:         "Yandex Resource Manager Resources",
			ShortDescription: `Allows management of the IAM policy for a Yandex Resource Manager folder.`,
			Description:      ``,
			Keywords: []string{
				"resource",
				"manager",
				"resourcemanager",
				"folder",
				"iam",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Required) ID of the folder that the policy is attached to.`,
				},
				resource.Attribute{
					Name:        "policy_data",
					Description: `(Required) The ` + "`" + `yandex_iam_policy` + "`" + ` data source that represents the IAM policy that will be applied to the folder. This policy overrides any existing policy applied to the folder.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_storage_bucket",
			Category:         "Yandex Object Storage",
			ShortDescription: `Allows management of a Yandex.Cloud Storage Bucket.`,
			Description:      ``,
			Keywords: []string{
				"object",
				"storage",
				"bucket",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Optional, Forces new resource) The name of the bucket. If omitted, Terraform will assign a random, unique name.`,
				},
				resource.Attribute{
					Name:        "bucket_prefix",
					Description: `(Optional, Forces new resource) Creates a unique bucket name beginning with the specified prefix. Conflicts with ` + "`" + `bucket` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Optional) The access key to use when applying changes. If omitted, ` + "`" + `storage_access_key` + "`" + ` specified in provider config is used.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Optional) The secret key to use when applying changes. If omitted, ` + "`" + `storage_secret_key` + "`" + ` specified in provider config is used.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The [predefined ACL](https://cloud.yandex.com/docs/storage/concepts/acl#predefined_acls) to apply. Defaults to ` + "`" + `private` + "`" + `. Conflicts with ` + "`" + `grant` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "grant",
					Description: `(Optional) An [ACL policy grant](https://cloud.yandex.com/docs/storage/concepts/acl#permissions-types). Conflicts with ` + "`" + `acl` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "force_destroy",
					Description: `(Optional, Default: ` + "`" + `false` + "`" + `) A boolean that indicates all objects should be deleted from the bucket so that the bucket can be destroyed without error. These objects are`,
				},
				resource.Attribute{
					Name:        "website",
					Description: `(Optional) A website object (documented below).`,
				},
				resource.Attribute{
					Name:        "cors_rule",
					Description: `(Optional) A rule of [Cross-Origin Resource Sharing](https://cloud.yandex.com/docs/storage/cors/) (documented below). The ` + "`" + `website` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "index_document",
					Description: `(Required) Storage returns this index document when requests are made to the root domain or any of the subfolders.`,
				},
				resource.Attribute{
					Name:        "error_document",
					Description: `(Optional) An absolute path to the document to return in case of a 4XX error. The ` + "`" + `CORS` + "`" + ` object supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The name of the bucket.`,
				},
				resource.Attribute{
					Name:        "bucket_domain_name",
					Description: `The bucket domain name.`,
				},
				resource.Attribute{
					Name:        "website_endpoint",
					Description: `The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.`,
				},
				resource.Attribute{
					Name:        "website_domain",
					Description: `The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. ## Import Storage bucket can be imported using the ` + "`" + `bucket` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_storage_bucket.bucket bucket-name ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The name of the bucket.`,
				},
				resource.Attribute{
					Name:        "bucket_domain_name",
					Description: `The bucket domain name.`,
				},
				resource.Attribute{
					Name:        "website_endpoint",
					Description: `The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.`,
				},
				resource.Attribute{
					Name:        "website_domain",
					Description: `The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. ## Import Storage bucket can be imported using the ` + "`" + `bucket` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_storage_bucket.bucket bucket-name ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_storage_object",
			Category:         "Yandex Object Storage",
			ShortDescription: `Allows management of a Yandex.Cloud Storage Object.`,
			Description:      ``,
			Keywords: []string{
				"object",
				"storage",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) The name of the containing bucket.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The name of the object once it is in the bucket.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional, conflicts with ` + "`" + `content` + "`" + ` and ` + "`" + `content_base64` + "`" + `) The path to a file that will be read and uploaded as raw bytes for the object content.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional, conflicts with ` + "`" + `source` + "`" + ` and ` + "`" + `content_base64` + "`" + `) Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.`,
				},
				resource.Attribute{
					Name:        "content_base64",
					Description: `(Optional, conflicts with ` + "`" + `source` + "`" + ` and ` + "`" + `content` + "`" + `) Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the ` + "`" + `gzipbase64` + "`" + ` function with small text strings. For larger objects, use ` + "`" + `source` + "`" + ` to stream the content from a disk file.`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `(Optional) The access key to use when applying changes. If omitted, ` + "`" + `storage_access_key` + "`" + ` specified in config is used.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Optional) The secret key to use when applying changes. If omitted, ` + "`" + `storage_secret_key` + "`" + ` specified in config is used.`,
				},
				resource.Attribute{
					Name:        "acl",
					Description: `(Optional) The [predefined ACL](https://cloud.yandex.com/docs/storage/concepts/acl#predefined_acls) to apply. Defaults to ` + "`" + `private` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ` + "`" + `key` + "`" + ` of the resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_vpc_network",
			Category:         "Yandex VPC Resources",
			ShortDescription: `Manages a network within Yandex.Cloud.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the network. Provided by the client when the network is created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource. Provide this property when you create the resource.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Labels to apply to this network. A list of key/value pairs. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the key. ## Import A network can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_vpc_network.default network_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the key. ## Import A network can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_vpc_network.default network_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_vpc_route_table",
			Category:         "Yandex VPC Resources",
			ShortDescription: `A VPC route table is a virtual version of the traditional route table on router device.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"route",
				"table",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) ID of the network this route table belongs to. - - -`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the route table. Provided by the client when the route table is created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of the route table. Provide this property when you create the resource.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the folder to which the resource belongs. If omitted, the provider folder is used.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Labels to assign to this route table. A list of key/value pairs.`,
				},
				resource.Attribute{
					Name:        "static_route",
					Description: `(Optional) A list of static route records for the route table. The structure is documented below. The ` + "`" + `static_route` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "destination_prefix",
					Description: `Route prefix in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "next_hop_address",
					Description: `Address of the next hop. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the route table. ## Import A route table can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_vpc_route_table.default route_table_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of the route table. ## Import A route table can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_vpc_route_table.default route_table_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_vpc_security_group",
			Category:         "Yandex VPC Resources",
			ShortDescription: `Yandex VPC Security Group.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"security",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `Status of this security group.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of this security group. --- The ` + "`" + `ingress` + "`" + ` and ` + "`" + `egress` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the rule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `Status of this security group.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation timestamp of this security group. --- The ` + "`" + `ingress` + "`" + ` and ` + "`" + `egress` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the rule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "yandex_vpc_subnet",
			Category:         "Yandex VPC Resources",
			ShortDescription: `A VPC network is a virtual version of the traditional physical networks that exist within and between physical data centers.`,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"subnet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) ID of the network this subnet belongs to. Only networks that are in the distributed mode can have subnets.`,
				},
				resource.Attribute{
					Name:        "v4_cidr_blocks",
					Description: `(Required) A list of blocks of internal IPv4 addresses that are owned by this subnet. Provide this property when you create the subnet. For example, 10.0.0.0/22 or 192.168.0.0/16. Blocks of addresses must be unique and non-overlapping within a network. Minimum subnet size is /28, and maximum subnet size is /16. Only IPv4 is supported.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) Name of the Yandex.Cloud zone for this subnet. - - -`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the subnet. Provided by the client when the subnet is created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of the subnet. Provide this property when you create the resource.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The ID of the folder to which the resource belongs. If omitted, the provider folder is used.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Labels to assign to this subnet. A list of key/value pairs.`,
				},
				resource.Attribute{
					Name:        "route_table_id",
					Description: `(Optional) The ID of the route table to assign to this subnet. Assigned route table should belong to the same network as this subnet.`,
				},
				resource.Attribute{
					Name:        "dhcp_options",
					Description: `(Optional) Options for DHCP client. The structure is documented below. --- The ` + "`" + `dhcp_options` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `(Optional) Domain name.`,
				},
				resource.Attribute{
					Name:        "domain_name_servers",
					Description: `(Optional) Domain name server IP addresses.`,
				},
				resource.Attribute{
					Name:        "ntp_servers",
					Description: `(Optional) NTP server IP addresses. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "v6_cidr_blocks",
					Description: `An optional list of blocks of IPv6 addresses that are owned by this subnet. ## Timeouts This resource provides the following configuration options for [timeouts](/docs/configuration/resources.html#timeouts): - ` + "`" + `create` + "`" + ` - Default is 3 minute. - ` + "`" + `update` + "`" + ` - Default is 3 minute. - ` + "`" + `delete` + "`" + ` - Default is 3 minute. ## Import A subnet can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_vpc_subnet.default subnet_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "v6_cidr_blocks",
					Description: `An optional list of blocks of IPv6 addresses that are owned by this subnet. ## Timeouts This resource provides the following configuration options for [timeouts](/docs/configuration/resources.html#timeouts): - ` + "`" + `create` + "`" + ` - Default is 3 minute. - ` + "`" + `update` + "`" + ` - Default is 3 minute. - ` + "`" + `delete` + "`" + ` - Default is 3 minute. ## Import A subnet can be imported using the ` + "`" + `id` + "`" + ` of the resource, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import yandex_vpc_subnet.default subnet_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"yandex_compute_disk":                          0,
		"yandex_compute_image":                         1,
		"yandex_compute_instance":                      2,
		"yandex_compute_instance_group":                3,
		"yandex_compute_snapshot":                      4,
		"yandex_cr_container_registry":                 5,
		"yandex_dataproc_cluster":                      6,
		"yandex_function":                              7,
		"yandex_function_iam_binding":                  8,
		"yandex_function_trigger":                      9,
		"yandex_iam_service_account":                   10,
		"yandex_iam_service_account_api_key":           11,
		"yandex_iam_service_account_iam_binding":       12,
		"yandex_iam_service_account_iam_member":        13,
		"yandex_iam_service_account_iam_policy":        14,
		"yandex_iam_service_account_key":               15,
		"yandex_iam_service_account_static_access_key": 16,
		"yandex_iot_core_device":                       17,
		"yandex_iot_core_registry":                     18,
		"yandex_kms_secret_ciphertext":                 19,
		"yandex_kms_symmetric_key":                     20,
		"yandex_kubernetes_cluster":                    21,
		"yandex_kubernetes_node_group":                 22,
		"yandex_lb_network_load_balancer":              23,
		"yandex_lb_target_group":                       24,
		"yandex_mdb_clickhouse_cluster":                25,
		"yandex_mdb_mongodb_cluster":                   26,
		"yandex_mdb_mysql_cluster":                     27,
		"yandex_mdb_postgresql_cluster":                28,
		"yandex_mdb_redis_cluster":                     29,
		"yandex_message_queue":                         30,
		"yandex_resourcemanager_cloud_iam_binding":     31,
		"yandex_resourcemanager_cloud_iam_member":      32,
		"yandex_resourcemanager_folder_iam_binding":    33,
		"yandex_resourcemanager_folder_iam_member":     34,
		"yandex_resourcemanager_folder_iam_policy":     35,
		"yandex_storage_bucket":                        36,
		"yandex_storage_object":                        37,
		"yandex_vpc_network":                           38,
		"yandex_vpc_route_table":                       39,
		"yandex_vpc_security_group":                    40,
		"yandex_vpc_subnet":                            41,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
