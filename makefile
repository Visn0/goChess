setup_nodejs:
	curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash -
	sudo apt-get install -y nodejs npm

install_project: setup_nodejs 
	npm --prefix frontend/ install

clean:
	# Remove the resulting files of transpiling TypeScript files
	rm -rf frontend/dist frontend/.parcel-cache

fmt-lint: frontend-fmt frontend-lint

frontend-fmt:
	npm --prefix frontend/ run fmt

frontend-lint:
	npm --prefix frontend/ run lint

up:
	npm --prefix frontend/ run dev

up.backend:
	cd backend; go run main.go