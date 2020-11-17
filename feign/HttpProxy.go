package feign

import (
	"github.com/duanchi/heron/feign/reflec"
	"log"
)

type HttpProxy struct {

}

func (p *HttpProxy)Invoke(proxy *reflec.Proxy, method *reflec.Method, args []interface{})(error){
	serviceName := args[0].(string)
	service := services[serviceName]

	if service.Debug == "true" {
		log.Println("----rpc begin----")
	}
	args = append(args, service.Url)
	args = append(args, service.Token)
	err:= method.Invoke(args...)
	if service.Debug == "true" {
		log.Println("----rpc finish----")
	}
	return err
}
