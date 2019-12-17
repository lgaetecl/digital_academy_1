package main

import "fmt"

func factorial(num uint64) (result uint64) {
	if num > 0 {
		result = num * factorial(num-1)
		return
	}
	return 1
}

func main() {
	var n uint64
	fmt.Println("Ingrese un número")
	fmt.Scanf("%d", &n)

	if n <= 50 {
		fmt.Printf("Factorial de %d es: %d \n", n, factorial(n))
	} else {
		fmt.Println("Favor ingresar un número menor a 50")
	}

}
