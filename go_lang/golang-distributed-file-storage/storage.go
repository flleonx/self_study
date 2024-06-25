package main

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"strings"
)

type (
	PathTransformFunc func(key string) PathKey

	PathKey struct {
		PathName string
		Filename string
	}

	StoreOpts struct {
		// Root is the folder name of the root, containing all the folders/files
		// of the system.
		Root              string
		PathTransformFunc PathTransformFunc
	}

	Store struct {
		StoreOpts
	}
)

const defaultRootFolderName = "network"

func (p PathKey) FullPath() string {
	return fmt.Sprintf("%s/%s", p.PathName, p.Filename)
}

func (p PathKey) FirstPathChunk() string {
	pathChunks := strings.Split(p.PathName, "/")
	if len(pathChunks) == 0 {
		return ""
	}
	return pathChunks[0]
}

func CASPathTransformFunc(key string) PathKey {
	hash := sha1.Sum([]byte(key))
	hashStr := hex.EncodeToString(hash[:])

	blocksize := 5
	sliceLen := len(hashStr) / blocksize

	paths := make([]string, sliceLen)

	for i := 0; i < sliceLen; i++ {
		from, to := i*blocksize, (i*blocksize)+blocksize
		paths[i] = hashStr[from:to]
	}

	return PathKey{
		PathName: strings.Join(paths, "/"),
		Filename: hashStr,
	}
}

func DefaultPathTransformFunc(key string) PathKey {
	return PathKey{
		PathName: key,
		Filename: key,
	}
}

func NewStore(opts StoreOpts) *Store {
	if opts.PathTransformFunc == nil {
		opts.PathTransformFunc = DefaultPathTransformFunc
	}
	if len(opts.Root) == 0 {
		opts.Root = defaultRootFolderName
	}
	return &Store{
		StoreOpts: opts,
	}
}

func (s *Store) Has(key string) bool {
	pathKey := s.PathTransformFunc(key)

	fullPathWithRoot := fmt.Sprintf("%s/%s", s.Root, pathKey.FullPath())

	_, err := os.Stat(fullPathWithRoot)

	return !errors.Is(err, fs.ErrNotExist)
}

func (s *Store) Write(key string, r io.Reader) (int64, error) {
	return s.writeStream(key, r)
}

func (s *Store) Read(key string) (int64, io.Reader, error) {
	return s.readStream(key)
}

func (s *Store) Delete(key string) error {
	pathKey := s.PathTransformFunc(key)

	defer func() {
		log.Printf("deleted [%s] from disk", pathKey.FullPath())
	}()

	firstPathChunkWithRoot := fmt.Sprintf("%s/%s", s.Root, pathKey.FirstPathChunk())

	return os.RemoveAll(firstPathChunkWithRoot)
}

func (s *Store) Clear() error {
	return os.RemoveAll(s.Root)
}

func (s *Store) readStream(key string) (int64, io.ReadCloser, error) {
	pathKey := s.PathTransformFunc(key)
	fullPathWithRoot := fmt.Sprintf("%s/%s", s.Root, pathKey.FullPath())

	// fi, err := os.Stat(fullPathWithRoot)
	// if err != nil {
	//     return nil, err
	// }

	file, err := os.Open(fullPathWithRoot)
	if err != nil {
		return 0, nil, err
	}

	fi, err := file.Stat()
	if err != nil {
		return 0, nil, err
	}

	return fi.Size(), file, nil
}

func (s *Store) writeStream(key string, r io.Reader) (int64, error) {
	pathKey := s.PathTransformFunc(key)

	pathNameWithRoot := fmt.Sprintf("%s/%s", s.Root, pathKey.PathName)

	if err := os.MkdirAll(pathNameWithRoot, os.ModePerm); err != nil {
		return 0, err
	}

	fullPathWithRoot := fmt.Sprintf("%s/%s", s.Root, pathKey.FullPath())

	f, err := os.Create(fullPathWithRoot)
	if err != nil {
		return 0, err
	}

	n, err := io.Copy(f, r)
	if err != nil {
		log.Printf("error writing into disk %s", err)
		return 0, err
	}

	return n, nil
}
