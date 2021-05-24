package main

/*
装饰模式(Decorator Pattern) ：动态地给一个对象增加一些额外的职责(Responsibility)，
就增加对象功能来说，装饰模式比生成子类实现更为灵活。其别名也可以称为包装器(Wrapper)，
与适配器模式的别名相同，但它们适用于不同的场合。根据翻译的不同，装饰模式也有人称之为“油漆工模式”，
它是一种对象结构型模式。

装饰模式包含如下角色：
Component: 抽象构件
ConcreteComponent: 具体构件
Decorator: 抽象装饰类
ConcreteDecorator: 具体装饰类

 */

// 面条
type Nodeles interface {
	Description() string
	Price() float32
}

// 拉面
type Ramen struct {
	name string
	price float32
}
func (r Ramen) Description() string {
	return r.name
}
func (r Ramen) Price() float32 {
	return r.price
}
var _ Nodeles = &Ramen{}

// 鸡蛋面
type Egg struct {
	noddles Nodeles
	name string
	price float32
}
func (e Egg) Description() string {
	return e.name + "+" + e.noddles.Description()
}
func (e Egg) Price() float32 {
	return e.price + e.noddles.Price()
}
func (e *Egg) SetNoddles(noddle Nodeles){
	e.noddles= noddle
}
var _ Nodeles = &Egg{}
