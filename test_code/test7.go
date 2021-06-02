package main

import "fmt"

type Human interface {
	Say()
	work()
}

//定义两个类，这两个类分别实现了 Human 接口的 Say 方法
type women struct {
	age int
}

type man struct {
	money int
}

func (m *women) work(location string, sarary string) string {
	fmt.Println(sarary, location)
	fmt.Println(m.age)
	return location
}

func (w *women) Say() {
	fmt.Printf("I'm a %d year old women\n\n", w.age)
}
func (m *man) Say() {
	fmt.Printf("I have money %d man\n\n", m.money)
}

type Child struct {
	book string
	song string
}

func (c *Child) Say(learn string) {
	fmt.Printf("I am learning sing %s and reading %s book", c.song, c.book)
	fmt.Printf("learn sill : %s \n", learn)
}
func main() {
	w := new(women)
	w.age = 30
	w.Say()
	m := new(man)
	m.money = 50
	m.Say()
	c := new(Child)
	c.book = "harry"
	c.song = "bird"
	skill := "swim"
	c.Say(skill)
}
