package api

import (
	"encoding/json"
	"http-scaffold/models"
	"net/http"
)

func (a *Api) handleGreet() http.HandlerFunc {
	type greetingResponse struct {
		Greeting string `json:"greeting"`
	}

	return func(rw http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		name := query.Get("name")
		greeting := a.greetingService.Greet(name)

		enc := json.NewEncoder(rw)
		err := enc.Encode(&greetingResponse{
			Greeting: greeting,
		})

		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}

		rw.WriteHeader(http.StatusOK)
	}
}

func (a *Api) handleForecast() http.HandlerFunc {
	type forecastResponse struct {
		Forecast models.Forecast `json:"forecast"`
	}

	return func(rw http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		location := query.Get("location")
		forecast, err := a.weatherService.GetForecastForLocation(location)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}

		enc := json.NewEncoder(rw)
		err = enc.Encode(&forecastResponse{
			forecast,
		})

		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}

		rw.WriteHeader(http.StatusAccepted)
	}
}
