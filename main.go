package main

import (
  "os"
  "net/http"
  "strings"
  "io/ioutil"
  "langgame/controllers"
  "github.com/krasnoukhov/beego"
  "github.com/krasnoukhov/train"
)

func main() {
  // Config
  if os.Getenv("GO_ENV") == "prod" {
    beego.RunMode = "prod"
  }
  // beego.SessionProvider = "redis"
  // beego.SessionSavePath = beego.AppConfig.String("redis")
  
  // Assets 
  beego.SetStaticPath("/public", "public")
  beego.AddFuncMap("javascript_tag", train.JavascriptTag)
  beego.AddFuncMap("stylesheet_tag", train.StylesheetTag)
  beego.AddFuncMap("app", func () (out string) {
    source, _ := ioutil.ReadFile("views/app.hbs")
    return string(source[:])
  })
  
  // Routes
  beego.Router("/", &controllers.MainController{})
  beego.RESTRouter("/game", &controllers.GameController{})
  beego.Router("/stats", &controllers.StatsController{})
  
  train.SetFileServer()
  beego.Errorhandler("404", func(w http.ResponseWriter, r *http.Request) {
    if strings.HasPrefix(r.URL.Path, "/assets") {
      train.ServeRequest(w, r)
    } else {
      http.Redirect(w, r, "/", 302)
    }
  })
  
  beego.Run()
}
