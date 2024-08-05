package geometry

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 20.0}
	got := Perimeter(rectangle)
	want := 60.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
		name  string
	}{
		{shape: Rectangle{Width: 10, Height: 20}, want: 200.0, name: "Rectangle area"},
		{shape: Circle{Radius: 10}, want: 314.1592653589793, name: "Circle area"},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0, name: "Triangle area"},
	}

	for _, areaTest := range areaTests {
		t.Run(areaTest.name, func(t *testing.T) {
			got := areaTest.shape.Area()
			if got != areaTest.want {
				t.Errorf("%#v got %g want %g", areaTest.shape, got, areaTest.want)
			}
		})

	}

}
