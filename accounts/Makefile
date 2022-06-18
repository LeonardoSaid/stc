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

start: ## Start pix-dict service and Its dependencies
	docker-compose up -d --build # Run docker-compose

mock-generate: ## Generate mocks
	go mod vendor
	docker run --rm -v "$(PWD):/app" -w /app -t vektra/mockery --all --dir ./internal/domain --case underscore --keeptree
	docker run --rm -v "$(PWD):/app" -w /app -t vektra/mockery --dir /app/vendor/github.com/hashlab/banking-accounts/api/infra/grpc/v1 --case underscore --name AccountServiceClient --output ./mocks/client
	docker run --rm -v "$(PWD):/app" -w /app -t vektra/mockery --dir /app/vendor/github.com/hashlab/banking-company-api/api/v1/genpb --case underscore --name CompanyServiceClient --output ./mocks/client
	docker run --rm -v "$(PWD):/app" -w /app -t vektra/mockery  --dir /app/pkg/btg --case underscore --name=ClientBTG --output ./mocks/pkg
	docker run --rm -v "$(PWD):/app" -w /app -t vektra/mockery  --dir /app/pkg/slack --case underscore --name=Slack --output ./mocks/pkg
	docker run --rm -v "$(PWD):/app" -w /app -t vektra/mockery  --dir /app/pkg/cloudstorage --case underscore --name=Storage --output ./mocks/pkg

migrate: ## Run database migrations
	migrate -path migrations/ -database postgres://dict_user:D1ct@localhost:5432/pix_dict?sslmode=disable up

migrate-create: ## Create migration file
	migrate create -ext sql -dir migrations -tz America/Sao_Paulo $(file_name)

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
