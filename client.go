package main

import (
	"encoding/json"
	"encoding/xml"
	"github.com/ChangjunZhao/protobuf-test/user"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"
)

func requestProtoc() {
	client := &http.Client{}
	resp, err := client.Get("http://localhost:8080/protoc")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	t := &user.ProtobufUser{}
	proto.Unmarshal(data, t)
}

func requestJson() {
	client := &http.Client{}
	resp, err := client.Get("http://localhost:8080/json")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	jsonuser := &User{}
	json.Unmarshal(data, jsonuser)
}

func requestXml() {
	client := &http.Client{}
	resp, err := client.Get("http://localhost:8080/xml")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	xmluser := &User{}
	xml.Unmarshal(data, xmluser)
}
