package admin

import (
	"fmt"

	"hand/pkg/admin/pb"
	auth "hand/pkg/auth/pb"
	"hand/pkg/config"

	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AdminServiceClient
	Auth auth.AuthServiceClient
}

// func InitServiceClient(c *config.Config) pb.AuthServiceClient {
// 	// Create transport credentials (insecure for this example)
// 	creds := credentials.NewClientTLSFromCert(nil, "")
// 	// Dial the gRPC server with the transport credentials
// 	conn, err := grpc.Dial(c.AuthSvcUrl, grpc.WithTransportCredentials(creds))
// 	if err != nil {
// 		return nil
// 	}
// 	// Create and return the gRPC client
// 	return pb.NewAuthServiceClient(conn)
// }

func InitServiceClient(c *config.Config) pb.AdminServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.AdminSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAdminServiceClient(cc)
}
