package core

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
)

type errorMessage struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func NewValidator(s interface{}) []errorMessage {
	en := en.New()
	uni = ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	validate = validator.New()
	entranslations.RegisterDefaultTranslations(validate, trans)

	return parse(trans, s)
}

func parse(trans ut.Translator, s interface{}) []errorMessage {
	errors := []errorMessage{}
	err := validate.Struct(s)

	if err != nil {
		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			errors = append(errors, errorMessage{
				Field: e.Namespace(),
				Error: e.Translate(trans),
			})
		}

		return errors
	}

	return nil
}
