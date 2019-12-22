.PHONY: build remove run

build:
		cd auth-gateway && make build
		cd users-gateway && make build
		cd users-service && make build

remove:
		docker rmi gomicroservices_auth-gateway -f
		docker rmi gomicroservices_users-gateway -f
		docker rmi gomicroservices_users-service -f

run:
		docker-compose up