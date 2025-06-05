package main

import "fmt"

type person struct{
	name string
	age int
}

func newPerson(name string) *person {
	p:=person{name: name}
	p.age = 42
	return &p
}

func personDemo(){
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Fred", age: 40})
	fmt.Println(person{name:"Alice"})
	fmt.Println(&person{name:"Anna",age :40})

	fmt.Println(newPerson("John"))

	s := person{name: "Sean",age :50}
	fmt.Println(s.name)
	fmt.Println(s.age)

	sp:= &s
	fmt.Println(sp.age)

	sp.age = 599
	fmt.Println(sp.age)
	fmt.Println((*(&sp)).age)
}