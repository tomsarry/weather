package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type weather struct {
	ID   int    `json:"id"`
	Main string `json:"main"`
	Desc string `json:"description"`
	Icon string `json:"icon"`
}

type mainResp struct {
	Temp     float32 `json:"temp"`
	FLike    float32 `json:"feels_like"`
	TMin     float32 `json:"temp_min"`
	TMax     float32 `json:"temp_max"`
	Pressure float32 `json:"pressure"`
}

type misc struct {
	Country string `json:"country"`
}

type response struct {
	Weather []weather `json:"weather"`
	Main    mainResp  `json:"main"`
	Misc    misc      `json:"sys"`
}

func init() {
	godotenv.Load(".env")
}

func getURL(city string) string {
	return "http://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + os.Getenv("API_KEY") + "&units=metric"
}

func main() {

	args := os.Args[1:]
	city := ""

	if len(args) == 0 {
		fmt.Println("Error: Expected arguments but received none.")
		// panic()
		os.Exit(1)
	}

	if args[0] == "-c" {
		if len(args) == 1 {
			fmt.Println("Error: Valid call is 'weather -c `City`', where `City` is the city you want to know the weather.")
			os.Exit(1)
		}
		city = args[1]
	}

	url := getURL(city)
	resp, err := http.Get(url)

	if err != nil {
		panic("Error with API call.")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("Error : Could not read response.")
	}

	res := response{}
	json.Unmarshal([]byte(body), &res)

	fmt.Printf("Results for %s, %s:\n", city, res.Misc.Country)
	fmt.Printf("Temperature: %.1f\n", res.Main.Temp)
	fmt.Printf("Weather: %s - (%s)\n", res.Weather[0].Main, res.Weather[0].Desc)
}
