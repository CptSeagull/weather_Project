package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type (
	Post struct {
		Api     string `json:"api_key"`
		Options `json:"options"`
	}

	Options struct {
		Location string `json:"location"`
	}

	RespData struct {
		Data `json:"data"`
	}

	Data struct {
		Index         int    `json:"index,omitempty"`
		Valid_time    string `json:"valid_time,omitempty"`
		Analysis_time string `json:"analysis_time,omitempty"`
	}
)

func main() {

	api, url, typ := setEnv()

	reqData := Post{
		Api: api,
		Options: Options{
			Location: "Sydney",
		},
	}

	jsonData := jsonMarshal(reqData)

	posturl := url + typ

	r := NewRequest(posturl, jsonData)
	AddHeader(r)

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Println("Issue with Client.Do:", err)
	}

	defer res.Body.Close()

	fmt.Println(res.StatusCode)

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
	// fmt.Println(data.Index)
	// fmt.Println(data.Valid_time)
	// fmt.Println(data.Analysis_time)
}
