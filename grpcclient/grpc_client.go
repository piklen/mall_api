package grpcclient

import (
	pb "github.com/piklen/pb/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"mall_api/pkg/log"
)

var (
	addr       = "127.0.0.1:8972"
	grpcConn   *grpc.ClientConn
	userClient pb.UserServiceClient
)

const (
	// 设置最大传输内容大小为50MB
	maxMsgSize = 50 * 1024 * 1024
)

func init() {
	var err error
	grpcConn, err = grpc.Dial(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		//设置最大传输请求大小
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxMsgSize), grpc.MaxCallSendMsgSize(maxMsgSize)),
	)
	if err != nil {
		log.LogrusObj.Infoln("Failed to connect to gRPC server: %v", err)
	}
	userClient = pb.NewUserServiceClient(grpcConn)
}

// GetUserClient 返回一个gRPC客户端实例
func GetUserClient() pb.UserServiceClient {
	return userClient
}

// CloseGrpcConn 关闭gRPC连接
func CloseGrpcConn() {
	if grpcConn != nil {
		grpcConn.Close()
	}
}
