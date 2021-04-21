//Package repository provides all the methods to persit data in a slice
package repository

import (
	"superheroe-gokit-api/src/entity"

	"github.com/go-kit/kit/log"
	"github.com/gofrs/uuid"
)

var resp []*entity.Superheroe

//Repository main repository interface
type Repository interface {
	GetSuperheroes() []*entity.Superheroe
	GetSuperheroeById(id string) *entity.Superheroe
	AddSuperheroe(c *entity.Superheroe) *entity.Superheroe
	EditSuperheroe(c *entity.Superheroe) *entity.Superheroe
	DeleteSuperheroe(id string)
	ClearRepository()
}

type repository struct {
	logger log.Logger
}

//NewRepository initialice a new repository with clean data
func NewRepository(logger log.Logger) Repository {
	uuid, _ := uuid.NewV4()
	a1 := entity.Superheroe{
		ID:    uuid.String(),
		Name:  "Thor",
		Alias: "Thor Odinson",
	}
	resp = append(resp, &a1)

	return &repository{
		logger: log.With(logger, "repo", "local"),
	}
}

//GetSuperheroes returns all the superheroes in the slice
func (r *repository) GetSuperheroes() []*entity.Superheroe {
	return resp
}

//GetSuperheroeById returns a single superheroe from the slice
func (r *repository) GetSuperheroeById(i string) *entity.Superheroe {
	for _, value := range resp {
		if value.ID == i {
			return value
		}
	}
	return nil
}

//AddSuperheroe add a new superheroe to the superheroes slice
func (r *repository) AddSuperheroe(c *entity.Superheroe) *entity.Superheroe {
	resp = append(resp, c)
	return c
}

//EditCharacter edit a superheroe with new information
func (r *repository) EditSuperheroe(c *entity.Superheroe) *entity.Superheroe {
	for index, value := range resp {
		if value.ID == c.ID {
			resp = append(resp[:index], resp[index+1:]...)
			resp = append(resp, c)
		}
	}
	return c
}

//DeleteSuperheroe remove a superheroe from the superheroes slice
func (r *repository) DeleteSuperheroe(id string) {
	for index, value := range resp {
		if value.ID == id {
			resp = append(resp[:index], resp[index+1:]...)
		}
	}
}

//ClearRepository remove all superheroes from the superheroes slice
func (r *repository) ClearRepository() {
	resp = nil
}
