IMAGE = dadangeuy/go-boilerplate
VERSION = $(shell git show -q --format=%H)

test:
	go test -v ./...

build-binary:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/web/application application/web/main.go
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/migrator/application application/migrator/main.go

build-image:
	docker build -t $(IMAGE)/web:$(VERSION) build/web
