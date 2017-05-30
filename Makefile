NAME=$(shell basename $(shell pwd) )

define DOCKERFILE_BUILD
FROM golang:onbuild
endef

define DOCKERFILE_RUN
FROM scratch
ADD $(NAME) /
ENTRYPOINT ["/$(NAME)"]
EXPOSE 8081
endef

export DOCKERFILE_BUILD
export DOCKERFILE_RUN

all: build-app build-app-container run

build-app:
	@echo "$$DOCKERFILE_BUILD" > Dockerfile.build
	@docker build -t $(NAME)-build -f Dockerfile.build .
	@docker run \
		--volume $(PWD):/data:rw \
		--env CGO_ENABLED=0 --env GOOS=linux \
		--interactive --tty --rm \
		$(NAME)-build \
		go build -a -installsuffix cgo -o /data/$(NAME) .
	@docker rmi $(NAME)-build
	@rm -f Dockerfile.build
	@strip $(NAME)

build-app-container:
	@echo "$$DOCKERFILE_RUN" > Dockerfile.run
	@docker build -t $(NAME)-app -f Dockerfile.run .
	@rm -f Dockerfile.run

run:
	@docker run --interactive --tty --rm --publish 8081:8081 $(NAME)-app || true

export:
	@docker create --hostname $(NAME) --name $(NAME) --publish 8081:8081 $(NAME)-app
	@docker export -o $(NAME).docker-export $(NAME)
	@docker rm $(NAME)
