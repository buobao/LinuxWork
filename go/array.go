package main

import (
	"fmt"
)

var a [7]int

func init() {

	for i:=0;i<len(a);i++ {
		a[i] = i
	}

}

func main() {
	
	reverse(&a)
	fmt.Println(a)
}

func reverse(a *[7]int) {
	for i,j:=0,len(a)-1;i<j;i,j=i+1,j-1 {
		a[i],a[j] = a[j],a[i]
	}
}