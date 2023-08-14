package validator

import (
	"fmt"
	"strings"

	"github.com/Padliwinata/iam-mfa/bin/pkg/errors"

	"gopkg.in/go-playground/validator.v9"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {

	if cv.Validator.Struct(i) != nil {
		errs := cv.Validator.Struct(i).(validator.ValidationErrors)
		errorMsg := fmt.Sprintf("\"%s\" is %s", strings.ToLower(errs[0].Field()), errs[0].Tag())
		return errors.Conflict(errorMsg)
	}

	return nil

}

func New() *validator.Validate {
	return validator.New()
}
