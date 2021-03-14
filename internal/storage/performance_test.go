package storage

import (
	"github.com/ehsundar/ghibli_people/internal/storage/filereader"
	"github.com/ehsundar/ghibli_people/internal/storage/memcached"
	"testing"
)

func BenchmarkFileReaderWithoutCache(b *testing.B) {
	storage := filereader.New("/home/amir/Workspace/github.com/ehsundar/ghibli_people/static_people.json")

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := storage.GetAll()
			if err != nil {
				panic(err)
			}
		}
	})
}

func BenchmarkMemCachedFileReader(b *testing.B) {
	storage := filereader.New("/home/amir/Workspace/github.com/ehsundar/ghibli_people/static_people.json")
	cachedStorage := memcached.New(storage)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := cachedStorage.GetAll()
			if err != nil {
				panic(err)
			}
		}
	})
}
