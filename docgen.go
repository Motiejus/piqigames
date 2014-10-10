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

func get_builtins(piqiSelf *piqi_doc_piqi.Piqi) map[string]bool {
  builtins := make(map[string]bool)
  for _, typedef := range piqiSelf.PiqiTypedef {
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

    var hreftype = func(usermod string, usertype *string) template.HTML {
        if usertype == nil {
            return "&lt;nil&gt;"
        }
        var tpl = `<a href="#module_%s_%s">%s/%s</a>`
        var mod, _ = fqtype(usermod, *usertype)
        var t = type2type(*usertype)
        return template.HTML(fmt.Sprintf(tpl, mod, t, mod, t))
    }

    var nameof = func(td *piqi_doc_piqi.PiqiTypedef) string {
      if td.Record != nil {
        return *td.Record.Name
      } else if td.Variant != nil {
        return *td.Variant.Name
      } else if td.PiqiEnum != nil {
        return *td.PiqiEnum.Name
      } else if td.Alias != nil {
        return *td.Alias.Name
      } else if td.List != nil {
        return *td.List.Name
      } else {
         return "unknown"
      }
    }

    funcmap := template.FuncMap{
        "type2type": type2type,
        "hreftype": hreftype,
        "nameof": nameof,
    }
    return template.New("module").Funcs(funcmap).Parse(templates.Module)
}

func main() {
    var in = flag.String("in", "/dev/stdin", "input file (protobuf encoded)")
    var out = flag.String("out", "/dev/stdout", "output HTML")
    var selfspec = flag.String("selfspec", "", "self-spec (pb)")
    flag.Parse()

    data, err := ioutil.ReadFile(*in)
    check(err)
    dataSelf, err := ioutil.ReadFile(*selfspec)
    check(err)

    piqiSelf := &piqi_doc_piqi.Piqi{}
    piqiL := &piqi_doc_piqi.PiqiList{}

    err = proto.Unmarshal(data, piqiL)
    if err != nil {
        log.Fatal("unmarshaling piqiL error: ", err)
    }
    err = proto.Unmarshal(dataSelf, piqiSelf)
    if err != nil {
        log.Fatal("unmarshaling piqiL error: ", err)
    }

    var piqiWithSelfSpec = append(piqiL.Piqi, piqiSelf)
    var builtins = get_builtins(piqiSelf)
    var tpl = template.Must(get_tpl(builtins))

    f, err := os.Create(*out)
    check(err)
    defer f.Close()
    tpl.Execute(f, piqiWithSelfSpec)
}
