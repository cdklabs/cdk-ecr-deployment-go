package cdkecrdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdklabs/cdk-ecr-deployment-go/cdkecrdeployment/v2/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdklabs/cdk-ecr-deployment-go/cdkecrdeployment/v2/internal"
)

type ECRDeployment interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	AddToPrincipalPolicy(statement awsiam.PolicyStatement) *awsiam.AddToPrincipalPolicyResult
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for ECRDeployment
type jsiiProxy_ECRDeployment struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_ECRDeployment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewECRDeployment(scope constructs.Construct, id *string, props *ECRDeploymentProps) ECRDeployment {
	_init_.Initialize()

	if err := validateNewECRDeploymentParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ECRDeployment{}

	_jsii_.Create(
		"cdk-ecr-deployment.ECRDeployment",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewECRDeployment_Override(e ECRDeployment, scope constructs.Construct, id *string, props *ECRDeploymentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-ecr-deployment.ECRDeployment",
		[]interface{}{scope, id, props},
		e,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func ECRDeployment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateECRDeployment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-ecr-deployment.ECRDeployment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ECRDeployment) AddToPrincipalPolicy(statement awsiam.PolicyStatement) *awsiam.AddToPrincipalPolicyResult {
	if err := e.validateAddToPrincipalPolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *awsiam.AddToPrincipalPolicyResult

	_jsii_.Invoke(
		e,
		"addToPrincipalPolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ECRDeployment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

