package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"sync"
)

type ConcurrentQueue struct {
	items []int
	mut   sync.Mutex
}

func (cq *ConcurrentQueue) enqueue(num int) {
	cq.mut.Lock()
	defer cq.mut.Unlock()
	cq.items = append(cq.items, num)
}

func (cq *ConcurrentQueue) dequeue() int {
	if len(cq.items) == 0 {
		panic("Queue is empty")
	}
	cq.mut.Lock()
	defer cq.mut.Unlock()
	item := cq.items[0]
	cq.items = cq.items[1:]
	return item
}

func (cq *ConcurrentQueue) len() int {
	return len(cq.items)
}

func main() {
	f, _ := os.Create("cpu.out")
	m, _ := os.Create("mem.out")
	pprof.StartCPUProfile(f)
	pprof.WriteHeapProfile(m)
	defer pprof.StopCPUProfile()
	performOperations()
}

func performOperations() {
	var MAX_ITEMS = 1000000

	var wg sync.WaitGroup
	queue := ConcurrentQueue{
		items: []int{},
	}
	for i := 1; i <= MAX_ITEMS; i++ {
		wg.Go(func() {
			queue.enqueue(i)
		})
	}
	wg.Wait()
	if queue.len() != MAX_ITEMS {
		panic("Incorrect queue length after enqueue")
	}

	for i := 1; i <= MAX_ITEMS; i++ {
		wg.Go(func() {
			queue.dequeue()
		})
	}
	wg.Wait()
	if queue.len() != 0 {
		panic("Incorrect queue length after dequeue")
	}

	fmt.Println("Operation Successfull")
}
