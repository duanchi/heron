package http

import (
	json2 "encoding/json"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type Request struct {
	initialed bool
	error error

	method string
	url string
	baseUrl string
	queryString string
	header http.Header
	payload []byte
	formData interface{}
}

const (
	POST = "POST"
	GET = "GET"
	PUT = "PUT"
	DELETE = "DELETE"
	HEAD = "HEAD"
	PATCH = "PATCH"
	OPTIONS = "OPTIONS"
	TRACE = "TRACE"
)

func (this *Request) POST(url string) *Request {
	this.Instance()

	this.method = POST
	this.url = url

	return this
}

func (this *Request) New() *Request {
	this.initialed = true
	return this
}

func (this *Request) Instance() {
	if !this.initialed {
		this.New()
	}
}

func (this *Request) BaseUrl (url string) *Request {
	this.Instance()

	this.baseUrl = url
	return this
}

func (this *Request) Body (data []byte) *Request {
	this.Instance()

	this.payload = data

	return this
}

func (this *Request) Json (json interface{}) *Request {
	this.Instance()

	this.Header("Content-Type", "application/json")
	this.payload, this.error = json2.Marshal(json)

	return this
}

func (this *Request) Form (formData interface{}) *Request {
	this.Instance()

	switch reflect.TypeOf(formData).Kind() {
	case reflect.String:
		this.queryString = formData.(string)
	case reflect.Map:
		this.queryString = buildQueryString(formData)
	}

	return this
}

func (this *Request) Query (query interface{}) *Request {
	this.Instance()

	switch reflect.TypeOf(query).Kind() {
	case reflect.String:
		this.queryString = query.(string)
	case reflect.Map:
		var slice []string
		for k, v := range query.(map[string]string) {
			slice = append(slice, k + "=" + v)
		}
		this.queryString = strings.Join(slice, "&")
	}

	return this
}

func (this *Request) Header (key string, value string) *Request {
	this.Instance()

	this.header.Set(key, value)

	return this
}

func (this *Request) Headers (headers http.Header) *Request {
	this.Instance()

	for k,v := range headers {
		this.header[k] = v
	}

	return this
}

func (this *Request) Response () (response Response, err error) {
	this.Instance()



	return
}

func buildQueryString(data interface{}) string {
	var slice []string
	for k, v := range data.(map[string]interface{}) {

		switch reflect.TypeOf(v).Kind() {
		case reflect.Array:
			for _, sub := range v.([]interface{}) {
				slice = append(slice, k + "[]=" + parseValue(sub))
			}
		default:
			slice = append(slice, k + "=" + parseValue(v))
		}
	}
	return strings.Join(slice, "&")
}

func parseValue (value interface{}) string {
	switch reflect.TypeOf(value).Kind() {
	case reflect.String:
		return value.(string)
	case reflect.Int:
		return strconv.Itoa(value.(int))
	case reflect.Int64:
		return strconv.FormatInt(value.(int64), 10)
	case reflect.Bool:
		if value.(bool) {
			return "true"
		}
		return "false"
	case reflect.Float64:
		return strconv.FormatFloat(value.(float64), 'f', -1, 64)
	}

	return ""
}