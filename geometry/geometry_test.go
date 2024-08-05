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
	t.Run("rectangles area", func(t *testing.T) {
		rectangle := Rectangle{10.0, 20.0}
		got := rectangle.Area()
		want := 200.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	t.Run("circles area",func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want:=314.15

		if got!= want{
			t.Errorf("got %g want %g",got,want)
		}
	})
}
