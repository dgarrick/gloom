# gloom [![Build Status](https://travis-ci.org/dgarrick/gloom.svg?branch=master)](https://travis-ci.org/dgarrick/gloom)
Simple bloom filter implemented in golang

## usage

 ```go
 var size int = 100 //the number of items to be inserted into the filter
 var fpos float64 = 0.1 //the acceptable percentage of lookups which can return a false positive
 bf, err := NewFilter(size,fpos) //size must be > 0 and 0.0 < fpos < 1.0
 if err != nil {
   ...
 }
 bf.PutString("hello").PutString("world") //returns a pointer to the filter object
 bf.HasString("hello") //true
 bf.HasString("world") //true
 ```
