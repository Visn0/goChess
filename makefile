BUILD_DIR=build
BACKEND_BUILD_DIR=backend/build

install_project: 
	npm --prefix frontend/ install

clean:
	# Remove the resulting files of transpiling TypeScript files
	rm -rf frontend/dist frontend/.parcel-cache

fmt-lint: frontend-fmt frontend-lint

frontend-fmt:
	npm --prefix frontend/ run fmt

frontend-lint:
	npm --prefix frontend/ run type-check
	npm --prefix frontend/ run lint

up:
	npm --prefix frontend/ run dev

up.backend:
	cd backend; go run main.go

backend-fmt-lint:
	make backend-fmt backend-lint	
	
backend-fmt:
	cd backend; go fmt ./...

backend-lint: 
	cd backend; docker run --rm -v $(PWD)/backend:/app -w /app golangci/golangci-lint:v1.50.1 golangci-lint run

builddir: 
	@if [ -d $(BUILD_DIR) ]; then\
		rm -rf $(BUILD_DIR);\
	fi;
	mkdir $(BUILD_DIR);

build: build.frontend build.backend
	mv frontend/dist $(BUILD_DIR)/dist
	mv backend/backendexec $(BUILD_DIR)/.

build.frontend:
	npm --prefix frontend/ run build	

build.backend: builddir
	cd backend; go build -o backendexec	

.PHONY: fmt-lint frontend-fmt frontend-lint up up.backend backend-fmt-lint builddir build build.frontend build.backend ec2
