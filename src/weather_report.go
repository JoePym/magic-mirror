package main

import (
	"github.com/mmcdole/gofeed"
)

type WeatherReport struct {
  Forecasts []string
  CurrentConditions string
  Warnings []string
}

func (report *WeatherReport) pullFrom (source string) {
  fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(source)
  for _, event := range feed.Items {
		switch event.Categories[0] {
		case "Warnings and Watches":
      report.Warnings = append(report.Warnings, event.Title)
		case "Current Conditions":
      report.CurrentConditions = event.Title
    case "Weather Forecasts":
      report.Forecasts = append(report.Forecasts, event.Title)
		}
	}
}
