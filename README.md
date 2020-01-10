![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/nicokahlert/metrodb?style=flat-square)  
![Logo](/assets/images/logo.png)  
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
