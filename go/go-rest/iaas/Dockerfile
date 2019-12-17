FROM golang:alpine AS builder

RUN apk add --no-cache build-base
RUN apk add --no-cache make
RUN apk add --no-cache git
RUN apk add --no-cache bash
RUN apk add --no-cache openssl

WORKDIR /app
COPY . .
RUN make build

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/build/bin/products /app
CMD ["./products"]⏎