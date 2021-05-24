package main

/*
享元模式(Flyweig'ht Pattern)：运用共享技术有效地支持大量细粒度对象的复用。
系统只使用少量的对'象，而这些对象都很相似，状态变化很小，可以实现对象的多次复用。
由于享元模式要求能够共享的对象必须是细粒度对象，因此它又称为轻量级模式，
它是一种对象结构型模式。

享元模式包含如下角色：
Flyweight: 抽象享元类
ConcreteFlyweight: 具体享元类
UnsharedConcreteFlyweight: 非共享具体享元类
FlyweightFactory: 享元工厂类
 */

type ChessPiceUnit struct {
	ID uint
	Name string
	color string
}
