package template

import "fmt"

type FishSoup struct {
	Cook
}

func NewFishSoup() *FishSoup {
	f := new(FishSoup)
	f.Cook.Preparation = f.Preparation
	f.Cook.Doing = f.Doing
	f.Cook.CarriedDishes = f.CarriedDishes
	return f
}

func (f FishSoup) Preparation() {
	fmt.Println("处理鱼...")
}

func (f FishSoup) Doing() {
	fmt.Println("烹饪鱼...")
}

func (f FishSoup) CarriedDishes() {
	fmt.Println("上鱼汤...")
}
