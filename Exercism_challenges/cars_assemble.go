package cars
// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return (successRate/100) * float64(productionRate)
}
// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var per_hr = CalculateWorkingCarsPerHour(productionRate, successRate)
    var per_min = per_hr/60
    return int(per_min)
}
// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var grp_ten = carsCount /10
    var indi_car = carsCount % 10
    var results =(grp_ten * 95000)  + (indi_car * 10000)
    return uint(results)
}