// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsSfnActivityInvalidNameRule checks the pattern is valid
type AwsSfnActivityInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsSfnActivityInvalidNameRule returns new rule with default attributes
func NewAwsSfnActivityInvalidNameRule() *AwsSfnActivityInvalidNameRule {
	return &AwsSfnActivityInvalidNameRule{
		resourceType:  "aws_sfn_activity",
		attributeName: "name",
		max:           80,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsSfnActivityInvalidNameRule) Name() string {
	return "aws_sfn_activity_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSfnActivityInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSfnActivityInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSfnActivityInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSfnActivityInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 80 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
