package main

import "fmt"

type observable interface {
	register(string, func())
	notify()
}

//Subscribers map
var subscribers = make(map[string]func())

type observerPoint struct {
	x int
	y int
}

//Example function on a point
func (point *observerPoint) changeXValue(x int) {
	point.x = x
	fmt.Printf("Change x value to %v\n", x)
	point.notify() //Notify the subscribers
}

//Register another subscriber
func (point *observerPoint) register(subscriber string, f func()) {
	subscribers[subscriber] = f
}

//Handle notifications
func (point *observerPoint) notify() {
	for k, v := range subscribers {
		fmt.Printf("Notifying subscriber %v\n", k)
		v()
	}
}

func main() {
	point := observerPoint{4, 2}
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
