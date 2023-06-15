package cdkecrdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IImageName interface {
	// The credentials of the docker image.
	//
	// Format `user:password` or `AWS Secrets Manager secret arn` or `AWS Secrets Manager secret name`.
	Creds() *string
	SetCreds(c *string)
	// The uri of the docker image.
	//
	// The uri spec follows https://github.com/containers/skopeo
	Uri() *string
}

// The jsii proxy for IImageName
type jsiiProxy_IImageName struct {
	_ byte // padding
}

func (j *jsiiProxy_IImageName) Creds() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IImageName)SetCreds(val *string) {
	_jsii_.Set(
		j,
		"creds",
		val,
	)
}

func (j *jsiiProxy_IImageName) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

