-include .env
-include .env.prod
export

.PHONY: gensql genswag gen lint dev-up dev-down prod-up prod-down clean

gensql:
	@chmod +x scripts/gen-sql.sh
	@./scripts/gen-sql.sh

genswag:
	@chmod +x scripts/gen-swagger.sh
	@./scripts/gen-swagger.sh

gen: gensql genswag

lint:
	@chmod +x scripts/lint.sh
	@./scripts/lint.sh

dev-up:
	docker compose --env-file .env up -d --build

dev-down:
	docker compose --env-file .env down -v

prod-up:
	docker compose --env-file .env.prod -f compose.yaml -f compose.prod.yaml up -d --build

prod-down:
	docker compose --env-file .env.prod -f compose.yaml -f compose.prod.yaml down -v

clean:
	rm -rf ./tmp