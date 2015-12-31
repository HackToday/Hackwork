package main

import (
	"errors"
	"github.com/vishvananda/netlink"
	"log"
	"os/exec"
	"strconv"
	"time"
)

func waitForIface() (netlink.Link, error) {
	log.Printf("Starting to wait for network interface")
	start := time.Now()
	for {
		if time.Since(start) > 5*time.Second {
			return nil, errors.New("Failed to find veth interfaces in 5s")
		}
		lst, err := netlink.LinkList()
		if err != nil {
			return nil, err
		}
		for _, l := range lst {
			if l.Type() == "veth" {
				return l, nil
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func setupIface(link netlink.Link) error {
	lo, err := netlink.LinkByName("lo")
	if err != nil {
		return err
	}
	log.Printf("Going to bring up interface loop")
	if err := netlink.LinkSetUp(lo); err != nil {
		return errors.New("loop interface up failed " + err.Error())
	}
	addr, err := netlink.ParseAddr("169.254.1.2/30")

	if err := netlink.LinkSetUp(link); err != nil {
		return errors.New("veth1 interface up failed " + err.Error())
	}
	return netlink.AddrAdd(link, addr)
}

func prepareIfacePair(pid int) error {
	cmd := exec.Command("net", strconv.Itoa(pid))
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New("veth pair command out: " + string(out) + " , err: " + err.Error())
	}
	return nil
}
