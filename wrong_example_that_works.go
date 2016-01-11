package main

import (
    "fmt"
  )

type Htmlize interface {
    ToHtml() string
}

func main() {
    var h Htmlize

    sl := []int{1,2,3}
    ms := MySlice{s:sl}

    h = &ms
    fmt.Println(h.ToHtml())
}

type MySlice struct {
    s []int
}

func (ms *MySlice) ToHtml() string {
  return fmt.Sprintf("<pre>%v</pre>",  ms.s)
}
