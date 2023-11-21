package main

import (
	"fmt"
	"gateway/cluster/rpc"
	"gateway/router"
	"gateway/utils/log"

	"os"
	"time"
)

func main() {
	time.Local = time.FixedZone("UTC", 0)

	// Initialize grpc client
	for _, rpcCli := range rpc.RpcCliList {
		rpc.InitGrpcCli(rpcCli)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	r := router.SetUpRouter()

	log.Info().Msgf("Http server start at 0.0.0.0:%s", httpPort)

	if err := r.Run(fmt.Sprintf(":%s", httpPort)); err != nil {
		panic(err)
	}
}
