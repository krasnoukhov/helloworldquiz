package controllers

import (
  "fmt"
  // "math"
  "langgame/models/variant"
  "langgame/initializers/redisPool"
  "github.com/astaxie/beego"
  "github.com/garyburd/redigo/redis"
)

type StatsController struct {
  beego.Controller
}

type StatsObject struct {
  Games           string
  CompletionRate  string
  Easiest         *StatsPair
  Hardest         *StatsPair
  Highest         string
}

type StatsPair struct {
  Key             string
  Value           int
}

func (this *StatsController) Get() {
  conn := redisPool.Get()
  defer conn.Close()
  
  var stats StatsObject
  games, _ := redis.Int(conn.Do("HLEN", "games"))
  stats.Games = fmt.Sprintf("%v", games)
  
  died, _ := redis.Int(conn.Do("GET", "died"))
  survived, _ := redis.Int(conn.Do("GET", "survived"))
  stats.CompletionRate = fmt.Sprintf("%.1f", (float64(died + survived) / float64(games)) * 100)
  
  highest, _ := redis.Int(conn.Do("GET", "highest"))
  stats.Highest = fmt.Sprintf("%v", highest)
  
  stats.Easiest = FindMaxVariant(conn, "success")
  stats.Hardest = FindMaxVariant(conn, "failure")
  
  this.Data["json"] = stats
  this.ServeJson()
}

func FindMaxVariant(conn redis.Conn, hash string) (response *StatsPair) {
  variants, err := redis.Values(conn.Do("HGETALL", hash))
  response = &StatsPair{}
  
  for k, _score := range variants {
    if k % 2 == 1 {
      score, _ := redis.Int(_score, err)
      if score > response.Value {
        key, _ := redis.String(variants[k-1], err)
        
        response.Key = variant.Objects[key].Name
        response.Value = score
      }
    }
  }
  
  return response
}
