package cars

const ProductionRate = 221

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
	if speed == 0 {
		return float64(0)
	} else if speed <= 4 {
		return float64(1)
	} else if speed <= 8 {
		return float64(.9)
	} else {
		return float64(.77)
	}
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	return SuccessRate(speed) * float64(speed) * ProductionRate
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	return int(CalculateProductionRatePerHour(speed) / 60)
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	theoretical := CalculateProductionRatePerHour(speed)
	if theoretical > limit {
		return limit
	} else {
		return theoretical
	}
}
