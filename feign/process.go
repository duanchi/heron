package feign

import (
	"go.heurd.com/heron-go/heron/yconfig"
	"log"
	"strconv"
	"strings"
	"time"
)

func refreshToken(serviceConf yconfig.Service, service *ServiceToken, debug string){
	for ; ; {
		reqBody := make(map[string]string)
		userName := strings.Split(serviceConf.Username,delimiter)
		password := strings.Split(serviceConf.Password,delimiter)
		reqBody[userName[0]] = userName[1]
		reqBody[password[0]] = password[1]
		respBody := make(map[string]string)

		err := Engin.Post(serviceConf.Name, serviceConf.Path, reqBody,&respBody)
		if(err != nil){
			log.Printf("Get [%s] Token Error, Error: %s\n", serviceConf.Name, err)
		} else {
			token := respBody[serviceConf.TokenKey]
			//fmt.Printf("Get [%s] Token Success, Token:%s\n", serviceConf.Name, token)
			if debug == "true" || debug == "info" {
				log.Printf("Get [%s] Token Success, Token:%s\n",  serviceConf.Name, token)
			}
			service.Token = token
		}
		interval,err := strconv.Atoi(serviceConf.Interval)
		if(err != nil ){
			panic("rpc token 更新间隔参数[interval]格式设置错误，必须为正整数")
		}
		time.Sleep(time.Second * time.Duration(interval))
	}
}

