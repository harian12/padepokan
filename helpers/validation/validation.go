package validation

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var validate *validator.Validate

func ValidationForDTO(data interface{}) (string, error) {
	validate = validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(data)
	if err != nil {
		var message string
		var errors = make(map[string]string)
		for i, erx := range err.(validator.ValidationErrors) {
			if i == 0 {
				message = erx.Translate(trans)
			}
			errors[erx.StructField()] = erx.Translate(trans)
		}

		return message, err
	}
	return "", nil

}
