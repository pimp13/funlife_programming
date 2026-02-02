package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Status string
	mu     sync.Mutex
}

func generateOrder(count int) []*Order {
	orders := make([]*Order, count)
	for i := 0; i < count; i++ {
		orders[i] = &Order{
			ID:     i + 1,
			Status: "Pending",
		}
	}
	return orders
}

func processOrders(order <-chan *Order, wg *sync.WaitGroup) {
	defer wg.Done()
	// order <- &Order{ID: 3, Status: "hello"}
	// for _, order := range order {
	time.Sleep(time.Duration(rand.IntN(500)) * time.Millisecond)
	fmt.Printf("Processing order %d\n", order)
	// }
}

func main() {
	// No Goroutine: ./goroutine-example  0.00s user 0.01s system 0% cpu 14.183 total
	// Yes Goroutine: ./goroutine-example  0.00s user 0.01s system 0% cpu 5.926 total
	/*
		NOTES
		WARNING: Data Race
		Choon ke chand goroutine hamzaman daran royae yek data minevisan va mikhonan
		error va warning race condition bevoojood miyad!!

		baraye daryaft va test kardan warning race condition:
		go run -race main.go

		Rahe Hal:
		baraye ok kardan barnameh be tori ke race condition bartaraf beshe bayad az:
		Mutex estefadeh kard.

		Rahe Hal Pishrafte tar:
		estefadeh az channel

		Mutex => Sadeh sarii
		Channel => Amn tar va memari mehvar tar

	*/
	/*
		Channel:
		channel ha baraye tabadoole data beiyne goroutine ha estenfadeh mishe

		baraye sakhte channel az dastore make estefadeh mikonim:
		myChan := make(chan DataType)

		baraye value dadan be channel:
		myChan <- Data

		baraye darvafteh data value az channel:
		myData := <-myChan
	*/

	var wg sync.WaitGroup
	wg.Add(2)
	orderChan := make(chan *Order)

	go func() {
		defer wg.Done()
		for _, order := range generateOrder(20) {
			orderChan <- order
		}

		fmt.Println("Done with generating orders")
	}()

	go processOrders(orderChan, &wg)

	wg.Wait()

	fmt.Println("All operations completed. Exiting.")
}
