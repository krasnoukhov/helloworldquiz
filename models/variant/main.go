package variant

import (
  "fmt"
  // "io"
  "io/ioutil"
  // "crypto/md5"
  "math/rand"
  "html"
  // "errors"
  // "github.com/astaxie/beego"
  "github.com/kylelemons/go-gypsy/yaml"
)

var (
  Objects    map[string]*Object
  Keys       []string
)

type Object struct {
  Key        string     `json:"-"`
  Name       string     `json:"-"`
  Snippet    string
  Variants   []string   `json:"-"`
  Options    []*Option
}

type DumpObject struct {
  Key        string
  Name       string
  Snippet    string
}

type Option struct {
  Key        string
  Name       string
}

func init() {
  Objects = make(map[string]*Object)
  Keys = []string{}
  file, _ := yaml.ReadFile("models/variant/data.yml")
  
  for key, node := range file.Root.(yaml.Map) {
    snippet, err := ioutil.ReadFile(fmt.Sprintf("models/variant/data/%s", key))
    
    if err == nil {
      name := node.(yaml.Map)["name"].(yaml.Scalar).String()
      
      variants := []string{}
      variants = append(variants, key)
      for _, variant := range node.(yaml.Map)["variants"].(yaml.List) {
        variants = append(variants, variant.(yaml.Scalar).String())
      }
      
      Objects[key] = &Object{ key, name, html.EscapeString(string(snippet[:])), variants, []*Option{} }
      Keys = append(Keys, key)
    }
  }
}

func ConvertToDumpObject(object *Object) (response *DumpObject) {
  return &DumpObject{ object.Key, object.Name, object.Snippet }
}

func Shuffle(object *Object) (response *Object) {
  response = &Object{ object.Key, object.Name, object.Snippet, append([]string{}, object.Variants...), []*Option{} }
  
  for i := range response.Variants {
    j := rand.Intn(i + 1)
    response.Variants[i], response.Variants[j] = response.Variants[j], response.Variants[i]
  }
  
  for _, variant := range response.Variants {
    response.Options = append(response.Options, &Option{ Objects[variant].Key, Objects[variant].Name })
  }
  
  return response
}
