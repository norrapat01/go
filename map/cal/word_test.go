package word

import (
	"reflect"
	"testing"
)

func TestWord_Count(t *testing.T)  {
	given :="hello world world"
	expected := map[string]int{
		"hello": 1,
		"world": 2,
	}

	actual :=WordCount(given)

	if (!reflect.DeepEqual(expected,actual)) {
		t.Fatalf("Expected %v but got %v",expected, actual)
	}
}