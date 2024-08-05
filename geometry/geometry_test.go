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
	}{
		{shape: Rectangle{Width: 10, Height: 20}, want: 200.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, areaTest := range areaTests {
		got := areaTest.shape.Area()
		if got != areaTest.want {
			t.Errorf("got %g want %g", got, areaTest.want)
		}

	}

}
