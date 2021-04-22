//Package service provides all the business logic for thie micro service
package service

import (
	"context"
	"fmt"
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
	s.logger.Log("get superheroes")
	return s.repo.GetSuperheroes(), nil
}

//GetAll return a single superheroe
func (s *service) GetByID(_ context.Context, id string) (*entity.Superheroe, error) {
	resp := s.repo.GetSuperheroeById(id)
	if resp == nil {
		return nil, fmt.Errorf("no superheroe with id %v found", id)
	}

	s.logger.Log("get superheroe by id", id)
	return resp, nil
}

//GetAll add a new superheroe
func (s *service) Add(_ context.Context, c *entity.Superheroe) (*entity.Superheroe, error) {
	resp := s.repo.GetSuperheroes()
	err := util.VerifySuperheroe(resp, *c)
	if err != nil {
		level.Error(s.logger).Log("add superheroe error:", err)
		return nil, err
	}

	uuid, _ := uuid.NewV4()
	c.ID = uuid.String()
	s.repo.AddSuperheroe(c)
	s.logger.Log("add superheroe", c.Name)

	return c, nil
}

//Edit a superheroe
func (s *service) Edit(_ context.Context, c *entity.Superheroe) (*entity.Superheroe, error) {
	resp := s.repo.GetSuperheroes()

	exists := util.SuperheroeExists(resp, c.ID)
	if exists == false {
		level.Error(s.logger).Log("edit superheroe error:", fmt.Sprintf("Superheroe with ID %v does not exist", c.ID))
		return nil, fmt.Errorf("Superheroe with ID %v does not exist", c.ID)
	}

	heroe := s.repo.EditSuperheroe(c)
	s.logger.Log("edit superheroe", heroe.ID, heroe.Name)

	return heroe, nil
}

//Delete delete a superheroe
func (s *service) Delete(_ context.Context, id string) (string, error) {
	resp := s.repo.GetSuperheroes()

	exists := util.SuperheroeExists(resp, id)
	if exists == false {
		level.Error(s.logger).Log("delete superheroe error:", fmt.Sprintf("Superheroe with ID %v does not exist", id))
		return "", fmt.Errorf("Superheroe with ID %v does not exist", id)
	}

	response := s.repo.DeleteSuperheroe(id)
	s.logger.Log("delete superheroe", id)

	return response, nil
}
