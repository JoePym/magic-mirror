package main

import (
	"fmt"
	"net/http"
  "github.com/joho/godotenv"
  "os"
  "strconv"
)

func currentLatandLong() (lat, long float64) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	lat, _ = strconv.ParseFloat(os.Getenv("LATITUDE"), 64)
	long, _ = strconv.ParseFloat(os.Getenv("LONGITUDE"), 64)

	return lat, long
}

func handler(writer http.ResponseWriter, r *http.Request) {
	report := new(WeatherReport)
	report.pullFrom("https://weather.gc.ca/rss/city/on-59_e.xml")
  lat, long := currentLatandLong()
	trafficIncidents := FetchCloseTrafficIncidents("http://www.mtocdn.ca/kml/events-en.kml", 10000,lat,long)
	for _, incident := range trafficIncidents {
		fmt.Fprintf(writer, "<h1>%s</h1>", incident)
	}
	fmt.Fprintf(writer, "<h1>%s</h1>", report.Warnings)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
