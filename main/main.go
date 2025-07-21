package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Создаём WaitGroup
	var wg sync.WaitGroup

	// Запускаем 3 горутины
	for i := 1; i <= 3; i++ {
		wg.Add(1)         // Говорим WaitGroup: "Добавь одну задачу в ожидание"
		go worker(i, &wg) // Запускаем горутину, передаём ей указатель на wg
	}

	fmt.Println("Главная горутина: Жду завершения всех горутин...")
	wg.Wait() // Блокируем здесь, пока wg.Done() не будет вызван 3 раза
}

// Функция, выполняемая в горутине (воркер)
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Важно! Гарантирует, что wg.Done() вызовется при ЛЮБОМ выходе из функции

	fmt.Printf("Воркер &d: Начинаю работу...\n", id)
	time.Sleep(time.Duration(id) * time.Second) // Имитация разной по времени работы
	fmt.Printf("Воркер %d: Закончил работу\n", id)
}
