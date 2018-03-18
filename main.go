package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	num := rand.Intn(2)
	if num == 0 {
		fmt.Println("Tails")
	} else {
		fmt.Println("Heads")
	}
}