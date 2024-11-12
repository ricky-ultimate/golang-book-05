package main

import "fmt"

func main() {
	question3()
}

func question2() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

func question3() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func question3_2() {
	for i := 0; i <= 100; i++ {
		switch {
		case (i != 0 && i%3 == 0 && i%5 == 0):
			fmt.Println("FizzBuzz")
		case (i != 0 && i%3 == 0):
			fmt.Println("Fizz")
		case (i != 0 && i%5 == 0):
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
