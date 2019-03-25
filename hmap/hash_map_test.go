package hmap_test

import (
	"fmt"
	"testing"

	"github.com/ReanGD/go-algo/hmap"
)

func sKey(i int) string {
	return fmt.Sprintf("string%02d", i)
}

func TestBigItems(t *testing.T) {
	var keys [256]string
	var values [256]int
	for i := 0; i < 256; i++ {
		keys[i] = sKey(i)
		values[i] = i
	}

	m := hmap.New(hmap.HashString)
	for i := 0; i != len(keys); i++ {
		m.Insert(keys[i], values[i])
	}
	for i := 0; i != len(keys); i++ {
		value, ok := m.Get(keys[i])
		if !ok {
			t.Errorf("#%d: missing key: %v", i, keys[i])
		}
		if values[i] != value {
			t.Errorf("#%d: missing value: %v", i, values[i])
		}
	}
}

func TestInsert(t *testing.T) {
	m := hmap.New(hmap.HashString)

	m.Insert(sKey(1), 1)
	m.Insert(sKey(2), 1)
	m.Insert(sKey(2), 2)
	m.Insert(sKey(100), 100)

	value, ok := m.Get(sKey(1))
	if !ok {
		t.Errorf("Map does not contain key: %v", sKey(1))
	}
	if value != 1 {
		t.Errorf("Map contains incorrect value: %v", value)
	}

	value, ok = m.Get(sKey(2))
	if !ok {
		t.Errorf("Map does not contain key: %v", sKey(2))
	}
	if value != 2 {
		t.Errorf("Map contains incorrect value: %v", value)
	}

	value, ok = m.Get(sKey(100))
	if !ok {
		t.Errorf("Map does not contain key: %v", sKey(100))
	}
	if value != 100 {
		t.Errorf("Map contains incorrect value: %v", value)
	}

	_, ok = m.Get(sKey(99))
	if ok {
		t.Errorf("Map contains wrong key: %v", sKey(99))
	}
}

func TestInsertDublicate(t *testing.T) {
	m := hmap.New(hmap.HashString)
	key1 := sKey(1)
	value1 := 1
	m.Insert(key1, value1)

	value, ok := m.Get(key1)
	if !ok {
		t.Errorf("Map does not contain key: %v", key1)
	}
	if value != value1 {
		t.Errorf("Map contains incorrect value: %v", value)
	}

	key2 := sKey(1)
	value2 := 2
	m.Insert(key2, value2)

	value, ok = m.Get(key1)
	if !ok {
		t.Errorf("Map does not contain key: %v", key1)
	}
	if value != value2 {
		t.Errorf("Map contains incorrect value: %v", value)
	}

	value, ok = m.Get(key2)
	if !ok {
		t.Errorf("Map does not contain key: %v", key2)
	}
	if value != value2 {
		t.Errorf("Map contains incorrect value: %v", value)
	}
}
