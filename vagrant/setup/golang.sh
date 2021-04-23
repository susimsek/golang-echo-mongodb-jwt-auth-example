#!/bin/bash

#!/bin/bash
export GO_VERSION=1.16.3

wget -c https://dl.google.com/go/go${GO_VERSION}.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local

sudo mkdir -p /vagrant/go

sudo chmod -R 777 /vagrant/go
# If GOPATH already set then DO Not set it again
if [ -z $GOPATH ]
then
    echo "export GOROOT=/usr/local/go" >> /home/vagrant/.profile
    echo "export GOPATH=/vagrant/go" >> /home/vagrant/.profile
    echo "export PATH=$PATH:/usr/local/go/bin" >> /home/vagrant/.profile
    echo "export GOROOT=/usr/local/go" >> /home/vagrant/.bashrc
    echo "export GOPATH=/vagrant/go" >> /home/vagrant/.bashrc
    echo "export PATH=$PATH:/usr/local/go/bin" >> /home/vagrant/.bashrc
else
    echo "======== No Change made to .profile, .bashrc  ====="
fi

echo "======= Done. PLEASE LOG OUT & LOG Back In ===="
echo "Then validate by executing    'go version'"