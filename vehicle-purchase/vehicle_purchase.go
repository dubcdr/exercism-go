package purchase

import (
	"fmt"
	"sort"
)

// NeedsLicense determines whether a license is need to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	switch {
	case kind == "car" || kind == "truck":
		return true
	default:
		return false
	}
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	strs := []string{option1, option2}
	sort.Strings(strs)
	return fmt.Sprintf("%s is clearly the better choice.", strs[0])
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var depreciationFactor float64
	switch {
	case age < 3:
		depreciationFactor = 0.8
	case age < 10:
		depreciationFactor = 0.7
	default:
		depreciationFactor = 0.5
	}
	return depreciationFactor * originalPrice
}
