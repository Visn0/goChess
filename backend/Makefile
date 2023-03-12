BUILD_DIR?=.
BINARY_NAME?=gochessbin

install_dependencies:
	go mod tidy	

up:
	go run main.go
	
fmt:
	go fmt ./...

lint:
	docker run --rm -v $(PWD)/backend:/app -w /app golangci/golangci-lint:v1.50.1 golangci-lint run

build:
	go build -o $(BINARY_NAME)		
	echo 'PORT=8081' > .env
	echo 'SERVE_SINGLE_PAGE_APP=true' >> .env	
	echo 'SINGLE_PAGE_APP_FOLDER=$(BUILD_DIR)/dist' >> .env
