FROM golang:1.16.4-alpine

RUN apk update && apk upgrade && apk add build-base

WORKDIR /sqlmigrate
RUN go get -v github.com/rubenv/sql-migrate/...

ENTRYPOINT [ "/go/bin/sql-migrate" ]
CMD [""]
