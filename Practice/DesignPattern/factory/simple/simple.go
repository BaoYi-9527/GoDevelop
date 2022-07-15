package simple

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("Hi！My name is %s", p.Name)
}

// NewPerson
// @Description:简单工厂模式
// @param name
// @param age
// @return *Person
func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}
