package cal

import (
	
	"testing"

)

func TestMaxDiff(t *testing.T)  {
	given := []int{4, 5, 1, 7, 9, 2}
	expect :=8

	actual := MaxDiff(given)

	if actual !=expect{
		t.Errorf("Expect %v but got %v",expect ,actual)
	}
}