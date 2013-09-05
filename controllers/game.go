package controllers

import (
  "fmt"
  "errors"
  // "encoding/json"
  "langgame/initializers/redisPool"
  "langgame/models/game"
  "langgame/models/variant"
  "github.com/astaxie/beego"
  "github.com/garyburd/redigo/redis"
)

type GameController struct {
  beego.Controller
}

type GameResponse struct {
  Game      *game.Object
  Variant   *variant.Object     `json:",omitempty"`
  Correct   *variant.DumpObject `json:",omitempty"`
  Status    string              `json:",omitempty"`
}

func (this *GameController) Post() {
  object, err := game.Add()
  if err == nil {
    this.SetSession("GameObjectId", object.ObjectId)
    this.Data["json"] = &GameResponse{ object, game.GetVariant(object), nil, "ready" }
  } else {
    this.Data["json"] = map[string]string{ "Error": fmt.Sprint(err) }
  }
  
  this.ServeJson()
}

func (this *GameController) Get() {
  object, err := Game(this)
  
  if err == nil {
    this.Data["json"] = &GameResponse{ object, game.GetVariant(object), nil, "ready" }
  } else {
    this.Data["json"] = map[string]string{ "Error": fmt.Sprint(err) }
  }
  
  this.ServeJson()
}

func (this *GameController) Put() {
  object, err := Game(this)
  
  if err == nil {
    conn := redisPool.Get()
    defer conn.Close()
    
    option := this.GetString("Option")
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
      conn.Do("SET", "highest", object.Score)
    }
    
    if object.Lives <= 0 {
      this.Data["json"] = &GameResponse{ object, nil, correct, "died" }
      conn.Do("INCR", "died")
    } else {
      variant := game.GetVariant(object)
      
      if variant != nil {
        this.Data["json"] = &GameResponse{ object, game.GetVariant(object), correct, "ready" }
      } else {
        this.Data["json"] = &GameResponse{ object, nil, correct, "survived" }
        conn.Do("INCR", "survived")
      }
    }
  } else {
    this.Data["json"] = map[string]string{ "Error": fmt.Sprint(err) }
  }
  
  this.ServeJson()
}

func Game(this *GameController) (object *game.Object, err error) {
  objectId := this.GetSession("GameObjectId")
  
  if objectId != nil {
    object, err := game.Get(string(objectId.([]byte)[:]))
    
    if err == nil {
      return object, nil
    } else {
      return nil, err
    }
  } else {
    return nil, errors.New("No game")
  }
}
