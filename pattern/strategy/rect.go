package strategy

type Rect struct {
	Width  float64
	Heigth float64
}

func (r Rect) Area() float64 {
	return r.Width * r.Heigth
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.Width + r.Heigth)
}
