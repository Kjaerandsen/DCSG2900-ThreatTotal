package Storage

import (
	"github.com/gomodule/redigo/redis"
)

// Alternatively use https://github.com/go-redis/redis

// TODO: move theese to the utils library when done testing
const Host = "10.0.0.42"
const Port = "6379"

// TODO: password in authentication file or as system args?
const Password = "admin"

// password can be set in the redis-cli using the command: 'CONFIG SET requirepass "password"'

// If the server is on another machine you need to set protected-mode to off
// in redis-cli 'CONFIG SET protected-mode no'

// InitPool initializes the storage pool used in the application
// Called from main, all other functions dealing with cache done in main?
func InitPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp",
				Host+":"+Port,
				redis.DialPassword(Password))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}
