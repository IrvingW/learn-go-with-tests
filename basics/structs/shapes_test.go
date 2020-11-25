package structs

import "testing"

func TestArea(t *testing.T) {
	testCases := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangle", shape: Rectangle{Height: 10, Width: 10}, want: 100.0},
		{name: "circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "triangle", shape: Triangle{2, 3}, want: 3},
	}

	for _, cc := range testCases {
		t.Run("area test for "+cc.name, func(t *testing.T) {
			got := cc.shape.Area()
			if got != cc.want {
				t.Errorf("got %.2f but expected %.2f", got, cc.want)
			}
		})
	}
}
