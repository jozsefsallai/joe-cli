package commands

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jozsefsallai/joe-cli/config"
	forecast "github.com/mlbright/darksky/v2"
	"github.com/urfave/cli"
)

func temperatureAt(temperature float64, t int64) string {
	timestamp := time.Unix(t, 0)
	hr, min, _ := timestamp.Clock()
	return fmt.Sprintf("%.2f°C at %d:%02d", temperature, hr, min)
}

func getWeatherEmoji(icon string) string {
	switch icon {
	case "rain", "sleet":
		return "🌧️"
	case "snow":
		return "❄️"
	case "wind":
		return "💨"
	case "fog":
		return "🌫️"
	case "cloudy", "partly-cloudy-day":
		return "⛅"
	case "partly-cloudy-night":
		return "🌙"
	case "thunderstorm":
		return "⛈️"
	default:
		return "☁️"
	}
}

func getWeather(time string, ctx *cli.Context) {
	conf := config.GetConfig()

	if ctx.NArg() != 0 && ctx.NArg() != 2 {
		fmt.Println("If you don't want to use the default coordinates, make sure to provide both of them!")
		os.Exit(1)
	}

	key := conf.Weather.Key

	lat := conf.Weather.Defaults.Latitude
	long := conf.Weather.Defaults.Longitude

	var err error

	if ctx.NArg() == 2 {
		lat, err = strconv.ParseFloat(ctx.Args().Get(0), 64)
		if err != nil {
			panic("The latitude must be a float.")
		}

		long, err = strconv.ParseFloat(ctx.Args().Get(1), 64)
		if err != nil {
			panic("The longitude must be a float.")
		}
	}

	f, err := forecast.Get(
		key,
		fmt.Sprintf("%f", lat),
		fmt.Sprintf("%f", long),
		time,
		forecast.CA,
		forecast.English,
	)

	if err != nil {
		panic(err)
	}

	currentDay := f.Daily.Data[0]

	emoji := getWeatherEmoji(f.Hourly.Icon)

	fmt.Println("Weather forecast for timezone", f.Timezone+":", f.Currently.Summary)
	fmt.Println(emoji, f.Hourly.Summary)
	fmt.Println(
		"Temperature:",
		fmt.Sprintf("%.2f°C", f.Currently.Temperature),
		"(max:",
		fmt.Sprintf("%s, min:", temperatureAt(currentDay.TemperatureMax, currentDay.TemperatureMaxTime)),
		temperatureAt(currentDay.TemperatureMin, currentDay.TemperatureMinTime)+")",
	)
	fmt.Println("Humidity:", f.Currently.Humidity)
}

func weatherCommandNow(ctx *cli.Context) {
	getWeather("now", ctx)
}

func weatherCommandTomorrow(ctx *cli.Context) {
	tomorrow := time.Now().AddDate(0, 0, 1).Unix()
	getWeather(strconv.FormatInt(tomorrow, 10), ctx)
}

// WeatherCommand gets the current weather information
var WeatherCommand = cli.Command{
	Name:  "weather",
	Usage: "Get weather information.",
	Subcommands: []cli.Command{
		{
			Name:    "now",
			Usage:   "Get the weather information at the current time.",
			Aliases: []string{"today"},
			Action:  weatherCommandNow,
		},
		{
			Name:   "tomorrow",
			Usage:  "Get the weather information for tomorrow.",
			Action: weatherCommandTomorrow,
		},
	},
}
