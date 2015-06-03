package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	git := exec.Command("git", os.Args[1:]...)
	git.Stdin = os.Stdin
	git.Stdout = os.Stdout
	git.Stderr = os.Stderr

	err := git.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "sd failed to exec: %s", err)
	}
}
