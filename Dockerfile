FROM golang:1.24-alpine

LABEL maintainer="maksimacx50@gmail.com"

WORKDIR /task_service

COPY . .

RUN go mod tidy

WORKDIR /task_service/cmd/app

VOLUME /task_service/logs

EXPOSE 8080

ENTRYPOINT ["go", "run", "main.go"]