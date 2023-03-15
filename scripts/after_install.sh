#!/bin/bash

# Load Go env variables
echo "=> Load Go env variables"
source /etc/profile.d/bash_go_config.sh

# Navigate to project directory
cd /home/ubuntu/go/goChess

# Install project dependencies
make install_dependencies

# Run project compilation
make build
