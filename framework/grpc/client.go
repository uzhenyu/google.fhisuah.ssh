package grpc

import (
	"encoding/json"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"zg5/lmonth/framework/nacos"
	"zg5/lmonth/framework/vipers"
)

func Client(fileName string) (*grpc.ClientConn, error) {
	err := vipers.GetViper(fileName)
	if err != nil {
		return nil, err
	}
	var consul struct {
		Consul struct {
			Name string `json:"Name"`
			Host string `json:"Host"`
			Port string `json:"Port"`
		} `json:"Consul"`
	}
	config, err := nacos.GetConfig(fileName)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(config), &consul)
	if err != nil {
		return nil, err
	}
	return grpc.Dial(fmt.Sprintf("%v://%v:%v/", consul.Consul.Name, consul.Consul.Host, consul.Consul.Port)+
		viper.GetString("Nacos.DataId")+"?wait=14s", grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"LoadBalancingPolicy": "round_robin"}`))
}
