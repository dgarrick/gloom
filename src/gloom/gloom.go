package gloom


import (
  "fmt"
  "hash"
  "hash/fnv"
)

type BloomFilter struct {
   m      int //number of bits in filter
   k      int //number of hash functions
   bits   []uint64
   hash   hash.Hash64
}

func NewFilter(m int, k int) *BloomFilter {
  size := m / 64
  if m % 64 > 0 {
    size++
  }
  bits := make([]uint64, size)
  return &BloomFilter{m,k,bits,fnv.New64a()}
}

func (bf * BloomFilter) GetHashes(key []byte) []uint64 {
  bf.hash.Reset()
  bf.hash.Write(key)
  h := bf.hash.Sum64()
  h1 := uint32(h)
  h2 := uint32(h >> 32) 
  hs := make([]uint64,bf.k)
  for i := 0; i < bf.k; i++ {
   //This distributes values so that one hash function
   //can effectively act as k hash functions 
   //https://www.eecs.harvard.edu/~michaelm/postscripts/rsa2008.pdf
   hs[i] = uint64(h1 + uint32(i)*h2)
  }
  return hs
}

func (bf *BloomFilter) HashLoc(h uint64) (chnk uint64,shft uint64) {
    bitInd := h % uint64(bf.m)
    chnkInd := bitInd / 64
    shftInd := bitInd - chnkInd*64
    return chnkInd, shftInd
}

func (bf *BloomFilter) Has(data []byte) bool {
  hs := bf.GetHashes(data)
  for i := 0; i < bf.k; i++ {
    hk := hs[i]
    chnk, shft := bf.HashLoc(hk)
    if bf.bits[chnk] ^ (1 << shft) == bf.bits[chnk] {
      return false
    }
  }
  return true
}

func (bf *BloomFilter) HasString(key string) bool {
  return bf.Has([]byte(key))
}

func (bf *BloomFilter) Put(data []byte) *BloomFilter {
  hs := bf.GetHashes(data)
  for i := 0; i < bf.k; i++ {
    hk := hs[i]
    chnk, shft := bf.HashLoc(hk)
    bf.bits[chnk] |= (1 << shft)
  }
  return bf
}

func (bf *BloomFilter) PutString(key string) *BloomFilter {
  return bf.Put([]byte(key))
}

func (bf *BloomFilter) Print() {
  for _, b := range bf.bits {
    fmt.Printf("%64b\n",b)
  }
}


