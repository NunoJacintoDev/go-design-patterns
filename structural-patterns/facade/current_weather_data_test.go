package facade

import (
	"strings"
	"testing"
)

func TestOpenWeatherMap_responseParser(t *testing.T) {
	r := getMockData()
	openWeatherMap := CurrentWeatherData{APIkey: ""}
	weather, err := openWeatherMap.responseParser(r)
	if err != nil {
		t.Fatal(err)
	}
	if weather.ID != 3117735 {
		t.Errorf("Madrid id is 3117735, not %d\n", weather.ID)
	}
}

func TestOpenWeatherMap_Invalid_API_KEY(t *testing.T) {
	weatherMap := CurrentWeatherData{APIkey: ""}
	_, err := weatherMap.GetByCityAndCountryCode("Madrid", "ES")
	if err == nil && strings.Contains(err.Error(), "Invalid API key") {
		t.Error("Expected Invalid API Key error got ", err)
	}
}

// TODO: Need to test with a valid API Key
