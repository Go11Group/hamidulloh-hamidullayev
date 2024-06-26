package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	pb "dodi/protos/genproto/protos" // Protokol paketini Go tiliga ko'chirish
)

const (
	address = "localhost:50051" // Qo'llab-quvvatlash serverining manzili
)

func main() {
	// Serverga ulanish uchun bir grpc ulanish yaratish
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewWeatherServiceClient(conn) // Protokol bilan gRPC Clientini yaratish

	// GetCurrentWeather testi
	currentWeatherResponse, err := c.GetCurrentWeather(context.Background(), &pb.CurrentWeatherRequest{Location: "New York"})
	if err != nil {
		log.Fatalf("Error calling GetCurrentWeather: %v", err)
	}
	log.Printf("Current weather in %s: %s, %.1f°C", currentWeatherResponse.Location, currentWeatherResponse.Obuhavo, currentWeatherResponse.Harorat)

	// GetWeatherForecast testi
	forecastResponse, err := c.GetWeatherForecast(context.Background(), &pb.ForecastRequest{Location: "London", Kunlar: 3})
	if err != nil {
		log.Fatalf("Error calling GetWeatherForecast: %v", err)
	}
	log.Printf("Weather forecast for %s:", forecastResponse.Location)
	for _, forecast := range forecastResponse.Forecast {
		log.Printf("Date: %s, Weather: %s", forecast.Date, forecast.Weather)
	}

	// ReportWeatherCondition testi
	reportResponse, err := c.ReportWeatherCondition(context.Background(), &pb.ConditionReport{Location: "Paris", Weather: "Cloudy", Temperature: 20})
	if err != nil {
		log.Fatalf("Error calling ReportWeatherCondition: %v", err)
	}
	log.Printf("Weather condition reported: %v", reportResponse.Success)
}
