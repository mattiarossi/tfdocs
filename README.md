# tfdocs

This package is autogenerated from the [Terraform Website](https://github.com/hashicorp/terraform-website) for all the providers in `ext/providers`

In the `cmd/` there is the logic to generate the actual documentation of each Provider.

The generated documentation is placed inside a package `providers/` with the name of the provider, so for `aws` would be `providers/aws`. Inside of it there are 2 files, one for the Resources (`r.go`) and another for the DataSources (`d.go`).

## Documentation

Each Provider package have the same 4 exported varables and functions

The exported variables and functions are:

### `var Resources []*Resource`:

List of all the Resources of the Provider

### `func GetResource(r string) (*resource.Resource, error)`:

Returns the specific Resource for the resource type, example: `aws_iam_user`

### `var DataSources []*Resource`:

List of all the DataSources of the Provider

### `func GetDataSource(r string) (*resource.Resource, error)`:

Returns the specific DataSource for the resource type, example: `aws_vpcs`

## Icons (`assets/`)

To add icons we need a JSON that connects the resource name (`aws_lb`) to the actual path (`Compute/Elastic-Load-Balancing-ELB.svg`). This is done inside the `assets/` directory, there we have a folder for each Provider that may have icons. The path to the actual JSON has to be `assets/{{ ProviderName }}/icons.json` and for the format it's a simple KV where K == resource name and V == path.

The path has to be the one from the official Provider source after download:
* aws: https://aws.amazon.com/architecture/icons/
