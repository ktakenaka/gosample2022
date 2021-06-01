FROM golang:1.16-alpine

WORKDIR /db
RUN apk add --no-cache --virtual=build-deps git
RUN CGO_ENABLED=0 go get -v github.com/rubenv/sql-migrate/...

ENTRYPOINT [ "/go/bin/sql-migrate" ]
CMD [""]
