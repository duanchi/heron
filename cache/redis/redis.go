package redis

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"go.heurd.com/heron-go/heron/abstract"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type RedisCache struct {
	abstract.Bean
	instance *redis.Client
}

func (this *RedisCache) Init() {

}

func (this *RedisCache) Instance(dsn *url.URL) {
	password, _ := dsn.User.Password()
	pathString := strings.Trim(dsn.Path, "/")
	if pathString == "" || pathString == "/" {
		pathString = "0"
	}
	path, _ := strconv.Atoi(dsn.Path)
	this.instance = redis.NewClient(&redis.Options{
		Addr:     dsn.Host,
		Password: password, // no password set
		DB:       path,  // use default DB
	})
	fmt.Printf("Redis %s connected at DB %d!\r\n", dsn.Host, path)
}

func (this *RedisCache) Get(key string) (value interface{}) {
	value, _ = this.instance.Get(key).Result()
	return
}

func (this *RedisCache) Has(key string) bool {

	has, _ := this.instance.Exists(key).Result()

	if has > 0 {
		return true
	}
	return false
}

func (this *RedisCache) Set(key string, value interface{}) {
	this.instance.Set(key, value, 0).Result()
}

func (this *RedisCache) SetWithTTL(key string, value interface{}, ttl int) {
	if ttl <= 0 {
		ttl = 0
	}

	this.instance.Set(key, value, time.Duration(ttl) * time.Second).Result()
}

func (this *RedisCache) Del(key string) {
	this.instance.Del(key).Result()
}

func (this *RedisCache) Flush() {
	this.instance.FlushDB().Result()
}
