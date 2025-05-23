package main

type (
	Post struct {
		Api     string `json:"api_key"`
		Options `json:"options"`
	}

	Options struct {
		Location string `json:"location"`
	}

	RespData struct {
		Data [][]Data `json:"data"`
	}

	Data struct {
		Index         int    `json:"index,string"`
		Valid_time    string `json:"valid_time,omitempty"`
		Analysis_time string `json:"analysis_time,omitempty"`
	}
)
