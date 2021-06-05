FROM golang:1.16.4-alpine

RUN apk update && && apk upgrade && apk add build-base

WORKDIR /gomsx
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
ENTRYPOINT [ "/go/bin/staticcheck" ]
CMD [ "" ]
