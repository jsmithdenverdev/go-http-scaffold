package api

import (
	"github.com/gorilla/mux"
	"http-scaffold/services"
	"log"
)

type Api struct {
	config Config

	infoLog  *log.Logger
	errorLog *log.Logger

	greetingService *services.GreetingService
	weatherService  *services.WeatherService
}

type Logging struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

type Services struct {
	Greeting *services.GreetingService
	Weather  *services.WeatherService
}

func New(config Config, logging Logging, services Services) *Api {
	return &Api{
		config:          config,
		infoLog:         logging.InfoLog,
		errorLog:        logging.ErrorLog,
		greetingService: services.Greeting,
		weatherService:  services.Weather,
	}
}

func (a *Api) Router() *mux.Router {
	router := mux.NewRouter()

	if a.config.LoggingEnabled {
		router.Use(useLogging(a.infoLog))
	}

	if a.config.TimingEnabled {
		router.Use(useTiming(a.errorLog, a.config.TimingThreshold))
	}

	// wire up routes
	routes(router, a)

	return router
}
