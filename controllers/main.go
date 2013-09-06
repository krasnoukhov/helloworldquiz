package controllers

import (
  "github.com/astaxie/beego"
)

type MainController struct {
  beego.Controller
}

func (this *MainController) Get() {
  this.Data["title"] = "Programming Languages Quiz"
  this.Data["host"] = this.Ctx.Request.Host
  this.TplNames = "index.tpl"
}
