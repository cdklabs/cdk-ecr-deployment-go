//go:build no_runtime_type_checking

package cdkecrdeployment

// Building without runtime type checking enabled, so all the below just return nil

func validateNewDockerImageNameParameters(name *string) error {
	return nil
}

