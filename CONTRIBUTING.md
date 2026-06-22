# How to Contribute

Thanks for your interest in contributing to this Terraform provider! Here are a few general guidelines on contributing and
reporting bugs that we ask you to review. Following these guidelines helps to communicate that you respect the time of
the contributors managing and developing this open source project. In return, they should reciprocate that respect in
addressing your issue, assessing changes, and helping you finalize your pull requests. In that spirit of mutual respect,
we endeavor to review incoming issues and pull requests within 10 days, and will close any lingering issues or pull
requests after 60 days of inactivity.

Please note that all of your interactions in the project are subject to our [Code of Conduct](/CODE_OF_CONDUCT.md). This
includes creation of issues or pull requests, commenting on issues or pull requests, and extends to all interactions in
any real-time space e.g., Slack, Discord, etc.

## Table Of Contents

- [Reporting Issues](#reporting-issues)
- [Development](#development)
  - [Building the Provider](#building-the-provider)
  - [Code Generation](#code-generation)
  - [Acceptance Tests](#acceptance-tests)
- [Sending Pull Requests](#sending-pull-requests)
- [Other Ways to Contribute](#other-ways-to-contribute)

## Reporting Issues

Before reporting a new issue, please ensure that the issue was not already reported or fixed by searching through our
[issues list](https://github.com/CiscoDevNet/terraform-provider-sdwan/issues).

When creating a new issue, please be sure to include a **title and clear description**, as much relevant information as
possible, and, if possible, a test case.

**If you discover a security bug, please do not report it through GitHub. Instead, please see security procedures in
[SECURITY.md](/SECURITY.md).**

## Development

### Building the Provider

1. Clone the repository
2. Enter the repository directory
3. Build the provider using the Go `install` command:

```shell
go install
```

### Code Generation

This provider heavily relies on code generation to create the necessary resources and data sources. The generator takes care of creating the necessary code, documentation and acceptance tests for a particular resource or data source. The generator is written in Go and can be found in the `gen` directory. There is a two step process to eventually generate the code for a new resource or data source. First, a "definition" is being created, which is a YAML file with all the necessary information to render the code artifacts. The second step is to run the generator with the "definition" file(s) as input. The generator will then render the code artifacts for the resources or data sources.

Definition files are being maintained in three directories under `gen/definitions`:
- `gen/definitions/feature_templates/` - Feature template definitions
- `gen/definitions/profile_parcels/` - Profile parcel definitions
- `gen/definitions/generic/` - Generic resource definitions

To generate the code for a specific resource or data source, run the following command:

```shell
make gen NAME="NAME"
```

Where `NAME` is the name of the definition (the `name:` field from the YAML file), e.g., `make gen NAME="Transport WAN VPN"`. Whenever the definition is updated, it is necessary to run the generator again.

In some cases it might also be required to update some of the generated Go code. This can be done by modifying the generated files in the `internal/provider` directory. Every code section has comment markers as shown below:

```go
// Section below is generated&owned by "gen/generator.go". //template:begin create

func Create() {
}

// End of section. //template:end create
```

As long as those markers remain in the code, the code will continue to be updated by the generator. If the markers are removed, the code will not be updated anymore and the code can be modified manually.

To regenerate and/or update the complete codebase for all resources and data sources, run the following command:

```shell
make gen
```

### Acceptance Tests

In order to run the full suite of Acceptance tests, run `make test`. Make sure the respective environment variables are set (e.g., `SDWAN_USERNAME`, `SDWAN_PASSWORD`, `SDWAN_URL`). You can copy `.env.sample` to `.env` and fill in your values.

> **Note**: Acceptance tests create real resources.

```shell
# Run all acceptance tests
make test

# Run tests for a specific definition (by name from gen/definitions/*.yaml)
make test NAME="Transport WAN VPN"
make test NAME="cEdge AAA"

# Run tests with debug logging
make test NAME="Cisco BFD" DEBUG=1
```

## Sending Pull Requests

Before sending a new pull request, take a look at existing pull requests and issues to see if the proposed change or fix
has been discussed in the past, or if the change was already implemented but not yet released.

We expect new pull requests to include tests for any affected behavior, and, as we follow semantic versioning, we may
reserve breaking changes until the next major version release.

## Other Ways to Contribute

We welcome anyone that wants to contribute to this Terraform provider to triage and reply to open issues to help troubleshoot
and fix existing bugs. Here is what you can do:

- Help ensure that existing issues follows the recommendations from the _[Reporting Issues](#reporting-issues)_ section,
  providing feedback to the issue's author on what might be missing.
- Review existing pull requests, and testing patches against real infrastructure.
- Write a test, or add a missing test case to an existing test.

Thanks again for your interest on contributing to this Terraform provider!

:heart:
