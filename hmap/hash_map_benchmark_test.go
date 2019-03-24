package hmap_test

import (
	"fmt"
	"testing"

	"github.com/ReanGD/go-algo/hmap"
)

func makeStringKeys(cnt int) []string {
	keys := make([]string, cnt)
	for i := 0; i != cnt; i++ {
		keys[i] = fmt.Sprintf("string%02d", i)
	}

	return keys
}

func BenchmarkRuntimeStringMap(b *testing.B) {
	keys := makeStringKeys(1024)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[string]bool)
		for keyInd := 0; keyInd != len(keys); keyInd++ {
			m[keys[keyInd]] = true
		}
	}
}

func BenchmarkHMapStringMap(b *testing.B) {
	keys := makeStringKeys(1024)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := hmap.New(hmap.HashString)
		for keyInd := 0; keyInd != len(keys); keyInd++ {
			m.Insert(keys[keyInd], true)
		}
	}
}
