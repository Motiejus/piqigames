package main

import (
    "io/ioutil"
    "strings"
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
    var builtin = func(value string)(bool) {
        var _, ok = builtins[value]
        return ok
    }

    var fqtype = func(currentmod, usertype string) (string, string) {
        /* Return (module, type) or ("", type) if type is local */
        var split = strings.Split(usertype, "/")
        if len(split) == 2 {
            return split[0], split[1]
        } else if builtin(split[0]) {
            return "piqi", split[0]
        } else {
            return currentmod, split[0]
        }
    }

    var type2type = func(usertype string) string {
        var split = strings.Split(usertype, "/")
        if len(split) == 2 {
            return split[1]
        } else {
            return usertype
        }
    }

    var type2mod = func(currentmod, usertype string) string {
        var mod, _ = fqtype(currentmod, usertype)
        return mod
    }

    var hreftype = func(usermod string, usertype *string) template.HTML {
        if usertype == nil {
            return "&lt;nil&gt;"
        }
        var tpl = `<a href="#module_%s_%s">%s/%s</a>`
        var mod = type2mod(usermod, *usertype)
        var t = type2type(*usertype)
        return template.HTML(fmt.Sprintf(tpl, mod, t, mod, t))
    }

    funcmap := template.FuncMap{
        "type2mod": type2mod,
        "type2type": type2type,
        "hreftype": hreftype,
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
