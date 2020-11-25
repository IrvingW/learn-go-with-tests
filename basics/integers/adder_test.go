package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	want := 4
	if sum != want {
		t.Errorf("got %d but expected %d", sum, want)
	}
}

// This is an example of Add func(ExampleXXX)
func ExampleAdd() {
	sum := Add(1, 6)
	fmt.Println(sum)
	// Output: 7
	// the below doc will be tested too
}
