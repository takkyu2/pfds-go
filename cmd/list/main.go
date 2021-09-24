package main

import (
	"fmt"
)

type link[T any] struct {
  elem T
  next *link[T]
}
