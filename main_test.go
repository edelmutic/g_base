package main

import (
	"net/http"
	"testing"
)


func TestHandleFib(t *testing.T) {
testCase:= []struct {
	name string
	num int
	want []byte
} {
	{
	name: "zero",
	num: 0,
	want: []byte("0"),
}, {
	name: "one",
	num: 1,
	want: []byte("1"),
},
}
handler := http.HandlerFunc(handleFib)
}