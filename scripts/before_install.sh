#!/bin/bash

# Update system
apt-get -y update

########################################
#####        INSTALL NODEJS        #####
########################################
# Remove older nodejs version 
sudo apt remove nodejs npm

# Install new nodejs version
cd ~
curl -sL https://deb.nodesource.com/setup_18.x -o nodesource_setup.sh

chmod +x nodesource_setup.sh
bash nodesource_setup.sh
sudo apt-get -y install nodejs
sudo dpkg -i --force-overwrite /var/cache/apt/archives/nodejs_18.15.0-deb-1nodesource1_amd64.deb

########################################
#####          INSTALL GO          #####
########################################
GO_VERSION="1.19.7"
GO_ROOT="/usr/local/go"

# Remove older Go version 
rm -rf /usr/local/go

# Download Go files
wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz -o /tmp/go.tar.gz

# Decompress and put Go files in the installation path
tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz

# Remove Go downloaded file
rm -rf go${GO_VERSION}.linux-amd64.tar.gz

# Set config file so everytime a shell is opened, it sets the env GOROOT and PATH 
# variables with the root of Go installation, so "go" command can be used
echo "# Go settings"                                  >  ~/.bash_go_config
echo "export GOROOT=${GO_ROOT}"                       >> ~/.bash_go_config
echo "export GOPATH=/home/ubuntu/go"                       >> ~/.bash_go_config
echo "export PATH=${GO_ROOT}/bin:${PATH}"             >> ~/.bash_go_config

# In case the machine is restarted, we want this env variables to be loaded automatically
cp ~/.bash_go_config /etc/profile.d/bash_go_config.sh

# Install other dependencies
DEBIAN_FRONTED=noninteractive apt-get install -y pkg-config libssl-dev
