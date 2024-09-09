package main

var serviceRegistry = make(map[string]string) // Service name -> VPC

func registerService(serviceName, vpcName string) {
	serviceRegistry[serviceName] = vpcName
}

func discoverService(serviceName string) (string, bool) {
	vpc, exists := serviceRegistry[serviceName]
	return vpc, exists
}
