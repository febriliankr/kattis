package main

import "fmt"

func main() {
	var input = 0
	var firstInt = 0
	var secondInt = 0

	fmt.Scan(&firstInt)
	fmt.Scan(&secondInt)
	fmt.Scanln(&input)
	for i := 1; i <= input; i++ {
		var output = ""

		if i%firstInt == 0 {
			output = "Fizz"
		}
		if i%secondInt == 0 {
			output += "Buzz"
		}
		if output != "" {
			fmt.Println(output)
		} else {
			fmt.Println(i)
		}

	}
}
