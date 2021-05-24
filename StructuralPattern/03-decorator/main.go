package main

import "fmt"

func main(){

	noddle := Ramen{"拉面", 12.4}
	noddle.Description()
	noddle.Price()

	egg := Egg{noddles:noddle,name: "鸡蛋", price: 4.5}
	//egg.SetNoddles(noddle)
	fmt.Println(egg.Description())
	fmt.Println(egg.Price())

}
