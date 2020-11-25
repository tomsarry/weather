package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/tomsarry/weather/models"
	"github.com/joho/godotenv"
)


func init() {
	godotenv.Load(".env")
}

func getURL(city string) string {
	// Madlad not using a .env to hide his API key :O
	return "http://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=0279d3934b72cf457f1a020a85d40371&units=metric"
}

func main() {

	pref := models.Preferences{}
	jsonFile, err := os.Open("pref.json")
		
	if err != nil {
		fmt.Println("[weather] Did not find `pref.json`, creating it now.")
	}

	defer jsonFile.Close()

	b, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(b, &pref)

	args := os.Args[1:]
	city := ""

	if len(args) == 0 {

		if pref.City == "" {
			fmt.Println("[weather] Error: Expected arguments but received none. (No prefered city saved)")
			os.Exit(1)
		}	
		fmt.Printf("[weather] Using prefered city: %s.\n", pref.City)	
		city = pref.City		
	} else {
		if args[0] == "-c" {
			if len(args) == 1 {
				fmt.Println("[weather] Error: Valid call is 'weather -c `City`', where `City` is the city you want to know the weather.")
				os.Exit(1)
			}
			city = args[1]
		} else if args[0] == "-f" {
			if len(args) == 1 {
				fmt.Println("[weather] Error: Valid call is 'weather -f `City`', where `City` is the city you want to save as preferred city.")
				os.Exit(1)
			}
			toWrite := models.Preferences{
				City: args[1],
			}
	
			file, err := json.MarshalIndent(toWrite, "", " ")
	
			if err != nil {
				panic(err.Error())
			}
	
			err = ioutil.WriteFile("pref.json", file, 0644)
	
			if err != nil {
				panic(err.Error())
			}
	
			fmt.Printf("[weather] New prefered city is %s.\n", args[1])
			return 
		}
	}

	url := getURL(city)
	resp, err := http.Get(url)

	if err != nil {
		panic("[weather] Error with API call.")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("[weather] Error: Could not read response.")
	}

	res := models.Response{}
	json.Unmarshal([]byte(body), &res)

	if len(res.Weather) == 0 {
		fmt.Printf("[weather] Did not receive a valid response, are you sure about the city name: %s ?\n", city)
		return
	}

	fmt.Printf("[weather] Results for %s, %s:\n", city, res.Misc.Country)
	fmt.Printf("[weather] Temperature: %.1f°C\n", res.Main.Temp)
	fmt.Printf("[weather] Weather: %s - (%s)\n", res.Weather[0].Main, res.Weather[0].Desc)
}
