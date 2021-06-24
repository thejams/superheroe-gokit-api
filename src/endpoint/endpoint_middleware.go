package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-playground/validator"
)

//ValidateFields Validate the request object fields
func ValidateFields() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, reqIn interface{}) (interface{}, error) {
			validate := validator.New()
			err := validate.Struct(reqIn)
			if err != nil {
				return nil, err
			}
			return next(ctx, reqIn)
		}
	}
}
