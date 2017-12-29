package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
  "encoding/xml"
  "math"
  "github.com/joho/godotenv"
  "os"
  "strconv"
  "strings"
)

type Placemark struct {
	Name		string	`xml:"name"`
	Description	string	`xml:"description"`
	Point		string	`xml:"Point>coordinates"`
}

type Folder struct {
  Name string `xml:"name"`
  Description string `xml:"description"`
  Placemark 	[]Placemark
}

type Document struct {
  Name string `xml:"name"`
  Description string `xml:"description"`
	Folder 	[]Folder
}

type KML struct {
	XMLName		xml.Name	`xml:"kml"`
	Namespace	string		`xml:"xmlns,attr"`
	Document 	Document
}

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

func distanceBetweenTwoPoints(lat1, lon1, lat2, lon2 float64) float64 {
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100

	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}

func currentLatandLong() (lat, long float64) {
  err := godotenv.Load()
  if err != nil {
    fmt.Println("Error loading .env file")
  }
  lat, _ = strconv.ParseFloat(os.Getenv("LATITUDE"), 64)
  long, _ = strconv.ParseFloat(os.Getenv("LONGITUDE"), 64)

  return lat, long
}

func (placemark *Placemark) getLatAndLong() (lat,long float64) {
  geo_data := strings.Split(placemark.Point, ",")

  long, _ = strconv.ParseFloat(geo_data[0], 64)
  lat, _ = strconv.ParseFloat(geo_data[1], 64)

  return lat,long
}

func FetchCloseTrafficIncidents(trafficSource string, radius float64) (closeIncidents []string){
  resp, _ := http.Get(trafficSource)
  body, _ := ioutil.ReadAll(resp.Body)
  defer resp.Body.Close()

  var k KML
  err := xml.Unmarshal(body, &k)

  if err != nil {
    fmt.Println(err)
  }
  var lat, long float64

  currentLat, currentLong := currentLatandLong()

  for _, folder := range k.Document.Folder {
    for _, incident := range folder.Placemark {
      lat, long = incident.getLatAndLong()
      if distanceBetweenTwoPoints(lat, long, currentLat, currentLong) <= radius {
        closeIncidents = append(closeIncidents, incident.Name)
      }
    }
  }
  return closeIncidents
}
