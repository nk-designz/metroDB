FROM golang:latest AS builder
WORKDIR /go/src/github.com/nk-designz/metrodb/services/logd/server
RUN go build "github.com/nk-designz/metrodb/services/logd/server"

FROM alpine
COPY --from=builder /go/src/github.com/nk-designz/metrodb/services/logdb/server/logd /bin/logd
RUN mkdir /opt/logd
EXPOSE 70558
ENV LOGD_DB_PATH=/opt/logd
CMD [ "logd" ]