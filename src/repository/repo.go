//Package repository provides all the methods to persit data in a slice
package repository

import (
	"fmt"
	"superheroe-gokit-api/src/entity"

	"github.com/go-kit/kit/log"
	"github.com/gofrs/uuid"
)

var resp []*entity.Superheroe

//Repository main repository interface
type Repository interface {
	GetSuperheroes() []*entity.Superheroe
	GetSuperheroeById(id string) (*entity.Superheroe, error)
	AddSuperheroe(c *entity.Superheroe) *entity.Superheroe
	EditSuperheroe(c *entity.Superheroe) (*entity.Superheroe, error)
	DeleteSuperheroe(id string) (string, error)
	ClearRepository()
}

type repository struct {
	logger log.Logger
}

//NewRepository initialice a new repository with clean data
func NewRepository(logger log.Logger) Repository {
	uuid, _ := uuid.NewV4()
	thor := entity.Superheroe{
		ID:        uuid.String(),
		Name:      "Thor",
		Publisher: "Marvel",
	}
	joker := entity.Superheroe{
		ID:        uuid.String(),
		Name:      "The Joker",
		Publisher: "DC",
	}
	resp = append(resp, &thor, &joker)

	return &repository{
		logger: log.With(logger, "repo", "local"),
	}
}

//GetSuperheroes returns all the superheroes in the slice
func (r *repository) GetSuperheroes() []*entity.Superheroe {
	return resp
}

//GetSuperheroeById returns a single superheroe from the slice
func (r *repository) GetSuperheroeById(i string) (*entity.Superheroe, error) {
	for _, value := range resp {
		if value.ID == i {
			return value, nil
		}
	}
	return nil, fmt.Errorf("no superheroe with id %v found", i)
}

//AddSuperheroe add a new superheroe to the superheroes slice
func (r *repository) AddSuperheroe(c *entity.Superheroe) *entity.Superheroe {
	resp = append(resp, c)
	return c
}

//EditCharacter edit a superheroe with new information
func (r *repository) EditSuperheroe(c *entity.Superheroe) (*entity.Superheroe, error) {
	for index, value := range resp {
		if value.ID == c.ID {
			resp = append(resp[:index], resp[index+1:]...)
			resp = append(resp, c)
			return c, nil
		}
	}
	return nil, fmt.Errorf("Superheroe with ID %v does not exist", c.ID)
}

//DeleteSuperheroe remove a superheroe from the superheroes slice
func (r *repository) DeleteSuperheroe(id string) (string, error) {
	for index, value := range resp {
		if value.ID == id {
			resp = append(resp[:index], resp[index+1:]...)
			return "Character deleted " + id, nil
		}
	}
	return "", fmt.Errorf("Superheroe with ID %v does not exist", id)
}

//ClearRepository remove all superheroes from the superheroes slice
func (r *repository) ClearRepository() {
	resp = nil
}
