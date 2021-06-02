package main

import "fmt"

type info struct {
	a int
}

func main() {
	var ss []*info
	ss = append(ss, &info{2})
	ss = append(ss, &info{2})
	ss = append(ss, &info{9})
	ss = append(ss, &info{2})
	ss = append(ss, &info{2})
	ss = append(ss, &info{2})
	fmt.Println(ss)
	for i, v := range ss {
		fmt.Println(i, v.a)
	}
}
