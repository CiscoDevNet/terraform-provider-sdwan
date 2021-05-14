# Terraform Provider SDWAN

Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) v0.13 and v0.14.

- [Go](https://golang.org/doc/install) Latest Version

## Building The Provider ##
Clone this repository to: `$GOPATH/src/github.com/CiscoDevNet/terraform-provider-sdwan`.
Clone [sdwan-go-client](https://github.com/CiscoDevNet/sdwan-go-client) to `$GOPATH/src/github.com/CiscoDevNet/sdwan-go-client`.

```sh
$ mkdir -p $GOPATH/src/github.com/CiscoDevNet; cd $GOPATH/src/github.com/CiscoDevNet
$ git clone https://github.com/CiscoDevNet/terraform-provider-sdwan.git
$ git clone https://github.com/CiscoDevNet/sdwan-go-client.git
```

Enter the provider directory, After that run make build to build the provider binary. This will build the provider with sanity checks present in scripts directory and put the provider binary in `$GOPATH/bin` directory.

```sh
$ cd $GOPATH/src/github.com/CiscoDevNet/terraform-provider-sdwan
$ make build

```


Using The Provider
------------------
If you are building the provider, follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin) After placing it into your plugins directory, run `terraform init` to initialize it.

ex.
```hcl
#configure provider with your SDWAN credentials.
terraform {
  required_providers {
    sdwan = {
      source = "CiscoDevNet/sdwan"
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