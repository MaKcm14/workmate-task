FROM golang:1.23-alpine3.21

LABEL maintainer="maksimacx50@gmail.com"

WORKDIR /task_service

COPY . .

RUN go mod tidy

WORKDIR /task_service/cmd/app

VOLUME /task_service/logs

ENTRYPOINT ["go", "run", "main.go"]