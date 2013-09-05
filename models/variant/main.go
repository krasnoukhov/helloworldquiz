package variant

import (
  "fmt"
  "io"
  "io/ioutil"
  "crypto/md5"
  "math/rand"
  // "errors"
  // "github.com/astaxie/beego"
  "github.com/kylelemons/go-gypsy/yaml"
)

var (
  Objects map[string]*Object
)

type Object struct {
  Key        string  `json:"-"`
  Name       string  `json:"-"`
  Hash       string
  Snippet    string
  Variants   [3]string
}

func init() {
  Objects = make(map[string]*Object)
  file, _ := yaml.ReadFile("models/variant/data.yml")
  
  for key, node := range file.Root.(yaml.Map) {
    snippet, err := ioutil.ReadFile(fmt.Sprintf("models/variant/data/%s", key))
    
    if err == nil {
      name := node.(yaml.Map)["name"].(yaml.Scalar).String()
      
      hash := md5.New()
      io.WriteString(hash, key)
        
      variants := [3]string{}
      variants[0] = key
      for idx, variant := range node.(yaml.Map)["variants"].(yaml.List) {
        variants[idx+1] = variant.(yaml.Scalar).String()
      }
      
      Objects[key] = &Object{ key, name, fmt.Sprintf("%x", hash.Sum(nil)), string(snippet[:]), variants }
    }
  }
}

func Get(originalObject *Object) (object *Object) {
  object = originalObject
  
  for i := range object.Variants {
    j := rand.Intn(i + 1)
    object.Variants[i], object.Variants[j] = object.Variants[j], object.Variants[i]
  }
  
  return object
}
