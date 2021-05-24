package main

import "sync"

/*
例模式(Singleton Pattern)：单例模式确保某一个类只有一个实例，
而且自行实例化并向整个系统提供这个实例，这个类称为单例类，它提供全局访问的方法。
 */

type Singleton struct {
}
var single *Singleton

func GetSingleton() *Singleton {
	once := sync.Once{}
	once.Do(func() {
		single = &Singleton{}
	})
	return single
}