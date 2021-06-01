FROM golang:1.16.4-alpine

WORKDIR /sqlboiler
RUN go install github.com/volatiletech/sqlboiler/v4@latest &&\
  go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest

ENTRYPOINT [ "/go/bin/sqlboiler" ]
CMD [""]
