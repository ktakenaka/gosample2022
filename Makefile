migrate-new:
	docker-compose run migration new ${name}

migrate-%:
	$(eval CMD:= $*)
	docker-compose run migration $(CMD)
