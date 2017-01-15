package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct { // 用户
	Id        int
	UserName  string
	Password  string
	Authority int
	HeadIcon  string
}

func init() {
	orm.RegisterModel(new(User))
}

func (user User) AddUser() {
	o := orm.NewOrm()
	o.Insert(user)
}

func (user User) UpdateUser() {
	o := orm.NewOrm()
	o.Update(user)
}

func GetUserByUname(userName string) User {
	o := orm.NewOrm()
	var user User
	user.UserName = userName
	o.Read(user)
	return user
}
