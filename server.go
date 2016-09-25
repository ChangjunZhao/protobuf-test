package main

import (
	"encoding/json"
	"encoding/xml"
	"github.com/golang/protobuf/proto"
	"github.com/venusource/protocdemo/user"
	"log"
	"net/http"
)

type User struct {
	Id   int
	Name string
}

func main() {
	http.HandleFunc("/protoc", testProtobuf)
	http.HandleFunc("/json", testJson)
	http.HandleFunc("/xml", testXml)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func testProtobuf(rw http.ResponseWriter, req *http.Request) {
	t := &user.ProtocUser{Id: 1, Name: "cjzhao"}
	data, err := proto.Marshal(t)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	rw.Write(data)
}

func testJson(rw http.ResponseWriter, req *http.Request) {
	user := User{Id: 1, Name: "cjzhao"}
	data, _ := json.Marshal(user)
	rw.Write(data)
}

func testXml(rw http.ResponseWriter, req *http.Request) {
	user := User{Id: 1, Name: "cjzhao"}
	data, _ := xml.Marshal(user)
	rw.Write(data)
}
