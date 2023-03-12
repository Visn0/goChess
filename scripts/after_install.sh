#!/bin/bash

# Load Go env variables
source ~/.bash_go_config

# Navigate to project directory
cd $HOME/go/goChess

# Install project dependencies
make install_dependencies

# Run project compilation
make build
