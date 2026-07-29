package main

import "fmt"

type Priority int

const (
	Low Priority = iota
	Medium
	High
)

func main() {
	priorities := []Priority{Low, Medium, High}

	urgentCount := 0

	for i, p := range priorities {
		if p == High {
			fmt.Printf("Задача #%d - срочная\n", i)
			urgentCount++
		}
	}
	fmt.Printf("Всего срочных задач: %d из %d\n", urgentCount, len(priorities))

	if urgentCount > len(priorities)/2 {
		fmt.Println("Слишком много срочных задач - пора расставлять приоритеты честнее")
	}

	//FizzBuzz exercise

	for i := 1; i <= 30; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
