package Rectangles
import "testing"

func TestPerimeter(t *testing.T){
	t.Run("rectangle with width 10 and height 5 should have perimeter of 30", func(t *testing.T){
		rectangle := Rectangle{12, 6}
		got := rectangle.Area()
		want := 72.0
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
	t.Run("area of circle with radius 5 should be 78.5", func(t *testing.T){
		circle := Circle{5}
		got := circle.Area()
		want := 78.53981633974483
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}

// We'll improve our code with the interface and the Area method.
func TestRefactorArea(t *testing.T){
	areaChecked := func(t testing.TB, shape Shape, want float64){
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("Got %g want %g", got, want)
		}
	}
	t.Run("rectangle area", func(t *testing.T){
		rectangle := Rectangle{12, 6}
		areaChecked(t, rectangle, 72.0)
	})
	t.Run("circle area", func(t *testing.T){
		circle := Circle{10}
		areaChecked(t, circle, 314.1592653589793)
	})
}


// Further refactoring

func TestAreaFR(t *testing.T){
	testCases := [] struct {
		shape Shape
		want float64
	}{
		{Rectangle{12,6},72.0},
		{Circle{10},314.1592653589793},
	}
	for _, tt := range testCases {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}	
