# Terraform Provider SDWAN
## NOTE: ## ### this is a fork of the CiscoDevNet provider by Apcela intended to be published independently. ###
### There are still many references to `CiscoDevNet` in the code and docs, many of these are historical at this point. ###


Requirements
------------
This provider is being developed and tested against:

- [Terraform](https://www.terraform.io/downloads.html) v1.0.11+

- [Go](https://golang.org/doc/install) 1.18+

## Building The Provider ##

Clone this repository and enter the provider directory.
```sh
$ git clone https://github.com/Apcela/terraform-provider-sdwan.git
```

Enter the provider directory.  Use the `make` command to build, test and validate the provider.  

```sh
$ cd terraform-provider-sdwan 
$ make
```

Using The Provider
------------------
Move the generated binary from the build step to the [plugin directory](https://www.terraform.io/docs/cli/config/config-file.html#implied-local-mirror-directories)/CiscoDevNet/sdwan/`<version>`/`<os>_<arch>`. Examples for `<os>_<arch>` are `windows_amd64`, `linux_arm`, `darwin_amd64`, etc. After placing it into your plugins directory, run `terraform init` to initialize it.

example:
```hcl
#configure provider with your SDWAN credentials.
terraform {
  required_providers {
    sdwan = {
      source = "Apcela/sdwan"
    }
  }
}

provider "sdwan" {
  # cisco-sdwan user name
  username = "admin"
  # cisco-sdwan password
  password = "password"
  # cisco-sdwan url
  url      = "https://my-cisco-sdwan.com"
  insecure = true
}

resource "sdwan_device_template" "test" {
  template_name = "example1"
  template_description = "For testing purpose"
  device_type = "vedge-cloud"
  device_role = "sdwan-edge"
  config_type = "file"
  factory_default = false
  template_configuration = "`"
}
```


Developing The Provider
-----------------------
If you want to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine. You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider with sanity checks present in scripts directory and put the provider binary in `$GOPATH/bin` directory.


## Debugging and Troubleshooting

- Set environment variable `TF_LOG` to one of the log levels `TRACE`, `DEBUG`, `INFO`, `WARN` or `ERROR`
- Set environment variable `TF_LOG_PATH` to write logs in a file. e.g. `TF_LOG_PATH=tf.log`

For more details visit - [Terraform Debugging](https://www.terraform.io/docs/internals/debugging.html)