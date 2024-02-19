package main

import (
	"fmt"

	word "github.com/norrapat01/map/cal"
)

type Vertex struct{
	Lat, Lon float64
}

var m map[string]Vertex

func main()  {
	fmt.Println(m)

	m=make(map[string]Vertex)

	m["Odds"]=Vertex{
		13.000, 50.000,
	}

	fmt.Println(m)

	json := map[string]string{}
	json["message"]= "hello world"
	fmt.Println(json)

	ss:="hello world world"

	counted := map[string]int{
		"hello" : 1,
		"world": 2,
	}
	
	word.WordCount(ss)
}

