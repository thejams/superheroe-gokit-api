package endpoint

import (
	"context"
	"fmt"
	"superheroe-gokit-api/src/util"

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
				// return nil, err
				return nil, &util.BadRequestError{Message: fmt.Sprintf("Los siguientes campos son requeridos: %v", err.Error())}
			}
			return next(ctx, reqIn)
		}
	}
}
