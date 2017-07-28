package main

import (
	momo "buobao/util"
	"fmt"
	love "kiki/util"
)

type Circle float64
type Height float64

func main() {
	name := "buobao"
	age := 20

	names := []string{"Jack", "Tom"}

	fmt.Printf("Hello,%s,%d\n", name, age)
	fmt.Printf("%d\n", len(names))

	for _, name := range names {
		fmt.Println(name)
	}

	var c Circle
	var h Height
	fmt.Println(c)
	fmt.Println(h)
	//fmt.Println(c==h)
	fmt.Println(c == Circle(h))
	love.Echo()
	momo.Echo()
}
