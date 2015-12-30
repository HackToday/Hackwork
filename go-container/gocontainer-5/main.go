package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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
		Cloneflags: syscall.CLONE_NEWPID | syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWNS | syscall.CLONE_NEWNET,
	}
	if err := cmd.Start(); err != nil {
		return err
	}
	log.Printf("container PID: %d", cmd.Process.Pid)
	if err := createVethPair(cmd.Process.Pid); err != nil {
		return err
	}
	return cmd.Wait()
}

type Mount struct {
	Source string
	Target string
	Fs     string
	Flags  int
	Data   string
}

type mountCfg struct {
	Mounts []Mount
	Rootfs string
}

var mountInfo = mountCfg{
	Mounts: []Mount{
		{
			Source: "proc",
			Target: "/proc",
			Fs:     "proc",
			Flags:  syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV,
		},
		{
			Source: "tmpfs",
			Target: "/dev",
			Fs:     "tmpfs",
			Flags:  syscall.MS_NOSUID | syscall.MS_STRICTATIME,
			Data:   "mode=755",
		},
	},
	Rootfs: "/home/ubuntu/busybox",
}

func mount(mntCfg mountCfg) error {
	for _, m := range mntCfg.Mounts {
		target := filepath.Join(mountInfo.Rootfs, m.Target)
		log.Printf("Mount %s to %s", m.Source, target)
		if err := syscall.Mount(m.Source, target, m.Fs, uintptr(m.Flags), m.Data); err != nil {
			return errors.New("Failed to mount " + m.Source)
		}
	}
	return nil
}

func pivotRoot(mntCfg mountCfg) error {
	root := mntCfg.Rootfs
	if err := syscall.Mount(root, root, "bind", syscall.MS_BIND|syscall.MS_REC, ""); err != nil {
		return errors.New("Mount rootfs to itself error: " + err.Error())
	}

	pivotDir := filepath.Join(root, ".pivot_root")
	if err := os.Mkdir(pivotDir, 0777); err != nil {
		return err
	}

	log.Printf("Pivot root dir: %s", pivotDir)
	log.Printf("Pivot root to %s", root)

	if err := syscall.PivotRoot(root, pivotDir); err != nil {
		return errors.New("pivot_root " + err.Error())
	}
	if err := syscall.Chdir("/"); err != nil {
		return errors.New("chdir / " + err.Error())
	}

	pivotDir = filepath.Join("/", ".pivot_root")

	if err := syscall.Unmount(pivotDir, syscall.MNT_DETACH); err != nil {

	}
	return os.Remove(pivotDir)
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
	if err := mount(mountInfo); err != nil {
		return err
	}
	if err := pivotRoot(mountInfo); err != nil {
		return errors.New("Pivot root error: " + err.Error())
	}
	link, err := waitForIface()
	if err := setupIface(link); err != nil {
		return err
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
