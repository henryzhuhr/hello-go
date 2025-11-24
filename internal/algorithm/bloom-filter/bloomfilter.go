package bloomfilter

import (
	"hash"
	"hash/fnv"
	"math"
)

type BloomFilter[T any] interface {
	// Add 添加元素
	Add(data T)

	// Contains 检查元素是否可能存在于布隆过滤器中
	Contains(data T) bool
}

// bloomFilter 布隆过滤器
type bloomFilter[T any] struct {
	bitSet  []uint64       // 位数组
	m       uint64         // 位数组大小 (bits)
	k       uint64         // 哈希函数数量
	hashFn1 hash.Hash64    // 哈希函数 1
	hashFn2 hash.Hash64    // 哈希函数 2
	keyMap  func(T) []byte // 将 T 转换为 []byte 的函数
}

// NewBloomFilter 创建一个新的布隆过滤器
// n: 预期元素数量
// p: 期望误判率 (0 < p < 1)
// keyMap: 将泛型 T 映射为 []byte 的函数
func NewBloomFilter[T any](n uint64, p float64, keyMap func(T) []byte) BloomFilter[T] {
	if n == 0 || p <= 0 || p >= 1 || keyMap == nil {
		return nil
	}

	// 计算最佳位数组大小 m
	// m = - (n * ln(p)) / (ln(2)^2)
	m := -float64(n) * math.Log(p) / math.Pow(math.Log(2), 2)
	mUint := uint64(math.Ceil(m))

	// 计算最佳哈希函数数量 k
	// k = (m / n) * ln(2)
	k := (m / float64(n)) * math.Log(2)
	kUint := uint64(math.Ceil(k))

	return &bloomFilter[T]{
		bitSet:  make([]uint64, (mUint+63)/64),
		m:       mUint,
		k:       kUint,
		hashFn1: fnv.New64a(),
		hashFn2: fnv.New64(),
		keyMap:  keyMap,
	}
}

// Add implements [BloomFilter.Add]
func (bf *bloomFilter[T]) Add(data T) {
	dataBytes := bf.keyMap(data)
	h1, h2 := bf.getHash(dataBytes)
	for i := uint64(0); i < bf.k; i++ {
		// 使用双重哈希模拟 k 个哈希函数
		// g_i(x) = (h1(x) + i * h2(x)) % m
		pos := (h1 + i*h2) % bf.m
		bf.setBit(pos)
	}
}

// Contains implements [BloomFilter.Contains]
func (bf *bloomFilter[T]) Contains(data T) bool {
	dataBytes := bf.keyMap(data)
	h1, h2 := bf.getHash(dataBytes)
	for i := uint64(0); i < bf.k; i++ {
		pos := (h1 + i*h2) % bf.m
		if !bf.testBit(pos) {
			return false
		}
	}
	return true
}

// getHash 计算两个基础哈希值
func (bf *bloomFilter[T]) getHash(data []byte) (uint64, uint64) {
	bf.hashFn1.Reset()
	bf.hashFn1.Write(data)
	h1 := bf.hashFn1.Sum64()

	bf.hashFn2.Reset()
	bf.hashFn2.Write(data)
	h2 := bf.hashFn2.Sum64()

	return h1, h2
}

// setBit 设置位
func (bf *bloomFilter[T]) setBit(pos uint64) {
	idx := pos / 64
	offset := pos % 64
	bf.bitSet[idx] |= (1 << offset)
}

// testBit 检查位
func (bf *bloomFilter[T]) testBit(pos uint64) bool {
	idx := pos / 64
	offset := pos % 64
	return (bf.bitSet[idx] & (1 << offset)) != 0
}
