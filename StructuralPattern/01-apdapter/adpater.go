package main

/*
	适配器模式(Adapter Pattern) ：将一个接口转换成客户希望的另一个接口，
适配器模式使接口不兼容的那些类可以一起工作，其别名为包装器(Wrapper)。
适配器模式既可以作为类结构型模式，也可以作为对象结构型模式。

适配器模式包含如下角色：
Target：目标抽象类
Adapter：适配器类
Adaptee：适配者类
Client：客户类
 */

// 适配器目标接口
type Target interface {
	Request() string
}

// 被适配的目标接口
type  Adaptee interface {
	SpecifiRequest() string
}

// 被适配接口的工厂函数
func NewAdeptee() Adaptee{
	return &AdapteeImpl{}
}

// 被适配的目标类
type AdapteeImpl struct {
}
func (AdapteeImpl) SpecifiRequest() string{
	return "adaptee method"
}

// 转换Adaptee为target接口的适配器
type adapter struct {
	Adaptee
}
func (a adapter) Request() string {
	return a.SpecifiRequest()
}

//接口实现
func NewAdapter(adaptee Adaptee) Target{
	return &adapter{Adaptee:adaptee}
}