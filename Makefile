DOCKER_PATH=docker-compose -f docker/docker-compose.yml
DOCKER_RUN=$(DOCKER_PATH) run --rm app

build:
	$(DOCKER_PATH) build
up:
	$(DOCKER_PATH) up -d
down:
	$(DOCKER_PATH) down
tidy:
	$(DOCKER_RUN) go mod tidy
wire-api:
	$(DOCKER_RUN) sh -c "cd server/ && wire"
grpc-gen:
	$(DOCKER_RUN) sh -c "sh tools/gen.sh"
sqlboiler:
	$(DOCKER_RUN) sh -c "cd ../../ && sqlboiler mysql -c database/sqlboiler/config.toml -o app/src/model/db -p model --add-global-variants --wipe"
local-run:
	cd app/src && go run cmd/api/main.go