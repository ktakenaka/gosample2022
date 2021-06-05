FROM golang:1.16.4-alpine

RUN apk update && apk upgrade

WORKDIR /sqlboiler
RUN go install github.com/volatiletech/sqlboiler/v4@latest &&\
  go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest

ENTRYPOINT [ "/go/bin/sqlboiler" ]
CMD [""]
