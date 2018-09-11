include ../Makefile.properties

PROJECT_NAME=$$(basename $$(pwd))
DOCKERREPO=$(REPO_NAME)/$(PROJECT_NAME)

# docker image to use in production
IMAGE_NAME_PRODUCTION=alpine:3.8

# ensure this value equals the value in ./config.yaml
DEV_ENV_PORT=9090

init: build
	@printf -- '\033[1m\nAttempting to initialise dependencies...\033[0m\n' \
		&& $(MAKE) dep.init \
		|| printf -- '\033[1m\nAttempting to install dependencies...\033[0m\n' \
			&& $(MAKE) dep.ensure \

dep.add:
	@printf -- '\033[1m\nInstalling dependency `${DEP}`...\033[0m\n'
	docker run \
		-v "$$(pwd):/go/src/$(PROJECT_NAME)" \
		$(DOCKERREPO) \
		dep ensure -v -add ${DEP}

dep.ensure:
	docker run \
		-v "$$(pwd):/go/src/$(PROJECT_NAME)" \
		$(DOCKERREPO) \
		dep ensure -v

dep.init:
	docker run \
		-v "$$(pwd):/go/src/$(PROJECT_NAME)" \
		$(DOCKERREPO) \
		dep init

dev:
	@printf -- '\033[1m\nSpinning up with live-reloading...\033[0m\n'
	docker run \
		-p $(DEV_ENV_PORT):$(DEV_ENV_PORT) \
		-v "$$(pwd):/go/src/$(PROJECT_NAME)" \
		$(DOCKERREPO) \
		realize start --run main.go

prod: compile
	@printf -- '\033[1m\nRunning in production...\033[0m\n'
	$(eval PATH_BINARY=/app/$$(PROJECT_NAME))
	docker run \
		--env BASIC_SERVER_ENV=production \
		--env BASIC_SERVER_PORT=8080 \
		-v "$$(pwd)/$(PROJECT_NAME):$(PATH_BINARY)" \
		-it \
		--entrypoint="$(PATH_BINARY)" \
		$(IMAGE_NAME_PRODUCTION)

compile:
	@printf -- '\033[1m\nCreating binary...\033[0m\n'
	docker run \
		-v "$$(pwd):/go/src/$(PROJECT_NAME)" \
		$(DOCKERREPO) \
		go build

build:
	@printf -- '\033[1m\nBuilding base image...\033[0m\n'
	docker build \
		--build-arg PROJECT_NAME=$(PROJECT_NAME) \
		--tag $(DOCKERREPO) \
		.
