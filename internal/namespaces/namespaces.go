package namespaces

import "fmt"

func CreateNamespace(ns string) {

}

func RemoveNamespace(ns string) {
	fmt.Printf("ip netns delete %s\n", ns)
}

func PrintNamespace(ns string) {
	fmt.Printf("ip netns add %s\n", ns)
	fmt.Printf("ip netns exec %s sysctl -w net.ipv4.ip_forward=1\n", ns)
	fmt.Printf("ip netns exec %s ip link set lo up\n", ns)
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
