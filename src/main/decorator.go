//Inspired by http://blog.ralch.com/tutorial/design-patterns/golang-decorator/
package main

import (
	"fmt"
)

const id = "id"
const retry = "retry"

type Retrieve interface {
	retrieve(map[string]int) (string, error)
}

type Database struct {
	retrier DatabaseRetry
}

type DatabaseRetry struct {
}

func (d *Database) retrieve(args map[string]int) (string, error) {
	fmt.Printf("Retrieve %v\n", args[id])
	if args[id] == 7 {
		return "", fmt.Errorf("Transient database error")
	} else {
		return "foo", nil
	}
}

func (d *DatabaseRetry) retrieve(args map[string]int) (string, error) {
	fmt.Printf("Retry with %v\n", args[id])
	//Do something to handle an automated retry
	return "bar", nil
}

func Decorator() {
	database := Database{}
	databaseRetrier := database.retrier
	ids := [2]int{0, 7}
	args := make(map[string]int)
	args[retry] = 3

	for _, v := range ids {
		var i Retrieve = &database

		args[id] = v
		name, ok := i.retrieve(args)

		if ok == nil {
			fmt.Printf("Result: %v\n", name)
		} else {
			i := &databaseRetrier
			name, _ := i.retrieve(args) //We do not test the errors here for the sake of clarity
			fmt.Printf("Result: %v\n", name)
		}
	}
}
