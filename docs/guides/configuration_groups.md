---
subcategory: "Guides"
page_title: "Getting Started - Configuration and Policy Groups"
description: |-
    Getting Started - Configuration and Policy Groups
---

# Getting Started - Configuration and Policy Groups

UX 2.0 replaces classic device and feature templates with an intent-based model, recommended for new deployments on SD-WAN Manager 20.15 or 20.18:

- **Feature profiles** are reusable configuration building blocks for a specific function (system, transport, service, or policy).
- **Configuration groups** combine feature profiles into a device's configuration.
- **Policy groups** combine policy feature profiles (application priority, security, and others) into a device's policy.

A device is associated with one configuration group and one policy group. If you are working with an existing classic deployment, see [Getting Started - Classic Templates](https://registry.terraform.io/providers/CiscoDevNet/sdwan/latest/docs/guides/getting_started) instead.

## Configuration Groups

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

Next, we create the corresponding configuration group, referencing the previously created feature profiles. The `feature_versions` attribute is a list of references to the `version` attribute of all used features. Whenever any of them change, the configuration group's computed `version` is incremented, which in turn can be used to trigger a redeployment of the associated devices (see below).

```terraform
resource "sdwan_configuration_group" "config_group_01" {
  name        = "config_group_01"
  description = "My config group"
  solution     = "sdwan"
  feature_profile_ids = [sdwan_system_feature_profile.system_01.id, sdwan_transport_feature_profile.transport_01.id]
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

Finally, device association, variables and deployment are managed by a separate `sdwan_configuration_group_devices` resource. It references the configuration group via `configuration_group_id` and provides a list of `devices`. Every device requires values to be provided for all the relevant variables used in the respective features, and has a `deploy` flag to enable/disable the deployment of configuration changes to that device. Setting `configuration_group_version` to `sdwan_configuration_group.<name>.version` ensures that a change to the configuration group (for example a changed feature) is deployed in the same apply: from SD-WAN Manager 20.15 to the devices the Manager reports out of date, on older versions to all devices of the resource.

Because device associations live in their own resource, the devices of a single configuration group can be spread across multiple Terraform state files (for example a few thousand devices per state) without one state impacting the devices managed by another.

```terraform
resource "sdwan_configuration_group_devices" "config_group_01_devices" {
  configuration_group_id      = sdwan_configuration_group.config_group_01.id
  solution                    = "sdwan"
  configuration_group_version = sdwan_configuration_group.config_group_01.version
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
}
```

See [Tracking Changes](#tracking-changes) below for why the `feature_versions` list on the configuration group is required.

## Policy Groups

Policy groups follow the same pattern, but reference a policy feature profile (for example application priority, topology, or security) instead of system/transport/service. We start with a feature profile and a policy feature attached to it.

```terraform
resource "sdwan_application_priority_feature_profile" "app_priority_01" {
  name        = "app_priority_01"
  description = "My application priority feature profile"
}

resource "sdwan_application_priority_qos_policy" "app_priority_01_qos" {
  name                        = "app_priority_01_qos"
  description                 = "QoS policy for application priority"
  feature_profile_id          = sdwan_application_priority_feature_profile.app_priority_01.id
  target_interfaces_variable  = "{{qos_interfaces}}"
}
```

The policy group itself references the feature profile and lists the `version` of every policy it uses via `policy_versions` - the policy equivalent of a configuration group's `feature_versions`. Like the configuration group, it has a computed `version` that is incremented whenever the group or any of its policy versions change.

```terraform
resource "sdwan_policy_group" "policy_group_01" {
  name                = "policy_group_01"
  description         = "My policy group"
  solution            = "sdwan"
  feature_profile_ids = [sdwan_application_priority_feature_profile.app_priority_01.id]
  policy_versions = [
    sdwan_application_priority_qos_policy.app_priority_01_qos.version,
  ]
}
```

Device association, variables and deployment are managed by a separate `sdwan_policy_group_devices` resource, which references the policy group via `policy_group_id`. Setting `policy_group_version` to `sdwan_policy_group.<name>.version` deploys a policy group change in the same apply, in the same way.

```terraform
resource "sdwan_policy_group_devices" "policy_group_01_devices" {
  policy_group_id      = sdwan_policy_group.policy_group_01.id
  solution             = "sdwan"
  policy_group_version = sdwan_policy_group.policy_group_01.version
  devices = [{
    id     = "C8K-40C0CCFD-9EA8-2B2E-E73B-32C5924EC79B"
    deploy = true
    variables = [
      {
        name       = "qos_interfaces"
        list_value = ["GigabitEthernet1", "GigabitEthernet2"]
      }
    ]
  }]
}
```

A device can be associated with both a configuration group and a policy group at the same time - each has its own feature profiles, features/policies, its own `feature_versions`/`policy_versions` list, and its own `*_devices` resource.

## Tracking Changes

Both `sdwan_configuration_group.feature_versions` and `sdwan_policy_group.policy_versions` require a list of references to the `version` attribute of every feature or policy they use. This is required so that:

- A change to any referenced feature/policy increments the group's `version`, which - referenced from `configuration_group_version`/`policy_group_version` - triggers a re-deployment of the devices in the corresponding `*_devices` resource.
- The devices are always deployed *after* the changes to the features/policies and the group have been applied.

We can simulate this by modifying a feature that is already deployed to a device, for example the password of the AAA feature resource from the Configuration Groups example above.

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

After running `terraform apply` the plan should show three changes: one for the feature resource that was just modified, one for the configuration group resource whose `version` is incremented as a result of the changed feature version, and one for the `sdwan_configuration_group_devices` resource which is required in order to trigger a re-deployment of the configuration to the devices.

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
      ~ version          = 0 -> (known after apply)
        # (3 unchanged attributes hidden)
    }

  # sdwan_configuration_group_devices.config_group_01_devices will be updated in-place
  ~ resource "sdwan_configuration_group_devices" "config_group_01_devices" {
        id                          = "1466cfc0-8104-4300-936b-3152aad1fe70"
      ~ configuration_group_version = 0 -> (known after apply)
        # (2 unchanged attributes hidden)
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

Plan: 0 to add, 3 to change, 0 to destroy.
```

The same applies to a policy group: modifying `sdwan_application_priority_qos_policy.app_priority_01_qos` would show a matching plan diff on `sdwan_policy_group.policy_group_01.policy_versions` and `version`, and on `sdwan_policy_group_devices.policy_group_01_devices.policy_group_version`.
