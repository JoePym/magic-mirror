package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, r *http.Request) {
	report := new(WeatherReport)
	report.pullFrom("https://weather.gc.ca/rss/city/on-59_e.xml")
	trafficIncidents := FetchCloseTrafficIncidents("http://www.mtocdn.ca/kml/events-en.kml", 10000)
	for _, incident := range trafficIncidents {
		fmt.Fprintf(writer, "<h1>%s</h1>", incident)
	}
	fmt.Fprintf(writer, "<h1>%s</h1>", report.Warnings)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
