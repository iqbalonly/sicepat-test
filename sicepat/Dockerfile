# TODO refactor into multistage build
FROM golang:1.16.1-alpine

WORKDIR /usr/src/app

ADD go.mod .

COPY . .

RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o main ./cmd && \
  chmod +x main

ARG port

EXPOSE $port

CMD ["/usr/src/app/main"]
