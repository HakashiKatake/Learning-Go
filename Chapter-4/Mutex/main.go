package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) increment(wg *sync.WaitGroup) {
	defer wg.Done()
	defer p.mu.Unlock() // Ensure the mutex is unlocked when done
	p.mu.Lock()         // Decrement the WaitGroup counter when done
	p.views += 1

}

func main() {
	var wg sync.WaitGroup

	myPost := post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.increment(&wg)

	}

	wg.Wait()

	fmt.Println("Post views:", myPost.views)
}
