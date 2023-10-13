package utils

// restricts number to a specific range
func RestrictNumber(number int, min int, max int) int {
	if number < min {
		return min
	} else if number > max {
		return max
	} else {
		return number
	}
}

func Validate(variable string, toType string) {
	
}