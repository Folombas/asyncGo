package main

// Пример ratelim

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	iterationsNum = 6
	goroutinesNum = 5
	quotaLimit    = 2
)

func startWorker(in int, wg *sync.WaitGroup, quotaCh chan struct{}) 
	{
		quotaCh <- struct{}{} // ratelim.go, берём свободный слот
		defer wg.Done()
		for j := 0; j < iterationsNum; j++ {
			fmt.Printf(formatWork(in, j))

			// if j%2 == 0 {
			// <-quotaCh               // ratelim.go, возвращаем слот
			// quotaCh <- struct{}{}   // ratelim.go, берём свободный слот
			//}

			runtime.Gosched()   // Даём поработать другим горутинам
		}
		<-quotaCh // ratelim.go, возвращаем слот
	}

func main() {
	wg := &sync.WaitGroup{}
	quotaCh := make(chan struct{}, quotaLimit)  // ratelim.go
	for i := 0; i < goroutinesNum; i++ {
		wg.Add(1)
		go startWorker(i, wg, quotaCh)
	}
	time.Sleep(time.Microsecond)
	wg.Wait()
	
}
