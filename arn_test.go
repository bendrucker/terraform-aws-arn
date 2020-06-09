package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
)

func TestAwsArn(t *testing.T) {
	cases := []struct {
		Name string
		ARN  arn.ARN
	}{
		{
			Name: "s3",
			ARN: arn.ARN{
				Partition: "aws",
				Service:   "s3",
				Resource:  "bucket/key",
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			opts := &terraform.Options{
				TerraformDir: "./",
				Vars: map[string]interface{}{
					"partition":   tc.ARN.Partition,
					"service":     tc.ARN.Service,
					"region":      tc.ARN.Region,
					"account_id":  tc.ARN.AccountID,
					"resource_id": tc.ARN.Resource,
				},
			}
			defer terraform.Destroy(t, opts)

			terraform.InitAndApply(t, opts)
			require.Equal(t, tc.ARN.String(), terraform.Output(t, opts, "arn"))
		})
	}
}
