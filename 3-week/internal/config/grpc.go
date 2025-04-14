package config

import (
	"net"
	// "github.com/pkg/errors"
)

const (
	grpcHostEnvName = "GRPC_HOST"
	grpcPortEnvName = "GRPC_PORT"
)

type GRPCConfig interface {
	Address() string
}

type grpcConfig struct {
	host string
	port string
}

func NewGRPCConfig() (GRPCConfig, error) {
	// host := os.Getenv(grpcHostEnvName)
	// if len(host) == 0 {
	// 	// os.Setenv(grpcHostEnvName, "localhost")
	// 	// fmt.Println("Default grpc host")
	// 	return nil, errors.New("grpc host not found")
	// }

	// port := os.Getenv(grpcPortEnvName)
	// if len(port) == 0 {
	// 	// os.Setenv(grpcPortEnvName, "50051")
	// 	// fmt.Println("Default grpc port")
	// 	return nil, errors.New("grpc port not found")
	// }

	return &grpcConfig{
		// host: host,
		// port: port,
		host: "localhost",
		port: "50051",
	}, nil
}

func (cfg *grpcConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
