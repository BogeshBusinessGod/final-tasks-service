package config

type GRPC struct {
	Port int `envconfig:"PORT" default:"50051"`
}
