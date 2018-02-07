// package main 

// import (
// 	//"encoding/json"
// 	"fmt"
// 	"github.com/garyburd/redigo/redis"
// 	// "io/ioutil"
// 	// "net/http"
// 	// "net/url"
// 	//"strings"
// )

// func main() {
// 	f, err := redis.Dial("tcp", "192.168.32.188:6380")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	// type Msg struct {
// 	// 	Id           string
// 	// 	Target_uid   string
// 	// 	Nickname     string
// 	// 	Xing         int
// 	// 	Gold         int
// 	// 	Time         string
// 	// 	Text_id      string
// 	// 	Build_name   int
// 	// 	Action       string
// 	// 	Is_attack    string
// 	// 	Item_text_id string
// 	// }

// 	//values, err := redis.Values(f.Do("lrange", "message_0", "0", "5"))
// 	//values, err := redis.StringMap(f.Do("smembers", "uid"))\
// 	values, err := redis.Values(f.Do("smembers", "uid"))
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	//var str map[string]interface{}
// 	// i := 0
// 	// sum := 0
// 	for _, v := range values {
// 		info := string(v.([]byte))
// 		// 	// //fmt.Println(info)
// 		// 	// err := json.Unmarshal([]byte(info), &str)
// 		// 	// if err != nil {
// 		// 	// 	fmt.Println(err)
// 		// 	// 	return
// 		// 	// }
// 		// 	// fmt.Println(str)
// 		// 	fmt.Println(v)
// 		// 	i++
// 		fmt.Println(info)
// 	}
// 	// fmt.Println(sum)
// }
