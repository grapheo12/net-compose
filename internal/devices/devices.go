package devices

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
