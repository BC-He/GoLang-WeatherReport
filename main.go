package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const apiKey = "123abc" // Replace this with your OpenWeatherMap API key

type WeatherResponse struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func main() {
	// OpenWeatherMap API endpoint for Taipei
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=Taipei&appid=%s&units=metric", apiKey)

	// Make HTTP request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error fetching weather data: ", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body: ", err)
	}

	// Parse JSON response
	var weatherData WeatherResponse
	if err := json.Unmarshal(body, &weatherData); err != nil {
		log.Fatal("Error parsing JSON: ", err)
	}

	// Display the weather data
	fmt.Printf("Weather in Taipei:\n")
	fmt.Printf("Temperature: %.2f°C\n", weatherData.Main.Temp)
	fmt.Printf("Humidity: %d%%\n", weatherData.Main.Humidity)
	fmt.Printf("Description: %s\n", weatherData.Weather[0].Description)
}
