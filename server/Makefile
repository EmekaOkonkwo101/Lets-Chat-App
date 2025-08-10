include .env
export

postgresinit:
	docker run --name $(POSTGRES_CONTAINER) -p $(POSTGRES_PORT):5432 \
		-e POSTGRES_USER=$(POSTGRES_USER) \
		-e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
		-d $(POSTGRES_IMAGE)

postgres:
	docker exec -it --user=$(POSTGRES_USER) $(POSTGRES_CONTAINER) psql

createdb:
	docker exec -it --user=$(POSTGRES_USER) $(POSTGRES_CONTAINER) \
		createdb -O $(POSTGRES_USER) $(POSTGRES_DB)

dropdb:
	docker exec -it --user=$(POSTGRES_USER) $(POSTGRES_CONTAINER) \
		dropdb $(POSTGRES_DB)

migrateup:
	migrate -path db/migrations \
		-database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" \
		-verbose up

migratedown:
	migrate -path db/migrations \
		-database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" \
		-verbose down

runserver:
	go run cmd/main.go

.PHONY: postgresinit postgres createdb dropdb migrateup migratedown runserver
