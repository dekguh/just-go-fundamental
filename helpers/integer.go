package helpers

func SumTotal(total []int) int {
	if len(total) == 0 {
		return 0
	}
	sum := 0
	for _, value := range total {
		sum += value
	}

	return sum
}
