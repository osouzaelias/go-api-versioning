package main

import (
	"encoding/json"
	"net/http"
)

type (
	PayloadV1 struct {
		Message string `json:"message"`
	}

	PayloadV2 struct {
		Message string `json:"message"`
		Author  string `json:"author"`
	}
)

type handlerFunc func(http.ResponseWriter, *http.Request) error

func handleV1(w http.ResponseWriter, _ *http.Request) error {
	data := getData()
	payload := PayloadV1{
		Message: data["message"],
	}
	return json.NewEncoder(w).Encode(payload)
}

func handleV2(w http.ResponseWriter, _ *http.Request) error {
	data := getData()
	payload := PayloadV2{
		Message: data["message"],
		Author:  data["author"],
	}
	return json.NewEncoder(w).Encode(payload)
}
