package main

import (
	"testing"
)

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
