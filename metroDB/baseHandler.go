package main

import (
  "net/http"
  "fmt"
)

func getIndex(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintln(
      w,
      "{\n\t\"message\" : \"Welcome to metroDB! Version 0.01\"\n\t\"doc\" : \"/doc\"\n}")
}

func getHealthCheck(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  fmt.Fprintln(
    w,
    "healthy")
}

func getApiTree(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintln(
      w,
      "{\n\"doc\" : [\n\t\"/\",\n\t\"/metrics\",\n\t\"/doc\",\n\t\"/api/\",\n\t\"/api/<stackName>\",\n\t\"/api/<stackName>/<blockID>\"\n\t]\n}")
}
