package storage

import (
	logging "dcsg2900-threattotal/logs"
	"dcsg2900-threattotal/utils"
	"fmt"
	"os"

	"github.com/gomodule/redigo/redis"
)

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
				os.Getenv("redisUrl"),
				redis.DialPassword(os.Getenv("redisPassword")))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

// Add to pool function takes a key, a timeout and some data and
// adds it to the redis pool as a new key value pair.
func AddToPool(key string, timeout int, data string) {
	response, err := utils.Conn.Do("SETEX", key, timeout, data)
	if err != nil {
		fmt.Println("Error:" + err.Error())
		logging.Logerror(err, "Error adding to redis pool redis.go:")

	}
	// Print the response to adding the data (should be "OK"
	fmt.Println(response)
}
