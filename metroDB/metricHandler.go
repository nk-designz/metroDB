package main

import (
  "net/http"
  "fmt"
  "time"
)

type MetricReport struct {
    Time     int64
    Stacks   uint32
    Blocks uint32
    Size    uint32
}

func getMetrics(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var MetricReport MetricReport
    MetricReport.Blocks = 0;
    MetricReport.Stacks = 0;

    for _, stackName := range stackMap {
      for _, block := range stackName.blocks {
          MetricReport.Size += uint32(cap(block.Data))
          MetricReport.Blocks++
      }
      MetricReport.Stacks++
    }
    MetricReport.Time = time.Now().Unix()


    returnValue := fmt.Sprintf("{\n\t\"sysTime\" : \"%v\",\n\t\"stacks\" : \"%v\",\n\t\"blocks\" : \"%v\",\n\t\"size\" : \"%v\"\n}",
      MetricReport.Time,
      MetricReport.Stacks,
      MetricReport.Blocks,
      MetricReport.Size)


    fmt.Fprintln(w, returnValue)
}
