# gpt-enterprise-ui

## Prerequisites

Running this project requires following dependencies

1. Go (to build go binary)
2. Docker (to orchestrate container)
3. Docker Compose (to manage containers)

## Steps  to install

1. Build go binary using (everytime we change something in go code) `GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build main.go`
2. Run `docker compose up -d`