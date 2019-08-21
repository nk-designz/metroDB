
FROM golang

ADD . /go/src/github.com/golang/nka/metroDB
RUN go get /go/src/github.com/golang/nka/metroDB
RUN go install /go/src/github.com/golang/nka/metroDB

EXPOSE 8756 8756
HEALTHCHECK CMD curl http://localhost:8756/health
ENTRYPOINT /go/bin/metroDB
