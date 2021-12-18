SQLMIGRATE_NAME=sqlmigrate:gomsx

up:
	docker-compose up -d app database

down:
	docker-compose down

run:
	docker-compose exec app go run app/cmd/main.go -c config/local.yml

mod:
	docker-compose exec app go mod tidy
	docker-compose exec app go mod vendor

fmt:
	docker-compose exec app go fmt ./...

build-tools:
	docker build --file ./tools/sqlmigrate/Dockerfile --tag $(SQLMIGRATE_NAME) .

lint:
	docker run --rm \
		--volume "$(PWD):/app" \
		-w /app golangci/golangci-lint:v1.43.0 \
		golangci-lint run

boil:
	docker-compose exec app sqlboiler mysql -c db/sqlboiler.toml

migrate-new:
	docker-compose run --rm migration new ${name}

local-migrate-%:
	$(eval CMD:= $*)
	docker-compose run --rm migration $(CMD) -env=mysql -config sqlmigrate.yml
