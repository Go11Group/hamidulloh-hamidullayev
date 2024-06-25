package main

import (
	"log"
	"net/rpc"
)

type TranslateRequest struct {
	Text string
}

type TranslateResponse struct {
	EnglishText string
}

func main() {
	uzbekText := TranslateRequest{
		Text: "salom",
	}

	response := TranslateResponse{}

	client, err := rpc.DialHTTP("tcp", ":8000")
	if err != nil {
		log.Fatal("Client connection error: ", err)
	}

	args := &uzbekText

	err = client.Call("TranslateServer.Translate", args, &response)
	if err != nil {
		log.Fatal("Client invocation error: ", err)
	}

	log.Println("Translation from Uzbek to English:", response.EnglishText)
}
