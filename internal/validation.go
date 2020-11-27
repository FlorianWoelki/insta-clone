package internal

import "github.com/go-playground/validator/v10"

// ValidationError wraps the validators FieldError
type ValidationError struct {
	validator.FieldError
}

// ValidationErrors is a collection of ValidationError
type ValidationErrors []ValidationError

// Validation contains following fields
type Validation struct {
	validate *validator.Validate
}

// NewValidation creates a new validation type
func NewValidation() *Validation {
	validate := validator.New()
	return &Validation{validate}
}

// Validate the item
func (v *Validation) Validate(i interface{}) ValidationErrors {
	errs := v.validate.Struct(i).(validator.ValidationErrors)

	if len(errs) == 0 {
		return nil
	}

	var returnErrs []ValidationError
	for _, err := range errs {
		ve := ValidationError{err.(validator.FieldError)}
		returnErrs = append(returnErrs, ve)
	}

	return returnErrs
}
