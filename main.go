package main

import (
	"fmt"
	"net/http"
)

func peerHandler(w http.ResponseWriter, r *http.Request) {
	source := r.URL.Query().Get("source")
	destination := r.URL.Query().Get("destination")

	vpcA, existsA := vpcs[source]
	vpcB, existsB := vpcs[destination]

	if !existsA || !existsB {
		fmt.Fprintf(w, "Invalid VPCs\n")
		return
	}

	logRequest(source, destination)

	success := peerVPCs(vpcA, vpcB)
	if success {
		fmt.Fprintf(w, "Successfully peered %s to %s\n", source, destination)
	} else {
		fmt.Fprintf(w, "Failed to peer %s to %s\n", source, destination)
	}
}

func main() {
	vpcs["vpcA"] = createVPC("vpcA")
	vpcs["vpcB"] = createVPC("vpcB")

	// Add a route for peering
	vpcs["vpcA"].addRoute("vpcB", "peering-connection")
	vpcs["vpcB"].addRoute("vpcA", "peering-connection")

	// Allow HTTP traffic
	vpcs["vpcA"].securityGroup.allow("http")
	vpcs["vpcB"].securityGroup.allow("http")

	// Register some services for service discovery
	registerService("webApp", "vpcA")

	http.HandleFunc("/peer", peerHandler)

	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", nil)
}
