tsp.build:
	docker compose build

tsp.init:
	docker compose run --rm typespec tsp init

tsp.compile:
	docker compose run --rm typespec tsp compile .
