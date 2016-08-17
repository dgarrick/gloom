package gloom

import (
  "hash/fnv"
  "hash"
  "fmt"
)

type BitSet []uint64
const FnvOffset = 14695981039346656037
const FnvPrime = 1099511628211
type BloomFilter struct {
   m      uint
   k      uint
   bits   BitSet
   hashes []uint64
}

func NewBloomFilter(m int, k int) *BloomFilter {
  bits := make(BitSet, m)
  for i := 0; i < k; i++ {
    hashes[i] = GetFnvHash
  }
  return &BloomFilter{m,k,bits}
}

func (bf *BloomFilter) PrintHashes() {
  for i := 0; i < bf.k; i++ {
    fmt.Print(bf.hashes[i])
    fmt.Println()
  }
}

func (bf *BloomFilter) Add(data []byte) *BloomFilter {
  
  for i := uint(0); i < k; i++ {
    bf.bits[
  }
}

//returns FNV64-1a hash for byte array
func GetFnvHash(data []byte) uint64 {
  hash := FnvOffset
  for _, byte := range data {
    hash ^=  uint64(byte)
    hash *= FnvPrime
  }
  return hash
}
