package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	f "github.com/kenhkelly/GoWeather/src/forecast"
	i "github.com/kenhkelly/GoWeather/src/ipinfo"
	"github.com/kenhkelly/GoWeather/src/objects"
)

var (
	version  string = "3.1"
	forecast bool
)

func help() {
	fmt.Printf("Usage: goweather [flags]\n")
	flag.PrintDefaults()
}

func exitHelp() {
	help()
	os.Exit(3)
}

func init() {

	flag.Usage = help

	forecastPtr := flag.Bool("forecast", false, "Show 8 day forecast")
	flag.Parse()

	forecast = *forecastPtr

}

func main() {
	fmt.Printf("\nGoWeather %s - @kenhkelly\n", version)

	key := objects.GetApiKey()

	lat, lng, err := i.GetLocation()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fc, err := f.GetForecast(key, lat, lng)
	if err != nil {
		fmt.Println("An error occured:", err)
		return
	}
	cur := fc.Currently
	dail := fc.Daily

	if !forecast {
		cur_time := time.Unix(cur.Time, 0).Format(time.RFC822)

		cur_str := `
      Current Weather: %s

        Summary     %s
        Temperature %f
        Humidity    %f
        WindSpeed   %f
        WindBearing %f
      `
		fmt.Printf(cur_str, cur_time, cur.Summary, cur.Temperature, cur.Humidity, cur.WindSpeed, cur.WindBearing)
	} else {
		var dail_str string

		for _, v := range dail.Data {
			d_time := time.Unix(v.Time, 0).Format(time.RFC822)
			dail_for_str := `
      Weather for %s

        Summary         %s
        Temperature Min %f
        Temperature Max %f
        Humidity        %f
        WindSpeed       %f
        WindBearing     %f
      `

			dail_str += fmt.Sprintf(dail_for_str, d_time, v.Summary, v.TemperatureMin, v.TemperatureMax, v.Humidity, v.WindSpeed, v.WindBearing)
		}

		fmt.Println(dail_str)
	}

	fmt.Println("")
}
