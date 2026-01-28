// Package weather contains variables of forecasting current condition and current location.
package weather

var (
	// CurrentCondition is the current weather condition.
	CurrentCondition string
	// CurrentLocation is the current weather location.
	CurrentLocation  string
)

// Forecast function use to forecast weather based on current city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
