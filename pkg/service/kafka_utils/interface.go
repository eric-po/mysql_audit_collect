package kafka_utils

import "fmt"

type IStore interface {
	call()
}


type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}
