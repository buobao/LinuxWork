package main

import (
	"fmt"
)

type Action interface {
	say()
}

type Zhangfei struct {
	name string
}

type Guanyu struct {
	name string
}

func (zhan *Zhangfei) say() {
	fmt.Printf("hello,%s\n", zhan.name)
}

func (guanyu *Guanyu) say() {
	fmt.Printf("hi,%s\n", guanyu.name)
}

func main() {
	var ptr Action
	zhang := Zhangfei{"zhang fei"}
	guan := Guanyu{"guan yu"}

	ptr = &zhang
	ptr.say()
	ptr = &guan
	ptr.say()
}
