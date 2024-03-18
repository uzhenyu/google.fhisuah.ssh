package grpc

import (
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log"
	"math/rand"
	"net"
	"time"
	"zg5/lmonth/framework/consul"
	"zg5/lmonth/framework/vipers"
)

func GetGrpc(port int64, reg func(s *grpc.Server), fileName string) error {
	rand.Seed(time.Now().Unix())
	str := rand.Intn(9000) + 1000
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", str))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	err = vipers.GetViper(fileName)
	if err != nil {
		return err
	}
	name := viper.GetString("Nacos.DataID")
	host := viper.GetString("Nacos.Host")
	//ports := viper.GetInt("Nacos.Port")
	err = consul.NewClient(name, host, int64(str))
	if err != nil {
		return err
	}
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())
	reflection.Register(s)
	reg(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return nil
}
