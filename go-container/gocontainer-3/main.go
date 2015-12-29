package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func start(args []string) error {
	cmd := &exec.Cmd{
		Path: os.Args[0],
		Args: append([]string{"forkChild"}, args...),
	}
	if len(args) == 0 {
		log.Printf("Hit error now, missing args.")
		return errors.New("Args missing")
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWPID | syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC,
	}
	return cmd.Run()
}

func forkChild() error {
	log.Printf("Fork child start now")
	name, err := exec.LookPath(os.Args[1])
	if err != nil {
		return errors.New("LookPath failure")
	}
	if err = syscall.Sethostname([]byte("Testhost")); err != nil {
		return errors.New("Sethostname failure")
	}

	return syscall.Exec(name, os.Args[1:], os.Environ())
}

func main() {
	log.Printf("Args are %s", os.Args)
	if os.Args[0] == "forkChild" {
		if err := forkChild(); err != nil {
			log.Fatalf("Fork child error: %v", err)
		}
		os.Exit(0)
	}
	args := os.Args[1:]
	if len(args) == 0 {
		args = []string{os.Getenv("SHELL")}
	}
	log.Printf("The args are %s", args)
	start(args)
}
