// Package weather provides info about weather in chosen city.
package weather

// CurrentCondition represents weather certain condition.
var CurrentCondition string
// CurrentLocation represents certain city.
var CurrentLocation string

// Forecast returns a string with current weather in city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
