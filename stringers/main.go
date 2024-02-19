package main

import "fmt"

func String(s interface{})string  {
	ss, ok :=s.(string)
	if !ok{
		return "unknown"
	}
	return ss
}

func main() {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(ok)

	f, ok := i.(float64)
	if !ok{
		fmt.Println("do something ...")
	}

	fmt.Println(f, ok)

	ff := i.(float64)
	fmt.Println(ff)

	fmt.Println(String(1))
	fmt.Println(String("aaa"))
	fmt.Println(String(map[string]string{}))

}