package feign

import (
	"github.com/duanchi/heron/feign/reflec"
	"github.com/duanchi/heron/types/config"
)
var services = make(map[string]*ServiceToken)
var Engin *reflec.Proxy
var delimiter = ":"
func Init(feignConf config.Feign){
	hc := &httpClient{}
	hp := new(HttpProxy)
	Engin = reflec.New(hc, hp)
	for _,serviceConf := range feignConf.Services{
		if serviceConf.Enabled == true && serviceConf.Path != "" {
			service := new(ServiceToken)
			service.Url = serviceConf.Url
			service.TokenHeader = serviceConf.TokenHeader
			service.Debug = feignConf.Debug
			services[serviceConf.Name] = service
			go refreshToken(serviceConf, service, feignConf.Debug)
		}
	}
}





