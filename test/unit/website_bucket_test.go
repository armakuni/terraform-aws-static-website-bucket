package test

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"

	"github.com/stretchr/testify/assert"
)

func TestTerraformAwsS3WebsiteBucketNameVariableCorrectlyAppliedNamed(t *testing.T) {
	/* ARRANGE */
	// Give this S3 Bucket a unique ID for a name tag so we can distinguish it from any other Buckets provisioned
	// in your AWS account
	expectedBucketName := fmt.Sprintf("terratest-website-bucket-test-%s", strings.ToLower(random.UniqueId()))

	/* ACTION */
	// This will run `terraform init` and `terraform plan` and fail the test if there are any errors
	options := terraform.Options{
		TerraformDir: "../../examples/website_bucket",
		Vars: map[string]interface{}{
			"name":   expectedBucketName,
			"region": "eu-west-3",
		},
	}
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &options)

	/* ACTION */
	plan := terraform.InitAndPlanAndShowWithStructNoLogTempPlanFile(t, terraformOptions)

	/* ASSERTIONS */
	// Verify that our Bucket name matches variable
	bucket := GetResourceChangeAfterByAddress("module.test_website_bucket.aws_s3_bucket.this", plan)
	assert.Equal(t, expectedBucketName, bucket["bucket"])
}

func TestTerraformAwsS3WebsiteBucketAclConiguredAsPerLocalConfig(t *testing.T) {
	/* ARRANGE */
	// expectBucketPublicBlock := "{\n  PublicAccessBlockConfiguration: {\n    BlockPublicAcls: false,\n    BlockPublicPolicy: false,\n    IgnorePublicAcls: false,\n    RestrictPublicBuckets: false\n  }\n}"
	expectedCanonicalUserPermissions := []string{"WRITE_ACP", "READ_ACP", "READ", "WRITE"}
	expectedGroupPermissions := []string{"READ", "READ_ACP", "WRITE"}

	// Construct the terraform options with default retryable errors to handle the most common retryable errors in
	// terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../../examples/website_bucket",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"name":   fmt.Sprintf("terratest-website-bucket-test-%s", strings.ToLower(random.UniqueId())),
			"region": "eu-west-3",
		},
	})

	/* ACTION */
	plan := terraform.InitAndPlanAndShowWithStructNoLogTempPlanFile(t, terraformOptions)

	/* ASSERTIONS */
	accessControlPolicy := GetBucketGrantList(t, plan)
	grantList := accessControlPolicy[0].Grant

	// Verify each grant permissions are contained in our expected user permissions list
	for _, grant := range grantList {
		granteeType := grant.Grantee[0].Type
		switch granteeType {
		case "CanonicalUser":
			assert.True(t, slices.Contains(expectedCanonicalUserPermissions, grant.Permission))
		case "Group":
			assert.True(t, slices.Contains(expectedGroupPermissions, grant.Permission))
		default:
			fmt.Printf("Incorrect grantee type: %s expcted CanonicalUser or Group\n", granteeType)
			t.Fail()
		}
	}
}
