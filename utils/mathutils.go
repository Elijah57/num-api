package utils

import (
	"math"
	"strconv"
)

// Function to check if a number is prime
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Function to check if a number is perfect
func IsPerfect(n int) bool {
	sum := 0
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum == n
}

// Function to check if a number is Armstrong
func IsArmstrong(n int) bool {
	sum := 0
	numStr := strconv.Itoa(n)
	numDigits := len(numStr)
	for _, digitChar := range numStr {
		digit, _ := strconv.Atoi(string(digitChar))
		sum += int(Pow(float64(digit), float64(numDigits)))
	}
	return sum == n
}

// cal power
func Pow(base, exponent float64) float64 {
	return math.Pow(base, exponent)
}

// digit sum
func DigitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}