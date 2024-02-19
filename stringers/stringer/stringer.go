package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type MyFloat float64

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func (f MyFloat) String() string {
	return "MyFloat"
}

func main() {
	a := Person{"A", 15}
	b := Person{"B", 25}

	var f MyFloat
	f = 3.141259

	fmt.Println(a, b, f)
}
