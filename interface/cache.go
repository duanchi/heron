package _interface

import "net/url"

type CacheInterface interface {
	Instance(dsn *url.URL)
	Has(key string) bool
	Get(key string) interface{}
	Set(key string, value interface{}, ttl int)
	Flush()
	Del(key string)
}
