package routain

import (
	"fmt"
	"sync"
)

type Result struct {
	Index int
	Value int
}

func computeSquare(idx, n int, ch chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	result := n * n
	ch <- Result{Index: idx, Value: result}
}

func Run() {
	numbers := []int{2, 4, 6, 8, 10}
	ch := make(chan Result, len(numbers))
	var wg sync.WaitGroup
	results := make([]int, len(numbers))

	for idx, n := range numbers {
		wg.Add(1)
		go computeSquare(idx, n, ch, &wg)
	}

	wg.Wait()
	close(ch)

	for r := range ch {
		results[r.Index] = r.Value
	}

	fmt.Println("입력 순서대로 결과:", results)
} 
