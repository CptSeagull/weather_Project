package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
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
		Data
	}

	Data struct {
		index         string `json:"index"`
		valid_time    string `json:"valid_time"`
		analysis_time string `json:"analysis_time"`
	}
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Unable to load the .env file")
	}

	api_Key := os.Getenv("api")

	reqData := Post{
		Api: api_Key,
		Options: Options{
			Location: "Sydney",
		},
	}

	fmt.Println(reqData)

	jsonData, err := json.Marshal(reqData)
	if err != nil {
		fmt.Println(("Having an error with Marshal"))
	}

	url := "https://sws-data.sws.bom.gov.au/api/v1/"
	typ := "get-k-index"

	posturl := url + typ

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Issue with POST request, %d")
	}

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Content-Type", "charset=UTF-8")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		fmt.Println("Issue with client.Do")
		return
	}

	defer res.Body.Close()

	fmt.Println(res.StatusCode)

	data := &Data{}
	derr := json.NewDecoder(res.Body).Decode(data)
	if derr != nil {
		panic(derr)
	}

	fmt.Println(string(data.index))
}
