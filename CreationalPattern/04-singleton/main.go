package main

import (
	"fmt"
	"sync"
)

func testSingle(){
	t1 := GetSingleton()
	t2 := GetSingleton()

	fmt.Println(t1 == t2)
}

func ProcessTestSingle(){
	const count = 100
	start := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(count)

	singles := [count]*Singleton{}

	for i:=0; i< count; i++{
		go func(index int) {
			//协程阻塞，等待channel被关闭才能继续运行
			<-start
			singles[index] = GetSingleton()
			wg.Done()
		}(i)
	}
	//关闭channel，所有协程同时开始运行，实现并行(parallel)
	close(start)
	wg.Wait()
	for i := 1; i < count; i++ {
		if singles[i] == singles[i-1] {
			fmt.Println("ok")
		}else {
			fmt.Println("error")
		}
	}
}

func main() {
	testSingle()
	ProcessTestSingle()
}