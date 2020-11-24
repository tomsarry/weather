package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
}

func getURL(city string) string {
	return "http://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + os.Getenv("API_KEY")
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

	fmt.Println(string(body))
}
