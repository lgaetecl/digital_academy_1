package fizzbuzz

import "fmt"

//FizzBuzz func
func FizzBuzz(input int) string {
	switch {
	case input%15 == 0:
		return "FizzBuzz"
	case input%3 == 0:
		return "Fizz"
	case input%5 == 0:
		return "Buzz"
	default:
		return fmt.Sprintf("%d", input)
	}
}
