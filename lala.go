package console

import (
	"fmt"

	"github.com/brianvoe/gofakeit"
)

func lala() {
	var person *gofakeit.PersonInfo
	person.FirstName = "John"
	if person.FirstName == "Danny" {
		person_age := 20
		ZERO := 0
		personAge := person_age / ZERO
		fmt.Println(personAge)
	}

}

func main() {
	lala()
}
