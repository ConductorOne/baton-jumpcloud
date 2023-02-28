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
	# podman run --rm -v \
	# 	"${PWD}/build/jcapi:/app/output" \
	# 	mcr.microsoft.com/openapi/kiota generate \
	# 	--openapi https://docs.jumpcloud.com/api/2.0/index.yaml \
	# 	--language Go -n jcapi
	# 	-o /app/output
	# podman run --rm -v \
	# 	"${PWD}/build/jcapi:/app/output" \
	# 	docker.io/swaggerapi/swagger-codegen-cli generate \
	# 	-i https://docs.jumpcloud.com/api/2.0/index.yaml \
	# 	-l go \
	# 	-o /app/output
	podman run --rm -v \
		"${PWD}/build/jcapi:/output" \
		docker.io/openapitools/openapi-generator-cli generate \
		-i https://docs.jumpcloud.com/api/2.0/index.yaml \
		-g go \
	    -o /output \
		--additional-properties=enumClassPrefix=true,hideGenerationTimestamp=true,structPrefix=true,disallowAdditionalPropertiesIfNotPresent=false,packageName=github.com/ConductorOne/baton-jumpcloud/pkg/jcapi,isGoSubmodule=true
	rm -rf pkg/jcapi
	mv build/jcapi pkg/jcapi

.PHONY: lint
lint:
	golangci-lint run
