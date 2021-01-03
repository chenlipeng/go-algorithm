package template

type Cook struct {
	Preparation   func()
	Doing         func()
	CarriedDishes func()
}

func (c Cook) Cooking() {
	c.Preparation()
	c.Doing()
	c.CarriedDishes()
}
