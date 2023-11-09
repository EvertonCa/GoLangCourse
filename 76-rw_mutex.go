package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type safeCounter76 struct {
	counts map[string]int
	mux    *sync.RWMutex
}

func (sc safeCounter76) inc(key string) {
	sc.mux.Lock()
	defer sc.mux.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounter76) val(key string) int {
	sc.mux.RLock()
	defer sc.mux.RUnlock()
	return sc.counts[key]
}

// don't touch below this line

func (sc safeCounter76) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

type emailTest76 struct {
	email string
	count int
}

func test76(sc safeCounter76, emailTests []emailTest76) {
	emails := make(map[string]struct{})

	var wg sync.WaitGroup
	for _, emailT := range emailTests {
		emails[emailT.email] = struct{}{}
		for i := 0; i < emailT.count; i++ {
			wg.Add(1)
			go func(emailT emailTest76) {
				sc.inc(emailT.email)
				wg.Done()
			}(emailT)
		}
	}
	wg.Wait()

	emailsSorted := make([]string, 0, len(emails))
	for email := range emails {
		emailsSorted = append(emailsSorted, email)
	}
	sort.Strings(emailsSorted)

	sc.mux.RLock()
	defer sc.mux.RUnlock()
	for _, email := range emailsSorted {
		fmt.Printf("Email: %s has %d emails\n", email, sc.val(email))
	}
	fmt.Println("=====================================")
}

func main() {
	sc := safeCounter76{
		counts: make(map[string]int),
		mux:    &sync.RWMutex{},
	}
	test76(sc, []emailTest76{
		{
			email: "john@example.com",
			count: 23,
		},
		{
			email: "john@example.com",
			count: 29,
		},
		{
			email: "jill@example.com",
			count: 31,
		},
		{
			email: "jill@example.com",
			count: 67,
		},
	})
	test76(sc, []emailTest76{
		{
			email: "kaden@example.com",
			count: 23,
		},
		{
			email: "george@example.com",
			count: 126,
		},
		{
			email: "kaden@example.com",
			count: 31,
		},
		{
			email: "george@example.com",
			count: 453,
		},
	})
}
