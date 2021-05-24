package main

import "fmt"

func main(){
	var b Builder

	b = &BuilderA{}
	d1 := NewDirector(b)
	d1.Construct()
	fmt.Println(b.GetResult())

	b = &BuilderB{}
	d2 := NewDirector(b)
	d2.Construct()
	fmt.Println(b.GetResult())
}
