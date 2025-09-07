package main

import (
	"fmt"
)

func testDefer() {
	fmt.Println("Function end")
	defer fmt.Println("Deferred call 1")
	defer fmt.Println("Deferred call 2")

	defer func() {
		fmt.Println("Deferred call 3")
		fmt.Println("Deferred call 4")
		fmt.Println("Deferred call 5")
	}()
}

func main() {
	var i int = 1

	if i >= 1 {
		fmt.Println("i is positive")
	} else if i < 0 {
		fmt.Println("i is negative")
	} else {
		fmt.Println("i is zero")
	}

	if i := 2; i == 2 {
		fmt.Println("i is 2")
	}

	var j int = 1

	for {
		fmt.Println(j)

		if j == 10 {
			break
		}

		j++
	}

	for j := 0; j < 5; j++ {
		if j == 3 {
			fmt.Println("Skipping 3")
			continue
		}

		fmt.Println(j)
	}

	fruits := [3]string{"Apple", "Banana", "Strawberry"}

	for index, fruit := range fruits {
		fmt.Printf("Fruit %d: %s\n", index, fruit)
	}

	for index, char := range "あいうえお" {
		fmt.Printf("Character %d: %d\n", index, char)
	}

	var k int = 5

	switch k {
	case 1, 2:
		fmt.Println("k is 1 or 2")
	case 3:
		fmt.Println("k is 3")
	case 4, 5:
		fmt.Println("k is 4 or 5")
	default:
		fmt.Println("k is something else")
	}

	var s string = "a"

	switch s {
	case "a":
		s += "b"
		fallthrough
	case "b":
		s += "c"
		fallthrough
	case "c":
		s += "d"
		fallthrough
	default:
		s += "e"
	}

	fmt.Println(s)

	switch n := 2; n {
	case 1, 2, 3, 4, 5:
		fmt.Println("n is between 1 and 5")
	case 6, 7, 8, 9, 10:
		fmt.Println("n is between 6 and 10")
	}

	var l interface{} = 1

	switch l.(type) {
	case int:
		fmt.Println("l is an integer")
	case string:
		fmt.Println("l is a string")
	default:
		fmt.Println("l is of another type")
	}

	var m interface{} = 10

	switch value := m.(type) {
	case int:
		fmt.Printf("m is an integer: %d\n", value)
	case string:
		fmt.Printf("m is a string: %s\n", value)
	}

	testDefer()
}
