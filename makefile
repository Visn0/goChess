BUILD_DIR=$(PWD)/build

# Frontend variables
FRONTEND_DIR=frontend
MAKE_FRONTEND=$(MAKE) -C $(FRONTEND_DIR)

# Backend variables
BACKEND_BINARY_NAME=gochessbin
BACKEND_DIR=backend
MAKE_BACKEND=$(MAKE) -C $(BACKEND_DIR)

install_dependencies: 
	$(MAKE_FRONTEND) install_dependencies
	$(MAKE_BACKEND) install_dependencies

up.frontend:
	$(MAKE_FRONTEND) up

up.backend:
	$(MAKE_BACKEND) up

fmt:
	$(MAKE_FRONTEND) fmt
	$(MAKE_BACKEND) fmt

lint: 
	$(MAKE_FRONTEND) lint
	$(MAKE_BACKEND) lint

builddir: 
	@if [ -d $(BUILD_DIR) ]; then\
		rm -rf $(BUILD_DIR);\
	fi;
	mkdir $(BUILD_DIR);

build: clean builddir
	$(MAKE_FRONTEND) build
	$(MAKE_BACKEND) build BUILD_DIR=$(BUILD_DIR) BINARY_NAME=$(BACKEND_BINARY_NAME)
	mv $(FRONTEND_DIR)/dist $(BUILD_DIR)/dist
	mv $(BACKEND_DIR)/$(BACKEND_BINARY_NAME) $(BUILD_DIR)/.
	mv $(BACKEND_DIR)/.env $(BUILD_DIR)/.

deploy.local: 
	echo 'VITE_APP_API_HOST=localhost:8081' > $(FRONTEND_DIR)/.env.production.local
	$(MAKE) build
	cd $(BUILD_DIR); ./$(BACKEND_BINARY_NAME)

clean:
	rm -rf $(FRONTEND_DIR)/dist $(FRONTEND_DIR)/.parcel-cache $(BUILD_DIR)

.PHONY: install_project up.frontend up.backend fmt lint builddir build deploy.local clean
