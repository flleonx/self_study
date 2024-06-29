package main

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathTransformFunc(t *testing.T) {
	key := "bestpicture"
	got := CASPathTransformFunc(key)

	wantPathName := "71056/ad8aa/24742/ea41e/a36fa/2e345/2a316/36e82"
	wantFileName := "71056ad8aa24742ea41ea36fa2e3452a31636e82"

	assert.Equal(t, wantPathName, got.PathName)
	assert.Equal(t, wantFileName, got.Filename)
}

func TestStoreDeleteKey(t *testing.T) {
	s := newStore()
	id := generateID()
	defer tearDown(t, s)

	key := "special"
	data := []byte("some jpg bytes")
	_, err := s.Write(id, key, bytes.NewReader(data))
	assert.Nil(t, err)

	ok := s.Has(id, key)
	assert.True(t, ok)

	err = s.Delete(id, key)
	assert.Nil(t, err)
}

func TestStore(t *testing.T) {
	s := newStore()
	id := generateID()
	defer tearDown(t, s)

	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("special_%d", i)

		data := []byte("some jpg bytes")
		_, err := s.Write(id, key, bytes.NewReader(data))
		assert.Nil(t, err)

		ok := s.Has(id, key)
		assert.True(t, ok)

		_, r, err := s.Read(id, key)
		assert.Nil(t, err)

		b, err := io.ReadAll(r)
		assert.Nil(t, err)
		assert.Equal(t, b, data)

		assert.Nil(t, s.Delete(id, key))

		assert.False(t, s.Has(id, key))
	}
}

func newStore() *Store {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	return NewStore(opts)
}

func tearDown(t *testing.T, s *Store) {
	if err := s.Clear(); err != nil {
		t.Error(err)
	}
}
