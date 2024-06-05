package utils

var Math = new(m)

type m struct {
}

/**
 * @Description: 取最大
 * @receiver m
 * @param y
 * @return int
 */
func (m) Max(y ...int) int {
	max := y[0]
	if len(y) == 1 {
		return max
	}
	for i := 1; i < len(y); i++ {
		if y[i] > max {
			max = y[i]
		}
	}
	return max
}

/**
 * @Description: 取最小
 * @receiver m
 * @param y
 * @return int
 */
func (m) Min(y ...int) int {
	min := y[0]
	if len(y) == 1 {
		return min
	}
	for i := 1; i < len(y); i++ {
		if y[i] < min {
			min = y[i]
		}
	}
	return min
}
