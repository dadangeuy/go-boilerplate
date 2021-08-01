PROJECT = $(shell basename $(CURDIR))
BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
COMMIT = $(shell git show -q --format=%H)
APPLICATIONS = $(shell ls application)
DOCKER_REPOSITORY ?= 'docker.io/dadangeuy'

generate-mock:
	@for domain in $(shell ls domain); do \
		mockery --dir domain/$$domain/ --output domain/$$domain/mock/ --outpkg mock --all --quiet || exit 1; \
	done

test:
	go test ./...

coverage-report:
	go test -coverpkg=./... -coverprofile=coverage.out ./...

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
  			docker build --quiet --tag $(PROJECT)-$$application:$(COMMIT) build/$$application || exit 1; \
  			echo $$application: build image finished \($(PROJECT)-$$application:$(COMMIT)\); \
		fi; \
  	done

push-image:
	@for application in $(APPLICATIONS); do \
  		if [ -f build/$$application/Dockerfile ]; then \
			echo $$application: push image started; \
			docker tag $(PROJECT)-$$application:$(COMMIT) $(DOCKER_REPOSITORY)/$(PROJECT)-$$application:$(COMMIT) || exit 1; \
			docker tag $(PROJECT)-$$application:$(COMMIT) $(DOCKER_REPOSITORY)/$(PROJECT)-$$application:$(BRANCH) || exit 1; \
			docker push $(DOCKER_REPOSITORY)/$(PROJECT)-$$application:$(COMMIT) || exit 1; \
			docker push $(DOCKER_REPOSITORY)/$(PROJECT)-$$application:$(BRANCH) || exit 1; \
  			if [ $(BRANCH) = 'master' ] || [ $(BRANCH) = 'main' ]; then \
  			  	docker tag $(PROJECT)-$$application:$(COMMIT) $(DOCKER_REPOSITORY)/$(PROJECT)-$$application:latest || exit 1; \
  			  	docker push $(DOCKER_REPOSITORY)/$(PROJECT)-$$application:latest || exit 1; \
		  	fi; \
			echo $$application: push image finished \($(DOCKER_REPOSITORY)/$(PROJECT)-$$application:$(COMMIT)\); \
  		fi; \
	done

postgres-migrate:
	build/migrator/application migrate

postgres-rollback:
	build/migrator/application rollback
