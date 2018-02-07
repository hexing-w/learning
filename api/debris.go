package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"io/ioutil"
)

type JsonFile struct {
	Debrisid int `json:"debrisid"`
}

func readFile(filename string) ([]JsonFile, error) {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("ReadFile:", err.Error())
		return nil, err
	}
	var r []JsonFile
	json.Unmarshal(bytes, &r)
	return r, nil

}
func main() {
	maps, err := readFile("./ssgame_debris.json")
	fmt.Println(maps)
	c, err := redis.Dial("tcp", "192.168.32.188:19001")
	if err != nil {
		return
	}
	defer c.Close()
	redivalue := `{
	  "total": 50,
	  "self": 50,
	  "use": 0
	}`
	for i := 0; i < len(maps); i++ {
		c.Do("hset", "debris_100001945", maps[i].Debrisid, redivalue)
	}
}
