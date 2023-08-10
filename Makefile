image := sultaniman/confetti
tag := latest
artifact := confetti
git_hash ?= $(shell git log -1 --pretty=format:%h)

.PHONY: build
build:
	docker build . \
	-t $(image):$(git_hash) \
	-t $(image):$(tag)

.PHONY: compile
compile:
	go build -o $(artifact)

.PHONY: push
push:
	docker image push $(image):$(git_hash)
	docker image push $(image):$(tag)

.PHONY: test
test:
	go test -v ./...

.PHONY: fmt
fmt:
	gofmt -w .

.PHONY: docker-test
docker-test:
	docker build --target builder .

swagger:
	swag init --parseDependency --parseInternal --parseDepth 2 --generalInfo cmd/serve.go
