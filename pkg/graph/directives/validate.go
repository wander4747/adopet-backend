package directives

import (
	"context"
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"

	"github.com/99designs/gqlgen/graphql"
)

var (
	validate  *validator.Validate
	CEPRegexp = regexp.MustCompile(`^\d{5}-?\d{3}$`)
)

func init() {
	validate = validator.New()
}

func Validate(ctx context.Context, obj interface{}, next graphql.Resolver, constraint string) (interface{}, error) {
	val, err := next(ctx)
	if err != nil {
		return nil, err
	}

	fieldName := *graphql.GetPathContext(ctx).Field

	validate.RegisterValidation("zipCode", zipCode)
	err = validate.Var(val, constraint)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		transErr := fmt.Errorf("%s%+v", fieldName, validationErrors[0].Error())
		return val, fmt.Errorf("the field %s with value: %s. Error:%+v", fieldName, val, transErr)
	}

	return val, nil
}

func zipCode(fl validator.FieldLevel) bool {
	if !CEPRegexp.MatchString(fl.Field().String()) {
		return false
	}
	return true
}
