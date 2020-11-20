package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type JsPort struct {
	Port int
}

func GetPort(filePath string, port *JsPort) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &port)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
