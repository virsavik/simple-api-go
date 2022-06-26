ifneq (,$(wildcard ./.env))
    include .env
    export
endif

run:
	docker-compose --env-file .env up -d

migrate-new:
	migrate create -ext sql -dir $(MIGRATION_LOCATION) -seq $(name)

migrate-force:
	migrate -path $(MIGRATION_LOCATION) -database $(DATABASE_URL) force $(V)

migrate-up:
	migrate -path=$(MIGRATION_LOCATION) -database $(DATABASE_URL) up

migrate-down:
	migrate -path=$(MIGRATION_LOCATION) -database $(DATABASE_URL) down

migrate-goto:
	migrate -path=$(MIGRATION_LOCATION) -database $(DATABASE_URL) goto $(v)

migrate-reset:
	migrate -path=$(MIGRATION_LOCATION) -database $(DATABASE_URL) drop

boil-gen:
	sqlboiler psql

doc-gen:
	swag init -g cmd/s3corp-golang-fresher/main.go -o api
