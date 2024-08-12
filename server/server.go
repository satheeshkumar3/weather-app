package main

import (
	"context"
	"log"
	"net"

	"github.com/satheeshkumar3/weather-app/generated"

	"google.golang.org/grpc"
)

type server struct {
	generated.UnimplementedWeatherServiceServer
}

func (s *server) GetWeather(ctx context.Context, req *generated.WeatherRequest) (*generated.WeatherResponse, error) {
	resp := &generated.WeatherResponse{
		City:        req.GetCity(),
		Description: "Sunny",
		Temperature: 25.3,
	}
	return resp, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	generated.RegisterWeatherServiceServer(s, &server{})

	log.Println("gRPC server is running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
