package main

import "fmt"

func main() {
	adaptee := NewAdeptee()
	target := NewAdapter(adaptee)

	res := target.Request()
	fmt.Println(res)
}


