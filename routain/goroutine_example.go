package routain

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Printf("[고루틴] %d\n", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func RunWg() {
	var wg sync.WaitGroup
	wg.Add(1)

	go printNumbers(&wg) // 고루틴으로 실행

	for i := 1; i <= 5; i++ {
		fmt.Printf("[메인] %d\n", i)
		time.Sleep(700 * time.Millisecond)
	}

	wg.Wait() // 고루틴이 끝날 때까지 대기
} 
