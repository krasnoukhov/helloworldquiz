package controllers

import (
  "language-game/initializers/redisPool"
  "github.com/astaxie/beego"
  "github.com/garyburd/redigo/redis"
)

type StatsController struct {
  beego.Controller
}

type StatsObject struct {
  Games int
}

func (this *StatsController) Get() {
  conn := redisPool.Get()
  defer conn.Close()
  
  var stats StatsObject
  
  games, _ := redis.Int(conn.Do("HLEN", "games"))
  stats.Games = games
  
  this.Data["json"] = stats
  this.ServeJson()
}
