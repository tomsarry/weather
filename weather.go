package main

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
	"github.com/tomsarry/weather/utils"
)

func init() {
	godotenv.Load(".env")
}

func main() {

	pref := utils.GetPref()

	flags := utils.GetFlag()
	// fmt.Println(flags)

	city := utils.GetCity()
	// fmt.Println(theCity)

	if city == "" && pref == "" {
		fmt.Println("[weather] Error: Expected city but received none. (No prefered city saved)")
		return
	}

	newPref := strings.ContainsAny(flags, "f")
	if newPref {
		err := utils.AddPref(city)
		if err == -1 {
			return
		}
	}

	if strings.ContainsAny(flags, "c") {
		// updated the new saved city
		if newPref {
			utils.DispPref(city)
		} else if city == "" {
			fmt.Println("[weather] Error: Expected city but received none.")
			return
		} else {
			utils.DispRes(city)
		}
	}

	if flags == "" {
		utils.DispPref(pref)
	}
}
