package main

import (
	"fmt"
	"strconv"
)

const (
	fala  string = "Fala"
	bella string = "Bella"
)

func falabella(in int) string {
	switch {
	case in%15 == 0:
		return fala + bella
	case in%3 == 0:
		return fala
	case in%5 == 0:
		return bella
	default:
		return strconv.Itoa(in)
	}
}

func main() {

	for index := 1; index <= 100; index++ {
		fmt.Println(falabella(index))
	}

}
