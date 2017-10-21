package main

import (
	"fmt"
)

type observable interface {
	register(string, func())
	notify()
}

type observerPoint struct {
	x           int
	y           int
	subscribers map[string]func()
}

//Sample method
func (point *observerPoint) changeXValue(x int) {
	point.x = x
	fmt.Printf("Change x value to %v\n", x)
	point.notify() //Notify the subscribers
}

//Register another subscriber
func (point *observerPoint) register(subscriber string, f func()) {
	point.subscribers[subscriber] = f
}

//Handle notifications
func (point *observerPoint) notify() {
	for k, v := range point.subscribers {
		fmt.Printf("Notifying subscriber %v\n", k)
		v()
	}
}

func Observer() {
	point := observerPoint{4, 2, make(map[string]func())}
	var pointObservable observable
	pointObservable = &point

	pointObservable.register("foo", func() {
		fmt.Println("Do something to notify foo")
	})

	pointObservable.register("bar", func() {
		fmt.Println("Do something to notify bar")
	})

	point.changeXValue(0)
}
