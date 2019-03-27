package hmap

import "hash/crc32"

const (
	bucketMaxSizeShift = 3
)

type KeyType interface{}
type ValueType interface{}

type item struct {
	key       KeyType
	value     ValueType
	hashValue uint32
}

type HashMap struct {
	items       []item
	bucketsSize []uint32
	hash        func(KeyType) uint32
	bucketMask  uint32
}

func New(hash func(KeyType) uint32) *HashMap {
	bucketsCount := uint32(2)
	return &HashMap{
		items:       make([]item, bucketsCount<<bucketMaxSizeShift),
		bucketsSize: make([]uint32, bucketsCount),
		hash:        hash,
		bucketMask:  bucketsCount - 1,
	}
}

func (m *HashMap) rehash() {
	oldItems := m.items
	oldBucketsSize := m.bucketsSize
	oldIBucketsCount := m.bucketMask + 1

	newBucketsCount := oldIBucketsCount << 1
	m.items = make([]item, newBucketsCount<<bucketMaxSizeShift)
	m.bucketsSize = make([]uint32, newBucketsCount)
	m.bucketMask = newBucketsCount - 1

	for bucketIndex := uint32(0); bucketIndex != oldIBucketsCount; bucketIndex++ {
		bucketStartInd := bucketIndex << bucketMaxSizeShift
		bucketEndInd := bucketStartInd + oldBucketsSize[bucketIndex]

		for i := bucketStartInd; i != bucketEndInd; i++ {
			newBucketIndex := oldItems[i].hashValue & m.bucketMask
			index := newBucketIndex<<bucketMaxSizeShift + m.bucketsSize[newBucketIndex]
			m.items[index] = oldItems[i]
			m.bucketsSize[newBucketIndex]++
		}
	}
}

func (m *HashMap) insert(element item) bool {
	bucketIndex := element.hashValue & m.bucketMask
	bucketStartInd := bucketIndex << bucketMaxSizeShift
	bucketEndInd := bucketStartInd + m.bucketsSize[bucketIndex]

	for i := bucketStartInd; i != bucketEndInd; i++ {
		if m.items[i].hashValue == element.hashValue && m.items[i].key == element.key {
			m.items[i].value = element.value
			return true
		}
	}

	if m.bucketsSize[bucketIndex] != 8 {
		m.items[bucketEndInd] = element
		m.bucketsSize[bucketIndex]++
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
	for !m.insert(element) {
		m.rehash()
	}
}

func (m *HashMap) Get(key KeyType) (ValueType, bool) {
	hashValue := m.hash(key)
	bucketIndex := hashValue & m.bucketMask
	bucketStartInd := bucketIndex << bucketMaxSizeShift
	bucketEndInd := bucketStartInd + m.bucketsSize[bucketIndex]

	for i := bucketStartInd; i != bucketEndInd; i++ {
		if m.items[i].hashValue == hashValue && m.items[i].key == key {
			return m.items[i].value, true
		}
	}

	return nil, false
}

func (m *HashMap) Remove(key KeyType) {
	hashValue := m.hash(key)
	bucketIndex := hashValue & m.bucketMask
	bucketStartInd := bucketIndex << bucketMaxSizeShift
	bucketEndInd := bucketStartInd + m.bucketsSize[bucketIndex]

	for i := bucketStartInd; i != bucketEndInd; i++ {
		if m.items[i].hashValue == hashValue && m.items[i].key == key {
			m.items[i] = m.items[bucketEndInd-1]
			m.bucketsSize[bucketIndex]--
			return
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
