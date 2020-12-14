package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	age  int
	city string
}

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
		{
			"struct with not string field",
			struct {
				Name string
				age  int
			}{"Irving", 32},
			[]string{"Irving"},
		},
		{
			"struct with nest struct",
			Person{
				"Irving",
				Profile{32, "Boston"},
			},
			[]string{"Irving", "Boston"},
		},
		{
			"pointer",
			&Person{
				"Irving",
				Profile{32, "Boston"},
			},
			[]string{"Irving", "Boston"},
		},
		{
			"slice",
			[]Person{
				{"Irving", Profile{32, "Boston"}},
				{"James", Profile{24, "NewYork"}},
			},
			[]string{"Irving", "Boston", "James", "NewYork"},
		},
		{
			"array",
			[2]Person{
				{"Irving", Profile{32, "Boston"}},
				{"James", Profile{24, "NewYork"}},
			},
			[]string{"Irving", "Boston", "James", "NewYork"},
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
