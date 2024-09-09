package main

import (
	"fmt"
)

func peerVPCs(vpcA, vpcB *VPC) bool {
	// Check if there's a route from VPC A to VPC B
	routeExists, _ := vpcA.checkRoute(vpcB.name)
	if !routeExists {
		fmt.Printf("Failed to reach %s from %s\n", vpcB.name, vpcA.name)
		return false
	}

	// Security check
	if !vpcA.securityGroup.isAllowed("http") || !vpcB.securityGroup.isAllowed("http") {
		fmt.Printf("Traffic blocked by security group between %s and %s\n", vpcA.name, vpcB.name)
		return false
	}

	// Simulate latency
	simulateLatency()

	fmt.Printf("Successfully reached %s from %s\n", vpcB.name, vpcA.name)
	return true
}
