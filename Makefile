test:
	go test -v ./...

build-binary:
	go build -o build/web/application application/web/main.go
	go build -o build/migrator/application application/migrator/main.go
