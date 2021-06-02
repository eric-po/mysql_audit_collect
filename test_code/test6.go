package main

import (
	"github.com/Shopify/sarama"
)

type Test struct {
	Value int
}

func buildStruct(count int) *[]Test {

	slice := make([]Test, 0)

	for i := 0; i < count; i++ {
		t := Test{Value: i}
		slice = append(slice, t)
	}
	return &slice
}

func main() {
	//fmt.Println("hehe haha")
	//a, _ := os.LookupEnv("KEY")
	//b, _ := strconv.Atoi(a)
	//slice := buildStruct(b)
	//for _, v := range *slice { //range 切片指针的正确方法
	//	//k为索引号,v为结构体
	//	fmt.Println(v.Value) //访问结构体的Value
	//}
	var bb *sarama.Message
	listOfNumberMessages := []*sarama.ConsumerMessage{}
	for i := 0; i < 10; i++ {
		var numberString string
		//numberString = fmt.Sprintf("#%s", strconv.Itoa(i))
		listOfNumberMessages = append(listOfNumberMessages, &numberString)
	}
}
