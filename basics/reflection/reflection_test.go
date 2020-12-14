package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cs := []struct {
		Name     string
		Input    interface{}
		Expected []string
	}{
		{
			"struct with only one field",
			struct {
				Name string
			}{"Irving"},
			[]string{"Irving"},
		},
		{
			"struct with two field",
			struct {
				Name  string
				Name2 string
			}{"Irving", "James"},
			[]string{"Irving", "James"},
		},
	}

	for _, c := range cs {
		t.Run(c.Name, func(t *testing.T) {
			var got []string
			Walk(c.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, c.Expected) {
				t.Errorf("got %v but want %v", got, c.Expected)
			}
		})
	}
}
