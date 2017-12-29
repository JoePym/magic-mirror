package main

import (
  "testing"
  "io/ioutil"
  "gopkg.in/jarcoal/httpmock.v1"
)

func TestPullFrom(t *testing.T) {
  httpmock.Activate()
  defer httpmock.DeactivateAndReset()

  response, _ := ioutil.ReadFile("fixtures/sample-weather.xml")

	httpmock.RegisterResponder("GET", "https://example.com/feed.xml",
		httpmock.NewStringResponder(200, string(response)))

  report := new(WeatherReport)
  report.pullFrom("https://example.com/feed.xml")

  if report.CurrentConditions != "Current Conditions: Mostly Cloudy, -14.7°C" {
    t.Error("expected 'Current Conditions: Mostly Cloudy, -14.7°C' got ", report.CurrentConditions)
  }

  if report.Warnings[0] != "No watches or warnings in effect, Richmond Hill" {
    t.Error("expected 'No watches or warnings in effect, Richmond Hill' got ", report.Warnings)
  }

}
