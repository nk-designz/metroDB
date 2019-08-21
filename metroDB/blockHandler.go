package main

import (
  "net/http"
  "fmt"
  "io/ioutil"

  "github.com/gorilla/mux"
)


func addNewBlock(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    vars := mux.Vars(r)
    stackName := vars["stackName"]

    data, err := ioutil.ReadAll(r.Body)
    if err != nil {
      panic(err)
    }

    returnValue := ""
    _ = returnValue

    if _, ok := stackMap[stackName]; ok {
      stackMap[stackName].AddBlock(string(data))
      hash := stackMap[stackName].blocks[len(stackMap[stackName].blocks) - 1].Hash
      returnValue = fmt.Sprintf("{\n\t\"hash\" : \"%x\",\n\t\"selfLink\" : \"/api/%s/%x\"\n}", hash ,stackName, hash)
    } else {
      returnValue = fmt.Sprintf("{\n\t\"error\" : \"NotCreatableStackNotExists\",\n\t\"name\" : \"%s\"\n}", stackName)
    }

    fmt.Fprintln(
      w,
      returnValue)
}

func getBlock(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  vars := mux.Vars(r)
  stackName := vars["stackName"]
  blockId := vars["blockId"]

  returnValue := ""
  _ = returnValue
  var thisBlock *Block = NewBlock("", []byte{})

  if _, ok := stackMap[stackName]; ok {
    for _, block := range stackMap[stackName].blocks {
      if( fmt.Sprintf("%x", block.Hash) == blockId) {
        thisBlock = block
      }
    }
  } else {
    returnValue = fmt.Sprintf("{\n\t\"error\" : \"NotReadableStackNotExists\",\n\t\"name\" : \"%s\"\n}", stackName)
  }

  if(thisBlock == NewBlock("", []byte{})) {
    returnValue = fmt.Sprintf("{\n\t\"error\" : \"NotReadableBlockIDNotExists\",\n\t\"hash\" : \"%s\"\n}", blockId)
  } else {
    returnValue = fmt.Sprintf(
      "{\n\t\"hash\" : \"%x\",\n\t\"selfLink\" : \"/api/%s/%x\",\n\t\"modified\" : \"%v\"\n,\n\t\"previousBlock\" : \"%x\"\n\t\"data\" : \"%s\"\n}",
      thisBlock.Hash,
      stackName,
      thisBlock.Timestamp,
      thisBlock.Hash,
      thisBlock.PrevBlockHash,
      thisBlock.Data)
  }

  fmt.Fprintln(
    w,
    returnValue)
}

func getBlockRaw(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  stackName := vars["stackName"]
  blockId := vars["blockId"]

  returnValue := ""
  _ = returnValue
  var thisBlock *Block = NewBlock("", []byte{} )

  if _, ok := stackMap[stackName]; ok {
    for _, block := range stackMap[stackName].blocks {
      if( fmt.Sprintf("%x", block.Hash) == blockId ) {
        thisBlock = block
      }
    }
  } else {
    returnValue = fmt.Sprintf("{\n\t\"error\" : \"NotReadableStackNotExists\",\n\t\"name\" : \"%s\"\n}", stackName)
  }

  if(thisBlock == NewBlock("", []byte{} )) {
    returnValue = fmt.Sprintf("{\n\t\"error\" : \"NotReadableBlockIDNotExists\",\n\t\"hash\" : \"%s\"\n}", blockId)
  } else {
    returnValue = fmt.Sprintf(
      "%s",
      thisBlock.Data)
  }

  fmt.Fprintln(w, returnValue)
}

func getBlockList(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  stackName := vars["stackName"]

  returnValue := "{\n"
  _ = returnValue

  if _, ok := stackMap[stackName]; ok {
    for _, block := range stackMap[stackName].blocks {
        returnValue += fmt.Sprintf("\t\"%x\" : {\n\t\t\"hash\" : \"%x\",\n\t\t\"timestamp\" : \"%v\",\n\t\t\"size\" : \"%v\",\n\t\t\"selfLink\" : \"/api/%s/%x\"\n\t},\n",
          block.Hash,
          block.Hash,
          block.Timestamp,
          cap(block.Data),
          stackName,
          block.Hash)
    }
    returnValue += fmt.Sprintf("\n\t\"stack\" : \"%s\"\n}", stackName)
  } else {
    returnValue = fmt.Sprintf("{\n\t\"error\" : \"NotReadableStackNotExists\",\n\t\"name\" : \"%s\"\n}", stackName)
  }

  fmt.Fprintln(w, returnValue)
}
