package cdkecrdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-ecr-deployment-go/cdkecrdeployment/v3/jsii"
)

type S3ArchiveName interface {
	IImageName
	// The credentials of the docker image.
	//
	// Format `user:password` or `AWS Secrets Manager secret arn` or `AWS Secrets Manager secret name`.
	Creds() *string
	SetCreds(val *string)
	// The uri of the docker image.
	//
	// The uri spec follows https://github.com/containers/skopeo
	Uri() *string
}

// The jsii proxy struct for S3ArchiveName
type jsiiProxy_S3ArchiveName struct {
	jsiiProxy_IImageName
}

func (j *jsiiProxy_S3ArchiveName) Creds() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ArchiveName) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}


func NewS3ArchiveName(p *string, ref *string, creds *string) S3ArchiveName {
	_init_.Initialize()

	if err := validateNewS3ArchiveNameParameters(p); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3ArchiveName{}

	_jsii_.Create(
		"cdk-ecr-deployment.S3ArchiveName",
		[]interface{}{p, ref, creds},
		&j,
	)

	return &j
}

func NewS3ArchiveName_Override(s S3ArchiveName, p *string, ref *string, creds *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-ecr-deployment.S3ArchiveName",
		[]interface{}{p, ref, creds},
		s,
	)
}

func (j *jsiiProxy_S3ArchiveName)SetCreds(val *string) {
	_jsii_.Set(
		j,
		"creds",
		val,
	)
}

