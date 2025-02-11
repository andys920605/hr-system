package shard

import (
	"encoding/binary"
	"hash/fnv"
)

const (
	Size = 100
)

func Sharding[T string | int64](in T) int64 {
	writing := make([]byte, 8)
	switch in := any(in).(type) {
	case string:
		writing = []byte(in)
	case int64:
		binary.BigEndian.PutUint64(writing, uint64(in))
	}

	hash := fnv.New32a()
	_, _ = hash.Write(writing)
	return int64(hash.Sum32() % Size)
}
