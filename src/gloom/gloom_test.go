package gloom

import (
  "testing"
)

func TestSimple(t *testing.T) {
  bf := NewFilter(128,3)
  bf.PutString("hello")
  bf.PutString("goodbye")
  if !bf.HasString("hello") {
    t.Error("hello not in filter!")
  }
  if !bf.HasString("goodbye") {
    t.Error("goodbye not in filter!")
  } 
}
