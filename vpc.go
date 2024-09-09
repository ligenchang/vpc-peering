package main

import (
	"sync"
)

type RouteTable struct {
	routes map[string]string // Destination VPC -> Peering connection
}

type VPC struct {
	name          string
	routeTable    *RouteTable
	securityGroup SecurityGroup
}

var vpcs = make(map[string]*VPC)
var mu sync.Mutex

func createVPC(name string) *VPC {
	return &VPC{
		name:          name,
		routeTable:    &RouteTable{routes: make(map[string]string)},
		securityGroup: SecurityGroup{rules: make(map[string]bool)},
	}
}

func (vpc *VPC) addRoute(destinationVPC string, peeringConnection string) {
	mu.Lock()
	defer mu.Unlock()
	vpc.routeTable.routes[destinationVPC] = peeringConnection
}

func (vpc *VPC) checkRoute(destinationVPC string) (bool, string) {
	mu.Lock()
	defer mu.Unlock()
	route, exists := vpc.routeTable.routes[destinationVPC]
	return exists, route
}
