package memory

import (
	"github.com/muesli/cache2go"
	"go.heurd.com/heron-go/heron/abstract"
	"net/url"
	"time"
)

type MemoryCache struct {
	abstract.Bean
	instance *cache2go.CacheTable
}

func (this *MemoryCache) Init() {

}

func (this *MemoryCache) Instance(dsn *url.URL) {
	this.instance = cache2go.Cache(dsn.Hostname())
}

func (this *MemoryCache) Get(key string) interface{} {
	result, err := this.instance.Value(key)

	if err != nil {
		return result
	}

	return nil
}

func (this *MemoryCache) Has(key string) bool {
	_, err := this.instance.Value(key)

	if err != nil {
		return true
	}

	return false
}

func (this *MemoryCache) Set(key string, value interface{}, ttl int) {
	if ttl <= 0 {
		ttl = 0
	}

	this.instance.Add(key, time.Duration(ttl) * time.Second, &value)
}

func (this *MemoryCache) Del(key string) {
	this.instance.Delete(key)
}

func (this *MemoryCache) Flush() {
	this.instance.Flush()
}

