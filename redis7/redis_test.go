package redis7

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/nuts-foundation/go-stoabs"
	"github.com/nuts-foundation/go-stoabs/kvtests"
	"testing"
)

func TestRedis(t *testing.T) {
	s := miniredis.RunT(t)
	defer s.Close()

	provider := func() (stoabs.KVStore, error) {
		return CreateRedisStore(&redis.Options{
			Addr: s.Addr(),
		})
	}

	kvtests.TestReadingAndWriting(t, provider)
	kvtests.TestRange(t, provider)
	kvtests.TestIterate(t, provider)
}
