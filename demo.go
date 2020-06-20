package main

import (
	"fmt"
	"time"
)

type X int

func (x *X) inc() {
	*x++
}

type user struct {
	name string
	age  byte
}

func (u user) toString() string {
	return fmt.Sprintf("%+v", u)
}

type manager struct {
	user
	title string
}

type Printer interface {
	toString() string
}

func task(id int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d: %d\n", id, i)
		time.Sleep(time.Second)
	}
}

func consumer(data chan int, done chan bool) {
	for x := range data {
		//time.Sleep(time.Second)
		println("rec:", x)
	}
	done <- true
}

func producer(data chan int) {
	for i := 0; i < 4; i++ {
		data <- i
	}
	close(data)
}

func foo() {
	done := make(chan bool)
	data := make(chan int)

	go consumer(data, done)
	go consumer(data, done)
	go consumer(data, done)
	go consumer(data, done)
	go producer(data)

	<-done
	<-done
	<-done
	<-done
}

func demo() {
	type (
		user struct {
			name string
			age  byte
		}
		event func(string) bool
	)
	u := user{"foo", 32}
	fmt.Println(u)
	var f event = func(s string) bool {
		println(s)
		return s != ""
	}
	b := f("abc")
	println(b)

}
