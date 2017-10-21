package main

import "fmt"

type strategyPoint struct {
	x int
	y int
}

//Main compute function
func compute(p strategyPoint, strategy func(*strategyPoint)) {
	//Execute the given strategy
	strategy(&p)

	fmt.Printf("%v, %v\n", p.x, p.y)
}

//Slice strategy
func slice(p *strategyPoint) {
	p.x = p.x / 2
	p.y = p.y / 2
}

//Add strategy
func add(p *strategyPoint) {
	p.x = p.x + 1
}

func main() {
	p := strategyPoint{4, 2}

	compute(p, slice) //2, 1
	compute(p, add)   //5, 2
}
