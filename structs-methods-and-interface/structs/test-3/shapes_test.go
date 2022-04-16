package shapes
import "testing"


func TestShape(t *testing.T){
	areaTest := [] struct {
		shape Shape
		want float64
	}{
		{Rectangle{12,6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12,6}, 36.0},
	}
	for _, test := range areaTest {
		got := test.shape.Area()
		if got != test.want{
			t.Errorf("Got %g, want %g", got, test.want)
		}
	}
}

// Refactor
// Again, the implementation is fine but our tests could do with some improvement.

func TestAreaR(t *testing.T){
	areaShape := [] struct {
		shape Shape
		outPut float64
	}{
		{shape: Rectangle{Width: 12, Height:6}, outPut: 72.0},
		{shape: Triangle{Base: 12, Height:6}, outPut: 36.0},
		{shape: Circle{Radius: 10}, outPut: 314.1592653589793},
	}
	for _, test := range areaShape {
		got := test.shape.Area()
		if got != test.outPut{
			t.Errorf("%#v got %g want %g", test.shape, got, test.outPut)
		}
	}
}