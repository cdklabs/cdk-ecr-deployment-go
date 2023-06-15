//go:build no_runtime_type_checking

package cdkecrdeployment

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ECRDeployment) validateAddToPrincipalPolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func validateECRDeployment_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewECRDeploymentParameters(scope constructs.Construct, id *string, props *ECRDeploymentProps) error {
	return nil
}

