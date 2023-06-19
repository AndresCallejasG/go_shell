package main

import (
	"fmt"
	"os"
	"strings"
)

func help() {
	fmt.Println("Help")
	fmt.Println("\t Available Builtins: help, cd, echo and exit")
}

func cd(arrayStrings []string) {

	if len(arrayStrings) > 1 {
		err := os.Chdir(arrayStrings[1])
		if err != nil {
			fmt.Println(err)
		}
	} else {
		homedir, _ := os.UserHomeDir()
		err := os.Chdir(homedir)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func echo(arrayStrings []string) {
	fmt.Println(strings.Join(arrayStrings[1:], " "))
}
