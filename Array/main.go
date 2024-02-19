package main

import "fmt"

func main()  {
	var a []string
	fmt.Println(a)

	if a == nil {
		fmt.Println("It nil!!")
	}

	prime := [6]int{
		2,3,5,7,11,13,
	}
	fmt.Println(prime)
	fmt.Println(prime[1:4] )
}