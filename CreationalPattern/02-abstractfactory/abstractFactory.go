package main

import "fmt"

/*
	象工厂模式用于生成产品族的工厂，所生成的对象是有关联的。
	如果抽象工厂退化成生成的对象无关联则成为工厂函数模式。
	比如本例子中使用RDB和XML存储订单信息，抽象工厂分别能生成相关的主订单信息和订单详情信息。
如果业务逻辑中需要替换使用的时候只需要改动工厂函数相关的类就能替换使用不同的存储方式了。

	提供一个创建一系列相关或相互依赖对象的接口，而无须指定它们具体的类。抽象工厂模式又称为Kit模式，
属于对象创建型模式。
*/

// 订单生成 事例

// 订单主记录
type OrderMainDAO interface {
	SaveOrderMain()
}
// 订单详情记录
type OrderDetailDAO interface {
	SaveOrderDetail()
}

// 抽象工厂接口
type DAOFactory interface {
	CreaterOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

// 关系数据库 OrderMainDAO 实现
type RDBMainDAO struct {
}
func (*RDBMainDAO) SaveOrderMain ()  {
	fmt.Println("rdb main save...")
}

// 关系数据库 OrderDetailDAO 实现
type RDBDetialDAO struct {
}

func (*RDBDetialDAO) SaveOrderDetail() {
	fmt.Println("rdb detail save... ")
}


// RDB  抽象工厂
type RDBDAOFactory struct {
}
func (*RDBDAOFactory) CreaterOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}
func (*RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetialDAO{}
}

// xml main 存储
type XMLMainDAO struct {
}
func (*XMLDetailDAO) SaveOrderMain() {
	fmt.Println("xml main save")
}

// xml detail 存储
type XMLDetailDAO struct {
}
func (*XMLDetailDAO) SaveOrderDetail(){
	fmt.Println("xml detail save")
}

// xml 抽象工厂实现
type XMLDAOFactory struct {
}
func (*XMLDAOFactory)CreaterOrderMainDAO() OrderMainDAO {
	return &XMLDetailDAO{}
}
func (*XMLDAOFactory)CreateOrderDetailDAO() OrderDetailDAO {
	return &XMLDetailDAO{}
}
