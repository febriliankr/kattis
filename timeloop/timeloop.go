package main

import (
	"fmt"
	"strconv"
)

func main() {
	var limit int
	fmt.Scanln(&limit)
	for i := 1; i <= limit; i++ {
		var number = strconv.Itoa(i)
		fmt.Println(number + " " + "Abracadabra")
	}
	fmt.Scanln(&limit)
	for i := 1; i<=limit; i++ {
		fmt.Println(strconv.Itoa(i))

	}
}


