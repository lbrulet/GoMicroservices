.PHONY: build remove run

build:
		cd auth && make build
		cd users && make build

remove:
		docker rmi gomicroservices_users -f
		docker rmi gomicroservices_api -f

run:
		docker-compose up