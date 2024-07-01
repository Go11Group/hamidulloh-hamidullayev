package handler

import "dodi/genproto/protos"

type Handler struct {
	Weather   protos.WeatherServiceClient
	Transport protos.TransportServiceClient
}

func NewHandler(Weather protos.WeatherServiceClient, Transport protos.TransportServiceClient) *Handler {
	return &Handler{Weather: Weather, Transport: Transport}
}