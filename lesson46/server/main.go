package main

import (
	"context"
	"log"
	"net"

	pb "dodi/protos/genproto/protos" // Protokol paketini Go tiliga ko'chirish

	"google.golang.org/grpc"
)

// Server struktura yaratish
type server struct {
	pb.UnsafeWeatherServiceServer
}

// GetCurrentWeather metodini serverga qo'shish
func (s *server) GetCurrentWeather(ctx context.Context, in *pb.CurrentWeatherRequest) (*pb.CurrentWeatherResponse, error) {
	log.Printf("Received GetCurrentWeather request for location: %s", in.Location)
	// Bu joyda asosiy loyihaga ma'lumotlarni olish va javob qaytarish kerak
	return &pb.CurrentWeatherResponse{
		Location: in.Location,
		Obuhavo:  "Sunny",
		Harorat:  24.5,
	}, nil
}

// GetWeatherForecast metodini serverga qo'shish
func (s *server) GetWeatherForecast(ctx context.Context, in *pb.ForecastRequest) (*pb.ForecastResponse, error) {
	log.Printf("Received GetWeatherForecast request for location: %s, days: %d", in.Location, in.Kunlar)
	// Bu joyda asosiy loyihaga ma'lumotlarni olish va javob qaytarish kerak
	forecasts := []*pb.Forecast{
		{Date: "2024-06-28", Weather: "Sunny"},
		{Date: "2024-06-29", Weather: "Cloudy"},
		{Date: "2024-06-30", Weather: "Rainy"},
	}
	return &pb.ForecastResponse{
		Location: in.Location,
		Forecast: forecasts,
	}, nil
}

// ReportWeatherCondition metodini serverga qo'shish
func (s *server) ReportWeatherCondition(ctx context.Context, in *pb.ConditionReport) (*pb.ConditionResponse, error) {
	log.Printf("Received ReportWeatherCondition request for location: %s, weather: %s, temperature: %d", in.Location, in.Weather, in.Temperature)
	// Bu joyda asosiy loyihaga ma'lumotlarni qaytarish kerak
	return &pb.ConditionResponse{
		Success: true,
		Message: "Weather condition reported successfully.",
	}, nil
}

// mustEmbedUnimplementedWeatherServiceServer metodini qo'shish
func (s *server) mustEmbedUnimplementedWeatherServiceServer() {}

func main() {
	listener, err := net.Listen("tcp", ":50051") // gRPC serverni eshitish uchun port ochish
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()                         // gRPC server yaratish
	pb.RegisterWeatherServiceServer(s, &server{}) // gRPC serverini serverimiz bilan birlashtirish
	log.Println("Starting gRPC server on port :50051")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
