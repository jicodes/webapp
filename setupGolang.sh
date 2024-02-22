#!/bin/bash

set -e 

sudo dnf install -y wget 

wget https://go.dev/dl/go1.21.6.linux-amd64.tar.gz

sudo tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz

rm go1.21.6.linux-amd64.tar.gz

echo "export PATH=$PATH:/usr/local/go/bin" | sudo tee -a /etc/profile

source /etc/profile

go version