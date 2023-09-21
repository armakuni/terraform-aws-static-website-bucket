package test

import (
	"encoding/json"
	"testing"

	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func GetResourceChangeAfterByAddress(address string, plan *terraform.PlanStruct) map[string]interface{} {
	for _, value := range plan.ResourceChangesMap {
		if value.Address == address {
			return value.Change.After.(map[string]interface{})
		}
	}
	return nil
}

// Retrieve ACL Bucket Resource in Terraform Change After from Terraform Plan
// Extract out complex nested data structure for AccessControlPolicy to be used throughout for testing grants
func GetBucketGrantList(t *testing.T, plan *terraform.PlanStruct) AccessControlPolicy {
	actualBucketACL := GetResourceChangeAfterByAddress("module.test_website_bucket.aws_s3_bucket_acl.this", plan)
	jsonData, _ := json.Marshal(actualBucketACL)

	var accessControlPolicy AccessControlPolicy
	//NOTE: Will always return a list, even if one element exists in the tree,
	//unless accessing a list and it will be that list, see function for more.
	k8s.UnmarshalJSONPath(
		t,
		[]byte(jsonData),
		"{ .access_control_policy[0] }",
		&accessControlPolicy,
	)

	return accessControlPolicy
}

type AccessControlPolicy []struct {
	Grant []struct {
		Grantee []struct {
			EmailAddress string `json:"email_address"`
			ID           string `json:"id"`
			Type         string `json:"type"`
			URI          string `json:"uri"`
		} `json:"grantee"`
		Permission string `json:"permission"`
	} `json:"grant"`
	Owner []struct {
		ID string `json:"id"`
	} `json:"owner"`
}
