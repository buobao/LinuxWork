package main

import (
	"fmt"
)

var pc [256]byte


func init() {

	for i := range pc {
	
		pc[i] = pc[i/2] + byte(i&1)
	
	}

}

func main() {


	for _,i := range pc {
		fmt.Printf("%04b\n",i)
	}

}