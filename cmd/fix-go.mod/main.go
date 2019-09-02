package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	cmdline := []string{"go", "mod", "edit"}
	for _, ver := range versions {
		cmdline = append(cmdline, "-require="+ver)
	}
	exe, err := exec.LookPath(cmdline[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	fmt.Println("running:", cmdline, "...")
	err = syscall.Exec(exe, cmdline, os.Environ())
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
