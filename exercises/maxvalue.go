package exercises

func MaxValue(array []int) int{
	max := array[0] // we expect to have at least one item in the array

	for _, element := range array[1:] {
		if element > max {
			max = element
		}
	}

	return max
}