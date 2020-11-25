package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/tomsarry/weather/models"
)

// GetURL creates the URL for the API request
func getURL(city string) string {
	// Madlad not using a .env to hide his API key :O
	return "http://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=0279d3934b72cf457f1a020a85d40371&units=metric"
}

// GetPref returns the prefered city, if it exists
func GetPref() string {
	pref := models.Preferences{}
	jsonFile, err := os.Open("pref.json")

	if err != nil {
		fmt.Println("[weather] Did not find `pref.json`, creating it now.")
	}

	defer jsonFile.Close()

	b, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(b, &pref)

	return pref.City
}

// GetFlags returns a string containing all the flags of the query
func GetFlag() string {
	args := os.Args[1:]
	res := ""

	for i := 0; i < len(args); i++ {
		if args[i][0] == '-' {
			res += args[i][1:]
		}
	}
	return res
}

// GetCity returns the city queried
func GetCity() string {
	args := os.Args[1:]

	if len(args) < 1 {
		return ""
	}

	city := args[len(args)-1]
	if city[0] == '-' {
		return ""
	}
	return city
}

// AddPref saves the new prefered city
func AddPref(city string) int {
	if city == "" {
		fmt.Println("[weather] Error: Please enter a valid city name to save.")
		return -1
	}
	toWrite := models.Preferences{
		City: city,
	}

	file, err := json.MarshalIndent(toWrite, "", " ")

	if err != nil {
		panic(err.Error())
	}

	err = ioutil.WriteFile("pref.json", file, 0644)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("[weather] New prefered city is %s.\n", city)
	return 0
}

func DispPref(pref string, newPref bool) {
	if !newPref {
		fmt.Printf("[weather] Using prefered city: %s.\n", pref)
	}

	makeRequest(pref)
}

// DispRes displays the result when the user does no have a prefered place
func DispRes(city string) {
	makeRequest(city)
}

func makeRequest(city string) {
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
