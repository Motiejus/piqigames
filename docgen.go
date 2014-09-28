package main

import (
    "io/ioutil"
    "fmt"
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
    var out = flag.String("out", "out", "output directory")
    flag.Parse()

    data, err := ioutil.ReadFile(*in)
    check(err)

    err = os.MkdirAll(*out, os.FileMode(0755))
    check(err)

    piqiL := &piqi_doc_piqi.PiqiList{}
    err = proto.Unmarshal(data, piqiL)

    if err != nil {
        log.Fatal("unmarshaling error: ", err)
    }

    for _, piqi := range piqiL.Piqi {
        f, err := os.Create(fmt.Sprintf("%s/%s.html", *out, *piqi.Module))
        check(err)
        defer f.Close()
        err = templates.Details.Execute(f, piqi)
        check(err)
    }

}
