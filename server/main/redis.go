//连接池
package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

//定义一个全局的 pool
var pool *redis.Pool

//不会主动调用(链接池)
func initPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {

	pool = &redis.Pool{
		MaxIdle:     maxIdle,     //最大空闲链接数
		MaxActive:   maxActive,   // 表示和数据库的最大链接数，0 表示没有限制
		IdleTimeout: idleTimeout, //最大空闲时间
		Dial: func() (conn redis.Conn, err error) { //初始化连接的代码，连接那个IP的redis
			return redis.Dial("tcp", address)
		},
	}

}
