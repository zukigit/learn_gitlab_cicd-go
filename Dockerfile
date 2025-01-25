# stage 1: build the project
FROM golang:1.23.5-alpine3.21 AS builder
WORKDIR /app
COPY . .
RUN go build -o test_cicd

# stage 2: minimal runtime image
FROM alpine:3.21
ENV TEST_PORT=123

WORKDIR /app
COPY --from=builder /app/test_cicd .
RUN chmod +x test_cicd

ENTRYPOINT [ "sh", "-c", "./test_cicd $TEST_PORT" ]