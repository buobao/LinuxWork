package main

import "fmt"

type People struct {
	age int
}

func (p *People) setAge(age int) {
	p.age = age
}

func (p *People) getAge() int {
	return p.age
}

type Student struct {
	People
	name  string
	clazz string
}

func (s *Student) setName(name string) {
	s.name = name
}

func (s *Student) getName() string {
	return s.name
}

func main() {
	s := Student{People{10}, "Jack", "clazz"}
	fmt.Println(s)
	s.setAge(20)
	s.setName("James")
	fmt.Printf("%s,%d", s.getName(), s.getAge())
}
