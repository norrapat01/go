package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type Employee struct{
	Name string
	Address
}

type Address struct{
	No string
}

func main() {
	e:=Employee{}

	e.Name = "son"
	e.Address.No = "2231 zzen"

	fmt.Println(e)
	fmt.Println(Vertex{
		X: 1,
		Y: 2,
	})

	v := Vertex{}

	v.X = 5
	v.Y = 10

	fmt.Println(v)
}
