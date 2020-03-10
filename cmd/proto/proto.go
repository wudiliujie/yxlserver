package main

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"yxlserver/cmd/proto/generate"
)

func main() {
	f, err := os.Open(`..\services\msg\proto.xml`)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	// v := make(map[string]interface{})
	var v = &generate.Packages{}
	err = xml.Unmarshal(data, v)
	v.Init()
	if err != nil {
		panic(err)
	}
	v.MakeGo(`..\services\msg\msg.pb.go`)

}
