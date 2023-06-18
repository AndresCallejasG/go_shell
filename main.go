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

	for {

		fmt.Print("$ ")
		buf := bufio.NewReader(os.Stdin)
		cmdInput, err := buf.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		cmdInput = strings.TrimSuffix(cmdInput, "\n")

		if cmdInput == "exit" {
			break
		}

		arrayCmd := strings.Fields(cmdInput)
		commandExec := exec.Command(arrayCmd[0], arrayCmd[1:]...)
		commandExec.Stderr = os.Stderr
		commandExec.Stdout = os.Stdout
		err = commandExec.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

}
