// Package goip provides a thin wrapper around the ip-api.com API to retrieve
// geolocation data for a specific IP address
package goip

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// Primary URI
const IpApiUri = "http://ip-api.com/json/"

// Location contains all the relevant data for an IP
type Location struct {
	As          string  `json:"as"`
	City        string  `json:"city"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Isp         string  `json:"isp"`
	Lat         float32 `json:"lat"`
	Lon         float32 `json:"lon"`
	Org         string  `json:"org"`
	Query       string  `json:"query"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	Status      string  `json:"status"`
	Timezone    string  `json:"timezone"`
	Zip         string  `json:"zip"`
}

// Client is the actor responsible for retrieving location data from the
// remote endpoint
type Client struct {
	URI        string
	HttpClient *http.Client
}

// GetLocation retrieves the current client's public IP address location
// information
func (g *Client) GetLocation() (*Location, error) {
	return getLocation(g.URI, g.HttpClient)
}

// GetLocationForIp retrieves the supplied IP address's location information
func (g *Client) GetLocationForIp(ip string) (*Location, error) {
	uri := g.URI + ip
	return getLocation(uri, g.HttpClient)
}

func getLocation(uri string, httpClient *http.Client) (*Location, error) {
	resp, err := httpClient.Get(uri)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 403 {
		return nil, errors.New("Exceeded maximum number of API calls")
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var l Location
	json.Unmarshal([]byte(contents), &l)

	if l.Status != "success" {
		err := errors.New("Failed to find location data")
		return nil, err
	}
	return &l, nil
}

// Returns a new client
func NewClient() *Client {
	return &Client{URI: IpApiUri, HttpClient: &http.Client{}}
}
