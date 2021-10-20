package connectors

import (
	"encoding/json"
	"fmt"
	"http-scaffold/models"
	"io"
	"net/http"
)

type OpenWeather struct {
	key    string
	client *http.Client
}

type openWeatherResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int64  `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int64   `json:"pressure"`
		Humidity  int64   `json:"humidity"`
	} `json:"main"`
	Visibility int64 `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int64   `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Clouds struct {
		All int64 `json:"all"`
	} `json:"clouds"`
	Dt  int64 `json:"dt"`
	Sys struct {
		Type    int64  `json:"type"`
		ID      int64  `json:"id"`
		Country string `json:"country"`
		Sunrise int64  `json:"sunrise"`
		Sunset  int64  `json:"sunset"`
	} `json:"sys"`
	Timezone int64  `json:"timezone"`
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Cod      int64  `json:"cod"`
}

func NewOpenWeather(key string) *OpenWeather {
	client := &http.Client{}

	return &OpenWeather{
		key,
		client,
	}
}

func (ow *OpenWeather) FetchForecast(query string) (forecast models.Forecast, err error) {
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&APPID=%s", query, ow.key)

	result, err := ow.client.Get(url)
	if err != nil {
		return models.Forecast{}, err
	}

	defer func(body io.ReadCloser) {
		err = body.Close()
	}(result.Body)

	res := openWeatherResponse{}
	dec := json.NewDecoder(result.Body)
	err = dec.Decode(&res)

	if err != nil {
		return models.Forecast{}, err
	}

	return models.Forecast{
		Location: models.Location{
			Latitude:  res.Coord.Lat,
			Longitude: res.Coord.Lon,
			Name:      res.Name,
		},
		Description:   res.Weather[0].Main,
		Temp:          res.Main.Temp,
		MaxTemp:       res.Main.TempMax,
		MinTemp:       res.Main.TempMin,
		RealFeelTemp:  res.Main.FeelsLike,
		Pressure:      res.Main.Pressure,
		Humidity:      res.Main.Humidity,
		Visibility:    res.Visibility,
		WindSpeed:     res.Wind.Speed,
		WindGust:      res.Wind.Gust,
		WindDirection: res.Wind.Deg,
		Sunrise:       res.Sys.Sunrise,
		Sunset:        res.Sys.Sunset,
	}, nil
}
