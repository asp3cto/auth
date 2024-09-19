package main

import (
	"context"
	"flag"

	// "fmt"
	"log"
	"net"

	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// sq "github.com/Masterminds/squirrel"

	"github.com/asp3cto/auth/internal/config"
	"github.com/asp3cto/auth/internal/converter"
	userRepo "github.com/asp3cto/auth/internal/repository/user"
	"github.com/asp3cto/auth/internal/service"
	userService "github.com/asp3cto/auth/internal/service/user"
	desc "github.com/asp3cto/auth/pkg/user_v1"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

type server struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

// Get ...
func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := s.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	log.Printf("%#v", user)

	return &desc.GetResponse{
		User: converter.ToUserFromService(user),
	}, nil
}

// Create ...
func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := s.userService.Create(ctx, converter.ToUserInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("Inserted user with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}

func main() {
	flag.Parse()
	ctx := context.Background()

	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	grpcConfig, err := config.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	pgConfig, err := config.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	userRepo := userRepo.NewRepository(pool)
	userServ := userService.NewService(userRepo)

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, &server{userService: userServ})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
