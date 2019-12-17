package main

import (
	"fmt"
	"strconv"
)

const (
	fala  string = "Fala"
	bella string = "Bella"
)

func falabella(num int) string {
	if num%15 == 0 {
		return fala + bella
	}
	if num%3 == 0 {
		return fala
	} else if num%5 == 0 {
		return bella
	}
	return strconv.Itoa(num)
}

func main() {

	for index := 1; index <= 100; index++ {
		fmt.Println(falabella(index))
	}

}
