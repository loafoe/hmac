FROM golang:1.13.7 as builder
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download

# Build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o hsdp-metrics-alert-collector

FROM alpine:latest 
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=builder /build/hsdp-metrics-alert-collector /app
EXPOSE 8080
CMD ["/app/hsdp-metrics-alert-collector"]
