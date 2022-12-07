FROM golang:1.19.4 as builder
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download

# Build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o hmac

FROM alpine:latest 
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=builder /build/hmac /app
EXPOSE 8080
CMD ["/app/hmac","server"]
