//go:build no_runtime_type_checking

package cdkecrdeployment

// Building without runtime type checking enabled, so all the below just return nil

func validateNewS3ArchiveNameParameters(p *string) error {
	return nil
}

