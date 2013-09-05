package main

import (
  "language-game/controllers"
  "github.com/astaxie/beego"
)

func main() {
  beego.SessionProvider = "redis"
  beego.SessionSavePath = beego.AppConfig.String("redis")
  
  beego.Router("/", &controllers.MainController{})
  beego.RESTRouter("/game", &controllers.GameController{})
  beego.Router("/stats", &controllers.StatsController{})
  beego.Run()
}
