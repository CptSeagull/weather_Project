package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Post struct {
	data string `json:"data"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading file")
	}
	// url := os.Getenv("url")
	// api := os.Getenv("api_key")
	// option := os.Getenv("option")

	url := os.Getenv("url2")
	api := os.Getenv("api_key2")
	typ := os.Getenv("type")
	q := os.Getenv("location")

	// body := []byte(`{
	// 	"locations": {
	// 		"q: "Sydney"}
	// }`)

	posturl := url + typ + api + q

	r, err := http.NewRequest("POST", posturl, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Body)

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Body)
	defer res.Body.Close()

	post := &Post{}
	deff := json.NewDecoder(res.Body).Decode(post)
	if deff != nil {
		panic(deff)
	}

	if res.StatusCode != http.StatusCreated {
		panic(res.Status)
	}
	fmt.Println(res.Body)
	fmt.Println("Data:", post.data)
}
