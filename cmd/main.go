package main

import (
	"fmt"
  "github.com/takkyu2/pfds/list"
)

func main() {
  var a,b list.link[int];
  a.next = &b
  b.elem = 1
  fmt.Println(a,a.next)
}
