package http

import (
	"bytes"
	json2 "encoding/json"
	"go.heurd.com/heron-go/heron/util/arrays"
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
	METHOD_POST = "POST"
	METHOD_GET = "GET"
	METHOD_PUT = "PUT"
	METHOD_DELETE = "DELETE"
	METHOD_HEAD = "HEAD"
	METHOD_PATCH = "PATCH"
	METHOD_OPTIONS = "OPTIONS"
	METHOD_TRACE = "TRACE"
)

func POST(url string) *Request {
	request := New()

	request.Method(METHOD_POST)
	request.Url(url)

	return &request
}

func GET(url string) *Request {
	request := New()

	request.Method(METHOD_GET)
	request.Url(url)

	return &request
}

func (this Request) PUT(url string) *Request {
	request := New()

	request.Method(METHOD_PUT)
	request.Url(url)

	return &request
}

func (this Request) DELETE(url string) *Request {
	request := New()

	request.Method(METHOD_DELETE)
	request.Url(url)

	return &request
}

func (this Request) HEAD(url string) *Request {
	request := New()

	request.Method(METHOD_HEAD)
	request.Url(url)

	return &request
}

func (this Request) PATCH(url string) *Request {
	request := New()

	request.Method(METHOD_PATCH)
	request.Url(url)

	return &request
}

func (this Request) OPTIONS(url string) *Request {
	request := New()

	request.Method(METHOD_OPTIONS)
	request.Url(url)

	return &request
}

func (this Request) TRACE(url string) *Request {
	request := New()

	request.Method(METHOD_TRACE)
	request.Url(url)

	return &request
}

func (this *Request) Url(url string) *Request {
	this.url = url

	return this
}

func (this *Request) Method(method string) *Request {
	if _, has := arrays.ContainsString([]string{ METHOD_GET, METHOD_POST, METHOD_PUT, METHOD_DELETE, METHOD_OPTIONS, METHOD_PATCH, METHOD_HEAD }, method); has {
		this.method = method
	} else {
		this.method = METHOD_GET
	}


	return this
}

func New() Request {
	instance := Request{
		initialed: true,
		header: http.Header{},
	}
	return instance
}

func (this *Request) BaseUrl (url string) *Request {
	this.baseUrl = url
	return this
}

func (this *Request) Body (data []byte) *Request {
	this.payload = data

	return this
}

func (this *Request) JSON (json interface{}) *Request {
	this.Header("Content-Type", "application/json")
	this.payload, this.error = json2.Marshal(json)

	return this
}

func (this *Request) Form (formData interface{}) *Request {
	switch reflect.TypeOf(formData).Kind() {
	case reflect.String:
		this.queryString = formData.(string)
	case reflect.Map:
		this.queryString = buildQueryString(formData)
	}

	return this
}

func (this *Request) Query (query interface{}) *Request {
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
	this.header.Set(key, value)

	return this
}

func (this *Request) Headers (headers http.Header) *Request {
	for k,v := range headers {
		this.header[k] = v
	}

	return this
}

func (this *Request) BearerToken (token string) *Request {

	this.Header("Authorization", token)

	return this
}

func (this *Request) Response () (response Response, err error) {
	if this.error != nil {
		return Response{}, this.error
	}

	url := this.baseUrl + this.url

	request, err := http.NewRequest(this.method, url, bytes.NewReader(this.payload))

	if err != nil {
		return
	}

	request.Header = this.header

	httpResponse, err := (&http.Client{}).Do(request)

	if err != nil {
		return
	}

	err = response.From(httpResponse)

	if err != nil { return }

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