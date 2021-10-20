package main

import (
	"fmt"
	"http-scaffold/api"
	"http-scaffold/connectors"
	"http-scaffold/services"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// initialize configuration
	var (
		portStr        = os.Getenv("PORT")
		openWeatherKey = os.Getenv("OPEN_WEATHER_KEY")
		port, err      = strconv.Atoi(portStr)
	)

	if err != nil {
		panic(err)
	}

	addr := fmt.Sprintf("localhost:%d", port)

	config := api.Config{
		Port:            int(port),
		LoggingEnabled:  true,
		TimingEnabled:   true,
		TimingThreshold: 30000,
	}

	// initialize logging
	infoLog := log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)

	// initialize connectors
	openWeatherConnector := connectors.NewOpenWeather(openWeatherKey)

	// initialize services
	apiServices := api.Services{
		Greeting: &services.GreetingService{},
		Weather:  services.NewWeatherService(openWeatherConnector, infoLog, errorLog),
	}

	// initialize api
	server := api.New(
		config,
		api.Logging{
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
		apiServices)

	router := server.Router()

	log.Fatal(http.ListenAndServe(addr, http.Handler(router)))
}
