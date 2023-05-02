package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"main/src/Config"
	"main/src/Geolocation"
	"main/src/RouteFinder"
	"main/src/TrafficController"
	"net/http"
	"strconv"
)

var port int64 = 8000

func getRootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Welcome to ECRO!")
	w.WriteHeader(http.StatusOK)
}

func getCurrentLocationHandler(w http.ResponseWriter, r *http.Request) {
	lat_lon_object := (&Geolocation.Geolocation{}).GetLatLonJSON()
	fmt.Fprintf(w, lat_lon_object)
}

func getLocationTrafficHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		log.Fatal("Method call has to be POST")
		return
	}

	var apiKey string = Config.Config{}.LoadConfig().HEREAPIKey[0]

	// Parse the request body to get the lat and lon values
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error reading request body:", err)
	}
	defer r.Body.Close()

	var requestBody struct {
		Lat float32 `json:"lat"`
		Lon float32 `json:"lon"`
		Rad int32   `json:"rad"`
	}

	err = json.Unmarshal(reqBody, &requestBody)
	if err != nil {
		log.Fatal("Error unmarshaling request body:", err)
	}

	// get the traffic data for the lat and lon values
	trafficData, err := TrafficController.TrafficController{}.GetTrafficData(requestBody.Lat, requestBody.Lon, int(requestBody.Rad), apiKey)

	// marshal the output map into a JSON object and send it in the response
	outputJSON, err := json.Marshal(trafficData)

	if err != nil {
		log.Fatal("Couldn't marshal the output into a JSON object:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(outputJSON)
}

func getRouteDataHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		log.Fatal("Method call has to be POST")
		return
	}

	r_body, err_0 := ioutil.ReadAll(r.Body)

	if err_0 != nil {
		log.Fatalf("Coudln't read the request body. Error: %s", err_0.Error())
		return
	}

	type body_struct struct {
		Mode       []string `json:"mode"`
		Waypoint0  string   `json:"waypoint0"`
		Waypoint1  string   `json:"waypoint1"`
		RouteMatch int32    `json:"routematch"`
	}

	var bs body_struct = body_struct{}

	err_1 := json.Unmarshal(r_body, &bs)

	if err_1 != nil {
		log.Fatalf("Coudln't unmarshall the request body. Error: %s", err_1.Error())
		return
	}

	var apiKey string = Config.Config{}.LoadConfig().HEREAPIKey[1]

	var routeResponse RouteFinder.RouteResponse = RouteFinder.RouteResponse{}

	routeResponse = routeResponse.GetRouteFromAtoB(apiKey, bs.Mode, bs.Waypoint0, bs.Waypoint1, bs.RouteMatch)

	routeResponseJSON, err_2 := json.Marshal(routeResponse)

	if err_2 != nil {
		log.Fatalf("Coudln't marshal the route response. Error: %s", err_2.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(routeResponseJSON)

}

func getTheShortestLocationHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		log.Fatal("Method call has to be POST")
		return
	}

	r_body, err_0 := ioutil.ReadAll(r.Body)

	if err_0 != nil {
		log.Fatalf("Coudln't read the request body. Error: %s", err_0.Error())
		return
	}

	type body_struct struct {
		Mode       []string `json:"mode"`
		Waypoint   string   `json:"waypoint"`
		Waypoints  []string `json:"waypoints"`
		RouteMatch int32    `json:"routematch"`
	}

	var bs body_struct = body_struct{}

	err_1 := json.Unmarshal(r_body, &bs)

	if err_1 != nil {
		log.Fatalf("Coudln't unmarshall the request body. Error: %s", err_1.Error())
		return
	}

	var apiKey string = Config.Config{}.LoadConfig().HEREAPIKey[1]

	var routeResponse RouteFinder.RouteResponse = RouteFinder.RouteResponse{}

	routeResponse = routeResponse.GetTheShortestLocation(apiKey, bs.Mode, bs.Waypoint, bs.Waypoints, int(bs.RouteMatch))

	routeResponseJSON, err_2 := json.Marshal(routeResponse)

	if err_2 != nil {
		log.Fatalf("Coudln't marshal the route response. Error: %s", err_2.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(routeResponseJSON)

}

func main() {

	http.HandleFunc("/", getRootHandler)
	http.HandleFunc("/getCurrentLocation", getCurrentLocationHandler)
	http.HandleFunc("/getLocationTraffic", getLocationTrafficHandler)
	http.HandleFunc("/getRouteData", getRouteDataHandler)
	http.HandleFunc("/getTheShortestLocationHandler", getTheShortestLocationHandler)
	fmt.Println("Listening on port", port)
	var s_port string = ":" + strconv.FormatInt(port, 10)
	if err := http.ListenAndServe(s_port, nil); err != nil {
		log.Fatalf("Failed to listed at %v", err)
	}

}
