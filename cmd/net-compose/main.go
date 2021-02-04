package main

import (
	"fmt"
	"net-compose/internal/devices"
	"net-compose/internal/namespaces"
	"net-compose/internal/parser"
	"net-compose/internal/routes"
	"net-compose/internal/shell"
	"os"
)

func usage() {
	fmt.Println("Usage: net-compose compose.yaml (up|down|dry-run|shell namespace)")
}

func main() {
	if len(os.Args) < 3 {
		usage()
		os.Exit(1)
	}

	fname := os.Args[1]
	cmd := os.Args[2]

	config := parser.Parse(fname)

	namespace, gotNS := config["namespaces"]
	if cmd == "shell" {
		if len(os.Args) < 4 {
			usage()
			os.Exit(1)
		}

		if !gotNS {
			fmt.Println("Unknown Namespace")
			os.Exit(1)
		}

		shell.HandleShell(namespace, os.Args[3])

		return
	}

	if gotNS {
		namespaces.HandleNamespace(namespace, cmd)
	}

	device, gotDevices := config["devices"]
	if gotDevices {
		devices.HandleDevices(device, cmd)
	}

	route, gotRoutes := config["routes"]
	if gotRoutes {
		routes.HandleRoutes(route, cmd)
	}
}
