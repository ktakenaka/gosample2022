STATICKCHECK_NAME=statickcheck:gomsx
SQLMIGRATE_NAME=sqlmigrate:gomsx
SQLBOILER_NAME=sqlboiler:gomsx

boiler:
	docker-compose run --rm sqlboiler mysql --config /sqlboiler.toml

up:
	docker-compose up -d app database

run-server:
	docker-compose exec app go run app/cmd/main.go -c config/local.yml

mod:
	docker-compose exec app go mod tidy
	docker-compose exec app go mod vendor

fmt:
	docker-compose exec app go fmt ./...

build-tools:
	docker build --file ./tools/sqlboiler/Dockerfile --tag $(SQLBOILER_NAME) .
	docker build --file ./tools/sqlmigrate/Dockerfile --tag $(SQLMIGRATE_NAME) .
	docker build --file ./tools/staticcheck/Dockerfile --tag $(STATICKCHECK_NAME) .

lint:
	@docker run\
		--rm\
		--volume "$(PWD):/gomsx"\
		$(STATICKCHECK_NAME) ./...

migrate-new:
	docker-compose run --rm migration new ${name}

local-migrate-%:
	$(eval CMD:= $*)
	docker-compose run --rm migration $(CMD)

# "host.docker.internal"がなかなか動かなかったので、localではdocker-composeで行く
remote-migrate-%:
	$(eval CMD:= $*)
	@docker run\
		--rm\
		--volume "$(PWD)/db/migrations:/sqlmigrate/migrations"\
		--volume "$(PWD)/tools/sqlmigrate/dbconfig.yml:/sqlmigrate/dbconfig.yml"\
		--env DB_USER=${DB_USER}\
		--env DB_PASSWORD=${DB_PASSWORD}\
		--env DB_HOST=${DB_HOST}\
		--env DB_NAME=${DB_NAME}\
		--env DB_PORT=${DB_PORT}\
		$(SQLMIGRATE_NAME) $(CMD)
