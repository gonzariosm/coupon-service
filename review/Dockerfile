# build stage
# Comment: use a fixed version of golang
FROM golang:1.23-alpine AS builder

# Commet: no-cache
RUN apk add --no-cache git gcc libc-dev

WORKDIR /go/src/coupon-service

# Commet: to create a layer of cache to next build
COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /go/src/coupon-service/cmd/coupon_service
RUN go build -ldflags="-s -w" -o main .

FROM alpine:3.20
RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/coupon-service/cmd/coupon_service/main /coupon-service

WORKDIR /

EXPOSE 8080

ENTRYPOINT ["/coupon-service"]