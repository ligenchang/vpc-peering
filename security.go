package main

type SecurityGroup struct {
	rules map[string]bool // Allowed services
}

func (sg *SecurityGroup) allow(service string) {
	sg.rules[service] = true
}

func (sg *SecurityGroup) isAllowed(service string) bool {
	allowed, exists := sg.rules[service]
	return exists && allowed
}
