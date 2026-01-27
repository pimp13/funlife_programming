package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var random = rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(random.Float64())
}
