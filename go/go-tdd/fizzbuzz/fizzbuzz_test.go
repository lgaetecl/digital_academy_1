package fizzbuzz_test

import (
	"testing"

	"go-tdd/fizzbuzz"
)

func TestFizzBuzz(t *testing.T) {

	t.Run("Si la entrada es 5 retornar Buzz", func(t *testing.T) {
		var result = fizzbuzz.FizzBuzz(5)

		if result != "Buzz" {
			t.Error("Resultado no es el esperado: " + result)
		}
	})

	t.Run("Si la entrada es 10 retornar Buzz", func(t *testing.T) {
		var result = fizzbuzz.FizzBuzz(10)

		if result != "Buzz" {
			t.Error("Resultado no es el esperado: " + result)
		}
	})

	t.Run("Si la entrada es 3 retornar Fizz", func(t *testing.T) {
		var result = fizzbuzz.FizzBuzz(3)

		if result != "Fizz" {
			t.Error("Resultado no es el esperado: " + result)
		}
	})

	t.Run("Si la entrada es 15 retornar FizzBuzz", func(t *testing.T) {
		var result = fizzbuzz.FizzBuzz(15)

		if result != "FizzBuzz" {
			t.Error("Resultado no es el esperado: " + result)
		}
	})

	t.Run("Si la entrada es 8 retornar '8'", func(t *testing.T) {
		var result = fizzbuzz.FizzBuzz(8)

		if result != "8" {
			t.Error("Resultado no es el esperado: " + result)
		}
	})

	t.Run("Si la entrada es 7 retornar '7'", func(t *testing.T) {
		var result = fizzbuzz.FizzBuzz(7)

		if result != "7" {
			t.Error("Resultado no es el esperado: " + result)
		}
	})
}

func BenchmarkFizzBuzz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fizzbuzz.FizzBuzz(i)
	}
}
