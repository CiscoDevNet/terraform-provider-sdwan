---
layout: "sdwan"
page_title: "SD-WAN: sdwan_feature_template_ids"
sidebar_current: "docs-sdwan-data-source-feature_template_ids"
description: |-
  Data source for SD-WAN Feature Templates List
---
# sdwan_feature_template_ids Data Source #
Data source for SD-WAN Feature Templates List

## Example Usage ##

```hcl
data "sdwan_feature_template_ids" "get_templates" {
  template_type = "cisco_omp"
}

```

## Argument Reference ##
* `template_type` - (Required) Type of the Feature Template

## Common Attribute Reference ##
* `template` - List of Templates on vManage server with the given Template Type (see [below for nested schema](#nestedblock--template))

<a id="nestedblock--template"></a>

## Nested Schema for template
* `template_name` - Name of the Feature Template
* `template_id` - Template ID of the Feature Template