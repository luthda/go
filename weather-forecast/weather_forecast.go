// Package weather provides functoinality make a forcast based on lcoation and condition.
package weather

// CurrentCondition represents weather actual condition.
var CurrentCondition string
// CurrentLocation represents location of weather.
var CurrentLocation string

// Forecast returns a weatherforecast based on city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
