package controllers

import (
  "github.com/astaxie/beego"
)

type MainController struct {
  beego.Controller
}

func (this *MainController) Get() {
  this.Data["title"] = "Hello World Quiz"
  this.Data["caption"] = "Programming Languages 'Hello World' Quiz"
  this.Data["host"] = this.Ctx.Request.Host
  this.TplNames = "index.tpl"
}
