package main

import (
	"encoding/json"
	"fmt"
	"io"
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

	if res.StatusCode != 200 {
		fmt.Println("Issue with HTTP request: ", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)

	fmt.Println(sb)

	data := &RespData{}
	derr := json.NewDecoder(res.Body).Decode(data)
	if derr != nil {
		log.Fatalln(derr)
	}
	fmt.Println(data.Data.Index)
	fmt.Println(data.Data.Valid_time)
	fmt.Println(data.Analysis_time)
}
