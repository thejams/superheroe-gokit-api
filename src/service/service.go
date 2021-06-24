//Package service provides all the business logic for thie micro service
package service

import (
	"context"
	"superheroe-gokit-api/src/entity"
	"superheroe-gokit-api/src/repository"
	"superheroe-gokit-api/src/util"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

//Service main interface for the service with the business logic
type Service interface {
	GetAll(context.Context) ([]*entity.Superheroe, error)
	GetByID(_ context.Context, id string) (*entity.Superheroe, error)
	Add(_ context.Context, c *entity.Superheroe) (*entity.Superheroe, error)
	Edit(_ context.Context, c *entity.Superheroe) (*entity.Superheroe, error)
	Delete(_ context.Context, id string) (string, error)
}

type service struct {
	repo   repository.Repository
	logger log.Logger
}

//NewService initialice a new service
func NewService(rep repository.Repository, logger log.Logger) Service {
	return &service{
		repo:   rep,
		logger: logger,
	}
}

//GetAll return all superheroes
func (s *service) GetAll(context.Context) ([]*entity.Superheroe, error) {
	logger := log.With(s.logger, "method", "GetAll")
	logger.Log("superheroes returned")
	return s.repo.GetSuperheroes(), nil
}

//GetAll return a single superheroe
func (s *service) GetByID(_ context.Context, id string) (*entity.Superheroe, error) {
	logger := log.With(s.logger, "method", "GetByID")
	resp, err := s.repo.GetSuperheroeById(id)
	if err != nil {
		level.Error(s.logger).Log("getById error:", err)
		return nil, err
	}

	logger.Log("superheroe returned", id)
	return resp, nil
}

//GetAll add a new superheroe
func (s *service) Add(_ context.Context, c *entity.Superheroe) (*entity.Superheroe, error) {
	logger := log.With(s.logger, "method", "Add")
	resp := s.repo.GetSuperheroes()
	err := util.VerifySuperheroe(resp, *c)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}

	uuid, _ := uuid.NewV4()
	c.ID = uuid.String()
	s.repo.AddSuperheroe(c)
	logger.Log("superheroe added", c.Name)
	return c, nil
}

//Edit a superheroe
func (s *service) Edit(_ context.Context, c *entity.Superheroe) (*entity.Superheroe, error) {
	logger := log.With(s.logger, "method", "Edit")
	heroe, err := s.repo.EditSuperheroe(c)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}

	logger.Log("superheroe edited", heroe.ID, heroe.Name)
	return heroe, nil
}

//Delete delete a superheroe
func (s *service) Delete(_ context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "Delete")
	response, err := s.repo.DeleteSuperheroe(id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("superheroe deleted", id)
	return response, nil
}
