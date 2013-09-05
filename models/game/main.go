package game

import (
  "errors"
  "encoding/json"
  "math/rand"
  "langgame/initializers/redisPool"
  "langgame/models/variant"
  "github.com/astaxie/beego"
  "github.com/dchest/uniuri"
  "github.com/garyburd/redigo/redis"
)

type Object struct {
  ObjectId   string
  Score      int
  Lives      int
  Completed  []string `json:"-"`
  Current    string   `json:"-"`
}

type DumpObject struct {
  ObjectId   string
  Score      int
  Lives      int
  Completed  []string
  Current    string
}

func Add() (object *Object, err error) {
  object = &Object{}
  object.ObjectId = uniuri.New()
  object.Score = 0
  object.Lives = 3
  object.Completed = []string{}
  
  err = Set(object)
  return object, err
}

func Get(objectId string) (response *Object, err error) {
  conn := redisPool.Get()
  defer conn.Close()
  
  data, err := redis.Bytes(conn.Do("HGET", "games", objectId))
  if err != nil {
    beego.Critical(err)
    return nil, errors.New("Can't get object")
  }
  
  object := &Object{}
  err = Restore(data, object)
  if err != nil {
    return nil, errors.New("Can't restore object")
  }
  
  return object, nil
}

func GetVariant(object *Object) (response *variant.Object) {
  if object.Current == "" {
    available := Diff(object.Completed, variant.Keys)
    object.Current = ""
    
    if len(available) > 0 {
      i := rand.Intn(len(available))
      object.Current = available[i]
    }
    
    Set(object)
  }
  
  if object.Current == "" {
    return nil
  }else{
    return variant.Shuffle(variant.Objects[object.Current])
  }
}

func SetVariant(object *Object, option string) {
  if object.Current != "" {
    conn := redisPool.Get()
    defer conn.Close()
    
    if object.Current == option {
      object.Score += 50
      conn.Do("HINCRBY", "success", object.Current, 1)
    } else {
      object.Lives -= 1
      conn.Do("HINCRBY", "failure", object.Current, 1)
    }
    
    object.Completed = append(object.Completed, object.Current)
    object.Current = ""
    
    Set(object)
  }
}

func Set(object *Object) (err error) {
  conn := redisPool.Get()
  defer conn.Close()
  
  data := Dump(object)
  _, err = conn.Do("HSET", "games", object.ObjectId, data)
  if err != nil {
    beego.Critical(err)
    return errors.New("Can't store object")
  }
  
  return nil
}

func Dump(object *Object) (dump string) {
  dumpObject := &DumpObject{ object.ObjectId, object.Score, object.Lives, append([]string{}, object.Completed...), object.Current }
  data, _ := json.Marshal(dumpObject)
  return string(data[:])
}

func Restore(dump []byte, object *Object) (err error) {
  dumpObject := &DumpObject{}
  err = json.Unmarshal(dump, &dumpObject)
  
  if err == nil {
    object.ObjectId = dumpObject.ObjectId
    object.Score = dumpObject.Score
    object.Lives = dumpObject.Lives
    object.Completed = dumpObject.Completed
    object.Current = dumpObject.Current
  }
  
  return err
}

func Diff(first, second []string) ([]string) { 
  count := make(map[string]int)
  for _, x := range first {
    _, present := count[x]
    if !present { 
      count[x] = 0
    } 
    count[x]++
  }
  
  for _, x := range second { 
    _, present := count[x]
    if !present {
      count[x] = 0;
    }
    count[x]++
  }
  
  var result []string
  for k, v := range count {
    if v < 2 {
      result = append(result, k)
    }
  }
  
  return result
}
