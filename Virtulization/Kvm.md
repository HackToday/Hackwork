KVM
=====

System Reqirements
------------------

Kvm requires some kind of HW processor support in virtulization. So before we proceed, we need to check if system
support it or not.


To check if you have kvm enable or default to qemu mode, can go through following ways

```
  1. the modules are correctly loaded lsmod|grep kvm
  2. you don't have a "KVM: disabled by BIOS" line in the output of dmesg
  3. /dev/kvm exists and you have the correct rights to use it
```

How to Manual Create VM  
-----------------------

1. You have networking configured, like bridge or ovs
2. If bridge, then use qemu tools to create disk images and boot VM

```
# create image like below, vmdisks/vm1 is my directory
sudo qemu-img create -f qcow2 vmdisks/vm1/mytestkvm.qcow2  35G

# Generate Mac addr for VM
printf 'DE:AD:BE:EF:%02X:%02X\n' $((RANDOM%256)) $((RANDOM%256))

```

# create qemu related script for network handy, we named it qemu-ifup, make sure make it executeable(+x)
```
#!/bin/sh
set -x

switch=br0

if [ -n "$1" ];then
        #tunctl -u `whoami` -t $1
        ip tuntap add $1 mode tap user `whoami`
        ip link set $1 up
        sleep 0.5s
        #brctl addif $switch $1
        ip link set $1 master $switch
        exit 0
else
        echo "Error: no interface specified"
        exit 1
fi
```

# start VM and install OS

```
sudo qemu-system-x86_64 -enable-kvm  -hda  vmdisks/vm1/mytestkvm.qcow2 -smp cpus=2 -m 3G -cdrom Downloads/ubuntu-16.04-server-amd64.iso  --device e1000,netdev=net0,mac=DE:AD:BE:EF:FB:E4 -netdev tap,id=net0,script=/home/test/qemu-ifup
```

After install complete, you can just start it without cdrom option like below:

```
sudo qemu-system-x86_64 -enable-kvm  -hda  vmdisks/vm1/mytestkvm.qcow2 -smp cpus=2 -m 3G   --device e1000,netdev=net0,mac=DE:AD:BE:EF:FB:E4 -netdev tap,id=net0,script=/home/test/qemu-ifup
```
