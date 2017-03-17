package gloom

import (
	"errors"
	"fmt"
	"hash"
	"hash/fnv"
	"math"
)

type BloomFilter struct {
	m      int //number of bits in filter
	k      int //number of hash functions
	size   int //number of elements to be inserted in filter
	chunks []uint64
	hash   hash.Hash64
}

func NewFilter(size int, fpos float64) (*BloomFilter, error) {
	if size < 1 || fpos <= 0.0 || fpos >= 1.0 {
		return nil, errors.New("Bad arguments for bloom filter: size must be >= 1 and 0.0 < fpos < 1.0!")
	}
	m := optimalM(size, fpos)
	k := optimalK(m, size)
	numchunks := m / 64
	if m%64 > 0 {
		numchunks++
	}
	chunks := make([]uint64, numchunks)
	return &BloomFilter{m, k, size, chunks, fnv.New64a()}, nil
}

func (bf *BloomFilter) estimateFalsePos() float64 {
	return math.Pow(1.0-math.Exp((-1.0*float64(bf.k*bf.size))/float64(bf.m)), float64(bf.k))
}

func optimalM(size int, fpos float64) int {
	return int(math.Ceil((float64(-1*size) * math.Log(fpos)) / math.Pow(math.Log(2.0), 2.0)))
}

func optimalK(m int, size int) int {
	return int(math.Ceil(float64(m) / float64(size) * math.Log(2.0)))
}

func (bf *BloomFilter) getHashes(key []byte) []uint64 {
	bf.hash.Write(key)
	h := bf.hash.Sum64()
	hs := make([]uint64, bf.k)
	for i := 0; i < bf.k; i++ {
		//This distributes values so that one hash function
		//can effectively act as k hash functions
		//for more info, see: Kirsch and Mitzenmacher
		//full disclosure: I discovered this tidbit from the Bloom Filter wikipedia entry:
		//https://en.wikipedia.org/wiki/Bloom_filter#cite_ref-3
		hs[i] = uint64(uint32(h) + uint32(i)*uint32(h>>32))
	}
	bf.hash.Reset()
	return hs
}

func (bf *BloomFilter) hashLoc(h uint64) (chnk uint64, shft uint64) {
	bitInd := h % uint64(bf.m)
	chnkInd := bitInd / 64
	shftInd := h % 64
	return chnkInd, shftInd
}

func (bf *BloomFilter) Has(data []byte) bool {
	hs := bf.getHashes(data)
	for i := 0; i < bf.k; i++ {
		hk := hs[i]
		chnk, shft := bf.hashLoc(hk)
		if bf.chunks[chnk]&(1<<shft) == 0 {
			return false
		}
	}
	return true
}

func (bf *BloomFilter) HasString(key string) bool {
	return bf.Has([]byte(key))
}

func (bf *BloomFilter) Put(data []byte) *BloomFilter {
	hs := bf.getHashes(data)
	for i := 0; i < bf.k; i++ {
		hk := hs[i]
		chnk, shft := bf.hashLoc(hk)
		bf.chunks[chnk] |= (1 << shft)
	}
	return bf
}

func (bf *BloomFilter) PutString(key string) *BloomFilter {
	return bf.Put([]byte(key))
}

func (bf *BloomFilter) Print() {
	for _, b := range bf.chunks {
		fmt.Printf("%64b\n", b)
	}
}
