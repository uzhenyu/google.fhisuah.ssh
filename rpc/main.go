package main

import (
	"flag"
	grpc2 "google.golang.org/grpc"
	"zg5/lmonth/framework/grpc"
	"zg5/lmonth/rpc/api"
	"zg5/lmonth/rpc/conts"
	"zg5/lmonth/rpc/model"
)

var (
	port = flag.Int("port", 8888, "wzy")
)

func main() {
	flag.Parse()
	err := model.AutoTable(conts.FilePath)
	if err != nil {
		return
	}

	err = grpc.GetGrpc(int64(*port), func(s *grpc2.Server) {
		api.Reg(s)
	}, conts.FilePath)
	if err != nil {
		return
	}
}
