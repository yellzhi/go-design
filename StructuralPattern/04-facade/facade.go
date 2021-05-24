package main

import (
	"fmt"
	"golang.org/x/xerrors"
)

/*
外观模式(Facade Pattern)：外部与一个子系统的通信必须通过一个统一的外观对象进行，
为子系统中的一组接口提供一个一致的界面，外观模式定义了一个高层接口，
这个接口使得这一子系统更加容易使用。外观模式又称为门面模式，它是一种对象结构型模式。

外观模式包含如下角色：
Facade: 外观角色
SubSystem:子系统角色

 */


/*
/假设现在我有一个网站，以前有登录和注册的流程，
//登录的时候调用用户的查询接口，注册时调用用户的创建接口。
//为了简化用户的使用流程，
//我们现在提供直接验证码登录/注册的功能，如果该手机号已注册那么我们就走登录流程，
如果该手机号未注册，那么我们就创建一个新的用户。
 */

type Iuser interface {
	Login(name string, pwd string)(id int, err error)
	Register(name string, pwd string)(id int, err error)
}

type User struct {
	Name string
	PWD string
	Id int
}
func (u User) Login(name string, pwd string) (id int, err error) {
	if name =="t1"{
		id, err = 1, nil
		return
	}
	return 0, xerrors.Errorf("new user")
}
func (u *User) Register(name string, pwd string) (id int, err error) {
	if "t1" != name{
		return 1, nil
	}
	return 0, xerrors.Errorf("old user")
}
var _ Iuser = &User{}

func (u *User) LoginRegister(name string, pwd string) (id int, err error){
	id, err = u.Login(name, pwd)
	if err == nil {
		return 0, nil
	}
	if id ==0{
		id , err = u.Register(name, pwd)
	}
	return
}

type UserFacade interface {
	LoginOrRegister(name string, pwd string)(int, error)
}
type NewUser struct {
	u User
}
func (n NewUser) LoginOrRegister(name string, pwd string) (int, error) {
	return  n.u.LoginRegister(name, pwd)
}

var _ UserFacade = &NewUser{}
func NewUserFacade() UserFacade{
	return &NewUser{}
}

func main()  {
	id , err:= NewUserFacade().LoginOrRegister("t1", "222")
	fmt.Println(id, err)

	id , err = NewUserFacade().LoginOrRegister("t2", "222")
	fmt.Println(id, err)
}