FROM golang:1.14 AS builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o healthcheck cmd/healthcheck/main.go 

FROM scratch

COPY --from=builder /app/healthcheck .

EXPOSE 8080

CMD ["/healthcheck"]