package strategy

//圆周率
var pai = 3.1415926

type Shape interface {
	Area() float64
	Perimeter() float64
}
