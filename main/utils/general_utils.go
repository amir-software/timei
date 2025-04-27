// Nested Package

package utils

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func init() { // init function package utils
	fmt.Print("Hello world")
	Validator = validator.New(validator.WithRequiredStructEnabled())
}

func ValidateStruct(structure any) error {
	// a function for validating the structs
	err := Validator.Struct(structure)

	return err
}

func GetFieldOfStructe(structure any, fieldName string) reflect.StructField {
	// retrun an object as struct and consist of the all info about a struct field
	dataType := reflect.TypeOf(structure)
	field, _ := dataType.FieldByName(fieldName)

	return field
} // temp := utils.GetFieldOfStructe(serviceA, "Title")
