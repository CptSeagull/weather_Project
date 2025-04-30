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

func setEnv() (envApi string, envUrl string, envTyp string) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Unable to load the .env file")
	}

	getApi := os.Getenv("api")
	getUrl := os.Getenv("url")
	getTpy := os.Getenv("typ")

	return getApi, getUrl, getTpy
}

func NewRequest(posturl string, jsonData []byte) (res *http.Request) {
	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Issue with POST request", err)
	}
	return r
}

func AddHeader(r *http.Request) {
	r.Header.Add("Content-Type", "application/json")
}

func jsonMarshal(reqData Post) []byte {
	jsonData, err := json.Marshal(reqData)
	if err != nil {
		fmt.Println(("Having an error with Marshal"))
	}
	return jsonData
}
