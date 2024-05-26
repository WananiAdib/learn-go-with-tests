package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("expected %.2f , received %.2f ", want, got)
	}
}

// func TestArea(t *testing.T) {
// 	checkArea := func( t *testing.T, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}

// 	}
// 	t.Run("chehcking rectangle's area", func(t *testing.T) {
// 		rectangle := Rectangle{10.0, 10.0}
// 		checkArea(t, rectangle, 100)
// 	})

// 	t.Run("circles", func(t *testing.T) {
// 		circle := Circle{10.0}
// 		checkArea(t, circle, 314.1592653589793)
// 	})
// }

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want float64
	}{
		{ shape: Rectangle{12.0, 6.0},want: 72.0},
		{ shape: Circle{10.0},want: 314.1592653589793},
		{ shape: Triangle{12, 6},want: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
