// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsCognitoUserGroupInvalidDescriptionRule checks the pattern is valid
type AwsCognitoUserGroupInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsCognitoUserGroupInvalidDescriptionRule returns new rule with default attributes
func NewAwsCognitoUserGroupInvalidDescriptionRule() *AwsCognitoUserGroupInvalidDescriptionRule {
	return &AwsCognitoUserGroupInvalidDescriptionRule{
		resourceType:  "aws_cognito_user_group",
		attributeName: "description",
		max:           2048,
	}
}

// Name returns the rule name
func (r *AwsCognitoUserGroupInvalidDescriptionRule) Name() string {
	return "aws_cognito_user_group_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCognitoUserGroupInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCognitoUserGroupInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCognitoUserGroupInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCognitoUserGroupInvalidDescriptionRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"description must be 2048 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
