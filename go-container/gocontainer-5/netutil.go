package main

import (
	"errors"
	"github.com/vishvananda/netlink"
	"log"
	"time"
)

func createVethPair(pid int) error {
	parentName := "veth0"
	peerName := "veth1"
	la := netlink.NewLinkAttrs()
	la.Name = parentName

	vp := &netlink.Veth{LinkAttrs: la, PeerName: peerName}
	if err := netlink.LinkAdd(vp); err != nil {
		return errors.New("veth pair creation failed with " + err.Error())
	}

	peer, err := netlink.LinkByName(peerName)
	if err != nil {
		return errors.New("Get peer interface failed with " + err.Error())
	}

	if err := netlink.LinkSetNsPid(peer, pid); err != nil {
		return errors.New("Move peer to ns failed with " + err.Error())
	}

	log.Printf("Going to bring up interface veth0")
	if err := netlink.LinkSetUp(vp); err != nil {
		return errors.New("Parent link up failed with " + err.Error())
	}

	addr, err := netlink.ParseAddr("169.254.1.1/30")
	if err := netlink.AddrAdd(vp, addr); err != nil {
		return err
	}
	return nil
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
