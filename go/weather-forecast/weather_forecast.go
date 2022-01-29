// Package weather provides forecasst functionality.
package weather

// CurrentCondition represents a city's current condition.
var CurrentCondition string

// CurrentLocation location represents a city.
var CurrentLocation string

// Forecast takes a city and a condition and returns a description of the condition of that specific city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
