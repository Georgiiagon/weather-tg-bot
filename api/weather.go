package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Weather struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

func GetByCoordinates(lat string, lon string) (weather Weather) {
	res, err := http.Get("http://api.openweathermap.org/data/2.5/weather?lat=" + lat + "&lon=" + lon + "&appid=" + os.Getenv("OPEN_WEATHER_KEY") + "&units=metric")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	_ = json.NewDecoder(res.Body).Decode(&weather)

	return weather
}

func GetByCity(city string) (weather Weather) {
	fmt.Println(city)
	res, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + os.Getenv("OPEN_WEATHER_KEY") + "&units=metric")

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	_ = json.NewDecoder(res.Body).Decode(&weather)

	return weather
}
