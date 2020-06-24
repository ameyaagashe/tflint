// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint/tflint"
)

func Test_AwsCloudwatchMetricAlarmInvalidComparisonOperatorRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cloudwatch_metric_alarm" "foo" {
	comparison_operator = "GreaterThanOrEqual"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsCloudwatchMetricAlarmInvalidComparisonOperatorRule(),
					Message: `"GreaterThanOrEqual" is an invalid value as comparison_operator`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cloudwatch_metric_alarm" "foo" {
	comparison_operator = "GreaterThanOrEqualToThreshold"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsCloudwatchMetricAlarmInvalidComparisonOperatorRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
