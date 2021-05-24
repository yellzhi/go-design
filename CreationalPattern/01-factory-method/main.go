package main

import "fmt"

func main()  {
	var factory OperatorFactory

	factory = PlusOperatorFactory{}
	compute(factory, 1, 2)

	factory = MinOperatorFactory{}
	compute(factory, 1,2)
}

func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	fmt.Println(op.Result())
	return op.Result()
}
