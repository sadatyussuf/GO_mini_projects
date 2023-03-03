// Package weather forecasts the current weather 
// conditions in several Goblinocus cities.
package weather
// CurrentCondition is a variable that returns the current weather condition.
var CurrentCondition string
// CurrentLocation is a variable that returns the current location.
var CurrentLocation string
// Forecast returna string value of the current location and  weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}