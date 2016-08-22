# gloom
Simple bloom filter implemented in golang

## usage

 ```go
 bf := NewFilter(m,k) //where m is the number of bits allocated to the filter, and k is the number of hash functions.
 bf.PutString("hello world!").PutString("goodnight") //returns a pointer to the filter object
 bf.HasString("hello world!") //true
 bf.HasString("goodnight") //true
 ```
