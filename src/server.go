package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, r *http.Request) {
  report := new(WeatherReport)
  report.pullFrom("https://weather.gc.ca/rss/city/on-59_e.xml")
	fmt.Fprintf(writer, "<h1>%s</h1>", report.Warnings)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
