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

build-jcapi:
	rm -rf build/jcapi
	mkdir -p build/jcapi
	podman run --rm -v \
		"${PWD}/build/jcapi:/output" \
		docker.io/openapitools/openapi-generator-cli generate \
		-i https://docs.jumpcloud.com/api/2.0/index.yaml \
		-g go \
	    -o /output \
		--additional-properties=enumClassPrefix=true,hideGenerationTimestamp=true,structPrefix=true,disallowAdditionalPropertiesIfNotPresent=false,packageName=jcapi,isGoSubmodule=true
	rm -rf build/jcapi/go.mod build/jcapi/go.sum
	rm -rf pkg/jcapi
	mv build/jcapi pkg/jcapi
	find pkg/jcapi \( -type d -name .git -prune \) -o -type f -print0 | xargs -0 sed -i 's/GIT_USER_ID\/GIT_REPO_ID/ConductorOne\/baton\-jumpcloud\/pkg/g'
	cd pkg/jcapi && go mod init github.com/ConductorOne/baton-jumpcloud/pkg/jcapi && go mod tidy -v

	

.PHONY: lint
lint:
	golangci-lint run
