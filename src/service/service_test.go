package service_test

import (
	"context"
	"fmt"
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
		ID:        "1",
		Name:      "Batman",
		Publisher: "DC",
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

		assert.Equal(t, "Batman", result.Superheroes[0].Name)
		assert.Equal(t, "DC", result.Superheroes[0].Publisher)
		assert.Equal(t, "1", result.Superheroes[0].ID)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("should return error when no superheroe is found", func(t *testing.T) {
		mockRepo := new(repoMock.Repository)
		svc := service.NewService(mockRepo, logger)
		mockRepo.On("GetSuperheroeById", "1").Return(nil, fmt.Errorf("no superheroe with id %v found", 1))
		res, _ := svc.GetByID(context.TODO(), "1")

		assert.NotNil(t, res.Error)
		assert.Equal(t, "no superheroe with id 1 found", res.Error.Error())
	})

	t.Run("should return a superheroe", func(t *testing.T) {
		mockRepo := new(repoMock.Repository)
		svc := service.NewService(mockRepo, logger)
		mockRepo.On("GetSuperheroeById", "1").Return(&batman, nil)
		result, _ := svc.GetByID(context.TODO(), "1")

		assert.Equal(t, "Batman", result.Superheroe.Name)
		assert.Equal(t, "DC", result.Superheroe.Publisher)
		assert.Equal(t, "1", result.Superheroe.ID)
	})
}

func TestAdd(t *testing.T) {
	t.Run("should return when name of heroe already exists", func(t *testing.T) {
		mockRepo := new(repoMock.Repository)
		svc := service.NewService(mockRepo, logger)
		nh := entity.Superheroe{
			Name:      "Batman",
			Publisher: "DC",
		}
		mockRepo.On("GetSuperheroes").Return(sh)
		_, err := svc.Add(context.TODO(), &nh)

		assert.NotNil(t, err)
		assert.Equal(t, "Name is already taken", err.Error())
	})

	t.Run("should add a new superheroe", func(t *testing.T) {
		mockRepo := new(repoMock.Repository)
		svc := service.NewService(mockRepo, logger)
		nh := entity.Superheroe{
			Name:      "Superman",
			Publisher: "DC",
		}
		mockRepo.On("GetSuperheroes").Return(sh)
		mockRepo.On("AddSuperheroe", &nh).Return(&nh)
		result, _ := svc.Add(context.TODO(), &nh)

		assert.Equal(t, "Superman", result.Name)
		assert.Equal(t, "DC", result.Publisher)
	})
}

func TestEdit(t *testing.T) {
	t.Run("should return error when heroe does not exists", func(t *testing.T) {
		nh := entity.Superheroe{
			ID:        "2",
			Name:      "Superman",
			Publisher: "DC",
		}
		mockRepo := new(repoMock.Repository)
		svc := service.NewService(mockRepo, logger)
		mockRepo.On("EditSuperheroe", &nh).Return(nil, fmt.Errorf("Superheroe with ID %v does not exist", 2))
		res, _ := svc.Edit(context.TODO(), &nh)

		assert.NotNil(t, res.Error)
		assert.Equal(t, "Superheroe with ID 2 does not exist", res.Error.Error())
	})

	t.Run("should edit a superheroe information", func(t *testing.T) {
		mockRepo := new(repoMock.Repository)
		svc := service.NewService(mockRepo, logger)
		nh := entity.Superheroe{
			ID:        "1",
			Name:      "Superman",
			Publisher: "DC",
		}
		mockRepo.On("EditSuperheroe", &nh).Return(&nh, nil)
		result, _ := svc.Edit(context.TODO(), &nh)

		assert.Equal(t, "Superman", result.Superheroe.Name)
		assert.Equal(t, "DC", result.Superheroe.Publisher)
	})
}

func TestDelete(t *testing.T) {
	t.Run("should return error when heroe does not exists", func(t *testing.T) {
		mockRepo := new(repoMock.Repository)
		svc := service.NewService(mockRepo, logger)
		mockRepo.On("DeleteSuperheroe", "2").Return("", fmt.Errorf("Superheroe with ID %v does not exist", 2))
		res, _ := svc.Delete(context.TODO(), "2")

		assert.NotNil(t, res.Error)
		assert.Equal(t, "Superheroe with ID 2 does not exist", res.Error.Error())
	})

	t.Run("should delete a superheroe", func(t *testing.T) {
		mockRepo := new(repoMock.Repository)
		svc := service.NewService(mockRepo, logger)
		mockRepo.On("DeleteSuperheroe", "1").Return("Character deleted 1", nil)
		result, _ := svc.Delete(context.TODO(), "1")

		assert.Equal(t, "Character deleted 1", result.Msg)
	})
}

func BenchmarkGetAll(b *testing.B) {
	mockRepo := new(repoMock.Repository)
	svc := service.NewService(mockRepo, logger)
	mockRepo.On("GetSuperheroeById", "1").Return(&batman, nil)

	for i := 0; i < b.N; i++ {
		svc.GetByID(context.TODO(), "1")
	}
}

func BenchmarkByID(b *testing.B) {
	mockRepo := new(repoMock.Repository)
	svc := service.NewService(mockRepo, logger)
	mockRepo.On("GetSuperheroes").Return([]*entity.Superheroe{&batman})

	for i := 0; i < b.N; i++ {
		svc.GetAll(context.TODO())
	}
}

func BenchmarkAdd(b *testing.B) {
	mockRepo := new(repoMock.Repository)
	svc := service.NewService(mockRepo, logger)
	nh := entity.Superheroe{
		Name:      "Superman",
		Publisher: "DC",
	}
	mockRepo.On("GetSuperheroes").Return(sh)
	mockRepo.On("AddSuperheroe", &nh).Return(&nh)

	for i := 0; i < b.N; i++ {
		svc.Add(context.TODO(), &nh)
	}
}

func BenchmarkEdit(b *testing.B) {
	mockRepo := new(repoMock.Repository)
	svc := service.NewService(mockRepo, logger)
	nh := entity.Superheroe{
		ID:        "1",
		Name:      "Superman",
		Publisher: "DC",
	}
	mockRepo.On("EditSuperheroe", &nh).Return(&nh, nil)
	mockRepo.On("GetSuperheroes").Return(sh)

	for i := 0; i < b.N; i++ {
		svc.Edit(context.TODO(), &nh)
	}
}

func BenchmarkDelete(b *testing.B) {
	mockRepo := new(repoMock.Repository)
	svc := service.NewService(mockRepo, logger)
	mockRepo.On("GetSuperheroes").Return(sh)
	mockRepo.On("DeleteSuperheroe", "1").Return("Character deleted 1", nil)

	for i := 0; i < b.N; i++ {
		svc.Delete(context.TODO(), "1")
	}
}
