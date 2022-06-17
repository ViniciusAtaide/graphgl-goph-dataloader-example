#!make

include .env
export $(shell sed 's/=.*//' .env)

up:
	migrate -path ./migrations/ -database $(URL_DB) up

force:
	migrate -path ./migrations/ -database $(URL_DB) force $(VERSION)

down:
	migrate -path ./migrations/ -database $(URL_DB) down

create:
	migrate create -ext sql -dir migrations -seq $(FILE)
