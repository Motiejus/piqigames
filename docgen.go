package main

import (
    "io/ioutil"
    "log"
    "code.google.com/p/goprotobuf/proto"

    "./definition"
)

func main() {
    data, err := ioutil.ReadFile("definition/demo.piqi.pb")
    if err != nil {
        panic(err)
    }

    newTest := &piqi_doc_piqi.Piqi{}
    err = proto.Unmarshal(data, newTest)
    if err != nil {
        log.Fatal("unmarshaling error: ", err)
    }
}
