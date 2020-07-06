module github.com/terraform-linters/tflint

go 1.14

require (
	github.com/aws/aws-sdk-go v1.33.1
	github.com/fatih/color v1.9.0
	github.com/golang/mock v1.4.3
	github.com/google/go-cmp v0.5.0
	github.com/hashicorp/aws-sdk-go-base v0.5.0
	github.com/hashicorp/go-plugin v1.3.0
	github.com/hashicorp/go-version v1.2.1
	github.com/hashicorp/hcl/v2 v2.6.0
	github.com/hashicorp/logutils v1.0.0
	github.com/hashicorp/terraform v0.12.28
	github.com/hashicorp/terraform-plugin-sdk v1.14.0
	github.com/jessevdk/go-flags v1.4.0
	github.com/mattn/go-colorable v0.1.7
	github.com/mitchellh/go-homedir v1.1.0
	github.com/serenize/snaker v0.0.0-20171204205717-a683aaf2d516
	github.com/sourcegraph/go-lsp v0.0.0-20181119182933-0c7d621186c1
	github.com/sourcegraph/jsonrpc2 v0.0.0-20190106185902-35a74f039c6a
	github.com/spf13/afero v1.3.1
	github.com/terraform-linters/tflint-plugin-sdk v0.2.0
	github.com/terraform-providers/terraform-provider-aws v2.68.0+incompatible
	github.com/zclconf/go-cty v1.5.1
	golang.org/x/lint v0.0.0-20191125180803-fdd1cda4f05f
)

replace github.com/terraform-providers/terraform-provider-aws v2.68.0+incompatible => github.com/terraform-providers/terraform-provider-aws v1.60.1-0.20200625234409-8688f3adfb43
