package routes

import (
	"fmt"
	"strings"
)

func AddRoutes(ns string, dev string, ips []interface{}) {

}

func PrintRoutes(ns string, dev string, ips []interface{}) {
	for _, ip := range ips {
		ip := ip.(string)
		if !strings.Contains(ip, ":") {
			fmt.Printf("ip netns exec %s ip route add %s dev %s\n", ns, ip, dev)
		} else {
			// ip address:gateway address
			parts := strings.Split(ip, ":")
			fmt.Printf("ip netns exec %s ip route add %s via %s dev %s\n", ns, parts[0], parts[1], dev)
		}
	}
}

func HandleRoutes(routes interface{}, cmd string) {
	// <ns>: <dev>: [array of routes]
	r, ok := routes.(map[interface{}]interface{})
	if !ok {
		panic("Wrong format")
	}
	for ns, devs := range r {
		ns := ns.(string)
		devmap := devs.(map[interface{}]interface{})
		for dev, ips := range devmap {
			dev := dev.(string)
			ips := ips.([]interface{})
			if cmd == "up" {
				AddRoutes(ns, dev, ips)
			} else if cmd == "dry-run" {
				PrintRoutes(ns, dev, ips)
			}
		}

	}
}
