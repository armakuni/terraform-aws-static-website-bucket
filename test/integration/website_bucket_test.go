package integration_test

import (
  "strings"
  "testing"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/s3"
  awsTerratest "github.com/gruntwork-io/terratest/modules/aws"
  "github.com/gruntwork-io/terratest/modules/random"
  "github.com/gruntwork-io/terratest/modules/terraform"
  "github.com/stretchr/testify/assert"
)

func TestTerraformAwsS3WebsiteBucket(t *testing.T) {
  /* ARRANGE */
  bucketName := "terratest-website-bucket-test-" + strings.ToLower(random.UniqueId())
  awsRegion := "eu-west-3"

  terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
    TerraformDir: "../../examples/integration",
    Vars: map[string]interface{}{
      "name":   bucketName,
      "region": awsRegion,
    },
  })

  /* ACTION */
  terraform.InitAndPlan(t, terraformOptions)
  defer terraform.Destroy(t, terraformOptions)
  terraform.InitAndApply(t, terraformOptions)

  /* ASSERTIONS */
  bucketID := terraform.Output(t, terraformOptions, "bucket_id")

  versioningStatus := awsTerratest.GetS3BucketVersioning(t, awsRegion, bucketID)
  assert.Equal(t, "Enabled", versioningStatus)

  sess, err := session.NewSession(&aws.Config{Region: &awsRegion})
  assert.Nil(t, err, "Failed to create AWS session")
  svc := s3.New(sess)

  actualBucketACL, err := svc.GetBucketAcl(&s3.GetBucketAclInput{Bucket: aws.String(bucketName)})
  assert.Nil(t, err, "Unable to GetBucketAclInput")

  var actualCanonicalUserPermissions []string
  var actualGroupPermissions []string
  for _, b := range actualBucketACL.Grants {
    if *b.Grantee.Type == "CanonicalUser" {
      actualCanonicalUserPermissions = append(actualCanonicalUserPermissions, *b.Permission)
    } else {
      actualGroupPermissions = append(actualGroupPermissions, *b.Permission)
    }
  }
  expectedCanonicalUserPermissions := []string{"WRITE_ACP", "READ_ACP", "READ", "WRITE"}
  assert.ElementsMatchf(t, expectedCanonicalUserPermissions, actualCanonicalUserPermissions, "Canonical permissions do not match")
  expectedGroupPermissions := []string{"READ", "READ_ACP", "WRITE"}
  assert.ElementsMatch(t, expectedGroupPermissions, actualGroupPermissions, "Group permissions do not match")

  actualPublicAccessBlock, err := svc.GetPublicAccessBlock(&s3.GetPublicAccessBlockInput{Bucket: aws.String(bucketName)})
  assert.Nil(t, err, "Unable to GetPublicAccessBlock")
  expectedBucketPublicBlock := "{\n  PublicAccessBlockConfiguration: {\n    BlockPublicAcls: false,\n    BlockPublicPolicy: false,\n    IgnorePublicAcls: false,\n    RestrictPublicBuckets: false\n  }\n}"
  assert.Equal(t, expectedBucketPublicBlock, actualPublicAccessBlock.String())
}
