.PHONY: dev

up:
	docker compose up --pull always \
		--build app

dev: 
	docker compose run \
		--rm app \
		bash

dev-down:
	docker compose down --volumes --remove-orphans
