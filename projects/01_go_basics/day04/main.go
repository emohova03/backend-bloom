package main

import "fmt"

func validateTitle(title string) (bool, error) {
	if title == "" {
		return false, fmt.Errorf("title не может быть пустым")
	}
	return true, nil
}

func countUrgent(priorities []int, urgentLevel int) int {
	count := 0
	for _, p := range priorities {
		if p == urgentLevel {
			count++
		}
	}
	return count
}

func filter(nums []int, predicate func(int) bool) []int {
	numResult := []int{}
	for _, n := range nums {
		if predicate(n) {
			numResult = append(numResult, n)
		}

	}
	return numResult

}

func main() {
	titles := []string{"Buy milk", "", "Write report"}
	for _, t := range titles {
		ok, err := validateTitle(t)
		if !ok {
			fmt.Println("x Невалидная задача: ", err)
			continue
		}
		fmt.Println("V Задача валидна: ", t)
	}

	priorities := []int{0, 2, 1, 2, 2}
	fmt.Println("Срочных задач: ", countUrgent(priorities, 2))

	evens := filter([]int{1, 2, 3, 4, 5, 6, 7, 8}, func(n int) bool {
		return n%2 == 0
	})

	fmt.Println(evens)

}
