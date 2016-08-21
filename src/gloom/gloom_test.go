package gloom

import (
  "testing"
)

func TestSimple(t *testing.T) {
  bf := NewFilter(300,0.001)
  if(bf.m != 4314) {
    t.Error("M is not optimal!")
  }
  if(bf.k != 10) {
    t.Error("K is not optimal!")
  }
  bf.PutString("hello")
  bf.PutString("goodbye")
  if !bf.HasString("hello") {
    t.Error("hello not in filter!")
  }
  if !bf.HasString("goodbye") {
    t.Error("goodbye not in filter!")
  }
  if(bf.EstimateFalsePos() > 0.001) {
    t.Error("false pos estimate is not optimal!")
  }
}

