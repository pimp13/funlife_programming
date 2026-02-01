package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
** Atomic **
baraye jologori az race condition,

vaghti chand goroutine roye yek variable write va read anjam bedan bedoone estefadeh az
atomic ya mutex race condition mishe va moshkel be voojod miyad

* baraye logic haye pichideh tar va variable haye bishtar bayad az Mutex estefadeh beshe!!
* atomic baraye yek ya doo variable khoob hast
*/

func main() {
	var counter atomic.Int64
	var wg sync.WaitGroup

	for i := 0; i < 1_000_000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Add(1)
		}()
	}
	wg.Wait()
	fmt.Printf("counter is: %d\n", counter.Load())
}
