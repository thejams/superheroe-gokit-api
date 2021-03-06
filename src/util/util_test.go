package util_test

import (
	"fmt"
	"superheroe-gokit-api/src/entity"
	"superheroe-gokit-api/src/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	batman = entity.Superheroe{
		ID:        "1",
		Name:      "Batman",
		Publisher: "DC",
	}
	superman = entity.Superheroe{
		ID:        "2",
		Name:      "Superman",
		Publisher: "DC",
	}
)

func TestVerifySuperheroe(t *testing.T) {
	t.Run("should return error when name is already taken", func(t *testing.T) {
		thor := entity.Superheroe{
			ID:        "1",
			Name:      "Thor",
			Publisher: "God of Thunder",
		}
		sh := []*entity.Superheroe{&thor}
		err := util.VerifySuperheroe(sh, thor)

		assert.NotNil(t, err)
		assert.Equal(t, "Name is already taken", err.Error())
	})

	t.Run("should not return error when a new heroe is verified", func(t *testing.T) {
		sh := []*entity.Superheroe{&batman}
		err := util.VerifySuperheroe(sh, superman)

		assert.Nil(t, err)
	})
}

func TestSuperheroeExists(t *testing.T) {
	t.Run("should return true if a heroe exists", func(t *testing.T) {
		hulk := entity.Superheroe{
			ID:        "1",
			Name:      "The Hulk",
			Publisher: "Bruce Banner",
		}
		sh := []*entity.Superheroe{&hulk}
		resp := util.SuperheroeExists(sh, "1")

		assert.True(t, resp)
	})

	t.Run("should return false if a heroe does not exists", func(t *testing.T) {
		wonderWoman := entity.Superheroe{
			ID:        "1",
			Name:      "Wonder Woman",
			Publisher: "Diana Prince",
		}
		sh := []*entity.Superheroe{&wonderWoman}
		resp := util.SuperheroeExists(sh, "3")

		assert.False(t, resp)
	})
}

func BenchmarkVerifySuperheroe(b *testing.B) {
	sh := []*entity.Superheroe{&batman}
	for i := 0; i < b.N; i++ {
		util.VerifySuperheroe(sh, superman)
	}
}

func BenchmarkSuperheroeExists(b *testing.B) {
	sh := []*entity.Superheroe{&batman}
	for i := 0; i < b.N; i++ {
		util.SuperheroeExists(sh, "3")
	}
}

func ExampleSuperheroeExists() {
	sh := []*entity.Superheroe{&batman}
	fmt.Println(util.SuperheroeExists(sh, "1"))
	fmt.Println(util.SuperheroeExists(sh, "2"))
	//Output:
	//true
	//false
}
