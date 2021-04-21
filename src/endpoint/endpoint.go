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
}

//MakeEndpoints initialice a new set of endpoints
func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		GetAll:  makeGetSuperheroesEndpoint(s),
		GetByID: makeGetSuperheroeByIdEndpoint(s),
		Add:     makeAddSuperheroeEndpoint(s),
		Edit:    makeEditSuperheroeEndpoint(s),
		Delete:  makeDeleteSuperheroeEndpoint(s),
	}
}

func makeGetSuperheroesEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, in interface{}) (interface{}, error) {
		s, _ := svc.GetAll(ctx)
		return entity.SuperheroesResponse{Superheroes: s}, nil
	}
}

func makeGetSuperheroeByIdEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, in interface{}) (interface{}, error) {
		req := in.(entity.GetIDRequest)
		s, _ := svc.GetByID(ctx, req.Id)
		return entity.SuperheroeResponse{Superheroe: s}, nil
	}
}

func makeAddSuperheroeEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, in interface{}) (interface{}, error) {
		req := in.(entity.Superheroe)
		s, _ := svc.Add(ctx, &req)
		return entity.SuperheroeResponse{Superheroe: s, Msg: "superheroe added"}, nil
	}
}

func makeEditSuperheroeEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, in interface{}) (interface{}, error) {
		req := in.(entity.Superheroe)
		s, _ := svc.Edit(ctx, &req)
		return entity.SuperheroeResponse{Superheroe: s, Msg: "superheroe updated"}, nil
	}
}

func makeDeleteSuperheroeEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, in interface{}) (interface{}, error) {
		req := in.(entity.GetIDRequest)
		s, _ := svc.Delete(ctx, req.Id)
		return entity.NormalResponse{Ok: s}, nil
	}
}
