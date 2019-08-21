package main

import (
  "net/http"
  "fmt"
  "strings"
  "io/ioutil"

  "github.com/gorilla/mux"
)

func addNewStack(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    vars := mux.Vars(r)
    stackName := vars["stackName"]

    returnValue := ""
    _ = returnValue

    if _, ok := stackMap[stackName]; ok {
      returnValue = fmt.Sprintf("{\n\t\"error\" : \"NotCreatableAlreadyExists\",\n\t\"name\" : \"%s\"\n}", stackName)
    } else {

      stackMap[stackName] = NewBlockchain()
      returnValue = fmt.Sprintf("{\n\t\"name\" : \"%s\",\n\t\"selfLink\" : \"/api/%s\"\n}", stackName, stackName)
    }

    fmt.Fprintln(
      w,
      returnValue)
}


func getStackInfo(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    vars := mux.Vars(r)
    stackName := vars["stackName"]

    returnValue := ""
    _ = returnValue

    if _, ok := stackMap[stackName]; ok {
      returnValue = fmt.Sprintf("{\n\t\"name\" : \"%s\",\n\t\"selfLink\" : \"/api/%s\",\n\t\"blocks\" : \"%v\",\n\t\"lastEntry\" : \"%x\",\n\t\"modified\" : \"%v\"\n}",
        stackName,
        stackName,
        len(stackMap[stackName].blocks),
        stackMap[stackName].blocks[(len(stackMap[stackName].blocks) - 1)].Hash,
        stackMap[stackName].blocks[(len(stackMap[stackName].blocks) - 1)].Timestamp)
    } else {
      returnValue = fmt.Sprintf("{\n\t\"error\" : \"NotExists\",\n\t\"name\" : \"%s\"\n}", stackName)
    }

    fmt.Fprintln(
      w,
      returnValue)
}


func getStack(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  vars := mux.Vars(r)
  stackName := vars["stackName"]

  returnValue := ""
  _ = returnValue

  var thisBlock *Block = NewBlock("", []byte{})

  if _, ok := stackMap[stackName]; ok {
    thisBlock = stackMap[stackName].blocks[(len(stackMap[stackName].blocks) - 1)]

    if(thisBlock == NewBlock("", []byte{})) {
      returnValue = fmt.Sprintf("{\n\t\"error\" : \"NotReadableBlockNotExists\",\n\t\"hash\" : \"-1\"\n}")
    } else {
      returnValue = fmt.Sprintf(
        "{\n\t\"hash\" : \"%x\",\n\t\"selfLink\" : \"/api/%s\",\n\t\"modified\" : \"%v\",\n\t\"previousBlock\" : \"%x\"\n\t\"data\" : \"%s\"\n}",
        string(thisBlock.Hash),
        stackName,
        stackMap[stackName].blocks[(len(stackMap[stackName].blocks) - 1)].Timestamp,
        thisBlock.PrevBlockHash,
        thisBlock.Data)
    }

  } else {
    returnValue = fmt.Sprintf("{\n\t\"error\" : \"NotReadableStackNotExists\",\n\t\"name\" : \"%s\"\n}", stackName)
  }

  fmt.Fprintln(
    w,
    returnValue)
}

func getStackList(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  returnValue := "{\n\t\"stacks\" : [\n\t\t"
  _ = returnValue

  keys := make([]string, 0)

  for key := range stackMap {
    keys = append(keys, fmt.Sprintf("\"%s\"", key))
  }

  returnValue += fmt.Sprintf(strings.Join(keys, ",\n\t\t"))
  returnValue += "\n\t]\n}"

  fmt.Fprintln(
    w,
    returnValue)
}

func setLabel(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    vars := mux.Vars(r)
    stackName := vars["stackName"]

    returnValue := ""
    _ = returnValue

    data, err := ioutil.ReadAll(r.Body)
    if err != nil {
      panic(err)
    }

    stackLabel := string(data)

    if _, ok := stackMap[stackName]; ok {
      stackMap[stackName].labels = append(stackMap[stackName].labels, stackLabel)
      returnValue = fmt.Sprintf("{\n\t\"name\" : \"%s\",\n\t\"label\" : \"%s\"\n}", stackName, stackLabel)
    } else {
      returnValue = fmt.Sprintf("{\n\t\"error\" : \"NotCreatableNotExists\",\n\t\"name\" : \"%s\"\n}", stackName)
    }

    fmt.Fprintln(
      w,
      returnValue)
}

func getLabels(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  vars := mux.Vars(r)
  stackName := vars["stackName"]
  returnValue := "{\n\t\"labels\" : [\n\t\t"
  _ = returnValue

  for label := range stackMap[stackName].labels {
      returnValue += fmt.Sprintf("\"%s\",\n\t\t", stackMap[stackName].labels[label])
  }

  returnValue += "\n\t]\n}"

  fmt.Fprintln(
    w,
    returnValue)
}
