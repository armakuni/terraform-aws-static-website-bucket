package test

import (
	"fmt"
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

func TestTerraformAwsS3WebsiteBucketPublicAccessConfig(t *testing.T) {
	/* ARRANGE */
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../../examples/website_bucket",

		Vars: map[string]interface{}{
			"name":   fmt.Sprintf("terratest-website-bucket-test-%s", strings.ToLower(random.UniqueId())),
			"region": "eu-west-3",
		},
	})

	/* ACTION */
	plan := terraform.InitAndPlanAndShowWithStructNoLogTempPlanFile(t, terraformOptions)
	bucketPublicAccess := GetResourceChangeAfterByAddress("module.test_website_bucket.aws_s3_bucket_public_access_block.this", plan)

	/* ASSERTIONS */
	assert.Equal(t, false, bucketPublicAccess["block_public_acls"])
	assert.Equal(t, false, bucketPublicAccess["block_public_policy"])
	assert.Equal(t, false, bucketPublicAccess["ignore_public_acls"])
	assert.Equal(t, false, bucketPublicAccess["restrict_public_buckets"])
}

func TestTerraformAwsS3WebsiteBucketOwnershipControls(t *testing.T) {
	/* ARRANGE */
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../../examples/website_bucket",

		Vars: map[string]interface{}{
			"name":   fmt.Sprintf("terratest-website-bucket-test-%s", strings.ToLower(random.UniqueId())),
			"region": "eu-west-3",
		},
	})

	/* ACTION */
	plan := terraform.InitAndPlanAndShowWithStructNoLogTempPlanFile(t, terraformOptions)
	bucketOwnershipControls := GetResourceChangeAfterByAddress("module.test_website_bucket.aws_s3_bucket_ownership_controls.this", plan)
	ownershipControlRules := bucketOwnershipControls["rule"].([]interface{})[0].(map[string]interface{})

	/* ASSERTIONS */
	assert.Equal(t, "ObjectWriter", ownershipControlRules["object_ownership"])
}

func TestTerraformAwsS3WebsiteBucketVersioningEnabled(t *testing.T) {
	/* ARRANGE */
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../../examples/website_bucket",

		Vars: map[string]interface{}{
			"name":   fmt.Sprintf("terratest-website-bucket-test-%s", strings.ToLower(random.UniqueId())),
			"region": "eu-west-3",
		},
	})

	/* ACTION */
	plan := terraform.InitAndPlanAndShowWithStructNoLogTempPlanFile(t, terraformOptions)
	bucketVersioning := GetResourceChangeAfterByAddress("module.test_website_bucket.aws_s3_bucket_versioning.this", plan)
	fmt.Println(bucketVersioning["versioning_configuration"].([]interface{})[0].(map[string]interface{})["status"])
	isVersionEnabled := bucketVersioning["versioning_configuration"].([]interface{})[0].(map[string]interface{})["status"]

	/* ASSERTIONS */
	assert.Equal(t, "Enabled", isVersionEnabled)
}
