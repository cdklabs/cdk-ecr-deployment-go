// CDK construct to deploy docker image to Amazon ECR
package cdkecrdeployment

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-ecr-deployment.DockerImageName",
		reflect.TypeOf((*DockerImageName)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "creds", GoGetter: "Creds"},
			_jsii_.MemberProperty{JsiiProperty: "uri", GoGetter: "Uri"},
		},
		func() interface{} {
			j := jsiiProxy_DockerImageName{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IImageName)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"cdk-ecr-deployment.ECRDeployment",
		reflect.TypeOf((*ECRDeployment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addToPrincipalPolicy", GoMethod: "AddToPrincipalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ECRDeployment{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-ecr-deployment.ECRDeploymentProps",
		reflect.TypeOf((*ECRDeploymentProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"cdk-ecr-deployment.IImageName",
		reflect.TypeOf((*IImageName)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "creds", GoGetter: "Creds"},
			_jsii_.MemberProperty{JsiiProperty: "uri", GoGetter: "Uri"},
		},
		func() interface{} {
			return &jsiiProxy_IImageName{}
		},
	)
	_jsii_.RegisterClass(
		"cdk-ecr-deployment.S3ArchiveName",
		reflect.TypeOf((*S3ArchiveName)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "creds", GoGetter: "Creds"},
			_jsii_.MemberProperty{JsiiProperty: "uri", GoGetter: "Uri"},
		},
		func() interface{} {
			j := jsiiProxy_S3ArchiveName{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IImageName)
			return &j
		},
	)
}
