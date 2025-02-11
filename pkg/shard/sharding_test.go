package shard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSharding(t *testing.T) {
	t.Run("sharding string", func(t *testing.T) {
		got := Sharding("ChIJHbyc6Z9eeUgR_G9U69yrDMI")
		assert.Equal(t, int64(22), got)
	})

	t.Run("sharding int64", func(t *testing.T) {
		got := Sharding(int64(659530501678518272))
		assert.Equal(t, int64(35), got)
	})
}
