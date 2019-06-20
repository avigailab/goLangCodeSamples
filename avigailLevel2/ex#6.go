package main

import "fmt"

func main() {
 const (
  _ = iota
  a =   2019 + iota
  b =   2019 + iota
  c =   2019 + iota
  d =   2019 + iota
 )
 fmt.Println(a, b, c, d)
}
