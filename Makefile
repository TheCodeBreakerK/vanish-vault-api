include .env
export

.PHONY: gensql genswag gen lint dev-up dev-down prod clean

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
	docker compose up --build

dev-down:
	docker compose down -v

prod:
	docker compose -f compose.yaml -f compose.prod.yaml up -d --build

clean:
	rm -rf ./tmp