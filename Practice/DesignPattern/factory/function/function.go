package function

import "fmt"

type Person struct {
	name string
	age  int
}

// newPersonFactory
// @Description:
// @param age
// @return func(name string) Person
func newPersonFactory(age int) func(name string) Person {
	return func(name string) Person {
		return Person{
			name: name,
			age:  age,
		}
	}
}

func test() {
	newBaby := newPersonFactory(1)
	baby := newBaby("John")

	newTeenager := newPersonFactory(16)
	teen := newTeenager("Jill")

	fmt.Println(baby, teen)
}
