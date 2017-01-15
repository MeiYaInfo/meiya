package controllers

import (
	"github.com/astaxie/beego"
)

type UploadController struct {
	beego.Controller
}

func (c *UploadController) Get() {
	f, h, err := c.GetFile("myFile")
	if err != nil {
		c.Data["error"] = err.Error()
		return
	}
	path := "/e/" + h.Filename
	f.Close()
	c.SaveToFile("myFile", path)
	c.Data["result"] = 1
}
