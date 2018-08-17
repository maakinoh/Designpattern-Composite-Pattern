package main

import (
	"./mobiles"
	"fmt"


)

type Mobi struct {
	Mob mobiles.Mobiles
}


func main()  {

	star1 := mobiles.Star{4}
	star2 := mobiles.Star{8}
	star3 := mobiles.Star{12}

	wire1 := mobiles.Wire{star1,star2,15}
	wire2 := mobiles.Wire{star3, wire1,17}

	mob := Mobi{wire2}
	fmt.Println(mob.Mob.Print())

}