package integration_test

import (
  "context"
  "github.com/armakuni/go-empty-s3-bucket"
  "github.com/aws/aws-sdk-go-v2/aws"
  "github.com/aws/aws-sdk-go-v2/config"
  "github.com/aws/aws-sdk-go-v2/service/s3"
  awsTerratest "github.com/gruntwork-io/terratest/modules/aws"
  "github.com/gruntwork-io/terratest/modules/random"
  "github.com/gruntwork-io/terratest/modules/terraform"
  "github.com/stretchr/testify/assert"
  "io"
  "net/http"
  "strings"
  "testing"
)

func TestTerraformAwsS3WebsiteBucket(t *testing.T) {
  ///* ARRANGE */
  // Note this bucket name is restricted by the Github Actions OIDC role
  bucketName := "terratest-website-bucket-test-" + strings.ToLower(random.UniqueId())
  const awsRegion = "eu-west-3"

  terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
    TerraformDir: "../../examples/complete",
    Vars: map[string]interface{}{
      "name":   bucketName,
      "region": awsRegion,
    },
  })

  /* ACTION */
  terraform.InitAndPlan(t, terraformOptions)
  defer terraform.Destroy(t, terraformOptions)
  terraform.InitAndApply(t, terraformOptions)

  bucketID := terraform.Output(t, terraformOptions, "bucket_id")

  cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(awsRegion))
  assert.Nil(t, err, "unable to load AWS v2 SDK config")

  svc := s3.NewFromConfig(cfg)

  defer emptys3bucket.EmptyBucket(svc, bucketName)

  _, err = svc.PutObject(context.TODO(), &s3.PutObjectInput{
    Bucket: aws.String(bucketName),
    Key:    aws.String("index.html"),
    Body:   strings.NewReader("<h1>Hello World</h1>"),
  })
  assert.Nil(t, err, "failed to create index.html in the bucket")

  /* ASSERTIONS */
  // File is accessible
  url := "http://" + bucketName + ".s3.amazonaws.com/index.html"
  response, err := http.Get(url)
  assert.Nil(t, err, "failed to GET "+url)
  assert.Equal(t, 200, response.StatusCode)
  defer response.Body.Close()
  body, err := io.ReadAll(response.Body)
  assert.Nil(t, err, "failed to read body")
  assert.Contains(t, string(body), "Hello World")

  versioningStatus := awsTerratest.GetS3BucketVersioning(t, awsRegion, bucketID)
  assert.Equal(t, "", versioningStatus)
}
