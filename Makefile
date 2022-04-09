SQLMIGRATE_NAME=sqlmigrate:gosample2022

up:
	docker-compose up -d

down:
	docker-compose down

run:
	docker-compose exec app go run cmd/srv/main.go

mod:
	docker-compose exec app go mod tidy
	docker-compose exec app go mod vendor

fmt:
	docker-compose exec app go fmt ./...

build-tools:
	docker build --file ./tools/sqlmigrate/Dockerfile --tag $(SQLMIGRATE_NAME) .

lint:
	@docker run --rm \
		--volume "$(PWD):/app" \
		-w /app golangci/golangci-lint:v1.43.0 \
		golangci-lint run

boil:
	docker-compose exec app sqlboiler mysql -c db/sqlboiler.toml --add-soft-deletes

protoc:
	docker-compose exec app protoc --go_out=./app/interface/grpc --go-grpc_out=./app/interface/grpc protos/${name}.proto

migrate-new:
	@docker run --rm \
	-w /sqlmigrate \
	-v "$(PWD):/sqlmigrate" \
	$(SQLMIGRATE_NAME) new -env=mysql -config db/sqlmigrate.yml ${name}

local-migrate-%:
	$(eval CMD:= $*)
	@docker run --rm \
	-w /sqlmigrate \
	-v "$(PWD):/sqlmigrate" \
	-e DB_USER=root -e DB_PASSWORD=root -e DB_HOST=docker.for.mac.localhost -e DB_PORT=3306 -e DB_NAME=gosample2022_development \
	$(SQLMIGRATE_NAME) \
	$(CMD) -env=mysql -config db/sqlmigrate.yml;
	@docker run --rm \
	-w /sqlmigrate \
	-v "$(PWD):/sqlmigrate" \
	-e DB_USER=root -e DB_PASSWORD=root -e DB_HOST=docker.for.mac.localhost -e DB_PORT=3306 -e DB_NAME="gosample2022_test" \
	$(SQLMIGRATE_NAME) \
	$(CMD) -env=mysql -config db/sqlmigrate.yml;

test-seed:
	docker-compose exec db mysql -uroot -proot gosample2022_test < ./testsupport/seed/seed.sql
