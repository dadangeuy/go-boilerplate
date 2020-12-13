IMAGE = go-boilerplate
VERSION = $(shell git show -q --format=%H)
APPLICATIONS = $(shell ls application)

test:
	go test -v ./...

build-binary:
	@for application in $(APPLICATIONS); do \
  		if [ -f application/$$application/main.go ]; then \
			echo $$application: build binary started; \
			GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
			go build -o build/$$application/application application/$$application/main.go || exit 1; \
			echo $$application: build binary finished \(build/$$application/application\); \
		fi; \
	done

build-image:
	@for application in $(APPLICATIONS); do \
		if [ -f build/$$application/Dockerfile ]; then \
			echo $$application: build image started; \
  			docker build -q -t $(IMAGE)/$$application:$(VERSION) build/$$application || exit 1; \
  			echo $$application: build image finished \($(IMAGE)/$$application:$(VERSION)\); \
		fi; \
  	done

postgres-migrate:
	build/migrator/application migrate

postgres-rollback:
	build/migrator/application rollback
