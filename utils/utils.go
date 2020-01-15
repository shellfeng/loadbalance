package utils



func Gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}

	if b == 0 {
		return a
	}else {
		return Gcd(b, a%b)
	}


}

func NGcd(numbers []int, n int) int {
	if n == 1 {
		return numbers[n-1]
	}

	return Gcd(numbers[n-1], NGcd(numbers, n-1))
}

//Min return min number
func Min(a, b int) int {
	if a < b {
		return a
	}
	if b < a {
		return b
	}
	return a
}

//Max return max number
func Max(a, b int) int {
	if a > b {
		return a
	}
	if b > a {
		return b
	}
	return a
}
