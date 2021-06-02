migrate-new:
	docker-compose run --rm migration new ${name}

migrate-%:
	$(eval CMD:= $*)
	docker-compose run --rm migration $(CMD)

sqlboiler:
	docker-compose run --rm sqlboiler mysql --config /sqlboiler.toml

mod:
	docker-compose exec app go mod tidy
	docker-compose exec app go mod vendor
