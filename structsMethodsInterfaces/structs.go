package structsMethodsInterfaces

import "math"

type (
	Shape interface {
		Area() float64
	}

	Rectangle struct {
		Width  float64
		Height float64
	}

	Circle struct {
		Radius float64
	}

	Triangle struct {
		Base   float64
		Height float64
	}
)

func Perimeter(rec Rectangle) float64 {
	return 2 * (rec.Width + rec.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}
