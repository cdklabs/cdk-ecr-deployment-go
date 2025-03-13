# cdk-ecr-deployment

[![Release](https://github.com/cdklabs/cdk-ecr-deployment/actions/workflows/release.yml/badge.svg)](https://github.com/cdklabs/cdk-ecr-deployment/actions/workflows/release.yml)
[![npm version](https://img.shields.io/npm/v/cdk-ecr-deployment)](https://www.npmjs.com/package/cdk-ecr-deployment)
[![PyPI](https://img.shields.io/pypi/v/cdk-ecr-deployment)](https://pypi.org/project/cdk-ecr-deployment)
[![npm](https://img.shields.io/npm/dw/cdk-ecr-deployment?label=npm%20downloads)](https://www.npmjs.com/package/cdk-ecr-deployment)
[![PyPI - Downloads](https://img.shields.io/pypi/dw/cdk-ecr-deployment?label=pypi%20downloads)](https://pypi.org/project/cdk-ecr-deployment)

CDK construct to synchronize single docker image between docker registries.

**Only use v3 of this package**

⚠️ Version 2.* is no longer supported, as the Go.1.x runtime is no longer supported in AWS Lambda.\
⚠️ Version 1.* is no longer supported, as CDK v1 has reached the end-of-life
stage.

## Features

* Copy image from ECR/external registry to (another) ECR/external registry
* Copy an archive tarball image from s3 to ECR/external registry

## Environment variables

Enable flags: `true`, `1`. e.g. `export CI=1`

* `CI` indicate if it's CI environment. This flag will enable building lambda from scratch.
* `NO_PREBUILT_LAMBDA` disable using prebuilt lambda.
* `FORCE_PREBUILT_LAMBDA` force using prebuilt lambda.

⚠️ If you want to force using prebuilt lambda in CI environment to reduce build time. Try `export FORCE_PREBUILT_LAMBDA=1`.

## Examples

```go
import "github.com/aws/aws-cdk-go/awscdk"


image := awscdk.NewDockerImageAsset(this, jsii.String("CDKDockerImage"), &DockerImageAssetProps{
	Directory: path.join(__dirname, jsii.String("docker")),
})

// Copy from cdk docker image asset to another ECR.
// Copy from cdk docker image asset to another ECR.
ecrdeploy.NewECRDeployment(this, jsii.String("DeployDockerImage1"), &ECRDeploymentProps{
	Src: ecrdeploy.NewDockerImageName(image.ImageUri),
	Dest: ecrdeploy.NewDockerImageName(fmt.Sprintf("%v.dkr.ecr.us-west-2.amazonaws.com/my-nginx:latest", cdk.Aws_ACCOUNT_ID())),
})

// Copy from docker registry to ECR.
// Copy from docker registry to ECR.
ecrdeploy.NewECRDeployment(this, jsii.String("DeployDockerImage2"), &ECRDeploymentProps{
	Src: ecrdeploy.NewDockerImageName(jsii.String("nginx:latest")),
	Dest: ecrdeploy.NewDockerImageName(fmt.Sprintf("%v.dkr.ecr.us-west-2.amazonaws.com/my-nginx2:latest", cdk.Aws_ACCOUNT_ID())),
})

// Copy from private docker registry to ECR.
// The format of secret in aws secrets manager must be either:
// - plain text in format <username>:<password>
// - json in format {"username":"<username>","password":"<password>"}
// Copy from private docker registry to ECR.
// The format of secret in aws secrets manager must be either:
// - plain text in format <username>:<password>
// - json in format {"username":"<username>","password":"<password>"}
ecrdeploy.NewECRDeployment(this, jsii.String("DeployDockerImage3"), &ECRDeploymentProps{
	Src: ecrdeploy.NewDockerImageName(jsii.String("javacs3/nginx:latest"), jsii.String("username:password")),
	// src: new ecrdeploy.DockerImageName('javacs3/nginx:latest', 'aws-secrets-manager-secret-name'),
	// src: new ecrdeploy.DockerImageName('javacs3/nginx:latest', 'arn:aws:secretsmanager:us-west-2:000000000000:secret:id'),
	Dest: ecrdeploy.NewDockerImageName(fmt.Sprintf("%v.dkr.ecr.us-west-2.amazonaws.com/my-nginx3:latest", cdk.Aws_ACCOUNT_ID())),
}).AddToPrincipalPolicy(awscdk.Aws_iam.NewPolicyStatement(&PolicyStatementProps{
	Effect: awscdk.*Aws_iam.Effect_ALLOW,
	Actions: []*string{
		jsii.String("secretsmanager:GetSecretValue"),
	},
	Resources: []*string{
		jsii.String("*"),
	},
}))
```

## Sample: [test/example.ecr-deployment.ts](./test/example.ecr-deployment.ts)

```shell
# Run the following command to try the sample.
NO_PREBUILT_LAMBDA=1 npx cdk deploy -a "npx ts-node -P tsconfig.dev.json --prefer-ts-exts test/example.ecr-deployment.ts"
```

## [API](./API.md)

## Tech Details & Contribution

The core of this project relies on [containers/image](https://github.com/containers/image) which is used by [Skopeo](https://github.com/containers/skopeo).
Please take a look at those projects before contribution.

To support a new docker image source(like docker tarball in s3), you need to implement [image transport interface](https://github.com/containers/image/blob/master/types/types.go). You could take a look at [docker-archive](https://github.com/containers/image/blob/ccb87a8d0f45cf28846e307eb0ec2b9d38a458c2/docker/archive/transport.go) transport for a good start.

To test the `lambda` folder, `make test`.
