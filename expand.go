package main

import (
	"fmt"
	"os"
	"time"
	"strconv"
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
		min := input[len(input) - 2:]
		hr := input[1:len(input) - 2]
		time := hr + ":" + min
		var stmt = "Pick up: " + time
		fmt.Println(stmt)
		output = output + stmt + "\n"
	case 'd':
		min := input[len(input) -2:]
		hr := input[1:len(input) - 2]
		time := hr + ":" + min
		var stmt = "Drop off: " + time
		fmt.Println(stmt)
		output = output + stmt + "\n"
	case 'w':
		stmt := "Weight: " + input[1:] + " lbs"
		fmt.Println(stmt)
		output = output + stmt + "\n"
	case 'r':
		stmt := "Rate: " + "$" + input[1:] + " per mile"
		fmt.Println(stmt)
		output = output + stmt + "\n"
	case 't':
		stmt := "Temperature: " + input[1:] + "°"
		fmt.Println(stmt)
		output = output + stmt + "\n"
	case 'o':
		month := input[1:3]
		day := input[3:5]
		var year string
		if len(input) == 5 {
			t := time.Now()
			year = strconv.Itoa(t.Year())
		} else {
			year = input[6:]
		}
		date := month + "/" + day + "/" + year
		fmt.Println(date)
		output = output + date + "\n"
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
