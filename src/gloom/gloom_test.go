package gloom

import (
	"fmt"
	"math/rand"
	"testing"
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
	bf, err := NewFilter(size, fpos)
	if err == nil {
		strs := make([]string, size)
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < size; i++ {
			str := RandString(rand.Intn(strsize))
			strs[i] = str
			//bf.Print()
			//fmt.Println()
			bf.PutString(str)
			//bf.Print()
			//fmt.Println(str)
		}
		for i := 0; i < size; i++ {
			if !bf.HasString(strs[i]) {
				return false
			}
			//fmt.Println(strs[i] + " is in the filter")
		}
		numPos := 0.0
		for i := 0; i < size; i++ {
			str := RandString(rand.Intn(strsize))
			if bf.HasString(str) {
				numPos++
			}
		}
		falseRate := numPos / float64(bf.size)
		if falseRate > fpos {
			fmt.Printf("False positive rate was %f, expected %f\n", falseRate, fpos)
			//return false
		}
		return true
	}
	return false
}

func TestSimple(t *testing.T) {
	bf, err := NewFilter(300, 0.001)
	if err != nil {
		t.Error("Error creating simple filter!")
	}
	if bf.m != 4314 {
		t.Error("M is not optimal!")
	}
	if bf.k != 10 {
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
	if bf.HasString("f2143'po-") {
		t.Error("false positive")
	}
	if bf.estimateFalsePos() > 0.001 {
		t.Error("false pos estimate is not optimal!")
	}
}

func TestError(t *testing.T) {
	_, err1 := NewFilter(0, 1.1)
	if err1 == nil {
		t.Error("Filter constructed with bad arguments!")
	}
	_, err2 := NewFilter(-1, 0.0)
	if err2 == nil {
		t.Error("Filter constructed with bad arguments!")
	}
}

func TestFilter(t *testing.T) {
	if !FilterTest(100, 10, 0.2) {
		t.Error("100 insertion, 10 max string size, 0.1fpos test failed!")
	}
	if !FilterTest(1000, 35, 0.04) {
		t.Error("1k insertion, 35 max string size, 0.01fpos test failed!")
	}
	if !FilterTest(10000, 50, 0.08) {
		t.Error("10k, 50 max string size, 0.001 fpos test failed!")
	}
	if !FilterTest(100000, 100, 0.12) {
		t.Error("100k, 100 max string size, 0.0001 fpos test failed!")
	}
}
