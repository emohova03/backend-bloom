package main

import "fmt"

func removeTask(tasks []string, index int) []string {
	if index < 0 || index >= len(tasks) {
		fmt.Println("Некорректный индекс")
		return tasks
	}

	fmt.Printf("\nЗадача %d выполнена, убираем из списка", index+1)
	return append(tasks[:index], tasks[index+1:]...)
}
func main() {
	var tasks []string
	tasks = append(tasks, "Buy milk")
	tasks = append(tasks, "Call mom")
	tasks = append(tasks, "Fix bug", "Write report")

	fmt.Println("Все задачи: ")
	for i, t := range tasks {
		fmt.Printf(" %d. %s\n", i+1, t)
	}

	fmt.Printf("\nВсего задач: %d (cap: %d)\n", len(tasks), cap(tasks))

	firstTwo := tasks[:2]
	fmt.Println("\nПервые две: ", firstTwo)

	tasks = removeTask(tasks, 1)

	fmt.Println("\nДействительные задачи: ")
	for i, t := range tasks {
		fmt.Printf(" %d. %s\n", i+1, t)
	}
}
