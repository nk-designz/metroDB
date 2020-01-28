#/usr/bin/bash
protoc --go_out=plugins=grpc,import_path=mapd:. *.proto