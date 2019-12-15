.PHONY: build run

build:
		cd auth && make build
		cd users && make build

run:
		docker-compose up