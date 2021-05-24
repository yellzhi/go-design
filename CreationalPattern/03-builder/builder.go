package main

/*
造者模式(Builder Pattern)：将一个复杂对象的构建与它的表示分离，
使得同样的构建过程可以创建不同的表示。

建造者模式是一步一步创建一个复杂的对象，
它允许用户只通过指定复杂对象的类型和内容就可以构建它们，
用户不需要知道内部的具体构建细节。
 */

// 生成器
type Builder interface {
	Part1()
	Part2()
	Part3()
	GetResult() interface{}
}

// 指挥者
type Director struct {
	builder Builder
}
func NewDirector(builder Builder) *Director {
	return &Director{
		builder:builder,
	}
}
//构造
func (d *Director) Construct(){
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}

// 具体构造者
type BuilderA struct {
	result string
}

func (b BuilderA) GetResult() interface{} {
	return b.result
}

func (b *BuilderA) Part1() {
	b.result += "p1--"
}
func (b *BuilderA) Part2() {
	b.result += "p2--"
}
func (b *BuilderA) Part3() {
	b.result += "p3"
}
var _ Builder = (*BuilderA)(nil)

type BuilderB struct {
	result int
}

func (b BuilderB) GetResult() interface{} {
	return b.result
}

func (b *BuilderB) Part1() {
	b.result += 1
}
func (b *BuilderB) Part2() {
	b.result += 1
}
func (b *BuilderB) Part3() {
	b.result += 1
}
var _ Builder = &BuilderB{}
