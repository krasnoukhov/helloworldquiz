package variant

import (
  "fmt"
  "io/ioutil"
  // "errors"
  // "github.com/astaxie/beego"
  "github.com/kylelemons/go-gypsy/yaml"
)

var (
  Objects map[string]*Object
)

type Object struct {
  Key        string
  Name       string
  Snippet    string
  Variants   [2]string
}

func init() {
  Objects = make(map[string]*Object)
  file, _ := yaml.ReadFile("models/variant/data.yml")
  
  for key, node := range file.Root.(yaml.Map) {
    snippet, err := ioutil.ReadFile(fmt.Sprintf("models/variant/data/%s", key))
    
    if err == nil {
      name := node.(yaml.Map)["name"].(yaml.Scalar).String()
      variants := [2]string{}
      for idx, variant := range node.(yaml.Map)["variants"].(yaml.List) {
        variants[idx] = variant.(yaml.Scalar).String()
      }
      
      Objects[key] = &Object{ key, name, string(snippet[:]), variants }
    }
  }
}
