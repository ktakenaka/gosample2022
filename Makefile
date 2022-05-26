SQLMIGRATE_NAME=sqlmigrate:gosample2022

up:
	docker-compose up -d

down:
	docker-compose down

run:
	ENV=local go run cmd/srv/main.go

cnsmrsample:
	go run cmd/cnsmrsample/main.go

mod:
	go mod tidy
	go mod vendor

fmt:
	go fmt ./...

build-tools:
	docker build --file ./tools/sqlmigrate/Dockerfile --tag $(SQLMIGRATE_NAME) .

lint:
	@docker run --rm \
		--volume "$(PWD):/app" \
		-w /app golangci/golangci-lint:v1.44.0 \
		golangci-lint run ./app/... --fix

boil:
	docker-compose exec app sqlboiler mysql -c db/sqlboiler.toml --add-soft-deletes --add-enum-types --add-global-variants --add-panic-variants

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
	docker-compose exec mysql mysql -uroot -proot gosample2022_test < ./testsupport/seed/seed.sql
