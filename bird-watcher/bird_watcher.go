package birdwatcher

func TotalBirdCount(birdsPerDay []int) int {
	count := 0

	for i := 0; i < len(birdsPerDay); i++ {
		count += birdsPerDay[i]
	}

	return count
}

func BirdsInWeek(birdsPerDay []int, week int) int {
	newIndex := (week - 1) * 7

	return TotalBirdCount(birdsPerDay[newIndex : newIndex + 7])
}

func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i] += 1
	}

	return birdsPerDay
}
