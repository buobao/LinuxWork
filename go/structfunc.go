package main

import (
	"fmt"
	"math"
)

// -----------------
//      Struct
// -----------------

type Person struct {
	name  string
	age   int
	email string
}

func (p *Person) getName() string {
	return p.name
}

// -------------------
//      Interface
// -------------------

type shape interface {
	area() float64
}

type rect struct {
	width  float64
	height float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// -----------------
//      Test
// -----------------

/**
 * 结构Person[{cdai 30 cdai@gmail.com}]，姓名[cdai]
 * 结构Person指针[&{cdai 30 cdai@gmail.com}]，姓名[cdai]
 * 用指针修改结构Person为[{carter 40 cdai@gmail.com}]
 *
 * Shape[0]周长为[13.920000]
 * Shape[1]周长为[58.088048]
 */
func main() {
	testStruct()
	testInterface()
}

func testStruct() {
	p1 := Person{"cdai", 30, "cdai@gmail.com"}
	p1 = Person{name: "cdai", age: 30, email: "cdai@gmail.com"}
	fmt.Printf("结构Person[%v]，姓名[%s]\n", p1, p1.getName())

	ptr1 := &p1
	fmt.Printf("结构Person指针[%v]，姓名[%s]\n", ptr1, ptr1.getName())

	ptr1.age = 40
	ptr1.name = "carter"
	fmt.Printf("用指针修改结构Person为[%v]\n", p1)
}

func testInterface() {
	r := rect{width: 2.9, height: 4.8}
	c := circle{radius: 4.3}

	s := []shape{&r, &c}
	for i, sh := range s {
		fmt.Printf("Shape[%d]周长为[%f]\n", i, sh.area())
	}
}
