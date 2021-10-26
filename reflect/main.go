package main

import (
	"fmt"
	"reflect"
	"reflect/config"
	"reflect/model"
)

type Test struct {
}

func (t *Test) F1() {
	println("f1")
}

func main() {
	t := Test{}
	r := reflect.ValueOf(&t)
	f := r.MethodByName("F1")
	f.Call([]reflect.Value{})
	return
	conf := model.Config{
		ServerConf: model.ServerConfig{
			Ip:   "10.238.2.2",
			Port: 8080,
		},
		MysqlConf: model.MysqlConfig{
			Username: "root",
			Password: "admin",
			Database: "test",
			Host:     "192.168.10.10",
			Port:     8080,
			Timeout:  1.2,
		},
	}
	config.MarshalFile("config.toml", conf)
	var conf2 model.Config
	config.UnMarshalFile("config.toml", &conf2)
	fmt.Println(conf2)
}
