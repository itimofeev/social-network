
test::
	go test -race -v ./...


lint::
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2 -v run ./...


.PHONY: gen
gen:
	@go generate ./...

.PHONY: up
up:
	docker compose up -d --remove-orphans
down:
	docker compose down

up-full:
	docker compose --file docker-compose.yc.yml up -d --remove-orphans

run-migrations:
	GOOSE_DRIVER=postgres \
	GOOSE_DBSTRING=postgres://admin:admin@127.0.0.1:5432/social-network?sslmode=disable \
	go run github.com/pressly/goose/v3/cmd/goose@v3.15.1 -dir migrations up

build-docker:
	docker build -f Dockerfile . \
		  --platform linux/amd64 \
          --tag social-network:local
	docker tag social-network:local itimofeev/social-network:2.0.0
	docker push itimofeev/social-network:2.0.0