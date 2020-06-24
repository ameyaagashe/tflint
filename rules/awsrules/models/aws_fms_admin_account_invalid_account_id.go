// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsFmsAdminAccountInvalidAccountIDRule checks the pattern is valid
type AwsFmsAdminAccountInvalidAccountIDRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsFmsAdminAccountInvalidAccountIDRule returns new rule with default attributes
func NewAwsFmsAdminAccountInvalidAccountIDRule() *AwsFmsAdminAccountInvalidAccountIDRule {
	return &AwsFmsAdminAccountInvalidAccountIDRule{
		resourceType:  "aws_fms_admin_account",
		attributeName: "account_id",
		max:           1024,
		min:           1,
		pattern:       regexp.MustCompile(`^[0-9]+$`),
	}
}

// Name returns the rule name
func (r *AwsFmsAdminAccountInvalidAccountIDRule) Name() string {
	return "aws_fms_admin_account_invalid_account_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsFmsAdminAccountInvalidAccountIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsFmsAdminAccountInvalidAccountIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsFmsAdminAccountInvalidAccountIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsFmsAdminAccountInvalidAccountIDRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"account_id must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"account_id must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[0-9]+$`),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
