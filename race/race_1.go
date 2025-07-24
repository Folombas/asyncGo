package main

import (
	"fmt"
	"time"
)

var counter int // Общая переменная

func main() {
	fmt.Printf("Начальное значение: %d\n", counter)

	go increment() // Горутина 1 (пишет)
	go increment() // Горутина 2 (пишет)

	// Даём горутинам время на выполнение
	time.Sleep(1 * time.Second)

	fmt.Printf("Конечное значение: %d\n", counter)
}

func increment() {
	temp := counter // Чтение
	temp++          // Изменение
	counter = temp  // Запись

	// Вывод внутри горутины для демонстрации  промежуточных состояний
	fmt.Printf("-> Горутина изменила значение: %d\n", counter)
}
