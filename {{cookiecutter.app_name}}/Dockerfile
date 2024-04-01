FROM golang:1.22 AS base

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o {{cookiecutter.app_name}} .

FROM golang:1.22-alpine
WORKDIR /
COPY --from=base /app   /
ENTRYPOINT ["./{{cookiecutter.app_name}}", "serve"]
