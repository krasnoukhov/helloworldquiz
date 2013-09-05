package redisPool

import (
  "time"
  "github.com/astaxie/beego"
  "github.com/garyburd/redigo/redis"
)

var (
  redisPool *redis.Pool
)

func init() {
  poolSize, err := beego.AppConfig.Int("redis-pool-size")
  if err != nil {
    beego.Error("Pool size is not specified")
  }
  
  redisPool = &redis.Pool{
    MaxIdle: poolSize,
    IdleTimeout: 240 * time.Second,
    Dial: func () (redis.Conn, error) {
      conn, err := redis.Dial("tcp", beego.AppConfig.String("redis"))
      if err != nil {
        return nil, err
      }
      return conn, err
    },
    TestOnBorrow: func (conn redis.Conn, t time.Time) (error) {
      _, err := conn.Do("PING")
      return err
    },
  }
}

func Get() (redis.Conn) {
  return redisPool.Get()
}
