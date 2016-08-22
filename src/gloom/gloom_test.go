package gloom

import (
  "testing"
  "math/rand"
  "time"
)

func RandString(n int) string {
  alphab := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123421;-=+/.,<>:[]{}()-_~`!@#$%^&*?")
  b := make([]rune, n)
  for i := range b {
      b[i] = alphab[rand.Intn(len(alphab))]
  }
  return string(b)
}

func FilterTest(size int, strsize int, fpos float64) bool {
  bf := NewFilter(size,fpos)
  strs := make([]string,size)
  rand.Seed(time.Now().UnixNano())
  for i := 0; i < size; i++ {
    str := RandString(rand.Intn(strsize))
    strs[i] = str
    bf.PutString(str)
  }
  for i := 0; i < size; i++ {
    if !bf.HasString(strs[i]) {
      return false
    }
  }
  return true
}

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

func TestFilter(t *testing.T) {
  if !FilterTest(100,10,0.1) {
    t.Error("100 insertion, 10 max string size, 0.1fpos test failed!")
  }
  if !FilterTest(1000, 35, 0.01) {
    t.Error("1k insertion, 35 max string size, 0.01fpos test failed!")
  }
  if !FilterTest(10000, 50, 0.001) {
    t.Error("10k, 50 max string size, 0.001 fpos test failed!")
  }
  if !FilterTest(100000, 100, 0.0001) {
    t.Error("100k, 100 max string size, 0.0001 fpos test failed!")
  }
}
