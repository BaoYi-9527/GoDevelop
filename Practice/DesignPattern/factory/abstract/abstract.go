package abstract

import "fmt"

type Person interface {
	Greet()
}

type person struct {
	name string
	age  int
}

func (p person) Greet() {
	fmt.Printf("Hi! My name is %s", p.name)
}

// NewPerson
// @Description: 抽象工厂模式 创建实例时返回的是接口而不是接头体
// 也就是使用结构体去实现接口 从而接口就是一个抽象工厂
// @param name
// @param age
// @return Person
func NewPerson(name string, age int) Person {
	return person{
		name: name,
		age:  age,
	}
}
