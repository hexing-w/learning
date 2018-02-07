package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type StarInfo struct {
	Attack_uid string
	// star       string
	// uid        string
	// build      string
}

func main() {
	// for i := 0; i < 10; i++ {
	param := make(map[string]interface{})
	param["attack_uid"] = "100001517"
	param["star"] = "60001"
	param["uid"] = "100001093"
	param["build"] = "130011"
	b, err := json.Marshal(param)
	if err != nil {
		fmt.Println(err)
	}
	body := bytes.NewBuffer([]byte(b))
	res, err := http.Post("http://192.168.40.34:85/AttackBuild", "application/json;charset=utf-8", body)
	if err != nil {
		log.Fatal(err)
		return
	}
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	res.Body.Close()
	fmt.Println(string(result))
	// }

}
