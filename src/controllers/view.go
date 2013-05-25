package controllers

import (
	"strconv"
	"github.com/astaxie/beego"
	"models"
)

type ViewController struct {
	beego.Controller
}

func (this *ViewController) Get() {
	id, _ := strconv.Atoi(this.Ctx.Params[":id"])
	this.Data["Post"] = models.GetBlog(id)
	this.Layout = "layout.tpl"
	this.TplNames = "view.tpl"
}