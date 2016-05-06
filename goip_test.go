package goip

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func getMockSuccessResponse() string {
	l := Location{
		As:          "500 North Pole",
		City:        "Timbuktu",
		Country:     "Fairy Tale Land",
		CountryCode: "FT",
		Isp:         "Acme ISP Service",
		Lat:         55.555,
		Lon:         33.333,
		Org:         "Who Knows?",
		Query:       "123.123.123.4",
		Region:      "CM",
		RegionName:  "Chocolate Mountain",
		Status:      "success",
		Timezone:    "America/Chicago",
		Zip:         "12312",
	}
	result, _ := json.Marshal(l)
	return string(result)
}

func getMockFailureResponse() string {
	l := Location{
		As:          "",
		City:        "",
		Country:     "",
		CountryCode: "",
		Isp:         "",
		Lat:         0.0,
		Lon:         0.0,
		Org:         "",
		Query:       "",
		Region:      "",
		RegionName:  "",
		Status:      "fail",
		Timezone:    "",
		Zip:         "",
	}
	result, _ := json.Marshal(l)
	return string(result)

}

func getMockServer(status int, responsePayload string) (*httptest.Server, *http.Client) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, responsePayload)
	}))

	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}
	httpClient := &http.Client{Transport: transport}

	return server, httpClient
}

func TestGetLocationSuccess(t *testing.T) {
	server, httpClient := getMockServer(200, getMockSuccessResponse())
	client := Client{server.URL, httpClient}

	location, err := client.GetLocation()
	if err != nil {
		t.Error(err)
	}
	if location.Status != "success" {
		t.Error(errors.New("Expected success"))
	}
	if location.City != "Timbuktu" {
		t.Error(errors.New("Expected Timbuktu"))
	}
}

func TestGetLocationFailure(t *testing.T) {
	// A failure response still returns a 200
	server, httpClient := getMockServer(200, getMockFailureResponse())
	client := Client{server.URL, httpClient}

	location, err := client.GetLocation()
	if location != nil {
		t.Error("Location should be nil")
	}
	if err == nil {
		t.Error("Should have returned an error")
	}
}

func TestGetLocationHttpError(t *testing.T) {
	server, httpClient := getMockServer(500, "")
	client := Client{server.URL, httpClient}

	location, err := client.GetLocation()
	if location != nil {
		t.Error("Location should be nil")
	}
	if err == nil {
		t.Error("Should have returned an error")
	}
}
