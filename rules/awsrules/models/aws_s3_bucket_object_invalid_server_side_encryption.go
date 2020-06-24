// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsS3BucketObjectInvalidServerSideEncryptionRule checks the pattern is valid
type AwsS3BucketObjectInvalidServerSideEncryptionRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsS3BucketObjectInvalidServerSideEncryptionRule returns new rule with default attributes
func NewAwsS3BucketObjectInvalidServerSideEncryptionRule() *AwsS3BucketObjectInvalidServerSideEncryptionRule {
	return &AwsS3BucketObjectInvalidServerSideEncryptionRule{
		resourceType:  "aws_s3_bucket_object",
		attributeName: "server_side_encryption",
		enum: []string{
			"AES256",
			"aws:kms",
		},
	}
}

// Name returns the rule name
func (r *AwsS3BucketObjectInvalidServerSideEncryptionRule) Name() string {
	return "aws_s3_bucket_object_invalid_server_side_encryption"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsS3BucketObjectInvalidServerSideEncryptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsS3BucketObjectInvalidServerSideEncryptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsS3BucketObjectInvalidServerSideEncryptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsS3BucketObjectInvalidServerSideEncryptionRule) Check(runner *tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as server_side_encryption`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
