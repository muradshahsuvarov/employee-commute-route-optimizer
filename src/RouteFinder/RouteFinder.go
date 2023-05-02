package RouteFinder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
)

type RouteResponse struct {
	Response struct {
		Route []struct {
			Mode struct {
				Type           string   `json:"type"`
				TransportModes []string `json:"transportModes"`
				TrafficMode    string   `json:"trafficMode"`
			} `json:"mode"`
			BoatFerry bool `json:"boatFerry"`
			RailFerry bool `json:"railFerry"`
			Waypoint  []struct {
				LinkID         string `json:"linkId"`
				MappedPosition struct {
					Latitude  float64 `json:"latitude"`
					Longitude float64 `json:"longitude"`
				} `json:"mappedPosition"`
				OriginalPosition struct {
					Latitude  float64 `json:"latitude"`
					Longitude float64 `json:"longitude"`
				} `json:"originalPosition"`
				Spot                        float64 `json:"spot"`
				ConfidenceValue             float64 `json:"confidenceValue"`
				Elevation                   float64 `json:"elevation"`
				HeadingDegreeNorthClockwise float64 `json:"headingDegreeNorthClockwise"`
				HeadingMatched              float64 `json:"headingMatched"`
				MatchDistance               float64 `json:"matchDistance"`
				MinError                    float64 `json:"minError"`
				RouteLinkSeqNrMatched       int     `json:"routeLinkSeqNrMatched"`
				SpeedMps                    float64 `json:"speedMps"`
				Timestamp                   int     `json:"timestamp"`
			} `json:"waypoint"`
			Leg []struct {
				Length     int `json:"length"`
				TravelTime int `json:"travelTime"`
				Link       []struct {
					LinkID          string    `json:"linkId"`
					Length          float64   `json:"length"`
					RemainDistance  int       `json:"remainDistance"`
					RemainTime      int       `json:"remainTime"`
					Shape           []float64 `json:"shape"`
					FunctionalClass int       `json:"functionalClass"`
					Confidence      float64   `json:"confidence"`
					SegmentRef      string    `json:"segmentRef"`
				} `json:"link"`
				TrafficTime     int `json:"trafficTime"`
				BaseTime        int `json:"baseTime"`
				RefReplacements struct {
					Zero string `json:"0"`
					One  string `json:"1"`
				} `json:"refReplacements"`
			} `json:"leg"`
			Summary struct {
				TravelTime  int           `json:"travelTime"`
				Distance    int           `json:"distance"`
				BaseTime    int           `json:"baseTime"`
				TrafficTime int           `json:"trafficTime"`
				Flags       []interface{} `json:"flags"`
			} `json:"summary"`
		} `json:"route"`
		Warnings []interface{} `json:"warnings"`
		Language string        `json:"language"`
	} `json:"response"`
}

func (rr *RouteResponse) GetRouteFromAtoB(apiKey string, mode []string, waypoint0 string, waypoint1 string, routeMatch int32) RouteResponse {

	var mode_str string = strings.Join(mode, ";")

	var url string = fmt.Sprintf("https://routematching.hereapi.com/v8/match/routelinks?apiKey=%s&mode=%s&waypoint0=%s&waypoint1=%s&routeMatch=%d", apiKey, mode_str, waypoint0, waypoint1, routeMatch)

	var jsonStr = []byte(``)

	req, err_0 := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))

	if err_0 != nil {
		log.Fatalf("Couldn't make a request. Error: %s", err_0.Error())
		return RouteResponse{}
	}

	var client *http.Client = &http.Client{}

	resp, err_1 := client.Do(req)

	if err_1 != nil {
		log.Fatalf("Couldn't call the request. Error: %s", err_1.Error())
		return RouteResponse{}
	}

	var routeResponse RouteResponse = RouteResponse{}

	resp_data, err_2 := ioutil.ReadAll(resp.Body)

	if err_2 != nil {
		log.Fatalf("Couldn't read the response. Error: %s", err_2.Error())
		return RouteResponse{}
	}

	err_3 := json.Unmarshal(resp_data, &routeResponse)

	if err_2 != nil {
		log.Fatalf("Couldn't unmarshall the response data. Error: %s", err_3.Error())
		return RouteResponse{}
	}

	return routeResponse
}

func (rr *RouteResponse) GetTheShortestLocation(apiKey string, mode []string, waitPoint string, waitPoints []string, routeMatch int) RouteResponse {

	var routeResponses []RouteResponse = make([]RouteResponse, 0)

	// Unbuffered channel
	ch := make(chan RouteResponse)

	for _, value := range waitPoints {
		go func(val string) {
			ch <- rr.GetRouteFromAtoB(apiKey, mode, waitPoint, val, int32(routeMatch))
		}(value)
	}

	// Blocking operation on the unbuffered channel ch
	for i := 0; i < len(waitPoints); i++ {
		routeResponses = append(routeResponses, <-ch)
	}

	// Sorting operation. Sorting is in ascending order
	sort.Slice(routeResponses, func(i, j int) bool {
		return routeResponses[i].Response.Route[0].Summary.TravelTime < routeResponses[j].Response.Route[0].Summary.TravelTime
	})

	return routeResponses[len(routeResponses)-1]
}
