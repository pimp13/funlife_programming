package main

import (
	"fmt"
	"math/rand/v2"
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
		order.Status = status
		fmt.Printf("Update order %d status: %s\n", order.ID, order.Status)
	}
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
	orders := generateOrder(20)
	fmt.Println("All operations completed. Exiting.")

	processOrders(orders)

	updateOrderStatuses(orders)
}
