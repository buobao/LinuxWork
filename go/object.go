package main

import (
	"fmt"
)

type Student struct {
	name  string
	age   int
	class string
}

func (s *Student) setName(name string) {
	s.name = name
}

func (s *Student) getName() string {
	return s.name
}

func (s *Student) say() {
	fmt.Println(s)
}

func main() {
	student := Student{age: 20}
	student.setName("Jack")
	fmt.Printf("%s,%d\n", student.getName(), student.age)
	student.say()
}
