package main

import (
	"os"
	"os/exec"
)

var cmdAlias map[string][]string

func init() {
	cmdAlias = make(map[string][]string, 0)
	cmdAlias["co"] = []string{"checkout"}
	cmdAlias["bl"] = []string{"Branch", "--list"}
	cmdAlias["pr"] = []string{"merge", "--no-ff"}
	//cmdAlias["dw"] = ["pull", "origin", "[branch]"]
	//cmdAlias["up"] = ["push", "origin", "[branch]"]
	cmdAlias["rev"] = []string{"reset", "--hard"}
	cmdAlias["hist"] = []string{"log", "-10"}
	cmdAlias["fhist"] = []string{"log", "--follow"}
	cmdAlias["sub"] = []string{"reset"}

}

func main() {
	git := exec.Command("git")
	git.Stdin = os.Stdin
	git.Stdout = os.Stdout
	git.Stderr = os.Stderr

	if len(os.Args) > 1 {

		cmd := os.Args[1]
		alias := cmdAlias[cmd]
		if len(alias) > 0 {
			git.Args = append(git.Args, alias...)
			git.Args = append(git.Args, os.Args[2:]...)
		} else {
			git.Args = append(git.Args, os.Args[1:]...)
		}

	} else {
		git.Args = append(git.Args, "status")
	}
	git.Run()

}
