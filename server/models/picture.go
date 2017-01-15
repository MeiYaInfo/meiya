package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Picture struct {
	Id         int
	PUrl       string
	UploadTime time.Time
}

func init() {
	orm.RegisterModel(new(Picture))
}
