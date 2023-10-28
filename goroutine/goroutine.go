package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

const apiKey = "1b43c3a98f656f8bcf3e10e6512106a9"

type WeatherData struct {
	City string
	Data interface{}
}

func fetchWeather(city string, ch chan<- WeatherData) {
	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
	res, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error while fetching weather for %s: %s\n", city, err)
		ch <- WeatherData{City: city, Data: nil}
		return
	}

	defer res.Body.Close()

	decodedErr := json.NewDecoder(res.Body).Decode(&data)
	if decodedErr != nil {
		fmt.Printf("Error while decoding weather for %s: %s\n", city, err)
		ch <- WeatherData{City: city, Data: nil}
		return
	}

	ch <- WeatherData{City: city, Data: data}
}

func main() {
	cities := []string{"London", "Paris", "Tokyo", "New York"}
	ch := make(chan WeatherData)

	var wg sync.WaitGroup
	wg.Add(len(cities))

	go func() {
		wg.Wait()
		close(ch)
	}()

	for _, city := range cities {
		go func(city string) {
			defer wg.Done()
			fetchWeather(city, ch)
		}(city)
	}

	for data := range ch {
		fmt.Printf("Weather in %s: %v\n", data.City, data.Data)
	}
}
