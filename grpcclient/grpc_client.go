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

func init() {
	var err error
	grpcConn, err = grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
