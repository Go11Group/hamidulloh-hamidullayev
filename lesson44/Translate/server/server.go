package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type TranslateResponse struct {
	EnglishText string
}

type TranslateRequest struct {
	Text string
}

type TranslateServer struct{}

func main() {
	translateServer := new(TranslateServer)

	rpc.Register(translateServer)
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal("Listener server: ", err)
	}
	http.Serve(listener, nil)
}

func (t *TranslateServer) Translate(req *TranslateRequest, res *TranslateResponse) error {
	text := req.Text

	var translator transform.Transformer = transform.Chain(
		norm.NFD, // Unicode-normalizatsiya (NFD forma)
		transform.RemoveFunc(func(r rune) bool {
			return !unicode.IsGraphic(r) // Tozalash uchun shart
		}),
		norm.NFC, // Unicode shakllantirish (NFC forma)
	)

	translatedText, _, err := transform.String(translator, text)
	if err != nil {
		log.Fatal("Translate error? ", err)
	}

	res.EnglishText = translatedText
	return nil
}
