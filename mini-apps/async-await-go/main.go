package main

import (
	"fmt"
	"time"
)

type Future[T any] struct {
	await func() (T, error)
}

func (f *Future[T]) Await() (T, error) {
	return f.await()
}

func Async[T any](f func() (T, error)) *Future[T] {
	var result T
	var err error

	done := make(chan struct{})

	go func() {
		defer close(done)
		result, err = f()
	}()

	return &Future[T]{
		await: func() (T, error) {
			<-done
			return result, err
		},
	}
}

func fetchUserData() (string, error) {
	fmt.Println("=> Starting to fetch user data...")
	time.Sleep(2 * time.Second)
	fmt.Println("=> Finished fetching user data.")
	return "User data: Dev Pouya Gh", nil
}

func main() {
	futureUser := Async(fetchUserData)

	fmt.Println("Darhale daryafte user data...")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Darahale kar kardan...")

	fmt.Println("Sabr kardan baraye daryafte user")
	userData, err := futureUser.Await()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Success! %v\n", userData)
	}
}
