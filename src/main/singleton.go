package main

import (
	"sync"
	"fmt"
)

type singletonPoint struct {
	x int
	y int
}

//Singleton instance
var instance singletonPoint = singletonPoint{}

//Lazy singleton instance
var lazyInstance singletonPoint

//Object performing exactly one action
var once sync.Once

func getPointInstance() singletonPoint {
	return instance
}

func getLazyPointInstance() singletonPoint {
	once.Do(func() {
		//The initialization can be done only once
		lazyInstance = singletonPoint{}
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
