package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var result string
	if option1 < option2 {
		result = fmt.Sprintf("%s is clearly the better choice.", option1)
	} else {
		result = fmt.Sprintf("%s is clearly the better choice.", option2)
	}
	return result
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var multiplier float64
	if age < 3 {
		multiplier = .8
	} else if age >= 3 && age < 10 {
		multiplier = .7
	} else {
		multiplier = .5
	}
	return originalPrice * multiplier
}
