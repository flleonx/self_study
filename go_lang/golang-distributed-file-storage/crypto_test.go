package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyEncrypt(t *testing.T) {
	src := bytes.NewReader([]byte("foo not bar"))
	dst := new(bytes.Buffer)
	key := newEncryptionKey()
	_, err := copyEncrypt(key, src, dst)
	assert.Nil(t, err)
	assert.Equal(t, len(dst.Bytes()), 27)
}

func TestCopyDecrypt(t *testing.T) {
	payload := "foo not bar"
	src := bytes.NewReader([]byte(payload))
	dst := new(bytes.Buffer)
	key := newEncryptionKey()
	_, err := copyEncrypt(key, src, dst)
	assert.Nil(t, err)

	out := new(bytes.Buffer)
	nn, err := copyDecrypt(key, dst, out)
	assert.Nil(t, err)
	assert.Equal(t, payload, out.String(), "decryotion failed")
	assert.Equal(t, nn, 16+len(payload))
}
