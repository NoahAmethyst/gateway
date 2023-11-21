package spider_svc

import (
	"gateway/cluster/rpc"
	"gateway/protocol/spider_pb"
)

func SvcCli() spider_pb.SpiderServiceClient {
	return spider_pb.NewSpiderServiceClient(rpc.GetConn(rpc.CliSpider))
}
