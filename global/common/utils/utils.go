package utils

import "github.com/go-playground/validator/v10"

func ObtainFirstValueOfValidationErrorsTranslations(errs validator.ValidationErrorsTranslations) string {
	for _, v := range errs {
		return v
	}
	return ""
}
