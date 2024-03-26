package console

import (
	"os"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestToTable(t *testing.T) {
	WriteTableTo(os.Stdout, 1)
	address := gofakeit.Address()
	WriteTableTo(os.Stdout, address, "Name", "Value")
	m := make(map[string]string)
	for i := 0; i < 5; i++ {
		m[gofakeit.FarmAnimal()] = gofakeit.FileExtension()
	}
	WriteTableTo(os.Stdout, m, "Key", "Value")
	WriteTableTo(os.Stdout, []string{"one", "two", "three"}, "Number")
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

func TestArray(t *testing.T) {
	var cars []gofakeit.CarInfo
	for i := 0; i < 10; i++ {
		cars = append(cars, *gofakeit.Car())
	}
	WriteTableTo(os.Stdout, cars)

	var cars1 []*gofakeit.CarInfo
	for i := 0; i < 10; i++ {
		cars1 = append(cars1, gofakeit.Car())
	}
	WriteTableTo(os.Stdout, cars1)
}

func TestStruct(t *testing.T) {
	WriteTable(gofakeit.Car())
	WriteTable(gofakeit.Person())
}
