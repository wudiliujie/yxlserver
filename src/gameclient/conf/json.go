package conf

import (
	"encoding/json"
	"io/ioutil"
	"leaf/log"
)

var Client struct {
	LoginAddr string
	ClientNum int
}

func init() {
	data, err := ioutil.ReadFile("conf/client.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Client)
	if err != nil {
		log.Fatal("%v", err)
	}
}
