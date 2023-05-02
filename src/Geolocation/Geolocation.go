package Geolocation

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"main/src/Config"
	"net/http"
)

type Geolocation struct {
	Latitude  float32
	Longitude float32
}

// getLatLonJSON sends a POST request to the Google Geocoding API to obtain the
// latitude and longitude as JSON.
func (g *Geolocation) GetLatLonJSON() string {

	config := Config.Config{}.LoadConfig()

	// Define the URL for the Google Geocoding API.
	url := "https://www.googleapis.com/geolocation/v1/geolocate?key=" + config.GoogleMapAPIKey

	// Define an empty JSON string and create a new HTTP request with it.
	// We use jsonStr as a request body in bytes.NewBuffer. bytes.NewBuffer creates an io.Reader from byte slice
	var jsonStr = []byte(``)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	// Set the necessary headers for the request.
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Vary", "Origin")
	req.Header.Set("Vary", "X-Origin")
	req.Header.Set("Vary", "Referer")
	req.Header.Set("Content-Encoding", "gzip")
	req.Header.Set("Date", "Wed, 01 Jul 2020 14:21:50 GMT")
	req.Header.Set("Server", "scaffolding on HTTPServer2")
	req.Header.Set("Cache-Control", "private")
	req.Header.Set("X-XSS-Protection", "0")
	req.Header.Set("X-Frame-Options", "SAMEORIGIN")
	req.Header.Set("X-Content-Type-Options", "nosniff")
	req.Header.Set("Alt-Svc", "h3-27="+":443"+"; ma=2592000,h3-25="+":443"+"; ma=2592000,h3-T050="+":443"+"; ma=2592000,h3-Q050="+":443"+"; ma=2592000,h3-Q046="+":443"+"; ma=2592000,h3-Q043="+":443"+"; ma=2592000,quic="+":443"+"; ma=2592000; v="+"46,43"+"")
	req.Header.Set("Transfer-Encoding", "chunked")

	// Check if there was an error creating the request and log it if so.
	if err != nil {
		log.Printf("error creating HTTP request: %v", err)
	}

	// Send the HTTP request and check for errors.
	var client *http.Client = &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("error sending HTTP request: %v", err)
		return ""
	}
	defer resp.Body.Close()

	// Read the response body and convert it to a string.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return ""
	}

	return string(body)
}

// UserGeolocation Represents user geolocation object with Latitude and Longitute attributes. Is used for parsing lat and lng.
type UserGeolocation struct {
	Accuracy int64 `json:"accuracy"`
	Location struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"location"`
}

// For parsing the latitude and longitude from strutured JSON
func (g *Geolocation) GetLatLon(locJSON string) (float64, float64) {

	var loc UserGeolocation
	err := json.Unmarshal([]byte(locJSON), &loc)

	if err != nil {
		log.Printf("Error in getting lat and lng: %v", err)
	}

	return loc.Location.Lat, loc.Location.Lng
}

func (g *Geolocation) GetCurrentLocation() (float32, float32) {
	var LAT, LNG = g.GetLatLon(g.GetLatLonJSON())
	return float32(LAT), float32(LNG)
}

func (g *Geolocation) Set(Latitude float32, Longitude float32) {
	g.Latitude = Latitude
	g.Longitude = Longitude
}

func (g *Geolocation) Get() (float32, float32) {
	return g.Latitude, g.Longitude
}
