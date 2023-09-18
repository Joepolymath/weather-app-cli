package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type WeatherDetail struct {
	Dt		int		`json:"dt"`
	Main	struct {
		Temp	float64	`json:"temp"`
		FeelsLike	float64	`json:"feels_like"`
		TempMin		float64	`json:"temp_min"`
		TempMax		float64	`json:"temp_max"`
		Pressure		int		`json:"pressure"`
		SeaLevel		int		`json:"sea_level"`
		GrndLevel	int		`json:"grnd_level"`
		Humidity		int		`json:"humidity"`
		TempKf		float64	`json:"temp_kf"`
	}	`json:"main"`
	Weather	[]struct {
		ID		uint64		`json:"id"`
		Main	string		`json:"main"`
		Description	string `json:"description"`
		Icon	string		`json:"icon"`
	}		`json:"weather"`
	Clouds	struct {
		All	int	`json:"all"`
	}	`json:"clouds"`
	Winds	struct {
		Speed	float32	`json:"speed"`
		Deg 	int		`json:"deg"`
		Gust	float64	`json:"gust"`
	}	`json:"winds"`
	Visibility	int	`json:"visibility"`
	Pop	int	`json:"pop"`
	Sys	struct {
		Pod	string	`json:"pod"`
	}
	DtTxt	string	`json:"dt_txt"`
}

type Weather struct {
	List	[]Weather
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	api_key := os.Getenv("API_KEY")
	url := "http://api.openweathermap.org/data/2.5/forecast?id=524901&appid=" + api_key
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic(res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}