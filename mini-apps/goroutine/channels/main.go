package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type UserService struct {
	ID        int
	Name      string
	CreatedAt time.Time
}

func makeUserService(count int) []*UserService {
	userServices := make([]*UserService, count)
	for i := 0; i < 20; i++ {
		userServices[i] = &UserService{
			ID:        i + 1,
			Name:      fmt.Sprintf("User Service - %d", i+1),
			CreatedAt: time.Now(),
		}
	}
	return userServices
}

func processInUserService(
	ctx context.Context,
	userServiceChan <-chan *UserService,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Processing canceled! in process")
			return
		case service, ok := <-userServiceChan:
			if !ok {
				fmt.Println("Channel close, Processing is end!")
				return
			}
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Processing in user service %s\n", service.Name)
		case <-time.After(1 * time.Second):
			fmt.Println("Timeout, wating for processing!")

		}
	}
}

func produce(
	ctx context.Context,
	userServiceChan chan<- *UserService,
	userServices []*UserService,
) {
	defer close(userServiceChan)

	for _, userService := range userServices {
		select {
		case <-ctx.Done():
			fmt.Println("Processing canceled! in product")
			return
		case userServiceChan <- userService:
			fmt.Printf("User Service added to channel: %d\n", userService.ID)

		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	services := makeUserService(20)
	userServiceChan := make(chan *UserService, 3)

	var wg sync.WaitGroup
	wg.Add(1)

	go produce(ctx, userServiceChan, services)
	go processInUserService(ctx, userServiceChan, &wg)

	wg.Wait()
}
