package purchase

import "strings"

func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
 }

func ChooseVehicle(option1, option2 string) string {
	decisionString := " is clearly the better choice."
	if strings.Compare(option1, option2) < 0 {
		return option1 + decisionString
	}
	
	return option2 + decisionString
}

func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3{
		return originalPrice * 0.8
	} else if age >= 10 {
		return originalPrice * 0.5
	} else {
		return originalPrice * 0.7
	}
}
