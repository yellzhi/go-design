package main

import "fmt"

/*
桥接模式(Bridge Pattern)：将抽象部分与它的实现部分分离，使它们都可以独立地变化。
它是一种对象结构型模式，又称为柄体(Handle and Body)模式或接口(Interface)模式。

桥接模式包含如下角色：
Abstraction：抽象类
RefinedAbstraction：扩充抽象类
Implementor：实现类接口
ConcreteImplementor：具体实现类
 */


// 形状
type Shape interface {
	Draw()
}

// 颜色
type Color interface{
	Paint()
}
// 白色
type Whate struct{
}
func (Whate) Paint()  {
	fmt.Println("白色")
}
func GetWhaste() Color{
	return & Whate{}
}
// 黑色
type Black struct{
}
func GetBalck() Color{
	return & Black{}
}
func (Black) Paint()  {
	fmt.Println("黑色")
}

// 图形
type ShapeRecl struct {
}
func (s ShapeRecl) Draw(){
	fmt.Println("圆形")
}
func GetRecl() Shape{
	return &ShapeRecl{}
}
//长方形
type ShapeRectangle struct {
}
func (s ShapeRectangle) Draw(){
	fmt.Println("长方形")
}
func GetRectangle() Shape{
	return & ShapeRectangle{}
}

type Canvar interface {
	Put()
}
type CanvasImpl struct{
	c Color
	s Shape
}
func (c CanvasImpl) Put(){
	fmt.Println("---------------")
	c.c.Paint()
	c.s.Draw()
	fmt.Println("---------------")
}

func NewCanvas(shape Shape, color Color) Canvar{
	return &CanvasImpl{
		c:color,
		s:shape,
	}
}
