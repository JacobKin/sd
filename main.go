package main

import (
	"os"
	"os/exec"
)

var cmdAlias map[string]string

func init() {
	cmdAlias = make(map[string]string, 0)
	cmdAlias["co"] = "checkout"
}

func main() {
	git := exec.Command("git")
	git.Stdin = os.Stdin
	git.Stdout = os.Stdout
	git.Stderr = os.Stderr

	if len(os.Args) > 1 {

		cmd := os.Args[1]
		alias := cmdAlias[cmd]
		if alias != "" {
			git.Args = append(git.Args, alias)
			git.Args = append(git.Args, os.Args[2:]...)
		} else {
			git.Args = append(git.Args, os.Args[1:]...)
		}

	} else {
		git.Args = append(git.Args, "status")
	}
	git.Run()

}
