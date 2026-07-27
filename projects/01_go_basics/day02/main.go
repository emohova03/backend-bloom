package main

import "fmt"

type Priority int

const (
	Low Priority = iota
	Medium
	High
)

func main() {
	var title string = "Buy milk"
	done := false
	priority := High
	estimatedMinutes := 15.5
	var title2 string = "Learn Go"
	done2 := false
	priority2 := Medium
	estimatedMinutes2 := 1030.0

	fmt.Printf("Задача: %s\n", title)
	fmt.Printf("Приоритет: %d (High = %d)\n", priority, High)
	fmt.Printf("Готово: %t\n", done)
	fmt.Printf("Оценка времени: %.1f минут\n", estimatedMinutes)
	fmt.Println("\n")

	fmt.Printf("Задача: %s\n", title2)
	fmt.Printf("Приоритет: %d (Medium = %d)\n", priority2, Medium)
	fmt.Printf("Готово: %t\n", done2)
	fmt.Printf("Оценка времени: %.1f минут\n", estimatedMinutes2)
	fmt.Printf("Тип: title %T\n", title)

	fmt.Printf("Всего: %.1f минут\n", estimatedMinutes+estimatedMinutes2)

}
