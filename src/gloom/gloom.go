package gloom


import (
  "fmt"
  "hash"
  "hash/fnv"
  "math"
)

type BloomFilter struct {
   m             int //number of chunks in filter
   k             int //number of hash functions
   size          int //number of elements to be inserted in set 
   chunks        []uint64
   hash          hash.Hash64
}

func NewFilter(size int, fpos float64) *BloomFilter {
  if size < 1 || fpos <= 0.0 || fpos >= 1.0 {
    panic("Invalid parameters for Bloom Filter! Size must be > 1 and 0.0 < fpos < 1.0!")
  }
  m := GetOptimalM(size, fpos)
  k := GetOptimalK(m, size)
  numchunks := m / 64
  if m % 64 > 0 {
    numchunks++
  }
  chunks := make([]uint64, numchunks)
  return &BloomFilter{m,k,size,chunks,fnv.New64a()}
}

func (bf *BloomFilter) EstimateFalsePos() float64 {
  return math.Pow(1.0 - math.Exp((-1.0 * float64(bf.k * bf.size)) / float64(bf.m)),  float64(bf.k))
}

func GetOptimalM (size int, fpos float64) int {
  c := math.Pow(math.Log(2.0),2.0)
  return int(math.Ceil((float64(-1 * size) * math.Log(fpos)) / c))
}

func GetOptimalK (m int, size int) int {
  return int(math.Ceil(float64(m) / float64(size) * math.Log(2.0)))
}

func (bf * BloomFilter) GetHashes(key []byte) []uint64 {
  bf.hash.Sum(key)
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
    if bf.chunks[chnk] ^ (1 << shft) == bf.chunks[chnk] {
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
    bf.chunks[chnk] |= (1 << shft)
  }
  return bf
}

func (bf *BloomFilter) PutString(key string) *BloomFilter {
  return bf.Put([]byte(key))
}

func (bf *BloomFilter) Print() {
  for _, b := range bf.chunks {
    fmt.Printf("%64b\n",b)
  }
}


