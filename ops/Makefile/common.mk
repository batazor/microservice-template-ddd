# APPLICATION TASKS ====================================================================================================

export CURRENT_UID=$(id -u):$(id -g)

run: ## Run this project in docker-compose
	@docker-compose \
		-f docker-compose.yaml \
		-f ops/docker-compose/application/api.yaml \
		-f ops/docker-compose/application/billing.yaml \
		-f ops/docker-compose/application/user.yaml \
		-f ops/docker-compose/application/book.yaml \
		-f ops/docker-compose/infrastructure/traefik.yaml \
		-f ops/docker-compose/infrastructure/redis.yaml \
		-f ops/docker-compose/infrastructure/opentracing.yaml \
		up -d --remove-orphans

down: ## Down docker-compose
	@docker-compose \
		-f docker-compose.yaml \
		-f ops/docker-compose/application/api.yaml \
		-f ops/docker-compose/application/billing.yaml \
		-f ops/docker-compose/application/user.yaml \
		-f ops/docker-compose/application/book.yaml \
		-f ops/docker-compose/infrastructure/traefik.yaml \
		-f ops/docker-compose/infrastructure/redis.yaml \
		-f ops/docker-compose/infrastructure/opentracing.yaml \
		down --remove-orphans
