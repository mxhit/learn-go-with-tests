package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{20.0, 10.0}

	got := Perimeter(rectangle)
	want := 60.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	// Table-driven Test
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 20.0, Height: 10.0}, hasArea: 200.0},
		{name: "Circle", shape: Circle{10.0}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 4.0, Height: 4.0}, hasArea: 8.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.hasArea {
				t.Errorf("%#v got %g, want %g", tt.name, got, tt.hasArea)
			}
		})
	}
}
