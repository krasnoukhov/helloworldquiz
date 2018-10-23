package controllers

import (
  "fmt"
  "errors"
  "net/http"
  "net/url"
  "time"
  // "encoding/json"
  "helloworldquiz/initializers/redisPool"
  "helloworldquiz/models/game"
  "helloworldquiz/models/variant"
  "github.com/krasnoukhov/beego"
  "github.com/garyburd/redigo/redis"
)

type GameController struct {
  beego.Controller
}

type GameResponse struct {
  Game      *game.Object        `json:"game"`
  Variant   *variant.Object     `json:"variant"`
  Correct   *variant.DumpObject `json:"correct"`
  Status    string              `json:"status"`
}

func (this *GameController) Post() {
  object, err := game.Add()
  if err == nil {
    cookie := http.Cookie{ Name: "GameObjectId",
      Value:    url.QueryEscape(object.ObjectId),
      Path:     "/",
      HttpOnly: true,
      Secure:   false,
      Expires:  time.Now().Add(time.Duration(31*24*3600) * time.Second) }
    http.SetCookie(this.Ctx.ResponseWriter, &cookie)
    
    this.Data["json"] = &GameResponse{ object, game.GetVariant(object), nil, "ready" }
  } else {
    this.Data["json"] = map[string]string{ "error": fmt.Sprint(err) }
  }
  
  this.ServeJson()
}

func (this *GameController) Get() {
  object, err := Game(this)
  
  if err == nil {
    status := "ready"
    if object.Lives <= 0 {
      status = "died"
    }else if len(object.Completed) == len(variant.Objects) {
      status = "survived"
    }
    
    this.Data["json"] = &GameResponse{ object, game.GetVariant(object), nil, status }
  } else {
    this.Data["json"] = map[string]string{ "error": fmt.Sprint(err) }
  }
  
  this.ServeJson()
}

func (this *GameController) Put() {
  object, err := Game(this)
  
  if err == nil {
    conn := redisPool.Get()
    defer conn.Close()
    
    option := this.GetString("option")
    correct := &variant.DumpObject{}
    
    if object.Current != "" {
      correct = variant.ConvertToDumpObject(variant.Objects[object.Current])
      if object.Current == option {
        correct = nil
      }
    } else {
      correct = nil
    }
    
    game.SetVariant(object, option)
    
    highest, _ := redis.Int(conn.Do("GET", "highest"))
    if object.Score > highest {
      if _, err := conn.Do("SET", "highest", object.Score); err != nil {
        beego.Critical(err)
      }
    }
    
    if object.Lives <= 0 {
      this.Data["json"] = &GameResponse{ object, nil, correct, "died" }
      if _, err := conn.Do("INCR", "died"); err != nil {
        beego.Critical(err)
      }
    } else {
      variant := game.GetVariant(object)
      
      if variant != nil {
        this.Data["json"] = &GameResponse{ object, game.GetVariant(object), correct, "ready" }
      } else {
        this.Data["json"] = &GameResponse{ object, nil, correct, "survived" }
        if _, err := conn.Do("INCR", "survived"); err != nil {
          beego.Critical(err)
        }
      }
    }
  } else {
    this.Data["json"] = map[string]string{ "error": fmt.Sprint(err) }
  }
  
  this.ServeJson()
}

func Game(this *GameController) (object *game.Object, err error) {
  cookie, err := this.Ctx.Request.Cookie("GameObjectId")
  
  if err == nil && cookie.Value != "" {
    objectId, _ := url.QueryUnescape(cookie.Value)
    object, err := game.Get(objectId)
    
    if err == nil {
      return object, nil
    } else {
      return nil, err
    }
  } else {
    return nil, errors.New("No game")
  }
}

