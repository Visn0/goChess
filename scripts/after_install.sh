# Load Go env variables
source ~/.bash_go_config

# Navigate to project directory
cd $HOME/go/goChess

# Install project dependencies
make install_project

# Run project compilation
make build
