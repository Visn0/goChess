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
	mv backend/.env $(BUILD_DIR)/.

build.frontend:
	npm --prefix frontend/ run build	

build.backend: builddir
	cd backend; go build -o backendexec		
	echo 'PORT=8081' > backend/.env
	echo 'SERVE_SINGLE_PAGE_APP=true' >> backend/.env	

deploy.local: 
	echo 'VITE_APP_API_HOST=localhost:8081' > frontend/.env.production.local
	make build
	cd $(BUILD_DIR); ./backendexec

deploy.aws: 
	rm -rf frontend/.env.production.local
	make build 
	scp -i "~/.ssh/chess-carlos-keypair.pem" -r ./build ubuntu@ec2-35-180-164-238.eu-west-3.compute.amazonaws.com:.

.PHONY: fmt-lint frontend-fmt frontend-lint up up.backend backend-fmt-lint builddir build build.frontend build.backend deploy.local
