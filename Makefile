HOSTNAME=local.com
NAMESPACE=ldcorentin
NAME=klayer
BINARY=terraform-provider-${NAME}
VERSION=1.0.0
OS_ARCH=darwin_amd64
DIST=~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
DOCKER_CMD = $(shell which docker)

# default: docker-build-install
default: build-install

init:
	go mod init terraform-provider-klayer
	go mod tidy
	
build:
	go build -o ${BINARY}

install: build
	mkdir -p ${DIST}
	mv ${BINARY} ${DIST}

build-install:
	$(MAKE) build
	$(MAKE) install

# docker-build-install:
# 	mkdir -p ${DIST}
# 	docker build --no-cache -t ${BINARY}:latest -f docker/Dockerfile . --build-arg BINARY=${BINARY}
# 	id=$$(docker create ${BINARY}) && \
# 	docker cp $$id:/app/${BINARY} ${DIST} && \
# 	docker rm -v $$id

go-test:
	TF_ACC=true go test -v klayers/*

terraform-test:
	cd ../terraform &&  \
	pwd && \
	terraform init && \
	terraform apply --auto-approve && \
	cd ../terraform && rm -rf terraform.tfstate .terraform .terraform.lock.hcl