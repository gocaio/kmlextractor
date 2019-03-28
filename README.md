# kmlExtractor

[![Go Report Card](https://goreportcard.com/badge/github.com/kevinborras/kmlextractor)](https://goreportcard.com/badge/github.com/kevinborras/kmlextractor)

KML data extractor written in Go.  

## How to use

```golang
package main

import (
    "fmt"
    "github.com/kevinborras/kmlextractor"
    "log"
    "os"
)

package main

import (
    "fmt"
    "github.com/kevinborras/kmlextractor"
    "log"
    "os"
    "reflect"
)

func main() {
    file, err := os.Open(`../samples/kml_3.kml`)
    if err != nil {
        log.Fatal(err)
    }
    content, err := kmlextractor.GetContent(file)
    if err != nil {
        log.Fatal(err)
    }

    s := reflect.ValueOf(&content).Elem()
    typeOfContent := s.Type()
    for i := 0; i < s.NumField(); i++ {
        f := s.Field(i)
        fmt.Printf("%d: %s = %v\n", i,
            typeOfContent.Field(i).Name, f.Interface())
    }
}
```

Output

```bash
0: Doc = { {[]} [] [] {180.0 0.0 1.5768778343115222E7 1.2742E7 0.0 0.0 clampToGround} {A sample ArcGrid file 1 1 {-1.4210854715202004E-14 1.4210854715202004E-14 1398874.1103087065 1130363.654413079 0.0 0.0 clampToGround}  {http://localhost:8080/geoserver/nurc/wms?service=wms&request=GetMap&version=1.1.1&format=application/vnd.google-earth.kml+xml&layers=nurc:Arc_Sample&styles=rain&height=768&width=768&transparent=false&srs=EPSG:4326 0.0 onStop 1.0 1.0}}}
1: Placemarks = []
2: F = {  0 [] {      } [] []}
3: NetLink = { 0 0 {      }  {    }}
```

## References

[KML Reference](http://dh.obdurodon.org/kml/kml-tutorial.xhtml)