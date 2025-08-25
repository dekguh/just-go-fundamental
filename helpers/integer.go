package helpers

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func SumTotal(total []float32) float32 {
	if len(total) == 0 {
		return 0
	}
	sum := float32(0)
	for _, value := range total {
		sum += value
	}

	return sum
}

func InterpolateFormula(pt1, pt2 float32, duration int) []float32 {
	result := []float32{pt1}
	lenNumber := Abs(int(pt1 - pt2))
	ratio := (float32(lenNumber) / float32(duration))

	var i int = 0
	for i < duration {
		result = append(result, result[i]+ratio)
		i++
	}

	return result
}
