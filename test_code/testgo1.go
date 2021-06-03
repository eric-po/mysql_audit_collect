package main

import (
	"fmt"
	"time"
)

func main() {
	//listOfNumberStrings := []*string{}
	//
	//for i := 0; i < 10; i++ {
	//	var numberString string
	//	numberString = fmt.Sprintf("#%s", strconv.Itoa(i))
	//	listOfNumberStrings = append(listOfNumberStrings, &numberString)
	//}
	//
	//for _, n := range listOfNumberStrings {
	//	fmt.Printf("%s\n", *n)
	//}
	//
	//return
	s := "20210603 13:48:59"
	//str := "2016-07-25 11:45:26"
	t, _ := time.Parse("20060102 15:04:05", s)
	a := t.Format("2006-01-02 15:04:05")
	fmt.Println(a)
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}
