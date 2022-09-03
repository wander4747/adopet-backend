package directives

import (
	"context"
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

	validate.RegisterValidation("zipCode", zipCode)
	err = validate.Var(val, constraint)
	if err != nil {
		return nil, err
	}

	return val, nil
}

func zipCode(fl validator.FieldLevel) bool {
	if !CEPRegexp.MatchString(fl.Field().String()) {
		return false
	}
	return true
}
