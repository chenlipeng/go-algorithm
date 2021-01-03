package strategy

type Context struct {
	Shape Shape
}

func (c Context) GetArea() float64 {
	return c.Shape.Area()
}

func (c Context) GetPerimeter() float64 {
	return c.Shape.Perimeter()
}
