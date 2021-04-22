package service_test

import (
	"context"
	"os"
	"superheroe-gokit-api/src/entity"
	repoMock "superheroe-gokit-api/src/repository/mocks"
	"superheroe-gokit-api/src/service"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
)

var (
	logger log.Logger
	batman entity.Superheroe
	sh     []*entity.Superheroe
)

func init() {
	batman = entity.Superheroe{
		ID:    "1",
		Name:  "Batman",
		Alias: "Bruce Wayne",
	}
	sh = []*entity.Superheroe{&batman}

	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "account",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
}

func TestGetAll(t *testing.T) {
	t.Run("should return an array with 1 superheroe", func(t *testing.T) {
		mockRepo := new(repoMock.Repository)
		svc := service.NewService(mockRepo, logger)
		mockRepo.On("GetSuperheroes").Return([]*entity.Superheroe{&batman})
		result, _ := svc.GetAll(context.TODO())

		assert.Equal(t, "Batman", result[0].Name)
		assert.Equal(t, "Bruce Wayne", result[0].Alias)
		assert.Equal(t, "1", result[0].ID)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("should return error when no superheroe is found", func(t *testing.T) {
		mockRepo := new(repoMock.Repository)
		svc := service.NewService(mockRepo, logger)
		mockRepo.On("GetSuperheroeById", "1").Return(nil)
		_, err := svc.GetByID(context.TODO(), "1")

		assert.NotNil(t, err)
		assert.Equal(t, "no superheroe with id 1 found", err.Error())
	})

	t.Run("should return a superheroe", func(t *testing.T) {
		mockRepo := new(repoMock.Repository)
		svc := service.NewService(mockRepo, logger)
		mockRepo.On("GetSuperheroeById", "1").Return(&batman)
		result, _ := svc.GetByID(context.TODO(), "1")

		assert.Equal(t, "Batman", result.Name)
		assert.Equal(t, "Bruce Wayne", result.Alias)
		assert.Equal(t, "1", result.ID)
	})
}

func TestAdd(t *testing.T) {
	t.Run("should return when alias or name of heroe already exists", func(t *testing.T) {
		mockRepo := new(repoMock.Repository)
		svc := service.NewService(mockRepo, logger)
		nh := entity.Superheroe{
			Name:  "Iron Man",
			Alias: "Bruce Wayne",
		}
		mockRepo.On("GetSuperheroes").Return(sh)
		_, err := svc.Add(context.TODO(), &nh)

		assert.NotNil(t, err)

		nh = entity.Superheroe{
			Name:  "Batman",
			Alias: "Tony Stark",
		}
		_, err = svc.Add(context.TODO(), &nh)

		assert.NotNil(t, err)
	})

	t.Run("should add a new superheroe", func(t *testing.T) {
		mockRepo := new(repoMock.Repository)
		svc := service.NewService(mockRepo, logger)
		nh := entity.Superheroe{
			Name:  "Superman",
			Alias: "Clark Kent",
		}
		mockRepo.On("GetSuperheroes").Return(sh)
		mockRepo.On("AddSuperheroe", &nh).Return(&nh)
		result, _ := svc.Add(context.TODO(), &nh)

		assert.Equal(t, "Superman", result.Name)
		assert.Equal(t, "Clark Kent", result.Alias)
	})
}

func TestEdit(t *testing.T) {
	t.Run("should return error when heroe does not exists", func(t *testing.T) {
		mockRepo := new(repoMock.Repository)
		svc := service.NewService(mockRepo, logger)
		mockRepo.On("GetSuperheroes").Return(sh)
		nh := entity.Superheroe{
			ID:    "2",
			Name:  "Superman",
			Alias: "Clark Kent",
		}
		_, err := svc.Edit(context.TODO(), &nh)

		assert.NotNil(t, err)
		assert.Equal(t, "Superheroe with ID 2 does not exist", err.Error())
	})

	t.Run("should edit a superheroe information", func(t *testing.T) {
		mockRepo := new(repoMock.Repository)
		svc := service.NewService(mockRepo, logger)
		nh := entity.Superheroe{
			ID:    "1",
			Name:  "Superman",
			Alias: "Clark Kent",
		}
		mockRepo.On("EditSuperheroe", &nh).Return(&nh)
		mockRepo.On("GetSuperheroes").Return(sh)
		result, _ := svc.Edit(context.TODO(), &nh)

		assert.Equal(t, "Superman", result.Name)
		assert.Equal(t, "Clark Kent", result.Alias)
	})
}

func TestDelete(t *testing.T) {
	t.Run("should return error when heroe does not exists", func(t *testing.T) {
		mockRepo := new(repoMock.Repository)
		svc := service.NewService(mockRepo, logger)
		mockRepo.On("GetSuperheroes").Return(sh)
		mockRepo.On("GetSuperheroes").Return(sh)
		_, err := svc.Delete(context.TODO(), "2")

		assert.NotNil(t, err)
		assert.Equal(t, "Superheroe with ID 2 does not exist", err.Error())
	})

	t.Run("should delete a superheroe", func(t *testing.T) {
		mockRepo := new(repoMock.Repository)
		svc := service.NewService(mockRepo, logger)
		mockRepo.On("GetSuperheroes").Return(sh)
		mockRepo.On("DeleteSuperheroe", "1").Return("Character deleted 1")
		result, _ := svc.Delete(context.TODO(), "1")

		assert.Equal(t, "Character deleted 1", result)
	})
}
