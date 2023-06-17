package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Simple shell in Go

func main() {

	buf := bufio.NewReader(os.Stdin)
	cmdInput, err := buf.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	cmdInput = strings.TrimSuffix(cmdInput, "\n")

	commandExec := exec.Command(cmdInput)
	commandExec.Stderr = os.Stderr
	commandExec.Stdout = os.Stdout
	err = commandExec.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

}
