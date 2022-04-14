package cards

func GetItem(slice []int, index int) (int, bool) {
	if isOutOfBound(slice, index) {
		return 0, false
	}

	return slice[index], true
}

func SetItem(slice []int, index, value int) []int {
	if isOutOfBound(slice, index) {
		return append(slice, value)
	}

	slice[index] = value

	return slice
}

func PrefilledSlice(value, length int) []int {
	var slice []int
	
	for i := 0; i < length; i++ {
		slice = append(slice, value)
	}

	return slice;
}

func RemoveItem(slice []int, index int) []int {
	if isOutOfBound(slice, index) {
		return slice
	}
	return append(slice[:index], slice[index + 1:]...)
}

func isOutOfBound(slice []int, index int) bool {
	return index < 0 || index >= len(slice)
}