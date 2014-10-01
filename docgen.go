package main

import (
    "io/ioutil"
    "fmt"
    "os"
    "log"
    "flag"
    "html/template"
    "code.google.com/p/goprotobuf/proto"

    "./definition"
    "./templates"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func get_builtins(piqiL *piqi_doc_piqi.PiqiList) map[string]bool {
  builtins := make(map[string]bool)
  selfspec := &piqi_doc_piqi.Piqi{}
  for _, piqi := range piqiL.Piqi {
    if *piqi.Module == "piqi" {
      selfspec = piqi
      break
    }
  }

  for _, typedef := range selfspec.PiqiTypedef {
    if typedef.Alias != nil && typedef.Alias.PiqiType != nil {
      builtins[*typedef.Alias.Name] = true
    }
  }
  return builtins
}

func get_tpl(builtins map[string]bool) (*template.Template, error) {
  funcmap := template.FuncMap{
    "builtin": func(value string)(bool) {
      var _, ok = builtins[value]
      fmt.Printf("Type: %s, builtin: %v\n", value, ok)
      return ok
    },
  }
  return template.New("module").Funcs(funcmap).Parse(templates.Module)
}

func main() {
    var in = flag.String("in", "/dev/stdin", "input file (protobuf encoded)")
    var out = flag.String("out", "/dev/stdout", "output HTML")
    flag.Parse()

    data, err := ioutil.ReadFile(*in)
    check(err)

    piqiL := &piqi_doc_piqi.PiqiList{}
    err = proto.Unmarshal(data, piqiL)
    if err != nil {
        log.Fatal("unmarshaling error: ", err)
    }

    var builtins = get_builtins(piqiL)
    var tpl = template.Must(get_tpl(builtins))

    f, err := os.Create(*out)
    check(err)
    defer f.Close()
    tpl.Execute(f, piqiL.Piqi)
}
