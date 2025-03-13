package cdkecrdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-ecr-deployment-go/cdkecrdeployment/v3/jsii"
)

type DockerImageName interface {
	IImageName
	// The credentials of the docker image.
	//
	// Format `user:password` or `AWS Secrets Manager secret arn` or `AWS Secrets Manager secret name`.
	//
	// If specifying an AWS Secrets Manager secret, the format of the secret should be either plain text (`user:password`) or
	// JSON (`{"username":"<username>","password":"<password>"}`).
	//
	// For more details on JSON format, see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/private-auth.html
	Creds() *string
	SetCreds(val *string)
	// The uri of the docker image.
	//
	// The uri spec follows https://github.com/containers/skopeo
	Uri() *string
}

// The jsii proxy struct for DockerImageName
type jsiiProxy_DockerImageName struct {
	jsiiProxy_IImageName
}

func (j *jsiiProxy_DockerImageName) Creds() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DockerImageName) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}


func NewDockerImageName(name *string, creds *string) DockerImageName {
	_init_.Initialize()

	if err := validateNewDockerImageNameParameters(name); err != nil {
		panic(err)
	}
	j := jsiiProxy_DockerImageName{}

	_jsii_.Create(
		"cdk-ecr-deployment.DockerImageName",
		[]interface{}{name, creds},
		&j,
	)

	return &j
}

func NewDockerImageName_Override(d DockerImageName, name *string, creds *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-ecr-deployment.DockerImageName",
		[]interface{}{name, creds},
		d,
	)
}

func (j *jsiiProxy_DockerImageName)SetCreds(val *string) {
	_jsii_.Set(
		j,
		"creds",
		val,
	)
}

