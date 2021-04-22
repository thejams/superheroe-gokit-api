package util_test

import (
	"superheroe-gokit-api/src/entity"
	"superheroe-gokit-api/src/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifySuperheroe(t *testing.T) {
	t.Run("should return error when name is already taken", func(t *testing.T) {
		thor := entity.Superheroe{
			ID:    "1",
			Name:  "Thor",
			Alias: "God of Thunder",
		}
		sh := []*entity.Superheroe{&thor}
		err := util.VerifySuperheroe(sh, thor)

		assert.NotNil(t, err)
		assert.Equal(t, "Name is already taken", err.Error())
	})

	t.Run("should return error when alias is already taken", func(t *testing.T) {
		thor := entity.Superheroe{
			ID:    "1",
			Name:  "Thor",
			Alias: "God of Thunder",
		}
		loki := entity.Superheroe{
			ID:    "2",
			Name:  "Loki",
			Alias: "God of Thunder",
		}
		sh := []*entity.Superheroe{&thor}
		err := util.VerifySuperheroe(sh, loki)

		assert.NotNil(t, err)
		assert.Equal(t, "Alias is already taken", err.Error())
	})

	t.Run("should not return error when a new heroe is verified", func(t *testing.T) {
		batman := entity.Superheroe{
			ID:    "1",
			Name:  "Batman",
			Alias: "Bruce Wayne",
		}
		superman := entity.Superheroe{
			ID:    "2",
			Name:  "Superman",
			Alias: "Clark Kent",
		}
		sh := []*entity.Superheroe{&batman}
		err := util.VerifySuperheroe(sh, superman)

		assert.Nil(t, err)
	})
}

func TestSuperheroeExists(t *testing.T) {
	t.Run("should return true if a heroe exists", func(t *testing.T) {
		hulk := entity.Superheroe{
			ID:    "1",
			Name:  "The Hulk",
			Alias: "Bruce Banner",
		}
		sh := []*entity.Superheroe{&hulk}
		resp := util.SuperheroeExists(sh, "1")

		assert.True(t, resp)
	})

	t.Run("should return false if a heroe does not exists", func(t *testing.T) {
		wonderWoman := entity.Superheroe{
			ID:    "1",
			Name:  "Wonder Woman",
			Alias: "Diana Prince",
		}
		sh := []*entity.Superheroe{&wonderWoman}
		resp := util.SuperheroeExists(sh, "3")

		assert.False(t, resp)
	})
}
