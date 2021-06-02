//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func loop(n int, mark string) {
//	for i := 0; i < n; i++ {
//		fmt.Println("haha : ", i, mark)
//	}
//	time.Sleep(1)
//}
//
//func main() {
//	go loop(5, "first")
//	loop(8, "2nd")
//}
//
///////////////////////////////////////////

//package main
//
//import "fmt"
//
//func main() {
//	message := make(chan string)
//	go func() {
//		message <- "ping"
//	}()
//	msg := <-message
//	fmt.Println(msg)
//}
///////////////////////////////////////////

package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func consumer(stop <-chan bool) {
	for {
		select {
		case <-stop:
			fmt.Println("exit sub goroutine")
			return
		default:
			fmt.Println("running...")
			time.Sleep(20 * time.Second)
			//time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	stop := make(chan bool)
	var wg sync.WaitGroup
	// Spawn example consumers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(stop <-chan bool) {
			defer wg.Done()
			consumer(stop)
		}(stop)
	}
	waitForSignal()
	close(stop)
	fmt.Println("stopping all jobs!")
	wg.Wait()
}
func waitForSignal() {
	sigs := make(chan os.Signal)
	fmt.Println("hehe")
	signal.Notify(sigs, os.Interrupt)
	signal.Notify(sigs, syscall.SIGTERM)
	fmt.Println("haha")
	<-sigs
	fmt.Println("haha again")
}
