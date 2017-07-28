package main

import (
	"fmt"
)

type Employee struct {
	name string
	age int
	address string
}


func main() {

	var em *Employee = new(Employee)
	em.name = "buobao"
	em.age = 10
	em.address = "china"
	
	fmt.Println(*em)

}