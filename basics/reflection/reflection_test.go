package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	name := "Irving"
	var got []string
	x := struct {
		Name string
	}{name}
	Walk(x, func(input string) {
		got = append(got, input)
	})
	want := []string{"Irving"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v but want %v", got, want)
	}
}
