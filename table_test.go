package console

import (
	"os"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestToTable(t *testing.T) {
	ToTable(os.Stdout, 1)
	address := gofakeit.Address()
	ToTable(os.Stdout, address, "Name", "Value")
	m := make(map[string]string)
	for i := 0; i < 5; i++ {
		m[gofakeit.FarmAnimal()] = gofakeit.FileExtension()
	}
	ToTable(os.Stdout, m, "Key", "Value")
}

func Test_convertMap(t *testing.T) {
	m := make(map[string]interface{})
	m[gofakeit.FarmAnimal()] = gofakeit.AchAccount()
	m1, ok := convertMap(m)
	assert.True(t, ok)
	assert.EqualValues(t, m, m1)

	m2 := make(map[string]string)
	m2[gofakeit.FarmAnimal()] = gofakeit.AchAccount()
	m3, ok := convertMap(m2)
	assert.True(t, ok)
	for k := range m3 {
		assert.EqualValues(t, m2[k], m3[k])
	}

}
