//go:build !no_runtime_type_checking

package cdkecrdeployment

import (
	"fmt"
)

func validateNewS3ArchiveNameParameters(p *string) error {
	if p == nil {
		return fmt.Errorf("parameter p is required, but nil was provided")
	}

	return nil
}

