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

// validate is a package level variable so that it can be used
var validate *validator.Validate

// ValidateRequest is a function to validate request body
func ValidateRequest(data interface{}) (string, error) {
	validate = validator.New()                                      // Initialize validator
	english := en.New()                                             // Initialize english translator
	uni := ut.New(english, english)                                 // Initialize universal translator
	trans, _ := uni.GetTranslator("en")                             // Get translator by language
	_ = enTranslations.RegisterDefaultTranslations(validate, trans) // Register default translation
	var errFields []model.ErrorInput                                // Initialize error fields
	elemType := reflect.TypeOf(data).Elem()                         // Get an element type of data
	err := validate.Struct(data)                                    // Validate data

	// If there is an error, return error message
	if err != nil {

		// Loop through all errors
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := err.Field()                        // Get field name
			field, _ := elemType.FieldByName(fieldName)     // Get field by name
			jsonTag := field.Tag.Get("json")                // Get json tag
			jsonFieldName := strings.Split(jsonTag, ",")[0] // Get json field name

			// If json field name is not empty, use json field name
			if jsonFieldName != "" {
				fieldName = jsonFieldName // Set field name to json field name
			}

			// Initialize error field struct
			var errField model.ErrorInput

			// Switch error tag
			switch err.Tag() {
			case "ipv4":
				errField.Field = fieldName
				errField.Message = "invalid " + fieldName + " address"
			case "required":
				errField.Field = fieldName
				errField.Message = fieldName + " is required"
			}

			// Append error field to error fields
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
