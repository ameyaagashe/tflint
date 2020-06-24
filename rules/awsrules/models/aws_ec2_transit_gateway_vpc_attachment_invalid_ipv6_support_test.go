// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint/tflint"
)

func Test_AwsEc2TransitGatewayVpcAttachmentInvalidIpv6SupportRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_ec2_transit_gateway_vpc_attachment" "foo" {
	ipv6_support = "on"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsEc2TransitGatewayVpcAttachmentInvalidIpv6SupportRule(),
					Message: `"on" is an invalid value as ipv6_support`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_ec2_transit_gateway_vpc_attachment" "foo" {
	ipv6_support = "enable"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsEc2TransitGatewayVpcAttachmentInvalidIpv6SupportRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
