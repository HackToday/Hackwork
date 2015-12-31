package main

import (
	"errors"
	"github.com/vishvananda/netlink"
	"log"
	"os"
	"strconv"
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

func main() {
	pid := 1
	if len(os.Args) > 1 {
		p, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		pid = p
	}

	if err := createVethPair(pid); err != nil {
		log.Fatal(err)
	}
}
