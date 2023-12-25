postgresinit:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:16-alpine

pg-createdb:
	docker exec -it postgres16 createdb --username=root --owner=root chatdb

pg-dropdb:
	docker exec -it postgres16 dropdb chatdb

pg-cli:
	docker exec -it postgres16 psql

.PHONY: postgresinit pg-createdb pg-dropdb pg-cli