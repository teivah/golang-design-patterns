package main

import (
	"sync"
	"fmt"
)

type SingletonPoint struct {
	x int
	y int
}

//Singleton instance
var instance SingletonPoint = SingletonPoint{}

//Lazy singleton instance
var lazyInstance SingletonPoint

//Object performing exactly one action
var once sync.Once

func getPointInstance() SingletonPoint {
	return instance
}

func getLazyPointInstance() SingletonPoint {
	once.Do(func() {
		//The initialization can be done only once
		lazyInstance = SingletonPoint{}
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
