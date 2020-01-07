FROM golang:latest AS logd-builder
RUN go get "github.com/nk-designz/metrodb/services/logd/server"
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o /bin/logd /go/src/github.com/nk-designz/metroDB/services/logd/server

FROM golang:latest AS mond-builder
RUN go get "github.com/nk-designz/metrodb/services/mond"
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o /bin/mond /go/src/github.com/nk-designz/metroDB/services/mond

FROM golang:latest AS mapd-builder
RUN go get "github.com/nk-designz/metrodb/services/mapd/server"
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o /bin/mapd /go/src/github.com/nk-designz/metroDB/services/mapd/server
RUN go get "github.com/nk-designz/metrodb/services/mapd/cmd"
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o /bin/mapd-util /go/src/github.com/nk-designz/metroDB/services/mapd/cmd

FROM registry.access.redhat.com/ubi8/ubi-minimal:latest
RUN mkdir /var/metrodb
# Install logd
COPY --from=logd-builder /bin/logd /bin/logd
RUN mkdir /var/metrodb/logd
ENV LOGD_DB_PATH=/opt/metrodb/logd
EXPOSE 7558
# Install mond
COPY --from=mond-builder /bin/mond /bin/mond
# Install mapd
COPY --from=mapd-builder /bin/mapd /bin/mapd
COPY --from=mapd-builder /bin/mapd-util /bin/mapd-util
RUN mkdir /var/metrodb/mapd
ENV MAPD_DB_PATH=/opt/metrodb/mapd
EXPOSE 7550
