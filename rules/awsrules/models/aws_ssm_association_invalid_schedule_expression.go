// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsSsmAssociationInvalidScheduleExpressionRule checks the pattern is valid
type AwsSsmAssociationInvalidScheduleExpressionRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsSsmAssociationInvalidScheduleExpressionRule returns new rule with default attributes
func NewAwsSsmAssociationInvalidScheduleExpressionRule() *AwsSsmAssociationInvalidScheduleExpressionRule {
	return &AwsSsmAssociationInvalidScheduleExpressionRule{
		resourceType:  "aws_ssm_association",
		attributeName: "schedule_expression",
		max:           256,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsSsmAssociationInvalidScheduleExpressionRule) Name() string {
	return "aws_ssm_association_invalid_schedule_expression"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmAssociationInvalidScheduleExpressionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmAssociationInvalidScheduleExpressionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmAssociationInvalidScheduleExpressionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmAssociationInvalidScheduleExpressionRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"schedule_expression must be 256 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"schedule_expression must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
