package services

import (
	"http-scaffold/models"
	"log"
)

type WeatherService struct {
	connector forecastFetcher
	infoLog   *log.Logger
	errorLog  *log.Logger
}

type forecastFetcher interface {
	FetchForecast(query string) (models.Forecast, error)
}

func NewWeatherService(fetcher forecastFetcher, infoLog, errorLog *log.Logger) *WeatherService {
	return &WeatherService{
		fetcher,
		infoLog,
		errorLog,
	}
}

func (w *WeatherService) GetForecastForLocation(location string) (models.Forecast, error) {
	w.infoLog.Printf("fetching forecast for %s\n", location)

	forecast, err := w.connector.FetchForecast(location)
	if err != nil {
		w.errorLog.Println(err.Error())
		return models.Forecast{}, nil
	}

	return forecast, nil
}
