package client

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"net/http"
)

func GetWeatherByRegion(region string) (model.MainWeather, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://rg-weather-api.fly.dev/weather?region="+region, nil)
	if err != nil {
		return model.MainWeather{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return model.MainWeather{}, err
	}

	defer resp.Body.Close()

	var weather model.MainWeather

	err = json.NewDecoder(resp.Body).Decode(&weather)
	if err != nil {
		return model.MainWeather{}, err
	}

	return weather, nil
}
