package main

import (
	"fmt"
	"math/rand"
	"time"
)

func simulateLatency() {
	latency := time.Duration(rand.Intn(500)) * time.Millisecond
	fmt.Printf("Simulated latency: %v\n", latency)
	time.Sleep(latency)
}
