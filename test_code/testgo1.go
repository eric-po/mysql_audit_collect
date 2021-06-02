package main

import (
	"fmt"
	"strconv"
)

func main() {
	listOfNumberStrings := []*string{}

	for i := 0; i < 10; i++ {
		var numberString string
		numberString = fmt.Sprintf("#%s", strconv.Itoa(i))
		listOfNumberStrings = append(listOfNumberStrings, &numberString)
	}

	for _, n := range listOfNumberStrings {
		fmt.Printf("%s\n", *n)
	}

	return
}
