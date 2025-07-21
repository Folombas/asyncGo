package main

import (
	"fmt"
	"sync"
	"time"
)

// RateLimiter - простой ограничитель частоты
type RateLimiter struct {
	interval time.Duration
	maxCount int
	count    int
	mu       sync.Mutex
}

func NewRateLimiter(interval time.Duration, maxCount int) *RateLimiter {
	rl := &RateLimiter{
		interval: interval,
		maxCount: maxCount,
	}

	// Автосброс счетчика
	go func() {
		for {
			time.Sleep(rl.interval)
			rl.mu.Lock()
			rl.count = 0
			rl.mu.Unlock()
		}
	}()

	return rl
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if rl.count < rl.maxCount {
		rl.count++
		return true
	}
	return false
}

func main() {
	// 5 операций в секунду
	limiter := NewRateLimiter(time.Second, 5)

	for i := 1; i <= 10; i++ {
		if limiter.Allow() {
			fmt.Printf("%s: Операция %d - разрешена\n", time.Now().Format("15:04:05"), i)
		} else {
			fmt.Printf("%s: Операция %d - ограничена!\n", time.Now().Format("15:04:05"), i)
		}
		time.Sleep(200 * time.Millisecond) // Искусственная задержка
	}
}
