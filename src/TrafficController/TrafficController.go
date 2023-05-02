package TrafficController

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type TrafficController struct {
	SourceUpdated string `json:"sourceUpdated"`
	Results       []struct {
		Location struct {
			Description string  `json:"description"`
			Length      float64 `json:"length"`
		} `json:"location"`
		CurrentFlow struct {
			JamFactor float64 `json:"jamFactor"`
		} `json:"currentFlow"`
	} `json:"results"`
}

func (t TrafficController) GetTrafficData(lat float32, lon float32, rad int, apiKey string) (TrafficController, error) {

	var url string = fmt.Sprintf("https://data.traffic.hereapi.com/v7/flow?in=circle:%f,%f;r=%d&locationReferencing=none&apiKey=%s", lat, lon, rad, apiKey)

	var jsonStr = []byte(``)

	var client *http.Client = &http.Client{}

	req, err_0 := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))

	if err_0 != nil {
		log.Fatalf("Could not create a request to %s, Error: %s", url, err_0.Error())
		return TrafficController{}, err_0
	}

	resp, err_1 := client.Do(req)

	if err_1 != nil {
		log.Fatalf("Could not make a request to %s, Error: %s", url, err_1.Error())
		return TrafficController{}, err_1
	}

	defer resp.Body.Close()

	body, err_2 := ioutil.ReadAll(resp.Body)
	if err_2 != nil {
		log.Fatalf("Could not read the response body, Error: %s", err_2.Error())
		return TrafficController{}, err_2
	}

	var response TrafficController = TrafficController{}
	err_3 := json.Unmarshal(body, &response)
	if err_3 != nil {
		log.Fatalf("Could not unmarshal %s", err_3.Error())
		return TrafficController{}, err_3
	}

	return response, nil
}
