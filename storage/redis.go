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

// Filter creation, creates filters for id, hash and urls
/*
func CreateFilters() {
	type Test5 struct {
		Values []string `json:"values"`
	}
	var test Test5
	var hashFilter [2]string
	var idFilter [2]string

	idFilter[0] = "id1"
	idFilter[1] = "id2"
	hashFilter[0] = "hash1"
	hashFilter[1] = "hash2"
	test.Values = make([]string, 3)
	test.Values[0] = "ntnu.edu"
	test.Values[1] = "ntnu.no"
	test.Values[2] = "testsafebrowsing.com/s/malware.html"

	testdata, _ := json.Marshal(test)
	fmt.Println(test)
	_, err := utils.Conn.Do("SET", "urlFilter", testdata)
	if err != nil {
		fmt.Println("Error adding data to redis:" + err.Error())
		logging.Logerror(err, "ERROR adding data to Redis, url filter list creation")
	}

	_, err = utils.Conn.Do("SETEX", "hashFilter", utils.CacheDurationFile, hashFilter)
	if err != nil {
		fmt.Println("Error adding data to redis:" + err.Error())
		logging.Logerror(err, "ERROR adding data to Redis, url filter list creation")
	}

	_, err = utils.Conn.Do("SETEX", "idFilter", utils.CacheDurationFile, idFilter)
	if err != nil {
		fmt.Println("Error adding data to redis:" + err.Error())
		logging.Logerror(err, "ERROR adding data to Redis, url filter list creation")
	}

}
*/

// Check filter takes a filter and an item and checks if the item is whitelisted
/*
func CheckFilter(filter string, item string) bool {

	type test5 struct {
		Values []string `json:"values"`
	}

	value, err := utils.Conn.Do("GET", filter)
	if value == nil {
		fmt.Println("No length")
		return false
	}
	if err != nil {
		return false
	}

	valueB, _ := json.Marshal(value)
	fmt.Println(value)
	var responseData test5
	json.Unmarshal(valueB, &responseData)

	fmt.Println(responseData)
	test32, err := json.Marshal(responseData)
	fmt.Printf("This is it, %q", value)
	fmt.Println("this is test32", string(test32))
	fmt.Println("this is test32", test32)
	fmt.Println("Length: ", len(responseData.Values))
	/*
		test2, err := json.Marshal(value)
		var test []string
		json.Unmarshal(test2, &test)
		fmt.Println(test2)

		fmt.Sprintf("%q", value)
*/
//	return true
//}
