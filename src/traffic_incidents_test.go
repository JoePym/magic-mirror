package main

import (
	"gopkg.in/jarcoal/httpmock.v1"
	"io/ioutil"
	"testing"
)

func TestFetchCloseTrafficIncidents(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

  requestURI := "http://example.com/feed.kml"

	response, _ := ioutil.ReadFile("fixtures/traffic-incidents.kml")
	httpmock.RegisterResponder("GET", requestURI,
		httpmock.NewStringResponder(200, string(response)))

  incidents := FetchCloseTrafficIncidents(requestURI, 10000, 43.691920, -79.595811)

	if len(incidents) != 1 {
		t.Error("expected 1 incidents got ", len(incidents))
	}
}
