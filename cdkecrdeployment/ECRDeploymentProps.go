package cdkecrdeployment

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

type ECRDeploymentProps struct {
	// The destination of the docker image.
	Dest IImageName `field:"required" json:"dest" yaml:"dest"`
	// The source of the docker image.
	Src IImageName `field:"required" json:"src" yaml:"src"`
	// Image to use to build Golang lambda for custom resource, if download fails or is not wanted.
	//
	// Might be needed for local build if all images need to come from own registry.
	//
	// Note that image should use yum as a package manager and have golang available.
	BuildImage *string `field:"optional" json:"buildImage" yaml:"buildImage"`
	// The environment variable to set.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// The amount of memory (in MiB) to allocate to the AWS Lambda function which replicates the files from the CDK bucket to the destination bucket.
	//
	// If you are deploying large files, you will need to increase this number
	// accordingly.
	MemoryLimit *float64 `field:"optional" json:"memoryLimit" yaml:"memoryLimit"`
	// Execution role associated with this function.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The VPC network to place the deployment lambda handler in.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
	// Where in the VPC to place the deployment lambda handler.
	//
	// Only used if 'vpc' is supplied.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

