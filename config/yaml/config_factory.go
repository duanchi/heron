package yaml

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"reflect"
)

var configInstance interface{}

var configFile = "./config/application.yaml"

func GetMapConfig()(conf map[string]interface{}, err error){
	configFile, err := readFile()
	conf = make(map[string]interface{})
	err = yaml.Unmarshal(configFile, conf)
	if err != nil {
		log.Println(err)
	}
	return
}

func GetYamlConfig(config interface{})(conf interface{}, err error){
	configYaml, err := readFile()
	if err != nil {
		return
	}
	configuration := reflect.New(reflect.TypeOf(config).Elem())
	err = yaml.Unmarshal(configYaml, configuration)
	fmt.Println(configuration)
	if err != nil {
		log.Println(err)
		panic(err.Error())
	} else {
		// fmt.Printf("读取配置文件: %+v\n", conf)
		parseConfig(config)
		fmt.Printf("%+v", config)
		// fmt.Printf("更新配置参数: %+v\n", conf)
		conf = config
	}
	// configInstance = conf
	return
}

func readFile() (config []byte, err error){
	config, err = ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("yamlFile.Get err %v ", err)
	}
	return
}

func Get(key string) interface{} {
	return getRaw(key).Interface()
}