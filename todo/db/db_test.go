package db

import (
	"io/ioutil"
	"testing"
)

func TestDB(t *testing.T) {
	rootDir, err := ioutil.TempDir("", "db-test-")
	t.Logf("rootDir = %s", rootDir)
	if err != nil {
		t.Fatal(err)
	}

	db, err := New(rootDir)
	if err != nil {
		t.Fatal(err)
	}

	keys := []string{"a", "b", "c", "d", "e"}
	for _, k := range keys {
		if err := db.Save(k, []byte(k)); err != nil {
			t.Fatal(err)
		}
	}

	dbKeys, err := db.Keys()
	if err != nil {
		t.Fatal(err)
	}

	if len(dbKeys) != len(keys) {
		t.Fatalf("wrong number of keys: got %d, expected %d", len(dbKeys), len(keys))
	}

	key := "b"
	data, err := db.Load(key)
	if err != nil {
		t.Fatal(err)
	}

	if string(data) != key {
		t.Fatalf("bad data for %s: %q", key, string(data))
	}

	if err := db.Delete(key); err != nil {
		t.Fatal(err)
	}

	dbKeys, err = db.Keys()
	if err != nil {
		t.Fatal(err)
	}

	if len(dbKeys) != len(keys)-1 {
		t.Fatalf("wrong number of keys after delete: got %d, expected %d", len(dbKeys), len(keys))
	}

}
