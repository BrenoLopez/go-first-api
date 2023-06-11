package validation

import (
	"encoding/json"
	"errors"

	httpError "github.com/BrenoLopez/go-first-api/src/config/handle-errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	pt_br_translations "github.com/go-playground/validator/v10/translations/pt_BR"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		pt_BR := pt_BR.New()
		unt := ut.New(pt_BR, pt_BR)
		transl, _ = unt.GetTranslator("pt_BR")
		pt_br_translations.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validationError error) *httpError.HttpError {
	var jsonErr *json.UnmarshalTypeError

	var jsonValidationError validator.ValidationErrors

	if errors.As(validationError, &jsonErr) {
		return httpError.NewBadRequestError("Invalid field type")
	} else if errors.As(validationError, &jsonValidationError) {
		errorsCauses := []httpError.Causes{}

		for _, err := range validationError.(validator.ValidationErrors) {
			cause := httpError.Causes{
				Message: err.Translate(transl),
				Field:   err.Field(),
			}
			errorsCauses = append(errorsCauses, cause)
		}
		return httpError.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return httpError.NewBadRequestError("Error trying to convert fields")
	}
}
