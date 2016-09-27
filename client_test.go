package main

import (
	"fmt"
	"github.com/ChangjunZhao/protobuf-test/user"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func Test_protoc(t *testing.T) {
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
	user := &user.ProtobufUser{}
	proto.Unmarshal(data, user)
	fmt.Printf("%d %s\n", user.Id, user.Name)
	for _, phone := range user.Phones {
		fmt.Printf("%s:%s\n", phone.PhoneType, phone.PhoneNumber)
	}
}

func Benchmark_protoc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		requestProtoc()
	}
}

func Benchmark_json(b *testing.B) {
	for i := 0; i < b.N; i++ {
		requestJson()
	}
}

func Benchmark_xml(b *testing.B) {
	for i := 0; i < b.N; i++ {
		requestXml()
	}
}
