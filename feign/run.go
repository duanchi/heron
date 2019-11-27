package feign

import (
	"go.heurd.com/heron-go/heron/yconfig"
	"go.heurd.com/heron-go/heron/feign/reflec"
)
var services = make(map[string]*ServiceToken)
var Engin *reflec.Proxy
var delimiter = ":"
func Init(feignConf yconfig.Feign){
	hc := &httpClient{}
	hp := new(HttpProxy)
	Engin = reflec.New(hc, hp)
	for _,serviceConf := range feignConf.Service{
		if serviceConf.Enabled == "true" && serviceConf.Path != "" {
			service := new(ServiceToken)
			service.Url = serviceConf.Url
			service.TokenHeader = serviceConf.TokenHeader
			service.Debug = feignConf.Debug
			services[serviceConf.Name] = service
			go refreshToken(serviceConf, service, feignConf.Debug)
		}
	}
}





