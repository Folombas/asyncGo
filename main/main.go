package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Воркер %d начал задание %d\n", id, job)
		time.Sleep(time.Second) // имитация работы
		results <- job * 2
		fmt.Printf("Воркер %d завершил задание %d\n", id, job)
	}
}

// Прикладное применение каналов и горутин
func main() {
	jobs := make(chan int, 3)
	results := make(chan int, 3)

	// Запуск 3 воркеров
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	// Отправка 5 заданий
	for i := 1; i <= 5; i++ {
		jobs <- i
		fmt.Printf("Задание %d отправлено\n", i)
	}
	close(jobs) // Закрыть канал после отправки

	// Сбор результатов
	for i := 1; i <= 5; i++ {
		fmt.Printf("Результат: %d\n", <-results)
	}
}
