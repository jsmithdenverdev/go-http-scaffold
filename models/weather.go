package models

type Forecast struct {
	Location    Location
	Description string

	Temp         float64
	MaxTemp      float64
	MinTemp      float64
	RealFeelTemp float64

	Pressure   int64
	Humidity   int64
	Visibility int64

	WindSpeed     float64
	WindGust      float64
	WindDirection int64

	Sunrise int64
	Sunset  int64
}
