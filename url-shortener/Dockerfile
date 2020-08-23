FROM golang:1.15 AS builder
WORKDIR /code
COPY ./ ./
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/url-shortener ./url-shortener/cmd/cli

FROM alpine:latest
RUN apk --no-cache -U add ca-certificates
COPY --from=builder /code/bin/url-shortener /usr/bin/url-shortener
RUN chmod a+x /usr/bin/url-shortener
ENTRYPOINT ["/usr/bin/url-shortener"]


