package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

// Parse the Yaml file given by fname
func Parse(fname string) map[interface{}]interface{} {
	config := make(map[interface{}]interface{})
	dat, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(dat, &config)
	if err != nil {
		panic(err)
	}

	return config
}

func CreateNamespace(ns string) {

}

func RemoveNamespace(ns string) {
	fmt.Printf("ip netns delete %s\n", ns)
}

func PrintNamespace(ns string) {
	fmt.Printf("ip netns add %s\n", ns)
	fmt.Printf("ip netns exec %s sysctl -w net.ipv4.ip_forward=1\n", ns)
}

func HandleNamespace(namespaces interface{}, cmd string) {
	nss, ok := namespaces.([]interface{})
	if !ok {
		panic("Wrong format")
	}
	for _, ns := range nss {
		if cmd == "up" {
			CreateNamespace(ns.(string))
		} else if cmd == "down" {
			RemoveNamespace(ns.(string))
		} else if cmd == "dry-run" {
			PrintNamespace(ns.(string))
		}
	}
}

func CreateVethPair(veth string, peer string) {

}

func RemoveVethPair(veth string) {

}

func PrintVethPair(veth string, peer string) {
	fmt.Printf("ip link add %s type veth peer name %s\n", veth, peer)
}

func AddVethToNamespace(veth string, ns string) {

}

func PrintVethToNamespace(veth string, ns string) {
	fmt.Printf("ip link set %s netns %s\n", veth, ns)
	fmt.Printf("ip netns exec %s ip link set %s up\n", ns, veth)
}

func AddIpToVeth(veth string, ip string, ns string) {

}

func PrintIpToVeth(veth string, ip string, ns string) {
	if ns != "default" {
		fmt.Printf("ip netns exec %s ip addr add %s dev %s\n", ns, ip, veth)
	} else {
		fmt.Printf("ip addr add %s dev %s\n", ip, veth)
	}
}

func HandleVeths(veth interface{}, cmd string) {
	devs, ok := veth.([]interface{})
	if !ok {
		panic("Wrong format")
	}

	for _, dev := range devs {
		dev := dev.(map[interface{}]interface{})
		for name, props := range dev {
			props := props.(map[interface{}]interface{})
			peer, gotPeer := props["peer"]
			if gotPeer {
				if cmd == "up" {
					CreateVethPair(name.(string), peer.(string))
				} else if cmd == "dry-run" {
					PrintVethPair(name.(string), peer.(string))
				}
			}

			ns, gotNS := props["namespace"]
			if gotNS {
				if cmd == "up" {
					AddVethToNamespace(name.(string), ns.(string))
				} else if cmd == "dry-run" {
					PrintVethToNamespace(name.(string), ns.(string))
				}
			}

			ip, gotIP := props["ip"]
			if gotIP {
				if cmd == "up" {
					if gotNS {
						AddIpToVeth(name.(string), ip.(string), ns.(string))
					} else {
						AddIpToVeth(name.(string), ip.(string), "default")
					}
				} else if cmd == "dry-run" {
					if gotNS {
						PrintIpToVeth(name.(string), ip.(string), ns.(string))
					} else {
						PrintIpToVeth(name.(string), ip.(string), "default")
					}
				}
			}
		}
	}
}

func HandleDevices(devices interface{}, cmd string) {
	// Only Veths supported for now
	devs, ok := devices.(map[interface{}]interface{})
	if !ok {
		return
	}
	veths, gotVeth := devs["veth"]
	if gotVeth {
		HandleVeths(veths, cmd)
	}
}

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

func main() {
	fname := os.Args[1]
	cmd := os.Args[2]
	config := Parse(fname)

	namespaces, gotNS := config["namespaces"]
	if gotNS {
		HandleNamespace(namespaces, cmd)
	}

	devices, gotDevices := config["devices"]
	if gotDevices {
		HandleDevices(devices, cmd)
	}

	routes, gotRoutes := config["routes"]
	if gotRoutes {
		HandleRoutes(routes, cmd)
	}
}
