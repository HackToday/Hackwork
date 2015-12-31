As unprivileged user, the network namepsace need some priviledged operations(especially the
parent namespace for create veth pairs). So we need some work to make it work, the following
specify what need to do to run the example correctly

Allow unprivileged user run:
============================

cd $GOPATH/src
go install github.com/user/gocontainer-6/net
sudo chown root:root  $GOPATH/bin/net
sudo chmod u+s $GOPATH/bin/net

Install Run handler:
====================

go install github.com/user/gocontainer-6

Run it:
========

$GOPATH/bin/gocontainer-6
