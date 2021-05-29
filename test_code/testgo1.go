package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hehe"
	s1 := [1]string{s}

	a := strings.Split(s, ",")
	a1 := strings.Split(s, ",")
	fmt.Println(a1)
	fmt.Println(a[1])
}
