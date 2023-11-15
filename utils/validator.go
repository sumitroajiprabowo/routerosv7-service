package utils

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/sumitroajiprabowo/routerosv7-service/model"
	"reflect"
	"strings"
)

var validate *validator.Validate

func ValidateRequest(data interface{}) (string, error) {
	validate = validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(validate, trans)
	var errFields []model.ErrorInput
	elemType := reflect.TypeOf(data).Elem()
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := err.Field()
			field, _ := elemType.FieldByName(fieldName)
			jsonTag := field.Tag.Get("json")
			jsonFieldName := strings.Split(jsonTag, ",")[0]
			if jsonFieldName != "" {
				fieldName = jsonFieldName
			}
			var errField model.ErrorInput
			switch err.Tag() {
			case "ipv4":
				errField.Field = fieldName
				errField.Message = "invalid " + fieldName + " address"
			case "required":
				errField.Field = fieldName
				errField.Message = fieldName + " is required"
			}
			errFields = append(errFields, errField)
		}
	}

	// If there is no error, return an empty string
	if len(errFields) == 0 {
		return "", nil
	}

	// Convert error fields to JSON
	marshaledErr, _ := json.Marshal(errFields)
	return string(marshaledErr), errors.New("error")
}
