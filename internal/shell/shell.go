package shell

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func HandleShell(namespace interface{}, ns string) {
	arr, ok := namespace.([]interface{})
	if !ok {
		fmt.Println("Wrong format")
		os.Exit(1)
	}
	found := false
	for _, name := range arr {
		if ns == name.(string) {
			found = true
		}
	}

	if !found {
		fmt.Println("Namespace not found")
		os.Exit(1)
	}

	var cmd string
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("[%s]> ", ns)
		cmd, _ = reader.ReadString('\n')
		cmd := strings.Trim(cmd, "\n")
		if cmd == "exit" {
			return
		}

		cmdArr := strings.Split(cmd, " ")

		path, _ := exec.LookPath("ip")

		command := &exec.Cmd{
			Path:   path,
			Args:   append([]string{path, "netns", "exec", ns}, cmdArr...),
			Stdin:  os.Stdin,
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		}

		command.Run()
	}

}
