// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsAppsyncGraphqlAPIInvalidAuthenticationTypeRule checks the pattern is valid
type AwsAppsyncGraphqlAPIInvalidAuthenticationTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsAppsyncGraphqlAPIInvalidAuthenticationTypeRule returns new rule with default attributes
func NewAwsAppsyncGraphqlAPIInvalidAuthenticationTypeRule() *AwsAppsyncGraphqlAPIInvalidAuthenticationTypeRule {
	return &AwsAppsyncGraphqlAPIInvalidAuthenticationTypeRule{
		resourceType:  "aws_appsync_graphql_api",
		attributeName: "authentication_type",
		enum: []string{
			"API_KEY",
			"AWS_IAM",
			"AMAZON_COGNITO_USER_POOLS",
			"OPENID_CONNECT",
		},
	}
}

// Name returns the rule name
func (r *AwsAppsyncGraphqlAPIInvalidAuthenticationTypeRule) Name() string {
	return "aws_appsync_graphql_api_invalid_authentication_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppsyncGraphqlAPIInvalidAuthenticationTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppsyncGraphqlAPIInvalidAuthenticationTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppsyncGraphqlAPIInvalidAuthenticationTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppsyncGraphqlAPIInvalidAuthenticationTypeRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as authentication_type`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
