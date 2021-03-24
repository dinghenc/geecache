package geecache

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestGetter(t *testing.T) {
	var f Getter = GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	})

	except := []byte("key")
	if v, _ := f.Get("key"); !reflect.DeepEqual(v, except) {
		t.Fatalf("callback failed")
	}
}

func TestGet(t *testing.T) {
	db := map[string]string{
		"Tom":  "630",
		"Jack": "589",
		"Sam":  "567",
	}

	loadCounts := make(map[string]int, len(db))
	g := NewGroup("scores", 2<<10, GetterFunc(func(key string) ([]byte, error) {
		log.Println("[SlowBD] search key: ", key)
		if v, ok := db[key]; ok {
			loadCounts[key] += 1
			return []byte(v), nil
		}
		return nil, fmt.Errorf("%s not exist", key)
	}))

	for k, v := range db {
		if view, err := g.Get(k); err != nil || view.String() != v {
			t.Fatalf("failed get value: %s", k)
		}
		if _, err := g.Get(k); err != nil || loadCounts[k] > 1 {
			t.Fatalf("cache miss %s", k)
		}
	}

	if view, err := g.Get("unkonwn"); err == nil {
		t.Fatalf("get the empty cache: %v", view)
	}
}
