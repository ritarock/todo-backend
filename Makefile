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

ogen:
	go run -mod=mod github.com/ogen-go/ogen/cmd/ogen@latest \
	-package rest \
	-target api/rest \
	-clean ./tsp-output/schema/openapi.yaml

create.table:
	sqlite3 data.sqlite < ./sqlc/schema.sql
