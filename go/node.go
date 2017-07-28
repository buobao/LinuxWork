package main

import "fmt"

type Node struct {

	value int
	provious,next *Node

}

var header *Node

func init() {
	var header *Node = new(Node)
	header.value = 0
	header.next = nil
	header.provious = nil
	
	for i:=1;i<10;i++ {
		tmp := new(Node)
		tmp.value = i
		
		tmp.provious = header
		tmp.next = header.next
		header.next = tmp
		if header.next != nil {
			header.next.provious = tmp
		}
		header = tmp
	}
}

func main() {
	tmp := header
	
	for i:=0;i<10;i++ {
		fmt.Println(tmp.value)
		tmp = tmp.next
	}
	
}