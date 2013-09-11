package controllers

import (
  "fmt"
  "errors"
  // "encoding/json"
  "langgame/initializers/redisPool"
  "langgame/models/game"
  "langgame/models/variant"
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
    this.SetSession("GameObjectId", object.ObjectId)
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
  objectId := this.GetSession("GameObjectId")
  
  if objectId != nil {
    object, err := game.Get(objectId.(string))
    
    if err == nil {
      return object, nil
    } else {
      return nil, err
    }
  } else {
    return nil, errors.New("No game")
  }
}
