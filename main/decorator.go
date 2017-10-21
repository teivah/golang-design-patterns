package main

import "fmt"

type decoratorPoint struct {
	x int
	y int
}

//Main compute function
func compute(p decoratorPoint, decorators ... func(*decoratorPoint)) {
	for _, v := range decorators {
		v(&p)
	}

	fmt.Printf("%v, %v", p.x, p.y)
}

//Slice decorator function
func slice(p *decoratorPoint) {
	p.x = p.x / 2
	p.y = p.y / 2
}

//Add decorator function
func add(p *decoratorPoint) {
	p.x = p.x + 1
}

func main() {
	p := decoratorPoint{4, 2}

	compute(p, slice, add) //3, 1
}
