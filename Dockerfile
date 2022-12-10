FROM golang:1.19-alpine as builder
WORKDIR /app
COPY . .
RUN apk add upx
RUN go mod tidy
RUN go build \
    -ldflags "-s -w" \
    -o /app/main main.go
RUN upx -9 /app/main

FROM alpine:latest
ENV APP_PORT=9000
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
EXPOSE ${APP_PORT}
CMD /app/main