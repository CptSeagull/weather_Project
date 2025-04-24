package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Post struct {
	Api     string `json:"api_key"`
	Options info   `json:"options,omitempty"`
}

type info struct {
	Location string `json:"location"`
}

type RespData struct {
	Data kIndex `json:"data"`
}

type kIndex struct {
	Index         string `json:"index"`
	Valid_time    string `json:"valid_time"`
	Analysis_time string `json:"analysis_time"`
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Unable to load the .env file")
	}

	api_Key := os.Getenv("api")

	os.Setenv("Key", "")
	os.Setenv("Loc", "Sydney")

	reqData := Post{
		Api: os.Getenv(api_Key),
		Options: info{
			Location: os.Getenv("Loc"),
		},
	}

	jsonData, err := json.Marshal(reqData)
	if err != nil {
		fmt.Println(("Having an error with Marshal"))
	}

	url := "https://sws-data.sws.bom.gov.au/api/v1/"
	typ := "get-k-index"

	posturl := url + typ

	res, err := http.NewRequest("POST", posturl, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Issue with POST request, %d")
	}

	res.Header.Add("Content-Type", "application/json")

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var data kIndex

	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	fmt.Println(data.Index, data.Valid_time, data.Analysis_time)
	fmt.Println(body)
}
