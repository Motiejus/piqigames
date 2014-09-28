package main

import (
    "io/ioutil"
    "os"
    "log"
    "flag"
    "code.google.com/p/goprotobuf/proto"

    "./definition"
    "./templates"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func main() {
    var in = flag.String("in", "/dev/stdin", "input file (protobuf encoded)")
    var out = flag.String("out", "/dev/stdout", "output HTML")
    flag.Parse()

    data, err := ioutil.ReadFile(*in)
    check(err)

    f, err := os.Create(*out)
    check(err)
    defer f.Close()

    piqiL := &piqi_doc_piqi.PiqiList{}
    err = proto.Unmarshal(data, piqiL)
    if err != nil {
        log.Fatal("unmarshaling error: ", err)
    }

    err = templates.Details.Execute(f, piqiL.Piqi)
    check(err)
}
