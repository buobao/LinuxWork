package main

import "fmt"

func main() {
	fc := ep(20)
	fc()
	fc()
	fc()
	fc()
	fc()
	fc()
	fc()
}

func ep(n int) func() {
	return func() {
		n++
		fmt.Println(n)
	}
}
