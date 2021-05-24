package main
// 工厂方法模式(Factory Method Pattern)又称为工厂模式
// 工厂父类负责定义创建产品对象的公共接口，
//而工厂子类则负责生成具体的产品对象，

// 实际操作类
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// 工厂接口
type OperatorFactory interface {
	Create() Operator
}

// 操作基类， 实现Operator接口，封装公用方法
type OperatorBase struct {
	a, b int
}
func (o *OperatorBase) SetA(a int)  {
	o.a = a
}
func (o*OperatorBase) SetB(b int)  {
	o.b = b
}

//PlusOperator Operator 的实际加法实现
type PlusOperator struct {
	*OperatorBase
}
func (p PlusOperator) Result() int {
	return p.a +p.b
}

// PlusOperatorFactory 是 PluOperator 的工厂类
type PlusOperatorFactory struct {
}
func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase:&OperatorBase{},
	}
}

// 减法操作实现
type MinOperator struct {
	*OperatorBase
}

func (m MinOperator) Result() int  {
	return m.a - m.b
}

//  减法工厂
type MinOperatorFactory struct {
}

func (MinOperatorFactory) Create() Operator {
	return &MinOperator{
		OperatorBase:&OperatorBase{},
	}
}
