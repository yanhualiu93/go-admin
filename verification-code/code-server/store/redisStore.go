package store

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"time"
)

func NewRedisStore(cli *redis.Client, prefix string, exxpirytime time.Duration) *redisStore {
	if prefix == "" {
		prefix = "code:"
	}
	if exxpirytime == 0 {
		exxpirytime = 120000 * time.Millisecond
	}
	return &redisStore{
		Prefix:     prefix,
		cli:        cli,
		Expirytime: exxpirytime,
	}
}

type redisStore struct {
	Prefix     string
	cli        *redis.Client
	Expirytime time.Duration
}

func (r *redisStore) Set(key, verifyValue string) {
	err := r.cli.Set(r.Prefix+key, verifyValue, r.Expirytime).Err()
	if err != nil {
		log.Error(err)
	}
}
func (r *redisStore) Get(key string, clear bool) (value string) {
	var err error
	value, err = r.cli.Get(r.Prefix + key).Result()
	if err != nil {
		log.Error(err)
		return
	}
	if clear {
		err = r.cli.Del(r.Prefix + key).Err()
		if err != nil {
			log.Error(err)
		}
	}
	return
}
