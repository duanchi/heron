package yconfig

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var configInstance interface{}

const configFile = "resource/application.yaml"

func GetMapConfig()(conf map[string]interface{}, err error){
	configFile := readFile()
	conf = make(map[string]interface{})
	err = yaml.Unmarshal(configFile, conf)
	if(err != nil ){
		log.Println(err)
	}
	return
}

func GetYamlConfig()(conf *Config, err error){
	configFile := readFile()
	err = yaml.Unmarshal(configFile, &conf)
	if(err != nil ){
		log.Println(err)
		panic(err.Error())
	} else {
		// fmt.Printf("读取配置文件: %+v\n", conf)
		parseConfig(conf)
		// fmt.Printf("更新配置参数: %+v\n", conf)
	}
	configInstance = conf
	return
}

func readFile() []byte{
	var err error;
	configFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("yamlFile.Get err %v ", err)
	}
	return configFile
}

func Get(key string) interface{} {
	return getRaw(key).Interface()
}