package validation

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/vnnyx/golang-api/exception"
	"github.com/vnnyx/golang-api/model"
)

func Validate(request model.CustomerCreateRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Username, validation.Required),
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.Password, validation.Required),
		validation.Field(&request.Gender, validation.Required),
	)
	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

func UpdateValidate(request model.CustomerUpdateRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Username, validation.Required),
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.Password, validation.Required),
		validation.Field(&request.Gender, validation.Required),
	)
	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
