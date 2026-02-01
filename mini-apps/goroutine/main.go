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

var (
	totalUpdates int
	updateMutex  sync.Mutex
)

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

func processOrders(orders []*Order) {
	for _, order := range orders {
		time.Sleep(time.Duration(rand.IntN(500)) * time.Millisecond)
		fmt.Printf("Processing order %d\n", order.ID)
	}
}

func updateOrderStatus(order *Order) {
	order.mu.Lock()

	time.Sleep(
		time.Duration(rand.IntN(300)) * time.Millisecond,
	)
	status := []string{"Processing", "Shipped", "Delivered"}[rand.IntN(3)]

	// WARNING: Data Race
	// Ok kardan ba Mutex
	order.Status = status
	fmt.Printf("Update order %d status: %s\n", order.ID, order.Status)

	order.mu.Unlock()

	updateMutex.Lock()
	defer updateMutex.Unlock()
	currentUpdates := totalUpdates
	time.Sleep(5 * time.Millisecond)
	totalUpdates = currentUpdates + 1
}

func reportOrderStatus(orders []*Order) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("\n--- Order Status Report ---")

		for _, order := range orders {
			fmt.Printf("Order %d: %s\n", order.ID, order.Status)
		}

		fmt.Println("--------------------")
		println()
	}
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

	orders := generateOrder(20)

	var wg sync.WaitGroup
	wg.Add(3)
	// go func() {
	// 	defer wg.Done()
	// 	processOrders(orders)
	// }()

	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			for _, order := range orders {
				updateOrderStatus(order)
			}
		}()
	}
	wg.Wait()

	reportOrderStatus(orders)

	fmt.Println("All operations completed. Exiting.")
	fmt.Printf("Total Updates: %d\n", totalUpdates)
}
