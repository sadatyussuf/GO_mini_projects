package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	res:= 0
    for i:=0;i<len(birdsPerDay);i++{
        res += birdsPerDay[i] 
    }
return res
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    end_day := 7 * week
    res := 0
    
	for start_day := end_day - 7; start_day < end_day;start_day++{
        res += birdsPerDay[start_day]
    }
return res 
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i:=0;i<len(birdsPerDay);i+=2{
        birdsPerDay[i] = birdsPerDay[i] + 1
    }
return birdsPerDay
}
