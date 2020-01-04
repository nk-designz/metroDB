FROM golang:latest AS builder
RUN go get "github.com/nk-designz/metrodb/services/logd/server"
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o /bin/logd /go/src/github.com/nk-designz/metroDB/services/logd/server

FROM alpine
COPY --from=builder /bin/logd /bin/logd
RUN mkdir /opt/logd
EXPOSE 7558
ENV LOGD_DB_PATH=/opt/logd-server
ENTRYPOINT [ "logd" ]