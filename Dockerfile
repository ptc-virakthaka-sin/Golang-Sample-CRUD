FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates

COPY bin/app .
COPY .env .

EXPOSE 8000
CMD ["./app"]
