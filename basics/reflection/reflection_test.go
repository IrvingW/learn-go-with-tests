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
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
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
				Age  int
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

	for _, c := range cases {
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

	t.Run("map", func(t *testing.T) {
		var got []string
		x := make(map[string]string, 2)
		x["a"] = "b"
		x["b"] = "c"
		Walk(x, func(input string) {
			got = append(got, input)
		})
		if len(got) != 2 {
			t.Errorf("expect size is 2 but got size is %d", len(got))
		}
		assertContain(t, got, "b")
		assertContain(t, got, "c")
	})
}

func assertContain(t *testing.T, input []string, key string) {
	t.Helper()
	contains := false
	for _, item := range input {
		if item == key {
			contains = true
			break
		}
	}
	if !contains {
		t.Errorf("could not find item %v", key)
	}
}
