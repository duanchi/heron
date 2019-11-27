package feign

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type httpClient struct {
}

func (this *httpClient)Pget(serviceName string, path string, params map[string]string, obj interface{}, host string, token string) error {
	service := services[serviceName]
	url := makeUrl(host, path, params, service.Debug)
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	k,v := getTokenHeader(service.TokenHeader)
	request.Header.Set(k, v + " " + token)
	resp, err := client.Do(request)
	if(err != nil){
		return nil
	}
	return parseResponse(resp, obj, service.Debug)
}

func (this *httpClient)Ppost(serviceName string, path string, params map[string]string, body interface{}, obj interface{}, host string, token string) error {
	service := services[serviceName]
	url := makeUrl(host, path, params, service.Debug)
	client := &http.Client{}
	reqBody,err := json.Marshal(body)
	if err != nil {
		return err
	}
	if service.Debug == "true" {
		log.Println("request Body:", string(reqBody))
	}
	reqBody_new := bytes.NewBuffer([]byte(reqBody))
	request, _ := http.NewRequest("POST", url, reqBody_new)
	request.Header.Set("Content-type", "application/json")
	k,v := getTokenHeader(service.TokenHeader)
	request.Header.Set(k, v + " " + token)
	resp, err := client.Do(request)
	if(err != nil){
		return err
	}
	return parseResponse(resp, obj, service.Debug)

}

func (this *httpClient)Pput(serviceName string, path string, params map[string]string, body interface{}, obj interface{}, host string, token string) error {
	service := services[serviceName]
	url := makeUrl(host, path, params, service.Debug)
	client := &http.Client{}
	reqBody,err := json.Marshal(body)
	if err != nil {
		return err
	}
	if service.Debug == "true" {
		log.Println("request Body:", string(reqBody))
	}
	reqBody_new := bytes.NewBuffer([]byte(reqBody))
	request, _ := http.NewRequest("PUT", url, reqBody_new)
	request.Header.Set("Content-type", "application/json")
	k,v := getTokenHeader(service.TokenHeader)
	request.Header.Set(k, v + " " + token)
	resp, err := client.Do(request)
	if(err != nil){
		return err
	}
	return parseResponse(resp, obj, service.Debug)
}

func (this *httpClient)Pdelete(serviceName string, path string, params map[string]string, obj interface{}, host string, token string) error {
	service := services[serviceName]
	url := makeUrl(host, path, params, service.Debug)
	client := &http.Client{}
	request, _ := http.NewRequest("DELETE", url, nil)
	k,v := getTokenHeader(service.TokenHeader)
	request.Header.Set(k, v + " " + token)
	resp, err := client.Do(request)
	if(err != nil){
		return err
	}
	return parseResponse(resp, obj, service.Debug)
}

//拼接url
func makeUrl(host string, path string, params map[string]string, debug string) (url string) {
	var buf bytes.Buffer
	buf.WriteString(host)
	buf.WriteString(path)
	if params != nil {
		buf.WriteString("?")
		var i uint8 = 0
		for k,v := range params {
			if i != 0 {
				buf.WriteString("&")
			}
			buf.WriteString(k)
			buf.WriteString("=")
			buf.WriteString(v)
			i++
		}
	}
	url = buf.String()
	if debug == "true" {
		log.Println("request Url:", url)
	}
	return
}

//解析返回结果
func parseResponse(resp *http.Response, obj interface{}, debug string) error{
	if debug == "true" {
		log.Println("response Status: ", resp.StatusCode)
	}

	respBody, _ := ioutil.ReadAll(resp.Body)
	if debug == "true" {
		log.Println("response Body: ", string(respBody))
	}

	if resp.StatusCode == 200 {
		if err := json.Unmarshal(respBody, obj); err != nil {
			return err
		}
	} else {
		respMap := make(map[string]string)
		json.Unmarshal(respBody, &respMap)
		return errors.New(respMap["message"])
	}
	return nil
}

func getTokenHeader(tokenHeader string) (k string, v string){
	tokenHeaders := strings.Split(tokenHeader, delimiter)
	if len(tokenHeaders) == 2 {
		k = tokenHeaders[0]
		v = tokenHeaders[1]
	} else {
		k = tokenHeaders[0]
		v = ""
	}
	return
}