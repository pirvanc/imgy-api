package validator

import (
	"github.com/asaskevich/govalidator"
)

// Validator is...
type Validator struct {
}

// NewValidator is
func NewValidator() Validator {
	return Validator{}
}

// IsEmptyString is
func (v Validator) IsEmptyString(s string) bool {

	if govalidator.IsNull(s) {
		return true
	}

	return false
}
