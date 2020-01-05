FROM golang:latest AS logd-builder
RUN go get "github.com/nk-designz/metrodb/services/logd/server"
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o /bin/logd /go/src/github.com/nk-designz/metroDB/services/logd/server

FROM golang:latest AS mond-builder
RUN go get "github.com/nk-designz/metrodb/services/mond"
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o /bin/mond /go/src/github.com/nk-designz/metroDB/services/mond

FROM alpine
COPY --from=logd-builder /bin/logd /bin/logd
COPY --from=mond-builder /bin/mond /bin/mond
RUN mkdir /opt/logd
EXPOSE 7558
ENV LOGD_DB_PATH=/opt/logd-server