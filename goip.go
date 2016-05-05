package goip

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Location struct {
	As          string `json:"as"`
	City        string `json:"city"`
	Country     string `json:"country"`
	CountryCode string `json:"countryCode"`
	Isp         string `json:"isp"`
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
	Org         string `json:"org"`
	Query       string `json:"query"`
	Region      string `json:"region"`
	RegionName  string `json:"regionName"`
	Status      string `json:"status"`
	Timezone    string `json:"timezone"`
	Zip         string `json:"zip"`
}

func GetLocation() (*Location, error) {
	resp, err := http.Get("http://ip-api.com/json")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var l Location
	json.Unmarshal([]byte(contents), &l)
	return &l, nil
}
