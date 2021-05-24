package main

import "fmt"

type API interface {
	Say(str string)
}

func NewAPI(t int ) API {
	if 1==t{
		return &hiAPI{}
	}
	if 2==t{
		return &helloAPI{}
	}
	return nil
}

type hiAPI struct {}

func (h hiAPI) Say(name string)  {
	fmt.Println( fmt.Sprintf("Hi, %s",name))
}

type helloAPI struct {
}

func (h helloAPI) Say(name string)  {
	fmt.Println( fmt.Sprintf("Hello, %s",name))
}

func main()  {
	api := NewAPI(1)
	api.Say("Tom")

	api1 := NewAPI(2)
	api1.Say("Tom")

	api2 := NewAPI(3)
	if api2 != nil{
		api1.Say("Tom")
	}
}


