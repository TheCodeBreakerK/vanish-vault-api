include .env
export

.PHONY: gensql genswag gen lint dev prod clean

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

dev:
	docker compose up --build

prod:
	docker compose -f compose.yaml -f compose.prod.yaml up -d --build

clean:
	rm -rf ./tmp