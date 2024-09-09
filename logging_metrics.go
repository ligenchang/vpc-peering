package main

import (
	"log"
	"time"
)

func logRequest(vpcSource, vpcDestination string) {
	log.Printf("Request from %s to %s at %s", vpcSource, vpcDestination, time.Now().Format(time.RFC3339))
}
