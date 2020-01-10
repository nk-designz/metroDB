![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/nicokahlert/metrodb?style=flat-square)  
![Logo](https://raw.githubusercontent.com/nk-designz/metroDB/master/assets/images/logo.png)  
# MetroDB
A distributed, scalable NoSQL Datastore based on immutable logs.
# Components
## logd
The Log-deamon is a immutable Logstore. It is managed by Mapd via GRPC calls on port 7558.
A byte-value is send to logd and an offset is received.
## mapd
The Map-deamon is a persistent in-memory keygrid.
A Value and Key is send to Mapd.
Mapd will send the Value to a Logd and store the received offset in the keygrid with its Key and Logd as a replica.
The mapd will communicate every change between its replicas.
## mond
The Monitor-daemon communicates with the map-daemons and validates the replicas and adds more if there are too many broken.
# Installation
## Local
For a testing enviroment we can use the [docker-compose](https://docs.docker.com/compose/) utility:
```bash
git clone https://github.com/nk-designz/metroDB.git
cd metroDB/deployment/docker-compose/
docker-compose up
```
This will build us a three node mapd instance with three logds.
## Kubernetes
Of course do we provide nice YAMLs for deploying metroDB on the worlds favorite orchestrator:
```bash
git clone https://github.com/nk-designz/metroDB.git
cd metroDB/deployment/kubernetes/example
kubectl apply -f . --namespace <the namespace you want to deploy in>
```
# Connecting
Now you can use the toolbox-container for the mapd-util:
```bash
docker exec -it docker-compose_toolbox_1 /bin/ash
# the docker-compose_toolbox_1 might differ
# try docker ps
mapd-util mapd1 set <key> <value>
```
OR use the started REST endpoint at localhost:7080
```bash
curl -X POST localhost:7080/<key> -d "<value>"
curl -X POST localhost:7080/<key> -d "@/<my>/<file>/<location>.<fileending>"

curl localhost:7080/<key>
curl localhost:7080/<key> --output /<my>/<file>/<location>.<fileending>
```
Or try the go-client
```go
import(
  "fmt"
  mapd "github.com/nk-designz/metroDB/services/mapd/client"
)

const(
  address = "127.0.0.1"
)

func main() {
  db := mapd.New(address)
  _, err := db.SetSafe("<key>", []byte("<value>")); if err != nil {
    panic(err)
  }
  _, err = db.Set("<key>", []byte("<value>")); if err != nil {
    panic(err)
  }
  value, err = db.Get("<key>"); if err != nil {
    panic(err)
  }
  fmt.Println(
    fmt.Sprintf("My value: %s", value))
}
```
