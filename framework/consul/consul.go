package consul

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
)

//func NewClient(name, host string, port int64) error {
//	c, err := api.NewClient(api.DefaultConfig())
//	if err != nil {
//		return err
//	}
//	err = c.Agent().ServiceRegister(&api.AgentServiceRegistration{
//		ID:      uuid.New().String(),
//		Name:    name,
//		Tags:    []string{"GRPC"},
//		Port:    int(port),
//		Address: host,
//		Check: &api.AgentServiceCheck{
//			Interval:                       "5s",
//			Timeout:                        "5s",
//			GRPC:                           fmt.Sprintf("%v:%v", host, port),
//			DeregisterCriticalServiceAfter: "10s",
//		},
//	})
//	if err != nil {
//		return err
//	}
//	return nil
//}

func NewClient(name string, host string, port int64) error {
	c, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return err
	}
	err = c.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      uuid.New().String(),
		Name:    name,
		Tags:    []string{"GRPc"},
		Port:    int(port),
		Address: host,
		Check: &api.AgentServiceCheck{
			Interval:                       "5s",
			Timeout:                        "5s",
			GRPC:                           fmt.Sprintf("%v:%v", host, port),
			DeregisterCriticalServiceAfter: "10s",
		},
	})
	if err != nil {
		return err
	}
	return nil
}
