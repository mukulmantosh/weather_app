package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	url2 "net/url"
	"os"
	"strings"
)

type Weather struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzId           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMm   float64 `json:"precip_mm"`
		PrecipIn   float64 `json:"precip_in"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		VisKm      float64 `json:"vis_km"`
		VisMiles   float64 `json:"vis_miles"`
		Uv         float64 `json:"uv"`
		GustMph    float64 `json:"gust_mph"`
		GustKph    float64 `json:"gust_kph"`
	} `json:"current"`
}

func main() {
	ApiKey := os.Getenv("API_KEY")
	if ApiKey == "" {
		fmt.Println("=*=*=* MAKE SURE TO SET THE API KEY IN ENVIRONMENT VARIABLES =*=*=*")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter name of the city: ")
	city, _ := reader.ReadString('\n')
	city = strings.Replace(city, "\n", "", -1)
	city = url2.QueryEscape(city)

	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s",
		ApiKey, city)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: Unable to fetch weather information ")
		return
	}
	defer response.Body.Close()
	var weather Weather
	err = json.NewDecoder(response.Body).Decode(&weather)
	if err != nil {
		fmt.Println("Error: Unable to decode weather information")

		return
	}

	fmt.Printf("The present temperature is %.2f degree celsius\n", weather.Current.FeelslikeC)

}
