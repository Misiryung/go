package services

import (
	"youtube-clone/file/pbs/file"
	// "youtube-clone/file/pbs/helper"

	"google.golang.org/grpc"
)

var (
	grpcServer *grpc.Server
)

func RegisterAllServices(server *grpc.Server) {
	grpcServer = server
	file.RegisterFileServiceServer(grpcServer, &fileServiceServer{})
}

// func newHttpError(m string, s int) *helper.HttpError {
// 	return &helper.HttpError{
// 		Message:    m,
// 		StatusCode: int32(s),
// 	}
// }
