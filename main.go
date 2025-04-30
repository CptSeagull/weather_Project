package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	api, url, typ := setEnv()

	reqData := Post{
		Api: api,
		Options: Options{
			Location: "Australian region",
		},
	}

	jsonData := jsonMarshal(reqData)

	posturl := url + typ

	r := NewRequest(posturl, jsonData)
	AddHeader(r)

	res := Client(r)

	defer res.Body.Close()

	data := &RespData{}
	derr := json.NewDecoder(res.Body).Decode(data)
	if derr != nil {
		log.Fatalln(derr)
	}
	fmt.Println(data.Data)
}
