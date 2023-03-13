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
GO_VERSION="1.19"
GO_ROOT="/usr/local/${GO_VERSION}"

# Remove older Go versions and installations
rm -rf /usr/local/go1.12
rm -rf /usr/local/go1.13
rm -rf /usr/local/go1.14
rm -rf /usr/local/go1.15
rm -rf /usr/local/go1.16
rm -rf /usr/local/go1.17
rm -rf /usr/local/go1.18
rm -rf /usr/local/go1.19

# Download Go files
wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz -o "/tmp/go.tar.gz"

# Decompress and put Go files in the installation path
tar -C /usr/local -xzf "/tmp/go.tar.gz"
mv /usr/local/go "${GO_ROOT}"

# Set config file so everytime a shell is opened, it sets the env GOROOT and PATH 
# variables with the root of Go installation, so "go" command can be used
echo "# Go settings"                                  >  ~/.bash_go_config
echo "export GOROOT=${GO_ROOT}"                       >> ~/.bash_go_config
echo "export GOPATH=${HOME}/go"                       >> ~/.bash_go_config
echo "export PATH=${GOROOT}/bin:$GOPATH/bin:${PATH}"  >> ~/.bash_go_config

# In case the machine is restarted, we want this env variables to be loaded automatically
cp ~/.bash_go_config /etc/profile.d/bash_go_config.sh

# Install other dependencies
DEBIAN_FRONTED=noninteractive apt-get install -y pkg-config libssl-dev
