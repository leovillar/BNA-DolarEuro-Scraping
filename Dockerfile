FROM golang:1.18 as builder
ENV GO111MODULE=on
WORKDIR /app
COPY . .
COPY .env .env
RUN go test -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dolar-bna-scraping

FROM scratch
WORKDIR /app
COPY --from=builder /app/dolar-bna-scraping /app/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["/app/dolar-bna-scraping"]
