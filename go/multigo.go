package main

import (
	"fmt"
	"time"
)

func show() {
	for {
		fmt.Println("child")
		time.Sleep(10000)
	}
}

func main() {
	go show()

	for {
		fmt.Println("parent")
		time.Sleep(10000)
	}
}
