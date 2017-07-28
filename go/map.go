package main

import (
	"fmt"
)

func main() {
	pair := make(map[string]int)
	
	pair["Jack"] = 30
	pair["Ocafor"] = 40
	
	for name,age := range pair {
		fmt.Println(name,age)
	
	}
 }