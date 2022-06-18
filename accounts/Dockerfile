#### development stage
FROM golang:1.17-buster AS development

# for clone private repos
ARG GITHUB_USER
ARG GITHUB_TOKEN

# set envvar
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GO111MODULE='on'

# set workdir
WORKDIR /source

# get project dependencies
COPY go.mod go.sum /source/
RUN go mod download

# copy files
COPY . /source

#### builder stage
FROM development AS builder
RUN go build -o ./app ./cmd/main.go

# Run
FROM alpine as run
ARG GRPC_HEALTH_PROBE_VERSION=v0.3.6
COPY --from=builder /source/app /
ADD https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 /bin/grpc_health_probe
RUN chmod +x /bin/grpc_health_probe
ENTRYPOINT ["/app"]
