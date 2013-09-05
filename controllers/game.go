package controllers

import (
  "fmt"
  "errors"
  "encoding/json"
  "language-game/models/game"
  "language-game/models/variant"
  "github.com/astaxie/beego"
)

type GameController struct {
  beego.Controller
}

type GameResponse struct {
  Game *game.Object
  Variant *variant.Object
}

func (this *GameController) Post() {
  beego.Debug(variant.Objects)
  var newObject game.Object
  json.Unmarshal(this.Ctx.RequestBody, &newObject)
  
  object, err := game.Add(&newObject)
  if err == nil {
    this.SetSession("GameObjectId", object.ObjectId)
    this.Data["json"] = &GameResponse{ object, game.GetVariant(object) }
  } else {
    this.Data["json"] = map[string]string{ "Error": fmt.Sprint(err) }
  }
  
  this.ServeJson()
}

func (this *GameController) Get() {
  object, err := Game(this)
  
  if err == nil {
    this.Data["json"] = &GameResponse{ object, game.GetVariant(object) }
  } else {
    this.Data["json"] = map[string]string{ "Error": fmt.Sprint(err) }
  }
  
  this.ServeJson()
}

func (this *GameController) Put() {
  var updatedObject game.Object
  json.Unmarshal(this.Ctx.RequestBody, &updatedObject)
  
  object, err := Game(this)
  if err == nil {
    object, err = game.Update(object.ObjectId, &updatedObject)
    
    if err == nil {
      this.Data["json"] = &GameResponse{ object, game.GetVariant(object) }
    } else {
      this.Data["json"] = map[string]string{ "Error": fmt.Sprint(err) }
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
