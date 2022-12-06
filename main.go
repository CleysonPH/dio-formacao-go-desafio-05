package main

func sum(values ...int) int {
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

func sub(values ...int) int {
	result := 0
	for _, v := range values {
		result -= v
	}
	return result
}

func mul(values ...int) int {
	result := 1
	for _, v := range values {
		result *= v
	}
	return result
}

func div(values ...int) int {
	result := 1
	for _, v := range values {
		result /= v
	}
	return result
}
