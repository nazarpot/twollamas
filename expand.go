package main

import (
	"fmt"
	"os"
)

func processer(input string) {
	var v = input[0]
	switch v {
	case 'p':
		fmt.Println("Pick up: ", input[1:])
	case 'd':
		fmt.Println("Drop off: ", input[1:])
	case 'w':
		fmt.Println("Weight: ", input[1:])
	case 'r':
		fmt.Println("Rate: ", input[1:])
	case 't':
		fmt.Println("Temperature:", input[1:])
	}

}

func reciever() {
	args := os.Args[1:]

	for i := 0; i < len(args); i++ {
		processer(args[i])
	}
}
func main() {
	reciever()
}
