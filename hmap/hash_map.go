package hmap

import "hash/crc32"

type KeyType interface{}
type ValueType interface{}

type item struct {
	key       KeyType
	value     ValueType
	hashValue uint32
}

type bucket struct {
	items  [8]item
	filled uint32
}

type HashMap struct {
	buckets    []bucket
	hash       func(KeyType) uint32
	bucketMask uint32
}

func New(hash func(KeyType) uint32) *HashMap {
	bucketsCount := uint32(2)
	return &HashMap{
		buckets:    make([]bucket, bucketsCount),
		hash:       hash,
		bucketMask: bucketsCount - 1,
	}
}

func (m *HashMap) rehash() {
	oldBuckets := m.buckets
	oldLen := len(oldBuckets)
	m.buckets = make([]bucket, oldLen<<1)
	m.bucketMask = uint32(len(m.buckets) - 1)
	for bucketIndex := 0; bucketIndex != oldLen; bucketIndex++ {
		filled := oldBuckets[bucketIndex].filled
		for i := uint32(0); i != filled; i++ {
			_ = m.insert(oldBuckets[bucketIndex].items[i])
		}
	}
}

func (m *HashMap) insert(element item) bool {
	bucketIndex := element.hashValue & m.bucketMask
	filled := m.buckets[bucketIndex].filled
	for i := uint32(0); i != filled; i++ {
		if m.buckets[bucketIndex].items[i].key == element.key {
			m.buckets[bucketIndex].items[i].value = element.value
			return true
		}
	}

	if filled != 8 {
		m.buckets[bucketIndex].items[filled] = element
		m.buckets[bucketIndex].filled++
		return true
	}

	return false
}

func (m *HashMap) Insert(key KeyType, value ValueType) {
	element := item{
		key:       key,
		value:     value,
		hashValue: m.hash(key),
	}
	if !m.insert(element) {
		m.rehash()
		m.insert(element)
	}
}

func (m *HashMap) Get(key KeyType) (ValueType, bool) {
	bucketIndex := m.hash(key) & m.bucketMask
	for i := uint32(0); i != m.buckets[bucketIndex].filled; i++ {
		if m.buckets[bucketIndex].items[i].key == key {
			return m.buckets[bucketIndex].items[i].value, true
		}
	}

	return nil, false
}

func (m *HashMap) Remove(key KeyType) {
	bucketIndex := m.hash(key) & m.bucketMask
	filled := m.buckets[bucketIndex].filled
	for i := uint32(0); i != filled; i++ {
		if m.buckets[bucketIndex].items[i].key == key {
			m.buckets[bucketIndex].items[i] = m.buckets[bucketIndex].items[filled-1]
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
