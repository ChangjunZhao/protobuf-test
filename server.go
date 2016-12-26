package main

import (
	"encoding/json"
	"encoding/xml"
	"github.com/ChangjunZhao/protobuf-test/user"
	"github.com/golang/protobuf/proto"
	"log"
	"net/http"
)

type User struct {
	Id     int
	Name   string
	Phones []*Phone
}

type Phone struct {
	PhoneType   PhoneType
	PhoneNumber string
}

type PhoneType int

const (
	PHONE_HOME PhoneType = iota
	PHONE_WORK
	PHONE_OTHER
)

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
	t := &user.ProtobufUser{Id: 1, Name: "cjzhao", Phones: []*user.ProtobufUser_Phone{{PhoneType: user.PhoneType_HOME, PhoneNumber: "01080308438"}, {PhoneType: user.PhoneType_WORK, PhoneNumber: "15517684328"}}}
	data, err := proto.Marshal(t)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	rw.Write(data)
}

func testJson(rw http.ResponseWriter, req *http.Request) {
	user := User{Id: 1, Name: "cjzhao", Phones: []*Phone{&Phone{PhoneType: PHONE_HOME, PhoneNumber: "01080308438"}, &Phone{PhoneType: PHONE_WORK, PhoneNumber: "15517684328"}}}
	data, _ := json.Marshal(user)
	rw.Write(data)
}

func testXml(rw http.ResponseWriter, req *http.Request) {
	user := User{Id: 1, Name: "cjzhao", Phones: []*Phone{&Phone{PhoneType: PHONE_HOME, PhoneNumber: "01080308438"}, &Phone{PhoneType: PHONE_WORK, PhoneNumber: "15517684328"}}}
	data, _ := xml.Marshal(user)
	rw.Write(data)
}
