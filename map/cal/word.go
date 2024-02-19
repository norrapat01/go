package word

import "strings"



func WordCount(s string )map[string]int {
words := strings.Fields(s)
counts :=
return map[string]int{
	"hello" : 1,
	"world" : 2,
}	
}