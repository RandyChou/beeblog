package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"models"
)

type DeleteController struct {
	beego.Controller
}

func (this *DeleteController) Get() {
	id, _ := strconv.Atoi(this.Ctx.Params[":id"])
	models.DelBlog(models.GetBlog(id))
	this.Ctx.Redirect(302, "/")
}