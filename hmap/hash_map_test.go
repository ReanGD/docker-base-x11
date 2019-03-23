package hmap_test

import (
	"fmt"
	"testing"

	"github.com/ReanGD/go-algo/hmap"
)

func TestBigItems(t *testing.T) {
	var keys [256]string
	var values [256]int
	for i := 0; i < 256; i++ {
		keys[i] = fmt.Sprintf("string%02d", i)
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
