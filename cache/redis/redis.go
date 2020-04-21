package redis

import "go.heurd.com/heron-go/heron/abstract"

type RedisCache struct {
	abstract.Bean
}

func (this *RedisCache) Init() {

}

func (this *RedisCache) Get(key string) interface{} {
	return nil
}

func (this *RedisCache) Has(key string) bool {
	return true
}

func (this *RedisCache) Set(key string, value interface{}, ttl int) {

}

func (this *RedisCache) Del(key string) {

}

func (this *RedisCache) Flush() {

}
