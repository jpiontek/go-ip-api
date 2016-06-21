// Package goip provides a thin wrapper around the ip-api.com API to retrieve
// geolocation data for a specific IP address
package goip

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type Client interface {
	GetLocation() (*Location, error)
	GetLocationForIp(ip string) (*Location, error)
}

// Primary URI
const STANDARD_URI = "http://ip-api.com/json/"

// Pro URI
const PRO_URI = "http://pro.ip-api.com/json/"

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

// ProClient is a commercial client for retrieving location data.
type ProClient struct {
	URI        string
	HttpClient *http.Client
	ApiKey     string
}

// GetLocation retrieves the current client's public IP address location
// information
func (g *ProClient) GetLocation() (*Location, error) {
	uri := buildProUri("", g.ApiKey)
	return getLocation(uri, g.HttpClient)
}

// GetLocationForIp retrieves the supplied IP address's location information
func (g *ProClient) GetLocationForIp(ip string) (*Location, error) {
	uri := buildProUri(ip, g.ApiKey)
	return getLocation(uri, g.HttpClient)
}

// StandardClient is a free client for retreiving location data with a
// 150 request per minute limit.
type StandardClient struct {
	URI        string
	HttpClient *http.Client
}

// GetLocation retrieves the current client's public IP address location
// information
func (g *StandardClient) GetLocation() (*Location, error) {
	return getLocation(g.URI, g.HttpClient)
}

// GetLocationForIp retrieves the supplied IP address's location information
func (g *StandardClient) GetLocationForIp(ip string) (*Location, error) {
	uri := buildStandardUri(ip)
	return getLocation(uri, g.HttpClient)
}

func getLocation(uri string, httpClient *http.Client) (*Location, error) {
	resp, err := httpClient.Get(uri)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 403 {
		if strings.Contains(uri, "?key=") {
			// Currently using a pro client, 403 is usually invalid API key
			return nil, errors.New("Invalid API key")
		}
		return nil, errors.New("Exceeded maximum number of API calls")

	}

	var l Location

	err = json.NewDecoder(resp.Body).Decode(&l)
	if err != nil {
		return nil, err
	}

	if l.Status != "success" {
		err := errors.New("Failed to find location data")
		return nil, err
	}
	return &l, nil
}

func buildStandardUri(ip string) string {
	return STANDARD_URI + ip
}

func buildProUri(ip string, apiKey string) string {
	if ip == "" {
		return PRO_URI + "?key=" + apiKey
	}
	return PRO_URI + ip + "?key=" + apiKey
}

// NewClient returns a new client
func NewClient() Client {
	return &StandardClient{URI: STANDARD_URI, HttpClient: &http.Client{}}
}

func NewClientWithApiKey(apiKey string) Client {
	return &ProClient{URI: PRO_URI, HttpClient: &http.Client{}, ApiKey: apiKey}
}
