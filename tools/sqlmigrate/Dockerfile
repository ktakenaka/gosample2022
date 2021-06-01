FROM golang:1.16.4-alpine

WORKDIR /db
RUN CGO_ENABLED=0 go get -v github.com/rubenv/sql-migrate/...

ENTRYPOINT [ "/go/bin/sql-migrate" ]
CMD [""]
