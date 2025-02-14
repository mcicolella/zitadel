package user

import (
	"google.golang.org/grpc"

	"github.com/zitadel/zitadel/internal/api/authz"
	"github.com/zitadel/zitadel/internal/api/grpc/server"
	"github.com/zitadel/zitadel/internal/command"
	"github.com/zitadel/zitadel/internal/crypto"
	"github.com/zitadel/zitadel/internal/query"
	user "github.com/zitadel/zitadel/pkg/grpc/user/v2alpha"
)

var _ user.UserServiceServer = (*Server)(nil)

type Server struct {
	user.UnimplementedUserServiceServer
	command     *command.Commands
	query       *query.Queries
	userCodeAlg crypto.EncryptionAlgorithm
}

type Config struct{}

func CreateServer(command *command.Commands, query *query.Queries, userCodeAlg crypto.EncryptionAlgorithm) *Server {
	return &Server{
		command:     command,
		query:       query,
		userCodeAlg: userCodeAlg,
	}
}

func (s *Server) RegisterServer(grpcServer *grpc.Server) {
	user.RegisterUserServiceServer(grpcServer, s)
}

func (s *Server) AppName() string {
	return user.UserService_ServiceDesc.ServiceName
}

func (s *Server) MethodPrefix() string {
	return user.UserService_ServiceDesc.ServiceName
}

func (s *Server) AuthMethods() authz.MethodMapping {
	return user.UserService_AuthMethods
}

func (s *Server) RegisterGateway() server.RegisterGatewayFunc {
	return user.RegisterUserServiceHandler
}
