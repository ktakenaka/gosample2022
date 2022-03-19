FROM golang:1.18-alpine3.15

RUN apk update && apk upgrade && apk add build-base
RUN go install github.com/rubenv/sql-migrate/...@latest
ENTRYPOINT [ "/go/bin/sql-migrate" ]
