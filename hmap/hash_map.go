package hmap

import "hash/crc32"

type KeyType interface{}
type ValueType interface{}

type bucket struct {
	keys   [8]KeyType
	values [8]ValueType
	filled uint32
}

type HashMap struct {
	buckets []bucket
	hash    func(KeyType) uint32
}

func New(hash func(KeyType) uint32) *HashMap {
	return &HashMap{
		buckets: make([]bucket, 2),
		hash:    hash,
	}
}

func (m *HashMap) rehash() {
	oldBuckets := m.buckets
	m.buckets = make([]bucket, len(oldBuckets)*2)
	for bucketIndex := 0; bucketIndex != len(oldBuckets); bucketIndex++ {
		filled := oldBuckets[bucketIndex].filled
		for i := uint32(0); i != filled; i++ {
			m.Insert(oldBuckets[bucketIndex].keys[i], oldBuckets[bucketIndex].values[i])
		}
	}
}

func (m *HashMap) Insert(key KeyType, value ValueType) {
	bucketIndex := m.hash(key) % uint32(len(m.buckets))
	filled := m.buckets[bucketIndex].filled
	for i := uint32(0); i != filled; i++ {
		if m.buckets[bucketIndex].keys[i] == key {
			m.buckets[bucketIndex].values[i] = value
			return
		}
	}

	if filled != 8 {
		m.buckets[bucketIndex].keys[filled] = key
		m.buckets[bucketIndex].values[filled] = value
		m.buckets[bucketIndex].filled++
		return
	}

	m.rehash()
	m.Insert(key, value)
}

func (m *HashMap) Get(key KeyType) (ValueType, bool) {
	bucketIndex := m.hash(key) % uint32(len(m.buckets))
	for i := uint32(0); i != m.buckets[bucketIndex].filled; i++ {
		if m.buckets[bucketIndex].keys[i] == key {
			return m.buckets[bucketIndex].values[i], true
		}
	}

	return nil, false
}

func (m *HashMap) Remove(key KeyType) {
	bucketIndex := m.hash(key) % uint32(len(m.buckets))
	filled := m.buckets[bucketIndex].filled
	for i := uint32(0); i != filled; i++ {
		if m.buckets[bucketIndex].keys[i] == key {
			m.buckets[bucketIndex].keys[i] = m.buckets[bucketIndex].keys[filled-1]
			m.buckets[bucketIndex].values[i] = m.buckets[bucketIndex].keys[filled-1]
			m.buckets[bucketIndex].filled--
		}
	}

}

func HashString(value KeyType) uint32 {
	s := value.(string)
	return crc32.ChecksumIEEE([]byte(s))
}

func HashUint32(value KeyType) uint32 {
	return value.(uint32)
}
