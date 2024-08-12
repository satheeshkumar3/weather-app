package main

import (
	"context"
	"fmt" // Use fmt for formatted printing
	"log"
	"time"

	"github.com/satheeshkumar3/weather-app/generated" // Import generated code

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := generated.NewWeatherServiceClient(conn) // Use generated client

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &generated.WeatherRequest{City: "San Francisco"} // Use generated message type
	res, err := client.GetWeather(ctx, req)
	if err != nil {
		log.Fatalf("failed to get weather: %v", err)
	}

	// Use fmt.Printf for formatted output
	fmt.Printf("Weather in %s: %s, %.2f°C\n", res.GetCity(), res.GetDescription(), res.GetTemperature())
}
