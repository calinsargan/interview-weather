package main

import "fmt"

type WeatherAPI interface {
	GetWeather(city string) (*WeatherResponse, error)
}

type WeatherResponse struct {
	Temperature float64
	Humidity    int
	Description string
}

// Task: Scrie funcția GetCityTemperature care folosește WeatherAPI
// și returnează doar temperatura.

func main() {
	fmt.Println("Rulați testele cu: go test")
}
