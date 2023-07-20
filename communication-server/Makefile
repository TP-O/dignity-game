PG_PATH=infrastructure/postgresql
PG_URL="postgresql://dgame:dgame@postgres:5432/dgame?sslmode=disable"

db-create:
	docker compose exec app migrate create -ext sql -dir ${PG_PATH}/migration -seq migration

db-migrate:
	docker compose exec app migrate -path ${PG_PATH}/migration \
        -verbose \
		-database ${PG_URL} up

db-migrate-1:
	docker compose exec app migrate -path ${PG_PATH}/migration \
        -verbose \
		-database ${PG_URL} up 1

db-rollback:
	docker compose exec app migrate -path ${PG_PATH}/migration \
        -verbose \
		-database ${PG_URL} down

db-rollback-1:
	docker compose exec app migrate -path ${PG_PATH}/migration \
        -verbose \
		-database ${PG_URL} down 1

db-drop:
	docker compose exec app migrate -path ${PG_PATH}/migration \
        -verbose \
		-database ${PG_URL} drop

db-seed:
	docker compose exec app echo "seed"

db-code: docker compose exec app sqlc generate
