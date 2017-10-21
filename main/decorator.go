package main

import "fmt"

type DecoratorPoint struct {
	x int
	y int
}

//Main compute function
func compute(p DecoratorPoint, f func(*DecoratorPoint)) {
	f(&p)

	fmt.Printf("%v, %v", p.x, p.y)
}

//Decorator function
func slice(p *DecoratorPoint) {
	p.x = p.x / 2
	p.y = p.y / 2
}

func main() {
	p := DecoratorPoint{2, 4}

	compute(p, slice) //1, 2
}
