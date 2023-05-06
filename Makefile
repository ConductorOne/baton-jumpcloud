GOOS = $(shell go env GOOS)
GOARCH = $(shell go env GOARCH)
BUILD_DIR = dist/${GOOS}_${GOARCH}
OUTPUT_PATH = ${BUILD_DIR}/baton-jumpcloud

.PHONY: build
build:
	rm -f ${OUTPUT_PATH}
	mkdir -p ${BUILD_DIR}
	go build -o ${OUTPUT_PATH} cmd/baton-jumpcloud/*.go

.PHONY: update-deps
update-deps:
	go get -d -u ./...
	go mod tidy -v
	go mod vendor

.PHONY: add-dep
add-dep:
	go mod tidy -v
	go mod vendor

.PHONY: build-jcapi
build-jcapi: build-jcapi1 build-jcapi2

.PHONY: build-jcapi2
build-jcapi2:
	rm -rf build/jcapi2
	mkdir -p build/jcapi2
	podman run --rm -v \
		"${PWD}/build/jcapi2:/output" \
		docker.io/openapitools/openapi-generator-cli generate \
		-i https://docs.jumpcloud.com/api/2.0/index.yaml \
		-g go \
	    -o /output \
		--additional-properties=enumClassPrefix=true,hideGenerationTimestamp=true,structPrefix=true,disallowAdditionalPropertiesIfNotPresent=false,packageName=jcapi2,isGoSubmodule=true
	rm -rf build/jcapi2/go.mod build/jcapi2/go.sum
	rm -rf pkg/jcapi2
	mv build/jcapi2 pkg/
	find pkg/jcapi2 \( -type d -name .git -prune \) -o -type f -print0 | xargs -0 sed -i 's/GIT_USER_ID\/GIT_REPO_ID/conductorone\/baton\-jumpcloud\/pkg/g'
	cd pkg/jcapi2 && go mod init github.com/conductorone/baton-jumpcloud/pkg/jcapi2 && go mod tidy -v

.PHONY: build-jcapi1
build-jcapi1:
	rm -rf build/jcapi1
	mkdir -p build/jcapi1
	podman run --rm -v \
		"${PWD}/build/jcapi1:/output" \
		docker.io/openapitools/openapi-generator-cli generate \
		-i https://docs.jumpcloud.com/api/1.0/index.yaml \
		-g go \
	    -o /output \
		--additional-properties=enumClassPrefix=true,hideGenerationTimestamp=true,structPrefix=true,disallowAdditionalPropertiesIfNotPresent=false,packageName=jcapi1,isGoSubmodule=true
	rm -rf build/jcapi1/go.mod build/jcapi1/go.sum
	rm -rf pkg/jcapi1
	mv build/jcapi1 pkg/
	find pkg/jcapi1 \( -type d -name .git -prune \) -o -type f -print0 | xargs -0 sed -i 's/GIT_USER_ID\/GIT_REPO_ID/conductorone\/baton\-jumpcloud\/pkg/g'
	cd pkg/jcapi1 && go mod init github.com/conductorone/baton-jumpcloud/pkg/jcapi1 && go mod tidy -v

.PHONY: lint
lint:
	golangci-lint run
