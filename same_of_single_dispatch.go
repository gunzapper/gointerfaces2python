package main

import (
    "fmt"
    "github.com/deckarep/golang-set" // mapset -> I need to use sets
    "math"
    "reflect" // to find the name of a function
    "runtime" // used together with this library
    "html" // to have & in html
    "strings" // to replace in strings
)

type Htmlize interface {
    ToHtml() string
}

func main() {
    var h Htmlize

    /*set := mapset.NewSet()
    set.Add(1)
    set.Add(2)
    set.Add(3)

    fmt.Println(set)*/

    ms := MySet{s:mapset.NewSet()}
    ms.s.Add(1)
    ms.s.Add(2)
    ms.s.Add(3)

    h = &ms
    fmt.Println(h.ToHtml())

    mf := MyFunc(math.Abs)

    h = mf
    fmt.Println(h.ToHtml())

    //ms := MyString("Heimlich & Co.\n- a game")
    // go is strong typed
    m_s := MyString("Heimlich & Co.\n- a game")
    h = m_s
    fmt.Println(h.ToHtml())

    // there is no way to make a generic heterogenous sequence
    // I have to use struct
    sub_e := make([]MyInt, 3)
    sub_e[0] = 3
    sub_e[1] = 2
    sub_e[2] = 1
    mos := MyOwnStruct{
      a:MyString("alpha"),
      i:MyInt(66),
      si:MyIntSeq{s:sub_e}}
    h = mos
    fmt.Println(h.ToHtml())
}

type MySet struct {
   s mapset.Set
}

func (v *MySet) ToHtml() string {
  return fmt.Sprintf("<pre>%v</pre>", v.s) // why is set an adress?
}

type MyFunc func(float64) float64

func (v MyFunc) ToHtml() string {
    //return fmt.Sprintf("<pre>%v</pre>", mf)
    // it puts something like <pre>0x473a70</pre>
    // looking on sof question 7052693
    name := runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name()
    return fmt.Sprintf("<pre>%s</pre>", name)
}

type MyString string

func (v MyString) ToHtml() string {
    return fmt.Sprintf("<pre>%s</pre>",
      strings.Replace(html.EscapeString(string(v)), "\n", "\n<br>", -1))
}

type MyInt int

func (i MyInt) ToHtml() string {
    return fmt.Sprintf("<pre>%v<pre>", i)
}

type MyIntSeq struct {
  s []MyInt
}

func (si MyIntSeq) ToHtml() string {
    return fmt.Sprintf("<pre>%v</pre>", si)
}

type MyOwnStruct struct {
  a MyString
  i MyInt
  si MyIntSeq
}

func (mos MyOwnStruct) ToHtml() string {
    out_s := ""
    /*
    v := reflect.ValueOf(mos)
    for i := 0; i < v.NumField(); i++ {
        e := v.Field(i).Interface()
        out_s = out_s + fmt.Sprintf("/n<li>%s</li>", e.ToHtml())
    }
    */
    out_s += fmt.Sprintf("\n<li>%s</li>", mos.a.ToHtml())
    out_s += fmt.Sprintf("\n<li>%v</li>", mos.i.ToHtml())
    out_s += fmt.Sprintf("\n<li>%v</li>", mos.si.ToHtml())
    return fmt.Sprintf("<ul>%s\n</ul>", out_s)
}
