package goip

import (
	"fmt"
	"testing"
)

func TestGetLocation(t *testing.T) {
	result, err := GetLocation()
	if err != nil {
		t.Fatalf("Unexpected error", err)
	}
	fmt.Println(result.City)
}
