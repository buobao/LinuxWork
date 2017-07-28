package main

import (
	"fmt"
)

type Actor struct {
	name string
	age int
}

type Movie struct {
	title string
	price float32	
	actors []Actor
}

func main() {

	var movies = []Movie{
		{title:"A",price:3200.0,actors:[]Actor{
			{name:"Jack",age:22},
			{name:"Tom",age:34}
		}},
		{title:"B",price:400.0,actors:[]Actor{
			{name:"Alice",age:17},
			{name:"James",age:26}
		}},
	}
	
	data, err := json.Marshal(movies)
	fmt.Println(data)
}