---
subcategory: "Guides"
page_title: "Configuration Groups"
description: |-
    Configuration Groups
---

# Configuration Groups

Configuration groups require a number of resources to be defined. We start by creating the required feature profiles.

```terraform
resource "sdwan_system_feature_profile" "system_01" {
  name        = "system_01"
  description = "My system feature profile"
}

resource "sdwan_transport_feature_profile" "transport_01" {
  name        = "transport_01"
  description = "My transport feature profile"
}
```

Next we would need to add the respective features to previously created feature profiles.

```terraform
resource "sdwan_system_basic_feature" "system_01_basic" {
  name               = "system_01_basic"
  feature_profile_id = sdwan_system_feature_profile.system_01.id
}

resource "sdwan_system_aaa_feature" "system_01_aaa" {
  name               = "system_01_aaa"
  feature_profile_id = sdwan_system_feature_profile.system_01.id
  server_auth_order  = ["local"]
  users = [{
    name     = "admin"
    password = "admin"
  }]
}

resource "sdwan_system_bfd_feature" "system_01_bfd" {
  name               = "system_01_bfd"
  feature_profile_id = sdwan_system_feature_profile.system_01.id
}

resource "sdwan_system_global_feature" "system_01_global" {
  name               = "system_01_global"
  feature_profile_id = sdwan_system_feature_profile.system_01.id
}

resource "sdwan_system_logging_feature" "system_01_logging" {
  name               = "system_01_logging"
  feature_profile_id = sdwan_system_feature_profile.system_01.id
}

resource "sdwan_system_omp_feature" "system_01_omp" {
  name               = "system_01_omp"
  feature_profile_id = sdwan_system_feature_profile.system_01.id
}

resource "sdwan_transport_wan_vpn_feature" "transport_01_wan_vpn" {
  name               = "transport_01_wan_vpn"
  feature_profile_id = sdwan_transport_feature_profile.transport_01.id
  vpn                = 0
}

resource "sdwan_transport_wan_vpn_interface_ethernet_feature" "transport_01_wan_vpn_interface_ethernet" {
  name                         = "transport_01_wan_vpn_interface_ethernet"
  feature_profile_id           = sdwan_transport_feature_profile.transport_01.id
  transport_wan_vpn_feature_id = sdwan_transport_wan_vpn_feature.transport_01_wan_vpn.id
  interface_name               = "GigabitEthernet1"
  shutdown                     = false
  ipv4_dhcp_distance           = 1
  tunnel_interface             = true
  tunnel_interface_encapsulations = [
    {
      encapsulation = "ipsec"
    }
  ]
}
```

Finally, we need to create the corresponding configuration group, referencing the previously created feature profiles. `devices` is a list of devices associated with the configuration group. Every device then requires values to be provided for all the relevant variables being used in the respective features. Every device has a flag `deploy` to enable/disable the deployment of configuration changes to the respective devices.

```terraform
resource "sdwan_configuration_group" "config_group_01" {
  name        = "config_group_01"
  description = "My config group"
  solution     = "sdwan"
  feature_profile_ids = [sdwan_system_feature_profile.system_01.id, sdwan_transport_feature_profile.transport_01.id]
  devices = [{
    id     = "C8K-40C0CCFD-9EA8-2B2E-E73B-32C5924EC79B"
    deploy = true
    variables = [
      {
        name = "host_name"
        value = "edge1"
      },
      {
        name = "pseudo_commit_timer"
        value = 0
      },
      {
        name = "site_id"
        value = 1
      },
      {
        name = "system_ip"
        value = "10.1.1.1"
      },
      {
        name = "ipv6_strict_control"
        value = "false"
      }
    ]
  }]
  feature_versions = [
    sdwan_system_basic_feature.system_01_basic.version,
    sdwan_system_aaa_feature.system_01_aaa.version,
    sdwan_system_bfd_feature.system_01_bfd.version,
    sdwan_system_global_feature.system_01_global.version,
    sdwan_system_logging_feature.system_01_logging.version,
    sdwan_system_omp_feature.system_01_omp.version,
    sdwan_transport_wan_vpn_feature.transport_01_wan_vpn.version,
    sdwan_transport_wan_vpn_interface_ethernet_feature.transport_01_wan_vpn_interface_ethernet.version,
  ]
}
```

Please note, at the end of the configuration group configuration, we need to provide a list of references to the `version` attribute of all used features. This is important in order to trigger a re-deployment of the configuration group in case any of the feature configurations change and is also required to ensure that the configuration group is always deployed *after* potential changes to any of the features have been made.

We can now simulate a change to a feature that is already deployed to a device, by for example modifying the password of the AAA feature resource.

```terraform
resource "sdwan_system_aaa_feature" "system_01_aaa" {
  name               = "system_01_aaa"
  feature_profile_id = sdwan_system_feature_profile.system_01.id
  server_auth_order  = ["local"]
  users = [{
    name     = "admin"
    password = "admin123"
  }]
}
```

After running `terraform apply` the plan should show two changes, one for the feature resource that was just modified and one for the configuration group resource which is required in order to trigger a re-deployment of the configuration.

```
  # sdwan_configuration_group.config_group_01 will be updated in-place
  ~ resource "sdwan_configuration_group" "config_group_01" {
      ~ feature_versions = [
            "0",
          + (known after apply),
            "0",
            # (3 unchanged elements hidden)
            "0",
          - "0",
        ]
        id               = "1466cfc0-8104-4300-936b-3152aad1fe70"
        name             = "config_group_01"
        # (4 unchanged attributes hidden)
    }

  # sdwan_system_aaa_feature.system_01_aaa will be updated in-place
  ~ resource "sdwan_system_aaa_feature" "system_01_aaa" {
        id                 = "5181c6ed-085c-4270-b482-f3cb284a862f"
        name               = "system_01_aaa"
      ~ users              = [
          ~ {
                name     = "admin"
              ~ password = "admin" -> "admin123"
            },
        ]
      ~ version            = 0 -> (known after apply)
        # (2 unchanged attributes hidden)
    }

Plan: 0 to add, 2 to change, 0 to destroy.
```
