package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func start(args []string) error {
	if len(args) == 0 {
		log.Printf("Hit error now, missing args.")
		return errors.New("Args missing")
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWPID | syscall.CLONE_NEWUTS,
	}
	return cmd.Run()
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		args = []string{os.Getenv("SHELL")}
	}
	log.Printf("The args are %s", args)
	start(args)
}
