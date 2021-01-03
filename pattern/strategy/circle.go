package strategy

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * pai
}

func (c Circle) Perimeter() float64 {
	return 2 * pai * c.Radius
}
