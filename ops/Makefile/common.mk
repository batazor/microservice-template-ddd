# APPLICATION TASKS ====================================================================================================

export CURRENT_UID=$(id -u):$(id -g)

run: ## Run this project in docker-compose
	@docker-compose \
		-f docker-compose.yaml \
		-f ops/docker-compose/traefik.yaml \
		up -d --remove-orphans

down: ## Down docker-compose
	@docker-compose \
		-f docker-compose.yaml \
		-f ops/docker-compose/traefik.yaml \
		down --remove-orphans
