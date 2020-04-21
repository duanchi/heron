package memory

import (
	"github.com/muesli/cache2go"
	"go.heurd.com/heron-go/heron/abstract"
	"net/url"
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
	return nil
}

func (this *MemoryCache) Has(key string) bool {
	return true
}

func (this *MemoryCache) Set(key string, value interface{}, ttl int) {

}

func (this *MemoryCache) Del(key string) {

}

func (this *MemoryCache) Flush() {

}

