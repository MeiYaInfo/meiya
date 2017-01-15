package models

import (
	"github.com/astaxie/beego/orm"
)

type NinePic struct {
	Id         int
	PidList    string
	Desc       string
	UpLoaderId int
	CreatorId  int
	UploadTime string
	Favors     int
	Visible    int
}

func init() {
	orm.RegisterModel(new(NinePic))
}
