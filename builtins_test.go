package main

import (
	"os"
	"testing"
)

func TestCdHome(t *testing.T) {

	homeDir, _ := os.UserHomeDir()

	arrayStrings := []string{"cd"}

	cd(arrayStrings)

	newDir, _ := os.Getwd()

	if newDir != homeDir {
		t.Errorf("builtin cd not working")
	}
}

func TestExit(t *testing.T) {
	input := "exit"

	exitFlag := execCommands(input)

	if exitFlag != true {
		t.Errorf("builtin exit not working")
	}
}
