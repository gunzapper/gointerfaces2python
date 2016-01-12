package main

import (
    "fmt"
    "github.com/deckarep/golang-set" // mapset -> I need to use sets
)

type Htmlize interface {
    ToHtml() string
}

func main() {
    var h Htmlize

    set := mapset.NewSet()
    set.Add(1)
    set.Add(2)
    set.Add(3)

    fmt.Println(set)

    ms := MySet{s:mapset.NewSet()}
    ms.s.Add(1)
    ms.s.Add(2)
    ms.s.Add(3)

    h = &ms
    fmt.Println(h.ToHtml)

}

type MySet struct {
   s mapset.Set
}

func (ms *MySet) ToHtml() string {
  return fmt.Sprintf("<pre>%v</pre>", ms.s) // why is set an adress?
}

