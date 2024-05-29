FROM golang:1.22.3 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go-weather-app ./cmd

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /go-weather-app /go-weather-app

COPY --from=builder /app/cmd/.env .

EXPOSE 8080

CMD ["/go-weather-app"]
