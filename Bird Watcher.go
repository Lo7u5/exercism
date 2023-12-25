package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sum := 0
    for _, v := range birdsPerDay {
        sum += v
    }
	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	cond := week * 7
    init := cond - 7
    sum := 0
    for i := init; i < cond; i++ {
        sum += birdsPerDay[i]
    }
	return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
        if i % 2 != 0 {
            continue
        }
    	birdsPerDay[i]++
    }
	return birdsPerDay
}
