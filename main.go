package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/joho/godotenv"
)

type WeatherResponse struct {
    Location struct {
        Name    string `json:"name"`
        Country string `json:"country"`
    } `json:"location"`
    Current struct {
        TempC    float64 `json:"temp_c"`
        Condition struct {
            Text string `json:"text"`
        } `json:"condition"`
        WindKph  float64 `json:"wind_kph"`
        Humidity int     `json:"humidity"`
    } `json:"current"`
}

func main() {
    // Load the .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    apiKey := os.Getenv("API_KEY")
    if apiKey == "" {
        log.Fatal("API_KEY not set in .env file")
    }

    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <city>")
        return
    }

    city := os.Args[1]
    url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, city)

    resp, err := http.Get(url)
    if err != nil {
        log.Fatal("Error fetching weather data:", err)
    }
    defer resp.Body.Close()

    var weatherData WeatherResponse
    if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
        log.Fatal("Error decoding response:", err)
    }

    // Display weather information
    fmt.Printf("Weather in %s, %s:\n", weatherData.Location.Name, weatherData.Location.Country)
    fmt.Printf("Temperature: %.1f°C\n", weatherData.Current.TempC)
    fmt.Printf("Condition: %s\n", weatherData.Current.Condition.Text)
    fmt.Printf("Wind: %.1f kph\n", weatherData.Current.WindKph)
    fmt.Printf("Humidity: %d%%\n", weatherData.Current.Humidity)
}
