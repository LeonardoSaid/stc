help: ## Show Help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

test: ## Run unit tests
	touch count.out
	go test -covermode=count -coverprofile=count.out -v ./...
	$(MAKE) coverage

coverage: ## Unit tests coverage
	go tool cover -func=count.out

lint: ## Run linter
	go mod vendor
	docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint:v1.39.0 golangci-lint run -v

copy-env: ## Copy .env.example to .env
	cp .env.sample .env

mock-generate: ## Generate mocks
	go mod vendor
	docker run --rm -v "$(PWD):/app" -w /app -t vektra/mockery --all --dir ./internal/infra/repositories --case underscore
	docker run --rm -v "$(PWD):/app" -w /app -t vektra/mockery --all --dir ./internal/usecases --case underscore
	docker run --rm -v "$(PWD):/app" -w /app -t vektra/mockery  --dir /app/pkg/stc-sdk --case underscore --all --output ./mocks/pkg
	docker run --rm -v "$(PWD):/app" -w /app -t vektra/mockery --all --dir ./internal/clients --case underscore

format: ## Format code
	docker run --rm -ti \
		--platform linux/amd64 \
		--name=pix-dict-fmt-format \
		-v `pwd`:/code \
		-w /code \
		golang:1.17-buster gofmt -d -w .

	docker run --rm -ti \
		--platform linux/amd64 \
		--name=pix-dict-goimports-format \
		-v `pwd`:/code \
		-w /code \
		-e GO111MODULE=off \
		golang:1.17-buster \
		sh -c "go get -u golang.org/x/tools/cmd/goimports; go run golang.org/x/tools/cmd/goimports -v -w -d -l ."

	docker run --rm -ti \
		--platform linux/amd64 \
		--name=pix-dict-golines-format \
		-v `pwd`:/code \
		-w /code \
		-e GO111MODULE=off \
		golang:1.17-buster \
		sh -c "go get -u github.com/segmentio/golines; golines . -m 120 -w --ignore-generated --no-reformat-tags --debug"
