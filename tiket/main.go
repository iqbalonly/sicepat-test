package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	cashier := initConcertTicket(1000)
	var wg sync.WaitGroup

	for i := 1; i <= 100000; i++ {
		wg.Add(1)

		go func(cashier *ConcertTicket, index int) {
			defer wg.Done()
			qty := rand.Intn(3-1) + 1
			cashier.Buy(fmt.Sprintf("user %d", index), qty)

		}(cashier, i)
	}

	wg.Wait()
}

type ConcertTicket struct {
	sync.Mutex
	Total int
}

func initConcertTicket(n int) *ConcertTicket {
	return &ConcertTicket{
		Total: n,
	}
}

func (t *ConcertTicket) Buy(user string, qty int) {
	t.Mutex.Lock()
	if t.Total > 0 && t.Total-qty >= 0 {
		t.Total -= qty
		fmt.Printf("user %s berhasil beli tiket sebanyak %d, sisa tiket tersedia %d\n", user, qty, t.Total)

	} else {
		fmt.Printf("user %s tidak dapat membeli tiket sebanyak %d, sisa tiket tersedia %d\n", user, qty, t.Total)
	}

	t.Mutex.Unlock()
}
