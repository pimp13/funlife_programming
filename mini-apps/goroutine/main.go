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

func processOrders(orders []*Order) {
	for _, order := range orders {
		time.Sleep(time.Duration(rand.IntN(500)) * time.Millisecond)
		fmt.Printf("Processing order %d\n", order.ID)
	}
}

func updateOrderStatuses(orders []*Order) {
	for _, order := range orders {
		time.Sleep(
			time.Duration(rand.IntN(300)) * time.Millisecond,
		)
		status := []string{"Processing", "Shipped", "Delivered"}[rand.IntN(3)]

		// WARNING: Data Race
		order.Status = status

		fmt.Printf("Update order %d status: %s\n", order.ID, order.Status)
	}
}

func reportOrderStatus(orders []*Order) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("\n--- Order Status Report ---")

		for _, order := range orders {
			// WARNING: Data Race
			fmt.Printf("Order %d: %s\n", order.ID, order.Status)
		}

		fmt.Println("--------------------")
		println()
	}
}

func main() {
	/*
		WARNING: Data Race
		Choon ke chand goroutine hamzaman daran royae yek data minevisan va mikhonan
		error va warning race condition bevoojood miyad!!

		baraye daryaft va test kardan warning race condition:
		go run -race main.go

		Rahe Hal:
		baraye ok kardan barname be tori ke race condition bartaraf beshe bayad az:
		Mutex estefadeh kard.

		Rahe Hal Pishrafte tar:
		estefadeh az channel

		Mutex => Sadeh sarii
		Channel => Amn tar va memari mehvar tar
	*/
	// No Goroutine: ./goroutine-example  0.00s user 0.01s system 0% cpu 14.183 total
	// Yes Goroutine: ./goroutine-example  0.00s user 0.01s system 0% cpu 5.926 total

	orders := generateOrder(20)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		processOrders(orders)
	}()

	go func() {
		defer wg.Done()
		updateOrderStatuses(orders)
	}()

	go func() {
		defer wg.Done()
		reportOrderStatus(orders)
	}()

	wg.Wait()

	fmt.Println("All operations completed. Exiting.")
}
