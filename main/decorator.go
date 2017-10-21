package main

import "fmt"

type DecoratorPoint struct {
	x int
	y int
}

//Main compute function
func compute(p DecoratorPoint, decorators ... func(*DecoratorPoint)) {
	for _, v := range decorators {
		v(&p)
	}

	fmt.Printf("%v, %v", p.x, p.y)
}

//Slice decorator function
func slice(p *DecoratorPoint) {
	p.x = p.x / 2
	p.y = p.y / 2
}

//Slice decorator function
func add(p *DecoratorPoint) {
	p.x = p.x + 1
}

func main() {
	p := DecoratorPoint{2, 4}

	compute(p, slice, add) //2, 2
}
