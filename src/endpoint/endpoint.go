//Package endpoint provides all go-kit endpoints for this micro service
package endpoint

import (
	"context"
	"superheroe-gokit-api/src/entity"
	"superheroe-gokit-api/src/service"

	"github.com/go-kit/kit/endpoint"
)

//Endpoints main endpoint struct
type Endpoints struct {
	GetAll  endpoint.Endpoint
	GetByID endpoint.Endpoint
	Add     endpoint.Endpoint
	Edit    endpoint.Endpoint
	Delete  endpoint.Endpoint
	Health  endpoint.Endpoint
}

//MakeEndpoints initialice a new set of endpoints
func MakeEndpoints(s service.Service) Endpoints {
	var (
		GetAll  endpoint.Endpoint
		GetByID endpoint.Endpoint
		Add     endpoint.Endpoint
		Edit    endpoint.Endpoint
		Delete  endpoint.Endpoint
		Health  endpoint.Endpoint
	)

	{
		Health = makeHealthEndpoint(s)
		GetAll = makeGetSuperheroesEndpoint(s)
		Delete = makeDeleteSuperheroeEndpoint(s)
		GetByID = makeGetSuperheroeByIdEndpoint(s)

		Add = makeAddSuperheroeEndpoint(s)
		Add = ValidateFields()(Add)

		Edit = makeEditSuperheroeEndpoint(s)
		Edit = ValidateFields()(Edit)
	}
	return Endpoints{
		GetAll,
		GetByID,
		Add,
		Edit,
		Delete,
		Health,
	}
}

func makeHealthEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, in interface{}) (interface{}, error) {
		return entity.NormalResponse{Ok: "service up"}, nil
	}
}

func makeGetSuperheroesEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, in interface{}) (interface{}, error) {
		return svc.GetAll(ctx)
	}
}

func makeGetSuperheroeByIdEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, in interface{}) (interface{}, error) {
		req := in.(entity.IDRequest)
		return svc.GetByID(ctx, req.Id)
	}
}

func makeAddSuperheroeEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, in interface{}) (interface{}, error) {
		req := in.(*entity.Superheroe)
		s, err := svc.Add(ctx, req)
		if err != nil {
			return nil, err
		}

		return entity.SuperheroeResponse{Superheroe: s, Msg: "superheroe added"}, nil
	}
}

func makeEditSuperheroeEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, in interface{}) (interface{}, error) {
		req := in.(*entity.Superheroe)
		return svc.Edit(ctx, req)
	}
}

func makeDeleteSuperheroeEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, in interface{}) (interface{}, error) {
		req := in.(entity.IDRequest)
		return svc.Delete(ctx, req.Id)
	}
}
