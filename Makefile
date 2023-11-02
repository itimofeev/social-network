
test::
	go test -race -v ./...


lint::
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.0 -v run ./...


.PHONY: gen
gen:
	@go generate ./...

.PHONY: up
up:
	docker-compose up
down:
	docker-compose down