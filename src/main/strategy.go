package main

import (
	"fmt"
	"../util"
)

//Main compute function
func compute(p util.Point, strategy func(*util.Point)) {
	//Execute the given strategy
	strategy(&p)

	fmt.Printf("%v, %v\n", p.X, p.X)
}

//Slice strategy
func slice(p *util.Point) {
	p.X = p.X / 2
	p.Y = p.Y / 2
}

//Add strategy
func add(p *util.Point) {
	p.X = p.X + 1
}

func main() {
	p := util.Point{4, 2}

	compute(p, slice) //2, 1
	compute(p, add)   //5, 2
}
