package controllers

import (
  "github.com/krasnoukhov/beego"
)

type MainController struct {
  beego.Controller
}

func (this *MainController) Get() {
  this.Data["title"] = "Hello World Quiz"
  this.Data["caption"] = "Guess programming language by „Hello, world!“ snippet"
  this.Data["description"] = "Guess programming language by „Hello, world!“ snippet"
  this.Data["host"] = this.Ctx.Request.Host
  this.TplNames = "index.tpl"
}
