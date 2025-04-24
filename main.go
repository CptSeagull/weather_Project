package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Post struct {
	Api     string `json:"api_key"`
	Options info   `json:"options,omitempty"`
}

type info struct {
	Location string `json:"location"`
}

type RespData struct {
	Data kIndex `json:"data,omitempty"`
}

type kIndex struct {
	Index      string `json:"index"`
	Valid_time string `json:"valid_time"`
}

func main() {

	reqData := Post{
		Api: "716b58a2-ff50-4cfd-aa72-aa81a50ef62d",
		Options: info{
			Location: "Sydney",
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
	// fmt.Println(r.Body)

	// r.Header.Add("Content-Type", "application/json")

	// client := &http.Client{}
	// res, err := client.Do(r)
	// if err != nil {
	// 	fmt.Println("Issue with client.Do, %d")
	// }

	defer res.Body.Close()

	presp, prerr := io.ReadAll(res.Body)
	if prerr != nil {
		panic(err)
	}
	fmt.Println(string(presp))

	// post := RespData{
	// 	Data: kIndex{
	// 		Index:      "",
	// 		Valid_time: "",
	// 	},
	// }
	// derr := json.NewDecoder(res.Body).Decode(post)
	// if derr != nil {
	// 	panic("This is the decoder error, %d")
	// }

	// fmt.Println("This is the status code: ", res.StatusCode)
	// fmt.Println("Data:", post.Data)
}
