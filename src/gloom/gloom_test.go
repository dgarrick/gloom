package gloom

import (
  "testing"
  "fmt"
)

func Test(t *testing.T) {
  bf := NewFilter(300,0.001)
  fmt.Println(bf.k)
  fmt.Println(bf.m)
  bf.PutString("hello")
  bf.PutString("goodbye")
  if !bf.HasString("hello") {
    t.Error("hello not in filter!")
  }
  if !bf.HasString("goodbye") {
    t.Error("goodbye not in filter!")
  }
  fmt.Print("Optimal test false pos est: ")
  fmt.Println(bf.EstimateFalsePos())
  
}
