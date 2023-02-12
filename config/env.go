package config

import "os"

type envConfig struct {
	GrpcPort string
}

func (config *envConfig) load() {
	config.GrpcPort = os.Getenv("GRPC_PORT")
}
