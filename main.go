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

		// print current dir and '$'
		path, _ := os.Getwd()
		fmt.Printf("%v$ ", path)

		// read from standard input
		buf := bufio.NewReader(os.Stdin)
		cmdInput, err := buf.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		cmdInput = strings.TrimSuffix(cmdInput, "\n")

		//Execute command
		exitFlag := execCommands(cmdInput)
		if exitFlag {
			break
		}
	}

}

func execCommands(cmdInput string) (ExitFlag bool) {

	// if cmdInput == "" || cmdInput == " " {
	// 	return false
	// }
	arrayCmd := strings.Fields(cmdInput)
	if len(arrayCmd) < 1 {
		return false
	}

	switch arrayCmd[0] {
	case "exit":
		return true
	case "help":
		help()
	case "cd":
		cd(arrayCmd)
	case "echo":
		echo(arrayCmd)
	default:
		commandExec := exec.Command(arrayCmd[0], arrayCmd[1:]...)
		commandExec.Stderr = os.Stderr
		commandExec.Stdout = os.Stdout
		err := commandExec.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
	return false
}
