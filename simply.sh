#!/bin/bash
#script for simplifying input

print_cleaner() {
	str=$1
	type="${str:0:1}"
	val="${str:1}"
	if [ "$type" = "p" ]; then
		echo "Pick up: ${val} " 
	elif [ "$type" = "d" ]; then
		echo "Drop off: ${val}"
	elif [ "$type" = "w" ]; then
		echo "Weight: ${val}"
	elif [ "$type" = "r" ]; then
		echo "Rate: ${val} pm"
	elif [ "$type" = "t" ]; then
		echo "Temp: ${val}"
	fi
}

for var in "$@"
do 
	print_cleaner $var
done

