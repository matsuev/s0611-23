# STEP-1
# build application from source

FROM golang:1.23.7-alpine3.21 AS builder

WORKDIR /appsource

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./go.mod ./

RUN go mod tidy

RUN go build -o app ./cmd/main.go


# STEP-2
# make container

FROM alpine:3.21

WORKDIR /myapp

COPY --from=builder /appsource/app ./

EXPOSE 7080

CMD [ "/myapp/app" ]