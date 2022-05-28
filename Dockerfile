FROM golang:1.17-alpine

ENV GIN_MODE=release
ENV port=8000

COPY . /go/src/convert-jenkinsfile
WORKDIR /go/src/convert-jenkinsfile


# dependency 다운로드를 위한 git 설치
RUN apk update && apk add --no-cache git && apk add build-base
RUN go install github.com/swaggo/swag/cmd/swag@v1.6.7
RUN swag init
RUN go get ./

RUN go build main.go

EXPOSE $PORT
ENTRYPOINT ["./main"]