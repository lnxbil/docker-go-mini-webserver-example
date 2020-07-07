NAME=lnxbil/docker-go-mini-webserver-example

build:
	@docker build --pull $(BUILD_ARGS) --tag $(NAME) .
	@make -s images

images:
	@docker images $(NAME)

rebuild:
	@make -s build BUILD_ARGS="$(BUILD_ARGS) --no-cache"

push:
	@docker push $(REPOSITORY)/$(NAME)

pull:
	@docker pull $(REPOSITORY)/$(NAME)

run:
	@docker run --interactive --tty --rm --publish 8081:8081 $(NAME) || true
