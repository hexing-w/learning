package main

import (
	// "bytes"
	// "encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	// "io/ioutil"
	// "log"
	// "net/http"
	// "net/url"
)

func main() {
	// u, _ := url.Parse("http://192.168.41.23:85/AttackBuild")

	redigo, err := redis.Dial("tcp", "192.168.32.188:6380")
	if err != nil {
		fmt.Println(err)
	}
	i := 0
	values, err := redis.Values(redigo.Do("smembers", "uid"))
	for _, v := range values {
		info := string(v.([]byte))
		if i < 1 {
			array := map[string]string{
				"attack_uid": "100004386",
				"star":       "60001",
				// "uid":"100001856",
				"build": "130011",
				"token": "bf1484b65a7aef02c40c2ad1c4ec1ac52955392c725c6523aa853f2e7dd23f83",
			}
			array["uid"] = info
			fmt.Println(array)
			i++

		}

	}

}
