package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/google/wire"
	"github.com/ijidan/jsocial/internal/pkg/config"
	"sync"
)

var (
	onceRd     sync.Once
	instanceRd redis.Conn
)

func GetRdInstance(cf *config.Redis) (redis.Conn, error) {
	onceRd.Do(func() {
		addr := fmt.Sprintf("%s:%d", cf.Host, cf.Port)
		instanceRd, _ = redis.Dial("tcp", addr)
	})
	return instanceRd, nil
}

var Provider = wire.NewSet(GetRdInstance)
