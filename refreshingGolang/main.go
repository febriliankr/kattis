package main

import (
	"errors"
	"fmt"
	"math"
)

func loop() {
	for i := 0; i < 5; i++ {
		fmt.Print(i)
	}
}

func arrayBasedLoop() {
	arr := []string{"one", "two", "three"}
	for index, value := range arr {
		fmt.Println("index", index, "value", value)
	}
}

func squareRoot(x float64) (float64, error) {

	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}
	return math.Sqrt(x), nil

}

func main() {
	// This is me learning data types all over again
	arrS := []int{1, 2, 3, 4}
	fmt.Println(arrS)

	arrS = append(arrS, 12, 2, 3)
	fmt.Println(arrS)

	vertices := make(map[string]int)
	vertices["triangle"] = 3
	vertices["square"] = 3
	vertices["octagon"] = 31

	// we can also use interface{} as value to set custom value
	// not best practice
	userObj := make(map[string]interface{})
	userObj["febrilian"] = map[string]interface{}{
		"name":  "febrilian kristiawan",
		"age":   22,
		"phone": "085281812202",
	}
	userObj["Rayhan"] = map[string]interface{}{
		"name":  "muhammad rayhan syahman",
		"age":   22,
		"tel":   2323,
		"phone": "128310283",
	}

	fmt.Println(userObj)
	delete(userObj, "Rayhan")
	fmt.Println(userObj)
	loop()
	arrayBasedLoop()
	s, err := squareRoot(-1)
	if err != nil {
		fmt.Println(err)
	}
	if err == nil {
		fmt.Println(s)
	}

	// slices
	superString := "superhuman"
	sliceOfString := superString[1:3]
	fmt.Println(sliceOfString)
}
