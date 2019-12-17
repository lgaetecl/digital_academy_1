package main

import "fmt"

func a(num int) (s string) {
	if num > 0 {
		s = "a" + (a(num - 1))
		return
	}
	return
}

func b(num int) (s string) {
	if num > 0 {
		s = "b" + (b(num - 1))
		return
	}
	return
}

func main() {
	var n int
	fmt.Println("Ingrese un número")
	fmt.Scanf("%d", &n)

	fmt.Println(a(n) + b(n))
}
