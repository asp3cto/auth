package config

import (
	"net"
	"os"

	"errors"
)

const (
	grpcHostEnvName = "GRPC_HOST"
	grpcPortEnvName = "GRPC_PORT"
)

// Interface for grpc config
type GRPCConfig interface {
	Address() string
}

type grpcConfig struct {
	host string
	port string
}

// constructor for GRPCConfig
func NewGRPCConfig() (GRPCConfig, error) {
	host := os.Getenv(grpcHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("grpc host not found")
	}

	port := os.Getenv(grpcPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("grpc port not found")
	}

	return &grpcConfig{
		host: host,
		port: port,
	}, nil
}

// Returns address of grpc host
func (cfg *grpcConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
