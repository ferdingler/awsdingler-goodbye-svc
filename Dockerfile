FROM golang:1.12

# GOPATH by default is /go
ADD app/ $GOPATH/src/app

RUN go get github.com/gin-gonic/gin
RUN go install app

ENTRYPOINT $GOPATH/bin/app