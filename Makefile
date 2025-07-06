tsp.build:
	docker compose build

tsp.init:
	docker compose run --rm typespec tsp init

tsp.compile:
	docker compose run --rm typespec tsp compile .

sqlc.pull:
	docker pull sqlc/sqlc

sqlc.generate:
	docker run --rm -v $(PWD):/src -w /src sqlc/sqlc generate
