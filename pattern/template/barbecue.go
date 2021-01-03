package template

import "fmt"

type Barbecue struct {
	Cook
}

func NewBarbecue() *Barbecue {
	f := new(Barbecue)
	f.Cook.Preparation = f.Preparation
	f.Cook.Doing = f.Doing
	f.Cook.CarriedDishes = f.CarriedDishes
	return f
}

func (b Barbecue) Preparation() {
	fmt.Println("准备烧烤架及配菜...")
}

func (f Barbecue) Doing() {
	fmt.Println("烧烤中...")
}

func (f Barbecue) CarriedDishes() {
	fmt.Println("烤熟了，吃吧...")
}
