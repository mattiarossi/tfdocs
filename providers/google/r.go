package google

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "google_access_context_manager_access_level",
			Category:         "Google Access Context Manager / VPC Service Control Resources",
			ShortDescription: `An AccessLevel is a label that can be applied to requests to GCP services, along with a list of requirements necessary for the label to be applied.`,
			Description:      ``,
			Keywords: []string{
				"access",
				"context",
				"manager",
				"vpc",
				"service",
				"control",
				"level",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "title",
					Description: `(Required) Human readable title. Must be unique within the Policy.`,
				},
				resource.Attribute{
					Name:        "parent",
					Description: `(Required) The AccessPolicy this AccessLevel lives in. Format: accessPolicies/{policy_id}`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Resource name for the Access Level. The short_name component must begin with a letter and only include alphanumeric and '_'. Format: accessPolicies/{policy_id}/accessLevels/{short_name} - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the AccessLevel and its use. Does not affect behavior.`,
				},
				resource.Attribute{
					Name:        "basic",
					Description: `(Optional) A set of predefined conditions for the access level and a combining function. Structure is documented below. The ` + "`" + `basic` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "combining_function",
					Description: `(Optional) How the conditions list should be combined to determine if a request is granted this AccessLevel. If AND is used, each Condition in conditions must be satisfied for the AccessLevel to be applied. If OR is used, at least one Condition in conditions must be satisfied for the AccessLevel to be applied. Defaults to AND if unspecified.`,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `(Required) A set of requirements for the AccessLevel to be granted. Structure is documented below. The ` + "`" + `conditions` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "ip_subnetworks",
					Description: `(Optional) A list of CIDR block IP subnetwork specification. May be IPv4 or IPv6. Note that for a CIDR IP address block, the specified IP address portion must be properly truncated (i.e. all the host bits must be zero) or the input is considered malformed. For example, "192.0.2.0/24" is accepted but "192.0.2.1/24" is not. Similarly, for IPv6, "2001:db8::/32" is accepted whereas "2001:db8::1/32" is not. The originating IP of a request must be in one of the listed subnets in order for this Condition to be true. If empty, all IP addresses are allowed.`,
				},
				resource.Attribute{
					Name:        "required_access_levels",
					Description: `(Optional) A list of other access levels defined in the same Policy, referenced by resource name. Referencing an AccessLevel which does not exist is an error. All access levels listed must be granted for the Condition to be true. Format: accessPolicies/{policy_id}/accessLevels/{short_name}`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Optional) An allowed list of members (users, groups, service accounts). The signed-in user originating the request must be a part of one of the provided members. If not specified, a request may come from any user (logged in/not logged in, not present in any groups, etc.). Formats: ` + "`" + `user:{emailid}` + "`" + `, ` + "`" + `group:{emailid}` + "`" + `, ` + "`" + `serviceAccount:{emailid}` + "`" + ``,
				},
				resource.Attribute{
					Name:        "negate",
					Description: `(Optional) Whether to negate the Condition. If true, the Condition becomes a NAND over its non-empty fields, each field must be false for the Condition overall to be satisfied. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "device_policy",
					Description: `(Optional) Device specific restrictions, all restrictions must hold for the Condition to be true. If not specified, all devices are allowed. Structure is documented below. The ` + "`" + `device_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "require_screen_lock",
					Description: `(Optional) Whether or not screenlock is required for the DevicePolicy to be true. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "allowed_encryption_statuses",
					Description: `(Optional) A list of allowed encryptions statuses. An empty list allows all statuses.`,
				},
				resource.Attribute{
					Name:        "allowed_device_management_levels",
					Description: `(Optional) A list of allowed device management levels. An empty list allows all management levels.`,
				},
				resource.Attribute{
					Name:        "os_constraints",
					Description: `(Optional) A list of allowed OS versions. An empty list allows all types and all versions. Structure is documented below. The ` + "`" + `os_constraints` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "minimum_version",
					Description: `(Optional) The minimum allowed OS version. If not set, any version of this OS satisfies the constraint. Format: "major.minor.patch" such as "10.5.301", "9.2.1".`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `(Optional) The operating system type of the device. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 6 minutes. - ` + "`" + `update` + "`" + ` - Default is 6 minutes. - ` + "`" + `delete` + "`" + ` - Default is 6 minutes. ## Import AccessLevel can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_access_context_manager_access_level.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_access_context_manager_access_policy",
			Category:         "Google Access Context Manager / VPC Service Control Resources",
			ShortDescription: `AccessPolicy is a container for AccessLevels (which define the necessary attributes to use GCP services) and ServicePerimeters (which define regions of services able to freely pass data within a perimeter).`,
			Description:      ``,
			Keywords: []string{
				"access",
				"context",
				"manager",
				"vpc",
				"service",
				"control",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "parent",
					Description: `(Required) The parent of this AccessPolicy in the Cloud Resource Hierarchy. Format: organizations/{organization_id}`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Required) Human readable title. Does not affect behavior. - - - ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Resource name of the AccessPolicy. Format: {policy_id}`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Time the AccessPolicy was created in UTC.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Time the AccessPolicy was updated in UTC. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 6 minutes. - ` + "`" + `update` + "`" + ` - Default is 6 minutes. - ` + "`" + `delete` + "`" + ` - Default is 6 minutes. ## Import AccessPolicy can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_access_context_manager_access_policy.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Resource name of the AccessPolicy. Format: {policy_id}`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Time the AccessPolicy was created in UTC.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Time the AccessPolicy was updated in UTC. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 6 minutes. - ` + "`" + `update` + "`" + ` - Default is 6 minutes. - ` + "`" + `delete` + "`" + ` - Default is 6 minutes. ## Import AccessPolicy can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_access_context_manager_access_policy.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_access_context_manager_service_perimeter",
			Category:         "Google Access Context Manager / VPC Service Control Resources",
			ShortDescription: `ServicePerimeter describes a set of GCP resources which can freely import and export data amongst themselves, but not export outside of the ServicePerimeter.`,
			Description:      ``,
			Keywords: []string{
				"access",
				"context",
				"manager",
				"vpc",
				"service",
				"control",
				"perimeter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "title",
					Description: `(Required) Human readable title. Must be unique within the Policy.`,
				},
				resource.Attribute{
					Name:        "parent",
					Description: `(Required) The AccessPolicy this ServicePerimeter lives in. Format: accessPolicies/{policy_id}`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Resource name for the ServicePerimeter. The short_name component must begin with a letter and only include alphanumeric and '_'. Format: accessPolicies/{policy_id}/servicePerimeters/{short_name} - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the ServicePerimeter and its use. Does not affect behavior.`,
				},
				resource.Attribute{
					Name:        "perimeter_type",
					Description: `(Optional) Specifies the type of the Perimeter. There are two types: regular and bridge. Regular Service Perimeter contains resources, access levels, and restricted services. Every resource can be in at most ONE regular Service Perimeter. In addition to being in a regular service perimeter, a resource can also be in zero or more perimeter bridges. A perimeter bridge only contains resources. Cross project operations are permitted if all effected resources share some perimeter (whether bridge or regular). Perimeter Bridge does not contain access levels or services: those are governed entirely by the regular perimeter that resource is in. Perimeter Bridges are typically useful when building more complex topologies with many independent perimeters that need to share some data with a common perimeter, but should not be able to share data among themselves.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) ServicePerimeter configuration. Specifies sets of resources, restricted services and access levels that determine perimeter content and boundaries. Structure is documented below. The ` + "`" + `status` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Optional) A list of GCP resources that are inside of the service perimeter. Currently only projects are allowed. Format: projects/{project_number}`,
				},
				resource.Attribute{
					Name:        "access_levels",
					Description: `(Optional) A list of AccessLevel resource names that allow resources within the ServicePerimeter to be accessed from the internet. AccessLevels listed must be in the same policy as this ServicePerimeter. Referencing a nonexistent AccessLevel is a syntax error. If no AccessLevel names are listed, resources within the perimeter can only be accessed via GCP calls with request origins within the perimeter. For Service Perimeter Bridge, must be empty. Format: accessPolicies/{policy_id}/accessLevels/{access_level_name}`,
				},
				resource.Attribute{
					Name:        "restricted_services",
					Description: `(Optional) GCP services that are subject to the Service Perimeter restrictions. Must contain a list of services. For example, if ` + "`" + `storage.googleapis.com` + "`" + ` is specified, access to the storage buckets inside the perimeter must meet the perimeter's access restrictions. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Time the AccessPolicy was created in UTC.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Time the AccessPolicy was updated in UTC. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 6 minutes. - ` + "`" + `update` + "`" + ` - Default is 6 minutes. - ` + "`" + `delete` + "`" + ` - Default is 6 minutes. ## Import ServicePerimeter can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_access_context_manager_service_perimeter.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `Time the AccessPolicy was created in UTC.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Time the AccessPolicy was updated in UTC. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 6 minutes. - ` + "`" + `update` + "`" + ` - Default is 6 minutes. - ` + "`" + `delete` + "`" + ` - Default is 6 minutes. ## Import ServicePerimeter can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_access_context_manager_service_perimeter.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_app_engine_application",
			Category:         "Google App Engine Resources",
			ShortDescription: `Allows management of an App Engine application.`,
			Description:      ``,
			Keywords: []string{
				"app",
				"engine",
				"application",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Required) The project ID to create the application under. ~>`,
				},
				resource.Attribute{
					Name:        "location_id",
					Description: `(Required) The [location](https://cloud.google.com/appengine/docs/locations) to serve the app from.`,
				},
				resource.Attribute{
					Name:        "auth_domain",
					Description: `(Optional) The domain to authenticate users with when using App Engine's User API.`,
				},
				resource.Attribute{
					Name:        "serving_status",
					Description: `(Optional) The serving status of the app.`,
				},
				resource.Attribute{
					Name:        "feature_settings",
					Description: `(Optional) A block of optional settings to configure specific App Engine features:`,
				},
				resource.Attribute{
					Name:        "split_health_checks",
					Description: `(Optional) Set to false to use the legacy health check instead of the readiness and liveness checks. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Unique name of the app, usually ` + "`" + `apps/{PROJECT_ID}` + "`" + ``,
				},
				resource.Attribute{
					Name:        "url_dispatch_rule",
					Description: `A list of dispatch rule blocks. Each block has a ` + "`" + `domain` + "`" + `, ` + "`" + `path` + "`" + `, and ` + "`" + `service` + "`" + ` field.`,
				},
				resource.Attribute{
					Name:        "code_bucket",
					Description: `The GCS bucket code is being stored in for this app.`,
				},
				resource.Attribute{
					Name:        "default_hostname",
					Description: `The default hostname for this app.`,
				},
				resource.Attribute{
					Name:        "default_bucket",
					Description: `The GCS bucket content is being stored in for this app.`,
				},
				resource.Attribute{
					Name:        "gcr_domain",
					Description: `The GCR domain used for storing managed Docker images for this app. ## Import Applications can be imported using the ID of the project the application belongs to, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_app_engine_application.app your-project-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Unique name of the app, usually ` + "`" + `apps/{PROJECT_ID}` + "`" + ``,
				},
				resource.Attribute{
					Name:        "url_dispatch_rule",
					Description: `A list of dispatch rule blocks. Each block has a ` + "`" + `domain` + "`" + `, ` + "`" + `path` + "`" + `, and ` + "`" + `service` + "`" + ` field.`,
				},
				resource.Attribute{
					Name:        "code_bucket",
					Description: `The GCS bucket code is being stored in for this app.`,
				},
				resource.Attribute{
					Name:        "default_hostname",
					Description: `The default hostname for this app.`,
				},
				resource.Attribute{
					Name:        "default_bucket",
					Description: `The GCS bucket content is being stored in for this app.`,
				},
				resource.Attribute{
					Name:        "gcr_domain",
					Description: `The GCR domain used for storing managed Docker images for this app. ## Import Applications can be imported using the ID of the project the application belongs to, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_app_engine_application.app your-project-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_app_engine_firewall_rule",
			Category:         "Google App Engine Resources",
			ShortDescription: `A single firewall rule that is evaluated against incoming traffic and provides an action to take on matched requests.`,
			Description:      ``,
			Keywords: []string{
				"app",
				"engine",
				"firewall",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_range",
					Description: `(Required) IP address or range, defined using CIDR notation, of requests that this rule applies to.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) The action to take if this rule matches. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional string description of this rule.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) A positive integer that defines the order of rule evaluation. Rules with the lowest priority are evaluated first. A default rule at priority Int32.MaxValue matches all IPv4 and IPv6 traffic when no previous rule matches. Only the action of this rule can be modified by the user.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import FirewallRule can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_app_engine_firewall_rule.default {{project}}/{{priority}} $ terraform import google_app_engine_firewall_rule.default {{priority}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_app_engine_standard_app_version",
			Category:         "Google App Engine Resources",
			ShortDescription: `Standard App Version resource to create a new version of standard GAE Application.`,
			Description:      ``,
			Keywords: []string{
				"app",
				"engine",
				"standard",
				"version",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "runtime",
					Description: `(Required) Desired runtime. Example python27. - - -`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `(Optional) Relative name of the version within the service. For example, ` + "`" + `v1` + "`" + `. Version names can contain only lowercase letters, numbers, or hyphens. Reserved names,"default", "latest", and any name with the prefix "ah-".`,
				},
				resource.Attribute{
					Name:        "threadsafe",
					Description: `(Optional) Whether multiple requests can be dispatched to this version at once.`,
				},
				resource.Attribute{
					Name:        "runtime_api_version",
					Description: `(Optional) The version of the API in the given runtime environment. Please see the app.yaml reference for valid values at https://cloud.google.com/appengine/docs/standard//config/appref`,
				},
				resource.Attribute{
					Name:        "handlers",
					Description: `(Optional) An ordered list of URL-matching patterns that should be applied to incoming requests. The first matching URL handles the request and other request handlers are not attempted. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "libraries",
					Description: `(Optional) Configuration for third-party Python runtime libraries that are required by the application. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "env_variables",
					Description: `(Optional) Environment variables available to the application.`,
				},
				resource.Attribute{
					Name:        "deployment",
					Description: `(Optional) Code and application artifacts that make up this version. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "entrypoint",
					Description: `(Optional) The entrypoint for the application. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Optional) AppEngine service resource`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "noop_on_destroy",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + `, the application version will not be deleted. The ` + "`" + `handlers` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "url_regex",
					Description: `(Optional) URL prefix. Uses regular expression syntax, which means regexp special characters must be escaped, but should not contain groupings. All URLs that begin with this prefix are handled by this handler, using the portion of the URL after the prefix as part of the file path.`,
				},
				resource.Attribute{
					Name:        "security_level",
					Description: `(Optional) Security (HTTPS) enforcement for this URL.`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `(Optional) Methods to restrict access to a URL based on login status.`,
				},
				resource.Attribute{
					Name:        "auth_fail_action",
					Description: `(Optional) Actions to take when the user is not logged in.`,
				},
				resource.Attribute{
					Name:        "redirect_http_response_code",
					Description: `(Optional) Redirect codes.`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `(Optional) Executes a script to handle the requests that match this URL pattern. Only the auto value is supported for Node.js in the App Engine standard environment, for example "script:" "auto". Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "static_files",
					Description: `(Optional) Files served directly to the user for a given URL, such as images, CSS stylesheets, or JavaScript source files. Static file handlers describe which files in the application directory are static files, and which URLs serve them. Structure is documented below. The ` + "`" + `script` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "script_path",
					Description: `(Optional) Path to the script from the application root directory. The ` + "`" + `static_files` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Path to the static files matched by the URL pattern, from the application root directory. The path can refer to text matched in groupings in the URL pattern.`,
				},
				resource.Attribute{
					Name:        "upload_path_regex",
					Description: `(Optional) Regular expression that matches the file paths for all files that should be referenced by this handler.`,
				},
				resource.Attribute{
					Name:        "http_headers",
					Description: `(Optional) HTTP headers to use for all responses from these URLs. An object containing a list of "key:value" value pairs.".`,
				},
				resource.Attribute{
					Name:        "mime_type",
					Description: `(Optional) MIME type used to serve all files served by this handler. Defaults to file-specific MIME types, which are derived from each file's filename extension.`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `(Optional) Time a static file served by this handler should be cached by web proxies and browsers. A duration in seconds with up to nine fractional digits, terminated by 's'. Example "3.5s".`,
				},
				resource.Attribute{
					Name:        "require_matching_file",
					Description: `(Optional) Whether this handler should match the request if the file referenced by the handler does not exist.`,
				},
				resource.Attribute{
					Name:        "application_readable",
					Description: `(Optional) Whether files should also be uploaded as code data. By default, files declared in static file handlers are uploaded as static data and are only served to end users; they cannot be read by the application. If enabled, uploads are charged against both your code and static data storage resource quotas. The ` + "`" + `libraries` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the library. Example "django".`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Version of the library to select, or "latest". The ` + "`" + `deployment` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "zip",
					Description: `(Optional) Zip File Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "files",
					Description: `(Optional) Manifest of the files stored in Google Cloud Storage that are included as part of this version. All files must be readable using the credentials supplied with this call. Structure is documented below. The ` + "`" + `zip` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "source_url",
					Description: `(Optional) Source URL`,
				},
				resource.Attribute{
					Name:        "files_count",
					Description: `(Optional) files count The ` + "`" + `files` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The identifier for this object. Format specified above.`,
				},
				resource.Attribute{
					Name:        "sha1_sum",
					Description: `(Optional) SHA1 checksum of the file`,
				},
				resource.Attribute{
					Name:        "source_url",
					Description: `(Optional) Source URL The ` + "`" + `entrypoint` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "shell",
					Description: `(Optional) The format should be a shell command that can be fed to bash -c. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Full path to the Version resource in the API. Example, "v1". ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import StandardAppVersion can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_app_engine_standard_app_version.default apps/{{project}}/services/{{service}}/versions/{{version_id}} $ terraform import google_app_engine_standard_app_version.default {{project}}/{{service}}/{{version_id}} $ terraform import google_app_engine_standard_app_version.default {{service}}/{{version_id}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Full path to the Version resource in the API. Example, "v1". ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import StandardAppVersion can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_app_engine_standard_app_version.default apps/{{project}}/services/{{service}}/versions/{{version_id}} $ terraform import google_app_engine_standard_app_version.default {{project}}/{{service}}/{{version_id}} $ terraform import google_app_engine_standard_app_version.default {{service}}/{{version_id}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_bigquery_dataset",
			Category:         "Google BigQuery Resources",
			ShortDescription: `Datasets allow you to organize and control access to your tables.`,
			Description:      ``,
			Keywords: []string{
				"bigquery",
				"dataset",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dataset_id",
					Description: `(Required) A unique ID for this dataset, without the project name. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum length is 1,024 characters. - - -`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `(Optional) An array of objects that define dataset access for one or more entities. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "default_table_expiration_ms",
					Description: `(Optional) The default lifetime of all tables in the dataset, in milliseconds. The minimum value is 3600000 milliseconds (one hour). Once this property is set, all newly-created tables in the dataset will have an ` + "`" + `expirationTime` + "`" + ` property set to the creation time plus the value in this property, and changing the value will only affect new tables, not existing ones. When the ` + "`" + `expirationTime` + "`" + ` for a given table is reached, that table will be deleted automatically. If a table's ` + "`" + `expirationTime` + "`" + ` is modified or removed before the table expires, or if you provide an explicit ` + "`" + `expirationTime` + "`" + ` when creating a table, that value takes precedence over the default expiration time indicated by this property.`,
				},
				resource.Attribute{
					Name:        "default_partition_expiration_ms",
					Description: `(Optional) The default partition expiration for all partitioned tables in the dataset, in milliseconds. Once this property is set, all newly-created partitioned tables in the dataset will have an ` + "`" + `expirationMs` + "`" + ` property in the ` + "`" + `timePartitioning` + "`" + ` settings set to this value, and changing the value will only affect new tables, not existing ones. The storage in a partition will have an expiration time of its partition time plus this value. Setting this property overrides the use of ` + "`" + `defaultTableExpirationMs` + "`" + ` for partitioned tables: only one of ` + "`" + `defaultTableExpirationMs` + "`" + ` and ` + "`" + `defaultPartitionExpirationMs` + "`" + ` will be used for any new partitioned table. If you provide an explicit ` + "`" + `timePartitioning.expirationMs` + "`" + ` when creating or updating a partitioned table, that value takes precedence over the default partition expiration time indicated by this property.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A user-friendly description of the dataset`,
				},
				resource.Attribute{
					Name:        "friendly_name",
					Description: `(Optional) A descriptive name for the dataset`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) The labels associated with this dataset. You can use these to organize and group your datasets`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Optional) The geographic location where the dataset should reside. See [official docs](https://cloud.google.com/bigquery/docs/dataset-locations). There are two types of locations, regional or multi-regional. A regional location is a specific geographic place, such as Tokyo, and a multi-regional location is a large geographic area, such as the United States, that contains at least two geographic places. Possible regional values include: ` + "`" + `asia-east1` + "`" + `, ` + "`" + `asia-northeast1` + "`" + `, ` + "`" + `asia-southeast1` + "`" + `, ` + "`" + `australia-southeast1` + "`" + `, ` + "`" + `europe-north1` + "`" + `, ` + "`" + `europe-west2` + "`" + ` and ` + "`" + `us-east4` + "`" + `. Possible multi-regional values: ` + "`" + `EU` + "`" + ` and ` + "`" + `US` + "`" + `. The default value is multi-regional location ` + "`" + `US` + "`" + `. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "delete_contents_on_destroy",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + `, delete all the tables in the dataset when destroying the resource; otherwise, destroying the resource will fail if tables are present. The ` + "`" + `access` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) A domain to grant access to. Any users signed in with the domain specified will be granted the specified access`,
				},
				resource.Attribute{
					Name:        "group_by_email",
					Description: `(Optional) An email address of a Google Group to grant access to.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Optional) Describes the rights granted to the user specified by the other member of the access object. Primitive, Predefined and custom roles are supported. Predefined roles that have equivalent primitive roles are swapped by the API to their Primitive counterparts, and will show a diff post-create. See [official docs](https://cloud.google.com/bigquery/docs/access-control).`,
				},
				resource.Attribute{
					Name:        "special_group",
					Description: `(Optional) A special group to grant access to. Possible values include:`,
				},
				resource.Attribute{
					Name:        "user_by_email",
					Description: `(Optional) An email address of a user to grant access to. For example: fred@example.com`,
				},
				resource.Attribute{
					Name:        "view",
					Description: `(Optional) A view from a different dataset to grant access to. Queries executed against that view will have read access to tables in this dataset. The role field is not required when this field is set. If that view is updated by any user, access to the view needs to be granted again via an update operation. Structure is documented below. The ` + "`" + `view` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "dataset_id",
					Description: `(Required) The ID of the dataset containing this table.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The ID of the project containing this table.`,
				},
				resource.Attribute{
					Name:        "table_id",
					Description: `(Required) The ID of the table. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum length is 1,024 characters. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `The time when this dataset was created, in milliseconds since the epoch.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `A hash of the resource.`,
				},
				resource.Attribute{
					Name:        "last_modified_time",
					Description: `The date when this dataset or any of its tables was last modified, in milliseconds since the epoch.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Dataset can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_bigquery_dataset.default {{project}}/{{dataset_id}} $ terraform import google_bigquery_dataset.default {{project}}:{{dataset_id}} $ terraform import google_bigquery_dataset.default {{dataset_id}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_time",
					Description: `The time when this dataset was created, in milliseconds since the epoch.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `A hash of the resource.`,
				},
				resource.Attribute{
					Name:        "last_modified_time",
					Description: `The date when this dataset or any of its tables was last modified, in milliseconds since the epoch.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Dataset can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_bigquery_dataset.default {{project}}/{{dataset_id}} $ terraform import google_bigquery_dataset.default {{project}}:{{dataset_id}} $ terraform import google_bigquery_dataset.default {{dataset_id}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_bigquery_table",
			Category:         "Google BigQuery Resources",
			ShortDescription: `Creates a table resource in a dataset for Google BigQuery.`,
			Description:      ``,
			Keywords: []string{
				"bigquery",
				"table",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dataset_id",
					Description: `(Required) The dataset ID to create the table in. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "table_id",
					Description: `(Required) A unique ID for the resource. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The field description.`,
				},
				resource.Attribute{
					Name:        "expiration_time",
					Description: `(Optional) The time when this table expires, in milliseconds since the epoch. If not present, the table will persist indefinitely. Expired tables will be deleted and their storage reclaimed.`,
				},
				resource.Attribute{
					Name:        "external_data_configuration",
					Description: `(Optional) Describes the data format, location, and other properties of a table stored outside of BigQuery. By defining these properties, the data source can then be queried as if it were a standard BigQuery table. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "friendly_name",
					Description: `(Optional) A descriptive name for the table.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A mapping of labels to assign to the resource.`,
				},
				resource.Attribute{
					Name:        "schema",
					Description: `(Optional) A JSON schema for the table. Schema is required for CSV and JSON formats and is disallowed for Google Cloud Bigtable, Cloud Datastore backups, and Avro formats when using external tables. For more information see the [BigQuery API documentation](https://cloud.google.com/bigquery/docs/reference/rest/v2/tables#resource).`,
				},
				resource.Attribute{
					Name:        "time_partitioning",
					Description: `(Optional) If specified, configures time-based partitioning for this table. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "view",
					Description: `(Optional) If specified, configures this table as a view. Structure is documented below. The ` + "`" + `external_data_configuration` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "autodetect",
					Description: `(Required) - Let BigQuery try to autodetect the schema and format of the table.`,
				},
				resource.Attribute{
					Name:        "source_uris",
					Description: `(Required) A list of the fully-qualified URIs that point to your data in Google Cloud. The ` + "`" + `csv_options` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "expiration_ms",
					Description: `(Optional) Number of milliseconds for which to keep the storage for a partition.`,
				},
				resource.Attribute{
					Name:        "field",
					Description: `(Optional) The field used to determine how to create a time-based partition. If time-based partitioning is enabled without this value, the table is partitioned based on the load time.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The only type supported is DAY, which will generate one partition per day based on data loading time.`,
				},
				resource.Attribute{
					Name:        "require_partition_filter",
					Description: `(Optional) If set to true, queries over this table require a partition filter that can be used for partition elimination to be specified. The ` + "`" + `view` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `(Required) A query that BigQuery executes when the view is referenced.`,
				},
				resource.Attribute{
					Name:        "use_legacy_sql",
					Description: `(Optional) Specifies whether to use BigQuery's legacy SQL for this view. The default value is true. If set to false, the view will use BigQuery's standard SQL. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `The time when this table was created, in milliseconds since the epoch.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `A hash of the resource.`,
				},
				resource.Attribute{
					Name:        "last_modified_time",
					Description: `The time when this table was last modified, in milliseconds since the epoch.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The geographic location where the table resides. This value is inherited from the dataset.`,
				},
				resource.Attribute{
					Name:        "num_bytes",
					Description: `The size of this table in bytes, excluding any data in the streaming buffer.`,
				},
				resource.Attribute{
					Name:        "num_long_term_bytes",
					Description: `The number of bytes in the table that are considered "long-term storage".`,
				},
				resource.Attribute{
					Name:        "num_rows",
					Description: `The number of rows of data in this table, excluding any data in the streaming buffer.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Describes the table type. ## Import BigQuery tables can be imported using the ` + "`" + `project` + "`" + `, ` + "`" + `dataset_id` + "`" + `, and ` + "`" + `table_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_bigquery_table.default gcp-project:foo.bar ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_time",
					Description: `The time when this table was created, in milliseconds since the epoch.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `A hash of the resource.`,
				},
				resource.Attribute{
					Name:        "last_modified_time",
					Description: `The time when this table was last modified, in milliseconds since the epoch.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The geographic location where the table resides. This value is inherited from the dataset.`,
				},
				resource.Attribute{
					Name:        "num_bytes",
					Description: `The size of this table in bytes, excluding any data in the streaming buffer.`,
				},
				resource.Attribute{
					Name:        "num_long_term_bytes",
					Description: `The number of bytes in the table that are considered "long-term storage".`,
				},
				resource.Attribute{
					Name:        "num_rows",
					Description: `The number of rows of data in this table, excluding any data in the streaming buffer.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Describes the table type. ## Import BigQuery tables can be imported using the ` + "`" + `project` + "`" + `, ` + "`" + `dataset_id` + "`" + `, and ` + "`" + `table_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_bigquery_table.default gcp-project:foo.bar ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_bigtable_instance",
			Category:         "Google Bigtable Resources",
			ShortDescription: `Creates a Google Bigtable instance.`,
			Description:      ``,
			Keywords: []string{
				"bigtable",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name (also called Instance Id in the Cloud Console) of the Cloud Bigtable instance.`,
				},
				resource.Attribute{
					Name:        "cluster",
					Description: `(Required) A block of cluster configuration options. This can be specified 1 or 2 times. See structure below. -----`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional) The instance type to create. One of ` + "`" + `"DEVELOPMENT"` + "`" + ` or ` + "`" + `"PRODUCTION"` + "`" + `. Defaults to ` + "`" + `"PRODUCTION"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The human-readable display name of the Bigtable instance. Defaults to the instance ` + "`" + `name` + "`" + `. ----- The ` + "`" + `cluster` + "`" + ` block supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The ID of the Cloud Bigtable cluster.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The zone to create the Cloud Bigtable cluster in. Each cluster must have a different zone in the same region. Zones that support Bigtable instances are noted on the [Cloud Bigtable locations page](https://cloud.google.com/bigtable/docs/locations).`,
				},
				resource.Attribute{
					Name:        "num_nodes",
					Description: `(Optional) The number of nodes in your Cloud Bigtable cluster. Required, with a minimum of ` + "`" + `3` + "`" + ` for a ` + "`" + `PRODUCTION` + "`" + ` instance. Must be left unset for a ` + "`" + `DEVELOPMENT` + "`" + ` instance.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `(Optional) The storage type to use. One of ` + "`" + `"SSD"` + "`" + ` or ` + "`" + `"HDD"` + "`" + `. Defaults to ` + "`" + `"SSD"` + "`" + `. ## Attributes Reference Only the arguments listed above are exposed as attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_bigtable_instance_iam_policy",
			Category:         "Google Bigtable Resources",
			ShortDescription: `Collection of resources to manage IAM policy for a Bigtable instance.`,
			Description:      ``,
			Keywords: []string{
				"bigtable",
				"instance",
				"iam",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance",
					Description: `(Required) The name or relative resource id of the instance to manage IAM policies for. For ` + "`" + `google_bigtable_instance_iam_member` + "`" + ` or ` + "`" + `google_bigtable_instance_iam_binding` + "`" + `:`,
				},
				resource.Attribute{
					Name:        "member/members",
					Description: `(Required) Identities that will be granted the privilege in ` + "`" + `role` + "`" + `. Each entry can have one of the following values:`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. Only one ` + "`" + `google_bigtable_instance_iam_binding` + "`" + ` can be used per role. Note that custom roles must be of the format ` + "`" + `[projects|organizations]/{parent-name}/roles/{role-name}` + "`" + `. ` + "`" + `google_bigtable_instance_iam_policy` + "`" + ` only:`,
				},
				resource.Attribute{
					Name:        "policy_data",
					Description: `(Required) The policy data generated by a ` + "`" + `google_iam_policy` + "`" + ` data source. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the instance belongs. If it is not provided, Terraform will use the provider default. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the instances's IAM policy. ## Import Instance IAM resources can be imported using the project, instance name, role and/or member. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_bigtable_instance_iam_policy.editor "projects/{project}/instances/{instance}" $ terraform import google_bigtable_instance_iam_binding.editor "projects/{project}/instances/{instance} roles/editor" $ terraform import google_bigtable_instance_iam_member.editor "projects/{project}/instances/{instance} roles/editor user:jane@example.com" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the instances's IAM policy. ## Import Instance IAM resources can be imported using the project, instance name, role and/or member. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_bigtable_instance_iam_policy.editor "projects/{project}/instances/{instance}" $ terraform import google_bigtable_instance_iam_binding.editor "projects/{project}/instances/{instance} roles/editor" $ terraform import google_bigtable_instance_iam_member.editor "projects/{project}/instances/{instance} roles/editor user:jane@example.com" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_bigtable_table",
			Category:         "Google Bigtable Resources",
			ShortDescription: `Creates a Google Cloud Bigtable table inside an instance.`,
			Description:      ``,
			Keywords: []string{
				"bigtable",
				"table",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the table.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Required) The name of the Bigtable instance.`,
				},
				resource.Attribute{
					Name:        "split_keys",
					Description: `(Optional) A list of predefined keys to split the table on.`,
				},
				resource.Attribute{
					Name:        "column_family",
					Description: `(Optional) A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ----- ` + "`" + `column_family` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `(Optional) The name of the column family. ## Attributes Reference Only the arguments listed above are exposed as attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_binary_authorization_attestor",
			Category:         "Google Binary Authorization Resources",
			ShortDescription: `An attestor that attests to container image artifacts.`,
			Description:      ``,
			Keywords: []string{
				"binary",
				"authorization",
				"attestor",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The resource name.`,
				},
				resource.Attribute{
					Name:        "attestation_authority_note",
					Description: `(Required) A Container Analysis ATTESTATION_AUTHORITY Note, created by the user. Structure is documented below. The ` + "`" + `attestation_authority_note` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "note_reference",
					Description: `(Required) The resource name of a ATTESTATION_AUTHORITY Note, created by the user. If the Note is in a different project from the Attestor, it should be specified in the format ` + "`" + `projects/`,
				},
				resource.Attribute{
					Name:        "public_keys",
					Description: `(Optional) Public keys that verify attestations signed by this attestor. This field may be updated. If this field is non-empty, one of the specified public keys must verify that an attestation was signed by this attestor for the image specified in the admission request. If this field is empty, this attestor always returns that no valid attestations exist. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "delegation_service_account_email",
					Description: `This field will contain the service account email address that this Attestor will use as the principal when querying Container Analysis. Attestor administrators must grant this service account the IAM role needed to read attestations from the noteReference in Container Analysis (containeranalysis.notes.occurrences.viewer). This email address is fixed for the lifetime of the Attestor, but callers should not make any other assumptions about the service account email; future versions may use an email based on a different naming pattern. The ` + "`" + `public_keys` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) A descriptive comment. This field may be updated.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of this public key. Signatures verified by BinAuthz must include the ID of the public key that can be used to verify them, and that ID must match the contents of this field exactly. Additional restrictions on this field can be imposed based on which public key type is encapsulated. See the documentation on publicKey cases below for details.`,
				},
				resource.Attribute{
					Name:        "ascii_armored_pgp_public_key",
					Description: `(Optional) ASCII-armored representation of a PGP public key, as the entire output by the command ` + "`" + `gpg --export --armor foo@example.com` + "`" + ` (either LF or CRLF line endings). When using this field, id should be left blank. The BinAuthz API handlers will calculate the ID and fill it in automatically. BinAuthz computes this ID as the OpenPGP RFC4880 V4 fingerprint, represented as upper-case hex. If id is provided by the caller, it will be overwritten by the API-calculated ID.`,
				},
				resource.Attribute{
					Name:        "pkix_public_key",
					Description: `(Optional) A raw PKIX SubjectPublicKeyInfo format public key. NOTE: id may be explicitly provided by the caller when using this type of public key, but it MUST be a valid RFC3986 URI. If id is left blank, a default one will be computed based on the digest of the DER encoding of the public key. Structure is documented below. The ` + "`" + `pkix_public_key` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "public_key_pem",
					Description: `(Optional) A PEM-encoded public key, as described in ` + "`" + `https://tools.ietf.org/html/rfc7468#section-13` + "`" + ``,
				},
				resource.Attribute{
					Name:        "signature_algorithm",
					Description: `(Optional) The signature algorithm used to verify a message against a signature using this key. These signature algorithm must match the structure and any object identifiers encoded in publicKeyPem (i.e. this algorithm must match that of the public key). - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A descriptive comment. This field may be updated. The field may be displayed in chooser dialogs.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Attestor can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_binary_authorization_attestor.default projects/{{project}}/attestors/{{name}} $ terraform import google_binary_authorization_attestor.default {{project}}/{{name}} $ terraform import google_binary_authorization_attestor.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_binary_authorization_policy",
			Category:         "Google Binary Authorization Resources",
			ShortDescription: `A policy for container image binary authorization.`,
			Description:      ``,
			Keywords: []string{
				"binary",
				"authorization",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "default_admission_rule",
					Description: `(Required) Default admission rule for a cluster without a per-cluster admission rule. Structure is documented below. The ` + "`" + `default_admission_rule` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "evaluation_mode",
					Description: `(Required) How this admission rule will be evaluated.`,
				},
				resource.Attribute{
					Name:        "require_attestations_by",
					Description: `(Optional) The resource names of the attestors that must attest to a container image. If the attestor is in a different project from the policy, it should be specified in the format ` + "`" + `projects/`,
				},
				resource.Attribute{
					Name:        "enforcement_mode",
					Description: `(Required) The action when a pod creation is denied by the admission rule. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A descriptive comment.`,
				},
				resource.Attribute{
					Name:        "global_policy_evaluation_mode",
					Description: `(Optional) Controls the evaluation of a Google-maintained global admission policy for common system-level images. Images not covered by the global policy will be subject to the project admission policy.`,
				},
				resource.Attribute{
					Name:        "admission_whitelist_patterns",
					Description: `(Optional) A whitelist of image patterns to exclude from admission rules. If an image's name matches a whitelist pattern, the image's admission requests will always be permitted regardless of your admission rules. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "cluster_admission_rules",
					Description: `(Optional) Per-cluster admission rules. An admission rule specifies either that all container images used in a pod creation request must be attested to by one or more attestors, that all pod creations will be allowed, or that all pod creations will be denied. There can be at most one admission rule per cluster spec. Identifier format: ` + "`" + `{{location}}.{{clusterId}}` + "`" + `. A location is either a compute zone (e.g. ` + "`" + `us-central1-a` + "`" + `) or a region (e.g. ` + "`" + `us-central1` + "`" + `). Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `admission_whitelist_patterns` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name_pattern",
					Description: `(Optional) An image name pattern to whitelist, in the form ` + "`" + `registry/path/to/image` + "`" + `. This supports a trailing`,
				},
				resource.Attribute{
					Name:        "cluster",
					Description: `(Required) The identifier for this object. Format specified above.`,
				},
				resource.Attribute{
					Name:        "evaluation_mode",
					Description: `(Optional) How this admission rule will be evaluated.`,
				},
				resource.Attribute{
					Name:        "require_attestations_by",
					Description: `(Optional) The resource names of the attestors that must attest to a container image. If the attestor is in a different project from the policy, it should be specified in the format ` + "`" + `projects/`,
				},
				resource.Attribute{
					Name:        "enforcement_mode",
					Description: `(Optional) The action when a pod creation is denied by the admission rule. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Policy can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_binary_authorization_policy.default projects/{{project}} $ terraform import google_binary_authorization_policy.default {{project}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_cloud_run_service",
			Category:         "Google Cloud Run Resources",
			ShortDescription: `**Note:** Cloud Run as a product is in beta, however the REST API is currently still an alpha.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"run",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name must be unique within a namespace, within a Cloud Run region. Is required when creating resources. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Required) RevisionSpec holds the desired state of the Revision (from the client). Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) Metadata associated with this Service, including name, namespace, labels, and annotations. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) The location of the cloud run instance. eg us-central1 The ` + "`" + `spec` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "containers",
					Description: `(Required) Container defines the unit of execution for this Revision. In the context of a Revision, we disallow a number of the fields of this Container, including: name, ports, and volumeMounts. The runtime contract is documented here: https://github.com/knative/serving/blob/master/docs/runtime-contract.md Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "container_concurrency",
					Description: `(Optional) ContainerConcurrency specifies the maximum allowed in-flight (concurrent) requests per container of the Revision. Values are: - ` + "`" + `0` + "`" + ` thread-safe, the system should manage the max concurrency. This is the default value. - ` + "`" + `1` + "`" + ` not-thread-safe. Single concurrency - ` + "`" + `2-N` + "`" + ` thread-safe, max concurrency of N`,
				},
				resource.Attribute{
					Name:        "serving_state",
					Description: `ServingState holds a value describing the state the resources are in for this Revision. It is expected that the system will manipulate this based on routability and load. The ` + "`" + `containers` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "working_dir",
					Description: `(Optional) Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image.`,
				},
				resource.Attribute{
					Name:        "args",
					Description: `(Optional) Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell`,
				},
				resource.Attribute{
					Name:        "env_from",
					Description: `(Optional) List of sources to populate environment variables in the container. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Required) Docker image name. This is most often a reference to a container located in the container registry, such as gcr.io/cloudrun/hello More info: https://kubernetes.io/docs/concepts/containers/images`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `(Optional) Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell`,
				},
				resource.Attribute{
					Name:        "env",
					Description: `(Optional) List of environment variables to set in the container. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `(Optional) Compute Resources required by this container. Used to set values such as max memory More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources Structure is documented below. The ` + "`" + `env_from` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Optional) An optional identifier to prepend to each key in the ConfigMap.`,
				},
				resource.Attribute{
					Name:        "config_map_ref",
					Description: `(Optional) The ConfigMap to select from Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "secret_ref",
					Description: `(Optional) The Secret to select from Structure is documented below. The ` + "`" + `config_map_ref` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "optional",
					Description: `(Optional) Specify whether the ConfigMap must be defined`,
				},
				resource.Attribute{
					Name:        "local_object_reference",
					Description: `(Optional) The ConfigMap to select from. Structure is documented below. The ` + "`" + `local_object_reference` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names The ` + "`" + `secret_ref` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "local_object_reference",
					Description: `(Optional) The Secret to select from. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "optional",
					Description: `(Optional) Specify whether the Secret must be defined The ` + "`" + `local_object_reference` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names The ` + "`" + `env` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the environment variable.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any route environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to "". The ` + "`" + `resources` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "limits",
					Description: `(Optional) Limits describes the maximum amount of compute resources allowed. The values of the map is string form of the 'quantity' k8s type: https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/api/resource/quantity.go`,
				},
				resource.Attribute{
					Name:        "requests",
					Description: `(Optional) Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. The values of the map is string form of the 'quantity' k8s type: https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/api/resource/quantity.go The ` + "`" + `metadata` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and routes. More info: http://kubernetes.io/docs/user-guide/labels`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `A sequence number representing a specific generation of the desired state.`,
				},
				resource.Attribute{
					Name:        "resource_version",
					Description: `An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed. May be used for optimistic concurrency, change detection, and the watch operation on a resource or set of resources. They may only be valid for a particular resource or set of resources. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#concurrency-control-and-consistency`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `SelfLink is a URL representing this object.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `UID is a unique id generated by the server on successful creation of a resource and is not allowed to change on PUT operations. More info: http://kubernetes.io/docs/user-guide/identifiers#uids`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Required) In Cloud Run the namespace must be equal to either the project ID or project number.`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) Annotations is a key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. More info: http://kubernetes.io/docs/user-guide/annotations - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the Service. Structure is documented below. The ` + "`" + `status` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `Array of observed Service Conditions, indicating the current ready state of the service. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `From RouteStatus. URL holds the url that will distribute traffic over the provided traffic targets. It generally has the form https://{route-hash}-{project-hash}-{cluster-level-suffix}.a.run.app`,
				},
				resource.Attribute{
					Name:        "observed_generation",
					Description: `ObservedGeneration is the 'Generation' of the Route that was last processed by the controller. Clients polling for completed reconciliation should poll until observedGeneration = metadata.generation and the Ready condition's status is True or False.`,
				},
				resource.Attribute{
					Name:        "latest_created_revision_name",
					Description: `From ConfigurationStatus. LatestCreatedRevisionName is the last revision that was created from this Service's Configuration. It might not be ready yet, for that use LatestReadyRevisionName.`,
				},
				resource.Attribute{
					Name:        "latest_ready_revision_name",
					Description: `From ConfigurationStatus. LatestReadyRevisionName holds the name of the latest Revision stamped out from this Service's Configuration that has had its "Ready" condition become "True". The ` + "`" + `conditions` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `Human readable message indicating details about the current status.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the condition, one of True, False, Unknown.`,
				},
				resource.Attribute{
					Name:        "reason",
					Description: `One-word CamelCase reason for the condition's current status.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of domain mapping condition. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Service can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import -provider=google-beta google_cloud_run_service.default projects/{{project}}/locations/{{location}}/services/{{name}} $ terraform import -provider=google-beta google_cloud_run_service.default {{project}}/{{location}}/{{name}} $ terraform import -provider=google-beta google_cloud_run_service.default {{location}}/{{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the Service. Structure is documented below. The ` + "`" + `status` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `Array of observed Service Conditions, indicating the current ready state of the service. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `From RouteStatus. URL holds the url that will distribute traffic over the provided traffic targets. It generally has the form https://{route-hash}-{project-hash}-{cluster-level-suffix}.a.run.app`,
				},
				resource.Attribute{
					Name:        "observed_generation",
					Description: `ObservedGeneration is the 'Generation' of the Route that was last processed by the controller. Clients polling for completed reconciliation should poll until observedGeneration = metadata.generation and the Ready condition's status is True or False.`,
				},
				resource.Attribute{
					Name:        "latest_created_revision_name",
					Description: `From ConfigurationStatus. LatestCreatedRevisionName is the last revision that was created from this Service's Configuration. It might not be ready yet, for that use LatestReadyRevisionName.`,
				},
				resource.Attribute{
					Name:        "latest_ready_revision_name",
					Description: `From ConfigurationStatus. LatestReadyRevisionName holds the name of the latest Revision stamped out from this Service's Configuration that has had its "Ready" condition become "True". The ` + "`" + `conditions` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `Human readable message indicating details about the current status.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the condition, one of True, False, Unknown.`,
				},
				resource.Attribute{
					Name:        "reason",
					Description: `One-word CamelCase reason for the condition's current status.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of domain mapping condition. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Service can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import -provider=google-beta google_cloud_run_service.default projects/{{project}}/locations/{{location}}/services/{{name}} $ terraform import -provider=google-beta google_cloud_run_service.default {{project}}/{{location}}/{{name}} $ terraform import -provider=google-beta google_cloud_run_service.default {{location}}/{{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_cloud_scheduler_job",
			Category:         "Google Cloud Scheduler Resources",
			ShortDescription: `A scheduled job that can publish a pubsub message or a http request every X interval of time, using crontab format string.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"scheduler",
				"job",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the job.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Region where the scheduler job resides - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-readable description for the job. This string must not contain more than 500 characters.`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `(Optional) Describes the schedule on which the job will be executed.`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `(Optional) Specifies the time zone to be used in interpreting schedule. The value of this field must be a time zone name from the tz database.`,
				},
				resource.Attribute{
					Name:        "retry_config",
					Description: `(Optional) By default, if a job does not complete successfully, meaning that an acknowledgement is not received from the handler, then it will be retried with exponential backoff according to the settings Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "pubsub_target",
					Description: `(Optional) Pub/Sub target If the job providers a Pub/Sub target the cron will publish a message to the provided topic Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "app_engine_http_target",
					Description: `(Optional) App Engine HTTP target. If the job providers a App Engine HTTP target the cron will send a request to the service instance Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "http_target",
					Description: `(Optional) HTTP target. If the job providers a http_target the cron will send a request to the targeted url Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `retry_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "retry_count",
					Description: `(Optional) The number of attempts that the system will make to run a job using the exponential backoff procedure described by maxDoublings. Values greater than 5 and negative values are not allowed.`,
				},
				resource.Attribute{
					Name:        "max_retry_duration",
					Description: `(Optional) The time limit for retrying a failed job, measured from time when an execution was first attempted. If specified with retryCount, the job will be retried until both limits are reached. A duration in seconds with up to nine fractional digits, terminated by 's'.`,
				},
				resource.Attribute{
					Name:        "min_backoff_duration",
					Description: `(Optional) The minimum amount of time to wait before retrying a job after it fails. A duration in seconds with up to nine fractional digits, terminated by 's'.`,
				},
				resource.Attribute{
					Name:        "max_backoff_duration",
					Description: `(Optional) The maximum amount of time to wait before retrying a job after it fails. A duration in seconds with up to nine fractional digits, terminated by 's'.`,
				},
				resource.Attribute{
					Name:        "max_doublings",
					Description: `(Optional) The time between retries will double maxDoublings times. A job's retry interval starts at minBackoffDuration, then doubles maxDoublings times, then increases linearly, and finally retries retries at intervals of maxBackoffDuration up to retryCount times. The ` + "`" + `pubsub_target` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "topic_name",
					Description: `(Required) The name of the Cloud Pub/Sub topic to which messages will be published when a job is delivered. The topic name must be in the same format as required by PubSub's PublishRequest.name, for example projects/PROJECT_ID/topics/TOPIC_ID.`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `(Optional) The message payload for PubsubMessage. Pubsub message must contain either non-empty data, or at least one attribute.`,
				},
				resource.Attribute{
					Name:        "attributes",
					Description: `(Optional) Attributes for PubsubMessage. Pubsub message must contain either non-empty data, or at least one attribute. The ` + "`" + `app_engine_http_target` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "http_method",
					Description: `(Optional) Which HTTP method to use for the request.`,
				},
				resource.Attribute{
					Name:        "app_engine_routing",
					Description: `(Optional) App Engine Routing setting for the job. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "relative_uri",
					Description: `(Required) The relative URI. The relative URL must begin with "/" and must be a valid HTTP relative URL. It can contain a path, query string arguments, and \# fragments. If the relative URL is empty, then the root path "/" will be used. No spaces are allowed, and the maximum length allowed is 2083 characters`,
				},
				resource.Attribute{
					Name:        "body",
					Description: `(Optional) HTTP request body. A request body is allowed only if the HTTP method is POST or PUT. It will result in invalid argument error to set a body on a job with an incompatible HttpMethod.`,
				},
				resource.Attribute{
					Name:        "headers",
					Description: `(Optional) HTTP request headers. This map contains the header field names and values. Headers can be set when the job is created. The ` + "`" + `app_engine_routing` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Optional) App service. By default, the job is sent to the service which is the default service when the job is attempted.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) App version. By default, the job is sent to the version which is the default version when the job is attempted.`,
				},
				resource.Attribute{
					Name:        "instance",
					Description: `(Optional) App instance. By default, the job is sent to an instance which is available when the job is attempted. The ` + "`" + `http_target` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `(Required) The full URI path that the request will be sent to.`,
				},
				resource.Attribute{
					Name:        "http_method",
					Description: `(Optional) Which HTTP method to use for the request.`,
				},
				resource.Attribute{
					Name:        "body",
					Description: `(Optional) HTTP request body. A request body is allowed only if the HTTP method is POST, PUT, or PATCH. It is an error to set body on a job with an incompatible HttpMethod.`,
				},
				resource.Attribute{
					Name:        "headers",
					Description: `(Optional) This map contains the header field names and values. Repeated headers are not supported, but a header value can contain commas.`,
				},
				resource.Attribute{
					Name:        "oauth_token",
					Description: `(Optional) Contains information needed for generating an OAuth token. This type of authorization should be used when sending requests to a GCP endpoint. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "oidc_token",
					Description: `(Optional) Contains information needed for generating an OpenID Connect token. This type of authorization should be used when sending requests to third party endpoints or Cloud Run. Structure is documented below. The ` + "`" + `oauth_token` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "service_account_email",
					Description: `(Optional) Service account email to be used for generating OAuth token. The service account must be within the same project as the job.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) OAuth scope to be used for generating OAuth access token. If not specified, "https://www.googleapis.com/auth/cloud-platform" will be used. The ` + "`" + `oidc_token` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "service_account_email",
					Description: `(Optional) Service account email to be used for generating OAuth token. The service account must be within the same project as the job.`,
				},
				resource.Attribute{
					Name:        "audience",
					Description: `(Optional) Audience to be used when generating OIDC token. If not specified, the URI specified in target will be used. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Job can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_cloud_scheduler_job.default projects/{{project}}/locations/{{region}}/jobs/{{name}} $ terraform import google_cloud_scheduler_job.default {{project}}/{{region}}/{{name}} $ terraform import google_cloud_scheduler_job.default {{region}}/{{name}} $ terraform import google_cloud_scheduler_job.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_cloudbuild_trigger",
			Category:         "Google Cloud Build Resources",
			ShortDescription: `Configuration for an automated build in response to source repository changes.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"build",
				"cloudbuild",
				"trigger",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description of the trigger.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Whether the trigger is disabled or not. If true, the trigger will never result in a build.`,
				},
				resource.Attribute{
					Name:        "substitutions",
					Description: `(Optional) Substitutions data for Build resource.`,
				},
				resource.Attribute{
					Name:        "filename",
					Description: `(Optional) Path, from the source root, to a file whose contents is used for the template. Either a filename or build template must be provided.`,
				},
				resource.Attribute{
					Name:        "ignored_files",
					Description: `(Optional) ignoredFiles and includedFiles are file glob matches using http://godoc/pkg/path/filepath#Match extended with support for ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "included_files",
					Description: `(Optional) ignoredFiles and includedFiles are file glob matches using http://godoc/pkg/path/filepath#Match extended with support for ` + "`" + ``,
				},
				resource.Attribute{
					Name:        "trigger_template",
					Description: `(Optional) Template describing the types of source changes to trigger a build. Branch and tag names in trigger templates are interpreted as regular expressions. Any branch or tag change that matches that regular expression will trigger a build. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "build",
					Description: `(Optional) Contents of the build template. Either a filename or build template must be provided. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `trigger_template` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) ID of the project that owns the Cloud Source Repository. If omitted, the project ID requesting the build is assumed.`,
				},
				resource.Attribute{
					Name:        "repo_name",
					Description: `(Optional) Name of the Cloud Source Repository. If omitted, the name "default" is assumed.`,
				},
				resource.Attribute{
					Name:        "dir",
					Description: `(Optional) Directory, relative to the source root, in which to run the build. This must be a relative path. If a step's dir is specified and is an absolute path, this value is ignored for that step's execution.`,
				},
				resource.Attribute{
					Name:        "branch_name",
					Description: `(Optional) Name of the branch to build. Exactly one a of branch name, tag, or commit SHA must be provided.`,
				},
				resource.Attribute{
					Name:        "tag_name",
					Description: `(Optional) Name of the tag to build. Exactly one of a branch name, tag, or commit SHA must be provided.`,
				},
				resource.Attribute{
					Name:        "commit_sha",
					Description: `(Optional) Explicit commit SHA to build. Exactly one of a branch name, tag, or commit SHA must be provided. The ` + "`" + `build` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags for annotation of a Build. These are not docker tags.`,
				},
				resource.Attribute{
					Name:        "images",
					Description: `(Optional) A list of images to be pushed upon the successful completion of all build steps. The images are pushed using the builder service account's credentials. The digests of the pushed images will be stored in the Build resource's results field. If any of the images fail to be pushed, the build status is marked FAILURE.`,
				},
				resource.Attribute{
					Name:        "step",
					Description: `(Optional) The operations to be performed on the workspace. Structure is documented below. The ` + "`" + `step` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the container image that will run this particular build step. If the image is available in the host's Docker daemon's cache, it will be run directly. If not, the host will attempt to pull the image first, using the builder service account's credentials if necessary. The Docker daemon's cache will already have the latest versions of all of the officially supported build steps (https://github.com/GoogleCloudPlatform/cloud-builders). The Docker daemon will also have cached many of the layers for some popular images, like "ubuntu", "debian", but they will be refreshed at the time you attempt to use them. If you built an image in a previous build step, it will be stored in the host's Docker daemon's cache and is available to use as the name for a later build step.`,
				},
				resource.Attribute{
					Name:        "args",
					Description: `(Optional) A list of arguments that will be presented to the step when it is started. If the image used to run the step's container has an entrypoint, the args are used as arguments to that entrypoint. If the image does not define an entrypoint, the first element in args is used as the entrypoint, and the remainder will be used as arguments.`,
				},
				resource.Attribute{
					Name:        "env",
					Description: `(Optional) A list of environment variable definitions to be used when running a step. The elements are of the form "KEY=VALUE" for the environment variable "KEY" being given the value "VALUE".`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Unique identifier for this build step, used in ` + "`" + `wait_for` + "`" + ` to reference this build step as a dependency.`,
				},
				resource.Attribute{
					Name:        "entrypoint",
					Description: `(Optional) Entrypoint to be used instead of the build step image's default entrypoint. If unset, the image's default entrypoint is used`,
				},
				resource.Attribute{
					Name:        "dir",
					Description: `(Optional) Working directory to use when running this step's container. If this value is a relative path, it is relative to the build's working directory. If this value is absolute, it may be outside the build's working directory, in which case the contents of the path may not be persisted across build step executions, unless a ` + "`" + `volume` + "`" + ` for that path is specified. If the build specifies a ` + "`" + `RepoSource` + "`" + ` with ` + "`" + `dir` + "`" + ` and a step with a ` + "`" + `dir` + "`" + `, which specifies an absolute path, the ` + "`" + `RepoSource` + "`" + ` ` + "`" + `dir` + "`" + ` is ignored for the step's execution.`,
				},
				resource.Attribute{
					Name:        "secret_env",
					Description: `(Optional) A list of environment variables which are encrypted using a Cloud Key Management Service crypto key. These values must be specified in the build's ` + "`" + `Secret` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Time limit for executing this build step. If not defined, the step has no time limit and will be allowed to continue to run until either it completes or the build itself times out.`,
				},
				resource.Attribute{
					Name:        "timing",
					Description: `(Optional) Output only. Stores timing information for executing this build step.`,
				},
				resource.Attribute{
					Name:        "volumes",
					Description: `(Optional) List of volumes to mount into the build step. Each volume is created as an empty volume prior to execution of the build step. Upon completion of the build, volumes and their contents are discarded. Using a named volume in only one step is not valid as it is indicative of a build request with an incorrect configuration. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "wait_for",
					Description: `(Optional) The ID(s) of the step(s) that this build step depends on. This build step will not start until all the build steps in ` + "`" + `wait_for` + "`" + ` have completed successfully. If ` + "`" + `wait_for` + "`" + ` is empty, this build step will start when all previous build steps in the ` + "`" + `Build.Steps` + "`" + ` list have completed successfully. The ` + "`" + `volumes` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the volume to mount. Volume names must be unique per build step and must be valid names for Docker volumes. Each named volume must be used by at least two build steps.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) Path at which to mount the volume. Paths must be absolute and cannot conflict with other volume paths on the same build step or with certain reserved volume paths. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "trigger_id",
					Description: `The unique identifier for the trigger.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Time when the trigger was created. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Trigger can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_cloudbuild_trigger.default projects/{{project}}/triggers/{{trigger_id}} $ terraform import google_cloudbuild_trigger.default {{project}}/{{trigger_id}} $ terraform import google_cloudbuild_trigger.default {{trigger_id}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "trigger_id",
					Description: `The unique identifier for the trigger.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Time when the trigger was created. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Trigger can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_cloudbuild_trigger.default projects/{{project}}/triggers/{{trigger_id}} $ terraform import google_cloudbuild_trigger.default {{project}}/{{trigger_id}} $ terraform import google_cloudbuild_trigger.default {{trigger_id}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_cloudfunctions_function",
			Category:         "Google Cloud Functions Resources",
			ShortDescription: `Creates a new Cloud Function.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"functions",
				"cloudfunctions",
				"function",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A user-defined name of the function. Function names must be unique globally.`,
				},
				resource.Attribute{
					Name:        "runtime",
					Description: `(Optional) The runtime in which the function is going to run. One of ` + "`" + `"nodejs6"` + "`" + `, ` + "`" + `"nodejs8"` + "`" + `, ` + "`" + `"nodejs10"` + "`" + `, ` + "`" + `"python37"` + "`" + `, ` + "`" + `"go111"` + "`" + `. If empty, defaults to ` + "`" + `"nodejs6"` + "`" + `. It's recommended that you override the default, as ` + "`" + `"nodejs6"` + "`" + ` is deprecated. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the function.`,
				},
				resource.Attribute{
					Name:        "available_memory_mb",
					Description: `(Optional) Memory (in MB), available to the function. Default value is 256MB. Allowed values are: 128MB, 256MB, 512MB, 1024MB, and 2048MB.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.`,
				},
				resource.Attribute{
					Name:        "entry_point",
					Description: `(Optional) Name of the function that will be executed when the Google Cloud Function is triggered.`,
				},
				resource.Attribute{
					Name:        "event_trigger",
					Description: `(Optional) A source that fires events in response to a condition in another service. Structure is documented below. Cannot be used with ` + "`" + `trigger_http` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "trigger_http",
					Description: `(Optional) Boolean variable. Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as ` + "`" + `https_trigger_url` + "`" + `. Cannot be used with ` + "`" + `trigger_bucket` + "`" + ` and ` + "`" + `trigger_topic` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs to assign to the function.`,
				},
				resource.Attribute{
					Name:        "service_account_email",
					Description: `(Optional) If provided, the self-provided service account to run the function with.`,
				},
				resource.Attribute{
					Name:        "environment_variables",
					Description: `(Optional) A set of key/value environment variable pairs to assign to the function.`,
				},
				resource.Attribute{
					Name:        "vpc_connector",
					Description: `(Optional) The VPC Network Connector that this cloud function can connect to. It can be either the fully-qualified URI, or the short name of the network connector resource. The format of this field is ` + "`" + `projects/`,
				},
				resource.Attribute{
					Name:        "source_archive_bucket",
					Description: `(Optional) The GCS bucket containing the zip archive which contains the function.`,
				},
				resource.Attribute{
					Name:        "source_archive_object",
					Description: `(Optional) The source archive object (file) in archive bucket.`,
				},
				resource.Attribute{
					Name:        "source_repository",
					Description: `(Optional) Represents parameters related to source repository where a function is hosted. Cannot be set alongside ` + "`" + `source_archive_bucket` + "`" + ` or ` + "`" + `source_archive_object` + "`" + `. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "max_instances",
					Description: `(Optional) The limit on the maximum number of function instances that may coexist at a given time. The ` + "`" + `event_trigger` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "event_type",
					Description: `(Required) The type of event to observe. For example: ` + "`" + `"google.storage.object.finalize"` + "`" + `. See the documentation on [calling Cloud Functions](https://cloud.google.com/functions/docs/calling/) for a full reference. Cloud Storage, Cloud Pub/Sub and Cloud Firestore triggers are supported at this time. Legacy triggers are supported, such as ` + "`" + `"providers/cloud.storage/eventTypes/object.change"` + "`" + `, ` + "`" + `"providers/cloud.pubsub/eventTypes/topic.publish"` + "`" + ` and ` + "`" + `"providers/cloud.firestore/eventTypes/document.create"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `(Required) Required. The name or partial URI of the resource from which to observe events. For example, ` + "`" + `"myBucket"` + "`" + ` or ` + "`" + `"projects/my-project/topics/my-topic"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "failure_policy",
					Description: `(Optional) Specifies policy for failed executions. Structure is documented below. The ` + "`" + `failure_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "retry",
					Description: `(Required) Whether the function should be retried on failure. Defaults to ` + "`" + `false` + "`" + `. The ` + "`" + `source_repository` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The URL pointing to the hosted repository where the function is defined. There are supported Cloud Source Repository URLs in the following formats:`,
				},
				resource.Attribute{
					Name:        "https_trigger_url",
					Description: `URL which triggers function execution. Returned only if ` + "`" + `trigger_http` + "`" + ` is used.`,
				},
				resource.Attribute{
					Name:        "source_repository.0.deployed_url",
					Description: `The URL pointing to the hosted repository where the function was defined at the time of deployment.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `Project of the function. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region of function. Currently can be only "us-central1". If it is not provided, the provider region is used. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 5 minutes. - ` + "`" + `update` + "`" + ` - Default is 5 minutes. - ` + "`" + `delete` + "`" + ` - Default is 5 minutes. ## Import Functions can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_cloudfunctions_function.default function-test ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "https_trigger_url",
					Description: `URL which triggers function execution. Returned only if ` + "`" + `trigger_http` + "`" + ` is used.`,
				},
				resource.Attribute{
					Name:        "source_repository.0.deployed_url",
					Description: `The URL pointing to the hosted repository where the function was defined at the time of deployment.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `Project of the function. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region of function. Currently can be only "us-central1". If it is not provided, the provider region is used. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 5 minutes. - ` + "`" + `update` + "`" + ` - Default is 5 minutes. - ` + "`" + `delete` + "`" + ` - Default is 5 minutes. ## Import Functions can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_cloudfunctions_function.default function-test ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_cloudiot_registry",
			Category:         "Google Cloud IoT Core",
			ShortDescription: `Creates a device registry in Google's Cloud IoT Core platform`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"iot",
				"core",
				"cloudiot",
				"registry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the resource, required by device registry. Changing this forces a new resource to be created. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The Region in which the created address should reside. If it is not provided, the provider region is used.`,
				},
				resource.Attribute{
					Name:        "event_notification_config",
					Description: `(Optional) A PubSub topics to publish device events. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "state_notification_config",
					Description: `(Optional) A PubSub topic to publish device state updates. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "mqtt_config",
					Description: `(Optional) Activate or deactivate MQTT. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "http_config",
					Description: `(Optional) Activate or deactivate HTTP. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `(Optional) List of public key certificates to authenticate devices. Structure is documented below. The ` + "`" + `event_notification_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "pubsub_topic_name",
					Description: `(Required) PubSub topic name to publish device events. The ` + "`" + `state_notification_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "pubsub_topic_name",
					Description: `(Required) PubSub topic name to publish device state updates. The ` + "`" + `mqtt_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "mqtt_enabled_state",
					Description: `(Required) The field allows ` + "`" + `MQTT_ENABLED` + "`" + ` or ` + "`" + `MQTT_DISABLED` + "`" + `. The ` + "`" + `http_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "http_enabled_state",
					Description: `(Required) The field allows ` + "`" + `HTTP_ENABLED` + "`" + ` or ` + "`" + `HTTP_DISABLED` + "`" + `. The ` + "`" + `credentials` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "public_key_certificate",
					Description: `(Required) The certificate format and data. The ` + "`" + `public_key_certificate` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Required) The field allows only ` + "`" + `X509_CERTIFICATE_PEM` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `(Required) The certificate data. ## Attributes Reference Only the arguments listed above are exposed as attributes. ## Import A device registry can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_cloudiot_registry.default-registry projects/{project}/locations/{region}/registries/{name} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_composer_environment",
			Category:         "Google Cloud Composer Resources",
			ShortDescription: `An environment for running orchestration tasks.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"composer",
				"environment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the environment - - -`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `(Optional) Configuration parameters for this environment Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) User-defined labels for this environment. The labels map can contain no more than 64 entries. Entries of the labels map are UTF8 strings that comply with the following restrictions: Label keys must be between 1 and 63 characters long and must conform to the following regular expression: ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The location or Compute Engine region for the environment.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `(Optional) The number of nodes in the Kubernetes Engine cluster that will be used to run this environment.`,
				},
				resource.Attribute{
					Name:        "node_config",
					Description: `(Optional) The configuration used for the Kubernetes Engine cluster. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "software_config",
					Description: `(Optional) The configuration settings for software inside the environment. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "private_environment_config",
					Description: `(Optional) The configuration used for the Private IP Cloud Composer environment. Structure is documented below. The ` + "`" + `node_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The Compute Engine zone in which to deploy the VMs running the Apache Airflow software, specified as the zone name or relative resource name (e.g. "projects/{project}/zones/{zone}"). Must belong to the enclosing environment's project and region.`,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `(Optional) The Compute Engine machine type used for cluster instances, specified as a name or relative resource name. For example: "projects/{project}/zones/{zone}/machineTypes/{machineType}". Must belong to the enclosing environment's project and region/zone.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) The Compute Engine network to be used for machine communications, specified as a self-link, relative resource name (e.g. "projects/{project}/global/networks/{network}"), by name. The network must belong to the environment's project. If unspecified, the "default" network ID in the environment's project is used. If a Custom Subnet Network is provided, subnetwork must also be provided.`,
				},
				resource.Attribute{
					Name:        "subnetwork",
					Description: `(Optional) The Compute Engine subnetwork to be used for machine communications, , specified as a self-link, relative resource name (e.g. "projects/{project}/regions/{region}/subnetworks/{subnetwork}"), or by name. If subnetwork is provided, network must also be provided and the subnetwork must belong to the enclosing environment's project and region.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `(Optional) The disk size in GB used for node VMs. Minimum size is 20GB. If unspecified, defaults to 100GB. Cannot be updated.`,
				},
				resource.Attribute{
					Name:        "oauth_scopes",
					Description: `(Optional) The set of Google API scopes to be made available on all node VMs. Cannot be updated. If empty, defaults to ` + "`" + `["https://www.googleapis.com/auth/cloud-platform"]` + "`" + ``,
				},
				resource.Attribute{
					Name:        "service_account",
					Description: `(Optional) The Google Cloud Platform Service Account to be used by the node VMs. If a service account is not specified, the "default" Compute Engine service account is used. Cannot be updated. If given, note that the service account must have ` + "`" + `roles/composer.worker` + "`" + ` for any GCP resources created under the Cloud Composer Environment.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The list of instance tags applied to all node VMs. Tags are used to identify valid sources or targets for network firewalls. Each tag within the list must comply with RFC1035. Cannot be updated.`,
				},
				resource.Attribute{
					Name:        "ip_allocation_policy",
					Description: `(Optional) Configuration for controlling how IPs are allocated in the GKE cluster. Structure is documented below. Cannot be updated. The ` + "`" + `software_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "airflow_config_overrides",
					Description: `(Optional) Apache Airflow configuration properties to override. Property keys contain the section and property names, separated by a hyphen, for example "core-dags_are_paused_at_creation". Section names must not contain hyphens ("-"), opening square brackets ("["), or closing square brackets ("]"). The property name must not be empty and cannot contain "=" or ";". Section and property names cannot contain characters: "." Apache Airflow configuration property names must be written in snake_case. Property values can contain any character, and can be written in any lower/upper case format. Certain Apache Airflow configuration property values are [blacklisted](https://cloud.google.com/composer/docs/concepts/airflow-configurations#airflow_configuration_blacklists), and cannot be overridden.`,
				},
				resource.Attribute{
					Name:        "pypi_packages",
					Description: `(Optional) Custom Python Package Index (PyPI) packages to be installed in the environment. Keys refer to the lowercase package name (e.g. "numpy"). Values are the lowercase extras and version specifier (e.g. "==1.12.0", "[devel,gcp_api]", "[devel]>=1.8.2, <1.9.2"). To specify a package without pinning it to a version specifier, use the empty string as the value.`,
				},
				resource.Attribute{
					Name:        "env_variables",
					Description: `(Optional) Additional environment variables to provide to the Apache Airflow scheduler, worker, and webserver processes. Environment variable names must match the regular expression ` + "`" + `[a-zA-Z_][a-zA-Z0-9_]`,
				},
				resource.Attribute{
					Name:        "enable_private_endpoint",
					Description: `If true, access to the public endpoint of the GKE cluster is denied.`,
				},
				resource.Attribute{
					Name:        "master_ipv4_cidr_block",
					Description: `(Optional) The IP range in CIDR notation to use for the hosted master network. This range is used for assigning internal IP addresses to the cluster master or set of masters and to the internal load balancer virtual IP. This range must not overlap with any other ranges in use within the cluster's network. If left blank, the default value of '172.16.0.0/28' is used. The ` + "`" + `ip_allocation_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "use_ip_aliases",
					Description: `(Optional) Whether or not to enable Alias IPs in the GKE cluster. If true, a VPC-native cluster is created. Defaults to true if the ` + "`" + `ip_allocation_block` + "`" + ` is present in config.`,
				},
				resource.Attribute{
					Name:        "cluster_secondary_range_name",
					Description: `(Optional) The name of the cluster's secondary range used to allocate IP addresses to pods. Specify either ` + "`" + `cluster_secondary_range_name` + "`" + ` or ` + "`" + `cluster_ipv4_cidr_block` + "`" + ` but not both. This field is applicable only when ` + "`" + `use_ip_aliases` + "`" + ` is true.`,
				},
				resource.Attribute{
					Name:        "services_secondary_range_name",
					Description: `(Optional) The name of the services' secondary range used to allocate IP addresses to the cluster. Specify either ` + "`" + `services_secondary_range_name` + "`" + ` or ` + "`" + `services_ipv4_cidr_block` + "`" + ` but not both. This field is applicable only when ` + "`" + `use_ip_aliases` + "`" + ` is true.`,
				},
				resource.Attribute{
					Name:        "cluster_ipv4_cidr_block",
					Description: `(Optional) The IP address range used to allocate IP addresses to pods in the cluster. Set to blank to have GKE choose a range with the default size. Set to /netmask (e.g. /14) to have GKE choose a range with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use. Specify either ` + "`" + `cluster_secondary_range_name` + "`" + ` or ` + "`" + `cluster_ipv4_cidr_block` + "`" + ` but not both.`,
				},
				resource.Attribute{
					Name:        "services_ipv4_cidr_block",
					Description: `(Optional) The IP address range used to allocate IP addresses in this cluster. Set to blank to have GKE choose a range with the default size. Set to /netmask (e.g. /14) to have GKE choose a range with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use. Specify either ` + "`" + `services_secondary_range_name` + "`" + ` or ` + "`" + `services_ipv4_cidr_block` + "`" + ` but not both. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "config.0.gke_cluster",
					Description: `The Kubernetes Engine cluster used to run this environment.`,
				},
				resource.Attribute{
					Name:        "config.0.dag_gcs_prefix",
					Description: `The Cloud Storage prefix of the DAGs for this environment. Although Cloud Storage objects reside in a flat namespace, a hierarchical file tree can be simulated using '/'-delimited object name prefixes. DAG objects for this environment reside in a simulated directory with this prefix.`,
				},
				resource.Attribute{
					Name:        "config.0.airflow_uri",
					Description: `The URI of the Apache Airflow Web UI hosted within this environment. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 60 minutes. - ` + "`" + `update` + "`" + ` - Default is 60 minutes. - ` + "`" + `delete` + "`" + ` - Default is 6 minutes. ## Import Environment can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_composer_environment.default projects/{{project}}/locations/{{region}}/environments/{{name}} $ terraform import google_composer_environment.default {{project}}/{{region}}/{{name}} $ terraform import google_composer_environment.default {{name}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "config.0.gke_cluster",
					Description: `The Kubernetes Engine cluster used to run this environment.`,
				},
				resource.Attribute{
					Name:        "config.0.dag_gcs_prefix",
					Description: `The Cloud Storage prefix of the DAGs for this environment. Although Cloud Storage objects reside in a flat namespace, a hierarchical file tree can be simulated using '/'-delimited object name prefixes. DAG objects for this environment reside in a simulated directory with this prefix.`,
				},
				resource.Attribute{
					Name:        "config.0.airflow_uri",
					Description: `The URI of the Apache Airflow Web UI hosted within this environment. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 60 minutes. - ` + "`" + `update` + "`" + ` - Default is 60 minutes. - ` + "`" + `delete` + "`" + ` - Default is 6 minutes. ## Import Environment can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_composer_environment.default projects/{{project}}/locations/{{region}}/environments/{{name}} $ terraform import google_composer_environment.default {{project}}/{{region}}/{{name}} $ terraform import google_composer_environment.default {{name}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_address",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Represents an Address resource.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"address",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Optional) The static external IP address represented by this resource. Only IPv4 is supported. An address may only be specified for INTERNAL address types. The IP address must be inside the specified subnetwork, if any.`,
				},
				resource.Attribute{
					Name:        "address_type",
					Description: `(Optional) The type of address to reserve, either INTERNAL or EXTERNAL. If unspecified, defaults to EXTERNAL.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource.`,
				},
				resource.Attribute{
					Name:        "network_tier",
					Description: `(Optional) The networking tier used for configuring this address. This field can take the following values: PREMIUM or STANDARD. If this field is not specified, it is assumed to be PREMIUM.`,
				},
				resource.Attribute{
					Name:        "subnetwork",
					Description: `(Optional) The URL of the subnetwork in which to reserve the address. If an IP address is specified, it must be within the subnetwork's IP range. This field can only be used with INTERNAL type with GCE_ENDPOINT/DNS_RESOLVER purposes.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) Labels to apply to this address. A list of key->value pairs.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The Region in which the created address should reside. If it is not provided, the provider region is used.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `The URLs of the resources that are using this address.`,
				},
				resource.Attribute{
					Name:        "label_fingerprint",
					Description: `The fingerprint used for optimistic locking of this resource. Used internally during updates.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The IP of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Address can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_address.default projects/{{project}}/regions/{{region}}/addresses/{{name}} $ terraform import google_compute_address.default {{project}}/{{region}}/{{name}} $ terraform import google_compute_address.default {{region}}/{{name}} $ terraform import google_compute_address.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `The URLs of the resources that are using this address.`,
				},
				resource.Attribute{
					Name:        "label_fingerprint",
					Description: `The fingerprint used for optimistic locking of this resource. Used internally during updates.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The IP of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Address can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_address.default projects/{{project}}/regions/{{region}}/addresses/{{name}} $ terraform import google_compute_address.default {{project}}/{{region}}/{{name}} $ terraform import google_compute_address.default {{region}}/{{name}} $ terraform import google_compute_address.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_attached_disk",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Resource that allows attaching existing persistent disks to compute instances.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"attached",
				"disk",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance",
					Description: `(Required) ` + "`" + `name` + "`" + ` or ` + "`" + `self_link` + "`" + ` of the compute instance that the disk will be attached to. If the ` + "`" + `self_link` + "`" + ` is provided then ` + "`" + `zone` + "`" + ` and ` + "`" + `project` + "`" + ` are extracted from the self link. If only the name is used then ` + "`" + `zone` + "`" + ` and ` + "`" + `project` + "`" + ` must be defined as properties on the resource or provider.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `(Required) ` + "`" + `name` + "`" + ` or ` + "`" + `self_link` + "`" + ` of the disk that will be attached. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project that the referenced compute instance is a part of. If ` + "`" + `instance` + "`" + ` is referenced by its ` + "`" + `self_link` + "`" + ` the project defined in the link will take precedence.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) The zone that the referenced compute instance is located within. If ` + "`" + `instance` + "`" + ` is referenced by its ` + "`" + `self_link` + "`" + ` the zone defined in the link will take precedence.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(Optional) Specifies a unique device name of your choice that is reflected into the /dev/disk/by-id/google-`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) The mode in which to attach this disk, either READ_WRITE or READ_ONLY. If not specified, the default is to attach the disk in READ_WRITE mode. Possible values: "READ_ONLY" "READ_WRITE" ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 5 minutes. - ` + "`" + `delete` + "`" + ` - Default is 5 minutes. ## Import Attached Disk can be imported the following ways: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_disk.default projects/{{project}}/zones/{{zone}}/disks/{{instance.name}}:{{disk.name}} $ terraform import google_compute_disk.default {{project}}/{{zone}}/{{instance.name}}:{{disk.name}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_autoscaler",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Represents an Autoscaler resource.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"autoscaler",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource. The name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "autoscaling_policy",
					Description: `(Required) The configuration parameters for the autoscaling algorithm. You can define one or more of the policies for an autoscaler: cpuUtilization, customMetricUtilizations, and loadBalancingUtilization. If none of these are specified, the default will be to autoscale based on cpuUtilization to 0.6 or 60%. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) URL of the managed instance group that this autoscaler will scale. The ` + "`" + `autoscaling_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "min_replicas",
					Description: `(Required) The minimum number of replicas that the autoscaler can scale down to. This cannot be less than 0. If not provided, autoscaler will choose a default value depending on maximum number of instances allowed.`,
				},
				resource.Attribute{
					Name:        "max_replicas",
					Description: `(Required) The maximum number of instances that the autoscaler can scale up to. This is required when creating or updating an autoscaler. The maximum number of replicas should not be lower than minimal number of replicas.`,
				},
				resource.Attribute{
					Name:        "cooldown_period",
					Description: `(Optional) The number of seconds that the autoscaler should wait before it starts collecting information from a new instance. This prevents the autoscaler from collecting information when the instance is initializing, during which the collected usage would not be reliable. The default time autoscaler waits is 60 seconds. Virtual machine initialization times might vary because of numerous factors. We recommend that you test how long an instance may take to initialize. To do this, create an instance and time the startup process.`,
				},
				resource.Attribute{
					Name:        "cpu_utilization",
					Description: `(Optional) Defines the CPU utilization policy that allows the autoscaler to scale based on the average CPU utilization of a managed instance group. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(Optional) Defines the CPU utilization policy that allows the autoscaler to scale based on the average CPU utilization of a managed instance group. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "load_balancing_utilization",
					Description: `(Optional) Configuration parameters of autoscaling based on a load balancer. Structure is documented below. The ` + "`" + `cpu_utilization` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) The target CPU utilization that the autoscaler should maintain. Must be a float value in the range (0, 1]. If not specified, the default is 0.6. If the CPU level is below the target utilization, the autoscaler scales down the number of instances until it reaches the minimum number of instances you specified or until the average CPU of your instances reaches the target utilization. If the average CPU is above the target utilization, the autoscaler scales up until it reaches the maximum number of instances you specified or until the average utilization reaches the target utilization. The ` + "`" + `metric` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The identifier (type) of the Stackdriver Monitoring metric. The metric cannot have negative values. The metric must have a value type of INT64 or DOUBLE.`,
				},
				resource.Attribute{
					Name:        "single_instance_assignment",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) If scaling is based on a per-group metric value that represents the total amount of work to be done or resource usage, set this value to an amount assigned for a single instance of the scaled group. The autoscaler will keep the number of instances proportional to the value of this metric, the metric itself should not change value due to group resizing. For example, a good metric to use with the target is ` + "`" + `pubsub.googleapis.com/subscription/num_undelivered_messages` + "`" + ` or a custom metric exporting the total number of requests coming to your instances. A bad example would be a metric exporting an average or median latency, since this value can't include a chunk assignable to a single instance, it could be better used with utilization_target instead.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Optional) The target value of the metric that autoscaler should maintain. This must be a positive value. A utilization metric scales number of virtual machines handling requests to increase or decrease proportionally to the metric. For example, a good metric to use as a utilizationTarget is www.googleapis.com/compute/instance/network/received_bytes_count. The autoscaler will work to keep this value constant for each of the instances.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Defines how target utilization value is expressed for a Stackdriver Monitoring metric. Either GAUGE, DELTA_PER_SECOND, or DELTA_PER_MINUTE.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) A filter string to be used as the filter string for a Stackdriver Monitoring TimeSeries.list API call. This filter is used to select a specific TimeSeries for the purpose of autoscaling and to determine whether the metric is exporting per-instance or per-group data. You can only use the AND operator for joining selectors. You can only use direct equality comparison operator (=) without any functions for each selector. You can specify the metric in both the filter string and in the metric field. However, if specified in both places, the metric must be identical. The monitored resource type determines what kind of values are expected for the metric. If it is a gce_instance, the autoscaler expects the metric to include a separate TimeSeries for each instance in a group. In such a case, you cannot filter on resource labels. If the resource type is any other value, the autoscaler expects this metric to contain values that apply to the entire autoscaled instance group and resource label filtering can be performed to point autoscaler at the correct TimeSeries to scale upon. This is called a per-group metric for the purpose of autoscaling. If not specified, the type defaults to gce_instance. You should provide a filter that is selective enough to pick just one TimeSeries for the autoscaled group or for each of the instances (if you are using gce_instance resource type). If multiple TimeSeries are returned upon the query execution, the autoscaler will sum their respective values to obtain its scaling value. The ` + "`" + `load_balancing_utilization` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) Fraction of backend capacity utilization (set in HTTP(s) load balancing configuration) that autoscaler should maintain. Must be a positive float value. If not defined, the default is 0.8. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) URL of the zone where the instance group resides.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Autoscaler can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_autoscaler.default projects/{{project}}/zones/{{zone}}/autoscalers/{{name}} $ terraform import google_compute_autoscaler.default {{project}}/{{zone}}/{{name}} $ terraform import google_compute_autoscaler.default {{zone}}/{{name}} $ terraform import google_compute_autoscaler.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Autoscaler can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_autoscaler.default projects/{{project}}/zones/{{zone}}/autoscalers/{{name}} $ terraform import google_compute_autoscaler.default {{project}}/{{zone}}/{{name}} $ terraform import google_compute_autoscaler.default {{zone}}/{{name}} $ terraform import google_compute_autoscaler.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_backend_bucket",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Backend buckets allow you to use Google Cloud Storage buckets with HTTP(S) load balancing.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"backend",
				"bucket",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket_name",
					Description: `(Required) Cloud Storage bucket name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "cdn_policy",
					Description: `(Optional) Cloud CDN configuration for this Backend Bucket. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional textual description of the resource; provided by the client when the resource is created.`,
				},
				resource.Attribute{
					Name:        "enable_cdn",
					Description: `(Optional) If true, enable Cloud CDN for this BackendBucket.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `cdn_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "signed_url_cache_max_age_sec",
					Description: `(Optional) Maximum number of seconds the response to a signed URL request will be considered fresh. Defaults to 1hr (3600s). After this time period, the response will be revalidated before being served. When serving responses to signed URL requests, Cloud CDN will internally behave as though all responses from this backend had a "Cache-Control: public, max-age=[TTL]" header, regardless of any existing Cache-Control header. The actual headers served in responses will not be altered. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import BackendBucket can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_backend_bucket.default projects/{{project}}/global/backendBuckets/{{name}} $ terraform import google_compute_backend_bucket.default {{project}}/{{name}} $ terraform import google_compute_backend_bucket.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import BackendBucket can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_backend_bucket.default projects/{{project}}/global/backendBuckets/{{name}} $ terraform import google_compute_backend_bucket.default {{project}}/{{name}} $ terraform import google_compute_backend_bucket.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_backend_bucket_signed_url_key",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `A key for signing Cloud CDN signed URLs for BackendBuckets.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"backend",
				"bucket",
				"signed",
				"url",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the signed URL key.`,
				},
				resource.Attribute{
					Name:        "key_value",
					Description: `(Required) 128-bit key value used for signing the URL. The key value must be a valid RFC 4648 Section 5 base64url encoded string.`,
				},
				resource.Attribute{
					Name:        "backend_bucket",
					Description: `(Required) The backend bucket this signed URL key belongs. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import BackendBucketSignedUrlKey can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_backend_bucket_signed_url_key.default projects/{{project}}/global/backendBuckets/{{backend_bucket}}/{{name}} $ terraform import google_compute_backend_bucket_signed_url_key.default {{project}}/{{backend_bucket}}/{{name}} $ terraform import google_compute_backend_bucket_signed_url_key.default {{backend_bucket}}/{{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_backend_service",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `A Backend Service defines a group of virtual machines that will serve traffic for load balancing.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"backend",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "health_checks",
					Description: `(Required) The set of URLs to the HttpHealthCheck or HttpsHealthCheck resource for health checking this BackendService. Currently at most one health check can be specified, and a health check is required. For internal load balancing, a URL to a HealthCheck resource must be specified instead.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "affinity_cookie_ttl_sec",
					Description: `(Optional) Lifetime of cookies in seconds if session_affinity is GENERATED_COOKIE. If set to 0, the cookie is non-persistent and lasts only until the end of the browser session (or equivalent). The maximum allowed value for TTL is one day. When the load balancing scheme is INTERNAL, this field is not used.`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) The set of backends that serve this BackendService. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "cdn_policy",
					Description: `(Optional) Cloud CDN configuration for this BackendService. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "connection_draining_timeout_sec",
					Description: `(Optional) Time for which instance will be drained (not accept new connections, but still work to finish started).`,
				},
				resource.Attribute{
					Name:        "custom_request_headers",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) Headers that the HTTP/S load balancer should add to proxied requests.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource.`,
				},
				resource.Attribute{
					Name:        "enable_cdn",
					Description: `(Optional) If true, enable Cloud CDN for this BackendService.`,
				},
				resource.Attribute{
					Name:        "iap",
					Description: `(Optional) Settings for enabling Cloud Identity Aware Proxy Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "load_balancing_scheme",
					Description: `(Optional) Indicates whether the backend service will be used with internal or external load balancing. A backend service created for one type of load balancing cannot be used with the other. Must be ` + "`" + `EXTERNAL` + "`" + ` or ` + "`" + `INTERNAL_SELF_MANAGED` + "`" + ` for a global backend service. Defaults to ` + "`" + `EXTERNAL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port_name",
					Description: `(Optional) Name of backend port. The same name should appear in the instance groups referenced by this service. Required when the load balancing scheme is EXTERNAL.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The protocol this BackendService uses to communicate with backends. Possible values are HTTP, HTTPS, HTTP2, TCP, and SSL. The default is HTTP.`,
				},
				resource.Attribute{
					Name:        "security_policy",
					Description: `(Optional) The security policy associated with this backend service.`,
				},
				resource.Attribute{
					Name:        "session_affinity",
					Description: `(Optional) Type of session affinity to use. The default is NONE. When the load balancing scheme is EXTERNAL, can be NONE, CLIENT_IP, or GENERATED_COOKIE. When the protocol is UDP, this field is not used.`,
				},
				resource.Attribute{
					Name:        "timeout_sec",
					Description: `(Optional) How many seconds to wait for the backend before considering it a failed request. Default is 30 seconds. Valid range is [1, 86400].`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `backend` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "balancing_mode",
					Description: `(Optional) Specifies the balancing mode for this backend. For global HTTP(S) or TCP/SSL load balancing, the default is UTILIZATION. Valid values are UTILIZATION, RATE (for HTTP(S)) and CONNECTION (for TCP/SSL).`,
				},
				resource.Attribute{
					Name:        "capacity_scaler",
					Description: `(Optional) A multiplier applied to the group's maximum servicing capacity (based on UTILIZATION, RATE or CONNECTION). Default value is 1, which means the group will serve up to 100% of its configured capacity (depending on balancingMode). A setting of 0 means the group is completely drained, offering 0% of its available Capacity. Valid range is [0.0,1.0].`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource. Provide this property when you create the resource.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) The fully-qualified URL of an Instance Group or Network Endpoint Group resource. In case of instance group this defines the list of instances that serve traffic. Member virtual machine instances from each instance group must live in the same zone as the instance group itself. No two backends in a backend service are allowed to use same Instance Group resource. For Network Endpoint Groups this defines list of endpoints. All endpoints of Network Endpoint Group must be hosted on instances located in the same zone as the Network Endpoint Group. Backend service can not contain mix of Instance Group and Network Endpoint Group backends. Note that you must specify an Instance Group or Network Endpoint Group resource using the fully-qualified URL, rather than a partial URL.`,
				},
				resource.Attribute{
					Name:        "max_connections",
					Description: `(Optional) The max number of simultaneous connections for the group. Can be used with either CONNECTION or UTILIZATION balancing modes. For CONNECTION mode, either maxConnections or one of maxConnectionsPerInstance or maxConnectionsPerEndpoint, as appropriate for group type, must be set.`,
				},
				resource.Attribute{
					Name:        "max_connections_per_instance",
					Description: `(Optional) The max number of simultaneous connections that a single backend instance can handle. This is used to calculate the capacity of the group. Can be used in either CONNECTION or UTILIZATION balancing modes. For CONNECTION mode, either maxConnections or maxConnectionsPerInstance must be set.`,
				},
				resource.Attribute{
					Name:        "max_connections_per_endpoint",
					Description: `(Optional) The max number of simultaneous connections that a single backend network endpoint can handle. This is used to calculate the capacity of the group. Can be used in either CONNECTION or UTILIZATION balancing modes. For CONNECTION mode, either maxConnections or maxConnectionsPerEndpoint must be set.`,
				},
				resource.Attribute{
					Name:        "max_rate",
					Description: `(Optional) The max requests per second (RPS) of the group. Can be used with either RATE or UTILIZATION balancing modes, but required if RATE mode. For RATE mode, either maxRate or one of maxRatePerInstance or maxRatePerEndpoint, as appropriate for group type, must be set.`,
				},
				resource.Attribute{
					Name:        "max_rate_per_instance",
					Description: `(Optional) The max requests per second (RPS) that a single backend instance can handle. This is used to calculate the capacity of the group. Can be used in either balancing mode. For RATE mode, either maxRate or maxRatePerInstance must be set.`,
				},
				resource.Attribute{
					Name:        "max_rate_per_endpoint",
					Description: `(Optional) The max requests per second (RPS) that a single backend network endpoint can handle. This is used to calculate the capacity of the group. Can be used in either balancing mode. For RATE mode, either maxRate or maxRatePerEndpoint must be set.`,
				},
				resource.Attribute{
					Name:        "max_utilization",
					Description: `(Optional) Used when balancingMode is UTILIZATION. This ratio defines the CPU utilization target for the group. The default is 0.8. Valid range is [0.0, 1.0]. The ` + "`" + `cdn_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "cache_key_policy",
					Description: `(Optional) The CacheKeyPolicy for this CdnPolicy. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "signed_url_cache_max_age_sec",
					Description: `(Optional) Maximum number of seconds the response to a signed URL request will be considered fresh, defaults to 1hr (3600s). After this time period, the response will be revalidated before being served. When serving responses to signed URL requests, Cloud CDN will internally behave as though all responses from this backend had a "Cache-Control: public, max-age=[TTL]" header, regardless of any existing Cache-Control header. The actual headers served in responses will not be altered. The ` + "`" + `cache_key_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "include_host",
					Description: `(Optional) If true requests to different hosts will be cached separately.`,
				},
				resource.Attribute{
					Name:        "include_protocol",
					Description: `(Optional) If true, http and https requests will be cached separately.`,
				},
				resource.Attribute{
					Name:        "include_query_string",
					Description: `(Optional) If true, include query string parameters in the cache key according to query_string_whitelist and query_string_blacklist. If neither is set, the entire query string will be included. If false, the query string will be excluded from the cache key entirely.`,
				},
				resource.Attribute{
					Name:        "query_string_blacklist",
					Description: `(Optional) Names of query string parameters to exclude in cache keys. All other parameters will be included. Either specify query_string_whitelist or query_string_blacklist, not both. '&' and '=' will be percent encoded and not treated as delimiters.`,
				},
				resource.Attribute{
					Name:        "query_string_whitelist",
					Description: `(Optional) Names of query string parameters to include in cache keys. All other parameters will be excluded. Either specify query_string_whitelist or query_string_blacklist, not both. '&' and '=' will be percent encoded and not treated as delimiters. The ` + "`" + `iap` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "oauth2_client_id",
					Description: `(Required) OAuth2 Client ID for IAP`,
				},
				resource.Attribute{
					Name:        "oauth2_client_secret",
					Description: `(Required) OAuth2 Client Secret for IAP`,
				},
				resource.Attribute{
					Name:        "oauth2_client_secret_sha256",
					Description: `OAuth2 Client Secret SHA-256 for IAP ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import BackendService can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_backend_service.default projects/{{project}}/global/backendServices/{{name}} $ terraform import google_compute_backend_service.default {{project}}/{{name}} $ terraform import google_compute_backend_service.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import BackendService can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_backend_service.default projects/{{project}}/global/backendServices/{{name}} $ terraform import google_compute_backend_service.default {{project}}/{{name}} $ terraform import google_compute_backend_service.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_backend_service_signed_url_key",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `A key for signing Cloud CDN signed URLs for Backend Services.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"backend",
				"service",
				"signed",
				"url",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the signed URL key.`,
				},
				resource.Attribute{
					Name:        "key_value",
					Description: `(Required) 128-bit key value used for signing the URL. The key value must be a valid RFC 4648 Section 5 base64url encoded string.`,
				},
				resource.Attribute{
					Name:        "backend_service",
					Description: `(Required) The backend service this signed URL key belongs. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import BackendServiceSignedUrlKey can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_backend_service_signed_url_key.default projects/{{project}}/global/backendServices/{{backend_service}}/{{name}} $ terraform import google_compute_backend_service_signed_url_key.default {{project}}/{{backend_service}}/{{name}} $ terraform import google_compute_backend_service_signed_url_key.default {{backend_service}}/{{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_disk",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Persistent disks are durable storage devices that function similarly to the physical disks in a desktop or a server.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"disk",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource. Provide this property when you create the resource.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Labels to apply to this disk. A list of key->value pairs.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Size of the persistent disk, specified in GB. You can specify this field when creating a persistent disk using the sourceImage or sourceSnapshot parameter, or specify it alone to create an empty persistent disk. If you specify this field along with sourceImage or sourceSnapshot, the value of sizeGb must not be less than the size of the sourceImage or the size of the snapshot.`,
				},
				resource.Attribute{
					Name:        "physical_block_size_bytes",
					Description: `(Optional) Physical block size of the persistent disk, in bytes. If not present in a request, a default value is used. Currently supported sizes are 4096 and 16384, other sizes may be added in the future. If an unsupported value is requested, the error message will list the supported values for the caller's project.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) URL of the disk type resource describing which disk type to use to create the disk. Provide this when creating the disk.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional) The image from which to initialize this disk. This can be one of: the image's ` + "`" + `self_link` + "`" + `, ` + "`" + `projects/{project}/global/images/{image}` + "`" + `, ` + "`" + `projects/{project}/global/images/family/{family}` + "`" + `, ` + "`" + `global/images/{image}` + "`" + `, ` + "`" + `global/images/family/{family}` + "`" + `, ` + "`" + `family/{family}` + "`" + `, ` + "`" + `{project}/{family}` + "`" + `, ` + "`" + `{project}/{image}` + "`" + `, ` + "`" + `{family}` + "`" + `, or ` + "`" + `{image}` + "`" + `. If referred by family, the images names must include the family name. If they don't, use the [google_compute_image data source](/docs/providers/google/d/datasource_compute_image.html). For instance, the image ` + "`" + `centos-6-v20180104` + "`" + ` includes its family name ` + "`" + `centos-6` + "`" + `. These images can be referred by family name here.`,
				},
				resource.Attribute{
					Name:        "resource_policies",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) Resource policies applied to this disk for automatic snapshot creations.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) A reference to the zone where the disk resides.`,
				},
				resource.Attribute{
					Name:        "source_image_encryption_key",
					Description: `(Optional) The customer-supplied encryption key of the source image. Required if the source image is protected by a customer-supplied encryption key. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "disk_encryption_key",
					Description: `(Optional) Encrypts the disk using a customer-supplied encryption key. After you encrypt a disk with a customer-supplied key, you must provide the same key if you use the disk later (e.g. to create a disk snapshot or an image, or to attach the disk to a virtual machine). Customer-supplied encryption keys do not protect access to metadata of the disk. If you do not provide an encryption key when creating the disk, then the disk will be encrypted using an automatically generated key and you do not need to provide a key to use the disk later. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "snapshot",
					Description: `(Optional) The source snapshot used to create this disk. You can provide this as a partial or full URL to the resource. If the snapshot is in another project than this disk, you must supply a full URL. For example, the following are valid values:`,
				},
				resource.Attribute{
					Name:        "source_snapshot_encryption_key",
					Description: `(Optional) The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `source_image_encryption_key` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "raw_key",
					Description: `(Optional) Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource.`,
				},
				resource.Attribute{
					Name:        "sha256",
					Description: `The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource.`,
				},
				resource.Attribute{
					Name:        "kms_key_self_link",
					Description: `(Optional) The self link of the encryption key used to encrypt the disk. Also called KmsKeyName in the cloud console. In order to use this additional IAM permissions need to be set on the Compute Engine Service Agent. See https://cloud.google.com/compute/docs/disks/customer-managed-encryption#encrypt_a_new_persistent_disk_with_your_own_keys The ` + "`" + `disk_encryption_key` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "raw_key",
					Description: `(Optional) Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource.`,
				},
				resource.Attribute{
					Name:        "sha256",
					Description: `The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource.`,
				},
				resource.Attribute{
					Name:        "kms_key_self_link",
					Description: `(Optional) The self link of the encryption key used to encrypt the disk. Also called KmsKeyName in the cloud console. In order to use this additional IAM permissions need to be set on the Compute Engine Service Agent. See https://cloud.google.com/compute/docs/disks/customer-managed-encryption#encrypt_a_new_persistent_disk_with_your_own_keys The ` + "`" + `source_snapshot_encryption_key` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "raw_key",
					Description: `(Optional) Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource.`,
				},
				resource.Attribute{
					Name:        "kms_key_self_link",
					Description: `(Optional) The self link of the encryption key used to encrypt the disk. Also called KmsKeyName in the cloud console. In order to use this additional IAM permissions need to be set on the Compute Engine Service Agent. See https://cloud.google.com/compute/docs/disks/customer-managed-encryption#encrypt_a_new_persistent_disk_with_your_own_keys`,
				},
				resource.Attribute{
					Name:        "sha256",
					Description: `The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "label_fingerprint",
					Description: `The fingerprint used for optimistic locking of this resource. Used internally during updates.`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "last_attach_timestamp",
					Description: `Last attach timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "last_detach_timestamp",
					Description: `Last detach timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `Links to the users of the disk (attached instances) in form: project/zones/zone/instances/instance`,
				},
				resource.Attribute{
					Name:        "source_image_id",
					Description: `The ID value of the image used to create this disk. This value identifies the exact image that was used to create this persistent disk. For example, if you created the persistent disk from an image that was later deleted and recreated under the same name, the source image ID would identify the exact version of the image that was used.`,
				},
				resource.Attribute{
					Name:        "source_snapshot_id",
					Description: `The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to create this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 5 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Disk can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_disk.default projects/{{project}}/zones/{{zone}}/disks/{{name}} $ terraform import google_compute_disk.default {{project}}/{{zone}}/{{name}} $ terraform import google_compute_disk.default {{zone}}/{{name}} $ terraform import google_compute_disk.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "label_fingerprint",
					Description: `The fingerprint used for optimistic locking of this resource. Used internally during updates.`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "last_attach_timestamp",
					Description: `Last attach timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "last_detach_timestamp",
					Description: `Last detach timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `Links to the users of the disk (attached instances) in form: project/zones/zone/instances/instance`,
				},
				resource.Attribute{
					Name:        "source_image_id",
					Description: `The ID value of the image used to create this disk. This value identifies the exact image that was used to create this persistent disk. For example, if you created the persistent disk from an image that was later deleted and recreated under the same name, the source image ID would identify the exact version of the image that was used.`,
				},
				resource.Attribute{
					Name:        "source_snapshot_id",
					Description: `The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to create this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 5 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Disk can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_disk.default projects/{{project}}/zones/{{zone}}/disks/{{name}} $ terraform import google_compute_disk.default {{project}}/{{zone}}/{{name}} $ terraform import google_compute_disk.default {{zone}}/{{name}} $ terraform import google_compute_disk.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_external_vpn_gateway",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Represents a VPN gateway managed outside of GCP.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"external",
				"vpn",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource.`,
				},
				resource.Attribute{
					Name:        "redundancy_type",
					Description: `(Optional) Indicates the redundancy type of this external VPN gateway`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Optional) A list of interfaces on this external VPN gateway. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `interface` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The numberic ID for this interface. Allowed values are based on the redundancy type of this external VPN gateway`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) IP address of the interface in the external VPN gateway. Only IPv4 is supported. This IP address can be either from your on-premise gateway or another Cloud provider’s VPN gateway, it cannot be an IP address from Google Compute Engine. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import ExternalVpnGateway can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import -provider=google-beta google_compute_external_vpn_gateway.default projects/{{project}}/global/externalVpnGateways/{{name}} $ terraform import -provider=google-beta google_compute_external_vpn_gateway.default {{project}}/{{name}} $ terraform import -provider=google-beta google_compute_external_vpn_gateway.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_firewall",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Each network has its own firewall controlling access to and from the instances.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"firewall",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) The name or self_link of the network to attach this firewall to. - - -`,
				},
				resource.Attribute{
					Name:        "allow",
					Description: `(Optional) The list of ALLOW rules specified by this firewall. Each rule specifies a protocol and port-range tuple that describes a permitted connection. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "deny",
					Description: `(Optional) The list of DENY rules specified by this firewall. Each rule specifies a protocol and port-range tuple that describes a denied connection. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource. Provide this property when you create the resource.`,
				},
				resource.Attribute{
					Name:        "destination_ranges",
					Description: `(Optional) If destination ranges are specified, the firewall will apply only to traffic that has destination IP address in these ranges. These ranges must be expressed in CIDR format. Only IPv4 is supported.`,
				},
				resource.Attribute{
					Name:        "direction",
					Description: `(Optional) Direction of traffic to which this firewall applies; default is INGRESS. Note: For INGRESS traffic, it is NOT supported to specify destinationRanges; For EGRESS traffic, it is NOT supported to specify sourceRanges OR sourceTags.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Denotes whether the firewall rule is disabled, i.e not applied to the network it is associated with. When set to true, the firewall rule is not enforced and the network behaves as if it did not exist. If this is unspecified, the firewall rule will be enabled.`,
				},
				resource.Attribute{
					Name:        "enable_logging",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) This field denotes whether to enable logging for a particular firewall rule. If logging is enabled, logs will be exported to Stackdriver.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority for this rule. This is an integer between 0 and 65535, both inclusive. When not specified, the value assumed is 1000. Relative priorities determine precedence of conflicting rules. Lower value of priority implies higher precedence (eg, a rule with priority 0 has higher precedence than a rule with priority 1). DENY rules take precedence over ALLOW rules having equal priority.`,
				},
				resource.Attribute{
					Name:        "source_ranges",
					Description: `(Optional) If source ranges are specified, the firewall will apply only to traffic that has source IP address in these ranges. These ranges must be expressed in CIDR format. One or both of sourceRanges and sourceTags may be set. If both properties are set, the firewall will apply to traffic that has source IP address within sourceRanges OR the source IP that belongs to a tag listed in the sourceTags property. The connection does not need to match both properties for the firewall to apply. Only IPv4 is supported.`,
				},
				resource.Attribute{
					Name:        "source_service_accounts",
					Description: `(Optional) If source service accounts are specified, the firewall will apply only to traffic originating from an instance with a service account in this list. Source service accounts cannot be used to control traffic to an instance's external IP address because service accounts are associated with an instance, not an IP address. sourceRanges can be set at the same time as sourceServiceAccounts. If both are set, the firewall will apply to traffic that has source IP address within sourceRanges OR the source IP belongs to an instance with service account listed in sourceServiceAccount. The connection does not need to match both properties for the firewall to apply. sourceServiceAccounts cannot be used at the same time as sourceTags or targetTags.`,
				},
				resource.Attribute{
					Name:        "source_tags",
					Description: `(Optional) If source tags are specified, the firewall will apply only to traffic with source IP that belongs to a tag listed in source tags. Source tags cannot be used to control traffic to an instance's external IP address. Because tags are associated with an instance, not an IP address. One or both of sourceRanges and sourceTags may be set. If both properties are set, the firewall will apply to traffic that has source IP address within sourceRanges OR the source IP that belongs to a tag listed in the sourceTags property. The connection does not need to match both properties for the firewall to apply.`,
				},
				resource.Attribute{
					Name:        "target_service_accounts",
					Description: `(Optional) A list of service accounts indicating sets of instances located in the network that may make network connections as specified in allowed[]. targetServiceAccounts cannot be used at the same time as targetTags or sourceTags. If neither targetServiceAccounts nor targetTags are specified, the firewall rule applies to all instances on the specified network.`,
				},
				resource.Attribute{
					Name:        "target_tags",
					Description: `(Optional) A list of instance tags indicating sets of instances located in the network that may make network connections as specified in allowed[]. If no targetTags are specified, the firewall rule applies to all instances on the specified network.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `allow` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The IP protocol to which this rule applies. The protocol type is required when creating a firewall rule. This value can either be one of the following well known protocol strings (tcp, udp, icmp, esp, ah, sctp), or the IP protocol number.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(Optional) An optional list of ports to which this rule applies. This field is only applicable for UDP or TCP protocol. Each entry must be either an integer or a range. If not specified, this rule applies to connections through any port. Example inputs include: ["22"], ["80","443"], and ["12345-12349"]. The ` + "`" + `deny` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) The IP protocol to which this rule applies. The protocol type is required when creating a firewall rule. This value can either be one of the following well known protocol strings (tcp, udp, icmp, esp, ah, sctp), or the IP protocol number.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(Optional) An optional list of ports to which this rule applies. This field is only applicable for UDP or TCP protocol. Each entry must be either an integer or a range. If not specified, this rule applies to connections through any port. Example inputs include: ["22"], ["80","443"], and ["12345-12349"]. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Firewall can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_firewall.default projects/{{project}}/global/firewalls/{{name}} $ terraform import google_compute_firewall.default {{project}}/{{name}} $ terraform import google_compute_firewall.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Firewall can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_firewall.default projects/{{project}}/global/firewalls/{{name}} $ terraform import google_compute_firewall.default {{project}}/{{name}} $ terraform import google_compute_firewall.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_forwarding_rule",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `A ForwardingRule resource.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"forwarding",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource. Provide this property when you create the resource.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) The IP address that this forwarding rule is serving on behalf of. Addresses are restricted based on the forwarding rule's load balancing scheme (EXTERNAL or INTERNAL) and scope (global or regional). When the load balancing scheme is EXTERNAL, for global forwarding rules, the address must be a global IP, and for regional forwarding rules, the address must live in the same region as the forwarding rule. If this field is empty, an ephemeral IPv4 address from the same scope (global or regional) will be assigned. A regional forwarding rule supports IPv4 only. A global forwarding rule supports either IPv4 or IPv6. When the load balancing scheme is INTERNAL, this can only be an RFC 1918 IP address belonging to the network/subnet configured for the forwarding rule. By default, if this field is empty, an ephemeral internal IP address will be automatically allocated from the IP range of the subnet or network configured for this forwarding rule. ~>`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `(Optional) The IP protocol to which this rule applies. Valid options are TCP, UDP, ESP, AH, SCTP or ICMP. When the load balancing scheme is INTERNAL, only TCP and UDP are valid.`,
				},
				resource.Attribute{
					Name:        "backend_service",
					Description: `(Optional) A BackendService to receive the matched traffic. This is used only for INTERNAL load balancing.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional, Deprecated) ipVersion is not a valid field for regional forwarding rules.`,
				},
				resource.Attribute{
					Name:        "load_balancing_scheme",
					Description: `(Optional) This signifies what the ForwardingRule will be used for and can only take the following values: INTERNAL, EXTERNAL The value of INTERNAL means that this will be used for Internal Network Load Balancing (TCP, UDP). The value of EXTERNAL means that this will be used for External Load Balancing (HTTP(S) LB, External TCP/UDP LB, SSL Proxy)`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) For internal load balancing, this field identifies the network that the load balanced IP should belong to for this Forwarding Rule. If this field is not specified, the default network will be used. This field is only used for INTERNAL load balancing.`,
				},
				resource.Attribute{
					Name:        "port_range",
					Description: `(Optional) This field is used along with the target field for TargetHttpProxy, TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway, TargetPool, TargetInstance. Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets addressed to ports in the specified range will be forwarded to target. Forwarding rules with the same [IPAddress, IPProtocol] pair must have disjoint port ranges. Some types of forwarding target have constraints on the acceptable ports:`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(Optional) This field is used along with the backend_service field for internal load balancing. When the load balancing scheme is INTERNAL, a single port or a comma separated list of ports can be configured. Only packets addressed to these ports will be forwarded to the backends configured with this forwarding rule. You may specify a maximum of up to 5 ports.`,
				},
				resource.Attribute{
					Name:        "subnetwork",
					Description: `(Optional) The subnetwork that the load balanced IP should belong to for this Forwarding Rule. This field is only used for INTERNAL load balancing. If the network specified is in auto subnet mode, this field is optional. However, if the network is in custom subnet mode, a subnetwork must be specified.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Optional) This field is only used for EXTERNAL load balancing. A reference to a TargetPool resource to receive the matched traffic. This target must live in the same region as the forwarding rule. The forwarded traffic must be of a type appropriate to the target object.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) Labels to apply to this forwarding rule. A list of key->value pairs.`,
				},
				resource.Attribute{
					Name:        "all_ports",
					Description: `(Optional) For internal TCP/UDP load balancing (i.e. load balancing scheme is INTERNAL and protocol is TCP/UDP), set this to true to allow packets addressed to any ports to be forwarded to the backends configured with this forwarding rule. Used with backend service. Cannot be set if port or portRange are set.`,
				},
				resource.Attribute{
					Name:        "network_tier",
					Description: `(Optional) The networking tier used for configuring this address. This field can take the following values: PREMIUM or STANDARD. If this field is not specified, it is assumed to be PREMIUM.`,
				},
				resource.Attribute{
					Name:        "service_label",
					Description: `(Optional) An optional prefix to the service name for this Forwarding Rule. If specified, will be the first label of the fully qualified service name. The label must be 1-63 characters long, and comply with RFC1035. Specifically, the label must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) A reference to the region where the regional forwarding rule resides. This field is not applicable to global forwarding rules.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "label_fingerprint",
					Description: `The fingerprint used for optimistic locking of this resource. Used internally during updates.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The internal fully qualified service name for this Forwarding Rule. This field is only used for INTERNAL load balancing.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import ForwardingRule can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_forwarding_rule.default projects/{{project}}/regions/{{region}}/forwardingRules/{{name}} $ terraform import google_compute_forwarding_rule.default {{project}}/{{region}}/{{name}} $ terraform import google_compute_forwarding_rule.default {{region}}/{{name}} $ terraform import google_compute_forwarding_rule.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "label_fingerprint",
					Description: `The fingerprint used for optimistic locking of this resource. Used internally during updates.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The internal fully qualified service name for this Forwarding Rule. This field is only used for INTERNAL load balancing.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import ForwardingRule can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_forwarding_rule.default projects/{{project}}/regions/{{region}}/forwardingRules/{{name}} $ terraform import google_compute_forwarding_rule.default {{project}}/{{region}}/{{name}} $ terraform import google_compute_forwarding_rule.default {{region}}/{{name}} $ terraform import google_compute_forwarding_rule.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_global_address",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Represents a Global Address resource.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"global",
				"address",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Optional) The IP address or beginning of the address range represented by this resource. This can be supplied as an input to reserve a specific address or omitted to allow GCP to choose a valid one for you.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) Labels to apply to this address. A list of key->value pairs.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) The IP Version that will be used by this address. Valid options are ` + "`" + `IPV4` + "`" + ` or ` + "`" + `IPV6` + "`" + `. The default value is ` + "`" + `IPV4` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "prefix_length",
					Description: `(Optional) The prefix length of the IP range. If not present, it means the address field is a single IP address. This field is not applicable to addresses with addressType=EXTERNAL.`,
				},
				resource.Attribute{
					Name:        "address_type",
					Description: `(Optional) The type of the address to reserve, default is EXTERNAL.`,
				},
				resource.Attribute{
					Name:        "purpose",
					Description: `(Optional) The purpose of the resource. For global internal addresses it can be`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) The URL of the network in which to reserve the IP range. The IP range must be in RFC1918 space. The network cannot be deleted if there are any reserved IP ranges referring to it. This should only be set when using an Internal address.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "label_fingerprint",
					Description: `The fingerprint used for optimistic locking of this resource. Used internally during updates.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import GlobalAddress can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_global_address.default projects/{{project}}/global/addresses/{{name}} $ terraform import google_compute_global_address.default {{project}}/{{name}} $ terraform import google_compute_global_address.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "label_fingerprint",
					Description: `The fingerprint used for optimistic locking of this resource. Used internally during updates.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import GlobalAddress can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_global_address.default projects/{{project}}/global/addresses/{{name}} $ terraform import google_compute_global_address.default {{project}}/{{name}} $ terraform import google_compute_global_address.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_global_forwarding_rule",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Represents a GlobalForwardingRule resource.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"global",
				"forwarding",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) The URL of the target resource to receive the matched traffic. The forwarded traffic must be of a type appropriate to the target object. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource. Provide this property when you create the resource.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) The IP address that this forwarding rule is serving on behalf of. Addresses are restricted based on the forwarding rule's load balancing scheme (EXTERNAL or INTERNAL) and scope (global or regional). When the load balancing scheme is EXTERNAL, for global forwarding rules, the address must be a global IP, and for regional forwarding rules, the address must live in the same region as the forwarding rule. If this field is empty, an ephemeral IPv4 address from the same scope (global or regional) will be assigned. A regional forwarding rule supports IPv4 only. A global forwarding rule supports either IPv4 or IPv6. When the load balancing scheme is INTERNAL, this can only be an RFC 1918 IP address belonging to the network/subnet configured for the forwarding rule. By default, if this field is empty, an ephemeral internal IP address will be automatically allocated from the IP range of the subnet or network configured for this forwarding rule. ~>`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `(Optional) The IP protocol to which this rule applies. Valid options are TCP, UDP, ESP, AH, SCTP or ICMP. When the load balancing scheme is INTERNAL_SELF_MANAGED, only TCP is valid.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) The IP Version that will be used by this global forwarding rule. Valid options are IPV4 or IPV6.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) Labels to apply to this forwarding rule. A list of key->value pairs.`,
				},
				resource.Attribute{
					Name:        "load_balancing_scheme",
					Description: `(Optional) This signifies what the GlobalForwardingRule will be used for. The value of INTERNAL_SELF_MANAGED means that this will be used for Internal Global HTTP(S) LB. The value of EXTERNAL means that this will be used for External Global Load Balancing (HTTP(S) LB, External TCP/UDP LB, SSL Proxy) NOTE: Currently global forwarding rules cannot be used for INTERNAL load balancing.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) This field is not used for external load balancing. For INTERNAL_SELF_MANAGED load balancing, this field identifies the network that the load balanced IP should belong to for this global forwarding rule. If this field is not specified, the default network will be used.`,
				},
				resource.Attribute{
					Name:        "port_range",
					Description: `(Optional) This field is used along with the target field for TargetHttpProxy, TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway, TargetPool, TargetInstance. Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets addressed to ports in the specified range will be forwarded to target. Forwarding rules with the same [IPAddress, IPProtocol] pair must have disjoint port ranges. Some types of forwarding target have constraints on the acceptable ports:`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "label_fingerprint",
					Description: `The fingerprint used for optimistic locking of this resource. Used internally during updates.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import GlobalForwardingRule can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_global_forwarding_rule.default projects/{{project}}/global/forwardingRules/{{name}} $ terraform import google_compute_global_forwarding_rule.default {{project}}/{{name}} $ terraform import google_compute_global_forwarding_rule.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "label_fingerprint",
					Description: `The fingerprint used for optimistic locking of this resource. Used internally during updates.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import GlobalForwardingRule can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_global_forwarding_rule.default projects/{{project}}/global/forwardingRules/{{name}} $ terraform import google_compute_global_forwarding_rule.default {{project}}/{{name}} $ terraform import google_compute_global_forwarding_rule.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_ha_vpn_gateway",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Represents a VPN gateway running in GCP.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"ha",
				"vpn",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) The network this VPN gateway is accepting traffic for. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region this gateway should sit in.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vpn_interfaces",
					Description: `A list of interfaces on this VPN gateway. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. The ` + "`" + `vpn_interfaces` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The numeric ID of this VPN gateway interface.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) The external IP address for this VPN gateway interface. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import HaVpnGateway can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import -provider=google-beta google_compute_ha_vpn_gateway.default projects/{{project}}/regions/{{region}}/vpnGateways/{{name}} $ terraform import -provider=google-beta google_compute_ha_vpn_gateway.default {{project}}/{{region}}/{{name}} $ terraform import -provider=google-beta google_compute_ha_vpn_gateway.default {{region}}/{{name}} $ terraform import -provider=google-beta google_compute_ha_vpn_gateway.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vpn_interfaces",
					Description: `A list of interfaces on this VPN gateway. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. The ` + "`" + `vpn_interfaces` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The numeric ID of this VPN gateway interface.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) The external IP address for this VPN gateway interface. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import HaVpnGateway can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import -provider=google-beta google_compute_ha_vpn_gateway.default projects/{{project}}/regions/{{region}}/vpnGateways/{{name}} $ terraform import -provider=google-beta google_compute_ha_vpn_gateway.default {{project}}/{{region}}/{{name}} $ terraform import -provider=google-beta google_compute_ha_vpn_gateway.default {{region}}/{{name}} $ terraform import -provider=google-beta google_compute_ha_vpn_gateway.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_health_check",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Health Checks determine whether instances are responsive and able to do work.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"health",
				"check",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "check_interval_sec",
					Description: `(Optional) How often (in seconds) to send a health check. The default value is 5 seconds.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource. Provide this property when you create the resource.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `(Optional) A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.`,
				},
				resource.Attribute{
					Name:        "timeout_sec",
					Description: `(Optional) How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to have greater value than checkIntervalSec.`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `(Optional) A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.`,
				},
				resource.Attribute{
					Name:        "http_health_check",
					Description: `(Optional) A nested object resource Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "https_health_check",
					Description: `(Optional) A nested object resource Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "tcp_health_check",
					Description: `(Optional) A nested object resource Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "ssl_health_check",
					Description: `(Optional) A nested object resource Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `http_health_check` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) The value of the host header in the HTTP health check request. If left empty (default value), the public IP on behalf of which this health check is performed will be used.`,
				},
				resource.Attribute{
					Name:        "request_path",
					Description: `(Optional) The request path of the HTTP health check request. The default value is /.`,
				},
				resource.Attribute{
					Name:        "response",
					Description: `(Optional) The bytes to match against the beginning of the response data. If left empty (the default value), any response will indicate health. The response data can only be ASCII.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The TCP port number for the HTTP health check request. The default value is 80.`,
				},
				resource.Attribute{
					Name:        "port_name",
					Description: `(Optional) Port name as defined in InstanceGroup#NamedPort#name. If both port and port_name are defined, port takes precedence.`,
				},
				resource.Attribute{
					Name:        "proxy_header",
					Description: `(Optional) Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is NONE.`,
				},
				resource.Attribute{
					Name:        "port_specification",
					Description: `(Optional) Specifies how port is selected for health checking, can be one of the following values:`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) The value of the host header in the HTTPS health check request. If left empty (default value), the public IP on behalf of which this health check is performed will be used.`,
				},
				resource.Attribute{
					Name:        "request_path",
					Description: `(Optional) The request path of the HTTPS health check request. The default value is /.`,
				},
				resource.Attribute{
					Name:        "response",
					Description: `(Optional) The bytes to match against the beginning of the response data. If left empty (the default value), any response will indicate health. The response data can only be ASCII.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The TCP port number for the HTTPS health check request. The default value is 443.`,
				},
				resource.Attribute{
					Name:        "port_name",
					Description: `(Optional) Port name as defined in InstanceGroup#NamedPort#name. If both port and port_name are defined, port takes precedence.`,
				},
				resource.Attribute{
					Name:        "proxy_header",
					Description: `(Optional) Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is NONE.`,
				},
				resource.Attribute{
					Name:        "port_specification",
					Description: `(Optional) Specifies how port is selected for health checking, can be one of the following values:`,
				},
				resource.Attribute{
					Name:        "request",
					Description: `(Optional) The application data to send once the TCP connection has been established (default value is empty). If both request and response are empty, the connection establishment alone will indicate health. The request data can only be ASCII.`,
				},
				resource.Attribute{
					Name:        "response",
					Description: `(Optional) The bytes to match against the beginning of the response data. If left empty (the default value), any response will indicate health. The response data can only be ASCII.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The TCP port number for the TCP health check request. The default value is 443.`,
				},
				resource.Attribute{
					Name:        "port_name",
					Description: `(Optional) Port name as defined in InstanceGroup#NamedPort#name. If both port and port_name are defined, port takes precedence.`,
				},
				resource.Attribute{
					Name:        "proxy_header",
					Description: `(Optional) Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is NONE.`,
				},
				resource.Attribute{
					Name:        "port_specification",
					Description: `(Optional) Specifies how port is selected for health checking, can be one of the following values:`,
				},
				resource.Attribute{
					Name:        "request",
					Description: `(Optional) The application data to send once the SSL connection has been established (default value is empty). If both request and response are empty, the connection establishment alone will indicate health. The request data can only be ASCII.`,
				},
				resource.Attribute{
					Name:        "response",
					Description: `(Optional) The bytes to match against the beginning of the response data. If left empty (the default value), any response will indicate health. The response data can only be ASCII.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The TCP port number for the SSL health check request. The default value is 443.`,
				},
				resource.Attribute{
					Name:        "port_name",
					Description: `(Optional) Port name as defined in InstanceGroup#NamedPort#name. If both port and port_name are defined, port takes precedence.`,
				},
				resource.Attribute{
					Name:        "proxy_header",
					Description: `(Optional) Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is NONE.`,
				},
				resource.Attribute{
					Name:        "port_specification",
					Description: `(Optional) Specifies how port is selected for health checking, can be one of the following values:`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the health check. One of HTTP, HTTPS, TCP, or SSL.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import HealthCheck can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_health_check.default projects/{{project}}/global/healthChecks/{{name}} $ terraform import google_compute_health_check.default {{project}}/{{name}} $ terraform import google_compute_health_check.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the health check. One of HTTP, HTTPS, TCP, or SSL.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import HealthCheck can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_health_check.default projects/{{project}}/global/healthChecks/{{name}} $ terraform import google_compute_health_check.default {{project}}/{{name}} $ terraform import google_compute_health_check.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_http_health_check",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `An HttpHealthCheck resource.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"http",
				"health",
				"check",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "check_interval_sec",
					Description: `(Optional) How often (in seconds) to send a health check. The default value is 5 seconds.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource. Provide this property when you create the resource.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `(Optional) A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) The value of the host header in the HTTP health check request. If left empty (default value), the public IP on behalf of which this health check is performed will be used.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The TCP port number for the HTTP health check request. The default value is 80.`,
				},
				resource.Attribute{
					Name:        "request_path",
					Description: `(Optional) The request path of the HTTP health check request. The default value is /.`,
				},
				resource.Attribute{
					Name:        "timeout_sec",
					Description: `(Optional) How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to have greater value than checkIntervalSec.`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `(Optional) A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import HttpHealthCheck can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_http_health_check.default projects/{{project}}/global/httpHealthChecks/{{name}} $ terraform import google_compute_http_health_check.default {{project}}/{{name}} $ terraform import google_compute_http_health_check.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import HttpHealthCheck can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_http_health_check.default projects/{{project}}/global/httpHealthChecks/{{name}} $ terraform import google_compute_http_health_check.default {{project}}/{{name}} $ terraform import google_compute_http_health_check.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_https_health_check",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `An HttpsHealthCheck resource.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"https",
				"health",
				"check",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "check_interval_sec",
					Description: `(Optional) How often (in seconds) to send a health check. The default value is 5 seconds.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource. Provide this property when you create the resource.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `(Optional) A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) The value of the host header in the HTTPS health check request. If left empty (default value), the public IP on behalf of which this health check is performed will be used.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The TCP port number for the HTTPS health check request. The default value is 80.`,
				},
				resource.Attribute{
					Name:        "request_path",
					Description: `(Optional) The request path of the HTTPS health check request. The default value is /.`,
				},
				resource.Attribute{
					Name:        "timeout_sec",
					Description: `(Optional) How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to have greater value than checkIntervalSec.`,
				},
				resource.Attribute{
					Name:        "unhealthy_threshold",
					Description: `(Optional) A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import HttpsHealthCheck can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_https_health_check.default projects/{{project}}/global/httpsHealthChecks/{{name}} $ terraform import google_compute_https_health_check.default {{project}}/{{name}} $ terraform import google_compute_https_health_check.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import HttpsHealthCheck can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_https_health_check.default projects/{{project}}/global/httpsHealthChecks/{{name}} $ terraform import google_compute_https_health_check.default {{project}}/{{name}} $ terraform import google_compute_https_health_check.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_image",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Represents an Image resource.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"image",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource. Provide this property when you create the resource.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `(Optional) Size of the image when restored onto a persistent disk (in GB).`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `(Optional) The name of the image family to which this image belongs. You can create disks by specifying an image family instead of a specific image name. The image family always returns its latest image that is not deprecated. The name of the image family must comply with RFC1035.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Labels to apply to this Image.`,
				},
				resource.Attribute{
					Name:        "licenses",
					Description: `(Optional) Any applicable license URI.`,
				},
				resource.Attribute{
					Name:        "raw_disk",
					Description: `(Optional) The parameters of the raw disk image. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "source_disk",
					Description: `(Optional) The source disk to create this image based on. You must provide either this property or the rawDisk.source property but not both to create an image.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `raw_disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "container_type",
					Description: `(Optional) The format used to encode and transmit the block device, which should be TAR. This is just a container and transmission format and not a runtime format. Provided by the client when the disk image is created.`,
				},
				resource.Attribute{
					Name:        "sha1",
					Description: `(Optional) An optional SHA1 checksum of the disk image before unpackaging. This is provided by the client when the disk image is created.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) The full Google Cloud Storage URL where disk storage is stored You must provide either this property or the sourceDisk property but not both. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "archive_size_bytes",
					Description: `Size of the image tar.gz archive stored in Google Cloud Storage (in bytes).`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "label_fingerprint",
					Description: `The fingerprint used for optimistic locking of this resource. Used internally during updates.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Image can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_image.default projects/{{project}}/global/images/{{name}} $ terraform import google_compute_image.default {{project}}/{{name}} $ terraform import google_compute_image.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "archive_size_bytes",
					Description: `Size of the image tar.gz archive stored in Google Cloud Storage (in bytes).`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "label_fingerprint",
					Description: `The fingerprint used for optimistic locking of this resource. Used internally during updates.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Image can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_image.default projects/{{project}}/global/images/{{name}} $ terraform import google_compute_image.default {{project}}/{{name}} $ terraform import google_compute_image.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_instance",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Manages a VM instance resource within GCE.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "boot_disk",
					Description: `(Required) The boot disk for the instance. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `(Required) The machine type to create.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the resource, required by GCE. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The zone that the machine should be created in.`,
				},
				resource.Attribute{
					Name:        "network_interface",
					Description: `(Required) Networks to attach to the instance. This can be specified multiple times. Structure is documented below. - - -`,
				},
				resource.Attribute{
					Name:        "allow_stopping_for_update",
					Description: `(Optional) If true, allows Terraform to stop the instance to update its properties. If you try to update a property that requires stopping the instance without setting this field, the update will fail.`,
				},
				resource.Attribute{
					Name:        "attached_disk",
					Description: `(Optional) Additional disks to attach to the instance. Can be repeated multiple times for multiple disks. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "can_ip_forward",
					Description: `(Optional) Whether to allow sending and receiving of packets with non-matching source or destination IPs. This defaults to false.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A brief description of this resource.`,
				},
				resource.Attribute{
					Name:        "deletion_protection",
					Description: `(Optional) Enable deletion protection on this instance. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Optional) A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid. Valid format is a series of labels 1-63 characters long matching the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "guest_accelerator",
					Description: `(Optional) List of the type and count of accelerator cards attached to the instance. Structure documented below.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A map of key/value label pairs to assign to the instance.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Metadata key/value pairs to make available from within the instance. Ssh keys attached in the Cloud Console will be removed. Add them to your config in order to keep them attached to your instance.`,
				},
				resource.Attribute{
					Name:        "metadata_startup_script",
					Description: `(Optional) An alternative to using the startup-script metadata key, except this one forces the instance to be recreated (thus re-running the script) if it is changed. This replaces the startup-script metadata key on the created instance and thus the two mechanisms are not allowed to be used simultaneously.`,
				},
				resource.Attribute{
					Name:        "min_cpu_platform",
					Description: `(Optional) Specifies a minimum CPU platform for the VM instance. Applicable values are the friendly names of CPU platforms, such as ` + "`" + `Intel Haswell` + "`" + ` or ` + "`" + `Intel Skylake` + "`" + `. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "scheduling",
					Description: `(Optional) The scheduling strategy to use. More details about this configuration option are detailed below.`,
				},
				resource.Attribute{
					Name:        "scratch_disk",
					Description: `(Optional) Scratch disks to attach to the instance. This can be specified multiple times for multiple scratch disks. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "service_account",
					Description: `(Optional) Service account to attach to the instance. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of tags to attach to the instance.`,
				},
				resource.Attribute{
					Name:        "shielded_instance_config",
					Description: `(Optional) Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "auto_delete",
					Description: `(Optional) Whether the disk will be auto-deleted when the instance is deleted. Defaults to true.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(Optional) Name with which attached disk will be accessible. On the instance, this device will be ` + "`" + `/dev/disk/by-id/google-{{device_name}}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_encryption_key_raw",
					Description: `(Optional) A 256-bit [customer-supplied encryption key] (https://cloud.google.com/compute/docs/disks/customer-supplied-encryption), encoded in [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4) to encrypt this disk. Only one of ` + "`" + `kms_key_self_link` + "`" + ` and ` + "`" + `disk_encryption_key_raw` + "`" + ` may be set.`,
				},
				resource.Attribute{
					Name:        "kms_key_self_link",
					Description: `(Optional) The self_link of the encryption key that is stored in Google Cloud KMS to encrypt this disk. Only one of ` + "`" + `kms_key_self_link` + "`" + ` and ` + "`" + `disk_encryption_key_raw` + "`" + ` may be set.`,
				},
				resource.Attribute{
					Name:        "initialize_params",
					Description: `(Optional) Parameters for a new disk that will be created alongside the new instance. Either ` + "`" + `initialize_params` + "`" + ` or ` + "`" + `source` + "`" + ` must be set. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) The name or self_link of the existing disk (such as those managed by ` + "`" + `google_compute_disk` + "`" + `) or disk image. To create an instance from a snapshot, first create a ` + "`" + `google_compute_disk` + "`" + ` from a snapshot and reference it here. The ` + "`" + `initialize_params` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) The size of the image in gigabytes. If not specified, it will inherit the size of its base image.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The GCE disk type. May be set to pd-standard or pd-ssd.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional) The image from which to initialize this disk. This can be one of: the image's ` + "`" + `self_link` + "`" + `, ` + "`" + `projects/{project}/global/images/{image}` + "`" + `, ` + "`" + `projects/{project}/global/images/family/{family}` + "`" + `, ` + "`" + `global/images/{image}` + "`" + `, ` + "`" + `global/images/family/{family}` + "`" + `, ` + "`" + `family/{family}` + "`" + `, ` + "`" + `{project}/{family}` + "`" + `, ` + "`" + `{project}/{image}` + "`" + `, ` + "`" + `{family}` + "`" + `, or ` + "`" + `{image}` + "`" + `. If referred by family, the images names must include the family name. If they don't, use the [google_compute_image data source](/docs/providers/google/d/datasource_compute_image.html). For instance, the image ` + "`" + `centos-6-v20180104` + "`" + ` includes its family name ` + "`" + `centos-6` + "`" + `. These images can be referred by family name here. The ` + "`" + `scratch_disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Optional) The disk interface to use for attaching this disk; either SCSI or NVME. Defaults to SCSI. The ` + "`" + `attached_disk` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) The name or self_link of the disk to attach to this instance.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(Optional) Name with which the attached disk will be accessible under ` + "`" + `/dev/disk/by-id/` + "`" + ``,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Either "READ_ONLY" or "READ_WRITE", defaults to "READ_WRITE" If you have a persistent disk with data that you want to share between multiple instances, detach it from any read-write instances and attach it to one or more instances in read-only mode.`,
				},
				resource.Attribute{
					Name:        "disk_encryption_key_raw",
					Description: `(Optional) A 256-bit [customer-supplied encryption key] (https://cloud.google.com/compute/docs/disks/customer-supplied-encryption), encoded in [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4) to encrypt this disk. Only one of ` + "`" + `kms_key_self_link` + "`" + ` and ` + "`" + `disk_encryption_key_raw` + "`" + ` may be set.`,
				},
				resource.Attribute{
					Name:        "kms_key_self_link",
					Description: `(Optional) The self_link of the encryption key that is stored in Google Cloud KMS to encrypt this disk. Only one of ` + "`" + `kms_key_self_link` + "`" + ` and ` + "`" + `disk_encryption_key_raw` + "`" + ` may be set. The ` + "`" + `network_interface` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) The name or self_link of the network to attach this interface to. Either ` + "`" + `network` + "`" + ` or ` + "`" + `subnetwork` + "`" + ` must be provided.`,
				},
				resource.Attribute{
					Name:        "network_ip",
					Description: `(Optional) The private IP address to assign to the instance. If empty, the address will be automatically assigned.`,
				},
				resource.Attribute{
					Name:        "access_config",
					Description: `(Optional) Access configurations, i.e. IPs via which this instance can be accessed via the Internet. Omit to ensure that the instance is not accessible from the Internet. If omitted, ssh provisioners will not work unless Terraform can send traffic to the instance's network (e.g. via tunnel or because it is running on another cloud instance on that network). This block can be repeated multiple times. Structure documented below.`,
				},
				resource.Attribute{
					Name:        "alias_ip_range",
					Description: `(Optional) An array of alias IP ranges for this network interface. Can only be specified for network interfaces on subnet-mode networks. Structure documented below. The ` + "`" + `access_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "nat_ip",
					Description: `(Optional) The IP address that will be 1:1 mapped to the instance's network ip. If not given, one will be generated.`,
				},
				resource.Attribute{
					Name:        "public_ptr_domain_name",
					Description: `(Optional) The DNS domain name for the public PTR record. To set this field on an instance, you must be verified as the owner of the domain. See [the docs](https://cloud.google.com/compute/docs/instances/create-ptr-record) for how to become verified as a domain owner.`,
				},
				resource.Attribute{
					Name:        "network_tier",
					Description: `(Optional) The [networking tier][network-tier] used for configuring this instance. This field can take the following values: PREMIUM or STANDARD. If this field is not specified, it is assumed to be PREMIUM. The ` + "`" + `alias_ip_range` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "ip_cidr_range",
					Description: `The IP CIDR range represented by this alias IP range. This IP CIDR range must belong to the specified subnetwork and cannot contain IP addresses reserved by system or used by other network interfaces. This range may be a single IP address (e.g. 10.2.3.4), a netmask (e.g. /24) or a CIDR format string (e.g. 10.1.2.0/24).`,
				},
				resource.Attribute{
					Name:        "subnetwork_range_name",
					Description: `(Optional) The subnetwork secondary range name specifying the secondary range from which to allocate the IP CIDR range for this alias IP range. If left unspecified, the primary range of the subnetwork will be used. The ` + "`" + `service_account` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) The service account e-mail address. If not given, the default Google Compute Engine service account is used.`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `(Required) A list of service scopes. Both OAuth2 URLs and gcloud short names are supported. To allow full access to all Cloud APIs, use the ` + "`" + `cloud-platform` + "`" + ` scope. See a complete list of scopes [here](https://cloud.google.com/sdk/gcloud/reference/alpha/compute/instances/set-scopes#--scopes).`,
				},
				resource.Attribute{
					Name:        "preemptible",
					Description: `(Optional) Specifies if the instance is preemptible. If this field is set to true, then ` + "`" + `automatic_restart` + "`" + ` must be set to false. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "on_host_maintenance",
					Description: `(Optional) Describes maintenance behavior for the instance. Can be MIGRATE or TERMINATE, for more info, read [here](https://cloud.google.com/compute/docs/instances/setting-instance-scheduling-options).`,
				},
				resource.Attribute{
					Name:        "automatic_restart",
					Description: `(Optional) Specifies if the instance should be restarted if it was terminated by Compute Engine (not a user). Defaults to true.`,
				},
				resource.Attribute{
					Name:        "node_affinities",
					Description: `(Optional) Specifies node affinities or anti-affinities to determine which sole-tenant nodes your instances and managed instance groups will use as host systems. Read more on sole-tenant node creation [here](https://cloud.google.com/compute/docs/nodes/create-nodes). Structure documented below. The ` + "`" + `guest_accelerator` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `The server-assigned unique identifier of this instance.`,
				},
				resource.Attribute{
					Name:        "metadata_fingerprint",
					Description: `The unique fingerprint of the metadata.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource.`,
				},
				resource.Attribute{
					Name:        "tags_fingerprint",
					Description: `The unique fingerprint of the tags.`,
				},
				resource.Attribute{
					Name:        "label_fingerprint",
					Description: `The unique fingerprint of the labels.`,
				},
				resource.Attribute{
					Name:        "cpu_platform",
					Description: `The CPU platform used by this instance.`,
				},
				resource.Attribute{
					Name:        "network_interface.0.network_ip",
					Description: `The internal ip address of the instance, either manually or dynamically assigned.`,
				},
				resource.Attribute{
					Name:        "network_interface.0.access_config.0.nat_ip",
					Description: `If the instance has an access config, either the given external ip (in the ` + "`" + `nat_ip` + "`" + ` field) or the ephemeral (generated) ip (if you didn't provide one).`,
				},
				resource.Attribute{
					Name:        "attached_disk.0.disk_encryption_key_sha256",
					Description: `The [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4) encoded SHA-256 hash of the [customer-supplied encryption key] (https://cloud.google.com/compute/docs/disks/customer-supplied-encryption) that protects this resource.`,
				},
				resource.Attribute{
					Name:        "boot_disk.disk_encryption_key_sha256",
					Description: `The [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4) encoded SHA-256 hash of the [customer-supplied encryption key] (https://cloud.google.com/compute/docs/disks/customer-supplied-encryption) that protects this resource.`,
				},
				resource.Attribute{
					Name:        "disk.0.disk_encryption_key_sha256",
					Description: `The [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4) encoded SHA-256 hash of the [customer-supplied encryption key] (https://cloud.google.com/compute/docs/disks/customer-supplied-encryption) that protects this resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 20 minutes. - ` + "`" + `update` + "`" + ` - Default is 20 minutes. - ` + "`" + `delete` + "`" + ` - Default is 20 minutes. ## Import ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_id",
					Description: `The server-assigned unique identifier of this instance.`,
				},
				resource.Attribute{
					Name:        "metadata_fingerprint",
					Description: `The unique fingerprint of the metadata.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource.`,
				},
				resource.Attribute{
					Name:        "tags_fingerprint",
					Description: `The unique fingerprint of the tags.`,
				},
				resource.Attribute{
					Name:        "label_fingerprint",
					Description: `The unique fingerprint of the labels.`,
				},
				resource.Attribute{
					Name:        "cpu_platform",
					Description: `The CPU platform used by this instance.`,
				},
				resource.Attribute{
					Name:        "network_interface.0.network_ip",
					Description: `The internal ip address of the instance, either manually or dynamically assigned.`,
				},
				resource.Attribute{
					Name:        "network_interface.0.access_config.0.nat_ip",
					Description: `If the instance has an access config, either the given external ip (in the ` + "`" + `nat_ip` + "`" + ` field) or the ephemeral (generated) ip (if you didn't provide one).`,
				},
				resource.Attribute{
					Name:        "attached_disk.0.disk_encryption_key_sha256",
					Description: `The [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4) encoded SHA-256 hash of the [customer-supplied encryption key] (https://cloud.google.com/compute/docs/disks/customer-supplied-encryption) that protects this resource.`,
				},
				resource.Attribute{
					Name:        "boot_disk.disk_encryption_key_sha256",
					Description: `The [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4) encoded SHA-256 hash of the [customer-supplied encryption key] (https://cloud.google.com/compute/docs/disks/customer-supplied-encryption) that protects this resource.`,
				},
				resource.Attribute{
					Name:        "disk.0.disk_encryption_key_sha256",
					Description: `The [RFC 4648 base64](https://tools.ietf.org/html/rfc4648#section-4) encoded SHA-256 hash of the [customer-supplied encryption key] (https://cloud.google.com/compute/docs/disks/customer-supplied-encryption) that protects this resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 20 minutes. - ` + "`" + `update` + "`" + ` - Default is 20 minutes. - ` + "`" + `delete` + "`" + ` - Default is 20 minutes. ## Import ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_instance_from_template",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Manages a VM instance resource within GCE.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"instance",
				"from",
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the resource, required by GCE. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "source_instance_template",
					Description: `(Required) Name or self link of an instance template to create the instance based on. - - -`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) The zone that the machine should be created in. If not set, the provider zone is used. In addition to these, all arguments from ` + "`" + `google_compute_instance` + "`" + ` are supported as a way to override the properties in the template. All exported attributes from ` + "`" + `google_compute_instance` + "`" + ` are likewise exported here. To support removal of Optional/Computed fields in Terraform 0.12 the following fields are marked [Attributes as Blocks](/docs/configuration/attr-as-blocks.html):`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_instance_group",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Manages an Instance Group within GCE.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"instance",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the instance group. Must be 1-63 characters long and comply with [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters include lowercase letters, numbers, and hyphens.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The zone that this instance group should be created in. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional textual description of the instance group.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `(Optional) List of instances in the group. They should be given as self_link URLs. When adding instances they must all be in the same network and zone as the instance group.`,
				},
				resource.Attribute{
					Name:        "named_port",
					Description: `(Optional) The named port configuration. See the section below for details on configuration.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) The URL of the network the instance group is in. If this is different from the network where the instances are in, the creation fails. Defaults to the network where the instances are in (if neither ` + "`" + `network` + "`" + ` nor ` + "`" + `instances` + "`" + ` is specified, this field will be blank). The ` + "`" + `named_port` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name which the port will be mapped to.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The port number to map the name to. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The number of instances in the group. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is ` + "`" + `6 minutes` + "`" + ` - ` + "`" + `update` + "`" + ` - Default is ` + "`" + `6 minutes` + "`" + ` - ` + "`" + `delete` + "`" + ` - Default is ` + "`" + `6 minutes` + "`" + ` ## Import Instance group can be imported using the ` + "`" + `zone` + "`" + ` and ` + "`" + `name` + "`" + ` with an optional ` + "`" + `project` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_instance_group.webservers us-central1-a/terraform-webservers $ terraform import google_compute_instance_group.webservers big-project/us-central1-a/terraform-webservers ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The number of instances in the group. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is ` + "`" + `6 minutes` + "`" + ` - ` + "`" + `update` + "`" + ` - Default is ` + "`" + `6 minutes` + "`" + ` - ` + "`" + `delete` + "`" + ` - Default is ` + "`" + `6 minutes` + "`" + ` ## Import Instance group can be imported using the ` + "`" + `zone` + "`" + ` and ` + "`" + `name` + "`" + ` with an optional ` + "`" + `project` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_instance_group.webservers us-central1-a/terraform-webservers $ terraform import google_compute_instance_group.webservers big-project/us-central1-a/terraform-webservers ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_instance_group_manager",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Manages an Instance Group within GCE.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"instance",
				"group",
				"manager",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "base_instance_name",
					Description: `(Required) The base instance name to use for instances in this group. The value must be a valid [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters are lowercase letters, numbers, and hyphens (-). Instances are named by appending a hyphen and a random four-character string to the base instance name.`,
				},
				resource.Attribute{
					Name:        "instance_template",
					Description: `(Required, [GA](https://terraform.io/docs/providers/google/provider_versions.html)) The full URL to an instance template from which all new instances will be created. This field is only present in the ` + "`" + `google` + "`" + ` provider.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) Application versions managed by this instance group. Each version deals with a specific instance template, allowing canary release scenarios. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the instance group manager. Must be 1-63 characters long and comply with [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters include lowercase letters, numbers, and hyphens.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The zone that instances in this group should be created in. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional textual description of the instance group manager.`,
				},
				resource.Attribute{
					Name:        "named_port",
					Description: `(Optional) The named port configuration. See the section below for details on configuration.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "update_strategy",
					Description: `(Optional, Default ` + "`" + `"REPLACE"` + "`" + `) If the ` + "`" + `instance_template` + "`" + ` resource is modified, a value of ` + "`" + `"NONE"` + "`" + ` will prevent any of the managed instances from being restarted by Terraform. A value of ` + "`" + `"REPLACE"` + "`" + ` will restart all of the instances at once. This field is only present in the ` + "`" + `google` + "`" + ` provider.`,
				},
				resource.Attribute{
					Name:        "target_size",
					Description: `(Optional) The target number of running instances for this managed instance group. This value should always be explicitly set unless this resource is attached to an autoscaler, in which case it should never be set. Defaults to ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "target_pools",
					Description: `(Optional) The full URL of all target pools to which new instances in the group are added. Updating the target pools attribute does not affect existing instances.`,
				},
				resource.Attribute{
					Name:        "wait_for_instances",
					Description: `(Optional) Whether to wait for all instances to be created/updated before returning. Note that if this is set to true and the operation does not succeed, Terraform will continue trying until it times out. ---`,
				},
				resource.Attribute{
					Name:        "auto_healing_policies",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) The autohealing policies for this managed instance group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).`,
				},
				resource.Attribute{
					Name:        "update_policy",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/beta/instanceGroupManagers/patch) - - - The ` + "`" + `update_policy` + "`" + ` block supports: ` + "`" + `` + "`" + `` + "`" + `hcl update_policy{ type = "PROACTIVE" minimal_action = "REPLACE" max_surge_percent = 20 max_unavailable_fixed = 2 min_ready_sec = 50 } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "minimal_action",
					Description: `(Required) - Minimal action to be taken on an instance. Valid values are ` + "`" + `"RESTART"` + "`" + `, ` + "`" + `"REPLACE"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) - The type of update. Valid values are ` + "`" + `"OPPORTUNISTIC"` + "`" + `, ` + "`" + `"PROACTIVE"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "max_surge_fixed",
					Description: `(Optional), The maximum number of instances that can be created above the specified targetSize during the update process. Conflicts with ` + "`" + `max_surge_percent` + "`" + `. If neither is set, defaults to 1`,
				},
				resource.Attribute{
					Name:        "max_surge_percent",
					Description: `(Optional), The maximum number of instances(calculated as percentage) that can be created above the specified targetSize during the update process. Conflicts with ` + "`" + `max_surge_fixed` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_unavailable_fixed",
					Description: `(Optional), The maximum number of instances that can be unavailable during the update process. Conflicts with ` + "`" + `max_unavailable_percent` + "`" + `. If neither is set, defaults to 1`,
				},
				resource.Attribute{
					Name:        "max_unavailable_percent",
					Description: `(Optional), The maximum number of instances(calculated as percentage) that can be unavailable during the update process. Conflicts with ` + "`" + `max_unavailable_fixed` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "min_ready_sec",
					Description: `(Optional), Minimum number of seconds to wait for after a newly created instance becomes available. This value must be from range [0, 3600] - - - The ` + "`" + `named_port` + "`" + ` block supports: (Include a ` + "`" + `named_port` + "`" + ` block for each named-port required).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the port.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The port number. - - - The ` + "`" + `auto_healing_policies` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `(Required) The health check resource that signals autohealing.`,
				},
				resource.Attribute{
					Name:        "initial_delay_sec",
					Description: `(Required) The number of seconds that the managed instance group waits before it applies autohealing policies to new instances or recently recreated instances. Between 0 and 3600. The ` + "`" + `version` + "`" + ` block supports: ` + "`" + `` + "`" + `` + "`" + `hcl version { name = "appserver-canary" instance_template = "${google_compute_instance_template.appserver-canary.self_link}" target_size { fixed = 1 } } ` + "`" + `` + "`" + `` + "`" + ` ` + "`" + `` + "`" + `` + "`" + `hcl version { name = "appserver-canary" instance_template = "${google_compute_instance_template.appserver-canary.self_link}" target_size { percent = 20 } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) - Version name.`,
				},
				resource.Attribute{
					Name:        "instance_template",
					Description: `(Required) - The full URL to an instance template from which all new instances of this version will be created.`,
				},
				resource.Attribute{
					Name:        "target_size",
					Description: `(Optional) - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below. -> Exactly one ` + "`" + `version` + "`" + ` you specify must not have a ` + "`" + `target_size` + "`" + ` specified. During a rolling update, the instance group manager will fulfill the ` + "`" + `target_size` + "`" + ` constraints of every other ` + "`" + `version` + "`" + `, and any remaining instances will be provisioned with the version where ` + "`" + `target_size` + "`" + ` is unset. The ` + "`" + `target_size` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "fixed",
					Description: `(Optional), The number of instances which are managed for this version. Conflicts with ` + "`" + `percent` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "percent",
					Description: `(Optional), The number of instances (calculated as percentage) which are managed for this version. Conflicts with ` + "`" + `fixed` + "`" + `. Note that when using ` + "`" + `percent` + "`" + `, rounding will be in favor of explicitly set ` + "`" + `target_size` + "`" + ` values; a managed instance group with 2 instances and 2 ` + "`" + `version` + "`" + `s, one of which has a ` + "`" + `target_size.percent` + "`" + ` of ` + "`" + `60` + "`" + ` will create 2 instances of that ` + "`" + `version` + "`" + `. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the instance group manager.`,
				},
				resource.Attribute{
					Name:        "instance_group",
					Description: `The full URL of the instance group created by the manager.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URL of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 5 minutes. - ` + "`" + `update` + "`" + ` - Default is 5 minutes. - ` + "`" + `delete` + "`" + ` - Default is 15 minutes. ## Import Instance group managers can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_instance_group_manager.appserver {{project}}/{{zone}}/{{name}} $ terraform import google_compute_instance_group_manager.appserver {{project}}/{{name}} $ terraform import google_compute_instance_group_manager.appserver {{name}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the instance group manager.`,
				},
				resource.Attribute{
					Name:        "instance_group",
					Description: `The full URL of the instance group created by the manager.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URL of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 5 minutes. - ` + "`" + `update` + "`" + ` - Default is 5 minutes. - ` + "`" + `delete` + "`" + ` - Default is 15 minutes. ## Import Instance group managers can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_instance_group_manager.appserver {{project}}/{{zone}}/{{name}} $ terraform import google_compute_instance_group_manager.appserver {{project}}/{{name}} $ terraform import google_compute_instance_group_manager.appserver {{name}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_instance_iam_policy",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Collection of resources to manage IAM policy for a GCE instance.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"instance",
				"iam",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Required) The name of the instance.`,
				},
				resource.Attribute{
					Name:        "member/members",
					Description: `(Required) Identities that will be granted the privilege in ` + "`" + `role` + "`" + `. Each entry can have one of the following values:`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. Only one ` + "`" + `google_compute_instance_iam_binding` + "`" + ` can be used per role. Note that custom roles must be of the format ` + "`" + `[projects|organizations]/{parent-name}/roles/{role-name}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy_data",
					Description: `(Required only by ` + "`" + `google_compute_instance_iam_policy` + "`" + `) The policy data generated by a ` + "`" + `google_iam_policy` + "`" + ` data source.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) The zone of the instance. If unspecified, this defaults to the zone configured in the provider. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the instance's IAM policy. ## Import For all import syntaxes, the "resource in question" can take any of the following forms:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the instance's IAM policy. ## Import For all import syntaxes, the "resource in question" can take any of the following forms:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_instance_template",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Manages a VM instance template resource within GCE.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"instance",
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "disk",
					Description: `(Required) Disks to attach to instances created from this template. This can be specified multiple times for multiple disks. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `(Required) The machine type to create. To create a machine with a [custom type][custom-vm-types] (such as extended memory), format the value like ` + "`" + `custom-VCPUS-MEM_IN_MB` + "`" + ` like ` + "`" + `custom-6-20480` + "`" + ` for 6 vCPU and 20GB of RAM. - - -`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the instance template. If you leave this blank, Terraform will auto-generate a unique name.`,
				},
				resource.Attribute{
					Name:        "name_prefix",
					Description: `(Optional) Creates a unique name beginning with the specified prefix. Conflicts with ` + "`" + `name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "can_ip_forward",
					Description: `(Optional) Whether to allow sending and receiving of packets with non-matching source or destination IPs. This defaults to false.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A brief description of this resource.`,
				},
				resource.Attribute{
					Name:        "instance_description",
					Description: `(Optional) A brief description to use for instances created from this template.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs to assign to instances created from this template,`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Metadata key/value pairs to make available from within instances created from this template.`,
				},
				resource.Attribute{
					Name:        "metadata_startup_script",
					Description: `(Optional) An alternative to using the startup-script metadata key, mostly to match the compute_instance resource. This replaces the startup-script metadata key on the created instance and thus the two mechanisms are not allowed to be used simultaneously.`,
				},
				resource.Attribute{
					Name:        "network_interface",
					Description: `(Required) Networks to attach to instances created from this template. This can be specified multiple times for multiple networks. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) An instance template is a global resource that is not bound to a zone or a region. However, you can still specify some regional resources in an instance template, which restricts the template to the region where that resource resides. For example, a custom ` + "`" + `subnetwork` + "`" + ` resource is tied to a specific region. Defaults to the region of the Provider if no value is given.`,
				},
				resource.Attribute{
					Name:        "scheduling",
					Description: `(Optional) The scheduling strategy to use. More details about this configuration option are detailed below.`,
				},
				resource.Attribute{
					Name:        "service_account",
					Description: `(Optional) Service account to attach to the instance. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Tags to attach to the instance.`,
				},
				resource.Attribute{
					Name:        "guest_accelerator",
					Description: `(Optional) List of the type and count of accelerator cards attached to the instance. Structure documented below.`,
				},
				resource.Attribute{
					Name:        "min_cpu_platform",
					Description: `(Optional) Specifies a minimum CPU platform. Applicable values are the friendly names of CPU platforms, such as ` + "`" + `Intel Haswell` + "`" + ` or ` + "`" + `Intel Skylake` + "`" + `. See the complete list [here](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).`,
				},
				resource.Attribute{
					Name:        "shielded_instance_config",
					Description: `(Optional) Enable [Shielded VM](https://cloud.google.com/security/shielded-cloud/shielded-vm) on this instance. Shielded VM provides verifiable integrity to prevent against malware and rootkits. Defaults to disabled. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "auto_delete",
					Description: `(Optional) Whether or not the disk should be auto-deleted. This defaults to true.`,
				},
				resource.Attribute{
					Name:        "boot",
					Description: `(Optional) Indicates that this is a boot disk.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(Optional) A unique device name that is reflected into the /dev/ tree of a Linux operating system running within the instance. If not specified, the server chooses a default device name to apply to this disk.`,
				},
				resource.Attribute{
					Name:        "disk_name",
					Description: `(Optional) Name of the disk. When not provided, this defaults to the name of the instance.`,
				},
				resource.Attribute{
					Name:        "source_image",
					Description: `(Required if source not set) The image from which to initialize this disk. This can be one of: the image's ` + "`" + `self_link` + "`" + `, ` + "`" + `projects/{project}/global/images/{image}` + "`" + `, ` + "`" + `projects/{project}/global/images/family/{family}` + "`" + `, ` + "`" + `global/images/{image}` + "`" + `, ` + "`" + `global/images/family/{family}` + "`" + `, ` + "`" + `family/{family}` + "`" + `, ` + "`" + `{project}/{family}` + "`" + `, ` + "`" + `{project}/{image}` + "`" + `, ` + "`" + `{family}` + "`" + `, or ` + "`" + `{image}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Optional) Specifies the disk interface to use for attaching this disk.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) The mode in which to attach this disk, either READ_WRITE or READ_ONLY. If you are attaching or creating a boot disk, this must read-write mode.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required if source_image not set) The name (`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Optional) The GCE disk type. Can be either ` + "`" + `"pd-ssd"` + "`" + `, ` + "`" + `"local-ssd"` + "`" + `, or ` + "`" + `"pd-standard"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `(Optional) The size of the image in gigabytes. If not specified, it will inherit the size of its base image.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of GCE disk, can be either ` + "`" + `"SCRATCH"` + "`" + ` or ` + "`" + `"PERSISTENT"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_encryption_key",
					Description: `(Optional) Encrypts or decrypts a disk using a customer-supplied encryption key. If you are creating a new disk, this field encrypts the new disk using an encryption key that you provide. If you are attaching an existing disk that is already encrypted, this field decrypts the disk using the customer-supplied encryption key. If you encrypt a disk using a customer-supplied key, you must provide the same key again when you attempt to use this resource at a later time. For example, you must provide the key when you create a snapshot or an image from the disk or when you attach the disk to a virtual machine instance. If you do not provide an encryption key, then the disk will be encrypted using an automatically generated key and you do not need to provide a key to use the disk later. Instance templates do not store customer-supplied encryption keys, so you cannot use your own keys to encrypt disks in a managed instance group. The ` + "`" + `disk_encryption_key` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "kms_key_self_link",
					Description: `(Optional) The self link of the encryption key that is stored in Google Cloud KMS The ` + "`" + `network_interface` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) The name or self_link of the network to attach this interface to. Use ` + "`" + `network` + "`" + ` attribute for Legacy or Auto subnetted networks and ` + "`" + `subnetwork` + "`" + ` for custom subnetted networks.`,
				},
				resource.Attribute{
					Name:        "subnetwork",
					Description: `(Optional) the name of the subnetwork to attach this interface to. The subnetwork must exist in the same ` + "`" + `region` + "`" + ` this instance will be created in. Either ` + "`" + `network` + "`" + ` or ` + "`" + `subnetwork` + "`" + ` must be provided.`,
				},
				resource.Attribute{
					Name:        "subnetwork_project",
					Description: `(Optional) The ID of the project in which the subnetwork belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "network_ip",
					Description: `(Optional) The private IP address to assign to the instance. If empty, the address will be automatically assigned.`,
				},
				resource.Attribute{
					Name:        "access_config",
					Description: `(Optional) Access configurations, i.e. IPs via which this instance can be accessed via the Internet. Omit to ensure that the instance is not accessible from the Internet (this means that ssh provisioners will not work unless you are running Terraform can send traffic to the instance's network (e.g. via tunnel or because it is running on another cloud instance on that network). This block can be repeated multiple times. Structure documented below.`,
				},
				resource.Attribute{
					Name:        "alias_ip_range",
					Description: `(Optional) An array of alias IP ranges for this network interface. Can only be specified for network interfaces on subnet-mode networks. Structure documented below. The ` + "`" + `access_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "nat_ip",
					Description: `(Optional) The IP address that will be 1:1 mapped to the instance's network ip. If not given, one will be generated.`,
				},
				resource.Attribute{
					Name:        "network_tier",
					Description: `(Optional) The [networking tier][network-tier] used for configuring this instance template. This field can take the following values: PREMIUM or STANDARD. If this field is not specified, it is assumed to be PREMIUM. The ` + "`" + `alias_ip_range` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "ip_cidr_range",
					Description: `The IP CIDR range represented by this alias IP range. This IP CIDR range must belong to the specified subnetwork and cannot contain IP addresses reserved by system or used by other network interfaces. At the time of writing only a netmask (e.g. /24) may be supplied, with a CIDR format resulting in an API error.`,
				},
				resource.Attribute{
					Name:        "subnetwork_range_name",
					Description: `(Optional) The subnetwork secondary range name specifying the secondary range from which to allocate the IP CIDR range for this alias IP range. If left unspecified, the primary range of the subnetwork will be used. The ` + "`" + `service_account` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) The service account e-mail address. If not given, the default Google Compute Engine service account is used.`,
				},
				resource.Attribute{
					Name:        "scopes",
					Description: `(Required) A list of service scopes. Both OAuth2 URLs and gcloud short names are supported. To allow full access to all Cloud APIs, use the ` + "`" + `cloud-platform` + "`" + ` scope. See a complete list of scopes [here](https://cloud.google.com/sdk/gcloud/reference/alpha/compute/instances/set-scopes#--scopes). The [service accounts documentation](https://cloud.google.com/compute/docs/access/service-accounts#accesscopesiam) explains that access scopes are the legacy method of specifying permissions for your instance. If you are following best practices and using IAM roles to grant permissions to service accounts, then you can define this field as an empty list. The ` + "`" + `scheduling` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "automatic_restart",
					Description: `(Optional) Specifies whether the instance should be automatically restarted if it is terminated by Compute Engine (not terminated by a user). This defaults to true.`,
				},
				resource.Attribute{
					Name:        "on_host_maintenance",
					Description: `(Optional) Defines the maintenance behavior for this instance.`,
				},
				resource.Attribute{
					Name:        "preemptible",
					Description: `(Optional) Allows instance to be preempted. This defaults to false. Read more on this [here](https://cloud.google.com/compute/docs/instances/preemptible).`,
				},
				resource.Attribute{
					Name:        "node_affinities",
					Description: `(Optional) Specifies node affinities or anti-affinities to determine which sole-tenant nodes your instances and managed instance groups will use as host systems. Read more on sole-tenant node creation [here](https://cloud.google.com/compute/docs/nodes/create-nodes). Structure documented below. The ` + "`" + `guest_accelerator` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "metadata_fingerprint",
					Description: `The unique fingerprint of the metadata.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource.`,
				},
				resource.Attribute{
					Name:        "tags_fingerprint",
					Description: `The unique fingerprint of the tags. [1]: /docs/providers/google/r/compute_instance_group_manager.html [2]: /docs/configuration/resources.html#lifecycle ## Import Instance templates can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_instance_template.default appserver-template ` + "`" + `` + "`" + `` + "`" + ` [custom-vm-types]: https://cloud.google.com/dataproc/docs/concepts/compute/custom-machine-types [network-tier]: https://cloud.google.com/network-tiers/docs/overview`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata_fingerprint",
					Description: `The unique fingerprint of the metadata.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource.`,
				},
				resource.Attribute{
					Name:        "tags_fingerprint",
					Description: `The unique fingerprint of the tags. [1]: /docs/providers/google/r/compute_instance_group_manager.html [2]: /docs/configuration/resources.html#lifecycle ## Import Instance templates can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_instance_template.default appserver-template ` + "`" + `` + "`" + `` + "`" + ` [custom-vm-types]: https://cloud.google.com/dataproc/docs/concepts/compute/custom-machine-types [network-tier]: https://cloud.google.com/network-tiers/docs/overview`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_interconnect_attachment",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Represents an InterconnectAttachment (VLAN attachment) resource.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"interconnect",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "router",
					Description: `(Required) URL of the cloud router to be used for dynamic routing. This router must be in the same region as this InterconnectAttachment. The InterconnectAttachment will automatically connect the Interconnect to the network & region within which the Cloud Router is configured.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "admin_enabled",
					Description: `(Optional) Whether the VLAN attachment is enabled or disabled. When using PARTNER type this will Pre-Activate the interconnect attachment`,
				},
				resource.Attribute{
					Name:        "interconnect",
					Description: `(Optional) URL of the underlying Interconnect object that this attachment's traffic will traverse through. Required if type is DEDICATED, must not be set if type is PARTNER.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional) Provisioned bandwidth capacity for the interconnect attachment. For attachments of type DEDICATED, the user can set the bandwidth. For attachments of type PARTNER, the Google Partner that is operating the interconnect must set the bandwidth. Output only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED, Defaults to BPS_10G`,
				},
				resource.Attribute{
					Name:        "edge_availability_domain",
					Description: `(Optional) Desired availability domain for the attachment. Only available for type PARTNER, at creation time. For improved reliability, customers should configure a pair of attachments with one per availability domain. The selected availability domain will be provided to the Partner via the pairing key so that the provisioned circuit will lie in the specified domain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of InterconnectAttachment you wish to create. Defaults to DEDICATED.`,
				},
				resource.Attribute{
					Name:        "candidate_subnets",
					Description: `(Optional) Up to 16 candidate prefixes that can be used to restrict the allocation of cloudRouterIpAddress and customerRouterIpAddress for this attachment. All prefixes must be within link-local address space (169.254.0.0/16) and must be /29 or shorter (/28, /27, etc). Google will attempt to select an unused /29 from the supplied candidate prefix(es). The request will fail if all possible /29s are in use on Google's edge. If not supplied, Google will randomly select an unused /29 from all of link-local space.`,
				},
				resource.Attribute{
					Name:        "vlan_tag8021q",
					Description: `(Optional) The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. When using PARTNER type this will be managed upstream.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Region where the regional interconnect attachment resides.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cloud_router_ip_address",
					Description: `IPv4 address + prefix length to be configured on Cloud Router Interface for this interconnect attachment.`,
				},
				resource.Attribute{
					Name:        "customer_router_ip_address",
					Description: `IPv4 address + prefix length to be configured on the customer router subinterface for this interconnect attachment.`,
				},
				resource.Attribute{
					Name:        "pairing_key",
					Description: `[Output only for type PARTNER. Not present for DEDICATED]. The opaque identifier of an PARTNER attachment used to initiate provisioning with a selected partner. Of the form "XXXXX/region/domain"`,
				},
				resource.Attribute{
					Name:        "partner_asn",
					Description: `[Output only for type PARTNER. Not present for DEDICATED]. Optional BGP ASN for the router that should be supplied by a layer 3 Partner if they configured BGP on behalf of the customer.`,
				},
				resource.Attribute{
					Name:        "private_interconnect_info",
					Description: `Information specific to an InterconnectAttachment. This property is populated if the interconnect that this is attached to is of type DEDICATED. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `[Output Only] The current state of this attachment's functionality.`,
				},
				resource.Attribute{
					Name:        "google_reference_id",
					Description: `Google reference ID, to be used when raising support tickets with Google or otherwise to debug backend connectivity issues.`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. The ` + "`" + `private_interconnect_info` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "tag8021q",
					Description: `802.1q encapsulation tag to be used for traffic between Google and the customer, going to and from this network and region. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import InterconnectAttachment can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_interconnect_attachment.default projects/{{project}}/regions/{{region}}/interconnectAttachments/{{name}} $ terraform import google_compute_interconnect_attachment.default {{project}}/{{region}}/{{name}} $ terraform import google_compute_interconnect_attachment.default {{region}}/{{name}} $ terraform import google_compute_interconnect_attachment.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_router_ip_address",
					Description: `IPv4 address + prefix length to be configured on Cloud Router Interface for this interconnect attachment.`,
				},
				resource.Attribute{
					Name:        "customer_router_ip_address",
					Description: `IPv4 address + prefix length to be configured on the customer router subinterface for this interconnect attachment.`,
				},
				resource.Attribute{
					Name:        "pairing_key",
					Description: `[Output only for type PARTNER. Not present for DEDICATED]. The opaque identifier of an PARTNER attachment used to initiate provisioning with a selected partner. Of the form "XXXXX/region/domain"`,
				},
				resource.Attribute{
					Name:        "partner_asn",
					Description: `[Output only for type PARTNER. Not present for DEDICATED]. Optional BGP ASN for the router that should be supplied by a layer 3 Partner if they configured BGP on behalf of the customer.`,
				},
				resource.Attribute{
					Name:        "private_interconnect_info",
					Description: `Information specific to an InterconnectAttachment. This property is populated if the interconnect that this is attached to is of type DEDICATED. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `[Output Only] The current state of this attachment's functionality.`,
				},
				resource.Attribute{
					Name:        "google_reference_id",
					Description: `Google reference ID, to be used when raising support tickets with Google or otherwise to debug backend connectivity issues.`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. The ` + "`" + `private_interconnect_info` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "tag8021q",
					Description: `802.1q encapsulation tag to be used for traffic between Google and the customer, going to and from this network and region. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import InterconnectAttachment can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_interconnect_attachment.default projects/{{project}}/regions/{{region}}/interconnectAttachments/{{name}} $ terraform import google_compute_interconnect_attachment.default {{project}}/{{region}}/{{name}} $ terraform import google_compute_interconnect_attachment.default {{region}}/{{name}} $ terraform import google_compute_interconnect_attachment.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_managed_ssl_certificate",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `An SslCertificate resource, used for HTTPS load balancing.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"managed",
				"ssl",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "managed",
					Description: `(Optional) Properties relevant to a managed certificate. These will be used if the certificate is managed (as indicated by a value of ` + "`" + `MANAGED` + "`" + ` in ` + "`" + `type` + "`" + `). Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Enum field whose value is always ` + "`" + `MANAGED` + "`" + ` - used to signal to the API which type this is.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `managed` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "domains",
					Description: `(Required) Domains for which a managed SSL certificate will be valid. Currently, there can only be one domain in this list. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `The unique identifier for the resource.`,
				},
				resource.Attribute{
					Name:        "subject_alternative_names",
					Description: `Domains associated with the certificate via Subject Alternative Name.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `Expire time of the certificate.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 6 minutes. - ` + "`" + `delete` + "`" + ` - Default is 30 minutes. ## Import ManagedSslCertificate can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import -provider=google-beta google_compute_managed_ssl_certificate.default projects/{{project}}/global/sslCertificates/{{name}} $ terraform import -provider=google-beta google_compute_managed_ssl_certificate.default {{project}}/{{name}} $ terraform import -provider=google-beta google_compute_managed_ssl_certificate.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `The unique identifier for the resource.`,
				},
				resource.Attribute{
					Name:        "subject_alternative_names",
					Description: `Domains associated with the certificate via Subject Alternative Name.`,
				},
				resource.Attribute{
					Name:        "expire_time",
					Description: `Expire time of the certificate.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 6 minutes. - ` + "`" + `delete` + "`" + ` - Default is 30 minutes. ## Import ManagedSslCertificate can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import -provider=google-beta google_compute_managed_ssl_certificate.default projects/{{project}}/global/sslCertificates/{{name}} $ terraform import -provider=google-beta google_compute_managed_ssl_certificate.default {{project}}/{{name}} $ terraform import -provider=google-beta google_compute_managed_ssl_certificate.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_network",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Manages a VPC network or legacy network resource on GCP.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource. The resource must be recreated to modify this field.`,
				},
				resource.Attribute{
					Name:        "ipv4_range",
					Description: `(Optional, Deprecated) If this field is specified, a deprecated legacy network is created. You will no longer be able to create a legacy network on Feb 1, 2020. See the [legacy network docs](https://cloud.google.com/vpc/docs/legacy) for more details. The range of internal addresses that are legal on this legacy network. This range is a CIDR specification, for example: ` + "`" + `192.168.0.0/16` + "`" + `. The resource must be recreated to modify this field.`,
				},
				resource.Attribute{
					Name:        "auto_create_subnetworks",
					Description: `(Optional) When set to ` + "`" + `true` + "`" + `, the network is created in "auto subnet mode" and it will create a subnet for each region automatically across the ` + "`" + `10.128.0.0/9` + "`" + ` address range. When set to ` + "`" + `false` + "`" + `, the network is created in "custom subnet mode" so the user can explicitly connect subnetwork resources.`,
				},
				resource.Attribute{
					Name:        "routing_mode",
					Description: `(Optional) The network-wide routing mode to use. If set to ` + "`" + `REGIONAL` + "`" + `, this network's cloud routers will only advertise routes with subnetworks of this network in the same region as the router. If set to ` + "`" + `GLOBAL` + "`" + `, this network's cloud routers will advertise routes with all subnetworks of this network, across regions.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "delete_default_routes_on_create",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + `, default routes (` + "`" + `0.0.0.0/0` + "`" + `) will be deleted immediately after network creation. Defaults to ` + "`" + `false` + "`" + `. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "gateway_ipv4",
					Description: `The gateway address for default routing out of the network. This value is selected by GCP.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Network can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_network.default projects/{{project}}/global/networks/{{name}} $ terraform import google_compute_network.default {{project}}/{{name}} $ terraform import google_compute_network.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "gateway_ipv4",
					Description: `The gateway address for default routing out of the network. This value is selected by GCP.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Network can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_network.default projects/{{project}}/global/networks/{{name}} $ terraform import google_compute_network.default {{project}}/{{name}} $ terraform import google_compute_network.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_network_endpoint",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `A Network endpoint represents a IP address and port combination that is part of a specific network endpoint group (NEG).`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"network",
				"endpoint",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance",
					Description: `(Required) The name for a specific VM instance that the IP address belongs to. This is required for network endpoints of type GCE_VM_IP_PORT. The instance must be in the same zone of network endpoint group.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Port number of network endpoint.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) IPv4 address of network endpoint. The IP address must belong to a VM in GCE (either the primary IP or as part of an aliased IP range).`,
				},
				resource.Attribute{
					Name:        "network_endpoint_group",
					Description: `(Required) The network endpoint group this endpoint is part of. - - -`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Zone where the containing network endpoint group is located.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import NetworkEndpoint can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_network_endpoint.default projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{network_endpoint_group}}/{{instance}}/{{ip_address}}/{{port}} $ terraform import google_compute_network_endpoint.default {{project}}/{{zone}}/{{network_endpoint_group}}/{{instance}}/{{ip_address}}/{{port}} $ terraform import google_compute_network_endpoint.default {{zone}}/{{network_endpoint_group}}/{{instance}}/{{ip_address}}/{{port}} $ terraform import google_compute_network_endpoint.default {{network_endpoint_group}}/{{instance}}/{{ip_address}}/{{port}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_network_endpoint_group",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Network endpoint groups (NEGs) are zonal resources that represent collections of IP address and port combinations for GCP resources within a single subnet.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"network",
				"endpoint",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) The network to which all network endpoints in the NEG belong. Uses "default" project network if unspecified. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource. Provide this property when you create the resource.`,
				},
				resource.Attribute{
					Name:        "network_endpoint_type",
					Description: `(Optional) Type of network endpoints in this network endpoint group. Currently the only supported value is GCE_VM_IP_PORT.`,
				},
				resource.Attribute{
					Name:        "subnetwork",
					Description: `(Optional) Optional subnetwork to which all network endpoints in the NEG belong.`,
				},
				resource.Attribute{
					Name:        "default_port",
					Description: `(Optional) The default port used if the port number is not specified in the network endpoint.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Zone where the network endpoint group is located.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Number of network endpoints in the network endpoint group.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import NetworkEndpointGroup can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_network_endpoint_group.default projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{name}} $ terraform import google_compute_network_endpoint_group.default {{project}}/{{zone}}/{{name}} $ terraform import google_compute_network_endpoint_group.default {{zone}}/{{name}} $ terraform import google_compute_network_endpoint_group.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "size",
					Description: `Number of network endpoints in the network endpoint group.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import NetworkEndpointGroup can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_network_endpoint_group.default projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{name}} $ terraform import google_compute_network_endpoint_group.default {{project}}/{{zone}}/{{name}} $ terraform import google_compute_network_endpoint_group.default {{zone}}/{{name}} $ terraform import google_compute_network_endpoint_group.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_network_peering",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Manages a network peering within GCE.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"network",
				"peering",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the peering.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) Resource link of the network to add a peering to.`,
				},
				resource.Attribute{
					Name:        "peer_network",
					Description: `(Required) Resource link of the peer network.`,
				},
				resource.Attribute{
					Name:        "auto_create_routes",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + `, the routes between the two networks will be created and managed automatically. Defaults to ` + "`" + `true` + "`" + `. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State for the peering.`,
				},
				resource.Attribute{
					Name:        "state_details",
					Description: `Details about the current state of the peering.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "state",
					Description: `State for the peering.`,
				},
				resource.Attribute{
					Name:        "state_details",
					Description: `Details about the current state of the peering.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_node_group",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Represents a NodeGroup resource to manage a group of sole-tenant nodes.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"node",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "node_template",
					Description: `(Required) The URL of the node template to which this node group belongs.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The total number of nodes in the node group. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional textual description of the resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the resource.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Zone where this node group is located`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import NodeGroup can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_node_group.default projects/{{project}}/zones/{{zone}}/nodeGroups/{{name}} $ terraform import google_compute_node_group.default {{project}}/{{zone}}/{{name}} $ terraform import google_compute_node_group.default {{zone}}/{{name}} $ terraform import google_compute_node_group.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import NodeGroup can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_node_group.default projects/{{project}}/zones/{{zone}}/nodeGroups/{{name}} $ terraform import google_compute_node_group.default {{project}}/{{zone}}/{{name}} $ terraform import google_compute_node_group.default {{zone}}/{{name}} $ terraform import google_compute_node_group.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_node_template",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Represents a NodeTemplate resource.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"node",
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional textual description of the resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the resource.`,
				},
				resource.Attribute{
					Name:        "node_affinity_labels",
					Description: `(Optional) Labels to use for node affinity, which will be used in instance scheduling.`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `(Optional) Node type to use for nodes group that are created from this template. Only one of nodeTypeFlexibility and nodeType can be specified.`,
				},
				resource.Attribute{
					Name:        "node_type_flexibility",
					Description: `(Optional) Flexible properties for the desired node type. Node groups that use this node template will create nodes of a type that matches these properties. Only one of nodeTypeFlexibility and nodeType can be specified. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "server_binding",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) The server binding policy for nodes using this template. Determines where the nodes should restart following a maintenance event. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Region where nodes using the node template will be created. If it is not provided, the provider region is used.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `node_type_flexibility` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "cpus",
					Description: `(Optional) Number of virtual CPUs to use.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Optional) Physical memory available to the node, defined in MB.`,
				},
				resource.Attribute{
					Name:        "local_ssd",
					Description: `Use local SSD The ` + "`" + `server_binding` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of server binding policy. If ` + "`" + `RESTART_NODE_ON_ANY_SERVER` + "`" + `, nodes using this template will restart on any physical server following a maintenance event. If ` + "`" + `RESTART_NODE_ON_MINIMAL_SERVER` + "`" + `, nodes using this template will restart on the same physical server following a maintenance event, instead of being live migrated to or restarted on a new physical server. This option may be useful if you are using software licenses tied to the underlying server characteristics such as physical sockets or cores, to avoid the need for additional licenses when maintenance occurs. However, VMs on such nodes will experience outages while maintenance is applied. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import NodeTemplate can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_node_template.default projects/{{project}}/regions/{{region}}/nodeTemplates/{{name}} $ terraform import google_compute_node_template.default {{project}}/{{region}}/{{name}} $ terraform import google_compute_node_template.default {{region}}/{{name}} $ terraform import google_compute_node_template.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import NodeTemplate can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_node_template.default projects/{{project}}/regions/{{region}}/nodeTemplates/{{name}} $ terraform import google_compute_node_template.default {{project}}/{{region}}/{{name}} $ terraform import google_compute_node_template.default {{region}}/{{name}} $ terraform import google_compute_node_template.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_project_default_network_tier",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Configures the default network tier for a project.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"project",
				"default",
				"network",
				"tier",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "network_tier",
					Description: `(Required) The default network tier to be configured for the project. This field can take the following values: ` + "`" + `PREMIUM` + "`" + ` or ` + "`" + `STANDARD` + "`" + `. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference Only the arguments listed above are exposed as attributes. ## Import This resource can be imported using the project ID: ` + "`" + `terraform import google_compute_project_default_network_tier.default project-id` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_project_metadata",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Manages common instance metadata`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"project",
				"metadata",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metadata",
					Description: `(Required) A series of key value pairs. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference Only the arguments listed above are exposed as attributes. ## Import This resource can be imported using the project ID: ` + "`" + `terraform import google_compute_project_metadata.foo my-project-id` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_project_metadata_item",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Manages a single key/value pair on common instance metadata`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"project",
				"metadata",
				"item",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The metadata key to set.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value to set for the given metadata key. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference Only the arguments listed above are exposed as attributes. ## Import Project metadata items can be imported using the ` + "`" + `key` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_project_metadata_item.default my_metadata ` + "`" + `` + "`" + `` + "`" + ` ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 5 minutes. - ` + "`" + `update` + "`" + ` - Default is 5 minutes. - ` + "`" + `delete` + "`" + ` - Default is 5 minutes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_region_autoscaler",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Represents an Autoscaler resource.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"region",
				"autoscaler",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource. The name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "autoscaling_policy",
					Description: `(Required) The configuration parameters for the autoscaling algorithm. You can define one or more of the policies for an autoscaler: cpuUtilization, customMetricUtilizations, and loadBalancingUtilization. If none of these are specified, the default will be to autoscale based on cpuUtilization to 0.6 or 60%. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) URL of the managed instance group that this autoscaler will scale. The ` + "`" + `autoscaling_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "min_replicas",
					Description: `(Required) The minimum number of replicas that the autoscaler can scale down to. This cannot be less than 0. If not provided, autoscaler will choose a default value depending on maximum number of instances allowed.`,
				},
				resource.Attribute{
					Name:        "max_replicas",
					Description: `(Required) The maximum number of instances that the autoscaler can scale up to. This is required when creating or updating an autoscaler. The maximum number of replicas should not be lower than minimal number of replicas.`,
				},
				resource.Attribute{
					Name:        "cooldown_period",
					Description: `(Optional) The number of seconds that the autoscaler should wait before it starts collecting information from a new instance. This prevents the autoscaler from collecting information when the instance is initializing, during which the collected usage would not be reliable. The default time autoscaler waits is 60 seconds. Virtual machine initialization times might vary because of numerous factors. We recommend that you test how long an instance may take to initialize. To do this, create an instance and time the startup process.`,
				},
				resource.Attribute{
					Name:        "cpu_utilization",
					Description: `(Optional) Defines the CPU utilization policy that allows the autoscaler to scale based on the average CPU utilization of a managed instance group. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(Optional) Defines the CPU utilization policy that allows the autoscaler to scale based on the average CPU utilization of a managed instance group. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "load_balancing_utilization",
					Description: `(Optional) Configuration parameters of autoscaling based on a load balancer. Structure is documented below. The ` + "`" + `cpu_utilization` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) The target CPU utilization that the autoscaler should maintain. Must be a float value in the range (0, 1]. If not specified, the default is 0.6. If the CPU level is below the target utilization, the autoscaler scales down the number of instances until it reaches the minimum number of instances you specified or until the average CPU of your instances reaches the target utilization. If the average CPU is above the target utilization, the autoscaler scales up until it reaches the maximum number of instances you specified or until the average utilization reaches the target utilization. The ` + "`" + `metric` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The identifier (type) of the Stackdriver Monitoring metric. The metric cannot have negative values. The metric must have a value type of INT64 or DOUBLE.`,
				},
				resource.Attribute{
					Name:        "single_instance_assignment",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) If scaling is based on a per-group metric value that represents the total amount of work to be done or resource usage, set this value to an amount assigned for a single instance of the scaled group. The autoscaler will keep the number of instances proportional to the value of this metric, the metric itself should not change value due to group resizing. For example, a good metric to use with the target is ` + "`" + `pubsub.googleapis.com/subscription/num_undelivered_messages` + "`" + ` or a custom metric exporting the total number of requests coming to your instances. A bad example would be a metric exporting an average or median latency, since this value can't include a chunk assignable to a single instance, it could be better used with utilization_target instead.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Optional) The target value of the metric that autoscaler should maintain. This must be a positive value. A utilization metric scales number of virtual machines handling requests to increase or decrease proportionally to the metric. For example, a good metric to use as a utilizationTarget is www.googleapis.com/compute/instance/network/received_bytes_count. The autoscaler will work to keep this value constant for each of the instances.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Defines how target utilization value is expressed for a Stackdriver Monitoring metric. Either GAUGE, DELTA_PER_SECOND, or DELTA_PER_MINUTE.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) A filter string to be used as the filter string for a Stackdriver Monitoring TimeSeries.list API call. This filter is used to select a specific TimeSeries for the purpose of autoscaling and to determine whether the metric is exporting per-instance or per-group data. You can only use the AND operator for joining selectors. You can only use direct equality comparison operator (=) without any functions for each selector. You can specify the metric in both the filter string and in the metric field. However, if specified in both places, the metric must be identical. The monitored resource type determines what kind of values are expected for the metric. If it is a gce_instance, the autoscaler expects the metric to include a separate TimeSeries for each instance in a group. In such a case, you cannot filter on resource labels. If the resource type is any other value, the autoscaler expects this metric to contain values that apply to the entire autoscaled instance group and resource label filtering can be performed to point autoscaler at the correct TimeSeries to scale upon. This is called a per-group metric for the purpose of autoscaling. If not specified, the type defaults to gce_instance. You should provide a filter that is selective enough to pick just one TimeSeries for the autoscaled group or for each of the instances (if you are using gce_instance resource type). If multiple TimeSeries are returned upon the query execution, the autoscaler will sum their respective values to obtain its scaling value. The ` + "`" + `load_balancing_utilization` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) Fraction of backend capacity utilization (set in HTTP(s) load balancing configuration) that autoscaler should maintain. Must be a positive float value. If not defined, the default is 0.8. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) URL of the region where the instance group resides.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import RegionAutoscaler can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_region_autoscaler.default projects/{{project}}/regions/{{region}}/autoscalers/{{name}} $ terraform import google_compute_region_autoscaler.default {{project}}/{{region}}/{{name}} $ terraform import google_compute_region_autoscaler.default {{region}}/{{name}} $ terraform import google_compute_region_autoscaler.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import RegionAutoscaler can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_region_autoscaler.default projects/{{project}}/regions/{{region}}/autoscalers/{{name}} $ terraform import google_compute_region_autoscaler.default {{project}}/{{region}}/{{name}} $ terraform import google_compute_region_autoscaler.default {{region}}/{{name}} $ terraform import google_compute_region_autoscaler.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_region_backend_service",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `A Region Backend Service defines a regionally-scoped group of virtual machines that will serve traffic for load balancing.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"region",
				"backend",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "health_checks",
					Description: `(Required) The list of HealthChecks for checking the health of the backend service. Currently at most one health check can be specified, and a health check is required. - - -`,
				},
				resource.Attribute{
					Name:        "backend",
					Description: `(Optional) The list of backends that serve this RegionBackendService. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource.`,
				},
				resource.Attribute{
					Name:        "failover_policy",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) Policy for failovers. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) The protocol this BackendService uses to communicate with backends. The possible values are TCP and UDP, and the default is TCP.`,
				},
				resource.Attribute{
					Name:        "session_affinity",
					Description: `(Optional) Type of session affinity to use. The default is NONE. Can be NONE, CLIENT_IP, CLIENT_IP_PROTO, or CLIENT_IP_PORT_PROTO. When the protocol is UDP, this field is not used.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The Region in which the created backend service should reside. If it is not provided, the provider region is used.`,
				},
				resource.Attribute{
					Name:        "timeout_sec",
					Description: `(Optional) How many seconds to wait for the backend before considering it a failed request. Default is 30 seconds. Valid range is [1, 86400].`,
				},
				resource.Attribute{
					Name:        "connection_draining_timeout_sec",
					Description: `(Optional) Time for which instance will be drained (not accept new connections, but still work to finish started).`,
				},
				resource.Attribute{
					Name:        "load_balancing_scheme",
					Description: `(Optional) This signifies what the ForwardingRule will be used for and can only be INTERNAL for RegionBackendServices`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `backend` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource. Provide this property when you create the resource.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) The fully-qualified URL of an Instance Group. This defines the list of instances that serve traffic. Member virtual machine instances from each instance group must live in the same zone as the instance group itself. No two backends in a backend service are allowed to use same Instance Group resource. Note that you must specify an Instance Group resource using the fully-qualified URL, rather than a partial URL. The instance group must be within the same region as the BackendService.`,
				},
				resource.Attribute{
					Name:        "failover",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) This field designates whether this is a failover backend. More than one failover backend can be configured for a given BackendService. The ` + "`" + `failover_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "disable_connection_drain_on_failover",
					Description: `(Optional) On failover or failback, this field indicates whether connection drain will be honored. Setting this to true has the following effect: connections to the old active pool are not drained. Connections to the new active pool use the timeout of 10 min (currently fixed). Setting to false has the following effect: both old and new connections will have a drain timeout of 10 min. This can be set to true only if the protocol is TCP. The default is false.`,
				},
				resource.Attribute{
					Name:        "drop_traffic_if_unhealthy",
					Description: `(Optional) This option is used only when no healthy VMs are detected in the primary and backup instance groups. When set to true, traffic is dropped. When set to false, new connections are sent across all VMs in the primary group. The default is false.`,
				},
				resource.Attribute{
					Name:        "failover_ratio",
					Description: `(Optional) The value of the field must be in [0, 1]. If the ratio of the healthy VMs in the primary backend is at or below this number, traffic arriving at the load-balanced IP will be directed to the failover backend. In case where 'failoverRatio' is not set or all the VMs in the backup backend are unhealthy, the traffic will be directed back to the primary backend in the "force" mode, where traffic will be spread to the healthy VMs with the best effort, or to all VMs when no VM is healthy. This field is only used with l4 load balancing. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import RegionBackendService can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_region_backend_service.default projects/{{project}}/regions/{{region}}/backendServices/{{name}} $ terraform import google_compute_region_backend_service.default {{project}}/{{region}}/{{name}} $ terraform import google_compute_region_backend_service.default {{region}}/{{name}} $ terraform import google_compute_region_backend_service.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "fingerprint",
					Description: `Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import RegionBackendService can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_region_backend_service.default projects/{{project}}/regions/{{region}}/backendServices/{{name}} $ terraform import google_compute_region_backend_service.default {{project}}/{{region}}/{{name}} $ terraform import google_compute_region_backend_service.default {{region}}/{{name}} $ terraform import google_compute_region_backend_service.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_region_disk",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Persistent disks are durable storage devices that function similarly to the physical disks in a desktop or a server.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"region",
				"disk",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "replica_zones",
					Description: `(Required) URLs of the zones where the disk should be replicated to. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource. Provide this property when you create the resource.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Labels to apply to this disk. A list of key->value pairs.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Size of the persistent disk, specified in GB. You can specify this field when creating a persistent disk using the sourceImage or sourceSnapshot parameter, or specify it alone to create an empty persistent disk. If you specify this field along with sourceImage or sourceSnapshot, the value of sizeGb must not be less than the size of the sourceImage or the size of the snapshot.`,
				},
				resource.Attribute{
					Name:        "physical_block_size_bytes",
					Description: `(Optional) Physical block size of the persistent disk, in bytes. If not present in a request, a default value is used. Currently supported sizes are 4096 and 16384, other sizes may be added in the future. If an unsupported value is requested, the error message will list the supported values for the caller's project.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) URL of the disk type resource describing which disk type to use to create the disk. Provide this when creating the disk.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) A reference to the region where the disk resides.`,
				},
				resource.Attribute{
					Name:        "disk_encryption_key",
					Description: `(Optional) Encrypts the disk using a customer-supplied encryption key. After you encrypt a disk with a customer-supplied key, you must provide the same key if you use the disk later (e.g. to create a disk snapshot or an image, or to attach the disk to a virtual machine). Customer-supplied encryption keys do not protect access to metadata of the disk. If you do not provide an encryption key when creating the disk, then the disk will be encrypted using an automatically generated key and you do not need to provide a key to use the disk later. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "snapshot",
					Description: `(Optional) The source snapshot used to create this disk. You can provide this as a partial or full URL to the resource. For example, the following are valid values:`,
				},
				resource.Attribute{
					Name:        "source_snapshot_encryption_key",
					Description: `(Optional) The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `disk_encryption_key` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "raw_key",
					Description: `(Optional) Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource.`,
				},
				resource.Attribute{
					Name:        "sha256",
					Description: `The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource.`,
				},
				resource.Attribute{
					Name:        "kms_key_name",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) The name of the encryption key that is stored in Google Cloud KMS. The ` + "`" + `source_snapshot_encryption_key` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "raw_key",
					Description: `(Optional) Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource.`,
				},
				resource.Attribute{
					Name:        "kms_key_name",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) The name of the encryption key that is stored in Google Cloud KMS.`,
				},
				resource.Attribute{
					Name:        "sha256",
					Description: `The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "label_fingerprint",
					Description: `The fingerprint used for optimistic locking of this resource. Used internally during updates.`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "last_attach_timestamp",
					Description: `Last attach timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "last_detach_timestamp",
					Description: `Last detach timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `Links to the users of the disk (attached instances) in form: project/zones/zone/instances/instance`,
				},
				resource.Attribute{
					Name:        "source_snapshot_id",
					Description: `The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to create this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 5 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import RegionDisk can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_region_disk.default projects/{{project}}/regions/{{region}}/disks/{{name}} $ terraform import google_compute_region_disk.default {{project}}/{{region}}/{{name}} $ terraform import google_compute_region_disk.default {{region}}/{{name}} $ terraform import google_compute_region_disk.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "label_fingerprint",
					Description: `The fingerprint used for optimistic locking of this resource. Used internally during updates.`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "last_attach_timestamp",
					Description: `Last attach timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "last_detach_timestamp",
					Description: `Last detach timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `Links to the users of the disk (attached instances) in form: project/zones/zone/instances/instance`,
				},
				resource.Attribute{
					Name:        "source_snapshot_id",
					Description: `The unique ID of the snapshot used to create this disk. This value identifies the exact snapshot that was used to create this persistent disk. For example, if you created the persistent disk from a snapshot that was later deleted and recreated under the same name, the source snapshot ID would identify the exact version of the snapshot that was used.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 5 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import RegionDisk can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_region_disk.default projects/{{project}}/regions/{{region}}/disks/{{name}} $ terraform import google_compute_region_disk.default {{project}}/{{region}}/{{name}} $ terraform import google_compute_region_disk.default {{region}}/{{name}} $ terraform import google_compute_region_disk.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_region_instance_group_manager",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Manages an Regional Instance Group within GCE.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"region",
				"instance",
				"group",
				"manager",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "base_instance_name",
					Description: `(Required) The base instance name to use for instances in this group. The value must be a valid [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters are lowercase letters, numbers, and hyphens (-). Instances are named by appending a hyphen and a random four-character string to the base instance name.`,
				},
				resource.Attribute{
					Name:        "instance_template",
					Description: `(Required, [GA](https://terraform.io/docs/providers/google/provider_versions.html)) The full URL to an instance template from which all new instances will be created. This field is only present in the ` + "`" + `google` + "`" + ` provider.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) Application versions managed by this instance group. Each version deals with a specific instance template, allowing canary release scenarios. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the instance group manager. Must be 1-63 characters long and comply with [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters include lowercase letters, numbers, and hyphens.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region where the managed instance group resides. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional textual description of the instance group manager.`,
				},
				resource.Attribute{
					Name:        "named_port",
					Description: `(Optional) The named port configuration. See the section below for details on configuration.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "target_size",
					Description: `(Optional) The target number of running instances for this managed instance group. This value should always be explicitly set unless this resource is attached to an autoscaler, in which case it should never be set. Defaults to ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "target_pools",
					Description: `(Optional) The full URL of all target pools to which new instances in the group are added. Updating the target pools attribute does not affect existing instances.`,
				},
				resource.Attribute{
					Name:        "wait_for_instances",
					Description: `(Optional) Whether to wait for all instances to be created/updated before returning. Note that if this is set to true and the operation does not succeed, Terraform will continue trying until it times out. ---`,
				},
				resource.Attribute{
					Name:        "auto_healing_policies",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) The autohealing policies for this managed instance group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).`,
				},
				resource.Attribute{
					Name:        "update_policy",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/beta/regionInstanceGroupManagers/patch)`,
				},
				resource.Attribute{
					Name:        "distribution_policy_zones",
					Description: `(Optional) The distribution policy for this managed instance group. You can specify one or more values. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups#selectingzones). - - - The ` + "`" + `update_policy` + "`" + ` block supports: ` + "`" + `` + "`" + `` + "`" + `hcl update_policy{ type = "PROACTIVE" instance_redistribution_type = "PROACTIVE" minimal_action = "REPLACE" max_surge_percent = 20 max_unavailable_fixed = 2 min_ready_sec = 50 } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "minimal_action",
					Description: `(Required) - Minimal action to be taken on an instance. Valid values are ` + "`" + `"RESTART"` + "`" + `, ` + "`" + `"REPLACE"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) - The type of update. Valid values are ` + "`" + `"OPPORTUNISTIC"` + "`" + `, ` + "`" + `"PROACTIVE"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "instance_redistribution_type",
					Description: `(Optional) - The instance redistribution policy for regional managed instance groups. Valid values are: ` + "`" + `"PROACTIVE"` + "`" + `, ` + "`" + `"NONE"` + "`" + `. If ` + "`" + `PROACTIVE` + "`" + ` (default), the group attempts to maintain an even distribution of VM instances across zones in the region. If ` + "`" + `NONE` + "`" + `, proactive redistribution is disabled.`,
				},
				resource.Attribute{
					Name:        "max_surge_fixed",
					Description: `(Optional), The maximum number of instances that can be created above the specified targetSize during the update process. Conflicts with ` + "`" + `max_surge_percent` + "`" + `. It has to be either 0 or at least equal to the number of zones. If fixed values are used, at least one of ` + "`" + `max_unavailable_fixed` + "`" + ` or ` + "`" + `max_surge_fixed` + "`" + ` must be greater than 0.`,
				},
				resource.Attribute{
					Name:        "max_surge_percent",
					Description: `(Optional), The maximum number of instances(calculated as percentage) that can be created above the specified targetSize during the update process. Conflicts with ` + "`" + `max_surge_fixed` + "`" + `. Percent value is only allowed for regional managed instance groups with size at least 10.`,
				},
				resource.Attribute{
					Name:        "max_unavailable_fixed",
					Description: `(Optional), The maximum number of instances that can be unavailable during the update process. Conflicts with ` + "`" + `max_unavailable_percent` + "`" + `. It has to be either 0 or at least equal to the number of zones. If fixed values are used, at least one of ` + "`" + `max_unavailable_fixed` + "`" + ` or ` + "`" + `max_surge_fixed` + "`" + ` must be greater than 0.`,
				},
				resource.Attribute{
					Name:        "max_unavailable_percent",
					Description: `(Optional), The maximum number of instances(calculated as percentage) that can be unavailable during the update process. Conflicts with ` + "`" + `max_unavailable_fixed` + "`" + `. Percent value is only allowed for regional managed instance groups with size at least 10.`,
				},
				resource.Attribute{
					Name:        "min_ready_sec",
					Description: `(Optional), Minimum number of seconds to wait for after a newly created instance becomes available. This value must be from range [0, 3600] - - - The ` + "`" + `named_port` + "`" + ` block supports: (Include a ` + "`" + `named_port` + "`" + ` block for each named-port required).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the port.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The port number. - - - The ` + "`" + `auto_healing_policies` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `(Required) The health check resource that signals autohealing.`,
				},
				resource.Attribute{
					Name:        "initial_delay_sec",
					Description: `(Required) The number of seconds that the managed instance group waits before it applies autohealing policies to new instances or recently recreated instances. Between 0 and 3600. The ` + "`" + `version` + "`" + ` block supports: ` + "`" + `` + "`" + `` + "`" + `hcl version { name = "appserver-canary" instance_template = "${google_compute_instance_template.appserver-canary.self_link}" target_size { fixed = 1 } } ` + "`" + `` + "`" + `` + "`" + ` ` + "`" + `` + "`" + `` + "`" + `hcl version { name = "appserver-canary" instance_template = "${google_compute_instance_template.appserver-canary.self_link}" target_size { percent = 20 } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) - Version name.`,
				},
				resource.Attribute{
					Name:        "instance_template",
					Description: `(Required) - The full URL to an instance template from which all new instances of this version will be created.`,
				},
				resource.Attribute{
					Name:        "target_size",
					Description: `(Optional) - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below. -> Exactly one ` + "`" + `version` + "`" + ` you specify must not have a ` + "`" + `target_size` + "`" + ` specified. During a rolling update, the instance group manager will fulfill the ` + "`" + `target_size` + "`" + ` constraints of every other ` + "`" + `version` + "`" + `, and any remaining instances will be provisioned with the version where ` + "`" + `target_size` + "`" + ` is unset. The ` + "`" + `target_size` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "fixed",
					Description: `(Optional), The number of instances which are managed for this version. Conflicts with ` + "`" + `percent` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "percent",
					Description: `(Optional), The number of instances (calculated as percentage) which are managed for this version. Conflicts with ` + "`" + `fixed` + "`" + `. Note that when using ` + "`" + `percent` + "`" + `, rounding will be in favor of explicitly set ` + "`" + `target_size` + "`" + ` values; a managed instance group with 2 instances and 2 ` + "`" + `version` + "`" + `s, one of which has a ` + "`" + `target_size.percent` + "`" + ` of ` + "`" + `60` + "`" + ` will create 2 instances of that ` + "`" + `version` + "`" + `. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the instance group manager.`,
				},
				resource.Attribute{
					Name:        "instance_group",
					Description: `The full URL of the instance group created by the manager.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URL of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 5 minutes. - ` + "`" + `update` + "`" + ` - Default is 5 minutes. - ` + "`" + `delete` + "`" + ` - Default is 15 minutes. ## Import Instance group managers can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_region_instance_group_manager.appserver appserver-igm ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of the instance group manager.`,
				},
				resource.Attribute{
					Name:        "instance_group",
					Description: `The full URL of the instance group created by the manager.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URL of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 5 minutes. - ` + "`" + `update` + "`" + ` - Default is 5 minutes. - ` + "`" + `delete` + "`" + ` - Default is 15 minutes. ## Import Instance group managers can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_region_instance_group_manager.appserver appserver-igm ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_resource_policy",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `A policy that can be attached to a resource to specify or schedule actions on that resource.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"resource",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "snapshot_schedule_policy",
					Description: `(Optional) Policy for creating snapshots of persistent disks. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Region where resource policy resides.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `snapshot_schedule_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `(Required) Contains one of an ` + "`" + `hourlySchedule` + "`" + `, ` + "`" + `dailySchedule` + "`" + `, or ` + "`" + `weeklySchedule` + "`" + `. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "retention_policy",
					Description: `(Optional) Retention policy applied to snapshots created by this resource policy. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "snapshot_properties",
					Description: `(Optional) Properties with which the snapshots are created, such as labels. Structure is documented below. The ` + "`" + `schedule` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "hourly_schedule",
					Description: `(Optional) The policy will execute every nth hour starting at the specified time. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "daily_schedule",
					Description: `(Optional) The policy will execute every nth day at the specified time. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "weekly_schedule",
					Description: `(Optional) Allows specifying a snapshot time for each day of the week. Structure is documented below. The ` + "`" + `hourly_schedule` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "hours_in_cycle",
					Description: `(Required) The number of hours between snapshots.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Required) Time within the window to start the operations. It must be in format "HH:MM", where HH : [00-23] and MM : [00-00] GMT. The ` + "`" + `daily_schedule` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "days_in_cycle",
					Description: `(Required) The number of days between snapshots.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Required) This must be in UTC format that resolves to one of 00:00, 04:00, 08:00, 12:00, 16:00, or 20:00. For example, both 13:00-5 and 08:00 are valid. The ` + "`" + `weekly_schedule` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "day_of_weeks",
					Description: `(Required) May contain up to seven (one for each day of the week) snapshot times. Structure is documented below. The ` + "`" + `day_of_weeks` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Required) Time within the window to start the operations. It must be in format "HH:MM", where HH : [00-23] and MM : [00-00] GMT.`,
				},
				resource.Attribute{
					Name:        "day",
					Description: `(Required) The day of the week to create the snapshot. e.g. MONDAY The ` + "`" + `retention_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "max_retention_days",
					Description: `(Required) Maximum age of the snapshot that is allowed to be kept.`,
				},
				resource.Attribute{
					Name:        "on_source_disk_delete",
					Description: `(Optional) Specifies the behavior to apply to scheduled snapshots when the source disk is deleted. Valid options are KEEP_AUTO_SNAPSHOTS and APPLY_RETENTION_POLICY The ` + "`" + `snapshot_properties` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key-value pairs.`,
				},
				resource.Attribute{
					Name:        "storage_locations",
					Description: `(Optional) GCS bucket location in which to store the snapshot (regional or multi-regional).`,
				},
				resource.Attribute{
					Name:        "guest_flush",
					Description: `(Optional) Whether to perform a 'guest aware' snapshot. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import ResourcePolicy can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import -provider=google-beta google_compute_resource_policy.default projects/{{project}}/regions/{{region}}/resourcePolicies/{{name}} $ terraform import -provider=google-beta google_compute_resource_policy.default {{project}}/{{region}}/{{name}} $ terraform import -provider=google-beta google_compute_resource_policy.default {{region}}/{{name}} $ terraform import -provider=google-beta google_compute_resource_policy.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_route",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Represents a Route resource.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"route",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dest_range",
					Description: `(Required) The destination range of outgoing packets that this route applies to. Only IPv4 is supported.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) The network that this route applies to. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource. Provide this property when you create the resource.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) The priority of this route. Priority is used to break ties in cases where there is more than one matching route of equal prefix length. In the case of two routes with equal prefix length, the one with the lowest-numbered priority value wins. Default value is 1000. Valid range is 0 through 65535.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of instance tags to which this route applies.`,
				},
				resource.Attribute{
					Name:        "next_hop_gateway",
					Description: `(Optional) URL to a gateway that should handle matching packets. Currently, you can only specify the internet gateway, using a full or partial valid URL:`,
				},
				resource.Attribute{
					Name:        "next_hop_instance",
					Description: `(Optional) URL to an instance that should handle matching packets. You can specify this as a full or partial URL. For example:`,
				},
				resource.Attribute{
					Name:        "next_hop_ip",
					Description: `(Optional) Network IP address of an instance that should handle matching packets.`,
				},
				resource.Attribute{
					Name:        "next_hop_vpn_tunnel",
					Description: `(Optional) URL to a VpnTunnel that should handle matching packets.`,
				},
				resource.Attribute{
					Name:        "next_hop_ilb",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) The URL to a forwarding rule of type loadBalancingScheme=INTERNAL that should handle matching packets. You can only specify the forwarding rule as a partial or full URL. For example, the following are all valid URLs: https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule regions/region/forwardingRules/forwardingRule Note that this can only be used when the destinationRange is a public (non-RFC 1918) IP CIDR range.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "next_hop_instance_zone",
					Description: `(Optional when ` + "`" + `next_hop_instance` + "`" + ` is specified) The zone of the instance specified in ` + "`" + `next_hop_instance` + "`" + `. Omit if ` + "`" + `next_hop_instance` + "`" + ` is specified as a URL. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "next_hop_network",
					Description: `URL to a Network that should handle matching packets.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Route can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_route.default projects/{{project}}/global/routes/{{name}} $ terraform import google_compute_route.default {{project}}/{{name}} $ terraform import google_compute_route.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "next_hop_network",
					Description: `URL to a Network that should handle matching packets.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Route can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_route.default projects/{{project}}/global/routes/{{name}} $ terraform import google_compute_route.default {{project}}/{{name}} $ terraform import google_compute_route.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_router",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Represents a Router resource.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"router",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) A reference to the network to which this router belongs. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource.`,
				},
				resource.Attribute{
					Name:        "bgp",
					Description: `(Optional) BGP information specific to this router. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Region where the router resides.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `bgp` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "asn",
					Description: `(Required) Local BGP Autonomous System Number (ASN). Must be an RFC6996 private ASN, either 16-bit or 32-bit. The value will be fixed for this router resource. All VPN tunnels that link to this router will have the same local ASN.`,
				},
				resource.Attribute{
					Name:        "advertise_mode",
					Description: `(Optional) User-specified flag to indicate which mode to use for advertisement. Valid values of this enum field are: DEFAULT, CUSTOM`,
				},
				resource.Attribute{
					Name:        "advertised_groups",
					Description: `(Optional) User-specified list of prefix groups to advertise in custom mode. This field can only be populated if advertiseMode is CUSTOM and is advertised to all peers of the router. These groups will be advertised in addition to any specified prefixes. Leave this field blank to advertise no custom groups. This enum field has the one valid value: ALL_SUBNETS`,
				},
				resource.Attribute{
					Name:        "advertised_ip_ranges",
					Description: `(Optional) User-specified list of individual IP ranges to advertise in custom mode. This field can only be populated if advertiseMode is CUSTOM and is advertised to all peers of the router. These IP ranges will be advertised in addition to any specified groups. Leave this field blank to advertise no custom IP ranges. Structure is documented below. The ` + "`" + `advertised_ip_ranges` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "range",
					Description: `(Optional) The IP range to advertise. The value must be a CIDR-formatted string.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) User-specified description for the IP range. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Router can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_router.default projects/{{project}}/regions/{{region}}/routers/{{name}} $ terraform import google_compute_router.default {{project}}/{{region}}/{{name}} $ terraform import google_compute_router.default {{region}}/{{name}} $ terraform import google_compute_router.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Router can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_router.default projects/{{project}}/regions/{{region}}/routers/{{name}} $ terraform import google_compute_router.default {{project}}/{{region}}/{{name}} $ terraform import google_compute_router.default {{region}}/{{name}} $ terraform import google_compute_router.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_router_interface",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Manages a Cloud Router interface.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"router",
				"interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the interface, required by GCE. Changing this forces a new interface to be created.`,
				},
				resource.Attribute{
					Name:        "router",
					Description: `(Required) The name of the router this interface will be attached to. Changing this forces a new interface to be created. In addition to the above required fields, a router interface must have specified either ` + "`" + `ip_range` + "`" + ` or exactly one of ` + "`" + `vpn_tunnel` + "`" + ` or ` + "`" + `interconnect_attachment` + "`" + `, or both. - - -`,
				},
				resource.Attribute{
					Name:        "ip_range",
					Description: `(Optional) IP address and range of the interface. The IP range must be in the RFC3927 link-local IP space. Changing this forces a new interface to be created.`,
				},
				resource.Attribute{
					Name:        "vpn_tunnel",
					Description: `(Optional) The name or resource link to the VPN tunnel this interface will be linked to. Changing this forces a new interface to be created. Only one of ` + "`" + `vpn_tunnel` + "`" + ` and ` + "`" + `interconnect_attachment` + "`" + ` can be specified.`,
				},
				resource.Attribute{
					Name:        "interconnect_attachment",
					Description: `(Optional) The name or resource link to the VLAN interconnect for this interface. Changing this forces a new interface to be created. Only one of ` + "`" + `vpn_tunnel` + "`" + ` and ` + "`" + `interconnect_attachment` + "`" + ` can be specified.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which this interface's router belongs. If it is not provided, the provider project is used. Changing this forces a new interface to be created.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region this interface's router sits in. If not specified, the project region will be used. Changing this forces a new interface to be created. ## Attributes Reference Only the arguments listed above are exposed as attributes. ## Import Router interfaces can be imported using the ` + "`" + `region` + "`" + `, ` + "`" + `router` + "`" + `, and ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_router_interface.foobar us-central1/router-1/interface-1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_router_nat",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Manages a Cloud NAT.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"router",
				"nat",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for Cloud NAT, required by GCE. Changing this forces a new NAT to be created.`,
				},
				resource.Attribute{
					Name:        "router",
					Description: `(Required) The name of the router in which this NAT will be configured. Changing this forces a new NAT to be created.`,
				},
				resource.Attribute{
					Name:        "nat_ip_allocate_option",
					Description: `(Required) How external IPs should be allocated for this NAT. Valid values are ` + "`" + `AUTO_ONLY` + "`" + ` or ` + "`" + `MANUAL_ONLY` + "`" + `. Changing this forces a new NAT to be created.`,
				},
				resource.Attribute{
					Name:        "source_subnetwork_ip_ranges_to_nat",
					Description: `(Required) How NAT should be configured per Subnetwork. Valid values include: ` + "`" + `ALL_SUBNETWORKS_ALL_IP_RANGES` + "`" + `, ` + "`" + `ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES` + "`" + `, ` + "`" + `LIST_OF_SUBNETWORKS` + "`" + `. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which this NAT's router belongs. If it is not provided, the provider project is used. Changing this forces a new NAT to be created.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region this NAT's router sits in. If not specified, the project region will be used. Changing this forces a new NAT to be created.`,
				},
				resource.Attribute{
					Name:        "nat_ips",
					Description: `(Optional) List of ` + "`" + `self_link` + "`" + `s of external IPs. Only valid if ` + "`" + `nat_ip_allocate_option` + "`" + ` is set to ` + "`" + `MANUAL_ONLY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subnetwork",
					Description: `(Optional) One or more subnetwork NAT configurations. Only used if ` + "`" + `source_subnetwork_ip_ranges_to_nat` + "`" + ` is set to ` + "`" + `LIST_OF_SUBNETWORKS` + "`" + `. See the section below for details on configuration.`,
				},
				resource.Attribute{
					Name:        "min_ports_per_vm",
					Description: `(Optional) Minimum number of ports allocated to a VM from this NAT config. If not set, a default number of ports is allocated to a VM.`,
				},
				resource.Attribute{
					Name:        "udp_idle_timeout_sec",
					Description: `(Optional) Timeout (in seconds) for UDP connections. Defaults to 30s if not set.`,
				},
				resource.Attribute{
					Name:        "icmp_idle_timeout_sec",
					Description: `(Optional) Timeout (in seconds) for ICMP connections. Defaults to 30s if not set.`,
				},
				resource.Attribute{
					Name:        "tcp_established_idle_timeout_sec",
					Description: `(Optional) Timeout (in seconds) for TCP established connections. Defaults to 1200s if not set.`,
				},
				resource.Attribute{
					Name:        "tcp_transitory_idle_timeout_sec",
					Description: `(Optional) Timeout (in seconds) for TCP transitory connections. Defaults to 30s if not set. The ` + "`" + `subnetwork` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The ` + "`" + `self_link` + "`" + ` of the subnetwork to NAT.`,
				},
				resource.Attribute{
					Name:        "source_ip_ranges_to_nat",
					Description: `(Required) List of options for which source IPs in the subnetwork should have NAT enabled. Supported values include: ` + "`" + `ALL_IP_RANGES` + "`" + `, ` + "`" + `LIST_OF_SECONDARY_IP_RANGES` + "`" + `, ` + "`" + `PRIMARY_IP_RANGE` + "`" + ``,
				},
				resource.Attribute{
					Name:        "secondary_ip_range_names",
					Description: `(Optional) List of the secondary ranges of the subnetwork that are allowed to use NAT. This can be populated only if ` + "`" + `LIST_OF_SECONDARY_IP_RANGES` + "`" + ` is one of the values in ` + "`" + `source_ip_ranges_to_nat` + "`" + `. The ` + "`" + `log_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) Specifies the desired filtering of logs on this NAT. Valid values include: ` + "`" + `ALL` + "`" + `, ` + "`" + `ERRORS_ONLY` + "`" + `, ` + "`" + `TRANSLATIONS_ONLY` + "`" + ``,
				},
				resource.Attribute{
					Name:        "enable",
					Description: `(Required) Whether to export logs. ## Import Router NATs can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_router_nat.default {{project}}/{{region}}/{{router}}/{{name}} $ terraform import google_compute_router_nat.default {{region}}/{{router}}/{{name}} $ terraform import google_compute_router_nat.default {{router}}/{{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_router_peer",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Manages a Cloud Router BGP peer.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"router",
				"peer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for BGP peer, required by GCE. Changing this forces a new peer to be created.`,
				},
				resource.Attribute{
					Name:        "router",
					Description: `(Required) The name of the router in which this BGP peer will be configured. Changing this forces a new peer to be created.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Required) The name of the interface the BGP peer is associated with. Changing this forces a new peer to be created.`,
				},
				resource.Attribute{
					Name:        "peer_ip_address",
					Description: `(Required) IP address of the BGP interface outside Google Cloud. Changing this forces a new peer to be created.`,
				},
				resource.Attribute{
					Name:        "peer_asn",
					Description: `(Required) Peer BGP Autonomous System Number (ASN). Changing this forces a new peer to be created. - - -`,
				},
				resource.Attribute{
					Name:        "advertised_route_priority",
					Description: `(Optional) The priority of routes advertised to this BGP peer. Changing this forces a new peer to be created.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which this peer's router belongs. If it is not provided, the provider project is used. Changing this forces a new peer to be created.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region this peer's router sits in. If not specified, the project region will be used. Changing this forces a new peer to be created. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `IP address of the interface inside Google Cloud Platform. ## Import Router BGP peers can be imported using the ` + "`" + `region` + "`" + `, ` + "`" + `router` + "`" + `, and ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_router_peer.foobar us-central1/router-1/peer-1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_address",
					Description: `IP address of the interface inside Google Cloud Platform. ## Import Router BGP peers can be imported using the ` + "`" + `region` + "`" + `, ` + "`" + `router` + "`" + `, and ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_router_peer.foobar us-central1/router-1/peer-1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_security_policy",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Creates a Security Policy resource for Google Compute Engine.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"security",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the security policy. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this security policy. Max size is 2048.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Optional) The set of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "\`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Action to take when ` + "`" + `match` + "`" + ` matches the request. Valid values:`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Required) An unique positive integer indicating the priority of evaluation for a rule. Rules are evaluated from highest priority (lowest numerically) to lowest priority (highest numerically) in order.`,
				},
				resource.Attribute{
					Name:        "match",
					Description: `(Required) A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding ` + "`" + `action` + "`" + ` is enforced. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this rule. Max size is 64.`,
				},
				resource.Attribute{
					Name:        "preview",
					Description: `(Optional) When set to true, the ` + "`" + `action` + "`" + ` specified above is not enforced. Stackdriver logs for requests that trigger a preview action are annotated as such. The ` + "`" + `match` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `(Required) The configuration options available when specifying ` + "`" + `versioned_expr` + "`" + `. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "versioned_expr",
					Description: `(Required) Predefined rule expression. Available options:`,
				},
				resource.Attribute{
					Name:        "src_ip_ranges",
					Description: `(Required) Set of IP addresses or ranges (IPV4 or IPV6) in CIDR notation to match against inbound traffic. There is a limit of 5 IP ranges per rule. A value of '\`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `Fingerprint of this resource.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Import Security policies can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_security_policy.policy my-policy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "fingerprint",
					Description: `Fingerprint of this resource.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Import Security policies can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_security_policy.policy my-policy ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_shared_vpc_host_project",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Enables the Google Compute Engine Shared VPC feature for a project, assigning it as a host project.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"shared",
				"vpc",
				"host",
				"project",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Required) The ID of the project that will serve as a Shared VPC host project ## Import Google Compute Engine Shared VPC host project feature can be imported using the ` + "`" + `project` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_shared_vpc_host_project.host host-project-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_shared_vpc_service_project",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Enables the Google Compute Engine Shared VPC feature for a project, assigning it as a service project.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"shared",
				"vpc",
				"service",
				"project",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "host_project",
					Description: `(Required) The ID of a host project to associate.`,
				},
				resource.Attribute{
					Name:        "service_project",
					Description: `(Required) The ID of the project that will serve as a Shared VPC service project. ## Import Google Compute Engine Shared VPC service project feature can be imported using the ` + "`" + `host_project` + "`" + ` and ` + "`" + `service_project` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_shared_vpc_service_project.service1 host-project-id/service-project-id-1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_snapshot",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Represents a Persistent Disk Snapshot resource.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"snapshot",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "source_disk",
					Description: `(Required) A reference to the disk used to create this snapshot. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Labels to apply to this Snapshot.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) A reference to the zone where the disk is hosted.`,
				},
				resource.Attribute{
					Name:        "snapshot_encryption_key",
					Description: `(Optional) The customer-supplied encryption key of the snapshot. Required if the source snapshot is protected by a customer-supplied encryption key. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "source_disk_encryption_key",
					Description: `(Optional) The customer-supplied encryption key of the source snapshot. Required if the source snapshot is protected by a customer-supplied encryption key. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `snapshot_encryption_key` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "raw_key",
					Description: `(Optional) Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource.`,
				},
				resource.Attribute{
					Name:        "sha256",
					Description: `The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied encryption key that protects this resource. The ` + "`" + `source_disk_encryption_key` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "raw_key",
					Description: `(Optional) Specifies a 256-bit customer-supplied encryption key, encoded in RFC 4648 base64 to either encrypt or decrypt this resource. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The unique identifier for the resource.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `Size of the snapshot, specified in GB.`,
				},
				resource.Attribute{
					Name:        "storage_bytes",
					Description: `A size of the the storage used by the snapshot. As snapshots share storage, this number is expected to change with snapshot creation/deletion.`,
				},
				resource.Attribute{
					Name:        "licenses",
					Description: `A list of public visible licenses that apply to this snapshot. This can be because the original image had licenses attached (such as a Windows image). snapshotEncryptionKey nested object Encrypts the snapshot using a customer-supplied encryption key.`,
				},
				resource.Attribute{
					Name:        "label_fingerprint",
					Description: `The fingerprint used for optimistic locking of this resource. Used internally during updates.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 5 minutes. - ` + "`" + `update` + "`" + ` - Default is 5 minutes. - ` + "`" + `delete` + "`" + ` - Default is 5 minutes. ## Import Snapshot can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_snapshot.default projects/{{project}}/global/snapshots/{{name}} $ terraform import google_compute_snapshot.default {{project}}/{{name}} $ terraform import google_compute_snapshot.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `The unique identifier for the resource.`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `Size of the snapshot, specified in GB.`,
				},
				resource.Attribute{
					Name:        "storage_bytes",
					Description: `A size of the the storage used by the snapshot. As snapshots share storage, this number is expected to change with snapshot creation/deletion.`,
				},
				resource.Attribute{
					Name:        "licenses",
					Description: `A list of public visible licenses that apply to this snapshot. This can be because the original image had licenses attached (such as a Windows image). snapshotEncryptionKey nested object Encrypts the snapshot using a customer-supplied encryption key.`,
				},
				resource.Attribute{
					Name:        "label_fingerprint",
					Description: `The fingerprint used for optimistic locking of this resource. Used internally during updates.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 5 minutes. - ` + "`" + `update` + "`" + ` - Default is 5 minutes. - ` + "`" + `delete` + "`" + ` - Default is 5 minutes. ## Import Snapshot can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_snapshot.default projects/{{project}}/global/snapshots/{{name}} $ terraform import google_compute_snapshot.default {{project}}/{{name}} $ terraform import google_compute_snapshot.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_ssl_certificate",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `An SslCertificate resource, used for HTTPS load balancing.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"ssl",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "certificate",
					Description: `(Required) The certificate in PEM format. The certificate chain must be no greater than 5 certs long. The chain must include at least one intermediate cert.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required) The write-only private key in PEM format. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "name_prefix",
					Description: `(Optional) Creates a unique name beginning with the specified prefix. Conflicts with ` + "`" + `name` + "`" + `. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `The unique identifier for the resource.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import SslCertificate can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_ssl_certificate.default projects/{{project}}/global/sslCertificates/{{name}} $ terraform import google_compute_ssl_certificate.default {{project}}/{{name}} $ terraform import google_compute_ssl_certificate.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `The unique identifier for the resource.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import SslCertificate can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_ssl_certificate.default projects/{{project}}/global/sslCertificates/{{name}} $ terraform import google_compute_ssl_certificate.default {{project}}/{{name}} $ terraform import google_compute_ssl_certificate.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_ssl_policy",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Represents a SSL policy.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"ssl",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource.`,
				},
				resource.Attribute{
					Name:        "profile",
					Description: `(Optional) Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. This can be one of ` + "`" + `COMPATIBLE` + "`" + `, ` + "`" + `MODERN` + "`" + `, ` + "`" + `RESTRICTED` + "`" + `, or ` + "`" + `CUSTOM` + "`" + `. If using ` + "`" + `CUSTOM` + "`" + `, the set of SSL features to enable must be specified in the ` + "`" + `customFeatures` + "`" + ` field. See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport) for information on what cipher suites each profile provides. If ` + "`" + `CUSTOM` + "`" + ` is used, the ` + "`" + `custom_features` + "`" + ` attribute`,
				},
				resource.Attribute{
					Name:        "min_tls_version",
					Description: `(Optional) The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer. This can be one of ` + "`" + `TLS_1_0` + "`" + `, ` + "`" + `TLS_1_1` + "`" + `, ` + "`" + `TLS_1_2` + "`" + `. Default is ` + "`" + `TLS_1_0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "custom_features",
					Description: `(Optional) Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. This can be one of ` + "`" + `COMPATIBLE` + "`" + `, ` + "`" + `MODERN` + "`" + `, ` + "`" + `RESTRICTED` + "`" + `, or ` + "`" + `CUSTOM` + "`" + `. If using ` + "`" + `CUSTOM` + "`" + `, the set of SSL features to enable must be specified in the ` + "`" + `customFeatures` + "`" + ` field. See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport) for which ciphers are available to use.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "enabled_features",
					Description: `The list of features enabled in the SSL policy.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import SslPolicy can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_ssl_policy.default projects/{{project}}/global/sslPolicies/{{name}} $ terraform import google_compute_ssl_policy.default {{project}}/{{name}} $ terraform import google_compute_ssl_policy.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "enabled_features",
					Description: `The list of features enabled in the SSL policy.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import SslPolicy can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_ssl_policy.default projects/{{project}}/global/sslPolicies/{{name}} $ terraform import google_compute_ssl_policy.default {{project}}/{{name}} $ terraform import google_compute_ssl_policy.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_subnetwork",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `A VPC network is a virtual version of the traditional physical networks that exist within and between physical data centers.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"subnetwork",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_cidr_range",
					Description: `(Required) The range of internal addresses that are owned by this subnetwork. Provide this property when you create the subnetwork. For example, 10.0.0.0/8 or 192.168.0.0/16. Ranges must be unique and non-overlapping within a network. Only IPv4 is supported.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the resource, provided by the client when initially creating the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) The network this subnet belongs to. Only networks that are in the distributed mode can have subnetworks. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource. Provide this property when you create the resource. This field can be set only at resource creation time.`,
				},
				resource.Attribute{
					Name:        "enable_flow_logs",
					Description: `(Optional) Whether to enable flow logging for this subnetwork.`,
				},
				resource.Attribute{
					Name:        "secondary_ip_range",
					Description: `(Optional) An array of configurations for secondary IP ranges for VM instances contained in this subnetwork. The primary IP of such VM must belong to the primary ipCidrRange of the subnetwork. The alias IPs may belong to either primary or secondary ranges. This field uses attr-as-block mode to avoid breaking users during the 0.12 upgrade. See [the Attr-as-Block page](https://www.terraform.io/docs/configuration/attr-as-blocks.html) for more details. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "private_ip_google_access",
					Description: `(Optional) When enabled, VMs in this subnetwork without external IP addresses can access Google APIs and services by using Private Google Access.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) URL of the GCP region for this subnetwork.`,
				},
				resource.Attribute{
					Name:        "log_config",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) Denotes the logging options for the subnetwork flow logs. If logging is enabled logs will be exported to Stackdriver. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `secondary_ip_range` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "range_name",
					Description: `(Required) The name associated with this subnetwork secondary range, used when adding an alias IP range to a VM instance. The name must be 1-63 characters long, and comply with RFC1035. The name must be unique within the subnetwork.`,
				},
				resource.Attribute{
					Name:        "ip_cidr_range",
					Description: `(Required) The range of IP addresses belonging to this subnetwork secondary range. Provide this property when you create the subnetwork. Ranges must be unique and non-overlapping with all primary and secondary IP ranges within a network. Only IPv4 is supported. The ` + "`" + `log_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "aggregation_interval",
					Description: `(Optional) Can only be specified if VPC flow logging for this subnetwork is enabled. Toggles the aggregation interval for collecting flow logs. Increasing the interval time will reduce the amount of generated flow logs for long lasting connections. Default is an interval of 5 seconds per connection. Possible values are INTERVAL_5_SEC, INTERVAL_30_SEC, INTERVAL_1_MIN, INTERVAL_5_MIN, INTERVAL_10_MIN, INTERVAL_15_MIN`,
				},
				resource.Attribute{
					Name:        "flow_sampling",
					Description: `(Optional) Can only be specified if VPC flow logging for this subnetwork is enabled. The value of the field must be in [0, 1]. Set the sampling rate of VPC flow logs within the subnetwork where 1.0 means all collected logs are reported and 0.0 means no logs are reported. Default is 0.5 which means half of all collected logs are reported.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Can only be specified if VPC flow logging for this subnetwork is enabled. Configures whether metadata fields should be added to the reported VPC flow logs. Default is ` + "`" + `INCLUDE_ALL_METADATA` + "`" + `. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "gateway_address",
					Description: `The gateway address for default routes to reach destination addresses outside this subnetwork.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `Fingerprint of this resource. This field is used internally during updates of this resource.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 6 minutes. - ` + "`" + `update` + "`" + ` - Default is 6 minutes. - ` + "`" + `delete` + "`" + ` - Default is 6 minutes. ## Import Subnetwork can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_subnetwork.default projects/{{project}}/regions/{{region}}/subnetworks/{{name}} $ terraform import google_compute_subnetwork.default {{project}}/{{region}}/{{name}} $ terraform import google_compute_subnetwork.default {{region}}/{{name}} $ terraform import google_compute_subnetwork.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "gateway_address",
					Description: `The gateway address for default routes to reach destination addresses outside this subnetwork.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `Fingerprint of this resource. This field is used internally during updates of this resource.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 6 minutes. - ` + "`" + `update` + "`" + ` - Default is 6 minutes. - ` + "`" + `delete` + "`" + ` - Default is 6 minutes. ## Import Subnetwork can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_subnetwork.default projects/{{project}}/regions/{{region}}/subnetworks/{{name}} $ terraform import google_compute_subnetwork.default {{project}}/{{region}}/{{name}} $ terraform import google_compute_subnetwork.default {{region}}/{{name}} $ terraform import google_compute_subnetwork.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_subnetwork_iam_policy",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Collection of resources to manage IAM policy for a GCE subnetwork.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"subnetwork",
				"iam",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "subnetwork",
					Description: `(Required) The name of the subnetwork.`,
				},
				resource.Attribute{
					Name:        "member/members",
					Description: `(Required) Identities that will be granted the privilege in ` + "`" + `role` + "`" + `. Each entry can have one of the following values:`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. Only one ` + "`" + `google_compute_subnetwork_iam_binding` + "`" + ` can be used per role. Note that custom roles must be of the format ` + "`" + `[projects|organizations]/{parent-name}/roles/{role-name}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy_data",
					Description: `(Required only by ` + "`" + `google_compute_subnetwork_iam_policy` + "`" + `) The policy data generated by a ` + "`" + `google_iam_policy` + "`" + ` data source.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region of the subnetwork. If unspecified, this defaults to the region configured in the provider. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the subnetwork's IAM policy. ## Import For all import syntaxes, the "resource in question" can take any of the following forms:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the subnetwork's IAM policy. ## Import For all import syntaxes, the "resource in question" can take any of the following forms:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_target_http_proxy",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Represents a TargetHttpProxy resource, which is used by one or more global forwarding rule to route incoming HTTP requests to a URL map.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"target",
				"http",
				"proxy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "url_map",
					Description: `(Required) A reference to the UrlMap resource that defines the mapping from URL to the BackendService. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "proxy_id",
					Description: `The unique identifier for the resource.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import TargetHttpProxy can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_target_http_proxy.default projects/{{project}}/global/targetHttpProxies/{{name}} $ terraform import google_compute_target_http_proxy.default {{project}}/{{name}} $ terraform import google_compute_target_http_proxy.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "proxy_id",
					Description: `The unique identifier for the resource.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import TargetHttpProxy can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_target_http_proxy.default projects/{{project}}/global/targetHttpProxies/{{name}} $ terraform import google_compute_target_http_proxy.default {{project}}/{{name}} $ terraform import google_compute_target_http_proxy.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_target_https_proxy",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Represents a TargetHttpsProxy resource, which is used by one or more global forwarding rule to route incoming HTTPS requests to a URL map.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"target",
				"https",
				"proxy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "ssl_certificates",
					Description: `(Required) A list of SslCertificate resources that are used to authenticate connections between users and the load balancer. Currently, exactly one SSL certificate must be specified.`,
				},
				resource.Attribute{
					Name:        "url_map",
					Description: `(Required) A reference to the UrlMap resource that defines the mapping from URL to the BackendService. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource.`,
				},
				resource.Attribute{
					Name:        "quic_override",
					Description: `(Optional) Specifies the QUIC override policy for this resource. This determines whether the load balancer will attempt to negotiate QUIC with clients or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is specified, uses the QUIC policy with no user overrides, which is equivalent to DISABLE. Not specifying this field is equivalent to specifying NONE.`,
				},
				resource.Attribute{
					Name:        "ssl_policy",
					Description: `(Optional) A reference to the SslPolicy resource that will be associated with the TargetHttpsProxy resource. If not set, the TargetHttpsProxy resource will not have any SSL policy configured.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "proxy_id",
					Description: `The unique identifier for the resource.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import TargetHttpsProxy can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_target_https_proxy.default projects/{{project}}/global/targetHttpsProxies/{{name}} $ terraform import google_compute_target_https_proxy.default {{project}}/{{name}} $ terraform import google_compute_target_https_proxy.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "proxy_id",
					Description: `The unique identifier for the resource.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import TargetHttpsProxy can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_target_https_proxy.default projects/{{project}}/global/targetHttpsProxies/{{name}} $ terraform import google_compute_target_https_proxy.default {{project}}/{{name}} $ terraform import google_compute_target_https_proxy.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_target_instance",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Represents a TargetInstance resource which defines an endpoint instance that terminates traffic of certain protocols.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"target",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "instance",
					Description: `(Required) The Compute instance VM handling traffic for this target instance. Accepts the instance self-link, relative path (e.g. ` + "`" + `projects/project/zones/zone/instances/instance` + "`" + `) or name. If name is given, the zone will default to the given zone or the provider-default zone and the project will default to the provider-level project. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource.`,
				},
				resource.Attribute{
					Name:        "nat_policy",
					Description: `(Optional) NAT option controlling how IPs are NAT'ed to the instance. Currently only NO_NAT (default value) is supported.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) URL of the zone where the target instance resides.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import TargetInstance can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_target_instance.default projects/{{project}}/zones/{{zone}}/targetInstances/{{name}} $ terraform import google_compute_target_instance.default {{project}}/{{zone}}/{{name}} $ terraform import google_compute_target_instance.default {{zone}}/{{name}} $ terraform import google_compute_target_instance.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import TargetInstance can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_target_instance.default projects/{{project}}/zones/{{zone}}/targetInstances/{{name}} $ terraform import google_compute_target_instance.default {{project}}/{{zone}}/{{name}} $ terraform import google_compute_target_instance.default {{zone}}/{{name}} $ terraform import google_compute_target_instance.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_target_pool",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Manages a Target Pool within GCE.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"target",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the resource, required by GCE. Changing this forces a new resource to be created. - - -`,
				},
				resource.Attribute{
					Name:        "backup_pool",
					Description: `(Optional) URL to the backup target pool. Must also set failover\_ratio.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Textual description field.`,
				},
				resource.Attribute{
					Name:        "failover_ratio",
					Description: `(Optional) Ratio (0 to 1) of failed nodes before using the backup pool (which must also be set).`,
				},
				resource.Attribute{
					Name:        "health_checks",
					Description: `(Optional) List of zero or one health check name or self_link. Only legacy ` + "`" + `google_compute_http_health_check` + "`" + ` is supported.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `(Optional) List of instances in the pool. They can be given as URLs, or in the form of "zone/name". Note that the instances need not exist at the time of target pool creation, so there is no need to use the Terraform interpolators to create a dependency on the instances from the target pool.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Where the target pool resides. Defaults to project region.`,
				},
				resource.Attribute{
					Name:        "session_affinity",
					Description: `(Optional) How to distribute load. Options are "NONE" (no affinity). "CLIENT\_IP" (hash of the source/dest addresses / ports), and "CLIENT\_IP\_PROTO" also includes the protocol (default "NONE"). ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Import Target pools can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_target_pool.default instance-pool ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Import Target pools can be imported using the ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_target_pool.default instance-pool ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_target_ssl_proxy",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Represents a TargetSslProxy resource, which is used by one or more global forwarding rule to route incoming SSL requests to a backend service.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"target",
				"ssl",
				"proxy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "backend_service",
					Description: `(Required) A reference to the BackendService resource.`,
				},
				resource.Attribute{
					Name:        "ssl_certificates",
					Description: `(Required) A list of SslCertificate resources that are used to authenticate connections between users and the load balancer. Currently, exactly one SSL certificate must be specified. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource.`,
				},
				resource.Attribute{
					Name:        "proxy_header",
					Description: `(Optional) Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is NONE.`,
				},
				resource.Attribute{
					Name:        "ssl_policy",
					Description: `(Optional) A reference to the SslPolicy resource that will be associated with the TargetSslProxy resource. If not set, the TargetSslProxy resource will not have any SSL policy configured.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "proxy_id",
					Description: `The unique identifier for the resource.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import TargetSslProxy can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_target_ssl_proxy.default projects/{{project}}/global/targetSslProxies/{{name}} $ terraform import google_compute_target_ssl_proxy.default {{project}}/{{name}} $ terraform import google_compute_target_ssl_proxy.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "proxy_id",
					Description: `The unique identifier for the resource.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import TargetSslProxy can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_target_ssl_proxy.default projects/{{project}}/global/targetSslProxies/{{name}} $ terraform import google_compute_target_ssl_proxy.default {{project}}/{{name}} $ terraform import google_compute_target_ssl_proxy.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_target_tcp_proxy",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Represents a TargetTcpProxy resource, which is used by one or more global forwarding rule to route incoming TCP requests to a Backend service.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"target",
				"tcp",
				"proxy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "backend_service",
					Description: `(Required) A reference to the BackendService resource. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource.`,
				},
				resource.Attribute{
					Name:        "proxy_header",
					Description: `(Optional) Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is NONE.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "proxy_id",
					Description: `The unique identifier for the resource.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import TargetTcpProxy can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_target_tcp_proxy.default projects/{{project}}/global/targetTcpProxies/{{name}} $ terraform import google_compute_target_tcp_proxy.default {{project}}/{{name}} $ terraform import google_compute_target_tcp_proxy.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "proxy_id",
					Description: `The unique identifier for the resource.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import TargetTcpProxy can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_target_tcp_proxy.default projects/{{project}}/global/targetTcpProxies/{{name}} $ terraform import google_compute_target_tcp_proxy.default {{project}}/{{name}} $ terraform import google_compute_target_tcp_proxy.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_url_map",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `UrlMaps are used to route requests to a backend service based on rules that you define for the host and path of an incoming URL.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"url",
				"map",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "default_service",
					Description: `(Required) The backend service or backend bucket to use when none of the given rules match.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource. Provide this property when you create the resource.`,
				},
				resource.Attribute{
					Name:        "host_rule",
					Description: `(Optional) The list of HostRules to use against the URL. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "path_matcher",
					Description: `(Optional) The list of named PathMatchers to use against the URL. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "test",
					Description: `(Optional) The list of expected URL mappings. Requests to update this UrlMap will succeed only if all of the test cases pass. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `host_rule` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this HostRule. Provide this property when you create the resource.`,
				},
				resource.Attribute{
					Name:        "hosts",
					Description: `(Required) The list of host patterns to match. They must be valid hostnames, except`,
				},
				resource.Attribute{
					Name:        "path_matcher",
					Description: `(Required) The name of the PathMatcher to use to match the path portion of the URL if the hostRule matches the URL's host portion. The ` + "`" + `path_matcher` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "default_service",
					Description: `(Required) The backend service or backend bucket to use when none of the given paths match.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name to which this PathMatcher is referred by the HostRule.`,
				},
				resource.Attribute{
					Name:        "path_rule",
					Description: `(Optional) The list of path rules. Structure is documented below. The ` + "`" + `path_rule` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "paths",
					Description: `(Required) The list of path patterns to match. Each must start with / and the only place a`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Required) The backend service or backend bucket to use if any of the given paths match. The ` + "`" + `test` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this test case.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required) Host portion of the URL.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) Path portion of the URL.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Required) The backend service or backend bucket link that should be matched by this test. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "map_id",
					Description: `The unique identifier for the resource.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `Fingerprint of this resource. This field is used internally during updates of this resource.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import UrlMap can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_url_map.default projects/{{project}}/global/urlMaps/{{name}} $ terraform import google_compute_url_map.default {{project}}/{{name}} $ terraform import google_compute_url_map.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "map_id",
					Description: `The unique identifier for the resource.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `Fingerprint of this resource. This field is used internally during updates of this resource.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import UrlMap can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_url_map.default projects/{{project}}/global/urlMaps/{{name}} $ terraform import google_compute_url_map.default {{project}}/{{name}} $ terraform import google_compute_url_map.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_vpn_gateway",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `Represents a VPN gateway running in GCP.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"vpn",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) The network this VPN gateway is accepting traffic for. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region this gateway should sit in.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import VpnGateway can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_vpn_gateway.default projects/{{project}}/regions/{{region}}/targetVpnGateways/{{name}} $ terraform import google_compute_vpn_gateway.default {{project}}/{{region}}/{{name}} $ terraform import google_compute_vpn_gateway.default {{region}}/{{name}} $ terraform import google_compute_vpn_gateway.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import VpnGateway can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_vpn_gateway.default projects/{{project}}/regions/{{region}}/targetVpnGateways/{{name}} $ terraform import google_compute_vpn_gateway.default {{project}}/{{region}}/{{name}} $ terraform import google_compute_vpn_gateway.default {{region}}/{{name}} $ terraform import google_compute_vpn_gateway.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_compute_vpn_tunnel",
			Category:         "Google Compute Engine Resources",
			ShortDescription: `VPN tunnel resource.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"vpn",
				"tunnel",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "shared_secret",
					Description: `(Required) Shared secret used to set the secure session between the Cloud VPN gateway and the peer VPN gateway. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description of this resource.`,
				},
				resource.Attribute{
					Name:        "target_vpn_gateway",
					Description: `(Optional) URL of the Target VPN gateway with which this VPN tunnel is associated.`,
				},
				resource.Attribute{
					Name:        "vpn_gateway",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) URL of the VPN gateway with which this VPN tunnel is associated. This must be used if a High Availability VPN gateway resource is created. This field must reference a ` + "`" + `google_compute_ha_vpn_gateway` + "`" + ` resource.`,
				},
				resource.Attribute{
					Name:        "vpn_gateway_interface",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) The interface ID of the VPN gateway with which this VPN tunnel is associated.`,
				},
				resource.Attribute{
					Name:        "peer_external_gateway",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) URL of the peer side external VPN gateway to which this VPN tunnel is connected.`,
				},
				resource.Attribute{
					Name:        "peer_external_gateway_interface",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) The interface ID of the external VPN gateway to which this VPN tunnel is connected.`,
				},
				resource.Attribute{
					Name:        "peer_gcp_gateway",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected. If provided, the VPN tunnel will automatically use the same vpn_gateway_interface ID in the peer GCP VPN gateway. This field must reference a ` + "`" + `google_compute_ha_vpn_gateway` + "`" + ` resource.`,
				},
				resource.Attribute{
					Name:        "router",
					Description: `(Optional) URL of router resource to be used for dynamic routing.`,
				},
				resource.Attribute{
					Name:        "peer_ip",
					Description: `(Optional) IP address of the peer VPN gateway. Only IPv4 is supported.`,
				},
				resource.Attribute{
					Name:        "ike_version",
					Description: `(Optional) IKE protocol version to use when establishing the VPN tunnel with peer VPN gateway. Acceptable IKE versions are 1 or 2. Default version is 2.`,
				},
				resource.Attribute{
					Name:        "local_traffic_selector",
					Description: `(Optional) Local traffic selector to use when establishing the VPN tunnel with peer VPN gateway. The value should be a CIDR formatted string, for example ` + "`" + `192.168.0.0/16` + "`" + `. The ranges should be disjoint. Only IPv4 is supported.`,
				},
				resource.Attribute{
					Name:        "remote_traffic_selector",
					Description: `(Optional) Remote traffic selector to use when establishing the VPN tunnel with peer VPN gateway. The value should be a CIDR formatted string, for example ` + "`" + `192.168.0.0/16` + "`" + `. The ranges should be disjoint. Only IPv4 is supported.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) Labels to apply to this VpnTunnel.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region where the tunnel is located. If unset, is set to the region of ` + "`" + `target_vpn_gateway` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "shared_secret_hash",
					Description: `Hash of the shared secret.`,
				},
				resource.Attribute{
					Name:        "label_fingerprint",
					Description: `The fingerprint used for optimistic locking of this resource. Used internally during updates.`,
				},
				resource.Attribute{
					Name:        "detailed_status",
					Description: `Detailed status message for the VPN tunnel.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import VpnTunnel can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_vpn_tunnel.default projects/{{project}}/regions/{{region}}/vpnTunnels/{{name}} $ terraform import google_compute_vpn_tunnel.default {{project}}/{{region}}/{{name}} $ terraform import google_compute_vpn_tunnel.default {{region}}/{{name}} $ terraform import google_compute_vpn_tunnel.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_timestamp",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "shared_secret_hash",
					Description: `Hash of the shared secret.`,
				},
				resource.Attribute{
					Name:        "label_fingerprint",
					Description: `The fingerprint used for optimistic locking of this resource. Used internally during updates.`,
				},
				resource.Attribute{
					Name:        "detailed_status",
					Description: `Detailed status message for the VPN tunnel.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import VpnTunnel can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_compute_vpn_tunnel.default projects/{{project}}/regions/{{region}}/vpnTunnels/{{name}} $ terraform import google_compute_vpn_tunnel.default {{project}}/{{region}}/{{name}} $ terraform import google_compute_vpn_tunnel.default {{region}}/{{name}} $ terraform import google_compute_vpn_tunnel.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_container_analysis_note",
			Category:         "Google Container Analysis Resources",
			ShortDescription: `Provides a detailed description of a Note.`,
			Description:      ``,
			Keywords: []string{
				"container",
				"analysis",
				"note",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the note.`,
				},
				resource.Attribute{
					Name:        "attestation_authority",
					Description: `(Required) Note kind that represents a logical attestation "role" or "authority". For example, an organization might have one AttestationAuthority for "QA" and one for "build". This Note is intended to act strictly as a grouping mechanism for the attached Occurrences (Attestations). This grouping mechanism also provides a security boundary, since IAM ACLs gate the ability for a principle to attach an Occurrence to a given Note. It also provides a single point of lookup to find all attached Attestation Occurrences, even if they don't all live in the same project. Structure is documented below. The ` + "`" + `attestation_authority` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "hint",
					Description: `(Required) This submessage provides human-readable hints about the purpose of the AttestationAuthority. Because the name of a Note acts as its resource reference, it is important to disambiguate the canonical name of the Note (which might be a UUID for security purposes) from "readable" names more suitable for debug output. Note that these hints should NOT be used to look up AttestationAuthorities in security sensitive contexts, such as when looking up Attestations to verify. Structure is documented below. The ` + "`" + `hint` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "human_readable_name",
					Description: `(Required) The human readable name of this Attestation Authority, for example "qa". - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Note can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import -provider=google-beta google_container_analysis_note.default projects/{{project}}/notes/{{name}} $ terraform import -provider=google-beta google_container_analysis_note.default {{project}}/{{name}} $ terraform import -provider=google-beta google_container_analysis_note.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_container_cluster",
			Category:         "Google Kubernetes (Container) Engine Resources",
			ShortDescription: `Creates a Google Kubernetes Engine (GKE) cluster.`,
			Description:      ``,
			Keywords: []string{
				"kubernetes",
				"container",
				"engine",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the cluster, unique within the project and location. - - -`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Optional) The location (region or zone) in which the cluster master will be created, as well as the default node location. If you specify a zone (such as ` + "`" + `us-central1-a` + "`" + `), the cluster will be a zonal cluster with a single cluster master. If you specify a region (such as ` + "`" + `us-west1` + "`" + `), the cluster will be a regional cluster with multiple masters spread across zones in the region, and with default node locations in those zones as well.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional, Deprecated) The zone that the cluster master and nodes should be created in. If specified, this cluster will be a zonal cluster. ` + "`" + `zone` + "`" + ` has been deprecated in favour of ` + "`" + `location` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "node_locations",
					Description: `(Optional) The list of zones in which the cluster's nodes should be located. These must be in the same region as the cluster zone for zonal clusters, or in the region of a regional cluster. In a multi-zonal cluster, the number of nodes specified in ` + "`" + `initial_node_count` + "`" + ` is created in all specified zones as well as the primary zone. If specified for a regional cluster, nodes will be created in only these zones. -> A "multi-zonal" cluster is a zonal cluster with at least one additional zone defined; in a multi-zonal cluster, the cluster master is only present in a single zone while nodes are present in each of the primary zone and the node locations. In contrast, in a regional cluster, cluster master nodes are present in multiple zones in the region. For that reason, regional clusters should be preferred.`,
				},
				resource.Attribute{
					Name:        "additional_zones",
					Description: `(Optional) The list of zones in which the cluster's nodes should be located. These must be in the same region as the cluster zone for zonal clusters, or in the region of a regional cluster. In a multi-zonal cluster, the number of nodes specified in ` + "`" + `initial_node_count` + "`" + ` is created in all specified zones as well as the primary zone. If specified for a regional cluster, nodes will only be created in these zones. ` + "`" + `additional_zones` + "`" + ` has been deprecated in favour of ` + "`" + `node_locations` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "addons_config",
					Description: `(Optional) The configuration for addons supported by GKE. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "cluster_ipv4_cidr",
					Description: `(Optional) The IP address range of the Kubernetes pods in this cluster in CIDR notation (e.g. 10.96.0.0/14). Leave blank to have one automatically chosen or specify a /14 block in 10.0.0.0/8. This field will only work if your cluster is not VPC-native- when an ` + "`" + `ip_allocation_policy` + "`" + ` block is not defined, or ` + "`" + `ip_allocation_policy.use_ip_aliases` + "`" + ` is set to false. If your cluster is VPC-native, use ` + "`" + `ip_allocation_policy.cluster_ipv4_cidr_block` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cluster_autoscaling",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) Per-cluster configuration of Node Auto-Provisioning with Cluster Autoscaler to automatically adjust the size of the cluster and create/delete node pools based on the current needs of the cluster's workload. See the [guide to using Node Auto-Provisioning](https://cloud.google.com/kubernetes-engine/docs/how-to/node-auto-provisioning) for more details. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "database_encryption",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)). Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the cluster.`,
				},
				resource.Attribute{
					Name:        "default_max_pods_per_node",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) The default maximum number of pods per node in this cluster. Note that this does not work on node pools which are "route-based" - that is, node pools belonging to clusters that do not have IP Aliasing enabled. See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr) for more information.`,
				},
				resource.Attribute{
					Name:        "enable_binary_authorization",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) Enable Binary Authorization for this cluster. If enabled, all container images will be validated by Google Binary Authorization.`,
				},
				resource.Attribute{
					Name:        "enable_kubernetes_alpha",
					Description: `(Optional) Whether to enable Kubernetes Alpha features for this cluster. Note that when this option is enabled, the cluster cannot be upgraded and will be automatically deleted after 30 days.`,
				},
				resource.Attribute{
					Name:        "enable_tpu",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) Whether to enable Cloud TPU resources in this cluster. See the [official documentation](https://cloud.google.com/tpu/docs/kubernetes-engine-setup).`,
				},
				resource.Attribute{
					Name:        "enable_legacy_abac",
					Description: `(Optional) Whether the ABAC authorizer is enabled for this cluster. When enabled, identities in the system, including service accounts, nodes, and controllers, will have statically granted permissions beyond those provided by the RBAC configuration or IAM. Defaults to ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "initial_node_count",
					Description: `(Optional) The number of nodes to create in this cluster's default node pool. Must be set if ` + "`" + `node_pool` + "`" + ` is not set. If you're using ` + "`" + `google_container_node_pool` + "`" + ` objects with no default node pool, you'll need to set this to a value of at least ` + "`" + `1` + "`" + `, alongside setting ` + "`" + `remove_default_node_pool` + "`" + ` to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip_allocation_policy",
					Description: `(Optional) Configuration for cluster IP allocation. As of now, only pre-allocated subnetworks (custom type with secondary ranges) are supported. This will activate IP aliases. See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/ip-aliases) Structure is documented below. This field is marked to use [Attribute as Block](/docs/configuration/attr-as-blocks.html) in order to support explicit removal with ` + "`" + `ip_allocation_policy = []` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logging_service",
					Description: `(Optional) The logging service that the cluster should write logs to. Available options include ` + "`" + `logging.googleapis.com` + "`" + `, ` + "`" + `logging.googleapis.com/kubernetes` + "`" + `, and ` + "`" + `none` + "`" + `. Defaults to ` + "`" + `logging.googleapis.com` + "`" + ``,
				},
				resource.Attribute{
					Name:        "maintenance_policy",
					Description: `(Optional) The maintenance policy to use for the cluster. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "master_auth",
					Description: `(Optional) The authentication information for accessing the Kubernetes master. Some values in this block are only returned by the API if your service account has permission to get credentials for your GKE cluster. If you see an unexpected diff removing a username/password or unsetting your client cert, ensure you have the ` + "`" + `container.clusters.getCredentials` + "`" + ` permission. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "master_authorized_networks_config",
					Description: `(Optional) The desired configuration options for master authorized networks. Omit the nested ` + "`" + `cidr_blocks` + "`" + ` attribute to disallow external access (except the cluster node IPs, which GKE automatically whitelists).`,
				},
				resource.Attribute{
					Name:        "min_master_version",
					Description: `(Optional) The minimum version of the master. GKE will auto-update the master to new versions, so this does not guarantee the current master version--use the read-only ` + "`" + `master_version` + "`" + ` field to obtain that. If unset, the cluster's version will be set by GKE to the version of the most recent official release (which is not necessarily the latest version). Most users will find the ` + "`" + `google_container_engine_versions` + "`" + ` data source useful - it indicates which versions are available, and can be use to approximate fuzzy versions in a Terraform-compatible way. If you intend to specify versions manually, [the docs](https://cloud.google.com/kubernetes-engine/versioning-and-upgrades#specifying_cluster_version) describe the various acceptable formats for this field. -> If you are using the ` + "`" + `google_container_engine_versions` + "`" + ` datasource with a regional cluster, ensure that you have provided a ` + "`" + `region` + "`" + ` to the datasource. A ` + "`" + `region` + "`" + ` can have a different set of supported versions than its corresponding ` + "`" + `zone` + "`" + `s, and not all ` + "`" + `zone` + "`" + `s in a ` + "`" + `region` + "`" + ` are guaranteed to support the same version.`,
				},
				resource.Attribute{
					Name:        "monitoring_service",
					Description: `(Optional) The monitoring service that the cluster should write metrics to. Automatically send metrics from pods in the cluster to the Google Cloud Monitoring API. VM metrics will be collected by Google Compute Engine regardless of this setting Available options include ` + "`" + `monitoring.googleapis.com` + "`" + `, ` + "`" + `monitoring.googleapis.com/kubernetes` + "`" + `, and ` + "`" + `none` + "`" + `. Defaults to ` + "`" + `monitoring.googleapis.com` + "`" + ``,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) The name or self_link of the Google Compute Engine network to which the cluster is connected. For Shared VPC, set this to the self link of the shared network.`,
				},
				resource.Attribute{
					Name:        "network_policy",
					Description: `(Optional) Configuration options for the [NetworkPolicy](https://kubernetes.io/docs/concepts/services-networking/networkpolicies/) feature. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "node_config",
					Description: `(Optional) Parameters used in creating the default node pool. Generally, this field should not be used at the same time as a ` + "`" + `google_container_node_pool` + "`" + ` or a ` + "`" + `node_pool` + "`" + ` block; this configuration manages the default node pool, which isn't recommended to be used with Terraform. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "node_pool",
					Description: `(Optional) List of node pools associated with this cluster. See [google_container_node_pool](container_node_pool.html) for schema.`,
				},
				resource.Attribute{
					Name:        "node_version",
					Description: `(Optional) The Kubernetes version on the nodes. Must either be unset or set to the same value as ` + "`" + `min_master_version` + "`" + ` on create. Defaults to the default version set by GKE which is not necessarily the latest version. This only affects nodes in the default node pool. While a fuzzy version can be specified, it's recommended that you specify explicit versions as Terraform will see spurious diffs when fuzzy versions are used. See the ` + "`" + `google_container_engine_versions` + "`" + ` data source's ` + "`" + `version_prefix` + "`" + ` field to approximate fuzzy versions in a Terraform-compatible way. To update nodes in other node pools, use the ` + "`" + `version` + "`" + ` attribute on the node pool.`,
				},
				resource.Attribute{
					Name:        "pod_security_policy_config",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) Configuration for the [PodSecurityPolicy](https://cloud.google.com/kubernetes-engine/docs/how-to/pod-security-policies) feature. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "authenticator_groups_config",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) Configuration for the [Google Groups for GKE](https://cloud.google.com/kubernetes-engine/docs/how-to/role-based-access-control#groups-setup-gsuite) feature. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "private_cluster_config",
					Description: `(Optional) A set of options for creating a private cluster. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "remove_default_node_pool",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, deletes the default node pool upon cluster creation. If you're using ` + "`" + `google_container_node_pool` + "`" + ` resources with no default node pool, this should be set to ` + "`" + `true` + "`" + `, alongside setting ` + "`" + `initial_node_count` + "`" + ` to at least ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resource_labels",
					Description: `(Optional) The GCE resource labels (a map of key/value pairs) to be applied to the cluster.`,
				},
				resource.Attribute{
					Name:        "resource_usage_export_config",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) Configuration for the [ResourceUsageExportConfig](https://cloud.google.com/kubernetes-engine/docs/how-to/cluster-usage-metering) feature. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "subnetwork",
					Description: `(Optional) The name or self_link of the Google Compute Engine subnetwork in which the cluster's instances are launched.`,
				},
				resource.Attribute{
					Name:        "vertical_pod_autoscaling",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) Vertical Pod Autoscaling automatically adjusts the resources of pods controlled by it. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "workload_identity_config",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) Workload Identity allows Kubernetes service accounts to act as a user-managed [Google IAM Service Account](https://cloud.google.com/iam/docs/service-accounts#user-managed_service_accounts). Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "enable_intranode_visibility",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) Whether Intra-node visibility is enabled for this cluster. This makes same node pod to pod traffic visible for VPC network. The ` + "`" + `addons_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "horizontal_pod_autoscaling",
					Description: `(Optional) The status of the Horizontal Pod Autoscaling addon, which increases or decreases the number of replica pods a replication controller has based on the resource usage of the existing pods. It ensures that a Heapster pod is running in the cluster, which is also used by the Cloud Monitoring service. It is enabled by default; set ` + "`" + `disabled = true` + "`" + ` to disable.`,
				},
				resource.Attribute{
					Name:        "http_load_balancing",
					Description: `(Optional) The status of the HTTP (L7) load balancing controller addon, which makes it easy to set up HTTP load balancers for services in a cluster. It is enabled by default; set ` + "`" + `disabled = true` + "`" + ` to disable.`,
				},
				resource.Attribute{
					Name:        "kubernetes_dashboard",
					Description: `(Optional) The status of the Kubernetes Dashboard add-on, which controls whether the Kubernetes Dashboard is enabled for this cluster. It is disabled by default; set ` + "`" + `disabled = false` + "`" + ` to enable.`,
				},
				resource.Attribute{
					Name:        "network_policy_config",
					Description: `(Optional) Whether we should enable the network policy addon for the master. This must be enabled in order to enable network policy for the nodes. To enable this, you must also define a [` + "`" + `network_policy` + "`" + `](#network_policy) block, otherwise nothing will happen. It can only be disabled if the nodes already do not have network policies enabled. Defaults to disabled; set ` + "`" + `disabled = false` + "`" + ` to enable.`,
				},
				resource.Attribute{
					Name:        "istio_config",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)). Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "cloudrun_config",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)). The status of the CloudRun addon. It requires ` + "`" + `istio_config` + "`" + ` enabled. It is disabled by default. Set ` + "`" + `disabled = false` + "`" + ` to enable. This addon can only be enabled at cluster creation time. This example ` + "`" + `addons_config` + "`" + ` disables two addons: ` + "`" + `` + "`" + `` + "`" + ` addons_config { http_load_balancing { disabled = true } horizontal_pod_autoscaling { disabled = true } } ` + "`" + `` + "`" + `` + "`" + ` The ` + "`" + `database_encryption` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Required) ` + "`" + `ENCRYPTED` + "`" + ` or ` + "`" + `DECRYPTED` + "`" + ``,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `(Required) the key to use to encrypt/decrypt secrets. See the [DatabaseEncryption definition](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1beta1/projects.locations.clusters#Cluster.DatabaseEncryption) for more information. The ` + "`" + `istio_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) The status of the Istio addon, which makes it easy to set up Istio for services in a cluster. It is disabled by default. Set ` + "`" + `disabled = false` + "`" + ` to enable.`,
				},
				resource.Attribute{
					Name:        "auth",
					Description: `(Optional) The authentication type between services in Istio. Available options include ` + "`" + `AUTH_MUTUAL_TLS` + "`" + `. The ` + "`" + `cluster_autoscaling` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Whether node auto-provisioning is enabled. Resource limits for ` + "`" + `cpu` + "`" + ` and ` + "`" + `memory` + "`" + ` must be defined to enable node auto-provisioning.`,
				},
				resource.Attribute{
					Name:        "resource_limits",
					Description: `(Optional) Global constraints for machine resources in the cluster. Configuring the ` + "`" + `cpu` + "`" + ` and ` + "`" + `memory` + "`" + ` types is required if node auto-provisioning is enabled. These limits will apply to node pool autoscaling in addition to node auto-provisioning. Structure is documented below. The ` + "`" + `resource_limits` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required) The type of the resource. For example, ` + "`" + `cpu` + "`" + ` and ` + "`" + `memory` + "`" + `. See the [guide to using Node Auto-Provisioning](https://cloud.google.com/kubernetes-engine/docs/how-to/node-auto-provisioning) for a list of types.`,
				},
				resource.Attribute{
					Name:        "minimum",
					Description: `(Optional) Minimum amount of the resource in the cluster.`,
				},
				resource.Attribute{
					Name:        "maximum",
					Description: `(Optional) Maximum amount of the resource in the cluster. The ` + "`" + `authenticator_groups_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "security_group",
					Description: `(Required) The name of the RBAC security group for use with Google security groups in Kubernetes RBAC. Group name must be in format ` + "`" + `gke-security-groups@yourdomain.com` + "`" + `. The ` + "`" + `maintenance_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "daily_maintenance_window",
					Description: `(Required) Time window specified for daily maintenance operations. Specify ` + "`" + `start_time` + "`" + ` in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) format "HH:MM”, where HH : \[00-23\] and MM : \[00-59\] GMT. For example: ` + "`" + `` + "`" + `` + "`" + ` maintenance_policy { daily_maintenance_window { start_time = "03:00" } } ` + "`" + `` + "`" + `` + "`" + ` The ` + "`" + `ip_allocation_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "use_ip_aliases",
					Description: `(Optional) Whether alias IPs will be used for pod IPs in the cluster. Defaults to ` + "`" + `true` + "`" + ` if the ` + "`" + `ip_allocation_policy` + "`" + ` block is defined, and to the API default otherwise. Prior to June 17th 2019, the default on the API is ` + "`" + `false` + "`" + `; afterwards, it's ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cluster_secondary_range_name",
					Description: `(Optional) The name of the secondary range to be used as for the cluster CIDR block. The secondary range will be used for pod IP addresses. This must be an existing secondary range associated with the cluster subnetwork.`,
				},
				resource.Attribute{
					Name:        "services_secondary_range_name",
					Description: `(Optional) The name of the secondary range to be used as for the services CIDR block. The secondary range will be used for service ClusterIPs. This must be an existing secondary range associated with the cluster subnetwork.`,
				},
				resource.Attribute{
					Name:        "cluster_ipv4_cidr_block",
					Description: `(Optional) The IP address range for the cluster pod IPs. Set to blank to have a range chosen with the default size. Set to /netmask (e.g. /14) to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use. This field will only work if your cluster is VPC-native- when ` + "`" + `ip_allocation_policy.use_ip_aliases` + "`" + ` is undefined or set to true. If your cluster is not VPC-native, use ` + "`" + `cluster_ipv4_cidr` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "node_ipv4_cidr_block",
					Description: `(Optional) The IP address range of the node IPs in this cluster. This should be set only if ` + "`" + `create_subnetwork` + "`" + ` is true. Set to blank to have a range chosen with the default size. Set to /netmask (e.g. /14) to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use.`,
				},
				resource.Attribute{
					Name:        "services_ipv4_cidr_block",
					Description: `(Optional) The IP address range of the services IPs in this cluster. Set to blank to have a range chosen with the default size. Set to /netmask (e.g. /14) to have a range chosen with a specific netmask. Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use.`,
				},
				resource.Attribute{
					Name:        "subnetwork_name",
					Description: `(Optional) A custom subnetwork name to be used if create_subnetwork is true. If this field is empty, then an automatic name will be chosen for the new subnetwork. The ` + "`" + `master_auth` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The password to use for HTTP basic authentication when accessing the Kubernetes master endpoint.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) The username to use for HTTP basic authentication when accessing the Kubernetes master endpoint. If not present basic auth will be disabled.`,
				},
				resource.Attribute{
					Name:        "client_certificate_config",
					Description: `(Optional) Whether client certificate authorization is enabled for this cluster. For example: ` + "`" + `` + "`" + `` + "`" + ` master_auth { client_certificate_config { issue_client_certificate = false } } ` + "`" + `` + "`" + `` + "`" + ` If this block is provided and both ` + "`" + `username` + "`" + ` and ` + "`" + `password` + "`" + ` are empty, basic authentication will be disabled. This block also contains several computed attributes, documented below. If this block is not provided, GKE will generate a password for you with the username ` + "`" + `admin` + "`" + `. The ` + "`" + `master_authorized_networks_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "cidr_blocks",
					Description: `(Optional) External networks that can access the Kubernetes cluster master through HTTPS. The ` + "`" + `master_authorized_networks_config.cidr_blocks` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Optional) External network that can access Kubernetes master through HTTPS. Must be specified in CIDR notation.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Field for users to identify CIDR blocks. The ` + "`" + `network_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "provider",
					Description: `(Optional) The selected network policy provider. Defaults to PROVIDER_UNSPECIFIED.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether network policy is enabled on the cluster. Defaults to false. The ` + "`" + `node_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "disk_size_gb",
					Description: `(Optional) Size of the disk attached to each node, specified in GB. The smallest allowed disk size is 10GB. Defaults to 100GB.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Optional) Type of the disk attached to each node (e.g. 'pd-standard' or 'pd-ssd'). If unspecified, the default disk type is 'pd-standard'`,
				},
				resource.Attribute{
					Name:        "guest_accelerator",
					Description: `(Optional) List of the type and count of accelerator cards attached to the instance. Structure documented below. To support removal of guest_accelerators in Terraform 0.12 this field is an [Attribute as Block](/docs/configuration/attr-as-blocks.html)`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `(Optional) The image type to use for this node. Note that changing the image type will delete and recreate all nodes in the node pool.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) The Kubernetes labels (key/value pairs) to be applied to each node.`,
				},
				resource.Attribute{
					Name:        "local_ssd_count",
					Description: `(Optional) The amount of local SSD disks that will be attached to each cluster node. Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `(Optional) The name of a Google Compute Engine machine type. Defaults to ` + "`" + `n1-standard-1` + "`" + `. To create a custom machine type, value should be set as specified [here](https://cloud.google.com/compute/docs/reference/latest/instances#machineType).`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) The metadata key/value pairs assigned to instances in the cluster. From GKE ` + "`" + `1.12` + "`" + ` onwards, ` + "`" + `disable-legacy-endpoints` + "`" + ` is set to ` + "`" + `true` + "`" + ` by the API; if ` + "`" + `metadata` + "`" + ` is set but that default value is not included, Terraform will attempt to unset the value. To avoid this, set the value in your config.`,
				},
				resource.Attribute{
					Name:        "min_cpu_platform",
					Description: `(Optional) Minimum CPU platform to be used by this instance. The instance may be scheduled on the specified or newer CPU platform. Applicable values are the friendly names of CPU platforms, such as ` + "`" + `Intel Haswell` + "`" + `. See the [official documentation](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform) for more information.`,
				},
				resource.Attribute{
					Name:        "oauth_scopes",
					Description: `(Optional) The set of Google API scopes to be made available on all of the node VMs under the "default" service account. These can be either FQDNs, or scope aliases. The following scopes are necessary to ensure the correct functioning of the cluster:`,
				},
				resource.Attribute{
					Name:        "preemptible",
					Description: `(Optional) A boolean that represents whether or not the underlying node VMs are preemptible. See the [official documentation](https://cloud.google.com/container-engine/docs/preemptible-vm) for more information. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "sandbox_config",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) [GKE Sandbox](https://cloud.google.com/kubernetes-engine/docs/how-to/sandbox-pods) configuration. When enabling this feature you must specify ` + "`" + `image_type = "COS_CONTAINERD"` + "`" + ` and ` + "`" + `node_version = "1.12.7-gke.17"` + "`" + ` or later to use it. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "service_account",
					Description: `(Optional) The service account to be used by the Node VMs. If not specified, the "default" service account is used. In order to use the configured ` + "`" + `oauth_scopes` + "`" + ` for logging and monitoring, the service account being used needs the [roles/logging.logWriter](https://cloud.google.com/iam/docs/understanding-roles#stackdriver_logging_roles) and [roles/monitoring.metricWriter](https://cloud.google.com/iam/docs/understanding-roles#stackdriver_monitoring_roles) roles. -> Projects that enable the [Cloud Compute Engine API](https://cloud.google.com/compute/) with Terraform may need these roles added manually to the service account. Projects that enable the API in the Cloud Console should have them added automatically.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The list of instance tags applied to all nodes. Tags are used to identify valid sources or targets for network firewalls.`,
				},
				resource.Attribute{
					Name:        "taint",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) List of [kubernetes taints](https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/) to apply to each node. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "workload_metadata_config",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) Metadata configuration to expose to workloads on the node pool. Structure is documented below. The ` + "`" + `guest_accelerator` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "private_endpoint",
					Description: `The internal IP address of this cluster's master endpoint.`,
				},
				resource.Attribute{
					Name:        "public_endpoint",
					Description: `The external IP address of this cluster's master endpoint. The ` + "`" + `sandbox_type` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `The IP address of this cluster's Kubernetes master.`,
				},
				resource.Attribute{
					Name:        "instance_group_urls",
					Description: `List of instance group URLs which have been assigned to the cluster.`,
				},
				resource.Attribute{
					Name:        "maintenance_policy.0.daily_maintenance_window.0.duration",
					Description: `Duration of the time window, automatically chosen to be smallest possible in the given scenario. Duration will be in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) format "PTnHnMnS".`,
				},
				resource.Attribute{
					Name:        "master_auth.0.client_certificate",
					Description: `Base64 encoded public certificate used by clients to authenticate to the cluster endpoint.`,
				},
				resource.Attribute{
					Name:        "master_auth.0.client_key",
					Description: `Base64 encoded private key used by clients to authenticate to the cluster endpoint.`,
				},
				resource.Attribute{
					Name:        "master_auth.0.cluster_ca_certificate",
					Description: `Base64 encoded public certificate that is the root of trust for the cluster.`,
				},
				resource.Attribute{
					Name:        "master_version",
					Description: `The current version of the master in the cluster. This may be different than the ` + "`" + `min_master_version` + "`" + ` set in the config if the master has been updated by GKE.`,
				},
				resource.Attribute{
					Name:        "tpu_ipv4_cidr_block",
					Description: `([Beta](https://terraform.io/docs/providers/google/provider_versions.html)) The IP address range of the Cloud TPUs in this cluster, in [CIDR](http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) notation (e.g. ` + "`" + `1.2.3.4/29` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "services_ipv4_cidr",
					Description: `The IP address range of the Kubernetes services in this cluster, in [CIDR](http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) notation (e.g. ` + "`" + `1.2.3.4/29` + "`" + `). Service addresses are typically put in the last ` + "`" + `/16` + "`" + ` from the container CIDR. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 30 minutes. - ` + "`" + `update` + "`" + ` - Default is 60 minutes. - ` + "`" + `delete` + "`" + ` - Default is 30 minutes. ## Import GKE clusters can be imported using the ` + "`" + `project` + "`" + ` , ` + "`" + `zone` + "`" + ` or ` + "`" + `region` + "`" + `, and ` + "`" + `name` + "`" + `. If the project is omitted, the default provider value will be used. Examples: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_container_cluster.mycluster my-gcp-project/us-east1-a/my-cluster $ terraform import google_container_cluster.mycluster us-east1-a/my-cluster ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "endpoint",
					Description: `The IP address of this cluster's Kubernetes master.`,
				},
				resource.Attribute{
					Name:        "instance_group_urls",
					Description: `List of instance group URLs which have been assigned to the cluster.`,
				},
				resource.Attribute{
					Name:        "maintenance_policy.0.daily_maintenance_window.0.duration",
					Description: `Duration of the time window, automatically chosen to be smallest possible in the given scenario. Duration will be in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) format "PTnHnMnS".`,
				},
				resource.Attribute{
					Name:        "master_auth.0.client_certificate",
					Description: `Base64 encoded public certificate used by clients to authenticate to the cluster endpoint.`,
				},
				resource.Attribute{
					Name:        "master_auth.0.client_key",
					Description: `Base64 encoded private key used by clients to authenticate to the cluster endpoint.`,
				},
				resource.Attribute{
					Name:        "master_auth.0.cluster_ca_certificate",
					Description: `Base64 encoded public certificate that is the root of trust for the cluster.`,
				},
				resource.Attribute{
					Name:        "master_version",
					Description: `The current version of the master in the cluster. This may be different than the ` + "`" + `min_master_version` + "`" + ` set in the config if the master has been updated by GKE.`,
				},
				resource.Attribute{
					Name:        "tpu_ipv4_cidr_block",
					Description: `([Beta](https://terraform.io/docs/providers/google/provider_versions.html)) The IP address range of the Cloud TPUs in this cluster, in [CIDR](http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) notation (e.g. ` + "`" + `1.2.3.4/29` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "services_ipv4_cidr",
					Description: `The IP address range of the Kubernetes services in this cluster, in [CIDR](http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) notation (e.g. ` + "`" + `1.2.3.4/29` + "`" + `). Service addresses are typically put in the last ` + "`" + `/16` + "`" + ` from the container CIDR. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 30 minutes. - ` + "`" + `update` + "`" + ` - Default is 60 minutes. - ` + "`" + `delete` + "`" + ` - Default is 30 minutes. ## Import GKE clusters can be imported using the ` + "`" + `project` + "`" + ` , ` + "`" + `zone` + "`" + ` or ` + "`" + `region` + "`" + `, and ` + "`" + `name` + "`" + `. If the project is omitted, the default provider value will be used. Examples: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_container_cluster.mycluster my-gcp-project/us-east1-a/my-cluster $ terraform import google_container_cluster.mycluster us-east1-a/my-cluster ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_container_node_pool",
			Category:         "Google Kubernetes (Container) Engine Resources",
			ShortDescription: `Manages a GKE NodePool resource.`,
			Description:      ``,
			Keywords: []string{
				"kubernetes",
				"container",
				"engine",
				"node",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster",
					Description: `(Required) The cluster to create the node pool for. Cluster must be present in ` + "`" + `zone` + "`" + ` provided for zonal clusters. - - -`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Optional) The location (region or zone) in which the cluster resides.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional, Deprecated) The zone in which the cluster resides. ` + "`" + `zone` + "`" + ` has been deprecated in favor of ` + "`" + `location` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, Deprecated) The region in which the cluster resides (for regional clusters). ` + "`" + `region` + "`" + ` has been deprecated in favor of ` + "`" + `location` + "`" + `. -> Note: You must specify a ` + "`" + `location` + "`" + ` for either cluster type or the type-specific ` + "`" + `region` + "`" + ` for regional clusters / ` + "`" + `zone` + "`" + ` for zonal clusters. - - -`,
				},
				resource.Attribute{
					Name:        "autoscaling",
					Description: `(Optional) Configuration required by cluster autoscaler to adjust the size of the node pool to the current cluster usage. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "initial_node_count",
					Description: `(Optional) The initial node count for the pool. Changing this will force recreation of the resource.`,
				},
				resource.Attribute{
					Name:        "management",
					Description: `(Optional) Node management configuration, wherein auto-repair and auto-upgrade is configured. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "max_pods_per_node",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) The maximum number of pods per node in this node pool. Note that this does not work on node pools which are "route-based" - that is, node pools belonging to clusters that do not have IP Aliasing enabled. See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr) for more information.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the node pool. If left blank, Terraform will auto-generate a unique name.`,
				},
				resource.Attribute{
					Name:        "node_config",
					Description: `(Optional) The node configuration of the pool. See [google_container_cluster](container_cluster.html) for schema.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `(Optional) The number of nodes per instance group. This field can be used to update the number of nodes per instance group but should not be used alongside ` + "`" + `autoscaling` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which to create the node pool. If blank, the provider-configured project will be used.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) The Kubernetes version for the nodes in this pool. Note that if this field and ` + "`" + `auto_upgrade` + "`" + ` are both specified, they will fight each other for what the node version should be, so setting both is highly discouraged. While a fuzzy version can be specified, it's recommended that you specify explicit versions as Terraform will see spurious diffs when fuzzy versions are used. See the ` + "`" + `google_container_engine_versions` + "`" + ` data source's ` + "`" + `version_prefix` + "`" + ` field to approximate fuzzy versions in a Terraform-compatible way. The ` + "`" + `autoscaling` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "min_node_count",
					Description: `(Required) Minimum number of nodes in the NodePool. Must be >=0 and <= ` + "`" + `max_node_count` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_node_count",
					Description: `(Required) Maximum number of nodes in the NodePool. Must be >= min_node_count. The ` + "`" + `management` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "auto_repair",
					Description: `(Optional) Whether the nodes will be automatically repaired.`,
				},
				resource.Attribute{
					Name:        "auto_upgrade",
					Description: `(Optional) Whether the nodes will be automatically upgraded. <a id="timeouts"></a> ## Timeouts ` + "`" + `google_container_node_pool` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `30 minutes` + "`" + `) Used for adding node pools - ` + "`" + `update` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for updates to node pools - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for removing node pools. ## Import Node pools can be imported using the ` + "`" + `project` + "`" + `, ` + "`" + `zone` + "`" + `, ` + "`" + `cluster` + "`" + ` and ` + "`" + `name` + "`" + `. If the project is omitted, the default provider value will be used. Examples: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_container_node_pool.mainpool my-gcp-project/us-east1-a/my-cluster/main-pool $ terraform import google_container_node_pool.mainpool us-east1-a/my-cluster/main-pool ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_dataflow_job",
			Category:         "Google Dataflow Resources",
			ShortDescription: `Creates a job in Dataflow according to a provided config file.`,
			Description:      ``,
			Keywords: []string{
				"dataflow",
				"job",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the resource, required by Dataflow.`,
				},
				resource.Attribute{
					Name:        "template_gcs_path",
					Description: `(Required) The GCS path to the Dataflow job template.`,
				},
				resource.Attribute{
					Name:        "temp_gcs_location",
					Description: `(Required) A writeable location on GCS for the Dataflow job to dump its temporary data. - - -`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Optional) Key/Value pairs to be passed to the Dataflow job (as used in the template).`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) User labels to be specified for the job. Keys and values should follow the restrictions specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions) page.`,
				},
				resource.Attribute{
					Name:        "max_workers",
					Description: `(Optional) The number of workers permitted to work on the job. More workers may improve processing speed at additional cost.`,
				},
				resource.Attribute{
					Name:        "on_delete",
					Description: `(Optional) One of "drain" or "cancel". Specifies behavior of deletion during ` + "`" + `terraform destroy` + "`" + `. See above note.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) The zone in which the created job should run. If it is not provided, the provider zone is used.`,
				},
				resource.Attribute{
					Name:        "service_account_email",
					Description: `(Optional) The Service Account email used to create the job.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) The network to which VMs will be assigned. If it is not provided, "default" will be used.`,
				},
				resource.Attribute{
					Name:        "subnetwork",
					Description: `(Optional) The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK".`,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `(Optional) The machine type to use for the job. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "state",
					Description: `The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_dataproc_autoscaling_policy",
			Category:         "Google Dataproc Resources",
			ShortDescription: `Describes an autoscaling policy for Dataproc cluster autoscaler.`,
			Description:      ``,
			Keywords: []string{
				"dataproc",
				"autoscaling",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Required) The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 50 characters. - - -`,
				},
				resource.Attribute{
					Name:        "worker_config",
					Description: `(Optional) Describes how the autoscaler will operate for primary workers. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "secondary_worker_config",
					Description: `(Optional) Describes how the autoscaler will operate for secondary workers. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "basic_algorithm",
					Description: `(Optional) Basic algorithm for autoscaling. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Optional) The location where the autoscaling poicy should reside. The default value is ` + "`" + `global` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `worker_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "min_instances",
					Description: `(Optional) Minimum number of instances for this group. Bounds: [2, maxInstances]. Defaults to 2.`,
				},
				resource.Attribute{
					Name:        "max_instances",
					Description: `(Required) Maximum number of instances for this group.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Weight for the instance group, which is used to determine the fraction of total workers in the cluster from this instance group. For example, if primary workers have weight 2, and secondary workers have weight 1, the cluster will have approximately 2 primary workers for each secondary worker. The cluster may not reach the specified balance if constrained by min/max bounds or other autoscaling settings. For example, if maxInstances for secondary workers is 0, then only primary workers will be added. The cluster can also be out of balance when created. If weight is not set on any instance group, the cluster will default to equal weight for all groups: the cluster will attempt to maintain an equal number of workers in each group within the configured size bounds for each group. If weight is set for one group only, the cluster will default to zero weight on the unset group. For example if weight is set only on primary workers, the cluster will use primary workers only and no secondary workers. The ` + "`" + `secondary_worker_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "min_instances",
					Description: `(Optional) Minimum number of instances for this group. Bounds: [0, maxInstances]. Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "max_instances",
					Description: `(Optional) Maximum number of instances for this group. Note that by default, clusters will not use secondary workers. Required for secondary workers if the minimum secondary instances is set. Bounds: [minInstances, ). Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `(Optional) Weight for the instance group, which is used to determine the fraction of total workers in the cluster from this instance group. For example, if primary workers have weight 2, and secondary workers have weight 1, the cluster will have approximately 2 primary workers for each secondary worker. The cluster may not reach the specified balance if constrained by min/max bounds or other autoscaling settings. For example, if maxInstances for secondary workers is 0, then only primary workers will be added. The cluster can also be out of balance when created. If weight is not set on any instance group, the cluster will default to equal weight for all groups: the cluster will attempt to maintain an equal number of workers in each group within the configured size bounds for each group. If weight is set for one group only, the cluster will default to zero weight on the unset group. For example if weight is set only on primary workers, the cluster will use primary workers only and no secondary workers. The ` + "`" + `basic_algorithm` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "cooldown_period",
					Description: `(Optional) Duration between scaling events. A scaling period starts after the update operation from the previous event has completed. Bounds: [2m, 1d]. Default: 2m.`,
				},
				resource.Attribute{
					Name:        "yarn_config",
					Description: `(Required) YARN autoscaling configuration. Structure is documented below. The ` + "`" + `yarn_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "graceful_decommission_timeout",
					Description: `(Required) Timeout for YARN graceful decommissioning of Node Managers. Specifies the duration to wait for jobs to complete before forcefully removing workers (and potentially interrupting jobs). Only applicable to downscaling operations. Bounds: [0s, 1d].`,
				},
				resource.Attribute{
					Name:        "scale_up_factor",
					Description: `(Required) Fraction of average pending memory in the last cooldown period for which to add workers. A scale-up factor of 1.0 will result in scaling up so that there is no pending memory remaining after the update (more aggressive scaling). A scale-up factor closer to 0 will result in a smaller magnitude of scaling up (less aggressive scaling). Bounds: [0.0, 1.0].`,
				},
				resource.Attribute{
					Name:        "scale_down_factor",
					Description: `(Required) Fraction of average pending memory in the last cooldown period for which to remove workers. A scale-down factor of 1 will result in scaling down so that there is no available memory remaining after the update (more aggressive scaling). A scale-down factor of 0 disables removing workers, which can be beneficial for autoscaling a single job. Bounds: [0.0, 1.0].`,
				},
				resource.Attribute{
					Name:        "scale_up_min_worker_fraction",
					Description: `(Optional) Minimum scale-up threshold as a fraction of total cluster size before scaling occurs. For example, in a 20-worker cluster, a threshold of 0.1 means the autoscaler must recommend at least a 2-worker scale-up for the cluster to scale. A threshold of 0 means the autoscaler will scale up on any recommended change. Bounds: [0.0, 1.0]. Default: 0.0.`,
				},
				resource.Attribute{
					Name:        "scale_down_min_worker_fraction",
					Description: `(Optional) Minimum scale-down threshold as a fraction of total cluster size before scaling occurs. For example, in a 20-worker cluster, a threshold of 0.1 means the autoscaler must recommend at least a 2 worker scale-down for the cluster to scale. A threshold of 0 means the autoscaler will scale down on any recommended change. Bounds: [0.0, 1.0]. Default: 0.0. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The "resource name" of the autoscaling policy. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import AutoscalingPolicy can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import -provider=google-beta google_dataproc_autoscaling_policy.default projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{policy_id}} $ terraform import -provider=google-beta google_dataproc_autoscaling_policy.default {{project}}/{{location}}/{{policy_id}} $ terraform import -provider=google-beta google_dataproc_autoscaling_policy.default {{location}}/{{policy_id}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The "resource name" of the autoscaling policy. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import AutoscalingPolicy can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import -provider=google-beta google_dataproc_autoscaling_policy.default projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{policy_id}} $ terraform import -provider=google-beta google_dataproc_autoscaling_policy.default {{project}}/{{location}}/{{policy_id}} $ terraform import -provider=google-beta google_dataproc_autoscaling_policy.default {{location}}/{{policy_id}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_dataproc_cluster",
			Category:         "Google Dataproc Resources",
			ShortDescription: `Manages a Cloud Dataproc cluster resource.`,
			Description:      ``,
			Keywords: []string{
				"dataproc",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the cluster, unique within the project and zone. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the ` + "`" + `cluster` + "`" + ` will exist. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which the cluster and associated nodes will be created in. Defaults to ` + "`" + `global` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional, Computed) The list of labels (key/value pairs) to be applied to instances in the cluster. GCP generates some itself including ` + "`" + `goog-dataproc-cluster-name` + "`" + ` which is the name of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_config",
					Description: `(Optional) Allows you to configure various aspects of the cluster. Structure defined below. - - - The ` + "`" + `cluster_config` + "`" + ` block supports: ` + "`" + `` + "`" + `` + "`" + `hcl cluster_config { gce_cluster_config { ... } master_config { ... } worker_config { ... } preemptible_worker_config { ... } software_config { ... } # You can define multiple initialization_action blocks initialization_action { ... } encryption_config { ... } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "staging_bucket",
					Description: `(Optional) The Cloud Storage staging bucket used to stage files, such as Hadoop jars, between client machines and the cluster. Note: If you don't explicitly specify a ` + "`" + `staging_bucket` + "`" + ` then GCP will auto create / assign one for you. However, you are not guaranteed an auto generated bucket which is solely dedicated to your cluster; it may be shared with other clusters in the same region/zone also choosing to use the auto generation option.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional, Computed) The GCP zone where your data is stored and used (i.e. where the master and the worker nodes will be created in). If ` + "`" + `region` + "`" + ` is set to 'global' (default) then ` + "`" + `zone` + "`" + ` is mandatory, otherwise GCP is able to make use of [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/auto-zone) to determine this automatically for you. Note: This setting additionally determines and restricts which computing resources are available for use with other configs such as ` + "`" + `cluster_config.master_config.machine_type` + "`" + ` and ` + "`" + `cluster_config.worker_config.machine_type` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional, Computed) The name or self_link of the Google Compute Engine network to the cluster will be part of. Conflicts with ` + "`" + `subnetwork` + "`" + `. If neither is specified, this defaults to the "default" network.`,
				},
				resource.Attribute{
					Name:        "subnetwork",
					Description: `(Optional) The name or self_link of the Google Compute Engine subnetwork the cluster will be part of. Conflicts with ` + "`" + `network` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_account",
					Description: `(Optional) The service account to be used by the Node VMs. If not specified, the "default" service account is used.`,
				},
				resource.Attribute{
					Name:        "service_account_scopes",
					Description: `(Optional, Computed) The set of Google API scopes to be made available on all of the node VMs under the ` + "`" + `service_account` + "`" + ` specified. These can be either FQDNs, or scope aliases. The following scopes are necessary to ensure the correct functioning of the cluster:`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The list of instance tags applied to instances in the cluster. Tags are used to identify valid sources or targets for network firewalls.`,
				},
				resource.Attribute{
					Name:        "internal_ip_only",
					Description: `(Optional) By default, clusters are not restricted to internal IP addresses, and will have ephemeral external IP addresses assigned to each instance. If set to true, all instances in the cluster will only have internal IP addresses. Note: Private Google Access (also known as ` + "`" + `privateIpGoogleAccess` + "`" + `) must be enabled on the subnetwork that the cluster will be launched in.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) A map of the Compute Engine metadata entries to add to all instances (see [Project and instance metadata](https://cloud.google.com/compute/docs/storing-retrieving-metadata#project_and_instance_metadata)). - - - The ` + "`" + `cluster_config.master_config` + "`" + ` block supports: ` + "`" + `` + "`" + `` + "`" + `hcl cluster_config { master_config { num_instances = 1 machine_type = "n1-standard-1" min_cpu_platform = "Intel Skylake" disk_config { boot_disk_type = "pd-ssd" boot_disk_size_gb = 15 num_local_ssds = 1 } } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `(Optional, Computed) The name of a Google Compute Engine machine type to create for the master. If not specified, GCP will default to a predetermined computed value (currently ` + "`" + `n1-standard-4` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "min_cpu_platform",
					Description: `(Optional, Computed, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) The name of a minimum generation of CPU family for the master. If not specified, GCP will default to a predetermined computed value for each zone. See [the guide](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform) for details about which CPU families are available (and defaulted) for each zone.`,
				},
				resource.Attribute{
					Name:        "boot_disk_type",
					Description: `(Optional) The disk type of the primary disk attached to each node. One of ` + "`" + `"pd-ssd"` + "`" + ` or ` + "`" + `"pd-standard"` + "`" + `. Defaults to ` + "`" + `"pd-standard"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "boot_disk_size_gb",
					Description: `(Optional, Computed) Size of the primary disk attached to each node, specified in GB. The primary disk contains the boot volume and system libraries, and the smallest allowed disk size is 10GB. GCP will default to a predetermined computed value if not set (currently 500GB). Note: If SSDs are not attached, it also contains the HDFS data blocks and Hadoop working directories.`,
				},
				resource.Attribute{
					Name:        "num_local_ssds",
					Description: `(Optional) The amount of local SSD disks that will be attached to each master cluster node. Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "accelerator_type",
					Description: `(Required) The short name of the accelerator type to expose to this instance. For example, ` + "`" + `nvidia-tesla-k80` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "accelerator_count",
					Description: `(Required) The number of the accelerator cards of this type exposed to this instance. Often restricted to one of ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `4` + "`" + `, or ` + "`" + `8` + "`" + `. ~> The Cloud Dataproc API can return unintuitive error messages when using accelerators; even when you have defined an accelerator, Auto Zone Placement does not exclusively select zones that have that accelerator available. If you get a 400 error that the accelerator can't be found, this is a likely cause. Make sure you check [accelerator availability by zone](https://cloud.google.com/compute/docs/reference/rest/v1/acceleratorTypes/list) if you are trying to use accelerators in a given zone. - - - The ` + "`" + `cluster_config.worker_config` + "`" + ` block supports: ` + "`" + `` + "`" + `` + "`" + `hcl cluster_config { worker_config { num_instances = 3 machine_type = "n1-standard-1" min_cpu_platform = "Intel Skylake" disk_config { boot_disk_type = "pd-standard" boot_disk_size_gb = 15 num_local_ssds = 1 } } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "machine_type",
					Description: `(Optional, Computed) The name of a Google Compute Engine machine type to create for the worker nodes. If not specified, GCP will default to a predetermined computed value (currently ` + "`" + `n1-standard-4` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "min_cpu_platform",
					Description: `(Optional, Computed, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) The name of a minimum generation of CPU family for the master. If not specified, GCP will default to a predetermined computed value for each zone. See [the guide](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform) for details about which CPU families are available (and defaulted) for each zone.`,
				},
				resource.Attribute{
					Name:        "boot_disk_type",
					Description: `(Optional) The disk type of the primary disk attached to each node. One of ` + "`" + `"pd-ssd"` + "`" + ` or ` + "`" + `"pd-standard"` + "`" + `. Defaults to ` + "`" + `"pd-standard"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "boot_disk_size_gb",
					Description: `(Optional, Computed) Size of the primary disk attached to each worker node, specified in GB. The smallest allowed disk size is 10GB. GCP will default to a predetermined computed value if not set (currently 500GB). Note: If SSDs are not attached, it also contains the HDFS data blocks and Hadoop working directories.`,
				},
				resource.Attribute{
					Name:        "num_local_ssds",
					Description: `(Optional) The amount of local SSD disks that will be attached to each worker cluster node. Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "accelerator_type",
					Description: `(Required) The short name of the accelerator type to expose to this instance. For example, ` + "`" + `nvidia-tesla-k80` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "accelerator_count",
					Description: `(Required) The number of the accelerator cards of this type exposed to this instance. Often restricted to one of ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `4` + "`" + `, or ` + "`" + `8` + "`" + `. ~> The Cloud Dataproc API can return unintuitive error messages when using accelerators; even when you have defined an accelerator, Auto Zone Placement does not exclusively select zones that have that accelerator available. If you get a 400 error that the accelerator can't be found, this is a likely cause. Make sure you check [accelerator availability by zone](https://cloud.google.com/compute/docs/reference/rest/v1/acceleratorTypes/list) if you are trying to use accelerators in a given zone. - - - The ` + "`" + `cluster_config.preemptible_worker_config` + "`" + ` block supports: ` + "`" + `` + "`" + `` + "`" + `hcl cluster_config { preemptible_worker_config { num_instances = 1 disk_config { boot_disk_type = "pd-standard" boot_disk_size_gb = 15 num_local_ssds = 1 } } } ` + "`" + `` + "`" + `` + "`" + ` Note: Unlike ` + "`" + `worker_config` + "`" + `, you cannot set the ` + "`" + `machine_type` + "`" + ` value directly. This will be set for you based on whatever was set for the ` + "`" + `worker_config.machine_type` + "`" + ` value.`,
				},
				resource.Attribute{
					Name:        "boot_disk_type",
					Description: `(Optional) The disk type of the primary disk attached to each preemptible worker node. One of ` + "`" + `"pd-ssd"` + "`" + ` or ` + "`" + `"pd-standard"` + "`" + `. Defaults to ` + "`" + `"pd-standard"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "boot_disk_size_gb",
					Description: `(Optional, Computed) Size of the primary disk attached to each preemptible worker node, specified in GB. The smallest allowed disk size is 10GB. GCP will default to a predetermined computed value if not set (currently 500GB). Note: If SSDs are not attached, it also contains the HDFS data blocks and Hadoop working directories.`,
				},
				resource.Attribute{
					Name:        "num_local_ssds",
					Description: `(Optional) The amount of local SSD disks that will be attached to each preemptible worker node. Defaults to 0. - - - The ` + "`" + `cluster_config.software_config` + "`" + ` block supports: ` + "`" + `` + "`" + `` + "`" + `hcl cluster_config { # Override or set some custom properties software_config { image_version = "1.3.7-deb9" override_properties = { "dataproc:dataproc.allow.zero.workers" = "true" } } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "image_version",
					Description: `(Optional, Computed) The Cloud Dataproc image version to use for the cluster - this controls the sets of software versions installed onto the nodes when you create clusters. If not specified, defaults to the latest version. For a list of valid versions see [Cloud Dataproc versions](https://cloud.google.com/dataproc/docs/concepts/dataproc-versions)`,
				},
				resource.Attribute{
					Name:        "override_properties",
					Description: `(Optional) A list of override and additional properties (key/value pairs) used to modify various aspects of the common configuration files used when creating a cluster. For a list of valid properties please see [Cluster properties](https://cloud.google.com/dataproc/docs/concepts/cluster-properties) - - - The ` + "`" + `initialization_action` + "`" + ` block (Optional) can be specified multiple times and supports: ` + "`" + `` + "`" + `` + "`" + `hcl cluster_config { # You can define multiple initialization_action blocks initialization_action { script = "gs://dataproc-initialization-actions/stackdriver/stackdriver.sh" timeout_sec = 500 } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "timeout_sec",
					Description: `(Optional, Computed) The maximum duration (in seconds) which ` + "`" + `script` + "`" + ` is allowed to take to execute its action. GCP will default to a predetermined computed value if not set (currently 300). - - - The ` + "`" + `encryption_config` + "`" + ` block supports: ` + "`" + `` + "`" + `` + "`" + `hcl cluster_config { encryption_config { kms_key_name = "projects/projectId/locations/region/keyRings/keyRingName/cryptoKeys/keyName" } } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "kms_key_name",
					Description: `(Required) The Cloud KMS key name to use for PD disk encryption for all instances in the cluster. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_config.0.master_config.0.instance_names",
					Description: `List of master instance names which have been assigned to the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_config.0.worker_config.0.instance_names",
					Description: `List of worker instance names which have been assigned to the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_config.0.preemptible_worker_config.0.instance_names",
					Description: `List of preemptible instance names which have been assigned to the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_config.0.bucket",
					Description: `The name of the cloud storage bucket ultimately used to house the staging data for the cluster. If ` + "`" + `staging_bucket` + "`" + ` is specified, it will contain this value, otherwise it will be the auto generated name.`,
				},
				resource.Attribute{
					Name:        "cluster_config.0.software_config.0.properties",
					Description: `A list of the properties used to set the daemon config files. This will include any values supplied by the user via ` + "`" + `cluster_config.software_config.override_properties` + "`" + ` ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 20 minutes. - ` + "`" + `update` + "`" + ` - Default is 20 minutes. - ` + "`" + `delete` + "`" + ` - Default is 20 minutes.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_config.0.master_config.0.instance_names",
					Description: `List of master instance names which have been assigned to the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_config.0.worker_config.0.instance_names",
					Description: `List of worker instance names which have been assigned to the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_config.0.preemptible_worker_config.0.instance_names",
					Description: `List of preemptible instance names which have been assigned to the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_config.0.bucket",
					Description: `The name of the cloud storage bucket ultimately used to house the staging data for the cluster. If ` + "`" + `staging_bucket` + "`" + ` is specified, it will contain this value, otherwise it will be the auto generated name.`,
				},
				resource.Attribute{
					Name:        "cluster_config.0.software_config.0.properties",
					Description: `A list of the properties used to set the daemon config files. This will include any values supplied by the user via ` + "`" + `cluster_config.software_config.override_properties` + "`" + ` ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 20 minutes. - ` + "`" + `update` + "`" + ` - Default is 20 minutes. - ` + "`" + `delete` + "`" + ` - Default is 20 minutes.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_dataproc_cluster_iam_policy",
			Category:         "Google Dataproc Resources",
			ShortDescription: `Collection of resources to manage IAM policy for a Dataproc cluster.`,
			Description:      ``,
			Keywords: []string{
				"dataproc",
				"cluster",
				"iam",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster",
					Description: `(Required) The name or relative resource id of the cluster to manage IAM policies for. For ` + "`" + `google_dataproc_cluster_iam_member` + "`" + ` or ` + "`" + `google_dataproc_cluster_iam_binding` + "`" + `:`,
				},
				resource.Attribute{
					Name:        "member/members",
					Description: `(Required) Identities that will be granted the privilege in ` + "`" + `role` + "`" + `. Each entry can have one of the following values:`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. Only one ` + "`" + `google_dataproc_cluster_iam_binding` + "`" + ` can be used per role. Note that custom roles must be of the format ` + "`" + `[projects|organizations]/{parent-name}/roles/{role-name}` + "`" + `. ` + "`" + `google_dataproc_cluster_iam_policy` + "`" + ` only:`,
				},
				resource.Attribute{
					Name:        "policy_data",
					Description: `(Required) The policy data generated by a ` + "`" + `google_iam_policy` + "`" + ` data source. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the cluster belongs. If it is not provided, Terraform will use the provider default.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which the cluster belongs. If it is not provided, Terraform will use the provider default. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the clusters's IAM policy. ## Import Cluster IAM resources can be imported using the project, region, cluster name, role and/or member. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_dataproc_cluster_iam_policy.editor "projects/{project}/regions/{region}/clusters/{cluster}" $ terraform import google_dataproc_cluster_iam_binding.editor "projects/{project}/regions/{region}/clusters/{cluster} roles/editor" $ terraform import google_dataproc_cluster_iam_member.editor "projects/{project}/regions/{region}/clusters/{cluster} roles/editor user:jane@example.com" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the clusters's IAM policy. ## Import Cluster IAM resources can be imported using the project, region, cluster name, role and/or member. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_dataproc_cluster_iam_policy.editor "projects/{project}/regions/{region}/clusters/{cluster}" $ terraform import google_dataproc_cluster_iam_binding.editor "projects/{project}/regions/{region}/clusters/{cluster} roles/editor" $ terraform import google_dataproc_cluster_iam_member.editor "projects/{project}/regions/{region}/clusters/{cluster} roles/editor user:jane@example.com" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_dataproc_job",
			Category:         "Google Dataproc Resources",
			ShortDescription: `Manages a job resource within a Dataproc cluster.`,
			Description:      ``,
			Keywords: []string{
				"dataproc",
				"job",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "placement.cluster_name",
					Description: `(Required) The name of the cluster where the job will be submitted.`,
				},
				resource.Attribute{
					Name:        "xxx_config",
					Description: `(Required) Exactly one of the specific job types to run on the cluster should be specified. If you want to submit multiple jobs, this will currently require the definition of multiple ` + "`" + `google_dataproc_job` + "`" + ` resources as shown in the example above, or by setting the ` + "`" + `count` + "`" + ` attribute. The following job configs are supported:`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the ` + "`" + `cluster` + "`" + ` can be found and jobs subsequently run against. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The Cloud Dataproc region. This essentially determines which clusters are available for this job to be submitted to. If not specified, defaults to ` + "`" + `global` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "force_delete",
					Description: `(Optional) By default, you can only delete inactive jobs within Dataproc. Setting this to true, and calling destroy, will ensure that the job is first cancelled before issuing the delete.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) The list of labels (key/value pairs) to add to the job.`,
				},
				resource.Attribute{
					Name:        "scheduling.max_failures_per_hour",
					Description: `(Optional) Maximum number of times per hour a driver may be restarted as a result of driver terminating with non-zero code before job is reported failed. The ` + "`" + `pyspark_config` + "`" + ` block supports: Submitting a pyspark job to the cluster. Below is an example configuration: ` + "`" + `` + "`" + `` + "`" + `hcl # Submit a pyspark job to the cluster resource "google_dataproc_job" "pyspark" { ... pyspark_config { main_python_file_uri = "gs://dataproc-examples-2f10d78d114f6aaec76462e3c310f31f/src/pyspark/hello-world/hello-world.py" properties = { "spark.logConf" = "true" } } } ` + "`" + `` + "`" + `` + "`" + ` For configurations requiring Hadoop Compatible File System (HCFS) references, the options below are generally applicable: - GCS files with the ` + "`" + `gs://` + "`" + ` prefix - HDFS files on the cluster with the ` + "`" + `hdfs://` + "`" + ` prefix - Local files on the cluster with the ` + "`" + `file://` + "`" + ` prefix`,
				},
				resource.Attribute{
					Name:        "args",
					Description: `(Optional) The arguments to pass to the driver.`,
				},
				resource.Attribute{
					Name:        "python_file_uris",
					Description: `(Optional) HCFS file URIs of Python files to pass to the PySpark framework. Supported file types: .py, .egg, and .zip.`,
				},
				resource.Attribute{
					Name:        "jar_file_uris",
					Description: `(Optional) HCFS URIs of jar files to add to the CLASSPATHs of the Python driver and tasks.`,
				},
				resource.Attribute{
					Name:        "file_uris",
					Description: `(Optional) HCFS URIs of files to be copied to the working directory of Python drivers and distributed tasks. Useful for naively parallel tasks.`,
				},
				resource.Attribute{
					Name:        "archive_uris",
					Description: `(Optional) HCFS URIs of archives to be extracted in the working directory of .jar, .tar, .tar.gz, .tgz, and .zip.`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `(Optional) A mapping of property names to values, used to configure PySpark. Properties that conflict with values set by the Cloud Dataproc API may be overwritten. Can include properties set in ` + "`" + `/etc/spark/conf/spark-defaults.conf` + "`" + ` and classes in user code.`,
				},
				resource.Attribute{
					Name:        "main_jar_file_uri",
					Description: `(Optional) The HCFS URI of jar file containing the driver jar. Conflicts with ` + "`" + `main_class` + "`" + ``,
				},
				resource.Attribute{
					Name:        "args",
					Description: `(Optional) The arguments to pass to the driver.`,
				},
				resource.Attribute{
					Name:        "jar_file_uris",
					Description: `(Optional) HCFS URIs of jar files to add to the CLASSPATHs of the Spark driver and tasks.`,
				},
				resource.Attribute{
					Name:        "file_uris",
					Description: `(Optional) HCFS URIs of files to be copied to the working directory of Spark drivers and distributed tasks. Useful for naively parallel tasks.`,
				},
				resource.Attribute{
					Name:        "archive_uris",
					Description: `(Optional) HCFS URIs of archives to be extracted in the working directory of .jar, .tar, .tar.gz, .tgz, and .zip.`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `(Optional) A mapping of property names to values, used to configure Spark. Properties that conflict with values set by the Cloud Dataproc API may be overwritten. Can include properties set in ` + "`" + `/etc/spark/conf/spark-defaults.conf` + "`" + ` and classes in user code.`,
				},
				resource.Attribute{
					Name:        "main_jar_file_uri",
					Description: `(Optional) The HCFS URI of the jar file containing the main class. Examples: 'gs://foo-bucket/analytics-binaries/extract-useful-metrics-mr.jar' 'hdfs:/tmp/test-samples/custom-wordcount.jar' 'file:///home/usr/lib/hadoop-mapreduce/hadoop-mapreduce-examples.jar'. Conflicts with ` + "`" + `main_class` + "`" + ``,
				},
				resource.Attribute{
					Name:        "args",
					Description: `(Optional) The arguments to pass to the driver. Do not include arguments, such as -libjars or -Dfoo=bar, that can be set as job properties, since a collision may occur that causes an incorrect job submission.`,
				},
				resource.Attribute{
					Name:        "jar_file_uris",
					Description: `(Optional) HCFS URIs of jar files to add to the CLASSPATHs of the Spark driver and tasks.`,
				},
				resource.Attribute{
					Name:        "file_uris",
					Description: `(Optional) HCFS URIs of files to be copied to the working directory of Hadoop drivers and distributed tasks. Useful for naively parallel tasks.`,
				},
				resource.Attribute{
					Name:        "archive_uris",
					Description: `(Optional) HCFS URIs of archives to be extracted in the working directory of .jar, .tar, .tar.gz, .tgz, and .zip.`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `(Optional) A mapping of property names to values, used to configure Hadoop. Properties that conflict with values set by the Cloud Dataproc API may be overwritten. Can include properties set in ` + "`" + `/etc/hadoop/conf/`,
				},
				resource.Attribute{
					Name:        "query_file_uri",
					Description: `(Optional) HCFS URI of file containing Hive script to execute as the job. Conflicts with ` + "`" + `query_list` + "`" + ``,
				},
				resource.Attribute{
					Name:        "continue_on_failure",
					Description: `(Optional) Whether to continue executing queries if a query fails. The default value is false. Setting to true can be useful when executing independent parallel queries. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "script_variables",
					Description: `(Optional) Mapping of query variable names to values (equivalent to the Hive command: ` + "`" + `SET name="value";` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `(Optional) A mapping of property names and values, used to configure Hive. Properties that conflict with values set by the Cloud Dataproc API may be overwritten. Can include properties set in ` + "`" + `/etc/hadoop/conf/`,
				},
				resource.Attribute{
					Name:        "jar_file_uris",
					Description: `(Optional) HCFS URIs of jar files to add to the CLASSPATH of the Hive server and Hadoop MapReduce (MR) tasks. Can contain Hive SerDes and UDFs. The ` + "`" + `pig_config` + "`" + ` block supports: ` + "`" + `` + "`" + `` + "`" + `hcl # Submit a pig job to the cluster resource "google_dataproc_job" "pig" { ... pig_config { query_list = [ "LNS = LOAD 'file:///usr/lib/pig/LICENSE.txt ' AS (line)", "WORDS = FOREACH LNS GENERATE FLATTEN(TOKENIZE(line)) AS word", "GROUPS = GROUP WORDS BY word", "WORD_COUNTS = FOREACH GROUPS GENERATE group, COUNT(WORDS)", "DUMP WORD_COUNTS" ] } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "query_file_uri",
					Description: `(Optional) HCFS URI of file containing Hive script to execute as the job. Conflicts with ` + "`" + `query_list` + "`" + ``,
				},
				resource.Attribute{
					Name:        "continue_on_failure",
					Description: `(Optional) Whether to continue executing queries if a query fails. The default value is false. Setting to true can be useful when executing independent parallel queries. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "script_variables",
					Description: `(Optional) Mapping of query variable names to values (equivalent to the Pig command: ` + "`" + `name=[value]` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `(Optional) A mapping of property names to values, used to configure Pig. Properties that conflict with values set by the Cloud Dataproc API may be overwritten. Can include properties set in ` + "`" + `/etc/hadoop/conf/`,
				},
				resource.Attribute{
					Name:        "jar_file_uris",
					Description: `(Optional) HCFS URIs of jar files to add to the CLASSPATH of the Pig Client and Hadoop MapReduce (MR) tasks. Can contain Pig UDFs.`,
				},
				resource.Attribute{
					Name:        "query_file_uri",
					Description: `(Optional) The HCFS URI of the script that contains SQL queries. Conflicts with ` + "`" + `query_list` + "`" + ``,
				},
				resource.Attribute{
					Name:        "script_variables",
					Description: `(Optional) Mapping of query variable names to values (equivalent to the Spark SQL command: ` + "`" + `SET name="value";` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `(Optional) A mapping of property names to values, used to configure Spark SQL's SparkConf. Properties that conflict with values set by the Cloud Dataproc API may be overwritten.`,
				},
				resource.Attribute{
					Name:        "jar_file_uris",
					Description: `(Optional) HCFS URIs of jar files to be added to the Spark CLASSPATH.`,
				},
				resource.Attribute{
					Name:        "reference.0.cluster_uuid",
					Description: `A cluster UUID generated by the Cloud Dataproc service when the job is submitted.`,
				},
				resource.Attribute{
					Name:        "status.0.state",
					Description: `A state message specifying the overall job state.`,
				},
				resource.Attribute{
					Name:        "status.0.details",
					Description: `Optional job state details, such as an error description if the state is ERROR.`,
				},
				resource.Attribute{
					Name:        "status.0.state_start_time",
					Description: `The time when this state was entered.`,
				},
				resource.Attribute{
					Name:        "status.0.substate",
					Description: `Additional state information, which includes status reported by the agent.`,
				},
				resource.Attribute{
					Name:        "driver_output_resource_uri",
					Description: `A URI pointing to the location of the stdout of the job's driver program.`,
				},
				resource.Attribute{
					Name:        "driver_controls_files_uri",
					Description: `If present, the location of miscellaneous control files which may be used as part of job setup and handling. If not present, control files may be placed in the same location as driver_output_uri. ## Timeouts ` + "`" + `google_dataproc_cluster` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for submitting a job to a dataproc cluster. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting a job from a dataproc cluster.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "reference.0.cluster_uuid",
					Description: `A cluster UUID generated by the Cloud Dataproc service when the job is submitted.`,
				},
				resource.Attribute{
					Name:        "status.0.state",
					Description: `A state message specifying the overall job state.`,
				},
				resource.Attribute{
					Name:        "status.0.details",
					Description: `Optional job state details, such as an error description if the state is ERROR.`,
				},
				resource.Attribute{
					Name:        "status.0.state_start_time",
					Description: `The time when this state was entered.`,
				},
				resource.Attribute{
					Name:        "status.0.substate",
					Description: `Additional state information, which includes status reported by the agent.`,
				},
				resource.Attribute{
					Name:        "driver_output_resource_uri",
					Description: `A URI pointing to the location of the stdout of the job's driver program.`,
				},
				resource.Attribute{
					Name:        "driver_controls_files_uri",
					Description: `If present, the location of miscellaneous control files which may be used as part of job setup and handling. If not present, control files may be placed in the same location as driver_output_uri. ## Timeouts ` + "`" + `google_dataproc_cluster` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for submitting a job to a dataproc cluster. - ` + "`" + `delete` + "`" + ` - (Default ` + "`" + `10 minutes` + "`" + `) Used for deleting a job from a dataproc cluster.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_dataproc_job_iam_policy",
			Category:         "Google Dataproc Resources",
			ShortDescription: `Collection of resources to manage IAM policy for a Dataproc job.`,
			Description:      ``,
			Keywords: []string{
				"dataproc",
				"job",
				"iam",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "job",
					Description: `(Required) The name or relative resource id of the job to manage IAM policies for. For ` + "`" + `google_dataproc_job_iam_member` + "`" + ` or ` + "`" + `google_dataproc_job_iam_binding` + "`" + `:`,
				},
				resource.Attribute{
					Name:        "member/members",
					Description: `(Required) Identities that will be granted the privilege in ` + "`" + `role` + "`" + `. Each entry can have one of the following values:`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. Only one ` + "`" + `google_dataproc_job_iam_binding` + "`" + ` can be used per role. Note that custom roles must be of the format ` + "`" + `[projects|organizations]/{parent-name}/roles/{role-name}` + "`" + `. ` + "`" + `google_dataproc_job_iam_policy` + "`" + ` only:`,
				},
				resource.Attribute{
					Name:        "policy_data",
					Description: `(Required) The policy data generated by a ` + "`" + `google_iam_policy` + "`" + ` data source. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the job belongs. If it is not provided, Terraform will use the provider default.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which the job belongs. If it is not provided, Terraform will use the provider default. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the jobs's IAM policy. ## Import Job IAM resources can be imported using the project, region, job id, role and/or member. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_dataproc_job_iam_policy.editor "projects/{project}/regions/{region}/jobs/{job_id}" $ terraform import google_dataproc_job_iam_binding.editor "projects/{project}/regions/{region}/jobs/{job_id} roles/editor" $ terraform import google_dataproc_job_iam_member.editor "projects/{project}/regions/{region}/jobs/{job_id} roles/editor user:jane@example.com" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the jobs's IAM policy. ## Import Job IAM resources can be imported using the project, region, job id, role and/or member. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_dataproc_job_iam_policy.editor "projects/{project}/regions/{region}/jobs/{job_id}" $ terraform import google_dataproc_job_iam_binding.editor "projects/{project}/regions/{region}/jobs/{job_id} roles/editor" $ terraform import google_dataproc_job_iam_member.editor "projects/{project}/regions/{region}/jobs/{job_id} roles/editor user:jane@example.com" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_dns_managed_zone",
			Category:         "Google DNS Resources",
			ShortDescription: `A zone is a subtree of the DNS namespace under one administrative responsibility.`,
			Description:      ``,
			Keywords: []string{
				"dns",
				"managed",
				"zone",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dns_name",
					Description: `(Required) The DNS name of this managed zone, for instance "example.com.".`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) User assigned name for this resource. Must be unique within the project. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A textual description field. Defaults to 'Managed by Terraform'.`,
				},
				resource.Attribute{
					Name:        "dnssec_config",
					Description: `(Optional) DNSSEC configuration Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs to assign to this ManagedZone.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `(Optional) The zone's visibility: public zones are exposed to the Internet, while private zones are visible only to Virtual Private Cloud resources. Must be one of: ` + "`" + `public` + "`" + `, ` + "`" + `private` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "private_visibility_config",
					Description: `(Optional) For privately visible zones, the set of Virtual Private Cloud resources that the zone is visible from. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "forwarding_config",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) The presence for this field indicates that outbound forwarding is enabled for this zone. The value of this field contains the set of destinations to forward to. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "peering_config",
					Description: `(Optional, [Beta](https://terraform.io/docs/providers/google/provider_versions.html)) The presence of this field indicates that DNS Peering is enabled for this zone. The value of this field contains the network to peer with. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `dnssec_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Optional) Identifies what kind of resource this is`,
				},
				resource.Attribute{
					Name:        "non_existence",
					Description: `(Optional) Specifies the mechanism used to provide authenticated denial-of-existence responses.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) Specifies whether DNSSEC is enabled, and what mode it is in`,
				},
				resource.Attribute{
					Name:        "default_key_specs",
					Description: `(Optional) Specifies parameters that will be used for generating initial DnsKeys for this ManagedZone. If you provide a spec for keySigning or zoneSigning, you must also provide one for the other. Structure is documented below. The ` + "`" + `default_key_specs` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `(Optional) String mnemonic specifying the DNSSEC algorithm of this key`,
				},
				resource.Attribute{
					Name:        "key_length",
					Description: `(Optional) Length of the keys in bits`,
				},
				resource.Attribute{
					Name:        "key_type",
					Description: `(Optional) Specifies whether this is a key signing key (KSK) or a zone signing key (ZSK). Key signing keys have the Secure Entry Point flag set and, when active, will only be used to sign resource record sets of type DNSKEY. Zone signing keys do not have the Secure Entry Point flag set and will be used to sign all other types of resource record sets.`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `(Optional) Identifies what kind of resource this is The ` + "`" + `private_visibility_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "networks",
					Description: `(Optional) The list of VPC networks that can see this zone. Until the provider updates to use the Terraform 0.12 SDK in a future release, you may experience issues with this resource while updating. If you've defined a ` + "`" + `networks` + "`" + ` block and add another ` + "`" + `networks` + "`" + ` block while keeping the old block, Terraform will see an incorrect diff and apply an incorrect update to the resource. If you encounter this issue, remove all ` + "`" + `networks` + "`" + ` blocks in an update and then apply another update adding all of them back simultaneously. Structure is documented below. The ` + "`" + `networks` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "network_url",
					Description: `(Optional) The fully qualified URL of the VPC network to bind to. This should be formatted like ` + "`" + `https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}` + "`" + ` The ` + "`" + `forwarding_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "target_name_servers",
					Description: `(Optional) List of target name servers to forward to. Cloud DNS will select the best available name server if more than one target is given. Structure is documented below. The ` + "`" + `target_name_servers` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `(Optional) IPv4 address of a target name server. The ` + "`" + `peering_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "target_network",
					Description: `(Optional) The network with which to peer. Structure is documented below. The ` + "`" + `target_network` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "network_url",
					Description: `(Optional) The fully qualified URL of the VPC network to forward queries to. This should be formatted like ` + "`" + `https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}` + "`" + ` ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `Delegate your managed_zone to these virtual name servers; defined by the server ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import ManagedZone can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_dns_managed_zone.default projects/{{project}}/managedZones/{{name}} $ terraform import google_dns_managed_zone.default {{project}}/{{name}} $ terraform import google_dns_managed_zone.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name_servers",
					Description: `Delegate your managed_zone to these virtual name servers; defined by the server ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import ManagedZone can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_dns_managed_zone.default projects/{{project}}/managedZones/{{name}} $ terraform import google_dns_managed_zone.default {{project}}/{{name}} $ terraform import google_dns_managed_zone.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_dns_policy",
			Category:         "Google DNS Resources",
			ShortDescription: `A policy is a collection of DNS rules applied to one or more Virtual Private Cloud resources.`,
			Description:      ``,
			Keywords: []string{
				"dns",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) User assigned name for this policy. - - -`,
				},
				resource.Attribute{
					Name:        "alternative_name_server_config",
					Description: `(Optional) Sets an alternative name server for the associated networks. When specified, all DNS queries are forwarded to a name server that you choose. Names such as .internal are not available when an alternative name server is specified. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A textual description field. Defaults to 'Managed by Terraform'.`,
				},
				resource.Attribute{
					Name:        "enable_inbound_forwarding",
					Description: `(Optional) Allows networks bound to this policy to receive DNS queries sent by VMs or applications over VPN connections. When enabled, a virtual IP address will be allocated from each of the sub-networks that are bound to this policy.`,
				},
				resource.Attribute{
					Name:        "enable_logging",
					Description: `(Optional) Controls whether logging is enabled for the networks bound to this policy. Defaults to no logging if not set.`,
				},
				resource.Attribute{
					Name:        "networks",
					Description: `(Optional) List of network names specifying networks to which this policy is applied. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `alternative_name_server_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "target_name_servers",
					Description: `(Optional) Sets an alternative name server for the associated networks. When specified, all DNS queries are forwarded to a name server that you choose. Names such as .internal are not available when an alternative name server is specified. Structure is documented below. The ` + "`" + `target_name_servers` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `(Optional) IPv4 address to forward to. The ` + "`" + `networks` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "network_url",
					Description: `(Optional) The fully qualified URL of the VPC network to bind to. This should be formatted like ` + "`" + `https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}` + "`" + ` ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Policy can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import -provider=google-beta google_dns_policy.default projects/{{project}}/policies/{{name}} $ terraform import -provider=google-beta google_dns_policy.default {{project}}/{{name}} $ terraform import -provider=google-beta google_dns_policy.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_dns_record_set",
			Category:         "Google DNS Resources",
			ShortDescription: `Manages a set of DNS records within Google Cloud DNS.`,
			Description:      ``,
			Keywords: []string{
				"dns",
				"record",
				"set",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "managed_zone",
					Description: `(Required) The name of the zone in which this record set will reside.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The DNS name this record set will apply to.`,
				},
				resource.Attribute{
					Name:        "rrdatas",
					Description: `(Required) The string data for the records in this record set whose meaning depends on the DNS type. For TXT record, if the string data contains spaces, add surrounding ` + "`" + `\"` + "`" + ` if you don't want your string to get split on spaces. To specify a single record value longer than 255 characters such as a TXT record for DKIM, add ` + "`" + `\"\"` + "`" + ` inside the Terraform configuration string (e.g. ` + "`" + `"first255characters\"\"morecharacters"` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Required) The time-to-live of this record set (seconds).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The DNS record set type. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference Only the arguments listed above are exposed as attributes. ## Import DNS record sets can be imported using either of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_dns_record_set.frontend {{project}}/{{zone}}/{{name}}/{{type}} $ terraform import google_dns_record_set.frontend {{zone}}/{{name}}/{{type}} ` + "`" + `` + "`" + `` + "`" + ` Note: The record name must include the trailing dot at the end.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_endpoints_service",
			Category:         "Google Endpoints Resources",
			ShortDescription: `Creates and rolls out a Google Endpoints service.`,
			Description:      ``,
			Keywords: []string{
				"endpoints",
				"service",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_filestore_instance",
			Category:         "Google Filestore Resources",
			ShortDescription: `A Google Cloud Filestore instance.`,
			Description:      ``,
			Keywords: []string{
				"filestore",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The resource name of the instance.`,
				},
				resource.Attribute{
					Name:        "tier",
					Description: `(Required) The service tier of the instance.`,
				},
				resource.Attribute{
					Name:        "file_shares",
					Description: `(Required) File system shares on the instance. For this version, only a single file share is supported. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "networks",
					Description: `(Required) VPC networks to which the instance is connected. For this version, only a single network is supported. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The name of the Filestore zone of the instance. The ` + "`" + `file_shares` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the fileshare (16 characters or less)`,
				},
				resource.Attribute{
					Name:        "capacity_gb",
					Description: `(Required) File share capacity in GB. The ` + "`" + `networks` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) The name of the GCE VPC network to which the instance is connected.`,
				},
				resource.Attribute{
					Name:        "modes",
					Description: `(Required) IP versions for which the instance has IP addresses assigned.`,
				},
				resource.Attribute{
					Name:        "reserved_ip_range",
					Description: `(Optional) A /29 CIDR block that identifies the range of IP addresses reserved for this instance.`,
				},
				resource.Attribute{
					Name:        "ip_addresses",
					Description: `A list of IPv4 or IPv6 addresses. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the instance.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Resource labels to represent user-provided metadata.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `Server-specified ETag for the instance resource to prevent simultaneous updates from overwriting each other. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 6 minutes. - ` + "`" + `update` + "`" + ` - Default is 6 minutes. - ` + "`" + `delete` + "`" + ` - Default is 6 minutes. ## Import Instance can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_filestore_instance.default projects/{{project}}/locations/{{zone}}/instances/{{name}} $ terraform import google_filestore_instance.default {{project}}/{{zone}}/{{name}} $ terraform import google_filestore_instance.default {{zone}}/{{name}} $ terraform import google_filestore_instance.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `Creation timestamp in RFC3339 text format.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `Server-specified ETag for the instance resource to prevent simultaneous updates from overwriting each other. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 6 minutes. - ` + "`" + `update` + "`" + ` - Default is 6 minutes. - ` + "`" + `delete` + "`" + ` - Default is 6 minutes. ## Import Instance can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_filestore_instance.default projects/{{project}}/locations/{{zone}}/instances/{{name}} $ terraform import google_filestore_instance.default {{project}}/{{zone}}/{{name}} $ terraform import google_filestore_instance.default {{zone}}/{{name}} $ terraform import google_filestore_instance.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_firestore_index",
			Category:         "Google Firestore Resources",
			ShortDescription: `Cloud Firestore indexes enable simple and complex queries against documents in a database.`,
			Description:      ``,
			Keywords: []string{
				"firestore",
				"index",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "collection",
					Description: `(Required) The collection being indexed.`,
				},
				resource.Attribute{
					Name:        "fields",
					Description: `(Required) The fields supported by this index. The last field entry is always for the field path ` + "`" + `__name__` + "`" + `. If, on creation, ` + "`" + `__name__` + "`" + ` was not specified as the last field, it will be added automatically with the same direction as that of the last field defined. If the final field in a composite index is not directional, the ` + "`" + `__name__` + "`" + ` will be ordered ` + "`" + `"ASCENDING"` + "`" + ` (unless explicitly specified otherwise). Structure is documented below. The ` + "`" + `fields` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "field_path",
					Description: `(Optional) Name of the field.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional) Indicates that this field supports ordering by the specified order or comparing using =, <, <=, >, >=. Only one of ` + "`" + `order` + "`" + ` and ` + "`" + `arrayConfig` + "`" + ` can be specified.`,
				},
				resource.Attribute{
					Name:        "array_config",
					Description: `(Optional) Indicates that this field supports operations on arrayValues. Only one of ` + "`" + `order` + "`" + ` and ` + "`" + `arrayConfig` + "`" + ` can be specified. - - -`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Optional) The Firestore database id. Defaults to ` + "`" + `"(default)"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "query_scope",
					Description: `(Optional) The scope at which a query is run. One of ` + "`" + `"COLLECTION"` + "`" + ` or ` + "`" + `"COLLECTION_GROUP"` + "`" + `. Defaults to ` + "`" + `"COLLECTION"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A server defined name for this index. Format: ` + "`" + `projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/indexes/{{server_generated_id}}` + "`" + ` ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 10 minutes. - ` + "`" + `delete` + "`" + ` - Default is 10 minutes. ## Import Index can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_firestore_index.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `A server defined name for this index. Format: ` + "`" + `projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/indexes/{{server_generated_id}}` + "`" + ` ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 10 minutes. - ` + "`" + `delete` + "`" + ` - Default is 10 minutes. ## Import Index can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_firestore_index.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_billing_account_iam_binding",
			Category:         "Google Cloud Platform Resources",
			ShortDescription: `Allows management of a single binding with an IAM policy for a Google Cloud Platform Billing Account.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"platform",
				"billing",
				"account",
				"iam",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "billing_account_id",
					Description: `(Required) The billing account id.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Required) A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the billing account's IAM policy. ## Import IAM binding imports use space-delimited identifiers; first the resource in question and then the role. These bindings can be imported using the ` + "`" + `billing_account_id` + "`" + ` and role, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_billing_account_iam_binding.binding "your-billing-account-id roles/viewer" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the billing account's IAM policy. ## Import IAM binding imports use space-delimited identifiers; first the resource in question and then the role. These bindings can be imported using the ` + "`" + `billing_account_id` + "`" + ` and role, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_billing_account_iam_binding.binding "your-billing-account-id roles/viewer" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_billing_account_iam_member",
			Category:         "Google Cloud Platform Resources",
			ShortDescription: `Allows management of a single member for a single binding on the IAM policy for a Google Cloud Platform Billing Account.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"platform",
				"billing",
				"account",
				"iam",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "billing_account_id",
					Description: `(Required) The billing account id.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The user that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the billing account's IAM policy. ## Import IAM member imports use space-delimited identifiers; the resource in question, the role, and the account. This member resource can be imported using the ` + "`" + `billing_account_id` + "`" + `, role, and member identity, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_billing_account_iam_member.binding "your-billing-account-id roles/viewer user:foo@example.com" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the billing account's IAM policy. ## Import IAM member imports use space-delimited identifiers; the resource in question, the role, and the account. This member resource can be imported using the ` + "`" + `billing_account_id` + "`" + `, role, and member identity, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_billing_account_iam_member.binding "your-billing-account-id roles/viewer user:foo@example.com" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_billing_account_iam_policy",
			Category:         "Google Cloud Platform Resources",
			ShortDescription: `Allows management of the entire IAM policy for a Google Cloud Platform Billing Account.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"platform",
				"billing",
				"account",
				"iam",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "billing_account_id",
					Description: `(Required) The billing account id.`,
				},
				resource.Attribute{
					Name:        "policy_data",
					Description: `(Required) The ` + "`" + `google_iam_policy` + "`" + ` data source that represents the IAM policy that will be applied to the billing account. This policy overrides any existing policy applied to the billing account. ## Import ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_billing_account_iam_policy.policy billing-account-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_folder",
			Category:         "Google Cloud Platform Resources",
			ShortDescription: `Allows management of a Google Cloud Platform folder.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"platform",
				"folder",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The folder’s display name. A folder’s display name must be unique amongst its siblings, e.g. no two folders with the same parent can share the same display name. The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters.`,
				},
				resource.Attribute{
					Name:        "parent",
					Description: `(Required) The resource name of the parent Folder or Organization. Must be of the form ` + "`" + `folders/{folder_id}` + "`" + ` or ` + "`" + `organizations/{org_id}` + "`" + `. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The resource name of the Folder. Its format is folders/{folder_id}.`,
				},
				resource.Attribute{
					Name:        "lifecycle_state",
					Description: `The lifecycle state of the folder such as ` + "`" + `ACTIVE` + "`" + ` or ` + "`" + `DELETE_REQUESTED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Timestamp when the Folder was created. Assigned by the server. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z". ## Import Folders can be imported using the folder autogenerated ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` # Both syntaxes are valid $ terraform import google_folder.department1 1234567 $ terraform import google_folder.department1 folders/1234567 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The resource name of the Folder. Its format is folders/{folder_id}.`,
				},
				resource.Attribute{
					Name:        "lifecycle_state",
					Description: `The lifecycle state of the folder such as ` + "`" + `ACTIVE` + "`" + ` or ` + "`" + `DELETE_REQUESTED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Timestamp when the Folder was created. Assigned by the server. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z". ## Import Folders can be imported using the folder autogenerated ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` # Both syntaxes are valid $ terraform import google_folder.department1 1234567 $ terraform import google_folder.department1 folders/1234567 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_folder_iam_binding",
			Category:         "Google Cloud Platform Resources",
			ShortDescription: `Allows management of a single binding with an IAM policy for a Google Cloud Platform folder.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"platform",
				"folder",
				"iam",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "folder",
					Description: `(Required) The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. Only one ` + "`" + `google_folder_iam_binding` + "`" + ` can be used per role. Note that custom roles must be of the format ` + "`" + `[projects|organizations]/{parent-name}/roles/{role-name}` + "`" + `. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the folder's IAM policy. ## Import IAM binding imports use space-delimited identifiers; first the resource in question and then the role. These bindings can be imported using the ` + "`" + `folder` + "`" + ` and role, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_folder_iam_binding.viewer "folder-name roles/viewer" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the folder's IAM policy. ## Import IAM binding imports use space-delimited identifiers; first the resource in question and then the role. These bindings can be imported using the ` + "`" + `folder` + "`" + ` and role, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_folder_iam_binding.viewer "folder-name roles/viewer" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_folder_iam_member",
			Category:         "Google Cloud Platform Resources",
			ShortDescription: `Allows management of a single member for a single binding on the IAM policy for a Google Cloud Platform folder.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"platform",
				"folder",
				"iam",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "folder",
					Description: `(Required) The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The identity that will be granted the privilege in the ` + "`" + `role` + "`" + `. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding This field can have one of the following values:`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. Note that custom roles must be of the format ` + "`" + `[projects|organizations]/{parent-name}/roles/{role-name}` + "`" + `. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the folder's IAM policy. ## Import IAM member imports use space-delimited identifiers; the resource in question, the role, and the account. This member resource can be imported using the ` + "`" + `folder` + "`" + `, role, and member identity e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_folder_iam_member.my_project "folder-name roles/viewer user:foo@example.com" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the folder's IAM policy. ## Import IAM member imports use space-delimited identifiers; the resource in question, the role, and the account. This member resource can be imported using the ` + "`" + `folder` + "`" + `, role, and member identity e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_folder_iam_member.my_project "folder-name roles/viewer user:foo@example.com" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_folder_iam_policy",
			Category:         "Google Cloud Platform Resources",
			ShortDescription: `Allows management of the IAM policy for a Google Cloud Platform folders.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"platform",
				"folder",
				"iam",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "folder",
					Description: `(Required) The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.`,
				},
				resource.Attribute{
					Name:        "policy_data",
					Description: `(Required) The ` + "`" + `google_iam_policy` + "`" + ` data source that represents the IAM policy that will be applied to the folder. This policy overrides any existing policy applied to the folder. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the folder's IAM policy. ` + "`" + `etag` + "`" + ` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. ## Import A policy can be imported using the ` + "`" + `folder` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_folder_iam_policy.my-folder-policy {{folder_id}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the folder's IAM policy. ` + "`" + `etag` + "`" + ` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. ## Import A policy can be imported using the ` + "`" + `folder` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_folder_iam_policy.my-folder-policy {{folder_id}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_folder_organization_policy",
			Category:         "Google Cloud Platform Resources",
			ShortDescription: `Allows management of Organization policies for a Google Folder.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"platform",
				"folder",
				"organization",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "folder",
					Description: `(Required) The resource name of the folder to set the policy for. Its format is folders/{folder_id}.`,
				},
				resource.Attribute{
					Name:        "constraint",
					Description: `(Required) The name of the Constraint the Policy is configuring, for example, ` + "`" + `serviceuser.services` + "`" + `. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints). - - -`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Version of the Policy. Default version is 0.`,
				},
				resource.Attribute{
					Name:        "boolean_policy",
					Description: `(Optional) A boolean policy is a constraint that is either enforced or not. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "list_policy",
					Description: `(Optional) A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "restore_policy",
					Description: `(Optional) A restore policy is a constraint to restore the default policy. Structure is documented below. ~>`,
				},
				resource.Attribute{
					Name:        "enforced",
					Description: `(Required) If true, then the Policy is enforced. If false, then any configuration is acceptable. The ` + "`" + `list_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "suggested_values",
					Description: `(Optional) The Google Cloud Console will try to default to a configuration that matches the value specified in this field.`,
				},
				resource.Attribute{
					Name:        "inherit_from_parent",
					Description: `(Optional) If set to true, the values from the effective Policy of the parent resource are inherited, meaning the values set in this Policy are added to the values inherited up the hierarchy. The ` + "`" + `allow` + "`" + ` or ` + "`" + `deny` + "`" + ` blocks support:`,
				},
				resource.Attribute{
					Name:        "all",
					Description: `(Optional) The policy allows or denies all values.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional) The policy can define specific values that are allowed or denied. The ` + "`" + `restore_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `(Required) May only be set to true. If set, then the default Policy is restored. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the organization policy. ` + "`" + `etag` + "`" + ` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `(Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z". ## Import Folder organization policies can be imported using any of the follow formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_folder_organization_policy.policy folders/folder-1234:constraints/serviceuser.services $ terraform import google_folder_organization_policy.policy folder-1234:serviceuser.services ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the organization policy. ` + "`" + `etag` + "`" + ` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `(Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z". ## Import Folder organization policies can be imported using any of the follow formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_folder_organization_policy.policy folders/folder-1234:constraints/serviceuser.services $ terraform import google_folder_organization_policy.policy folder-1234:serviceuser.services ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_iap_tunnel_instance_iam_policy",
			Category:         "Google IAP Resources",
			ShortDescription: `Collection of resources to manage IAM policy for a IAP Tunnel Instance.`,
			Description:      ``,
			Keywords: []string{
				"iap",
				"tunnel",
				"instance",
				"iam",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance",
					Description: `(Required) The name of the instance.`,
				},
				resource.Attribute{
					Name:        "member/members",
					Description: `(Required) Identities that will be granted the privilege in ` + "`" + `role` + "`" + `. Each entry can have one of the following values:`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. Only one ` + "`" + `google_iap_tunnel_instance_iam_binding` + "`" + ` can be used per role. Note that custom roles must be of the format ` + "`" + `[projects|organizations]/{parent-name}/roles/{role-name}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy_data",
					Description: `(Required only by ` + "`" + `google_iap_tunnel_instance_iam_policy` + "`" + `) The policy data generated by a ` + "`" + `google_iam_policy` + "`" + ` data source.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) The zone of the instance. If unspecified, this defaults to the zone configured in the provider. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the instance's IAM policy. ## Import For all import syntaxes, the "resource in question" can take any of the following forms:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the instance's IAM policy. ## Import For all import syntaxes, the "resource in question" can take any of the following forms:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_kms_crypto_key_iam_binding",
			Category:         "Google Key Management Service Resources",
			ShortDescription: `Allows management of a single binding with an IAM policy for a Google Cloud KMS crypto key`,
			Description:      ``,
			Keywords: []string{
				"key",
				"management",
				"service",
				"kms",
				"crypto",
				"iam",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "members",
					Description: `(Required) A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. Only one ` + "`" + `google_kms_crypto_key_iam_binding` + "`" + ` can be used per role. Note that custom roles must be of the format ` + "`" + `[projects|organizations]/{parent-name}/roles/{role-name}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "crypto_key_id",
					Description: `(Required) The crypto key ID, in the form ` + "`" + `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` + "`" + ` or ` + "`" + `{location_name}/{key_ring_name}/{crypto_key_name}` + "`" + `. In the second form, the provider's project setting will be used as a fallback. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the crypto key's IAM policy. ## Import IAM binding imports use space-delimited identifiers; first the resource in question and then the role. These bindings can be imported using the ` + "`" + `crypto_key_id` + "`" + ` and role, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_kms_crypto_key_iam_binding.crypto_key "my-gcp-project/us-central1/my-key-ring/my-crypto-key roles/editor" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the crypto key's IAM policy. ## Import IAM binding imports use space-delimited identifiers; first the resource in question and then the role. These bindings can be imported using the ` + "`" + `crypto_key_id` + "`" + ` and role, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_kms_crypto_key_iam_binding.crypto_key "my-gcp-project/us-central1/my-key-ring/my-crypto-key roles/editor" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_kms_crypto_key_iam_member",
			Category:         "Google Key Management Service Resources",
			ShortDescription: `Allows management of a single member for a single binding on the IAM policy for a Google Cloud KMS crypto key.`,
			Description:      ``,
			Keywords: []string{
				"key",
				"management",
				"service",
				"kms",
				"crypto",
				"iam",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The user that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. Note that custom roles must be of the format ` + "`" + `[projects|organizations]/{parent-name}/roles/{role-name}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "crypto_key_id",
					Description: `(Required) The key ring ID, in the form ` + "`" + `{project_id}/{location_name}/{key_ring_name}/{crypto_key_name}` + "`" + ` or ` + "`" + `{location_name}/{key_ring_name}/{crypto_key_name}` + "`" + `. In the second form, the provider's project setting will be used as a fallback. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the project's IAM policy. ## Import IAM member imports use space-delimited identifiers; the resource in question, the role, and the account. This member resource can be imported using the ` + "`" + `crypto_key_id` + "`" + `, role, and member identity e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_kms_crypto_key_iam_member.member "your-project-id/location-name/key-ring-name/key-name roles/viewer user:foo@example.com" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the project's IAM policy. ## Import IAM member imports use space-delimited identifiers; the resource in question, the role, and the account. This member resource can be imported using the ` + "`" + `crypto_key_id` + "`" + `, role, and member identity e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_kms_crypto_key_iam_member.member "your-project-id/location-name/key-ring-name/key-name roles/viewer user:foo@example.com" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_kms_key_ring_iam_policy",
			Category:         "Google Key Management Service Resources",
			ShortDescription: `Collection of resources to manage IAM policy for a Google Cloud KMS key ring.`,
			Description:      ``,
			Keywords: []string{
				"key",
				"management",
				"service",
				"kms",
				"ring",
				"iam",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key_ring_id",
					Description: `(Required) The key ring ID, in the form ` + "`" + `{project_id}/{location_name}/{key_ring_name}` + "`" + ` or ` + "`" + `{location_name}/{key_ring_name}` + "`" + `. In the second form, the provider's project setting will be used as a fallback.`,
				},
				resource.Attribute{
					Name:        "member/members",
					Description: `(Required) Identities that will be granted the privilege in ` + "`" + `role` + "`" + `. Each entry can have one of the following values:`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. Only one ` + "`" + `google_kms_key_ring_iam_binding` + "`" + ` can be used per role. Note that custom roles must be of the format ` + "`" + `[projects|organizations]/{parent-name}/roles/{role-name}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy_data",
					Description: `(Required only by ` + "`" + `google_kms_key_ring_iam_policy` + "`" + `) The policy data generated by a ` + "`" + `google_iam_policy` + "`" + ` data source. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the key ring's IAM policy. ## Import IAM member imports use space-delimited identifiers; the resource in question, the role, and the account. This member resource can be imported using the ` + "`" + `key_ring_id` + "`" + `, role, and account e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_kms_key_ring_iam_member.key_ring_iam "your-project-id/location-name/key-ring-name roles/viewer user:foo@example.com" ` + "`" + `` + "`" + `` + "`" + ` IAM binding imports use space-delimited identifiers; the resource in question and the role. This binding resource can be imported using the ` + "`" + `key_ring_id` + "`" + ` and role, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_kms_key_ring_iam_binding.key_ring_iam "your-project-id/location-name/key-ring-name roles/viewer" ` + "`" + `` + "`" + `` + "`" + ` IAM policy imports use the identifier of the resource in question. This policy resource can be imported using the ` + "`" + `key_ring_id` + "`" + `, role, and account e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_kms_key_ring_iam_policy.key_ring_iam your-project-id/location-name/key-ring-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the key ring's IAM policy. ## Import IAM member imports use space-delimited identifiers; the resource in question, the role, and the account. This member resource can be imported using the ` + "`" + `key_ring_id` + "`" + `, role, and account e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_kms_key_ring_iam_member.key_ring_iam "your-project-id/location-name/key-ring-name roles/viewer user:foo@example.com" ` + "`" + `` + "`" + `` + "`" + ` IAM binding imports use space-delimited identifiers; the resource in question and the role. This binding resource can be imported using the ` + "`" + `key_ring_id` + "`" + ` and role, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_kms_key_ring_iam_binding.key_ring_iam "your-project-id/location-name/key-ring-name roles/viewer" ` + "`" + `` + "`" + `` + "`" + ` IAM policy imports use the identifier of the resource in question. This policy resource can be imported using the ` + "`" + `key_ring_id` + "`" + `, role, and account e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_kms_key_ring_iam_policy.key_ring_iam your-project-id/location-name/key-ring-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_organization_iam_binding",
			Category:         "Google Cloud Platform Resources",
			ShortDescription: `Allows management of a single binding with an IAM policy for a Google Cloud Platform Organization.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"platform",
				"organization",
				"iam",
				"binding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org_id",
					Description: `(Required) The numeric ID of the organization in which you want to create a custom role.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. Only one ` + "`" + `google_organization_iam_binding` + "`" + ` can be used per role. Note that custom roles must be of the format ` + "`" + `[projects|organizations]/{parent-name}/roles/{role-name}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Required) A list of users that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the organization's IAM policy. ## Import IAM binding imports use space-delimited identifiers; first the resource in question and then the role. These bindings can be imported using the ` + "`" + `org_id` + "`" + ` and role, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_organization_iam_binding.my_org "your-org-id roles/viewer" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the organization's IAM policy. ## Import IAM binding imports use space-delimited identifiers; first the resource in question and then the role. These bindings can be imported using the ` + "`" + `org_id` + "`" + ` and role, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_organization_iam_binding.my_org "your-org-id roles/viewer" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_organization_iam_custom_role",
			Category:         "Google Cloud Platform Resources",
			ShortDescription: `Allows management of a customized Cloud IAM organization role.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"platform",
				"organization",
				"iam",
				"custom",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "role_id",
					Description: `(Required) The role id to use for this role.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `(Required) The numeric ID of the organization in which you want to create a custom role.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Required) A human-readable title for the role.`,
				},
				resource.Attribute{
					Name:        "stage",
					Description: `(Optional) The current launch stage of the role. Defaults to ` + "`" + `GA` + "`" + `. List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-readable description for the role. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "deleted",
					Description: `(Optional) The current deleted state of the role. ## Import Customized IAM organization role can be imported using their URI, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_organization_iam_custom_role.my-custom-role organizations/123456789/roles/myCustomRole ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "deleted",
					Description: `(Optional) The current deleted state of the role. ## Import Customized IAM organization role can be imported using their URI, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_organization_iam_custom_role.my-custom-role organizations/123456789/roles/myCustomRole ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_organization_iam_member",
			Category:         "Google Cloud Platform Resources",
			ShortDescription: `Allows management of a single member for a single binding on the IAM policy for a Google Cloud Platform Organization.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"platform",
				"organization",
				"iam",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org_id",
					Description: `(Required) The numeric ID of the organization in which you want to create a custom role.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. Note that custom roles must be of the format ` + "`" + `[projects|organizations]/{parent-name}/roles/{role-name}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Required) The user that the role should apply to. For more details on format and restrictions see https://cloud.google.com/billing/reference/rest/v1/Policy#Binding ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the organization's IAM policy. ## Import IAM member imports use space-delimited identifiers; the resource in question, the role, and the account. This member resource can be imported using the ` + "`" + `org_id` + "`" + `, role, and member identity, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_organization_iam_member.my_org "your-org-id roles/viewer user:foo@example.com" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the organization's IAM policy. ## Import IAM member imports use space-delimited identifiers; the resource in question, the role, and the account. This member resource can be imported using the ` + "`" + `org_id` + "`" + `, role, and member identity, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_organization_iam_member.my_org "your-org-id roles/viewer user:foo@example.com" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_organization_iam_policy",
			Category:         "Google Cloud Platform Resources",
			ShortDescription: `Allows management of the entire IAM policy for a Google Cloud Platform Organization.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"platform",
				"organization",
				"iam",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org_id",
					Description: `(Required) The numeric ID of the organization in which you want to create a custom role.`,
				},
				resource.Attribute{
					Name:        "policy_data",
					Description: `(Required) The ` + "`" + `google_iam_policy` + "`" + ` data source that represents the IAM policy that will be applied to the organization. This policy overrides any existing policy applied to the organization. ## Import ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_organization_iam_policy.my_org your-org-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_organization_policy",
			Category:         "Google Cloud Platform Resources",
			ShortDescription: `Allows management of Organization policies for a Google Organization.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"platform",
				"organization",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org_id",
					Description: `(Required) The numeric ID of the organization to set the policy for.`,
				},
				resource.Attribute{
					Name:        "constraint",
					Description: `(Required) The name of the Constraint the Policy is configuring, for example, ` + "`" + `serviceuser.services` + "`" + `. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints). - - -`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Version of the Policy. Default version is 0.`,
				},
				resource.Attribute{
					Name:        "boolean_policy",
					Description: `(Optional) A boolean policy is a constraint that is either enforced or not. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "list_policy",
					Description: `(Optional) A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "restore_policy",
					Description: `(Optional) A restore policy is a constraint to restore the default policy. Structure is documented below. ~>`,
				},
				resource.Attribute{
					Name:        "enforced",
					Description: `(Required) If true, then the Policy is enforced. If false, then any configuration is acceptable. The ` + "`" + `list_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "suggested_values",
					Description: `(Optional) The Google Cloud Console will try to default to a configuration that matches the value specified in this field.`,
				},
				resource.Attribute{
					Name:        "inherit_from_parent",
					Description: `(Optional) If set to true, the values from the effective Policy of the parent resource are inherited, meaning the values set in this Policy are added to the values inherited up the hierarchy. The ` + "`" + `allow` + "`" + ` or ` + "`" + `deny` + "`" + ` blocks support:`,
				},
				resource.Attribute{
					Name:        "all",
					Description: `(Optional) The policy allows or denies all values.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional) The policy can define specific values that are allowed or denied. The ` + "`" + `restore_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `(Required) May only be set to true. If set, then the default Policy is restored. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the organization policy. ` + "`" + `etag` + "`" + ` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `(Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z". ## Import Organization Policies can be imported using the ` + "`" + `org_id` + "`" + ` and the ` + "`" + `constraint` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_organization_policy.services_policy 123456789:constraints/serviceuser.services`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the organization policy. ` + "`" + `etag` + "`" + ` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `(Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z". ## Import Organization Policies can be imported using the ` + "`" + `org_id` + "`" + ` and the ` + "`" + `constraint` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_organization_policy.services_policy 123456789:constraints/serviceuser.services`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_project",
			Category:         "Google Cloud Platform Resources",
			ShortDescription: `Allows management of a Google Cloud Platform project.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"platform",
				"project",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The display name of the project.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The project ID. Changing this forces a new project to be created.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `(Optional) The numeric ID of the organization this project belongs to. Changing this forces a new project to be created. Only one of ` + "`" + `org_id` + "`" + ` or ` + "`" + `folder_id` + "`" + ` may be specified. If the ` + "`" + `org_id` + "`" + ` is specified then the project is created at the top level. Changing this forces the project to be migrated to the newly specified organization.`,
				},
				resource.Attribute{
					Name:        "folder_id",
					Description: `(Optional) The numeric ID of the folder this project should be created under. Only one of ` + "`" + `org_id` + "`" + ` or ` + "`" + `folder_id` + "`" + ` may be specified. If the ` + "`" + `folder_id` + "`" + ` is specified, then the project is created under the specified folder. Changing this forces the project to be migrated to the newly specified folder.`,
				},
				resource.Attribute{
					Name:        "billing_account",
					Description: `(Optional) The alphanumeric ID of the billing account this project belongs to. The user or service account performing this operation with Terraform must have Billing Account Administrator privileges (` + "`" + `roles/billing.admin` + "`" + `) in the organization. See [Google Cloud Billing API Access Control](https://cloud.google.com/billing/v1/how-tos/access-control) for more details.`,
				},
				resource.Attribute{
					Name:        "skip_delete",
					Description: `(Optional) If true, the Terraform resource can be deleted without deleting the Project via the Google API.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs to assign to the project.`,
				},
				resource.Attribute{
					Name:        "auto_create_network",
					Description: `(Optional) Create the 'default' network automatically. Default ` + "`" + `true` + "`" + `. If set to ` + "`" + `false` + "`" + `, the default network will be deleted. Note that, for quota purposes, you will still need to have 1 network slot available to create the project successfully, even if you set ` + "`" + `auto_create_network` + "`" + ` to ` + "`" + `false` + "`" + `, since the network will exist momentarily. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "number",
					Description: `The numeric identifier of the project. ## Import Projects can be imported using the ` + "`" + `project_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_project.my_project your-project-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "number",
					Description: `The numeric identifier of the project. ## Import Projects can be imported using the ` + "`" + `project_id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_project.my_project your-project-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_project_iam_policy",
			Category:         "Google Cloud Platform Resources",
			ShortDescription: `Collection of resources to manage IAM policy for a project.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"platform",
				"project",
				"iam",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "member/members",
					Description: `(Required) Identities that will be granted the privilege in ` + "`" + `role` + "`" + `. Each entry can have one of the following values:`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. Only one ` + "`" + `google_project_iam_binding` + "`" + ` can be used per role. Note that custom roles must be of the format ` + "`" + `[projects|organizations]/{parent-name}/roles/{role-name}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy_data",
					Description: `(Required only by ` + "`" + `google_project_iam_policy` + "`" + `) The ` + "`" + `google_iam_policy` + "`" + ` data source that represents the IAM policy that will be applied to the project. The policy will be merged with any existing policy applied to the project. Changing this updates the policy. Deleting this removes all policies from the project, locking out users without organization-level access.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project ID. If not specified for ` + "`" + `google_project_iam_binding` + "`" + ` or ` + "`" + `google_project_iam_member` + "`" + `, uses the ID of the project configured with the provider. Required for ` + "`" + `google_project_iam_policy` + "`" + ` - you must explicitly set the project, and it will not be inferred from the provider. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the project's IAM policy. ## Import IAM member imports use space-delimited identifiers; the resource in question, the role, and the account. This member resource can be imported using the ` + "`" + `project_id` + "`" + `, role, and member e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_project_iam_member.my_project "your-project-id roles/viewer user:foo@example.com" ` + "`" + `` + "`" + `` + "`" + ` IAM binding imports use space-delimited identifiers; the resource in question and the role. This binding resource can be imported using the ` + "`" + `project_id` + "`" + ` and role, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import google_project_iam_binding.my_project "your-project-id roles/viewer" ` + "`" + `` + "`" + `` + "`" + ` IAM policy imports use the identifier of the resource in question. This policy resource can be imported using the ` + "`" + `project_id` + "`" + `. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_project_iam_policy.my_project your-project-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the project's IAM policy. ## Import IAM member imports use space-delimited identifiers; the resource in question, the role, and the account. This member resource can be imported using the ` + "`" + `project_id` + "`" + `, role, and member e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_project_iam_member.my_project "your-project-id roles/viewer user:foo@example.com" ` + "`" + `` + "`" + `` + "`" + ` IAM binding imports use space-delimited identifiers; the resource in question and the role. This binding resource can be imported using the ` + "`" + `project_id` + "`" + ` and role, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import google_project_iam_binding.my_project "your-project-id roles/viewer" ` + "`" + `` + "`" + `` + "`" + ` IAM policy imports use the identifier of the resource in question. This policy resource can be imported using the ` + "`" + `project_id` + "`" + `. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_project_iam_policy.my_project your-project-id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_project_iam_custom_role",
			Category:         "Google Cloud Platform Resources",
			ShortDescription: `Allows management of a customized Cloud IAM project role.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"platform",
				"project",
				"iam",
				"custom",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "role_id",
					Description: `(Required) The role id to use for this role.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Required) A human-readable title for the role.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project that the service account will be created in. Defaults to the provider project configuration.`,
				},
				resource.Attribute{
					Name:        "stage",
					Description: `(Optional) The current launch stage of the role. Defaults to ` + "`" + `GA` + "`" + `. List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-readable description for the role. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "deleted",
					Description: `(Optional) The current deleted state of the role. ## Import Customized IAM project role can be imported using their URI, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_project_iam_custom_role.my-custom-role projects/my-project/roles/myCustomRole ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "deleted",
					Description: `(Optional) The current deleted state of the role. ## Import Customized IAM project role can be imported using their URI, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_project_iam_custom_role.my-custom-role projects/my-project/roles/myCustomRole ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_project_organization_policy",
			Category:         "Google Cloud Platform Resources",
			ShortDescription: `Allows management of Organization policies for a Google Project.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"platform",
				"project",
				"organization",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Required) The project id of the project to set the policy for.`,
				},
				resource.Attribute{
					Name:        "constraint",
					Description: `(Required) The name of the Constraint the Policy is configuring, for example, ` + "`" + `serviceuser.services` + "`" + `. Check out the [complete list of available constraints](https://cloud.google.com/resource-manager/docs/organization-policy/understanding-constraints#available_constraints). - - -`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Version of the Policy. Default version is 0.`,
				},
				resource.Attribute{
					Name:        "boolean_policy",
					Description: `(Optional) A boolean policy is a constraint that is either enforced or not. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "list_policy",
					Description: `(Optional) A policy that can define specific values that are allowed or denied for the given constraint. It can also be used to allow or deny all values. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "restore_policy",
					Description: `(Optional) A restore policy is a constraint to restore the default policy. Structure is documented below. ~>`,
				},
				resource.Attribute{
					Name:        "enforced",
					Description: `(Required) If true, then the Policy is enforced. If false, then any configuration is acceptable. The ` + "`" + `list_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "suggested_values",
					Description: `(Optional) The Google Cloud Console will try to default to a configuration that matches the value specified in this field.`,
				},
				resource.Attribute{
					Name:        "inherit_from_parent",
					Description: `(Optional) If set to true, the values from the effective Policy of the parent resource are inherited, meaning the values set in this Policy are added to the values inherited up the hierarchy. The ` + "`" + `allow` + "`" + ` or ` + "`" + `deny` + "`" + ` blocks support:`,
				},
				resource.Attribute{
					Name:        "all",
					Description: `(Optional) The policy allows or denies all values.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Optional) The policy can define specific values that are allowed or denied. The ` + "`" + `restore_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `(Required) May only be set to true. If set, then the default Policy is restored. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the organization policy. ` + "`" + `etag` + "`" + ` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `(Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z". ## Import Project organization policies can be imported using any of the follow formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_project_organization_policy.policy projects/test-project:constraints/serviceuser.services $ terraform import google_project_organization_policy.policy test-project:constraints/serviceuser.services $ terraform import google_project_organization_policy.policy test-project:serviceuser.services ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the organization policy. ` + "`" + `etag` + "`" + ` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `(Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z". ## Import Project organization policies can be imported using any of the follow formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_project_organization_policy.policy projects/test-project:constraints/serviceuser.services $ terraform import google_project_organization_policy.policy test-project:constraints/serviceuser.services $ terraform import google_project_organization_policy.policy test-project:serviceuser.services ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_project_service",
			Category:         "Google Cloud Platform Resources",
			ShortDescription: `Allows management of a single API service for a Google Cloud Platform project.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"platform",
				"project",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service",
					Description: `(Required) The service to enable.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project ID. If not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "disable_dependent_services",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, services that are enabled and which depend on this service should also be disabled when this service is destroyed. If ` + "`" + `false` + "`" + ` or unset, an error will be generated if any enabled services depend on this service when destroying it.`,
				},
				resource.Attribute{
					Name:        "disable_on_destroy",
					Description: `(Optional) If true, disable the service when the terraform resource is destroyed. Defaults to true. May be useful in the event that a project is long-lived but the infrastructure running in that project changes frequently. ## Import Project services can be imported using the ` + "`" + `project_id` + "`" + ` and ` + "`" + `service` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_project_service.my_project your-project-id/iam.googleapis.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_project_services",
			Category:         "Google Cloud Platform Resources",
			ShortDescription: `Allows management of API services for a Google Cloud Platform project.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"platform",
				"project",
				"services",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project",
					Description: `(Required) The project ID. Changing this forces Terraform to attempt to disable all previously managed API services in the previous project.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Required) The list of services that are enabled. Supports update.`,
				},
				resource.Attribute{
					Name:        "disable_on_destroy",
					Description: `(Optional) Whether or not to disable APIs on project when destroyed. Defaults to true.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_service_account",
			Category:         "Google Cloud Platform Resources",
			ShortDescription: `Allows management of a Google Cloud Platform service account.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"platform",
				"service",
				"account",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_id",
					Description: `(Required) The account id that is used to generate the service account email address and a stable unique id. It is unique within a project, must be 6-30 characters long, and match the regular expression ` + "`" + `[a-z]([-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The display name for the service account. Can be updated without creating a new resource.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project that the service account will be created in. Defaults to the provider project configuration. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The e-mail address of the service account. This value should be referenced from any ` + "`" + `google_iam_policy` + "`" + ` data sources that would grant the service account privileges.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The fully-qualified name of the service account.`,
				},
				resource.Attribute{
					Name:        "unique_id",
					Description: `The unique id of the service account. ## Import Service accounts can be imported using their URI, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_service_account.my_sa projects/my-project/serviceAccounts/my-sa@my-project.iam.gserviceaccount.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "email",
					Description: `The e-mail address of the service account. This value should be referenced from any ` + "`" + `google_iam_policy` + "`" + ` data sources that would grant the service account privileges.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The fully-qualified name of the service account.`,
				},
				resource.Attribute{
					Name:        "unique_id",
					Description: `The unique id of the service account. ## Import Service accounts can be imported using their URI, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_service_account.my_sa projects/my-project/serviceAccounts/my-sa@my-project.iam.gserviceaccount.com ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_service_account_iam_policy",
			Category:         "Google Cloud Platform Resources",
			ShortDescription: `Collection of resources to manage IAM policy for a service account.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"platform",
				"service",
				"account",
				"iam",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_account_id",
					Description: `(Required) The fully-qualified name of the service account to apply policy to.`,
				},
				resource.Attribute{
					Name:        "member/members",
					Description: `(Required) Identities that will be granted the privilege in ` + "`" + `role` + "`" + `. Each entry can have one of the following values:`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. Only one ` + "`" + `google_service_account_iam_binding` + "`" + ` can be used per role. Note that custom roles must be of the format ` + "`" + `[projects|organizations]/{parent-name}/roles/{role-name}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy_data",
					Description: `(Required only by ` + "`" + `google_service_account_iam_policy` + "`" + `) The policy data generated by a ` + "`" + `google_iam_policy` + "`" + ` data source. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the service account IAM policy. ## Import Service account IAM resources can be imported using the project, service account email, role and member identity. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_service_account_iam_policy.admin-account-iam projects/{your-project-id}/serviceAccounts/{your-service-account-email} $ terraform import google_service_account_iam_binding.admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} roles/editor" $ terraform import google_service_account_iam_member.admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} roles/editor user:foo@example.com" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the service account IAM policy. ## Import Service account IAM resources can be imported using the project, service account email, role and member identity. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_service_account_iam_policy.admin-account-iam projects/{your-project-id}/serviceAccounts/{your-service-account-email} $ terraform import google_service_account_iam_binding.admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} roles/editor" $ terraform import google_service_account_iam_member.admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} roles/editor user:foo@example.com" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_service_account_key",
			Category:         "Google Cloud Platform Resources",
			ShortDescription: `Allows management of a Google Cloud Platform service account Key Pair`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"platform",
				"service",
				"account",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_account_id",
					Description: `(Required) The Service account id of the Key Pair. This can be a string in the format ` + "`" + `{ACCOUNT}` + "`" + ` or ` + "`" + `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}` + "`" + `, where ` + "`" + `{ACCOUNT}` + "`" + ` is the email address or unique id of the service account. If the ` + "`" + `{ACCOUNT}` + "`" + ` syntax is used, the project will be inferred from the account.`,
				},
				resource.Attribute{
					Name:        "key_algorithm",
					Description: `(Optional) The algorithm used to generate the key. KEY_ALG_RSA_2048 is the default algorithm. Valid values are listed at [ServiceAccountPrivateKeyType](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKeyAlgorithm) (only used on create)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name used for this key pair`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `The public key, base64 encoded`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The private key in JSON format, base64 encoded. This is what you normally get as a file when creating service account keys through the CLI or web console. This is only populated when creating a new key, and when no ` + "`" + `pgp_key` + "`" + ` is provided.`,
				},
				resource.Attribute{
					Name:        "private_key_fingerprint",
					Description: `The MD5 public key fingerprint for the encrypted private key. This is only populated when creating a new key and ` + "`" + `pgp_key` + "`" + ` is supplied`,
				},
				resource.Attribute{
					Name:        "valid_after",
					Description: `The key can be used after this timestamp. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".`,
				},
				resource.Attribute{
					Name:        "valid_before",
					Description: `The key can be used before this timestamp. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name used for this key pair`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `The public key, base64 encoded`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The private key in JSON format, base64 encoded. This is what you normally get as a file when creating service account keys through the CLI or web console. This is only populated when creating a new key, and when no ` + "`" + `pgp_key` + "`" + ` is provided.`,
				},
				resource.Attribute{
					Name:        "private_key_fingerprint",
					Description: `The MD5 public key fingerprint for the encrypted private key. This is only populated when creating a new key and ` + "`" + `pgp_key` + "`" + ` is supplied`,
				},
				resource.Attribute{
					Name:        "valid_after",
					Description: `The key can be used after this timestamp. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".`,
				},
				resource.Attribute{
					Name:        "valid_before",
					Description: `The key can be used before this timestamp. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_healthcare_dataset",
			Category:         "Google Healthcare Resources",
			ShortDescription: `A Healthcare ` + "`" + `Dataset` + "`" + ` is a toplevel logical grouping of ` + "`" + `dicomStores` + "`" + `, ` + "`" + `fhirStores` + "`" + ` and ` + "`" + `hl7V2Stores` + "`" + `.`,
			Description:      ``,
			Keywords: []string{
				"healthcare",
				"dataset",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The resource name for the Dataset.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) The location for the Dataset. - - -`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `(Optional) The default timezone used by this dataset. Must be a either a valid IANA time zone name such as "America/New_York" or empty, which defaults to UTC. This is used for parsing times in resources (e.g., HL7 messages) where no explicit timezone is specified.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The fully qualified name of this dataset ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Dataset can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import -provider=google-beta google_healthcare_dataset.default projects/{{project}}/locations/{{location}}/datasets/{{name}} $ terraform import -provider=google-beta google_healthcare_dataset.default {{project}}/{{location}}/{{name}} $ terraform import -provider=google-beta google_healthcare_dataset.default {{location}}/{{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "self_link",
					Description: `The fully qualified name of this dataset ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Dataset can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import -provider=google-beta google_healthcare_dataset.default projects/{{project}}/locations/{{location}}/datasets/{{name}} $ terraform import -provider=google-beta google_healthcare_dataset.default {{project}}/{{location}}/{{name}} $ terraform import -provider=google-beta google_healthcare_dataset.default {{location}}/{{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_healthcare_dataset_iam_policy",
			Category:         "Google Healthcare Resources",
			ShortDescription: `Collection of resources to manage IAM policy for a Google Cloud Healthcare dataset.`,
			Description:      ``,
			Keywords: []string{
				"healthcare",
				"dataset",
				"iam",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dataset_id",
					Description: `(Required) The dataset ID, in the form ` + "`" + `{project_id}/{location_name}/{dataset_name}` + "`" + ` or ` + "`" + `{location_name}/{dataset_name}` + "`" + `. In the second form, the provider's project setting will be used as a fallback.`,
				},
				resource.Attribute{
					Name:        "member/members",
					Description: `(Required) Identities that will be granted the privilege in ` + "`" + `role` + "`" + `. Each entry can have one of the following values:`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. Only one ` + "`" + `google_healthcare_dataset_iam_binding` + "`" + ` can be used per role. Note that custom roles must be of the format ` + "`" + `[projects|organizations]/{parent-name}/roles/{role-name}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy_data",
					Description: `(Required only by ` + "`" + `google_healthcare_dataset_iam_policy` + "`" + `) The policy data generated by a ` + "`" + `google_iam_policy` + "`" + ` data source. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the dataset's IAM policy. ## Import IAM member imports use space-delimited identifiers; the resource in question, the role, and the account. This member resource can be imported using the ` + "`" + `dataset_id` + "`" + `, role, and account e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_healthcare_dataset_iam_member.dataset_iam "your-project-id/location-name/dataset-name roles/viewer user:foo@example.com" ` + "`" + `` + "`" + `` + "`" + ` IAM binding imports use space-delimited identifiers; the resource in question and the role. This binding resource can be imported using the ` + "`" + `dataset_id` + "`" + ` and role, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_healthcare_dataset_iam_binding.dataset_iam "your-project-id/location-name/dataset-name roles/viewer" ` + "`" + `` + "`" + `` + "`" + ` IAM policy imports use the identifier of the resource in question. This policy resource can be imported using the ` + "`" + `dataset_id` + "`" + `, role, and account e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_healthcare_dataset_iam_policy.dataset_iam your-project-id/location-name/dataset-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the dataset's IAM policy. ## Import IAM member imports use space-delimited identifiers; the resource in question, the role, and the account. This member resource can be imported using the ` + "`" + `dataset_id` + "`" + `, role, and account e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_healthcare_dataset_iam_member.dataset_iam "your-project-id/location-name/dataset-name roles/viewer user:foo@example.com" ` + "`" + `` + "`" + `` + "`" + ` IAM binding imports use space-delimited identifiers; the resource in question and the role. This binding resource can be imported using the ` + "`" + `dataset_id` + "`" + ` and role, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_healthcare_dataset_iam_binding.dataset_iam "your-project-id/location-name/dataset-name roles/viewer" ` + "`" + `` + "`" + `` + "`" + ` IAM policy imports use the identifier of the resource in question. This policy resource can be imported using the ` + "`" + `dataset_id` + "`" + `, role, and account e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_healthcare_dataset_iam_policy.dataset_iam your-project-id/location-name/dataset-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_healthcare_dicom_store",
			Category:         "Google Healthcare Resources",
			ShortDescription: `A DicomStore is a datastore inside a Healthcare dataset that conforms to the DICOM (https://www.`,
			Description:      ``,
			Keywords: []string{
				"healthcare",
				"dicom",
				"store",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The resource name for the DicomStore.`,
				},
				resource.Attribute{
					Name:        "dataset",
					Description: `(Required) Identifies the dataset addressed by this request. Must be in the format 'projects/{project}/locations/{location}/datasets/{dataset}' - - -`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) User-supplied key-value pairs used to organize DICOM stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store. An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.`,
				},
				resource.Attribute{
					Name:        "notification_config",
					Description: `(Optional) A nested object resource Structure is documented below. The ` + "`" + `notification_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "pubsub_topic",
					Description: `(Required) The Cloud Pub/Sub topic that notifications of changes are published on. Supplied by the client. PubsubMessage.Data will contain the resource name. PubsubMessage.MessageId is the ID of this message. It is guaranteed to be unique within the topic. PubsubMessage.PublishTime is the time at which the message was published. Notifications are only sent if the topic is non-empty. Topic names must be scoped to a project. cloud-healthcare@system.gserviceaccount.com must have publisher permissions on the given Cloud Pub/Sub topic. Not having adequate permissions will cause the calls that send notifications to fail. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The fully qualified name of this dataset ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import DicomStore can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import -provider=google-beta google_healthcare_dicom_store.default {{dataset}}/dicomStores/{{name}} $ terraform import -provider=google-beta google_healthcare_dicom_store.default {{dataset}}/{{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "self_link",
					Description: `The fully qualified name of this dataset ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import DicomStore can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import -provider=google-beta google_healthcare_dicom_store.default {{dataset}}/dicomStores/{{name}} $ terraform import -provider=google-beta google_healthcare_dicom_store.default {{dataset}}/{{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_healthcare_dicom_store_iam_policy",
			Category:         "Google Healthcare Resources",
			ShortDescription: `Collection of resources to manage IAM policy for a Google Cloud Healthcare DICOM store.`,
			Description:      ``,
			Keywords: []string{
				"healthcare",
				"dicom",
				"store",
				"iam",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dicom_store_id",
					Description: `(Required) The DICOM store ID, in the form ` + "`" + `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` + "`" + ` or ` + "`" + `{location_name}/{dataset_name}/{dicom_store_name}` + "`" + `. In the second form, the provider's project setting will be used as a fallback.`,
				},
				resource.Attribute{
					Name:        "member/members",
					Description: `(Required) Identities that will be granted the privilege in ` + "`" + `role` + "`" + `. Each entry can have one of the following values:`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. Only one ` + "`" + `google_healthcare_dicom_store_iam_binding` + "`" + ` can be used per role. Note that custom roles must be of the format ` + "`" + `[projects|organizations]/{parent-name}/roles/{role-name}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy_data",
					Description: `(Required only by ` + "`" + `google_healthcare_dicom_store_iam_policy` + "`" + `) The policy data generated by a ` + "`" + `google_iam_policy` + "`" + ` data source. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the DICOM store's IAM policy. ## Import IAM member imports use space-delimited identifiers; the resource in question, the role, and the account. This member resource can be imported using the ` + "`" + `dicom_store_id` + "`" + `, role, and account e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_healthcare_dicom_store_iam_member.dicom_store_iam "your-project-id/location-name/dataset-name/dicom-store-name roles/viewer user:foo@example.com" ` + "`" + `` + "`" + `` + "`" + ` IAM binding imports use space-delimited identifiers; the resource in question and the role. This binding resource can be imported using the ` + "`" + `dicom_store_id` + "`" + ` and role, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_healthcare_dicom_store_iam_binding.dicom_store_iam "your-project-id/location-name/dataset-name/dicom-store-name roles/viewer" ` + "`" + `` + "`" + `` + "`" + ` IAM policy imports use the identifier of the resource in question. This policy resource can be imported using the ` + "`" + `dicom_store_id` + "`" + `, role, and account e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_healthcare_dicom_store_iam_policy.dicom_store_iam your-project-id/location-name/dataset-name/dicom-store-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the DICOM store's IAM policy. ## Import IAM member imports use space-delimited identifiers; the resource in question, the role, and the account. This member resource can be imported using the ` + "`" + `dicom_store_id` + "`" + `, role, and account e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_healthcare_dicom_store_iam_member.dicom_store_iam "your-project-id/location-name/dataset-name/dicom-store-name roles/viewer user:foo@example.com" ` + "`" + `` + "`" + `` + "`" + ` IAM binding imports use space-delimited identifiers; the resource in question and the role. This binding resource can be imported using the ` + "`" + `dicom_store_id` + "`" + ` and role, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_healthcare_dicom_store_iam_binding.dicom_store_iam "your-project-id/location-name/dataset-name/dicom-store-name roles/viewer" ` + "`" + `` + "`" + `` + "`" + ` IAM policy imports use the identifier of the resource in question. This policy resource can be imported using the ` + "`" + `dicom_store_id` + "`" + `, role, and account e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_healthcare_dicom_store_iam_policy.dicom_store_iam your-project-id/location-name/dataset-name/dicom-store-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_healthcare_fhir_store",
			Category:         "Google Healthcare Resources",
			ShortDescription: `A FhirStore is a datastore inside a Healthcare dataset that conforms to the FHIR (https://www.`,
			Description:      ``,
			Keywords: []string{
				"healthcare",
				"fhir",
				"store",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The resource name for the FhirStore.`,
				},
				resource.Attribute{
					Name:        "dataset",
					Description: `(Required) Identifies the dataset addressed by this request. Must be in the format 'projects/{project}/locations/{location}/datasets/{dataset}' - - -`,
				},
				resource.Attribute{
					Name:        "enable_update_create",
					Description: `(Optional) Whether this FHIR store has the updateCreate capability. This determines if the client can use an Update operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through the Create operation and attempts to Update a non-existent resource will return errors. Please treat the audit logs with appropriate levels of care if client-specified resource IDs contain sensitive data such as patient identifiers, those IDs will be part of the FHIR resource path recorded in Cloud audit logs and Cloud Pub/Sub notifications.`,
				},
				resource.Attribute{
					Name:        "disable_referential_integrity",
					Description: `(Optional) Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store creation. The default value is false, meaning that the API will enforce referential integrity and fail the requests that will result in inconsistent state in the FHIR store. When this field is set to true, the API will skip referential integrity check. Consequently, operations that rely on references, such as Patient.get$everything, will not return all the results if broken references exist.`,
				},
				resource.Attribute{
					Name:        "disable_resource_versioning",
					Description: `(Optional) Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation of FHIR store. If set to false, which is the default behavior, all write operations will cause historical versions to be recorded automatically. The historical versions can be fetched through the history APIs, but cannot be updated. If set to true, no historical versions will be kept. The server will send back errors for attempts to read the historical versions.`,
				},
				resource.Attribute{
					Name:        "enable_history_import",
					Description: `(Optional) Whether to allow the bulk import API to accept history bundles and directly insert historical resource versions into the FHIR store. Importing resource histories creates resource interactions that appear to have occurred in the past, which clients may not want to allow. If set to false, history bundles within an import will fail with an error.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) User-supplied key-value pairs used to organize FHIR stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store. An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.`,
				},
				resource.Attribute{
					Name:        "notification_config",
					Description: `(Optional) A nested object resource Structure is documented below. The ` + "`" + `notification_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "pubsub_topic",
					Description: `(Required) The Cloud Pub/Sub topic that notifications of changes are published on. Supplied by the client. PubsubMessage.Data will contain the resource name. PubsubMessage.MessageId is the ID of this message. It is guaranteed to be unique within the topic. PubsubMessage.PublishTime is the time at which the message was published. Notifications are only sent if the topic is non-empty. Topic names must be scoped to a project. cloud-healthcare@system.gserviceaccount.com must have publisher permissions on the given Cloud Pub/Sub topic. Not having adequate permissions will cause the calls that send notifications to fail. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The fully qualified name of this dataset ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import FhirStore can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import -provider=google-beta google_healthcare_fhir_store.default {{dataset}}/fhirStores/{{name}} $ terraform import -provider=google-beta google_healthcare_fhir_store.default {{dataset}}/{{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "self_link",
					Description: `The fully qualified name of this dataset ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import FhirStore can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import -provider=google-beta google_healthcare_fhir_store.default {{dataset}}/fhirStores/{{name}} $ terraform import -provider=google-beta google_healthcare_fhir_store.default {{dataset}}/{{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_healthcare_fhir_store_iam_policy",
			Category:         "Google Healthcare Resources",
			ShortDescription: `Collection of resources to manage IAM policy for a Google Cloud Healthcare FHIR store.`,
			Description:      ``,
			Keywords: []string{
				"healthcare",
				"fhir",
				"store",
				"iam",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "fhir_store_id",
					Description: `(Required) The FHIR store ID, in the form ` + "`" + `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` + "`" + ` or ` + "`" + `{location_name}/{dataset_name}/{fhir_store_name}` + "`" + `. In the second form, the provider's project setting will be used as a fallback.`,
				},
				resource.Attribute{
					Name:        "member/members",
					Description: `(Required) Identities that will be granted the privilege in ` + "`" + `role` + "`" + `. Each entry can have one of the following values:`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. Only one ` + "`" + `google_healthcare_fhir_store_iam_binding` + "`" + ` can be used per role. Note that custom roles must be of the format ` + "`" + `[projects|organizations]/{parent-name}/roles/{role-name}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy_data",
					Description: `(Required only by ` + "`" + `google_healthcare_fhir_store_iam_policy` + "`" + `) The policy data generated by a ` + "`" + `google_iam_policy` + "`" + ` data source. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the FHIR store's IAM policy. ## Import IAM member imports use space-delimited identifiers; the resource in question, the role, and the account. This member resource can be imported using the ` + "`" + `fhir_store_id` + "`" + `, role, and account e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_healthcare_fhir_store_iam_member.fhir_store_iam "your-project-id/location-name/dataset-name/fhir-store-name roles/viewer user:foo@example.com" ` + "`" + `` + "`" + `` + "`" + ` IAM binding imports use space-delimited identifiers; the resource in question and the role. This binding resource can be imported using the ` + "`" + `fhir_store_id` + "`" + ` and role, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_healthcare_fhir_store_iam_binding.fhir_store_iam "your-project-id/location-name/dataset-name/fhir-store-name roles/viewer" ` + "`" + `` + "`" + `` + "`" + ` IAM policy imports use the identifier of the resource in question. This policy resource can be imported using the ` + "`" + `fhir_store_id` + "`" + `, role, and account e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_healthcare_fhir_store_iam_policy.fhir_store_iam your-project-id/location-name/dataset-name/fhir-store-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the FHIR store's IAM policy. ## Import IAM member imports use space-delimited identifiers; the resource in question, the role, and the account. This member resource can be imported using the ` + "`" + `fhir_store_id` + "`" + `, role, and account e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_healthcare_fhir_store_iam_member.fhir_store_iam "your-project-id/location-name/dataset-name/fhir-store-name roles/viewer user:foo@example.com" ` + "`" + `` + "`" + `` + "`" + ` IAM binding imports use space-delimited identifiers; the resource in question and the role. This binding resource can be imported using the ` + "`" + `fhir_store_id` + "`" + ` and role, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_healthcare_fhir_store_iam_binding.fhir_store_iam "your-project-id/location-name/dataset-name/fhir-store-name roles/viewer" ` + "`" + `` + "`" + `` + "`" + ` IAM policy imports use the identifier of the resource in question. This policy resource can be imported using the ` + "`" + `fhir_store_id` + "`" + `, role, and account e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_healthcare_fhir_store_iam_policy.fhir_store_iam your-project-id/location-name/dataset-name/fhir-store-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_healthcare_hl7_v2_store",
			Category:         "Google Healthcare Resources",
			ShortDescription: `A Hl7V2Store is a datastore inside a Healthcare dataset that conforms to the FHIR (https://www.`,
			Description:      ``,
			Keywords: []string{
				"healthcare",
				"hl7",
				"v2",
				"store",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The resource name for the Hl7V2Store.`,
				},
				resource.Attribute{
					Name:        "dataset",
					Description: `(Required) Identifies the dataset addressed by this request. Must be in the format 'projects/{project}/locations/{location}/datasets/{dataset}' - - -`,
				},
				resource.Attribute{
					Name:        "parser_config",
					Description: `(Optional) A nested object resource Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) User-supplied key-value pairs used to organize HL7v2 stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store. An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.`,
				},
				resource.Attribute{
					Name:        "notification_config",
					Description: `(Optional) A nested object resource Structure is documented below. The ` + "`" + `parser_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "allow_null_header",
					Description: `(Optional) Determines whether messages with no header are allowed.`,
				},
				resource.Attribute{
					Name:        "segment_terminator",
					Description: `(Optional) Byte(s) to be used as the segment terminator. If this is unset, '\r' will be used as segment terminator. A base64-encoded string. The ` + "`" + `notification_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "pubsub_topic",
					Description: `(Required) The Cloud Pub/Sub topic that notifications of changes are published on. Supplied by the client. PubsubMessage.Data will contain the resource name. PubsubMessage.MessageId is the ID of this message. It is guaranteed to be unique within the topic. PubsubMessage.PublishTime is the time at which the message was published. Notifications are only sent if the topic is non-empty. Topic names must be scoped to a project. cloud-healthcare@system.gserviceaccount.com must have publisher permissions on the given Cloud Pub/Sub topic. Not having adequate permissions will cause the calls that send notifications to fail. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The fully qualified name of this dataset ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Hl7V2Store can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import -provider=google-beta google_healthcare_hl7_v2_store.default {{dataset}}/hl7V2Stores/{{name}} $ terraform import -provider=google-beta google_healthcare_hl7_v2_store.default {{dataset}}/{{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "self_link",
					Description: `The fully qualified name of this dataset ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Hl7V2Store can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import -provider=google-beta google_healthcare_hl7_v2_store.default {{dataset}}/hl7V2Stores/{{name}} $ terraform import -provider=google-beta google_healthcare_hl7_v2_store.default {{dataset}}/{{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_healthcare_hl7_v2_store_iam_policy",
			Category:         "Google Healthcare Resources",
			ShortDescription: `Collection of resources to manage IAM policy for a Google Cloud Healthcare HL7v2 store.`,
			Description:      ``,
			Keywords: []string{
				"healthcare",
				"hl7",
				"v2",
				"store",
				"iam",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "hl7_v2_store_id",
					Description: `(Required) The HL7v2 store ID, in the form ` + "`" + `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` + "`" + ` or ` + "`" + `{location_name}/{dataset_name}/{hl7_v2_store_name}` + "`" + `. In the second form, the provider's project setting will be used as a fallback.`,
				},
				resource.Attribute{
					Name:        "member/members",
					Description: `(Required) Identities that will be granted the privilege in ` + "`" + `role` + "`" + `. Each entry can have one of the following values:`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. Only one ` + "`" + `google_healthcare_hl7_v2_store_iam_binding` + "`" + ` can be used per role. Note that custom roles must be of the format ` + "`" + `[projects|organizations]/{parent-name}/roles/{role-name}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy_data",
					Description: `(Required only by ` + "`" + `google_healthcare_hl7_v2_store_iam_policy` + "`" + `) The policy data generated by a ` + "`" + `google_iam_policy` + "`" + ` data source. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the HL7v2 store's IAM policy. ## Import IAM member imports use space-delimited identifiers; the resource in question, the role, and the account. This member resource can be imported using the ` + "`" + `hl7_v2_store_id` + "`" + `, role, and account e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_healthcare_hl7_v2_store_iam_member.hl7_v2_store_iam "your-project-id/location-name/dataset-name/hl7-v2-store-name roles/viewer user:foo@example.com" ` + "`" + `` + "`" + `` + "`" + ` IAM binding imports use space-delimited identifiers; the resource in question and the role. This binding resource can be imported using the ` + "`" + `hl7_v2_store_id` + "`" + ` and role, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_healthcare_hl7_v2_store_iam_binding.hl7_v2_store_iam "your-project-id/location-name/dataset-name/hl7-v2-store-name roles/viewer" ` + "`" + `` + "`" + `` + "`" + ` IAM policy imports use the identifier of the resource in question. This policy resource can be imported using the ` + "`" + `hl7_v2_store_id` + "`" + `, role, and account e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_healthcare_hl7_v2_store_iam_policy.hl7_v2_store_iam your-project-id/location-name/dataset-name/hl7-v2-store-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the HL7v2 store's IAM policy. ## Import IAM member imports use space-delimited identifiers; the resource in question, the role, and the account. This member resource can be imported using the ` + "`" + `hl7_v2_store_id` + "`" + `, role, and account e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_healthcare_hl7_v2_store_iam_member.hl7_v2_store_iam "your-project-id/location-name/dataset-name/hl7-v2-store-name roles/viewer user:foo@example.com" ` + "`" + `` + "`" + `` + "`" + ` IAM binding imports use space-delimited identifiers; the resource in question and the role. This binding resource can be imported using the ` + "`" + `hl7_v2_store_id` + "`" + ` and role, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_healthcare_hl7_v2_store_iam_binding.hl7_v2_store_iam "your-project-id/location-name/dataset-name/hl7-v2-store-name roles/viewer" ` + "`" + `` + "`" + `` + "`" + ` IAM policy imports use the identifier of the resource in question. This policy resource can be imported using the ` + "`" + `hl7_v2_store_id` + "`" + `, role, and account e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_healthcare_hl7_v2_store_iam_policy.hl7_v2_store_iam your-project-id/location-name/dataset-name/hl7-v2-store-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_kms_crypto_key",
			Category:         "Google Key Management Service Resources",
			ShortDescription: `A ` + "`" + `CryptoKey` + "`" + ` represents a logical key that can be used for cryptographic operations.`,
			Description:      ``,
			Keywords: []string{
				"key",
				"management",
				"service",
				"kms",
				"crypto",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The resource name for the CryptoKey.`,
				},
				resource.Attribute{
					Name:        "key_ring",
					Description: `(Required) The KeyRing that this key belongs to. Format: ` + "`" + `'projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}'` + "`" + `. - - -`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Labels with user-defined metadata to apply to this resource.`,
				},
				resource.Attribute{
					Name:        "purpose",
					Description: `(Optional) The immutable purpose of this CryptoKey. See the [purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose) for possible inputs.`,
				},
				resource.Attribute{
					Name:        "rotation_period",
					Description: `(Optional) Every time this period passes, generate a new CryptoKeyVersion and set it as the primary. The first rotation will take place after the specified period. The rotation period has the format of a decimal number with up to 9 fractional digits, followed by the letter ` + "`" + `s` + "`" + ` (seconds). It must be greater than a day (ie, 86400).`,
				},
				resource.Attribute{
					Name:        "version_template",
					Description: `(Optional) A template describing settings for new crypto key versions. Structure is documented below. The ` + "`" + `version_template` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `(Required) The algorithm to use when creating a version based on this template. See the [algorithm reference](https://cloud.google.com/kms/docs/reference/rest/v1/CryptoKeyVersionAlgorithm) for possible inputs.`,
				},
				resource.Attribute{
					Name:        "protection_level",
					Description: `(Optional) The protection level to use when creating a version based on this template. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_kms_key_ring",
			Category:         "Google Key Management Service Resources",
			ShortDescription: `A ` + "`" + `KeyRing` + "`" + ` is a toplevel logical grouping of ` + "`" + `CryptoKeys` + "`" + `.`,
			Description:      ``,
			Keywords: []string{
				"key",
				"management",
				"service",
				"kms",
				"ring",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The resource name for the KeyRing.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) The location for the KeyRing. A full list of valid locations can be found by running ` + "`" + `gcloud kms locations list` + "`" + `. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_logging_billing_account_exclusion",
			Category:         "Google Stackdriver Logging Resources",
			ShortDescription: `Manages a billing_account-level logging exclusion.`,
			Description:      ``,
			Keywords: []string{
				"stackdriver",
				"logging",
				"billing",
				"account",
				"exclusion",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "billing_account",
					Description: `(Required) The billing account to create the exclusion for.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the logging exclusion.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-readable description.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Whether this exclusion rule should be disabled or not. This defaults to false.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) The filter to apply when excluding logs. Only log entries that match the filter are excluded. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to write a filter. ## Import Billing account logging exclusions can be imported using their URI, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_logging_billing_account_exclusion.my_exclusion billingAccounts/my-billing_account/exclusions/my-exclusion ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_logging_billing_account_sink",
			Category:         "Google Stackdriver Logging Resources",
			ShortDescription: `Manages a billing account logging sink.`,
			Description:      ``,
			Keywords: []string{
				"stackdriver",
				"logging",
				"billing",
				"account",
				"sink",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the logging sink.`,
				},
				resource.Attribute{
					Name:        "billing_account",
					Description: `(Required) The billing account exported to the sink.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Required) The destination of the sink (or, in other words, where logs are written to). Can be a Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples: ` + "`" + `` + "`" + `` + "`" + ` "storage.googleapis.com/[GCS_BUCKET]" "bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]" "pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]" ` + "`" + `` + "`" + `` + "`" + ` The writer associated with the sink must have access to write to the above resource.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) The filter to apply when exporting logs. Only log entries that match the filter are exported. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to write a filter. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "writer_identity",
					Description: `The identity associated with this sink. This identity must be granted write access to the configured ` + "`" + `destination` + "`" + `. ## Import Billing account logging sinks can be imported using this format: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_logging_billing_account_sink.my_sink billingAccounts/{{billing_account_id}}/sinks/{{sink_id}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "writer_identity",
					Description: `The identity associated with this sink. This identity must be granted write access to the configured ` + "`" + `destination` + "`" + `. ## Import Billing account logging sinks can be imported using this format: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_logging_billing_account_sink.my_sink billingAccounts/{{billing_account_id}}/sinks/{{sink_id}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_logging_folder_exclusion",
			Category:         "Google Stackdriver Logging Resources",
			ShortDescription: `Manages a folder-level logging exclusion.`,
			Description:      ``,
			Keywords: []string{
				"stackdriver",
				"logging",
				"folder",
				"exclusion",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "folder",
					Description: `(Required) The folder to be exported to the sink. Note that either [FOLDER_ID] or "folders/[FOLDER_ID]" is accepted.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the logging exclusion.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-readable description.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Whether this exclusion rule should be disabled or not. This defaults to false.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) The filter to apply when excluding logs. Only log entries that match the filter are excluded. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to write a filter. ## Import Folder-level logging exclusions can be imported using their URI, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_logging_folder_exclusion.my_exclusion folders/my-folder/exclusions/my-exclusion ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_logging_folder_sink",
			Category:         "Google Stackdriver Logging Resources",
			ShortDescription: `Manages a folder-level logging sink.`,
			Description:      ``,
			Keywords: []string{
				"stackdriver",
				"logging",
				"folder",
				"sink",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the logging sink.`,
				},
				resource.Attribute{
					Name:        "folder",
					Description: `(Required) The folder to be exported to the sink. Note that either [FOLDER_ID] or "folders/[FOLDER_ID]" is accepted.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Required) The destination of the sink (or, in other words, where logs are written to). Can be a Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples: ` + "`" + `` + "`" + `` + "`" + ` "storage.googleapis.com/[GCS_BUCKET]" "bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]" "pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]" ` + "`" + `` + "`" + `` + "`" + ` The writer associated with the sink must have access to write to the above resource.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) The filter to apply when exporting logs. Only log entries that match the filter are exported. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to write a filter.`,
				},
				resource.Attribute{
					Name:        "include_children",
					Description: `(Optional) Whether or not to include children folders in the sink export. If true, logs associated with child projects are also exported; otherwise only logs relating to the provided folder are included. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "writer_identity",
					Description: `The identity associated with this sink. This identity must be granted write access to the configured ` + "`" + `destination` + "`" + `. ## Import Folder-level logging sinks can be imported using this format: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_logging_folder_sink.my_sink folders/{{folder_id}}/sinks/{{sink_id}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "writer_identity",
					Description: `The identity associated with this sink. This identity must be granted write access to the configured ` + "`" + `destination` + "`" + `. ## Import Folder-level logging sinks can be imported using this format: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_logging_folder_sink.my_sink folders/{{folder_id}}/sinks/{{sink_id}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_logging_metric",
			Category:         "Google Stackdriver Logging Resources",
			ShortDescription: `Logs-based metric can also be used to extract values from logs and create a a distribution of the values.`,
			Description:      ``,
			Keywords: []string{
				"stackdriver",
				"logging",
				"metric",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The client-assigned metric identifier. Examples - "error_count", "nginx/requests". Metric identifiers are limited to 100 characters and can include only the following characters A-Z, a-z, 0-9, and the special characters _-.,+!`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-filters) which is used to match log entries.`,
				},
				resource.Attribute{
					Name:        "metric_descriptor",
					Description: `(Required) The metric descriptor associated with the logs-based metric. Structure is documented below. The ` + "`" + `metric_descriptor` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "value_type",
					Description: `(Required) Whether the measurement is an integer, a floating-point number, etc. Some combinations of metricKind and valueType might not be supported. For counter metrics, set this to INT64.`,
				},
				resource.Attribute{
					Name:        "metric_kind",
					Description: `(Required) Whether the metric records instantaneous values, changes to a value, etc. Some combinations of metricKind and valueType might not be supported. For counter metrics, set this to DELTA.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) The set of labels that can be used to describe a specific instance of this metric type. For example, the appengine.googleapis.com/http/server/response_latencies metric type has a label for the HTTP response code, response_code, so you can look at latencies for successful responses or just for responses that failed. Structure is documented below. The ` + "`" + `labels` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The label key.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-readable description for the label.`,
				},
				resource.Attribute{
					Name:        "value_type",
					Description: `(Optional) The type of data that can be assigned to the label. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of this metric, which is used in documentation. The maximum length of the description is 8000 characters.`,
				},
				resource.Attribute{
					Name:        "label_extractors",
					Description: `(Optional) A map from a label key string to an extractor expression which is used to extract data from a log entry field and assign as the label value. Each label key specified in the LabelDescriptor must have an associated extractor expression in this map. The syntax of the extractor expression is the same as for the valueExtractor field.`,
				},
				resource.Attribute{
					Name:        "value_extractor",
					Description: `(Optional) A valueExtractor is required when using a distribution logs-based metric to extract the values to record from a log entry. Two functions are supported for value extraction - EXTRACT(field) or REGEXP_EXTRACT(field, regex). The argument are 1. field - The name of the log entry field from which the value is to be extracted. 2. regex - A regular expression using the Google RE2 syntax (https://github.com/google/re2/wiki/Syntax) with a single capture group to extract data from the specified log entry field. The value of the field is converted to a string before applying the regex. It is an error to specify a regex that does not include exactly one capture group.`,
				},
				resource.Attribute{
					Name:        "bucket_options",
					Description: `(Optional) The bucketOptions are required when the logs-based metric is using a DISTRIBUTION value type and it describes the bucket boundaries used to create a histogram of the extracted values. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `bucket_options` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "linear_buckets",
					Description: `(Optional) Specifies a linear sequence of buckets that all have the same width (except overflow and underflow). Each bucket represents a constant absolute uncertainty on the specific value in the bucket. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "exponential_buckets",
					Description: `(Optional) Specifies an exponential sequence of buckets that have a width that is proportional to the value of the lower bound. Each bucket represents a constant relative uncertainty on a specific value in the bucket. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "explicit",
					Description: `(Optional) Specifies a set of buckets with arbitrary widths. Structure is documented below. The ` + "`" + `linear_buckets` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "num_finite_buckets",
					Description: `(Optional) Must be greater than 0.`,
				},
				resource.Attribute{
					Name:        "width",
					Description: `(Optional) Must be greater than 0.`,
				},
				resource.Attribute{
					Name:        "offset",
					Description: `(Optional) Lower bound of the first bucket. The ` + "`" + `exponential_buckets` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "num_finite_buckets",
					Description: `(Optional) Must be greater than 0.`,
				},
				resource.Attribute{
					Name:        "growth_factor",
					Description: `(Optional) Must be greater than 1.`,
				},
				resource.Attribute{
					Name:        "scale",
					Description: `(Optional) Must be greater than 0. The ` + "`" + `explicit` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "bounds",
					Description: `(Optional) The values must be monotonically increasing. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Metric can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_logging_metric.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_logging_organization_exclusion",
			Category:         "Google Stackdriver Logging Resources",
			ShortDescription: `Manages a organization-level logging exclusion.`,
			Description:      ``,
			Keywords: []string{
				"stackdriver",
				"logging",
				"organization",
				"exclusion",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the logging exclusion.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `(Required) The organization to create the exclusion in.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-readable description.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Whether this exclusion rule should be disabled or not. This defaults to false.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) The filter to apply when excluding logs. Only log entries that match the filter are excluded. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to write a filter. ## Import Organization-level logging exclusions can be imported using their URI, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_logging_organization_exclusion.my_exclusion organizations/my-organization/exclusions/my-exclusion ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_logging_organization_sink",
			Category:         "Google Stackdriver Logging Resources",
			ShortDescription: `Manages a organization-level logging sink.`,
			Description:      ``,
			Keywords: []string{
				"stackdriver",
				"logging",
				"organization",
				"sink",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the logging sink.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `(Required) The numeric ID of the organization to be exported to the sink.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Required) The destination of the sink (or, in other words, where logs are written to). Can be a Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples: ` + "`" + `` + "`" + `` + "`" + ` "storage.googleapis.com/[GCS_BUCKET]" "bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]" "pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]" ` + "`" + `` + "`" + `` + "`" + ` The writer associated with the sink must have access to write to the above resource.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) The filter to apply when exporting logs. Only log entries that match the filter are exported. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to write a filter.`,
				},
				resource.Attribute{
					Name:        "include_children",
					Description: `(Optional) Whether or not to include children organizations in the sink export. If true, logs associated with child projects are also exported; otherwise only logs relating to the provided organization are included. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "writer_identity",
					Description: `The identity associated with this sink. This identity must be granted write access to the configured ` + "`" + `destination` + "`" + `. ## Import Organization-level logging sinks can be imported using this format: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_logging_organization_sink.my_sink organizations/{{organization_id}}/sinks/{{sink_id}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "writer_identity",
					Description: `The identity associated with this sink. This identity must be granted write access to the configured ` + "`" + `destination` + "`" + `. ## Import Organization-level logging sinks can be imported using this format: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_logging_organization_sink.my_sink organizations/{{organization_id}}/sinks/{{sink_id}} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_logging_project_exclusion",
			Category:         "Google Stackdriver Logging Resources",
			ShortDescription: `Manages a project-level logging exclusion.`,
			Description:      ``,
			Keywords: []string{
				"stackdriver",
				"logging",
				"project",
				"exclusion",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) The filter to apply when excluding logs. Only log entries that match the filter are excluded. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to write a filter.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the logging exclusion.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-readable description.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Whether this exclusion rule should be disabled or not. This defaults to false.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project to create the exclusion in. If omitted, the project associated with the provider is used. ## Import Project-level logging exclusions can be imported using their URI, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_logging_project_exclusion.my_exclusion projects/my-project/exclusions/my-exclusion ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_logging_project_sink",
			Category:         "Google Stackdriver Logging Resources",
			ShortDescription: `Manages a project-level logging sink.`,
			Description:      ``,
			Keywords: []string{
				"stackdriver",
				"logging",
				"project",
				"sink",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the logging sink.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Required) The destination of the sink (or, in other words, where logs are written to). Can be a Cloud Storage bucket, a PubSub topic, or a BigQuery dataset. Examples: ` + "`" + `` + "`" + `` + "`" + ` "storage.googleapis.com/[GCS_BUCKET]" "bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]" "pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]" ` + "`" + `` + "`" + `` + "`" + ` The writer associated with the sink must have access to write to the above resource.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) The filter to apply when exporting logs. Only log entries that match the filter are exported. See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to write a filter.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project to create the sink in. If omitted, the project associated with the provider is used.`,
				},
				resource.Attribute{
					Name:        "unique_writer_identity",
					Description: `(Optional) Whether or not to create a unique identity associated with this sink. If ` + "`" + `false` + "`" + ` (the default), then the ` + "`" + `writer_identity` + "`" + ` used is ` + "`" + `serviceAccount:cloud-logs@system.gserviceaccount.com` + "`" + `. If ` + "`" + `true` + "`" + `, then a unique service account is created and used for this sink. If you wish to publish logs across projects, you must set ` + "`" + `unique_writer_identity` + "`" + ` to true. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "writer_identity",
					Description: `The identity associated with this sink. This identity must be granted write access to the configured ` + "`" + `destination` + "`" + `. ## Import Project-level logging sinks can be imported using their URI, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_logging_project_sink.my_sink projects/my-project/sinks/my-sink ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "writer_identity",
					Description: `The identity associated with this sink. This identity must be granted write access to the configured ` + "`" + `destination` + "`" + `. ## Import Project-level logging sinks can be imported using their URI, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_logging_project_sink.my_sink projects/my-project/sinks/my-sink ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_ml_engine_model",
			Category:         "Google ML Engine Resources",
			ShortDescription: `Represents a machine learning solution.`,
			Description:      ``,
			Keywords: []string{
				"ml",
				"engine",
				"model",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name specified for the model. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description specified for the model when it was created.`,
				},
				resource.Attribute{
					Name:        "default_version",
					Description: `(Optional) The default version of the model. This version will be used to handle prediction requests that do not specify a version. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `(Optional) The list of regions where the model is going to be deployed. Currently only one region per model is supported`,
				},
				resource.Attribute{
					Name:        "online_prediction_logging",
					Description: `(Optional) If true, online prediction access logs are sent to StackDriver Logging.`,
				},
				resource.Attribute{
					Name:        "online_prediction_console_logging",
					Description: `(Optional) If true, online prediction nodes send stderr and stdout streams to Stackdriver Logging`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) One or more labels that you can add, to organize your models.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `default_version` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name specified for the version when it was created. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Model can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_ml_engine_model.default projects/{{project}}/models/{{name}} $ terraform import google_ml_engine_model.default {{project}}/{{name}} $ terraform import google_ml_engine_model.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_monitoring_alert_policy",
			Category:         "Google Stackdriver Monitoring Resources",
			ShortDescription: `A description of the conditions under which some aspect of your system is considered to be "unhealthy" and the ways to notify people or services about this state.`,
			Description:      ``,
			Keywords: []string{
				"stackdriver",
				"monitoring",
				"alert",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid confusion, don't use the same display name for multiple policies in the same project. The name is limited to 512 Unicode characters.`,
				},
				resource.Attribute{
					Name:        "combiner",
					Description: `(Required) How to combine the results of multiple conditions to determine if an incident should be opened.`,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `(Required) A list of conditions for the policy. The conditions are combined by AND or OR according to the combiner field. If the combined conditions evaluate to true, then an incident is created. A policy can have from one to six conditions. Structure is documented below. The ` + "`" + `conditions` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "condition_absent",
					Description: `(Optional) A condition that checks that a time series continues to receive new data points. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The unique resource name for this condition. Its syntax is: projects/[PROJECT_ID]/alertPolicies/[POLICY_ID]/conditions/[CONDITION_ID] [CONDITION_ID] is assigned by Stackdriver Monitoring when the condition is created as part of a new or updated alerting policy.`,
				},
				resource.Attribute{
					Name:        "condition_threshold",
					Description: `(Optional) A condition that compares a time series against a threshold. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) A short name or phrase used to identify the condition in dashboards, notifications, and incidents. To avoid confusion, don't use the same display name for multiple conditions in the same policy. The ` + "`" + `condition_absent` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "aggregations",
					Description: `(Optional) Specifies the alignment of data points in individual time series as well as how to combine the retrieved time series together (such as when aggregating multiple streams on each resource to a single stream for each resource or when aggregating streams across all members of a group of resources). Multiple aggregations are applied in the order specified. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "trigger",
					Description: `(Optional) The number/percent of time series for which the comparison must hold in order for the condition to trigger. If unspecified, then the condition will trigger if the comparison is true for any of the time series that have been identified by filter and aggregations. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Required) The amount of time that a time series must fail to report new data to be considered failing. Currently, only values that are a multiple of a minute--e.g. 60s, 120s, or 300s --are supported.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A filter that identifies which time series should be compared with the threshold.The filter is similar to the one that is specified in the MetricService.ListTimeSeries request (that call is useful to verify the time series that will be retrieved / processed) and must specify the metric type and optionally may contain restrictions on resource type, resource labels, and metric labels. This field may not exceed 2048 Unicode characters in length. The ` + "`" + `aggregations` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "per_series_aligner",
					Description: `(Optional) The approach to be used to align individual time series. Not all alignment functions may be applied to all time series, depending on the metric type and value type of the original time series. Alignment may change the metric type or the value type of the time series.Time series data must be aligned in order to perform cross- time series reduction. If crossSeriesReducer is specified, then perSeriesAligner must be specified and not equal ALIGN_NONE and alignmentPeriod must be specified; otherwise, an error is returned.`,
				},
				resource.Attribute{
					Name:        "group_by_fields",
					Description: `(Optional) The set of fields to preserve when crossSeriesReducer is specified. The groupByFields determine how the time series are partitioned into subsets prior to applying the aggregation function. Each subset contains time series that have the same value for each of the grouping fields. Each individual time series is a member of exactly one subset. The crossSeriesReducer is applied to each subset of time series. It is not possible to reduce across different resource types, so this field implicitly contains resource.type. Fields not specified in groupByFields are aggregated away. If groupByFields is not specified and all the time series have the same resource type, then the time series are aggregated into a single output time series. If crossSeriesReducer is not defined, this field is ignored.`,
				},
				resource.Attribute{
					Name:        "alignment_period",
					Description: `(Optional) The alignment period for per-time series alignment. If present, alignmentPeriod must be at least 60 seconds. After per-time series alignment, each time series will contain data points only on the period boundaries. If perSeriesAligner is not specified or equals ALIGN_NONE, then this field is ignored. If perSeriesAligner is specified and does not equal ALIGN_NONE, then this field must be defined; otherwise an error is returned.`,
				},
				resource.Attribute{
					Name:        "cross_series_reducer",
					Description: `(Optional) The approach to be used to combine time series. Not all reducer functions may be applied to all time series, depending on the metric type and the value type of the original time series. Reduction may change the metric type of value type of the time series.Time series data must be aligned in order to perform cross- time series reduction. If crossSeriesReducer is specified, then perSeriesAligner must be specified and not equal ALIGN_NONE and alignmentPeriod must be specified; otherwise, an error is returned. The ` + "`" + `trigger` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "percent",
					Description: `(Optional) The percentage of time series that must fail the predicate for the condition to be triggered.`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `(Optional) The absolute number of time series that must fail the predicate for the condition to be triggered. The ` + "`" + `condition_threshold` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "threshold_value",
					Description: `(Optional) A value against which to compare the time series.`,
				},
				resource.Attribute{
					Name:        "denominator_filter",
					Description: `(Optional) A filter that identifies a time series that should be used as the denominator of a ratio that will be compared with the threshold. If a denominator_filter is specified, the time series specified by the filter field will be used as the numerator.The filter is similar to the one that is specified in the MetricService.ListTimeSeries request (that call is useful to verify the time series that will be retrieved / processed) and must specify the metric type and optionally may contain restrictions on resource type, resource labels, and metric labels. This field may not exceed 2048 Unicode characters in length.`,
				},
				resource.Attribute{
					Name:        "denominator_aggregations",
					Description: `(Optional) Specifies the alignment of data points in individual time series selected by denominatorFilter as well as how to combine the retrieved time series together (such as when aggregating multiple streams on each resource to a single stream for each resource or when aggregating streams across all members of a group of resources).When computing ratios, the aggregations and denominator_aggregations fields must use the same alignment period and produce time series that have the same periodicity and labels.This field is similar to the one in the MetricService.ListTimeSeries request. It is advisable to use the ListTimeSeries method when debugging this field. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Required) The amount of time that a time series must violate the threshold to be considered failing. Currently, only values that are a multiple of a minute--e.g., 0, 60, 120, or 300 seconds--are supported. If an invalid value is given, an error will be returned. When choosing a duration, it is useful to keep in mind the frequency of the underlying time series data (which may also be affected by any alignments specified in the aggregations field); a good duration is long enough so that a single outlier does not generate spurious alerts, but short enough that unhealthy states are detected and alerted on quickly.`,
				},
				resource.Attribute{
					Name:        "comparison",
					Description: `(Required) The comparison to apply between the time series (indicated by filter and aggregation) and the threshold (indicated by threshold_value). The comparison is applied on each time series, with the time series on the left-hand side and the threshold on the right-hand side. Only COMPARISON_LT and COMPARISON_GT are supported currently.`,
				},
				resource.Attribute{
					Name:        "trigger",
					Description: `(Optional) The number/percent of time series for which the comparison must hold in order for the condition to trigger. If unspecified, then the condition will trigger if the comparison is true for any of the time series that have been identified by filter and aggregations, or by the ratio, if denominator_filter and denominator_aggregations are specified. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "aggregations",
					Description: `(Optional) Specifies the alignment of data points in individual time series as well as how to combine the retrieved time series together (such as when aggregating multiple streams on each resource to a single stream for each resource or when aggregating streams across all members of a group of resources). Multiple aggregations are applied in the order specified.This field is similar to the one in the MetricService.ListTimeSeries request. It is advisable to use the ListTimeSeries method when debugging this field. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) A filter that identifies which time series should be compared with the threshold.The filter is similar to the one that is specified in the MetricService.ListTimeSeries request (that call is useful to verify the time series that will be retrieved / processed) and must specify the metric type and optionally may contain restrictions on resource type, resource labels, and metric labels. This field may not exceed 2048 Unicode characters in length. The ` + "`" + `denominator_aggregations` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "per_series_aligner",
					Description: `(Optional) The approach to be used to align individual time series. Not all alignment functions may be applied to all time series, depending on the metric type and value type of the original time series. Alignment may change the metric type or the value type of the time series.Time series data must be aligned in order to perform cross- time series reduction. If crossSeriesReducer is specified, then perSeriesAligner must be specified and not equal ALIGN_NONE and alignmentPeriod must be specified; otherwise, an error is returned.`,
				},
				resource.Attribute{
					Name:        "group_by_fields",
					Description: `(Optional) The set of fields to preserve when crossSeriesReducer is specified. The groupByFields determine how the time series are partitioned into subsets prior to applying the aggregation function. Each subset contains time series that have the same value for each of the grouping fields. Each individual time series is a member of exactly one subset. The crossSeriesReducer is applied to each subset of time series. It is not possible to reduce across different resource types, so this field implicitly contains resource.type. Fields not specified in groupByFields are aggregated away. If groupByFields is not specified and all the time series have the same resource type, then the time series are aggregated into a single output time series. If crossSeriesReducer is not defined, this field is ignored.`,
				},
				resource.Attribute{
					Name:        "alignment_period",
					Description: `(Optional) The alignment period for per-time series alignment. If present, alignmentPeriod must be at least 60 seconds. After per-time series alignment, each time series will contain data points only on the period boundaries. If perSeriesAligner is not specified or equals ALIGN_NONE, then this field is ignored. If perSeriesAligner is specified and does not equal ALIGN_NONE, then this field must be defined; otherwise an error is returned.`,
				},
				resource.Attribute{
					Name:        "cross_series_reducer",
					Description: `(Optional) The approach to be used to combine time series. Not all reducer functions may be applied to all time series, depending on the metric type and the value type of the original time series. Reduction may change the metric type of value type of the time series.Time series data must be aligned in order to perform cross- time series reduction. If crossSeriesReducer is specified, then perSeriesAligner must be specified and not equal ALIGN_NONE and alignmentPeriod must be specified; otherwise, an error is returned. The ` + "`" + `trigger` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "percent",
					Description: `(Optional) The percentage of time series that must fail the predicate for the condition to be triggered.`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `(Optional) The absolute number of time series that must fail the predicate for the condition to be triggered. The ` + "`" + `aggregations` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "per_series_aligner",
					Description: `(Optional) The approach to be used to align individual time series. Not all alignment functions may be applied to all time series, depending on the metric type and value type of the original time series. Alignment may change the metric type or the value type of the time series.Time series data must be aligned in order to perform cross- time series reduction. If crossSeriesReducer is specified, then perSeriesAligner must be specified and not equal ALIGN_NONE and alignmentPeriod must be specified; otherwise, an error is returned.`,
				},
				resource.Attribute{
					Name:        "group_by_fields",
					Description: `(Optional) The set of fields to preserve when crossSeriesReducer is specified. The groupByFields determine how the time series are partitioned into subsets prior to applying the aggregation function. Each subset contains time series that have the same value for each of the grouping fields. Each individual time series is a member of exactly one subset. The crossSeriesReducer is applied to each subset of time series. It is not possible to reduce across different resource types, so this field implicitly contains resource.type. Fields not specified in groupByFields are aggregated away. If groupByFields is not specified and all the time series have the same resource type, then the time series are aggregated into a single output time series. If crossSeriesReducer is not defined, this field is ignored.`,
				},
				resource.Attribute{
					Name:        "alignment_period",
					Description: `(Optional) The alignment period for per-time series alignment. If present, alignmentPeriod must be at least 60 seconds. After per-time series alignment, each time series will contain data points only on the period boundaries. If perSeriesAligner is not specified or equals ALIGN_NONE, then this field is ignored. If perSeriesAligner is specified and does not equal ALIGN_NONE, then this field must be defined; otherwise an error is returned.`,
				},
				resource.Attribute{
					Name:        "cross_series_reducer",
					Description: `(Optional) The approach to be used to combine time series. Not all reducer functions may be applied to all time series, depending on the metric type and the value type of the original time series. Reduction may change the metric type of value type of the time series.Time series data must be aligned in order to perform cross- time series reduction. If crossSeriesReducer is specified, then perSeriesAligner must be specified and not equal ALIGN_NONE and alignmentPeriod must be specified; otherwise, an error is returned. - - -`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether or not the policy is enabled. The default is true.`,
				},
				resource.Attribute{
					Name:        "notification_channels",
					Description: `(Optional) Identifies the notification channels to which notifications should be sent when incidents are opened or closed or when new violations occur on an already opened incident. Each element of this array corresponds to the name field in each of the NotificationChannel objects that are returned from the notificationChannels.list method. The syntax of the entries in this field is ` + "`" + `projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID]` + "`" + ``,
				},
				resource.Attribute{
					Name:        "user_labels",
					Description: `(Optional) This field is intended to be used for organizing and identifying the AlertPolicy objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.`,
				},
				resource.Attribute{
					Name:        "documentation",
					Description: `(Optional) A short name or phrase used to identify the policy in dashboards, notifications, and incidents. To avoid confusion, don't use the same display name for multiple policies in the same project. The name is limited to 512 Unicode characters. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `documentation` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional) The text of the documentation, interpreted according to mimeType. The content may not exceed 8,192 Unicode characters and may not exceed more than 10,240 bytes when encoded in UTF-8 format, whichever is smaller.`,
				},
				resource.Attribute{
					Name:        "mime_type",
					Description: `(Optional) The format of the content field. Presently, only the value "text/markdown" is supported. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The unique resource name for this policy. Its syntax is: projects/[PROJECT_ID]/alertPolicies/[ALERT_POLICY_ID]`,
				},
				resource.Attribute{
					Name:        "creation_record",
					Description: `A read-only record of the creation of the alerting policy. If provided in a call to create or update, this field will be ignored. Structure is documented below. The ` + "`" + `creation_record` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "mutate_time",
					Description: `When the change occurred.`,
				},
				resource.Attribute{
					Name:        "mutated_by",
					Description: `The email address of the user making the change. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import AlertPolicy can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_monitoring_alert_policy.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The unique resource name for this policy. Its syntax is: projects/[PROJECT_ID]/alertPolicies/[ALERT_POLICY_ID]`,
				},
				resource.Attribute{
					Name:        "creation_record",
					Description: `A read-only record of the creation of the alerting policy. If provided in a call to create or update, this field will be ignored. Structure is documented below. The ` + "`" + `creation_record` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "mutate_time",
					Description: `When the change occurred.`,
				},
				resource.Attribute{
					Name:        "mutated_by",
					Description: `The email address of the user making the change. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import AlertPolicy can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_monitoring_alert_policy.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_monitoring_group",
			Category:         "Google Stackdriver Monitoring Resources",
			ShortDescription: `The description of a dynamic collection of monitored resources.`,
			Description:      ``,
			Keywords: []string{
				"stackdriver",
				"monitoring",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) A user-assigned name for this group, used only for display purposes.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) The filter used to determine which monitored resources belong to this group. - - -`,
				},
				resource.Attribute{
					Name:        "parent_name",
					Description: `(Optional) The name of the group's parent, if it has one. The format is "projects/{project_id_or_number}/groups/{group_id}". For groups with no parent, parentName is the empty string, "".`,
				},
				resource.Attribute{
					Name:        "is_cluster",
					Description: `(Optional) If true, the members of this group are considered to be a cluster. The system can perform additional analysis on groups that are clusters.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A unique identifier for this group. The format is "projects/{project_id_or_number}/groups/{group_id}". ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Group can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_monitoring_group.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `A unique identifier for this group. The format is "projects/{project_id_or_number}/groups/{group_id}". ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Group can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_monitoring_group.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_monitoring_notification_channel",
			Category:         "Google Stackdriver Monitoring Resources",
			ShortDescription: `A NotificationChannel is a medium through which an alert is delivered when a policy violation is detected.`,
			Description:      ``,
			Keywords: []string{
				"stackdriver",
				"monitoring",
				"notification",
				"channel",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field. See https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list to get the list of valid values such as "email", "slack", etc...`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and unique name in order to make it easier to identify the channels in your project, though this is not enforced. The display name is limited to 512 Unicode characters. - - -`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Configuration fields that define the channel and its behavior. The permissible and required labels are specified in the NotificationChannelDescriptor corresponding to the type field.`,
				},
				resource.Attribute{
					Name:        "user_labels",
					Description: `(Optional) User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema, unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional human-readable description of this notification channel. This description may provide additional details, beyond the display name, for the channel. This may not exceed 1024 Unicode characters.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of notifications to a particular channel without removing the channel from all alerting policies that reference the channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the same set of alerting policies on the channel at some point in the future.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The full REST resource name for this channel. The syntax is: projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID] The [CHANNEL_ID] is automatically assigned by the server on creation.`,
				},
				resource.Attribute{
					Name:        "verification_status",
					Description: `Indicates whether this channel has been verified or not. On a ListNotificationChannels or GetNotificationChannel operation, this field is expected to be populated.If the value is UNVERIFIED, then it indicates that the channel is non-functioning (it both requires verification and lacks verification); otherwise, it is assumed that the channel works.If the channel is neither VERIFIED nor UNVERIFIED, it implies that the channel is of a type that does not require verification or that this specific channel has been exempted from verification because it was created prior to verification being required for channels of this type.This field cannot be modified using a standard UpdateNotificationChannel operation. To change the value of this field, you must call VerifyNotificationChannel. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import NotificationChannel can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_monitoring_notification_channel.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The full REST resource name for this channel. The syntax is: projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID] The [CHANNEL_ID] is automatically assigned by the server on creation.`,
				},
				resource.Attribute{
					Name:        "verification_status",
					Description: `Indicates whether this channel has been verified or not. On a ListNotificationChannels or GetNotificationChannel operation, this field is expected to be populated.If the value is UNVERIFIED, then it indicates that the channel is non-functioning (it both requires verification and lacks verification); otherwise, it is assumed that the channel works.If the channel is neither VERIFIED nor UNVERIFIED, it implies that the channel is of a type that does not require verification or that this specific channel has been exempted from verification because it was created prior to verification being required for channels of this type.This field cannot be modified using a standard UpdateNotificationChannel operation. To change the value of this field, you must call VerifyNotificationChannel. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import NotificationChannel can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_monitoring_notification_channel.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_monitoring_uptime_check_config",
			Category:         "Google Stackdriver Monitoring Resources",
			ShortDescription: `This message configures which resources and services to monitor for availability.`,
			Description:      ``,
			Keywords: []string{
				"stackdriver",
				"monitoring",
				"uptime",
				"check",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) A human-friendly name for the uptime check configuration. The display name should be unique within a Stackdriver Workspace in order to make it easier to identify; however, uniqueness is not enforced.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Required) The maximum amount of time to wait for the request to complete (must be between 1 and 60 seconds). Accepted formats https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#google.protobuf.Duration - - -`,
				},
				resource.Attribute{
					Name:        "period",
					Description: `(Optional) How often, in seconds, the uptime check is performed. Currently, the only supported values are 60s (1 minute), 300s (5 minutes), 600s (10 minutes), and 900s (15 minutes). Optional, defaults to 300s.`,
				},
				resource.Attribute{
					Name:        "content_matchers",
					Description: `(Optional) The expected content on the page the check is run against. Currently, only the first entry in the list is supported, and other entries will be ignored. The server will look for an exact match of the string in the page response's content. This field is optional and should only be specified if a content match is required. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "selected_regions",
					Description: `(Optional) The list of regions from which the check will be run. Some regions contain one location, and others contain more than one. If this field is specified, enough regions to include a minimum of 3 locations must be provided, or an error message is returned. Not specifying this field will result in uptime checks running from all regions.`,
				},
				resource.Attribute{
					Name:        "http_check",
					Description: `(Optional) Contains information needed to make an HTTP or HTTPS check. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "tcp_check",
					Description: `(Optional) Contains information needed to make a TCP check. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "resource_group",
					Description: `(Optional) The group resource associated with the configuration. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "monitored_resource",
					Description: `(Optional) The monitored resource (https://cloud.google.com/monitoring/api/resources) associated with the configuration. The following monitored resource types are supported for uptime checks: uptime_url gce_instance gae_app aws_ec2_instance aws_elb_load_balancer Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `content_matchers` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional) String or regex content to match (max 1024 bytes) The ` + "`" + `http_check` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "auth_info",
					Description: `(Optional) The authentication information. Optional when creating an HTTP check; defaults to empty. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port to the page to run the check against. Will be combined with host (specified within the MonitoredResource) and path to construct the full URL. Optional (defaults to 80 without SSL, or 443 with SSL).`,
				},
				resource.Attribute{
					Name:        "headers",
					Description: `(Optional) The list of headers to send as part of the uptime check request. If two headers have the same key and different values, they should be entered as a single header, with the value being a comma-separated list of all the desired values as described at https://www.w3.org/Protocols/rfc2616/rfc2616.txt (page 31). Entering two separate headers with the same key in a Create call will cause the first to be overwritten by the second. The maximum number of headers allowed is 100.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The path to the page to run the check against. Will be combined with the host (specified within the MonitoredResource) and port to construct the full URL. Optional (defaults to "/").`,
				},
				resource.Attribute{
					Name:        "use_ssl",
					Description: `(Optional) If true, use HTTPS instead of HTTP to run the check.`,
				},
				resource.Attribute{
					Name:        "mask_headers",
					Description: `(Optional) Boolean specifying whether to encrypt the header information. Encryption should be specified for any headers related to authentication that you do not wish to be seen when retrieving the configuration. The server will be responsible for encrypting the headers. On Get/List calls, if mask_headers is set to True then the headers will be obscured with`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The password to authenticate.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) The username to authenticate. The ` + "`" + `tcp_check` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) The port to the page to run the check against. Will be combined with host (specified within the MonitoredResource) to construct the full URL. The ` + "`" + `resource_group` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Optional) The resource type of the group members.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `(Optional) The group of resources being monitored. Should be the ` + "`" + `name` + "`" + ` of a group The ` + "`" + `monitored_resource` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The monitored resource type. This field must match the type field of a MonitoredResourceDescriptor (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.monitoredResourceDescriptors#MonitoredResourceDescriptor) object. For example, the type of a Compute Engine VM instance is gce_instance. For a list of types, see Monitoring resource types (https://cloud.google.com/monitoring/api/resources) and Logging resource types (https://cloud.google.com/logging/docs/api/v2/resource-list).`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Required) Values for all of the labels listed in the associated monitored resource descriptor. For example, Compute Engine VM instances use the labels "project_id", "instance_id", and "zone". ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A unique resource name for this UptimeCheckConfig. The format is projects/[PROJECT_ID]/uptimeCheckConfigs/[UPTIME_CHECK_ID].`,
				},
				resource.Attribute{
					Name:        "uptime_check_id",
					Description: `The id of the uptime check ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import UptimeCheckConfig can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_monitoring_uptime_check_config.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `A unique resource name for this UptimeCheckConfig. The format is projects/[PROJECT_ID]/uptimeCheckConfigs/[UPTIME_CHECK_ID].`,
				},
				resource.Attribute{
					Name:        "uptime_check_id",
					Description: `The id of the uptime check ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import UptimeCheckConfig can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_monitoring_uptime_check_config.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_pubsub_subscription",
			Category:         "Google PubSub Resources",
			ShortDescription: `A named resource representing the stream of messages from a single, specific topic, to be delivered to the subscribing application.`,
			Description:      ``,
			Keywords: []string{
				"pubsub",
				"subscription",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the subscription.`,
				},
				resource.Attribute{
					Name:        "topic",
					Description: `(Required) A reference to a Topic resource. - - -`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs to assign to this Subscription.`,
				},
				resource.Attribute{
					Name:        "push_config",
					Description: `(Optional) If push delivery is used with this subscription, this field is used to configure it. An empty pushConfig signifies that the subscriber will pull and ack messages using API methods. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "ack_deadline_seconds",
					Description: `(Optional) This value is the maximum time after a subscriber receives a message before the subscriber should acknowledge the message. After message delivery but before the ack deadline expires and before the message is acknowledged, it is an outstanding message and will not be delivered again during that time (on a best-effort basis). For pull subscriptions, this value is used as the initial value for the ack deadline. To override this value for a given message, call subscriptions.modifyAckDeadline with the corresponding ackId if using pull. The minimum custom deadline you can specify is 10 seconds. The maximum custom deadline you can specify is 600 seconds (10 minutes). If this parameter is 0, a default value of 10 seconds is used. For push delivery, this value is also used to set the request timeout for the call to the push endpoint. If the subscriber never acknowledges the message, the Pub/Sub system will eventually redeliver the message.`,
				},
				resource.Attribute{
					Name:        "message_retention_duration",
					Description: `(Optional) How long to retain unacknowledged messages in the subscription's backlog, from the moment a message is published. If retainAckedMessages is true, then this also configures the retention of acknowledged messages, and thus configures how far back in time a subscriptions.seek can be done. Defaults to 7 days. Cannot be more than 7 days (` + "`" + `"604800s"` + "`" + `) or less than 10 minutes (` + "`" + `"600s"` + "`" + `). A duration in seconds with up to nine fractional digits, terminated by 's'. Example: ` + "`" + `"600.5s"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "retain_acked_messages",
					Description: `(Optional) Indicates whether to retain acknowledged messages. If ` + "`" + `true` + "`" + `, then messages are not expunged from the subscription's backlog, even if they are acknowledged, until they fall out of the messageRetentionDuration window.`,
				},
				resource.Attribute{
					Name:        "expiration_policy",
					Description: `(Optional) A policy that specifies the conditions for this subscription's expiration. A subscription is considered active as long as any connected subscriber is successfully consuming messages from the subscription or is issuing operations on the subscription. If expirationPolicy is not set, a default policy with ttl of 31 days will be used. The minimum allowed value for expirationPolicy.ttl is 1 day. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `push_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "push_endpoint",
					Description: `(Required) A URL locating the endpoint to which messages should be pushed. For example, a Webhook endpoint might use "https://example.com/push".`,
				},
				resource.Attribute{
					Name:        "attributes",
					Description: `(Optional) Endpoint configuration attributes. Every endpoint has a set of API supported attributes that can be used to control different aspects of the message delivery. The currently supported attribute is x-goog-version, which you can use to change the format of the pushed message. This attribute indicates the version of the data expected by the endpoint. This controls the shape of the pushed message (i.e., its fields and metadata). The endpoint version is based on the version of the Pub/Sub API. If not present during the subscriptions.create call, it will default to the version of the API used to make such call. If not present during a subscriptions.modifyPushConfig call, its value will not be changed. subscriptions.get calls will always return a valid version, even if the subscription was created without this attribute. The possible values for this attribute are: - v1beta1: uses the push format defined in the v1beta1 Pub/Sub API. - v1 or v1beta2: uses the push format defined in the v1 Pub/Sub API. The ` + "`" + `expiration_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) Specifies the "time-to-live" duration for an associated resource. The resource expires if it is not active for a period of ttl. The definition of "activity" depends on the type of the associated resource. The minimum and maximum allowed values for ttl depend on the type of the associated resource, as well. If ttl is not set, the associated resource never expires. A duration in seconds with up to nine fractional digits, terminated by 's'. Example - "3.5s". ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_pubsub_subscription_iam_policy",
			Category:         "Google PubSub Resources",
			ShortDescription: `Collection of resources to manage IAM policy for a Pubsub subscription.`,
			Description:      ``,
			Keywords: []string{
				"pubsub",
				"subscription",
				"iam",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "subscription",
					Description: `(Required) The subscription name or id to bind to attach IAM policy to.`,
				},
				resource.Attribute{
					Name:        "member/members",
					Description: `(Required) Identities that will be granted the privilege in ` + "`" + `role` + "`" + `. Each entry can have one of the following values:`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. Only one ` + "`" + `google_pubsub_subscription_iam_binding` + "`" + ` can be used per role. Note that custom roles must be of the format ` + "`" + `[projects|organizations]/{parent-name}/roles/{role-name}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy_data",
					Description: `(Required only by ` + "`" + `google_pubsub_subscription_iam_policy` + "`" + `) The policy data generated by a ` + "`" + `google_iam_policy` + "`" + ` data source. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the subscription's IAM policy. ## Import Pubsub subscription IAM resources can be imported using the project, subscription name, role and member. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_pubsub_subscription_iam_policy.editor projects/{your-project-id}/subscriptions/{your-subscription-name} $ terraform import google_pubsub_subscription_iam_binding.editor "projects/{your-project-id}/subscriptions/{your-subscription-name} roles/editor" $ terraform import google_pubsub_subscription_iam_member.editor "projects/{your-project-id}/subscriptions/{your-subscription-name} roles/editor jane@example.com" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the subscription's IAM policy. ## Import Pubsub subscription IAM resources can be imported using the project, subscription name, role and member. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_pubsub_subscription_iam_policy.editor projects/{your-project-id}/subscriptions/{your-subscription-name} $ terraform import google_pubsub_subscription_iam_binding.editor "projects/{your-project-id}/subscriptions/{your-subscription-name} roles/editor" $ terraform import google_pubsub_subscription_iam_member.editor "projects/{your-project-id}/subscriptions/{your-subscription-name} roles/editor jane@example.com" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_pubsub_topic",
			Category:         "Google PubSub Resources",
			ShortDescription: `A named resource to which messages are sent by publishers.`,
			Description:      ``,
			Keywords: []string{
				"pubsub",
				"topic",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the topic. - - -`,
				},
				resource.Attribute{
					Name:        "kms_key_name",
					Description: `(Optional) The resource name of the Cloud KMS CryptoKey to be used to protect access to messages published on this topic. Your project's PubSub service account (` + "`" + `service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com` + "`" + `) must have ` + "`" + `roles/cloudkms.cryptoKeyEncrypterDecrypter` + "`" + ` to use this feature. The expected format is ` + "`" + `projects/`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs to assign to this Topic.`,
				},
				resource.Attribute{
					Name:        "message_storage_policy",
					Description: `(Optional) Policy constraining the set of Google Cloud Platform regions where messages published to the topic may be stored. If not present, then no constraints are in effect. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `message_storage_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "allowed_persistence_regions",
					Description: `(Required) A list of IDs of GCP regions where messages that are published to the topic may be persisted in storage. Messages published by publishers running in non-allowed GCP regions (or running outside of GCP altogether) will be routed for storage in one of the allowed regions. An empty list means that no regions are allowed, and is not a valid configuration. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Topic can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_pubsub_topic.default projects/{{project}}/topics/{{name}} $ terraform import google_pubsub_topic.default {{project}}/{{name}} $ terraform import google_pubsub_topic.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_pubsub_topic_iam_policy",
			Category:         "Google PubSub Resources",
			ShortDescription: `Collection of resources to manage IAM policy for PubsubTopic`,
			Description:      ``,
			Keywords: []string{
				"pubsub",
				"topic",
				"iam",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "topic",
					Description: `(Required) The topic name or id to bind to attach IAM policy to.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "member/members",
					Description: `(Required) Identities that will be granted the privilege in ` + "`" + `role` + "`" + `. Each entry can have one of the following values:`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. Only one ` + "`" + `google_pubsub_topic_iam_binding` + "`" + ` can be used per role. Note that custom roles must be of the format ` + "`" + `[projects|organizations]/{parent-name}/roles/{role-name}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy_data",
					Description: `(Required only by ` + "`" + `google_pubsub_topic_iam_policy` + "`" + `) The policy data generated by a ` + "`" + `google_iam_policy` + "`" + ` data source. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the topic's IAM policy. ## Import Pubsub topic IAM resources can be imported using the project, topic name, role and member. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_pubsub_topic_iam_policy.editor projects/{{project}}/topics/{{topic}} $ terraform import google_pubsub_topic_iam_binding.editor "projects/{{project}}/topics/{{topic}} roles/editor" $ terraform import google_pubsub_topic_iam_member.editor "projects/{{project}}/topics/{{topic}} roles/editor jane@example.com" ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the topic's IAM policy. ## Import Pubsub topic IAM resources can be imported using the project, topic name, role and member. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_pubsub_topic_iam_policy.editor projects/{{project}}/topics/{{topic}} $ terraform import google_pubsub_topic_iam_binding.editor "projects/{{project}}/topics/{{topic}} roles/editor" $ terraform import google_pubsub_topic_iam_member.editor "projects/{{project}}/topics/{{topic}} roles/editor jane@example.com" ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_redis_instance",
			Category:         "Google Redis (Cloud Memorystore) Resources",
			ShortDescription: `A Google Cloud Redis instance.`,
			Description:      ``,
			Keywords: []string{
				"redis",
				"cloud",
				"memorystore",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The ID of the instance or a fully qualified identifier for the instance.`,
				},
				resource.Attribute{
					Name:        "memory_size_gb",
					Description: `(Required) Redis memory size in GiB. - - -`,
				},
				resource.Attribute{
					Name:        "alternative_location_id",
					Description: `(Optional) Only applicable to STANDARD_HA tier which protects the instance against zonal failures by provisioning it across two zones. If provided, it must be a different zone from the one provided in [locationId].`,
				},
				resource.Attribute{
					Name:        "authorized_network",
					Description: `(Optional) The full name of the Google Compute Engine network to which the instance is connected. If left unspecified, the default network will be used.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) An arbitrary and optional user-provided name for the instance.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Resource labels to represent user provided metadata.`,
				},
				resource.Attribute{
					Name:        "redis_configs",
					Description: `(Optional) Redis configuration parameters, according to http://redis.io/topics/config. Please check Memorystore documentation for the list of supported parameters: https://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances#Instance.FIELDS.redis_configs`,
				},
				resource.Attribute{
					Name:        "location_id",
					Description: `(Optional) The zone where the instance will be provisioned. If not provided, the service will choose a zone for the instance. For STANDARD_HA tier, instances will be created across two zones for protection against zonal failures. If [alternativeLocationId] is also provided, it must be different from [locationId].`,
				},
				resource.Attribute{
					Name:        "redis_version",
					Description: `(Optional) The version of Redis software. If not provided, latest supported version will be used. Currently, the supported values are: - REDIS_4_0 for Redis 4.0 compatibility - REDIS_3_2 for Redis 3.2 compatibility`,
				},
				resource.Attribute{
					Name:        "reserved_ip_range",
					Description: `(Optional) The CIDR range of internal addresses that are reserved for this instance. If not provided, the service will choose an unused /29 block, for example, 10.0.0.0/29 or 192.168.0.0/29. Ranges must be unique and non-overlapping with existing subnets in an authorized network.`,
				},
				resource.Attribute{
					Name:        "tier",
					Description: `(Optional) The service tier of the instance. Must be one of these values: - BASIC: standalone instance - STANDARD_HA: highly available primary/replica instances`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The name of the Redis region of the instance.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time the instance was created in RFC3339 UTC "Zulu" format, accurate to nanoseconds.`,
				},
				resource.Attribute{
					Name:        "current_location_id",
					Description: `The current zone where the Redis endpoint is placed. For Basic Tier instances, this will always be the same as the [locationId] provided by the user at creation time. For Standard Tier instances, this can be either [locationId] or [alternativeLocationId] and can change after a failover event.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `Hostname or IP address of the exposed Redis endpoint used by clients to connect to the service.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port number of the exposed Redis endpoint. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 10 minutes. - ` + "`" + `update` + "`" + ` - Default is 10 minutes. - ` + "`" + `delete` + "`" + ` - Default is 10 minutes. ## Import Instance can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_redis_instance.default projects/{{project}}/locations/{{region}}/instances/{{name}} $ terraform import google_redis_instance.default {{project}}/{{region}}/{{name}} $ terraform import google_redis_instance.default {{region}}/{{name}} $ terraform import google_redis_instance.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `The time the instance was created in RFC3339 UTC "Zulu" format, accurate to nanoseconds.`,
				},
				resource.Attribute{
					Name:        "current_location_id",
					Description: `The current zone where the Redis endpoint is placed. For Basic Tier instances, this will always be the same as the [locationId] provided by the user at creation time. For Standard Tier instances, this can be either [locationId] or [alternativeLocationId] and can change after a failover event.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `Hostname or IP address of the exposed Redis endpoint used by clients to connect to the service.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port number of the exposed Redis endpoint. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 10 minutes. - ` + "`" + `update` + "`" + ` - Default is 10 minutes. - ` + "`" + `delete` + "`" + ` - Default is 10 minutes. ## Import Instance can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_redis_instance.default projects/{{project}}/locations/{{region}}/instances/{{name}} $ terraform import google_redis_instance.default {{project}}/{{region}}/{{name}} $ terraform import google_redis_instance.default {{region}}/{{name}} $ terraform import google_redis_instance.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_resource_manager_lien",
			Category:         "Google Cloud Platform Resources",
			ShortDescription: `A Lien represents an encumbrance on the actions that can be performed on a resource.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"platform",
				"resource",
				"manager",
				"lien",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "reason",
					Description: `(Required) Concise user-visible strings indicating why an action cannot be performed on a resource. Maximum length of 200 characters.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `(Required) A stable, user-visible/meaningful string identifying the origin of the Lien, intended to be inspected programmatically. Maximum length of 200 characters.`,
				},
				resource.Attribute{
					Name:        "parent",
					Description: `(Required) A reference to the resource this Lien is attached to. The server will validate the parent against those for which Liens are supported. Since a variety of objects can have Liens against them, you must provide the type prefix (e.g. "projects/my-project-name").`,
				},
				resource.Attribute{
					Name:        "restrictions",
					Description: `(Required) The types of operations which should be blocked as a result of this Lien. Each value should correspond to an IAM permission. The server will validate the permissions against those for which Liens are supported. An empty list is meaningless and will be rejected. e.g. ['resourcemanager.projects.delete'] - - - ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A system-generated unique identifier for this Lien.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Time of creation ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Lien can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_resource_manager_lien.default {{parent}}/{{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `A system-generated unique identifier for this Lien.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Time of creation ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Lien can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_resource_manager_lien.default {{parent}}/{{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_runtimeconfig_config",
			Category:         "Google RuntimeConfig Resources",
			ShortDescription: `Manages a RuntimeConfig resource in Google Cloud.`,
			Description:      ``,
			Keywords: []string{
				"runtimeconfig",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the runtime config. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description to associate with the runtime config. ## Import Runtime Configs can be imported using the ` + "`" + `name` + "`" + ` or full config name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_runtimeconfig_config.myconfig myconfig ` + "`" + `` + "`" + `` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_runtimeconfig_config.myconfig projects/my-gcp-project/configs/myconfig ` + "`" + `` + "`" + `` + "`" + ` When importing using only the name, the provider project must be set.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_runtimeconfig_variable",
			Category:         "Google RuntimeConfig Resources",
			ShortDescription: `Manages a RuntimeConfig variable in Google Cloud.`,
			Description:      ``,
			Keywords: []string{
				"runtimeconfig",
				"variable",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the variable to manage. Note that variable names can be hierarchical using slashes (e.g. "prod-variables/hostname").`,
				},
				resource.Attribute{
					Name:        "parent",
					Description: `(Required) The name of the RuntimeConfig resource containing this variable.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `(Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z". ## Import Runtime Config Variables can be imported using the ` + "`" + `name` + "`" + ` or full variable name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_runtimeconfig_variable.myvariable myconfig/myvariable ` + "`" + `` + "`" + `` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_runtimeconfig_variable.myvariable projects/my-gcp-project/configs/myconfig/variables/myvariable ` + "`" + `` + "`" + `` + "`" + ` When importing using only the name, the provider project must be set.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "update_time",
					Description: `(Computed) The timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds, representing when the variable was last updated. Example: "2016-10-09T12:33:37.578138407Z". ## Import Runtime Config Variables can be imported using the ` + "`" + `name` + "`" + ` or full variable name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_runtimeconfig_variable.myvariable myconfig/myvariable ` + "`" + `` + "`" + `` + "`" + ` ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_runtimeconfig_variable.myvariable projects/my-gcp-project/configs/myconfig/variables/myvariable ` + "`" + `` + "`" + `` + "`" + ` When importing using only the name, the provider project must be set.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_scc_source",
			Category:         "Google Cloud Security Command Center (SCC) Resources",
			ShortDescription: `A Cloud Security Command Center's (Cloud SCC) finding source.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"security",
				"command",
				"center",
				"scc",
				"source",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The source’s display name. A source’s display name must be unique amongst its siblings, for example, two sources with the same parent can't share the same display name. The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens, and underscores, and can be no longer than 32 characters.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Required) The organization whose Cloud Security Command Center the Source lives in. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the source (max of 1024 characters). ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The resource name of this source, in the format ` + "`" + `organizations/{{organization}}/sources/{{source}}` + "`" + `. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Source can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_scc_source.default organizations/{{organization}}/sources/{{name}} $ terraform import google_scc_source.default {{organization}}/{{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The resource name of this source, in the format ` + "`" + `organizations/{{organization}}/sources/{{source}}` + "`" + `. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Source can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_scc_source.default organizations/{{organization}}/sources/{{name}} $ terraform import google_scc_source.default {{organization}}/{{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_security_scanner_scan_config",
			Category:         "Google Cloud Security Scanner Resources",
			ShortDescription: `A ScanConfig resource contains the configurations to launch a scan.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"security",
				"scanner",
				"scan",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The user provider display name of the ScanConfig.`,
				},
				resource.Attribute{
					Name:        "starting_urls",
					Description: `(Required) The starting URLs from which the scanner finds site pages. - - -`,
				},
				resource.Attribute{
					Name:        "max_qps",
					Description: `(Optional) The maximum QPS during scanning. A valid value ranges from 5 to 20 inclusively. Defaults to 15.`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `(Optional) The authentication configuration. If specified, service will use the authentication configuration during scanning. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "user_agent",
					Description: `(Optional) Type of the user agents used for scanning`,
				},
				resource.Attribute{
					Name:        "blacklist_patterns",
					Description: `(Optional) The blacklist URL patterns as described in https://cloud.google.com/security-scanner/docs/excluded-urls`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `(Optional) The schedule of the ScanConfig Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "target_platforms",
					Description: `(Optional) Set of Cloud Platforms targeted by the scan. If empty, APP_ENGINE will be used as a default.`,
				},
				resource.Attribute{
					Name:        "export_to_security_command_center",
					Description: `(Optional) Controls export of scan configurations and results to Cloud Security Command Center.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `authentication` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "google_account",
					Description: `(Optional) Describes authentication configuration that uses a Google account. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "custom_account",
					Description: `(Optional) Describes authentication configuration that uses a custom account. Structure is documented below. The ` + "`" + `google_account` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The user name of the Google account.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The password of the Google account. The credential is stored encrypted in GCP. The ` + "`" + `custom_account` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The user name of the custom account.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The password of the custom account. The credential is stored encrypted in GCP.`,
				},
				resource.Attribute{
					Name:        "login_url",
					Description: `(Required) The login form URL of the website. The ` + "`" + `schedule` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "schedule_time",
					Description: `(Optional) A timestamp indicates when the next run will be scheduled. The value is refreshed by the server after each run. If unspecified, it will default to current server time, which means the scan will be scheduled to start immediately.`,
				},
				resource.Attribute{
					Name:        "interval_duration_days",
					Description: `(Required) The duration of time between executions in days ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A server defined name for this index. Format: ` + "`" + `projects/{{project}}/scanConfigs/{{server_generated_id}}` + "`" + ` ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import ScanConfig can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import -provider=google-beta google_security_scanner_scan_config.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `A server defined name for this index. Format: ` + "`" + `projects/{{project}}/scanConfigs/{{server_generated_id}}` + "`" + ` ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import ScanConfig can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import -provider=google-beta google_security_scanner_scan_config.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_service_networking_connection",
			Category:         "Google Service Networking Resources",
			ShortDescription: `Manages creating a private VPC connection to a service provider.`,
			Description:      ``,
			Keywords: []string{
				"service",
				"networking",
				"connection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "network",
					Description: `(Required) Name of VPC network connected with service producers using VPC peering.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Required) Provider peering service that is managing peering connectivity for a service provider organization. For Google services that support this functionality it is 'servicenetworking.googleapis.com'.`,
				},
				resource.Attribute{
					Name:        "reserved_peering_ranges",
					Description: `(Required) Named IP address range(s) of PEERING type reserved for this service provider. Note that invoking this method with a different range when connection is already established will not reallocate already provisioned service producer subnetworks.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_sourcerepo_repository",
			Category:         "Google Source Repositories Resources",
			ShortDescription: `A repository (or repo) is a Git repository storing versioned source content.`,
			Description:      ``,
			Keywords: []string{
				"source",
				"repositories",
				"sourcerepo",
				"repository",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Resource name of the repository, of the form ` + "`" + `{{repo}}` + "`" + `. The repo name may contain slashes. eg, ` + "`" + `name/with/slash` + "`" + ` - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `URL to clone the repository from Google Cloud Source Repositories.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The disk usage of the repo, in bytes. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Repository can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_sourcerepo_repository.default projects/{{project}}/repos/{{name}} $ terraform import google_sourcerepo_repository.default {{project}}/{{name}} $ terraform import google_sourcerepo_repository.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "url",
					Description: `URL to clone the repository from Google Cloud Source Repositories.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The disk usage of the repo, in bytes. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Repository can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_sourcerepo_repository.default projects/{{project}}/repos/{{name}} $ terraform import google_sourcerepo_repository.default {{project}}/{{name}} $ terraform import google_sourcerepo_repository.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_sourcerepo_repository_iam_policy",
			Category:         "Google Source Repositories Resources",
			ShortDescription: `Collection of resources to manage IAM policy for SourceRepoRepository`,
			Description:      ``,
			Keywords: []string{
				"source",
				"repositories",
				"sourcerepo",
				"repository",
				"iam",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "repository",
					Description: `(Required) The repository name or id to bind to attach IAM policy to.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "member/members",
					Description: `(Required) Identities that will be granted the privilege in ` + "`" + `role` + "`" + `. Each entry can have one of the following values:`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. Only one ` + "`" + `google_sourcerepo_repository_iam_binding` + "`" + ` can be used per role. Note that custom roles must be of the format ` + "`" + `[projects|organizations]/{parent-name}/roles/{role-name}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy_data",
					Description: `(Required only by ` + "`" + `google_sourcerepo_repository_iam_policy` + "`" + `) The policy data generated by a ` + "`" + `google_iam_policy` + "`" + ` data source. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the repository's IAM policy. ## Import SourceRepo repository IAM resources can be imported using the project, repository name, role and member. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_sourcerepo_repository_iam_policy.editor {{project}}/{{repository}} $ terraform import google_sourcerepo_repository_iam_binding.editor "{{project}}/{{repository}} roles/editor" $ terraform import google_sourcerepo_repository_iam_member.editor "{{project}}/{{repository}} roles/editor jane@example.com" ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the repository's IAM policy. ## Import SourceRepo repository IAM resources can be imported using the project, repository name, role and member. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_sourcerepo_repository_iam_policy.editor {{project}}/{{repository}} $ terraform import google_sourcerepo_repository_iam_binding.editor "{{project}}/{{repository}} roles/editor" $ terraform import google_sourcerepo_repository_iam_member.editor "{{project}}/{{repository}} roles/editor jane@example.com" ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_spanner_database",
			Category:         "Google Spanner Resources",
			ShortDescription: `A Cloud Spanner Database which is hosted on a Spanner instance.`,
			Description:      ``,
			Keywords: []string{
				"spanner",
				"database",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique identifier for the database, which cannot be changed after the instance is created. Values are of the form [a-z][-a-z0-9]`,
				},
				resource.Attribute{
					Name:        "instance",
					Description: `(Required) The instance to create the database on. - - -`,
				},
				resource.Attribute{
					Name:        "ddl",
					Description: `(Optional) An optional list of DDL statements to run inside the newly created database. Statements can create tables, indexes, etc. These statements execute atomically with the creation of the database: if there is an error in any statement, the database is not created.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `An explanation of the status of the database. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Database can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_spanner_database.default projects/{{project}}/instances/{{instance}}/databases/{{name}} $ terraform import google_spanner_database.default instances/{{instance}}/databases/{{name}} $ terraform import google_spanner_database.default {{project}}/{{instance}}/{{name}} $ terraform import google_spanner_database.default {{instance}}/{{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "state",
					Description: `An explanation of the status of the database. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Database can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_spanner_database.default projects/{{project}}/instances/{{instance}}/databases/{{name}} $ terraform import google_spanner_database.default instances/{{instance}}/databases/{{name}} $ terraform import google_spanner_database.default {{project}}/{{instance}}/{{name}} $ terraform import google_spanner_database.default {{instance}}/{{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_spanner_database_iam_policy",
			Category:         "Google Spanner Resources",
			ShortDescription: `Collection of resources to manage IAM policy for a Spanner database.`,
			Description:      ``,
			Keywords: []string{
				"spanner",
				"database",
				"iam",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "database",
					Description: `(Required) The name of the Spanner database.`,
				},
				resource.Attribute{
					Name:        "instance",
					Description: `(Required) The name of the Spanner instance the database belongs to.`,
				},
				resource.Attribute{
					Name:        "member/members",
					Description: `(Required) Identities that will be granted the privilege in ` + "`" + `role` + "`" + `. Each entry can have one of the following values:`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. Only one ` + "`" + `google_spanner_database_iam_binding` + "`" + ` can be used per role. Note that custom roles must be of the format ` + "`" + `[projects|organizations]/{parent-name}/roles/{role-name}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy_data",
					Description: `(Required only by ` + "`" + `google_spanner_database_iam_policy` + "`" + `) The policy data generated by a ` + "`" + `google_iam_policy` + "`" + ` data source.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the database's IAM policy. ## Import For all import syntaxes, the "resource in question" can take any of the following forms:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the database's IAM policy. ## Import For all import syntaxes, the "resource in question" can take any of the following forms:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_spanner_instance",
			Category:         "Google Spanner Resources",
			ShortDescription: `An isolated set of Cloud Spanner resources on which databases can be hosted.`,
			Description:      ``,
			Keywords: []string{
				"spanner",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique identifier for the instance, which cannot be changed after the instance is created. The name must be between 6 and 30 characters in length. If not provided, a random string starting with ` + "`" + `tf-` + "`" + ` will be selected.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `(Required) The name of the instance's configuration (similar but not quite the same as a region) which defines defines the geographic placement and replication of your databases in this instance. It determines where your data is stored. Values are typically of the form ` + "`" + `regional-europe-west1` + "`" + ` , ` + "`" + `us-central` + "`" + ` etc. In order to obtain a valid list please consult the [Configuration section of the docs](https://cloud.google.com/spanner/docs/instances).`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The descriptive name for this instance as it appears in UIs. Must be unique per project and between 4 and 30 characters in length. - - -`,
				},
				resource.Attribute{
					Name:        "num_nodes",
					Description: `(Optional) The number of nodes allocated to this instance.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Instance status: ` + "`" + `CREATING` + "`" + ` or ` + "`" + `READY` + "`" + `. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Instance can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_spanner_instance.default projects/{{project}}/instances/{{name}} $ terraform import google_spanner_instance.default {{project}}/{{name}} $ terraform import google_spanner_instance.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "state",
					Description: `Instance status: ` + "`" + `CREATING` + "`" + ` or ` + "`" + `READY` + "`" + `. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import Instance can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_spanner_instance.default projects/{{project}}/instances/{{name}} $ terraform import google_spanner_instance.default {{project}}/{{name}} $ terraform import google_spanner_instance.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_spanner_instance_iam_policy",
			Category:         "Google Spanner Resources",
			ShortDescription: `Collection of resources to manage IAM policy for a Spanner instance.`,
			Description:      ``,
			Keywords: []string{
				"spanner",
				"instance",
				"iam",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance",
					Description: `(Required) The name of the instance.`,
				},
				resource.Attribute{
					Name:        "member/members",
					Description: `(Required) Identities that will be granted the privilege in ` + "`" + `role` + "`" + `. Each entry can have one of the following values:`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. Only one ` + "`" + `google_spanner_instance_iam_binding` + "`" + ` can be used per role. Note that custom roles must be of the format ` + "`" + `[projects|organizations]/{parent-name}/roles/{role-name}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy_data",
					Description: `(Required only by ` + "`" + `google_spanner_instance_iam_policy` + "`" + `) The policy data generated by a ` + "`" + `google_iam_policy` + "`" + ` data source.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the instance's IAM policy. ## Import For all import syntaxes, the "resource in question" can take any of the following forms:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the instance's IAM policy. ## Import For all import syntaxes, the "resource in question" can take any of the following forms:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_sql_database",
			Category:         "Google SQL Resources",
			ShortDescription: `Represents a SQL database inside the Cloud SQL instance, hosted in Google's cloud.`,
			Description:      ``,
			Keywords: []string{
				"sql",
				"database",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the database in the Cloud SQL instance. This does not include the project ID or instance name.`,
				},
				resource.Attribute{
					Name:        "instance",
					Description: `(Required) The name of the Cloud SQL instance. This does not include the project ID. - - -`,
				},
				resource.Attribute{
					Name:        "charset",
					Description: `(Optional) The charset value. See MySQL's [Supported Character Sets and Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html) and Postgres' [Character Set Support](https://www.postgresql.org/docs/9.6/static/multibyte.html) for more details and supported values. Postgres databases only support a value of ` + "`" + `UTF8` + "`" + ` at creation time.`,
				},
				resource.Attribute{
					Name:        "collation",
					Description: `(Optional) The collation value. See MySQL's [Supported Character Sets and Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html) and Postgres' [Collation Support](https://www.postgresql.org/docs/9.6/static/collation.html) for more details and supported values. Postgres databases only support a value of ` + "`" + `en_US.UTF8` + "`" + ` at creation time.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 15 minutes. - ` + "`" + `update` + "`" + ` - Default is 10 minutes. - ` + "`" + `delete` + "`" + ` - Default is 10 minutes. ## Import Database can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_sql_database.default projects/{{project}}/instances/{{instance}}/databases/{{name}} $ terraform import google_sql_database.default instances/{{instance}}/databases/{{name}} $ terraform import google_sql_database.default {{project}}/{{instance}}/{{name}} $ terraform import google_sql_database.default {{instance}}/{{name}} $ terraform import google_sql_database.default {{instance}}:{{name}} $ terraform import google_sql_database.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_sql_database_instance",
			Category:         "Google SQL Resources",
			ShortDescription: `Creates a new SQL database instance in Google Cloud SQL.`,
			Description:      ``,
			Keywords: []string{
				"sql",
				"database",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region the instance will sit in. Note, first-generation Cloud SQL instance regions do not line up with the Google Compute Engine (GCE) regions, and Cloud SQL is not available in all regions - choose from one of the options listed [here](https://cloud.google.com/sql/docs/mysql/instance-locations). A valid region must be provided to use this resource. If a region is not provided in the resource definition, the provider region will be used instead, but this will be an apply-time error for all first-generation instances`,
				},
				resource.Attribute{
					Name:        "settings",
					Description: `(Required) The settings to use for the database. The configuration is detailed below. - - -`,
				},
				resource.Attribute{
					Name:        "database_version",
					Description: `(Optional, Default: ` + "`" + `MYSQL_5_6` + "`" + `) The MySQL or PostgreSQL version to use. Can be ` + "`" + `MYSQL_5_6` + "`" + `, ` + "`" + `MYSQL_5_7` + "`" + `, ` + "`" + `POSTGRES_9_6` + "`" + ` or ` + "`" + `POSTGRES_11` + "`" + ` (beta) for second-generation instances, or ` + "`" + `MYSQL_5_5` + "`" + ` or ` + "`" + `MYSQL_5_6` + "`" + ` for first-generation instances. See [Second Generation Capabilities](https://cloud.google.com/sql/docs/1st-2nd-gen-differences) for more information.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, Computed) The name of the instance. If the name is left blank, Terraform will randomly generate one when the instance is first created. This is done because after a name is used, it cannot be reused for up to [one week](https://cloud.google.com/sql/docs/delete-instance).`,
				},
				resource.Attribute{
					Name:        "master_instance_name",
					Description: `(Optional) The name of the instance that will act as the master in the replication setup. Note, this requires the master to have ` + "`" + `binary_log_enabled` + "`" + ` set, as well as existing backups.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "replica_configuration",
					Description: `(Optional) The configuration for replication. The configuration is detailed below. The required ` + "`" + `settings` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "tier",
					Description: `(Required) The machine tier (First Generation) or type (Second Generation) to use. See [tiers](https://cloud.google.com/sql/docs/admin-api/v1beta4/tiers) for more details and supported versions. Postgres supports only shared-core machine types such as ` + "`" + `db-f1-micro` + "`" + `, and custom machine types such as ` + "`" + `db-custom-2-13312` + "`" + `. See the [Custom Machine Type Documentation](https://cloud.google.com/compute/docs/instances/creating-instance-with-custom-machine-type#create) to learn about specifying custom machine types.`,
				},
				resource.Attribute{
					Name:        "activation_policy",
					Description: `(Optional) This specifies when the instance should be active. Can be either ` + "`" + `ALWAYS` + "`" + `, ` + "`" + `NEVER` + "`" + ` or ` + "`" + `ON_DEMAND` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "authorized_gae_applications",
					Description: `(Optional) A list of Google App Engine (GAE) project names that are allowed to access this instance.`,
				},
				resource.Attribute{
					Name:        "availability_type",
					Description: `(Optional) This specifies whether a PostgreSQL instance should be set up for high availability (` + "`" + `REGIONAL` + "`" + `) or single zone (` + "`" + `ZONAL` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "crash_safe_replication",
					Description: `(Optional) Specific to read instances, indicates when crash-safe replication flags are enabled.`,
				},
				resource.Attribute{
					Name:        "disk_autoresize",
					Description: `(Optional, Second Generation, Default: ` + "`" + `true` + "`" + `) Configuration to increase storage size automatically.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional, Second Generation, Default: ` + "`" + `10` + "`" + `) The size of data disk, in GB. Size of a running instance cannot be reduced but can be increased.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Optional, Second Generation, Default: ` + "`" + `PD_SSD` + "`" + `) The type of data disk: PD_SSD or PD_HDD.`,
				},
				resource.Attribute{
					Name:        "pricing_plan",
					Description: `(Optional, First Generation) Pricing plan for this instance, can be one of ` + "`" + `PER_USE` + "`" + ` or ` + "`" + `PACKAGE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "replication_type",
					Description: `(Optional) Replication type for this instance, can be one of ` + "`" + `ASYNCHRONOUS` + "`" + ` or ` + "`" + `SYNCHRONOUS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_labels",
					Description: `(Optional) A set of key/value user label pairs to assign to the instance. The optional ` + "`" + `settings.database_flags` + "`" + ` sublist supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the flag.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) Value of the flag. The optional ` + "`" + `settings.backup_configuration` + "`" + ` subblock supports:`,
				},
				resource.Attribute{
					Name:        "binary_log_enabled",
					Description: `(Optional) True if binary logging is enabled. If ` + "`" + `settings.backup_configuration.enabled` + "`" + ` is false, this must be as well. Cannot be used with Postgres.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) True if backup configuration is enabled.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Optional) ` + "`" + `HH:MM` + "`" + ` format time indicating when backup configuration starts. The optional ` + "`" + `settings.ip_configuration` + "`" + ` subblock supports:`,
				},
				resource.Attribute{
					Name:        "ipv4_enabled",
					Description: `(Optional) Whether this Cloud SQL instance should be assigned a public IPV4 address. Either ` + "`" + `ipv4_enabled` + "`" + ` must be enabled or a ` + "`" + `private_network` + "`" + ` must be configured.`,
				},
				resource.Attribute{
					Name:        "private_network",
					Description: `(Optional) The VPC network from which the Cloud SQL instance is accessible for private IP. For example, /projects/myProject/global/networks/default. Specifying a network enables private IP. Either ` + "`" + `ipv4_enabled` + "`" + ` must be enabled or a ` + "`" + `private_network` + "`" + ` must be configured. This setting can be updated, but it cannot be removed after it is set.`,
				},
				resource.Attribute{
					Name:        "require_ssl",
					Description: `(Optional) True if mysqld should default to ` + "`" + `REQUIRE X509` + "`" + ` for users connecting over IP. The optional ` + "`" + `settings.ip_configuration.authorized_networks[]` + "`" + ` sublist supports:`,
				},
				resource.Attribute{
					Name:        "expiration_time",
					Description: `(Optional) The [RFC 3339](https://tools.ietf.org/html/rfc3339) formatted date time string indicating when this whitelist expires.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A name for this whitelist entry.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) A CIDR notation IPv4 or IPv6 address that is allowed to access this instance. Must be set even if other two attributes are not for the whitelist to become active. The optional ` + "`" + `settings.location_preference` + "`" + ` subblock supports:`,
				},
				resource.Attribute{
					Name:        "follow_gae_application",
					Description: `(Optional) A GAE application whose zone to remain in. Must be in the same region as this instance.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) The preferred compute engine [zone](https://cloud.google.com/compute/docs/zones?hl=en). The optional ` + "`" + `settings.maintenance_window` + "`" + ` subblock for Second Generation instances declares a one-hour [maintenance window](https://cloud.google.com/sql/docs/instance-settings?hl=en#maintenance-window-2ndgen) when an Instance can automatically restart to apply updates. The maintenance window is specified in UTC time. It supports:`,
				},
				resource.Attribute{
					Name:        "day",
					Description: `(Optional) Day of week (` + "`" + `1-7` + "`" + `), starting on Monday`,
				},
				resource.Attribute{
					Name:        "hour",
					Description: `(Optional) Hour of day (` + "`" + `0-23` + "`" + `), ignored if ` + "`" + `day` + "`" + ` not set`,
				},
				resource.Attribute{
					Name:        "update_track",
					Description: `(Optional) Receive updates earlier (` + "`" + `canary` + "`" + `) or later (` + "`" + `stable` + "`" + `) The optional ` + "`" + `replica_configuration` + "`" + ` block must have ` + "`" + `master_instance_name` + "`" + ` set to work, cannot be updated, and supports:`,
				},
				resource.Attribute{
					Name:        "ca_certificate",
					Description: `(Optional) PEM representation of the trusted CA's x509 certificate.`,
				},
				resource.Attribute{
					Name:        "client_certificate",
					Description: `(Optional) PEM representation of the slave's x509 certificate.`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `(Optional) PEM representation of the slave's private key. The corresponding public key in encoded in the ` + "`" + `client_certificate` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "connect_retry_interval",
					Description: `(Optional, Default: 60) The number of seconds between connect retries.`,
				},
				resource.Attribute{
					Name:        "dump_file_path",
					Description: `(Optional) Path to a SQL file in GCS from which slave instances are created. Format is ` + "`" + `gs://bucket/filename` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "failover_target",
					Description: `(Optional) Specifies if the replica is the failover target. If the field is set to true the replica will be designated as a failover replica. If the master instance fails, the replica instance will be promoted as the new master instance.`,
				},
				resource.Attribute{
					Name:        "master_heartbeat_period",
					Description: `(Optional) Time in ms between replication heartbeats.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password for the replication connection.`,
				},
				resource.Attribute{
					Name:        "sslCipher",
					Description: `(Optional) Permissible ciphers for use in SSL encryption.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) Username for replication connection.`,
				},
				resource.Attribute{
					Name:        "verify_server_certificate",
					Description: `(Optional) True if the master's common name value is checked during the SSL handshake. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource.`,
				},
				resource.Attribute{
					Name:        "connection_name",
					Description: `The connection name of the instance to be used in connection strings. For example, when connecting with [Cloud SQL Proxy](https://cloud.google.com/sql/docs/mysql/connect-admin-proxy).`,
				},
				resource.Attribute{
					Name:        "service_account_email_address",
					Description: `The service account email address assigned to the instance. This property is applicable only to Second Generation instances.`,
				},
				resource.Attribute{
					Name:        "ip_address.0.ip_address",
					Description: `The IPv4 address assigned.`,
				},
				resource.Attribute{
					Name:        "ip_address.0.time_to_retire",
					Description: `The time this IP address will be retired, in RFC 3339 format.`,
				},
				resource.Attribute{
					Name:        "ip_address.0.type",
					Description: `The type of this IP address.`,
				},
				resource.Attribute{
					Name:        "first_ip_address",
					Description: `The first IPv4 address of any type assigned. This is to support accessing the [first address in the list in a terraform output](https://github.com/terraform-providers/terraform-provider-google/issues/912) when the resource is configured with a ` + "`" + `count` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "public_ip_address",
					Description: `The first public (` + "`" + `PRIMARY` + "`" + `) IPv4 address assigned. This is a workaround for an [issue fixed in Terraform 0.12](https://github.com/hashicorp/terraform/issues/17048) but also provides a convenient way to access an IP of a specific type without performing filtering in a Terraform config.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `The first private (` + "`" + `PRIVATE` + "`" + `) IPv4 address assigned. This is a workaround for an [issue fixed in Terraform 0.12](https://github.com/hashicorp/terraform/issues/17048) but also provides a convenient way to access an IP of a specific type without performing filtering in a Terraform config.`,
				},
				resource.Attribute{
					Name:        "settings.version",
					Description: `Used to make sure changes to the ` + "`" + `settings` + "`" + ` block are atomic.`,
				},
				resource.Attribute{
					Name:        "server_ca_cert.0.cert",
					Description: `The CA Certificate used to connect to the SQL Instance via SSL.`,
				},
				resource.Attribute{
					Name:        "server_ca_cert.0.common_name",
					Description: `The CN valid for the CA Cert.`,
				},
				resource.Attribute{
					Name:        "server_ca_cert.0.create_time",
					Description: `Creation time of the CA Cert.`,
				},
				resource.Attribute{
					Name:        "server_ca_cert.0.expiration_time",
					Description: `Expiration time of the CA Cert.`,
				},
				resource.Attribute{
					Name:        "server_ca_cert.0.sha1_fingerprint",
					Description: `SHA Fingerprint of the CA Cert. ## Timeouts ` + "`" + `google_sql_database_instance` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 20 minutes. - ` + "`" + `update` + "`" + ` - Default is 20 minutes. - ` + "`" + `delete` + "`" + ` - Default is 20 minutes. ## Import Database instances can be imported using one of any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_sql_database_instance.master projects/{{project}}/instances/{{name}} $ terraform import google_sql_database_instance.master {{project}}/{{name}} $ terraform import google_sql_database_instance.master {{name}} ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource.`,
				},
				resource.Attribute{
					Name:        "connection_name",
					Description: `The connection name of the instance to be used in connection strings. For example, when connecting with [Cloud SQL Proxy](https://cloud.google.com/sql/docs/mysql/connect-admin-proxy).`,
				},
				resource.Attribute{
					Name:        "service_account_email_address",
					Description: `The service account email address assigned to the instance. This property is applicable only to Second Generation instances.`,
				},
				resource.Attribute{
					Name:        "ip_address.0.ip_address",
					Description: `The IPv4 address assigned.`,
				},
				resource.Attribute{
					Name:        "ip_address.0.time_to_retire",
					Description: `The time this IP address will be retired, in RFC 3339 format.`,
				},
				resource.Attribute{
					Name:        "ip_address.0.type",
					Description: `The type of this IP address.`,
				},
				resource.Attribute{
					Name:        "first_ip_address",
					Description: `The first IPv4 address of any type assigned. This is to support accessing the [first address in the list in a terraform output](https://github.com/terraform-providers/terraform-provider-google/issues/912) when the resource is configured with a ` + "`" + `count` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "public_ip_address",
					Description: `The first public (` + "`" + `PRIMARY` + "`" + `) IPv4 address assigned. This is a workaround for an [issue fixed in Terraform 0.12](https://github.com/hashicorp/terraform/issues/17048) but also provides a convenient way to access an IP of a specific type without performing filtering in a Terraform config.`,
				},
				resource.Attribute{
					Name:        "private_ip_address",
					Description: `The first private (` + "`" + `PRIVATE` + "`" + `) IPv4 address assigned. This is a workaround for an [issue fixed in Terraform 0.12](https://github.com/hashicorp/terraform/issues/17048) but also provides a convenient way to access an IP of a specific type without performing filtering in a Terraform config.`,
				},
				resource.Attribute{
					Name:        "settings.version",
					Description: `Used to make sure changes to the ` + "`" + `settings` + "`" + ` block are atomic.`,
				},
				resource.Attribute{
					Name:        "server_ca_cert.0.cert",
					Description: `The CA Certificate used to connect to the SQL Instance via SSL.`,
				},
				resource.Attribute{
					Name:        "server_ca_cert.0.common_name",
					Description: `The CN valid for the CA Cert.`,
				},
				resource.Attribute{
					Name:        "server_ca_cert.0.create_time",
					Description: `Creation time of the CA Cert.`,
				},
				resource.Attribute{
					Name:        "server_ca_cert.0.expiration_time",
					Description: `Expiration time of the CA Cert.`,
				},
				resource.Attribute{
					Name:        "server_ca_cert.0.sha1_fingerprint",
					Description: `SHA Fingerprint of the CA Cert. ## Timeouts ` + "`" + `google_sql_database_instance` + "`" + ` provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 20 minutes. - ` + "`" + `update` + "`" + ` - Default is 20 minutes. - ` + "`" + `delete` + "`" + ` - Default is 20 minutes. ## Import Database instances can be imported using one of any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_sql_database_instance.master projects/{{project}}/instances/{{name}} $ terraform import google_sql_database_instance.master {{project}}/{{name}} $ terraform import google_sql_database_instance.master {{name}} ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_sql_ssl_cert",
			Category:         "Google SQL Resources",
			ShortDescription: `Creates a new SQL Ssl Cert in Google Cloud SQL.`,
			Description:      ``,
			Keywords: []string{
				"sql",
				"ssl",
				"cert",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance",
					Description: `(Required) The name of the Cloud SQL instance. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "common_name",
					Description: `(Required) The common name to be used in the certificate to identify the client. Constrained to [a-zA-Z.-_ ]+. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "sha1_fingerprint",
					Description: `The SHA1 Fingerprint of the certificate.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The private key associated with the client certificate.`,
				},
				resource.Attribute{
					Name:        "server_ca_cert",
					Description: `The CA cert of the server this client cert was generated from.`,
				},
				resource.Attribute{
					Name:        "cert",
					Description: `The actual certificate data for this client certificate.`,
				},
				resource.Attribute{
					Name:        "cert_serial_number",
					Description: `The serial number extracted from the certificate data.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the certificate was created in RFC 3339 format, for example 2012-11-15T16:19:00.094Z.`,
				},
				resource.Attribute{
					Name:        "expiration_time",
					Description: `The time when the certificate expires in RFC 3339 format, for example 2012-11-15T16:19:00.094Z. ## Import Since the contents of the certificate cannot be accessed after its creation, this resource cannot be imported.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "sha1_fingerprint",
					Description: `The SHA1 Fingerprint of the certificate.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The private key associated with the client certificate.`,
				},
				resource.Attribute{
					Name:        "server_ca_cert",
					Description: `The CA cert of the server this client cert was generated from.`,
				},
				resource.Attribute{
					Name:        "cert",
					Description: `The actual certificate data for this client certificate.`,
				},
				resource.Attribute{
					Name:        "cert_serial_number",
					Description: `The serial number extracted from the certificate data.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The time when the certificate was created in RFC 3339 format, for example 2012-11-15T16:19:00.094Z.`,
				},
				resource.Attribute{
					Name:        "expiration_time",
					Description: `The time when the certificate expires in RFC 3339 format, for example 2012-11-15T16:19:00.094Z. ## Import Since the contents of the certificate cannot be accessed after its creation, this resource cannot be imported.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_sql_user",
			Category:         "Google SQL Resources",
			ShortDescription: `Creates a new SQL user in Google Cloud SQL.`,
			Description:      ``,
			Keywords: []string{
				"sql",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance",
					Description: `(Required) The name of the Cloud SQL instance. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the user. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) The password for the user. Can be updated. - - -`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) The host the user can connect from. This is only supported for MySQL instances. Don't set this field for PostgreSQL instances. Can be an IP address. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. ## Attributes Reference Only the arguments listed above are exposed as attributes. ## Import SQL users for MySQL databases can be imported using the ` + "`" + `project` + "`" + `, ` + "`" + `instance` + "`" + `, ` + "`" + `host` + "`" + ` and ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_sql_user.users my-project/master-instance/my-domain.com/me ` + "`" + `` + "`" + `` + "`" + ` SQL users for PostgreSQL databases can be imported using the ` + "`" + `project` + "`" + `, ` + "`" + `instance` + "`" + ` and ` + "`" + `name` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_sql_user.users my-project/master-instance/me ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_storage_bucket",
			Category:         "Google Storage Resources",
			ShortDescription: `Creates a new bucket in Google Cloud Storage.`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"bucket",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the bucket. - - -`,
				},
				resource.Attribute{
					Name:        "force_destroy",
					Description: `(Optional, Default: false) When deleting a bucket, this boolean option will delete all contained objects. If you try to delete a bucket that contains objects, Terraform will fail that run.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Optional, Default: 'US') The [GCS location](https://cloud.google.com/storage/docs/bucket-locations)`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `(Optional) The [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of the new bucket. Supported values include: ` + "`" + `MULTI_REGIONAL` + "`" + `, ` + "`" + `REGIONAL` + "`" + `, ` + "`" + `NEARLINE` + "`" + `, ` + "`" + `COLDLINE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lifecycle_rule",
					Description: `(Optional) The bucket's [Lifecycle Rules](https://cloud.google.com/storage/docs/lifecycle#configuration) configuration. Multiple blocks of this type are permitted. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "versioning",
					Description: `(Optional) The bucket's [Versioning](https://cloud.google.com/storage/docs/object-versioning) configuration.`,
				},
				resource.Attribute{
					Name:        "website",
					Description: `(Optional) Configuration if the bucket acts as a website. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "cors",
					Description: `(Optional) The bucket's [Cross-Origin Resource Sharing (CORS)](https://www.w3.org/TR/cors/) configuration. Multiple blocks of this type are permitted. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "retention_policy",
					Description: `(Optional) Configuration of the bucket's data retention policy for how long objects in the bucket should be retained. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A set of key/value label pairs to assign to the bucket.`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `(Optional) The bucket's [Access & Storage Logs](https://cloud.google.com/storage/docs/access-logs) configuration.`,
				},
				resource.Attribute{
					Name:        "encryption",
					Description: `(Optional) The bucket's encryption configuration.`,
				},
				resource.Attribute{
					Name:        "requester_pays",
					Description: `(Optional, Default: false) Enables [Requester Pays](https://cloud.google.com/storage/docs/requester-pays) on a storage bucket.`,
				},
				resource.Attribute{
					Name:        "bucket_policy_only",
					Description: `(Optional, Default: false) Enables [Bucket Policy Only](https://cloud.google.com/storage/docs/bucket-policy-only) access to a bucket. The ` + "`" + `lifecycle_rule` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) The Lifecycle Rule's action configuration. A single block of this type is supported. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `(Required) The Lifecycle Rule's condition configuration. A single block of this type is supported. Structure is documented below. The ` + "`" + `action` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the action of this Lifecycle Rule. Supported values include: ` + "`" + `Delete` + "`" + ` and ` + "`" + `SetStorageClass` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `(Required if action type is ` + "`" + `SetStorageClass` + "`" + `) The target [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of objects affected by this Lifecycle Rule. Supported values include: ` + "`" + `MULTI_REGIONAL` + "`" + `, ` + "`" + `REGIONAL` + "`" + `, ` + "`" + `NEARLINE` + "`" + `, ` + "`" + `COLDLINE` + "`" + `. The ` + "`" + `condition` + "`" + ` block supports the following elements, and requires at least one to be defined:`,
				},
				resource.Attribute{
					Name:        "age",
					Description: `(Optional) Minimum age of an object in days to satisfy this condition.`,
				},
				resource.Attribute{
					Name:        "created_before",
					Description: `(Optional) Creation date of an object in RFC 3339 (e.g. ` + "`" + `2017-06-13` + "`" + `) to satisfy this condition.`,
				},
				resource.Attribute{
					Name:        "with_state",
					Description: `(Optional) Match to live and/or archived objects. Unversioned buckets have only live objects. Supported values include: ` + "`" + `"LIVE"` + "`" + `, ` + "`" + `"ARCHIVED"` + "`" + `, ` + "`" + `"ANY"` + "`" + `. Unset or empty strings will be treated as ` + "`" + `ARCHIVED` + "`" + ` to maintain backwards compatibility with ` + "`" + `is_live` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "is_live",
					Description: `(Optional, Deprecated) Defaults to ` + "`" + `false` + "`" + ` to match archived objects. If ` + "`" + `true` + "`" + `, this condition matches live objects. Unversioned buckets have only live objects.`,
				},
				resource.Attribute{
					Name:        "matches_storage_class",
					Description: `(Optional) [Storage Class](https://cloud.google.com/storage/docs/storage-classes) of objects to satisfy this condition. Supported values include: ` + "`" + `MULTI_REGIONAL` + "`" + `, ` + "`" + `REGIONAL` + "`" + `, ` + "`" + `NEARLINE` + "`" + `, ` + "`" + `COLDLINE` + "`" + `, ` + "`" + `STANDARD` + "`" + `, ` + "`" + `DURABLE_REDUCED_AVAILABILITY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "num_newer_versions",
					Description: `(Optional) Relevant only for versioned objects. The number of newer versions of an object to satisfy this condition. The ` + "`" + `versioning` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) While set to ` + "`" + `true` + "`" + `, versioning is fully enabled for this bucket. The ` + "`" + `website` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "main_page_suffix",
					Description: `(Optional) Behaves as the bucket's directory index where missing objects are treated as potential directories.`,
				},
				resource.Attribute{
					Name:        "not_found_page",
					Description: `(Optional) The custom object to return when a requested resource is not found. The ` + "`" + `cors` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `(Optional) The list of [Origins](https://tools.ietf.org/html/rfc6454) eligible to receive CORS response headers. Note: "`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Optional) The list of HTTP methods on which to include CORS response headers, (GET, OPTIONS, POST, etc) Note: "`,
				},
				resource.Attribute{
					Name:        "response_header",
					Description: `(Optional) The list of HTTP headers other than the [simple response headers](https://www.w3.org/TR/cors/#simple-response-header) to give permission for the user-agent to share across domains.`,
				},
				resource.Attribute{
					Name:        "max_age_seconds",
					Description: `(Optional) The value, in seconds, to return in the [Access-Control-Max-Age header](https://www.w3.org/TR/cors/#access-control-max-age-response-header) used in preflight responses. The ` + "`" + `retention_policy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "is_locked",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + `, the bucket will be [locked](https://cloud.google.com/storage/docs/using-bucket-lock#lock-bucket) and permanently restrict edits to the bucket's retention policy. Caution: Locking a bucket is an irreversible action.`,
				},
				resource.Attribute{
					Name:        "retention_period",
					Description: `(Optional) The period of time, in seconds, that objects in the bucket must be retained and cannot be deleted, overwritten, or archived. The value must be less than 3,155,760,000 seconds. The ` + "`" + `logging` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "log_bucket",
					Description: `(Required) The bucket that will receive log objects.`,
				},
				resource.Attribute{
					Name:        "log_object_prefix",
					Description: `(Optional, Computed) The object prefix for log objects. If it's not provided, by default GCS sets this to this bucket's name. The ` + "`" + `encryption` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The base URL of the bucket, in the format ` + "`" + `gs://<bucket-name>` + "`" + `. ## Import Storage buckets can be imported using the ` + "`" + `name` + "`" + ` or ` + "`" + `project/name` + "`" + `. If the project is not passed to the import command it will be inferred from the provider block or environment variables. If it cannot be inferred it will be queried from the Compute API (this will fail if the API is not enabled). e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_storage_bucket.image-store image-store-bucket $ terraform import google_storage_bucket.image-store tf-test-project/image-store-bucket ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The base URL of the bucket, in the format ` + "`" + `gs://<bucket-name>` + "`" + `. ## Import Storage buckets can be imported using the ` + "`" + `name` + "`" + ` or ` + "`" + `project/name` + "`" + `. If the project is not passed to the import command it will be inferred from the provider block or environment variables. If it cannot be inferred it will be queried from the Compute API (this will fail if the API is not enabled). e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_storage_bucket.image-store image-store-bucket $ terraform import google_storage_bucket.image-store tf-test-project/image-store-bucket ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_storage_bucket_acl",
			Category:         "Google Storage Resources",
			ShortDescription: `Creates a new bucket ACL in Google Cloud Storage.`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"bucket",
				"acl",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) The name of the bucket it applies to. - - -`,
				},
				resource.Attribute{
					Name:        "predefined_acl",
					Description: `(Optional) The [canned GCS ACL](https://cloud.google.com/storage/docs/access-control/lists#predefined-acl) to apply. Must be set if ` + "`" + `role_entity` + "`" + ` is not.`,
				},
				resource.Attribute{
					Name:        "role_entity",
					Description: `(Optional) List of role/entity pairs in the form ` + "`" + `ROLE:entity` + "`" + `. See [GCS Bucket ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/bucketAccessControls) for more details. Must be set if ` + "`" + `predefined_acl` + "`" + ` is not.`,
				},
				resource.Attribute{
					Name:        "default_acl",
					Description: `(Optional) Configure this ACL to be the default ACL. ## Attributes Reference Only the arguments listed above are exposed as attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_storage_bucket_iam_member",
			Category:         "Google Storage Resources",
			ShortDescription: `Collection of resources to manage IAM policy for a Google storage bucket.`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"bucket",
				"iam",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) The name of the bucket it applies to.`,
				},
				resource.Attribute{
					Name:        "member/members",
					Description: `(Required) Identities that will be granted the privilege in ` + "`" + `role` + "`" + `. Each entry can have one of the following values:`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role that should be applied. Note that custom roles must be of the format ` + "`" + `[projects|organizations]/{parent-name}/roles/{role-name}` + "`" + `. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the storage bucket's IAM policy. ## Import For ` + "`" + `google_storage_bucket_iam_policy` + "`" + `: IAM member imports use space-delimited identifiers - generally the resource in question, the role, and the member identity (i.e. ` + "`" + `serviceAccount: my-sa@my-project.iam.gserviceaccount.com` + "`" + ` or ` + "`" + `user:foo@example.com` + "`" + `). Policies, bindings, and members can be respectively imported as follows: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_storage_bucket_iam_policy.policy "my-bucket user:foo@example.com" $ terraform import google_storage_bucket_iam_binding.binding "my-bucket roles/my-role " $ terraform import google_storage_bucket_iam_member.member "my-bucket roles/my-role user:foo@example.com" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "etag",
					Description: `(Computed) The etag of the storage bucket's IAM policy. ## Import For ` + "`" + `google_storage_bucket_iam_policy` + "`" + `: IAM member imports use space-delimited identifiers - generally the resource in question, the role, and the member identity (i.e. ` + "`" + `serviceAccount: my-sa@my-project.iam.gserviceaccount.com` + "`" + ` or ` + "`" + `user:foo@example.com` + "`" + `). Policies, bindings, and members can be respectively imported as follows: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_storage_bucket_iam_policy.policy "my-bucket user:foo@example.com" $ terraform import google_storage_bucket_iam_binding.binding "my-bucket roles/my-role " $ terraform import google_storage_bucket_iam_member.member "my-bucket roles/my-role user:foo@example.com" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_storage_bucket_object",
			Category:         "Google Storage Resources",
			ShortDescription: `Creates a new object inside a specified bucket`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"bucket",
				"object",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) The name of the containing bucket.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the object. If you're interpolating the name of this object, see ` + "`" + `output_name` + "`" + ` instead. One of the following is required:`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional, Sensitive) Data as ` + "`" + `string` + "`" + ` to be uploaded. Must be defined if ` + "`" + `source` + "`" + ` is not.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) A path to the data you want to upload. Must be defined if ` + "`" + `content` + "`" + ` is not. - - -`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `(Optional) [Cache-Control](https://tools.ietf.org/html/rfc7234#section-5.2) directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `(Optional) [Content-Disposition](https://tools.ietf.org/html/rfc6266) of the object data.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `(Optional) [Content-Encoding](https://tools.ietf.org/html/rfc7231#section-3.1.2.2) of the object data.`,
				},
				resource.Attribute{
					Name:        "content_language",
					Description: `(Optional) [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object data.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `(Optional) [Content-Type](https://tools.ietf.org/html/rfc7231#section-3.1.1.5) of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `(Optional) The [StorageClass](https://cloud.google.com/storage/docs/storage-classes) of the new bucket object. Supported values include: ` + "`" + `MULTI_REGIONAL` + "`" + `, ` + "`" + `REGIONAL` + "`" + `, ` + "`" + `NEARLINE` + "`" + `, ` + "`" + `COLDLINE` + "`" + `. If not provided, this defaults to the bucket's default storage class or to a [standard](https://cloud.google.com/storage/docs/storage-classes#standard) class. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "crc32c",
					Description: `(Computed) Base 64 CRC32 hash of the uploaded data.`,
				},
				resource.Attribute{
					Name:        "md5hash",
					Description: `(Computed) Base 64 MD5 hash of the uploaded data.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `(Computed) A url reference to this object.`,
				},
				resource.Attribute{
					Name:        "output_name",
					Description: `(Computed) The name of the object. Use this field in interpolations with ` + "`" + `google_storage_object_acl` + "`" + ` to recreate ` + "`" + `google_storage_object_acl` + "`" + ` resources when your ` + "`" + `google_storage_bucket_object` + "`" + ` is recreated.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "crc32c",
					Description: `(Computed) Base 64 CRC32 hash of the uploaded data.`,
				},
				resource.Attribute{
					Name:        "md5hash",
					Description: `(Computed) Base 64 MD5 hash of the uploaded data.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `(Computed) A url reference to this object.`,
				},
				resource.Attribute{
					Name:        "output_name",
					Description: `(Computed) The name of the object. Use this field in interpolations with ` + "`" + `google_storage_object_acl` + "`" + ` to recreate ` + "`" + `google_storage_object_acl` + "`" + ` resources when your ` + "`" + `google_storage_bucket_object` + "`" + ` is recreated.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_storage_default_object_access_control",
			Category:         "Google Storage Resources",
			ShortDescription: `The DefaultObjectAccessControls resources represent the Access Control Lists (ACLs) applied to a new object within a Google Cloud Storage bucket when no ACL was provided for that object.`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"default",
				"object",
				"access",
				"control",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) The name of the bucket.`,
				},
				resource.Attribute{
					Name:        "entity",
					Description: `(Required) The entity holding the permission, in one of the following forms:`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The access permission for the entity. - - -`,
				},
				resource.Attribute{
					Name:        "object",
					Description: `(Optional) The name of the object, if applied to an object. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `The domain associated with the entity.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The email address associated with the entity.`,
				},
				resource.Attribute{
					Name:        "entity_id",
					Description: `The ID for the entity`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `The content generation of the object, if applied to an object.`,
				},
				resource.Attribute{
					Name:        "project_team",
					Description: `The project team associated with the entity Structure is documented below. The ` + "`" + `project_team` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "project_number",
					Description: `(Optional) The project team associated with the entity`,
				},
				resource.Attribute{
					Name:        "team",
					Description: `(Optional) The team. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import DefaultObjectAccessControl can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_storage_default_object_access_control.default {{bucket}}/{{entity}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `The domain associated with the entity.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The email address associated with the entity.`,
				},
				resource.Attribute{
					Name:        "entity_id",
					Description: `The ID for the entity`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `The content generation of the object, if applied to an object.`,
				},
				resource.Attribute{
					Name:        "project_team",
					Description: `The project team associated with the entity Structure is documented below. The ` + "`" + `project_team` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "project_number",
					Description: `(Optional) The project team associated with the entity`,
				},
				resource.Attribute{
					Name:        "team",
					Description: `(Optional) The team. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import DefaultObjectAccessControl can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_storage_default_object_access_control.default {{bucket}}/{{entity}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_storage_default_object_acl",
			Category:         "Google Storage Resources",
			ShortDescription: `Authoritatively manages the default object ACLs for a Google Cloud Storage bucket`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"default",
				"object",
				"acl",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) The name of the bucket it applies to. ---`,
				},
				resource.Attribute{
					Name:        "role_entity",
					Description: `(Optional) List of role/entity pairs in the form ` + "`" + `ROLE:entity` + "`" + `. See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details. Omitting the field is the same as providing an empty list. ## Attributes Reference Only the arguments listed above are exposed as attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_storage_notification",
			Category:         "Google Storage Resources",
			ShortDescription: `Creates a new notification configuration on a specified bucket.`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"notification",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) The name of the bucket.`,
				},
				resource.Attribute{
					Name:        "payload_format",
					Description: `(Required) The desired content of the Payload. One of ` + "`" + `"JSON_API_V1"` + "`" + ` or ` + "`" + `"NONE"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "topic",
					Description: `(Required) The Cloud PubSub topic to which this subscription publishes. Expects either the topic name, assumed to belong to the default GCP provider project, or the project-level name, i.e. ` + "`" + `projects/my-gcp-project/topics/my-topic` + "`" + ` or ` + "`" + `my-topic` + "`" + `. - - -`,
				},
				resource.Attribute{
					Name:        "custom_attributes",
					Description: `(Optional) A set of key/value attribute pairs to attach to each Cloud PubSub message published for this notification subscription`,
				},
				resource.Attribute{
					Name:        "event_types",
					Description: `(Optional) List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: ` + "`" + `"OBJECT_FINALIZE"` + "`" + `, ` + "`" + `"OBJECT_METADATA_UPDATE"` + "`" + `, ` + "`" + `"OBJECT_DELETE"` + "`" + `, ` + "`" + `"OBJECT_ARCHIVE"` + "`" + ``,
				},
				resource.Attribute{
					Name:        "object_name_prefix",
					Description: `(Optional) Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Import Storage notifications can be imported using the notification ` + "`" + `id` + "`" + ` in the format ` + "`" + `<bucket_name>/notificationConfigs/<id>` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_storage_notification.notification default_bucket/notificationConfigs/102 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "self_link",
					Description: `The URI of the created resource. ## Import Storage notifications can be imported using the notification ` + "`" + `id` + "`" + ` in the format ` + "`" + `<bucket_name>/notificationConfigs/<id>` + "`" + ` e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_storage_notification.notification default_bucket/notificationConfigs/102 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_storage_object_access_control",
			Category:         "Google Storage Resources",
			ShortDescription: `The ObjectAccessControls resources represent the Access Control Lists (ACLs) for objects within Google Cloud Storage.`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"object",
				"access",
				"control",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) The name of the bucket.`,
				},
				resource.Attribute{
					Name:        "entity",
					Description: `(Required) The entity holding the permission, in one of the following forms:`,
				},
				resource.Attribute{
					Name:        "object",
					Description: `(Required) The name of the object to apply the access control to.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The access permission for the entity. - - - ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `The domain associated with the entity.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The email address associated with the entity.`,
				},
				resource.Attribute{
					Name:        "entity_id",
					Description: `The ID for the entity`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `The content generation of the object, if applied to an object.`,
				},
				resource.Attribute{
					Name:        "project_team",
					Description: `The project team associated with the entity Structure is documented below. The ` + "`" + `project_team` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "project_number",
					Description: `(Optional) The project team associated with the entity`,
				},
				resource.Attribute{
					Name:        "team",
					Description: `(Optional) The team. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import ObjectAccessControl can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_storage_object_access_control.default {{bucket}}/{{object}}/{{entity}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `The domain associated with the entity.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The email address associated with the entity.`,
				},
				resource.Attribute{
					Name:        "entity_id",
					Description: `The ID for the entity`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `The content generation of the object, if applied to an object.`,
				},
				resource.Attribute{
					Name:        "project_team",
					Description: `The project team associated with the entity Structure is documented below. The ` + "`" + `project_team` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "project_number",
					Description: `(Optional) The project team associated with the entity`,
				},
				resource.Attribute{
					Name:        "team",
					Description: `(Optional) The team. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 4 minutes. - ` + "`" + `update` + "`" + ` - Default is 4 minutes. - ` + "`" + `delete` + "`" + ` - Default is 4 minutes. ## Import ObjectAccessControl can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_storage_object_access_control.default {{bucket}}/{{object}}/{{entity}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_storage_object_acl",
			Category:         "Google Storage Resources",
			ShortDescription: `Creates a new object ACL in Google Cloud Storage.`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"object",
				"acl",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) The name of the bucket the object is stored in.`,
				},
				resource.Attribute{
					Name:        "object",
					Description: `(Required) The name of the object to apply the acl to. - - -`,
				},
				resource.Attribute{
					Name:        "predefined_acl",
					Description: `(Optional) The "canned" [predefined ACL](https://cloud.google.com/storage/docs/access-control#predefined-acl) to apply. Must be set if ` + "`" + `role_entity` + "`" + ` is not.`,
				},
				resource.Attribute{
					Name:        "role_entity",
					Description: `(Optional) List of role/entity pairs in the form ` + "`" + `ROLE:entity` + "`" + `. See [GCS Object ACL documentation](https://cloud.google.com/storage/docs/json_api/v1/objectAccessControls) for more details. Must be set if ` + "`" + `predefined_acl` + "`" + ` is not. -> The object's creator will always have ` + "`" + `OWNER` + "`" + ` permissions for their object, and any attempt to modify that permission would return an error. Instead, Terraform automatically adds that role/entity pair to your ` + "`" + `terraform plan` + "`" + ` results when it is omitted in your config; ` + "`" + `terraform plan` + "`" + ` will show the correct final state at every point except for at ` + "`" + `Create` + "`" + ` time, where the object role/entity pair is omitted if not explicitly set. ## Attributes Reference Only the arguments listed above are exposed as attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_storage_transfer_job",
			Category:         "Google Storage Transfer Resources",
			ShortDescription: `Creates a new Transfer Job in Google Cloud Storage Transfer.`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"transfer",
				"job",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Unique description to identify the Transfer Job.`,
				},
				resource.Attribute{
					Name:        "transfer_spec",
					Description: `(Required) Transfer specification. Structure documented below.`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `(Required) Schedule specification defining when the Transfer Job should be scheduled to start, end and and what time to run. Structure documented below. - - -`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The project in which the resource belongs. If it is not provided, the provider project is used.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Status of the job. Default: ` + "`" + `ENABLED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "gcs_data_sink",
					Description: `(Required) A Google Cloud Storage data sink. Structure documented below.`,
				},
				resource.Attribute{
					Name:        "object_conditions",
					Description: `(Optional) Only objects that satisfy these object conditions are included in the set of data source and data sink objects. Object conditions based on objects' ` + "`" + `last_modification_time` + "`" + ` do not exclude objects in a data sink. Structure documented below.`,
				},
				resource.Attribute{
					Name:        "transfer_options",
					Description: `(Optional) Characteristics of how to treat files from datasource and sink during job. If the option ` + "`" + `delete_objects_unique_in_sink` + "`" + ` is true, object conditions based on objects' ` + "`" + `last_modification_time` + "`" + ` are ignored and do not exclude objects in a data source or a data sink. Structure documented below.`,
				},
				resource.Attribute{
					Name:        "gcs_data_source",
					Description: `(Optional) A Google Cloud Storage data source. Structure documented below.`,
				},
				resource.Attribute{
					Name:        "aws_s3_data_source",
					Description: `(Optional) An AWS S3 data source. Structure documented below.`,
				},
				resource.Attribute{
					Name:        "http_data_source",
					Description: `(Optional) An HTTP URL data source. Structure documented below. The ` + "`" + `schedule` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "schedule_start_date",
					Description: `(Required) The first day the recurring transfer is scheduled to run. If ` + "`" + `schedule_start_date` + "`" + ` is in the past, the transfer will run for the first time on the following day. Structure documented below.`,
				},
				resource.Attribute{
					Name:        "schedule_end_date",
					Description: `(Optional) The last day the recurring transfer will be run. If ` + "`" + `schedule_end_date` + "`" + ` is the same as ` + "`" + `schedule_start_date` + "`" + `, the transfer will be executed only once. Structure documented below.`,
				},
				resource.Attribute{
					Name:        "start_time_of_day",
					Description: `(Optional) The time in UTC at which the transfer will be scheduled to start in a day. Transfers may start later than this time. If not specified, recurring and one-time transfers that are scheduled to run today will run immediately; recurring transfers that are scheduled to run on a future date will start at approximately midnight UTC on that date. Note that when configuring a transfer with the Cloud Platform Console, the transfer's start time in a day is specified in your local timezone. Structure documented below. The ` + "`" + `object_conditions` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "max_time_elapsed_since_last_modification",
					Description: `(Optional) A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".`,
				},
				resource.Attribute{
					Name:        "min_time_elapsed_since_last_modification",
					Description: `(Optional) A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".`,
				},
				resource.Attribute{
					Name:        "include_prefixes",
					Description: `(Optional) If ` + "`" + `include_refixes` + "`" + ` is specified, objects that satisfy the object conditions must have names that start with one of the ` + "`" + `include_prefixes` + "`" + ` and that do not start with any of the ` + "`" + `exclude_prefixes` + "`" + `. If ` + "`" + `include_prefixes` + "`" + ` is not specified, all objects except those that have names starting with one of the ` + "`" + `exclude_prefixes` + "`" + ` must satisfy the object conditions. See [Requirements](https://cloud.google.com/storage-transfer/docs/reference/rest/v1/TransferSpec#ObjectConditions).`,
				},
				resource.Attribute{
					Name:        "exclude_prefixes",
					Description: `(Optional) ` + "`" + `exclude_prefixes` + "`" + ` must follow the requirements described for ` + "`" + `include_prefixes` + "`" + `. See [Requirements](https://cloud.google.com/storage-transfer/docs/reference/rest/v1/TransferSpec#ObjectConditions). The ` + "`" + `transfer_options` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "overwrite_objects_already_existing_in_sink",
					Description: `(Optional) Whether overwriting objects that already exist in the sink is allowed.`,
				},
				resource.Attribute{
					Name:        "delete_objects_unique_in_sink",
					Description: `(Optional) Whether objects that exist only in the sink should be deleted. Note that this option and ` + "`" + `delete_objects_from_source_after_transfer` + "`" + ` are mutually exclusive.`,
				},
				resource.Attribute{
					Name:        "delete_objects_from_source_after_transfer",
					Description: `(Optional) Whether objects should be deleted from the source after they are transferred to the sink. Note that this option and ` + "`" + `delete_objects_unique_in_sink` + "`" + ` are mutually exclusive. The ` + "`" + `gcs_data_sink` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `(Required) Google Cloud Storage bucket name. The ` + "`" + `gcs_data_source` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `(Required) Google Cloud Storage bucket name. The ` + "`" + `aws_s3_data_source` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `(Required) S3 Bucket name.`,
				},
				resource.Attribute{
					Name:        "aws_access_key",
					Description: `(Required) AWS credentials block. The ` + "`" + `aws_access_key` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "access_key_id",
					Description: `(Required) AWS Key ID.`,
				},
				resource.Attribute{
					Name:        "secret_access_key",
					Description: `(Required) AWS Secret Access Key. The ` + "`" + `http_data_source` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "list_url",
					Description: `(Required) The URL that points to the file that stores the object list entries. This file must allow public access. Currently, only URLs with HTTP and HTTPS schemes are supported. The ` + "`" + `schedule_start_date` + "`" + ` and ` + "`" + `schedule_end_date` + "`" + ` blocks support:`,
				},
				resource.Attribute{
					Name:        "year",
					Description: `(Required) Year of date. Must be from 1 to 9999.`,
				},
				resource.Attribute{
					Name:        "month",
					Description: `(Required) Month of year. Must be from 1 to 12.`,
				},
				resource.Attribute{
					Name:        "day",
					Description: `(Required) Day of month. Must be from 1 to 31 and valid for the year and month. The ` + "`" + `start_time_of_day` + "`" + ` blocks support:`,
				},
				resource.Attribute{
					Name:        "hours",
					Description: `(Required) Hours of day in 24 hour format. Should be from 0 to 23`,
				},
				resource.Attribute{
					Name:        "minutes",
					Description: `(Required) Minutes of hour of day. Must be from 0 to 59.`,
				},
				resource.Attribute{
					Name:        "seconds",
					Description: `(Optional) Seconds of minutes of the time. Must normally be from 0 to 59.`,
				},
				resource.Attribute{
					Name:        "nanos",
					Description: `(Required) Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Transfer Job.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `When the Transfer Job was created.`,
				},
				resource.Attribute{
					Name:        "last_modification_time",
					Description: `When the Transfer Job was last modified.`,
				},
				resource.Attribute{
					Name:        "deletion_time",
					Description: `When the Transfer Job was deleted. ## Import Storage buckets can be imported using the Transfer Job's ` + "`" + `project` + "`" + ` and ` + "`" + `name` + "`" + ` without the ` + "`" + `transferJob/` + "`" + ` prefix, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_storage_transfer_job.nightly-backup-transfer-job my-project-1asd32/8422144862922355674 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Transfer Job.`,
				},
				resource.Attribute{
					Name:        "creation_time",
					Description: `When the Transfer Job was created.`,
				},
				resource.Attribute{
					Name:        "last_modification_time",
					Description: `When the Transfer Job was last modified.`,
				},
				resource.Attribute{
					Name:        "deletion_time",
					Description: `When the Transfer Job was deleted. ## Import Storage buckets can be imported using the Transfer Job's ` + "`" + `project` + "`" + ` and ` + "`" + `name` + "`" + ` without the ` + "`" + `transferJob/` + "`" + ` prefix, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_storage_transfer_job.nightly-backup-transfer-job my-project-1asd32/8422144862922355674 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_tpu_node",
			Category:         "Google Cloud TPU",
			ShortDescription: `A Cloud TPU instance.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"tpu",
				"node",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The immutable name of the TPU.`,
				},
				resource.Attribute{
					Name:        "accelerator_type",
					Description: `(Required) The type of hardware accelerators associated with this node.`,
				},
				resource.Attribute{
					Name:        "tensorflow_version",
					Description: `(Required) The version of Tensorflow running in the Node.`,
				},
				resource.Attribute{
					Name:        "cidr_block",
					Description: `(Required) The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network that is using that CIDR block.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The GCP location for the TPU. - - -`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The user-supplied description of the TPU. Maximum of 512 characters.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) The name of a network to peer the TPU node to. It must be a preexisting Compute Engine network inside of the project on which this API has been activated. If none is provided, "default" will be used.`,
				},
				resource.Attribute{
					Name:        "scheduling_config",
					Description: `(Optional) Sets the scheduling options for this TPU instance. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Resource labels to represent user provided metadata.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `(Optional) The ID of the project in which the resource belongs. If it is not provided, the provider project is used. The ` + "`" + `scheduling_config` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "preemptible",
					Description: `(Optional) Defines whether the TPU instance is preemptible. ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "service_account",
					Description: `The service account used to run the tensor flow services within the node. To share resources, including Google Cloud Storage data, with the Tensorflow job running in the Node, this account must have permissions to that data.`,
				},
				resource.Attribute{
					Name:        "network_endpoints",
					Description: `The network endpoints where TPU workers can be accessed and sent work. It is recommended that Tensorflow clients of the node first reach out to the first (index 0) entry. Structure is documented below. The ` + "`" + `network_endpoints` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP address of this network endpoint.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port of this network endpoint. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 15 minutes. - ` + "`" + `update` + "`" + ` - Default is 15 minutes. - ` + "`" + `delete` + "`" + ` - Default is 15 minutes. ## Import Node can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_tpu_node.default projects/{{project}}/locations/{{zone}}/nodes/{{name}} $ terraform import google_tpu_node.default {{project}}/{{zone}}/{{name}} $ terraform import google_tpu_node.default {{zone}}/{{name}} $ terraform import google_tpu_node.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "service_account",
					Description: `The service account used to run the tensor flow services within the node. To share resources, including Google Cloud Storage data, with the Tensorflow job running in the Node, this account must have permissions to that data.`,
				},
				resource.Attribute{
					Name:        "network_endpoints",
					Description: `The network endpoints where TPU workers can be accessed and sent work. It is recommended that Tensorflow clients of the node first reach out to the first (index 0) entry. Structure is documented below. The ` + "`" + `network_endpoints` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP address of this network endpoint.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port of this network endpoint. ## Timeouts This resource provides the following [Timeouts](/docs/configuration/resources.html#timeouts) configuration options: - ` + "`" + `create` + "`" + ` - Default is 15 minutes. - ` + "`" + `update` + "`" + ` - Default is 15 minutes. - ` + "`" + `delete` + "`" + ` - Default is 15 minutes. ## Import Node can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import google_tpu_node.default projects/{{project}}/locations/{{zone}}/nodes/{{name}} $ terraform import google_tpu_node.default {{project}}/{{zone}}/{{name}} $ terraform import google_tpu_node.default {{zone}}/{{name}} $ terraform import google_tpu_node.default {{name}} ` + "`" + `` + "`" + `` + "`" + ` -> If you're importing a resource with beta features, make sure to include ` + "`" + `-provider=google-beta` + "`" + ` as an argument so that Terraform uses the correct provider to import your resource. ## User Project Overrides This resource supports [User Project Overrides](https://www.terraform.io/docs/providers/google/provider_reference.html#user_project_override).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "google_project_usage_export_bucket",
			Category:         "Google Cloud Platform Resources",
			ShortDescription: `Manages a project's usage export bucket.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"platform",
				"project",
				"usage",
				"export",
				"bucket",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"google_access_context_manager_access_level":      0,
		"google_access_context_manager_access_policy":     1,
		"google_access_context_manager_service_perimeter": 2,
		"google_app_engine_application":                   3,
		"google_app_engine_firewall_rule":                 4,
		"google_app_engine_standard_app_version":          5,
		"google_bigquery_dataset":                         6,
		"google_bigquery_table":                           7,
		"google_bigtable_instance":                        8,
		"google_bigtable_instance_iam_policy":             9,
		"google_bigtable_table":                           10,
		"google_binary_authorization_attestor":            11,
		"google_binary_authorization_policy":              12,
		"google_cloud_run_service":                        13,
		"google_cloud_scheduler_job":                      14,
		"google_cloudbuild_trigger":                       15,
		"google_cloudfunctions_function":                  16,
		"google_cloudiot_registry":                        17,
		"google_composer_environment":                     18,
		"google_compute_address":                          19,
		"google_compute_attached_disk":                    20,
		"google_compute_autoscaler":                       21,
		"google_compute_backend_bucket":                   22,
		"google_compute_backend_bucket_signed_url_key":    23,
		"google_compute_backend_service":                  24,
		"google_compute_backend_service_signed_url_key":   25,
		"google_compute_disk":                             26,
		"google_compute_external_vpn_gateway":             27,
		"google_compute_firewall":                         28,
		"google_compute_forwarding_rule":                  29,
		"google_compute_global_address":                   30,
		"google_compute_global_forwarding_rule":           31,
		"google_compute_ha_vpn_gateway":                   32,
		"google_compute_health_check":                     33,
		"google_compute_http_health_check":                34,
		"google_compute_https_health_check":               35,
		"google_compute_image":                            36,
		"google_compute_instance":                         37,
		"google_compute_instance_from_template":           38,
		"google_compute_instance_group":                   39,
		"google_compute_instance_group_manager":           40,
		"google_compute_instance_iam_policy":              41,
		"google_compute_instance_template":                42,
		"google_compute_interconnect_attachment":          43,
		"google_compute_managed_ssl_certificate":          44,
		"google_compute_network":                          45,
		"google_compute_network_endpoint":                 46,
		"google_compute_network_endpoint_group":           47,
		"google_compute_network_peering":                  48,
		"google_compute_node_group":                       49,
		"google_compute_node_template":                    50,
		"google_compute_project_default_network_tier":     51,
		"google_compute_project_metadata":                 52,
		"google_compute_project_metadata_item":            53,
		"google_compute_region_autoscaler":                54,
		"google_compute_region_backend_service":           55,
		"google_compute_region_disk":                      56,
		"google_compute_region_instance_group_manager":    57,
		"google_compute_resource_policy":                  58,
		"google_compute_route":                            59,
		"google_compute_router":                           60,
		"google_compute_router_interface":                 61,
		"google_compute_router_nat":                       62,
		"google_compute_router_peer":                      63,
		"google_compute_security_policy":                  64,
		"google_compute_shared_vpc_host_project":          65,
		"google_compute_shared_vpc_service_project":       66,
		"google_compute_snapshot":                         67,
		"google_compute_ssl_certificate":                  68,
		"google_compute_ssl_policy":                       69,
		"google_compute_subnetwork":                       70,
		"google_compute_subnetwork_iam_policy":            71,
		"google_compute_target_http_proxy":                72,
		"google_compute_target_https_proxy":               73,
		"google_compute_target_instance":                  74,
		"google_compute_target_pool":                      75,
		"google_compute_target_ssl_proxy":                 76,
		"google_compute_target_tcp_proxy":                 77,
		"google_compute_url_map":                          78,
		"google_compute_vpn_gateway":                      79,
		"google_compute_vpn_tunnel":                       80,
		"google_container_analysis_note":                  81,
		"google_container_cluster":                        82,
		"google_container_node_pool":                      83,
		"google_dataflow_job":                             84,
		"google_dataproc_autoscaling_policy":              85,
		"google_dataproc_cluster":                         86,
		"google_dataproc_cluster_iam_policy":              87,
		"google_dataproc_job":                             88,
		"google_dataproc_job_iam_policy":                  89,
		"google_dns_managed_zone":                         90,
		"google_dns_policy":                               91,
		"google_dns_record_set":                           92,
		"google_endpoints_service":                        93,
		"google_filestore_instance":                       94,
		"google_firestore_index":                          95,
		"google_billing_account_iam_binding":              96,
		"google_billing_account_iam_member":               97,
		"google_billing_account_iam_policy":               98,
		"google_folder":                                   99,
		"google_folder_iam_binding":                       100,
		"google_folder_iam_member":                        101,
		"google_folder_iam_policy":                        102,
		"google_folder_organization_policy":               103,
		"google_iap_tunnel_instance_iam_policy":           104,
		"google_kms_crypto_key_iam_binding":               105,
		"google_kms_crypto_key_iam_member":                106,
		"google_kms_key_ring_iam_policy":                  107,
		"google_organization_iam_binding":                 108,
		"google_organization_iam_custom_role":             109,
		"google_organization_iam_member":                  110,
		"google_organization_iam_policy":                  111,
		"google_organization_policy":                      112,
		"google_project":                                  113,
		"google_project_iam_policy":                       114,
		"google_project_iam_custom_role":                  115,
		"google_project_organization_policy":              116,
		"google_project_service":                          117,
		"google_project_services":                         118,
		"google_service_account":                          119,
		"google_service_account_iam_policy":               120,
		"google_service_account_key":                      121,
		"google_healthcare_dataset":                       122,
		"google_healthcare_dataset_iam_policy":            123,
		"google_healthcare_dicom_store":                   124,
		"google_healthcare_dicom_store_iam_policy":        125,
		"google_healthcare_fhir_store":                    126,
		"google_healthcare_fhir_store_iam_policy":         127,
		"google_healthcare_hl7_v2_store":                  128,
		"google_healthcare_hl7_v2_store_iam_policy":       129,
		"google_kms_crypto_key":                           130,
		"google_kms_key_ring":                             131,
		"google_logging_billing_account_exclusion":        132,
		"google_logging_billing_account_sink":             133,
		"google_logging_folder_exclusion":                 134,
		"google_logging_folder_sink":                      135,
		"google_logging_metric":                           136,
		"google_logging_organization_exclusion":           137,
		"google_logging_organization_sink":                138,
		"google_logging_project_exclusion":                139,
		"google_logging_project_sink":                     140,
		"google_ml_engine_model":                          141,
		"google_monitoring_alert_policy":                  142,
		"google_monitoring_group":                         143,
		"google_monitoring_notification_channel":          144,
		"google_monitoring_uptime_check_config":           145,
		"google_pubsub_subscription":                      146,
		"google_pubsub_subscription_iam_policy":           147,
		"google_pubsub_topic":                             148,
		"google_pubsub_topic_iam_policy":                  149,
		"google_redis_instance":                           150,
		"google_resource_manager_lien":                    151,
		"google_runtimeconfig_config":                     152,
		"google_runtimeconfig_variable":                   153,
		"google_scc_source":                               154,
		"google_security_scanner_scan_config":             155,
		"google_service_networking_connection":            156,
		"google_sourcerepo_repository":                    157,
		"google_sourcerepo_repository_iam_policy":         158,
		"google_spanner_database":                         159,
		"google_spanner_database_iam_policy":              160,
		"google_spanner_instance":                         161,
		"google_spanner_instance_iam_policy":              162,
		"google_sql_database":                             163,
		"google_sql_database_instance":                    164,
		"google_sql_ssl_cert":                             165,
		"google_sql_user":                                 166,
		"google_storage_bucket":                           167,
		"google_storage_bucket_acl":                       168,
		"google_storage_bucket_iam_member":                169,
		"google_storage_bucket_object":                    170,
		"google_storage_default_object_access_control":    171,
		"google_storage_default_object_acl":               172,
		"google_storage_notification":                     173,
		"google_storage_object_access_control":            174,
		"google_storage_object_acl":                       175,
		"google_storage_transfer_job":                     176,
		"google_tpu_node":                                 177,
		"google_project_usage_export_bucket":              178,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
