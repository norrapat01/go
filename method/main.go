package main

import "fmt"

type Square struct{
	X, Y int
}

type Rectangle struct{
	X, Y int
}

func(s Square)Area()int  {
	return s.X * s.Y
}

func (s Square)ToString()string  {
	return fmt.Sprintf("X: %v, Y: %v", s.X, s.Y)
}

func (s Square) MapToRec() Rectangle  {
	return Rectangle{
		X: s.X,
		Y: s.Y,
	}
}
func main()  {
	sq := Square{6,6}

	fmt.Println(sq.Area())
	fmt.Println(sq.ToString())
	fmt.Println(sq.MapToRec())
}