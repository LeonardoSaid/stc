# Stone Tech Challenge - Go(lang)

Repositório dedicado à resolução do desafio: [https://gist.github.com/guilhermebr/fb0d5896d76634703d385a4c68b730d8](https://gist.github.com/guilhermebr/fb0d5896d76634703d385a4c68b730d8)

## **Requisitos**

- [Docker](https://docs.docker.com/install/)
- [Docker-compose](https://docs.docker.com/compose/install/)
- Arquivo .env para configurar as variáveis default do projeto
- Ambiente Unix/Linux-based para utilizar os comandos do Makefile
## Setup

1. Configurar os arquivos .env dos serviços (login/transfer/account) com os valores da porta e host dos serviços descritos no `docker-compose.yml`, para utilizar o arquivo .env.sample como base basta executar o comando do Makefile:
```
make copy-env
```

2. Com as variáveis configuradas gere o build e suba o sistema:
```
docker compose up -d --build
```

## Exemplo
Com os containers ativos, a API estará disponível para receber requests localmente, no diretório /docs foi disponibilizado uma collection e environment do Postman com as rotas configuradas para teste das APIs.

## Comandos úteis

O arquivo Makefile que acompanha cada serviço também possui outros comandos úteis para o desenvolvimento da aplicação, como:

- `make lint` - Aplica o lint no código
- `make test` - Executa os testes unitários e de integração
- `make coverage` - Gera o relatório de code coverage
- `make mock-generate` - Gera os arquivos de mock utilizados para os testes

## **Desenvolvido com**

- `[golang] (https://go.dev/)`
- `[fx] (https://github.com/uber-go/fx)`
- `[bun] (https://bun.uptrace.dev/)`
- `[echo] (https://echo.labstack.com/)`
- `[postgres] (https://www.postgresql.org/)`
