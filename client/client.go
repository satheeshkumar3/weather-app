package main

import (
	"context"
	"log"
	"time"

	"github.com/satheeshkumar3/weather-app/generated"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := generated.NewWeatherServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &generated.WeatherRequest{City: "San Francisco"}
	res, err := client.GetWeather(ctx, req)
	if err != nil {
		log.Fatalf("failed to get weather: %v", err)
	}

	log.Printf("Weather in %s: %s, %.2f°C", res.GetCity(), res.GetDescription(), res.GetTemperature())
}
