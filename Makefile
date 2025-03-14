# Makefile
export

compose-build:
	docker-compose build

compose-with-debug: compose-build
	@echo "Starting in the debug mode for container"
	@docker compose up 

compose-without-app: compose-build
	@echo "Starting in the debug mode for container"
	@docker compose up --scale app=0 -d

compose-up: compose-build
	@docker compose up -d

compose-stop:
	@echo "stopping docker compose in background"
	@docker compose down

compose-clean: compose-stop
	docker-compose rm -f

migrate-up:
	@echo "migrating DB with Latest Changes"
	$(MIG_UP)

migrate-down:
	@echo "migrating down DB with Latest Changes"
	$(MIG_DOWN)