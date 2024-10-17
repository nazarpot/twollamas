package main

import (
	"fmt"
	"os"
	"golang.design/x/clipboard"
)

func copy(output string) {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}

	clipboard.Write(clipboard.FmtText, []byte(output))
}

func processer(input string) string {
	output := ""
	var v = input[0]
	switch v {
	case 'p':
		var stmt = "Pick up: " + input[1:]
		fmt.Println(stmt)
		output = output + stmt + "\n"
	case 'd':
		var stmt = "Drop off: " + input[1:]
		fmt.Println(stmt)
		output = output + stmt + "\n"
	case 'w':
		stmt := "Weight: " + input[1:]
		fmt.Println(stmt)
		output = output + stmt + "\n"
	case 'r':
		stmt := "Rate: " + input[1:]
		fmt.Println(stmt)
		output = output + stmt + "\n"
	case 't':
		stmt := "Temperature: " + input[1:]
		fmt.Println(stmt)
		output = output + stmt + "\n"
	}
	return output
}

func reciever() {
	args := os.Args[1:]

	output := ""
	for i := 0; i < len(args); i++ {
		line := processer(args[i])
		output += line
	}
	copy(output)
}
func main() {
	reciever()
}
