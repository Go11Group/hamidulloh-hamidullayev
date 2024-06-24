package main

import "github.com/Go11Group/lesson43/api_gateway_service/api"

func main() {
	api.Routes().ListenAndServe()
}
