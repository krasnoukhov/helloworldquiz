package game

import (
  "errors"
  "encoding/json"
  // "strconv"
  // "time"
  "langgame/initializers/redisPool"
  "langgame/models/variant"
  // "github.com/astaxie/beego"
  "github.com/dchest/uniuri"
  "github.com/garyburd/redigo/redis"
)

type Object struct {
  ObjectId   string
  Score      int32
  Lives      int32
}

func Add(newObject *Object) (object *Object, err error) {
  newObject.ObjectId = uniuri.New()
  newObject.Score = 0
  newObject.Lives = 3
  
  return Set(newObject)
}

func Get(objectId string) (object *Object, err error) {
  conn := redisPool.Get()
  defer conn.Close()
  
  data, err := redis.Bytes(conn.Do("HGET", "games", objectId))
  if err != nil {
    return nil, errors.New("Can't get object")
  }
  
  err = json.Unmarshal(data, &object)
  if err != nil {
    return nil, errors.New("Can't restore object")
  }
  
  return object, nil
}

func GetVariant(object *Object) (response *variant.Object) {
  return variant.Get(variant.Objects["c"])
}

func Update(ObjectId string, updatedObject *Object) (object *Object, err error) {
  object, err = Get(ObjectId)
  if err != nil {
    return nil, err
  }
  
  object.Score = updatedObject.Score
  object.Lives = updatedObject.Lives
  
  return Set(object)
}

func Set(setObject *Object) (object *Object, err error) {
  conn := redisPool.Get()
  defer conn.Close()
  
  data := Dump(setObject)
  _, err = conn.Do("HSET", "games", setObject.ObjectId, data)
  if err != nil {
    return nil, errors.New("Can't store object")
  }
  
  return setObject, nil
}

func Dump(object *Object) (dump string) {
  data, _ := json.Marshal(object)
  return string(data[:])
}
