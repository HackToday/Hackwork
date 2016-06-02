Ubuntu
=======

NetworkManager
-------------
Usually we want to disable that instead of uninstall it, for ubuntu it had one bug for that,

```
sudo stop network-manager
echo "manual" | sudo tee /etc/init/network-manager.override
```

Manual Bridge Setup
-------------------

Usually we want to permenat change bridge setup instead of reboot disapper, we can change interface file as below

```
# Use old eth0 config for br0, plus bridge stuff
iface br0 inet dhcp
    bridge_ports    eth0
    bridge_stp      off
    bridge_maxwait  0
    bridge_fd       0
```

And then reboot system, some online guide said just `/etc/init.d/networking restart`, it seems can not work correctly,
it may because of Network-Manager issue, but not confirmed it yet.


If you want to use manual ways,(just hack and check, not make it permenat, use this ways like or did it in script for handy.

```
ip addr flush dev eth0

# create the bridge
ip link add br0 type bridge

# add related physical device to that bridge 
ip link set eth0 master br0

# assign addr to bridge, and up it
ip addr add $addr dev br0
ip link set br0 up

# set default gateway
ip route add default via $gw dev br0

```

Reference
-------------
1. https://help.ubuntu.com/community/NetworkManager
2. http://www.linux-kvm.org/page/Networking
3. https://help.ubuntu.com/community/NetworkConnectionBridge
