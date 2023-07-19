FROM golang:latest AS builder

WORKDIR /app/mockserver

COPY docker/mockserver /app/mockserver
COPY go.mod /app/mockserver
COPY go.sum /app/mockserver

ENV GOPROXY https://goproxy.cn,direct

RUN go mod download

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build mockserver.go

FROM alpine AS runner

WORKDIR /app/mockserver

COPY --from=builder /app/mockserver/mockserver .

EXPOSE 8080

ENTRYPOINT ["./mockserver"]
