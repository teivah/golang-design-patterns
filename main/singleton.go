package main

import (
	"sync"
	"fmt"
)

type point struct {
	x int
	y int
}

//Singleton instance
var instance point = point{}

//Lazy singleton instance
var lazyInstance point

//Object performing exactly one action
var once sync.Once

func getPointInstance() point {
	return instance
}

func getLazyPointInstance() point {
	once.Do(func() {
		//The initialization can be done only once
		lazyInstance = point{}
	})

	return lazyInstance
}

func main() {
	i1 := getPointInstance()
	fmt.Printf("%p\n", &i1) //0xc042008160

	i1 = getPointInstance()
	fmt.Printf("%p\n", &i1) //0xc042008160

	i2 := getLazyPointInstance()
	fmt.Printf("%p\n", &i2) //0xc042008190

	i2 = getLazyPointInstance()
	fmt.Printf("%p\n", &i2) //0xc042008190
}
