package devices

import "fmt"

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
