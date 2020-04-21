package cache

import (
	"go.heurd.com/heron-go/heron/cache/memory"
	"go.heurd.com/heron-go/heron/cache/redis"
	"go.heurd.com/heron-go/heron/config"
	_interface "go.heurd.com/heron-go/heron/interface"
	"net/url"
	"reflect"
)

var engine _interface.CacheInterface

var CacheEngines map[string]reflect.Value

func Init() {
	if config.Get("Cache.Enabled").(bool) {

		CacheEngines["memory"] = reflect.ValueOf(&memory.MemoryCache{})
		CacheEngines["redis"] = reflect.ValueOf(&redis.RedisCache{})

		dsn := config.Get("Cache.Dsn").(string)
		dsnUrl, _ := url.Parse(dsn)

		if _, ok := CacheEngines[dsnUrl.Scheme]; !ok {
			dsnUrl, _ = url.Parse("memory://heron")
		}

		engine = CacheEngines[dsnUrl.Scheme].Interface().(_interface.CacheInterface)

		engine.Instance(dsnUrl)
	}
}

func Has(key string) bool {
	return engine.Has(key)
}

func Get(key string) interface{} {
	return engine.Get(key)
}

func Set(key string, value interface{}, ttl int) {
	engine.Set(key, value, ttl)
}

func Del (key string) {
	engine.Del(key)
}

func Flush() {
	engine.Flush()
}