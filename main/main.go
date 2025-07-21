package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1) // +1 к счетчику
		go func(id int) {
			defer wg.Done() // Гарантированный -1 при выходе
			fmt.Printf("Горутина %d завершена\n", id)
		}(i)
	}

	wg.Wait() // Ждём обнуления счетчика
	fmt.Println("Все горутины завершились")
}
