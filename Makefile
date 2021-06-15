PROTOGEN_IMAGE = indrasaputra/protogen:0.0.1

.PHONY: format
format:
	bin/format.sh

.PHONY: genproto
genproto:
	bin/generate-proto.sh

.PHONY: genprotodocker
genprotodocker:
	docker run -it --rm \
    --mount "type=bind,source=$(PWD),destination=/work" \
    --mount "type=volume,source=orvosi-go-mod-cache,destination=/go,consistency=cached" \
    --mount "type=volume,source=orvosi-buf-cache,destination=/home/.cache,consistency=cached" \
    -w /work $(PROTOGEN_IMAGE) make -e -f Makefile genproto pretty

.PHONY: check-import
check-import:
	bin/check-import.sh

.PHONY: mockgen
mockgen:
	bin/generate-mock.sh

.PHONY: cleanlintcache
cleanlintcache:
	golangci-lint cache clean

.PHONY: lint
lint: cleanlintcache
	buf lint
	golangci-lint run ./...

.PHONY: tidy
tidy:
	GO111MODULE=on go mod tidy

.PHONY: pretty
pretty: tidy format lint

.PHONY: compile
compile:
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o orvosi cmd/server/main.go

.PHONY: cover
cover:
	go test -v -race ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	go tool cover -func coverage.out

.PHONY: coverhtml
coverhtml:
	go test -v -race ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

.PHONY: cleantestcache
cleantestcache:
	go clean -testcache

.PHONY: test.unit
test.unit: cleantestcache
	go test -v -race ./...
